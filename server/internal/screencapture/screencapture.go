package screencapture

import (
	"bufio"
	"bytes"
	"context"
	"fmt"
	"os"
	"runtime"
	"strings"

	berror "github.com/hongsam14/boxer-remote-control/error"
	"github.com/hongsam14/boxer-remote-control/server/internal/exec"
)

const (
	_FFMPEG_BIN = "ffmpeg"
	_LINUX      = "-f x11grab -video_size 1280x720 -i :0.0 -f mjpeg pipe:1 -q:v 5"
	_WINDOWS    = "-f gdigrab -framerate 15 -i desktop -f mjpeg pipe:1 -q:v 5"
	_DARWIN     = "-f avfoundation -framerate 15 -i 1 -f mjpeg pipe:1 -q:v 5"
)

type Screencapture interface {
	Start() error
	Wait() (int, error)
	Stop() error
	FrameChan() <-chan []byte
}

type screencapture struct {
	ctx    context.Context
	cancel context.CancelFunc
	// promise to run the screencapture command
	promise exec.Promise
	// frame reader
	reader *bufio.Reader
	// pipe
	framePipe *bytes.Buffer
	// output channel for frames
	frameChan chan []byte
}

func (sc *screencapture) FrameChan() <-chan []byte {
	return sc.frameChan
}

func NewScreencapture() (Screencapture, error) {
	var err error

	newSC := new(screencapture)

	// create a context for the screencapture
	newSC.ctx, newSC.cancel = context.WithCancel(context.Background())

	// create a channel to send frames
	newSC.frameChan = make(chan []byte, 10) // buffer size of 10 frames

	// create a pipe to get output from ffmpeg
	// input buffer 1280 * 720 * 3 + 2 bytes (RGB)
	newSC.framePipe = bytes.NewBuffer(make([]byte, 0, 1280*720*3+2))
	if err != nil {
		return nil, berror.BoxerError{
			Code:   berror.SystemError,
			Msg:    "error while NewScreencapture",
			Origin: err,
		}
	}
	// create a reader for the frame output
	newSC.reader = bufio.NewReader(newSC.framePipe)

	newSC.promise = nil

	return newSC, nil
}

func (sc *screencapture) Start() error {
	var (
		err error
	)

	// create cmdline for ffmpeg
	switch runtime.GOOS {
	case "linux":
		sc.promise, err = exec.Run(
			os.Stdin,
			sc.framePipe,
			_FFMPEG_BIN,
			strings.Split(_LINUX, " ")...)
	case "darwin":
		sc.promise, err = exec.Run(
			os.Stdin,
			sc.framePipe,
			_FFMPEG_BIN,
			strings.Split(_DARWIN, " ")...)
	case "windows":
		sc.promise, err = exec.Run(
			os.Stdin,
			sc.framePipe,
			_FFMPEG_BIN,
			strings.Split(_WINDOWS, " ")...)
	default:
		return berror.BoxerError{
			Code:   berror.InternalError,
			Msg:    "error in start function",
			Origin: fmt.Errorf("unsupported OS: %s", runtime.GOOS),
		}
	}
	if err != nil {
		return berror.BoxerError{
			Code:   berror.SystemError,
			Msg:    "error while starting screencapture",
			Origin: err,
		}
	}
	// start the scanner thread to read frames
	go sc.scannerThread()
	return nil
}

func (sc *screencapture) Wait() (int, error) {
	if sc.promise == nil {
		return 0, berror.BoxerError{
			Code:   berror.InvalidState,
			Msg:    "error while Wait() function",
			Origin: fmt.Errorf("screencapture promise is not initialized"),
		}
	}

	exitCode, err := sc.promise.Wait()
	if err != nil {
		return exitCode, berror.BoxerError{
			Code:   berror.SystemError,
			Msg:    "error while Wait() function",
			Origin: err,
		}
	}
	return exitCode, nil
}

func (sc *screencapture) Stop() error {
	if sc.promise == nil {
		return berror.BoxerError{
			Code:   berror.InvalidState,
			Msg:    "error while Stop() function",
			Origin: fmt.Errorf("screencapture promise is not initialized"),
		}
	}

	// cancel the promise
	if err := sc.promise.Cancel(); err != nil {
		return berror.BoxerError{
			Code:   berror.SystemError,
			Msg:    "error while Stop() function",
			Origin: err,
		}
	}

	// close the pipes
	close(sc.frameChan)

	// cancel the context
	sc.cancel()

	return nil
}

func (sc *screencapture) scannerThread() {

	var (
		buf       bytes.Buffer
		b, b1, b2 byte
		frame     []byte
		err       error
	)

	for {
		select {
		case <-sc.ctx.Done():
			// if the context is done, we stop the scanner thread.
			// this is to ensure that we don't leak goroutines
			return
		default:
			// read frame from the frame pipe buffer
			// find jpeg frame starting with 0xFFD8
			b1, err = sc.reader.ReadByte()
			// check if we can read the next byte
			// if not, we will break the loop
			if err != nil {
				// if we reach the end of the stream, break
				if err.Error() == "EOF" {
					return
				}
				return
			}
			if b1 != 0xFF {
				continue
			}
			b2, err = sc.reader.ReadByte()
			// check if we can read the next byte
			// if not, we will break the loop
			if err != nil {
				// if we reach the end of the stream, break
				if err.Error() == "EOF" {
					return
				}
				return
			}
			if b2 != 0xD8 {
				continue
			}
			// reset the buffer before writing the frame
			buf.Reset()

			// write the first two bytes to the buffer
			// which are the start of the JPEG frame
			// 0xFFD8 is the JPEG SOI (Start of Image) marker
			// and we need to include it in the frame
			buf.Write([]byte{b1, b2})

			for {
				b, err = sc.reader.ReadByte()
				if err != nil {
					// if we reach the end of the stream, break
					if err.Error() == "EOF" {
						if buf.Len() > 0 {
							// send the last frame if it is not empty
							sc.frameChan <- buf.Bytes()
						}
						return
					}
					return
				}
				buf.WriteByte(b)

				if buf.Len() > 2 &&
					buf.Bytes()[buf.Len()-2] == 0xFF &&
					buf.Bytes()[buf.Len()-1] == 0xD9 {
					break
				}
			}
			// create a frame from the buffer
			frame = buf.Bytes()
			// send the frame to the channel
			sc.frameChan <- frame
		}
	}
}
