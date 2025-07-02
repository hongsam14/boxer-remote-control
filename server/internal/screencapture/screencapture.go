package screencapture

import (
	"context"
	"fmt"
	"image"
	"io"
	"os"
	"time"

	berror "github.com/hongsam14/boxer-remote-control/error"
	"github.com/kbinani/screenshot"
	"golang.org/x/sync/errgroup"
)

type Screencapture interface {
	Start() error
	Wait() error
	Stop() error
	FrameChan() <-chan *image.RGBA
}

type screencapture struct {
	ctx    context.Context
	cancel context.CancelFunc
	// scanner context cancel
	scannerCtx    context.Context
	scannerCancel context.CancelFunc
	// errgroup
	captureGrp *errgroup.Group
	// screen infos
	numActivateScreens int
	bounds             image.Rectangle
	// config
	frameRate uint
	// pipe
	frameReader *io.PipeReader
	frameWriter *io.PipeWriter
	// output channel for frames
	frameChan chan *image.RGBA
}

func (sc *screencapture) FrameChan() <-chan *image.RGBA {
	return sc.frameChan
}

func NewScreencapture(frameRate uint) (Screencapture, error) {
	newSC := new(screencapture)

	// create a context for the screencapture
	newSC.ctx, newSC.cancel = context.WithCancel(context.Background())
	// create a context for the scanner
	newSC.scannerCtx, newSC.scannerCancel = context.WithCancel(context.Background())

	// create a channel to send frames
	newSC.frameChan = make(chan *image.RGBA, 10) // buffer size of 10 frames

	// get screen information
	newSC.numActivateScreens = screenshot.NumActiveDisplays()
	if newSC.numActivateScreens <= 0 {
		return nil, berror.BoxerError{
			Code:   berror.InvalidState,
			Msg:    "error while NewScreencapture",
			Origin: fmt.Errorf("no active screens found"),
		}
	}
	newSC.bounds = screenshot.GetDisplayBounds(0)
	if newSC.bounds.Empty() {
		return nil, berror.BoxerError{
			Code:   berror.InvalidState,
			Msg:    "error while NewScreencapture",
			Origin: fmt.Errorf("failed to get display bounds for screen 0"),
		}
	}

	fmt.Fprintf(os.Stderr, "Screencapture: %d active screens found, bounds: %v\n", newSC.numActivateScreens, newSC.bounds)

	// create a pipe to get output from ffmpeg
	newSC.frameReader, newSC.frameWriter = io.Pipe()
	// config
	newSC.frameRate = frameRate
	// create errgroup for captureThread
	newSC.captureGrp, _ = errgroup.WithContext(newSC.ctx)
	return newSC, nil
}

func (sc *screencapture) Start() error {
	// start capture thread
	sc.captureGrp.Go(sc.captureThread)
	// start the scanner thread to read frames
	sc.captureGrp.Go(sc.scannerThread)
	return nil
}

func (sc *screencapture) Wait() error {
	err := sc.captureGrp.Wait()
	if err != nil {
		return berror.BoxerError{
			Code:   berror.SystemError,
			Msg:    "error while Wait() function",
			Origin: err,
		}
	}
	// close the frame channel
	close(sc.frameChan)
	return nil
}

func (sc *screencapture) Stop() error {
	// cancel the context
	sc.cancel()
	fmt.Fprintf(os.Stderr, "Screencapture: capture thread called\n")
	// sc.scannerCancel()
	// fmt.Fprintf(os.Stderr, "Screencapture: scanner thread stopped\n")
	return nil
}

func (sc *screencapture) captureThread() error {
	var (
		err error
		img *image.RGBA
	)

	// calculate the frame rate in miliseconds
	frameRateMs := 1000 / int(sc.frameRate)

	// create a ticker to capture frames at the specified frame rate
	ticker := time.NewTicker(time.Duration(frameRateMs) * time.Millisecond)
	defer ticker.Stop()
	defer sc.frameWriter.Close()

	for {
		select {
		case <-sc.ctx.Done():
			// if the context is done, we stop the capture thread.
			// this is to ensure that we don't leak goroutines
			// close the pipes
			fmt.Fprintf(os.Stderr, "Screencapture: capture thread stopped\n")
			return nil
		case <-ticker.C:
			// capture the screen
			img, err = screenshot.CaptureRect(sc.bounds)
			if err != nil {
				// create a channel to send framesp()
				return berror.BoxerError{
					Code:   berror.InternalError,
					Origin: err,
					Msg:    "error in captureThread function",
				}
			}
			if len(img.Pix) <= 0 {
				// if the captured image is empty, we skip this frame
				return berror.BoxerError{
					Code:   berror.InternalError,
					Origin: fmt.Errorf("captured image is empty"),
					Msg:    "error in captureThread function",
				}
			}
			// write the captured image to the pipe
			_, err = sc.frameWriter.Write(img.Pix)
			if err != nil {
				return berror.BoxerError{
					Code:   berror.InternalError,
					Origin: err,
					Msg:    "error in captureThread function",
				}
			}
			fmt.Fprintf(os.Stderr, "Screencapture: captured frame at %v\n", time.Now())
		}
	}
}

func (sc *screencapture) scannerThread() error {

	var (
		l     int
		frame []byte
		img   *image.RGBA
		err   error
	)

	frame = make([]byte, sc.bounds.Dx()*sc.bounds.Dy()*4)
	defer sc.frameReader.Close()
	for {
		select {
		case <-sc.scannerCtx.Done():
			// if the context is done, we stop the scanner thread.
			// this is to ensure that we don't leak goroutines
			fmt.Fprintf(os.Stderr, "Screencapture: scanner thread stopped\n")
			return nil
		default:
			// read frame from the frame pipe buffer
			l, err = io.ReadFull(sc.frameReader, frame)
			if err != nil {
				if err == io.EOF {
					// if we reach EOF, we stop the scanner thread
					fmt.Fprintf(os.Stderr, "Screencapture: scanner thread reached EOF %v\n", l)
					return nil
				}
				return berror.BoxerError{
					Code:   berror.InternalError,
					Origin: err,
					Msg:    "error in scannerThread",
				}
			}
			if l < 0 {
				// if we read 0 bytes, we continue to the next iteration
				fmt.Fprintf(os.Stderr, "Screencapture: scanner thread read 0 bytes\n")
				continue
			}
			// send the frame to the channel
			img = &image.RGBA{
				Pix:    frame,
				Stride: sc.bounds.Dx() * 4,
				Rect:   sc.bounds,
			}
			sc.frameChan <- img
		}
	}
}
