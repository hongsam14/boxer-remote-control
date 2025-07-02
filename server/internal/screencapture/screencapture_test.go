package screencapture_test

import (
	"fmt"
	"image/png"
	"os"
	"path/filepath"
	"testing"
	"time"

	"github.com/hongsam14/boxer-remote-control/server/internal/screencapture"
)

func TestNewScreencapture(t *testing.T) {
	sc, err := screencapture.NewScreencapture(3)
	if err != nil {
		t.Fatalf("Failed to create new screencapture: %v", err)
		return
	}
	err = sc.Start()
	if err != nil {
		t.Fatalf("Failed to start screencapture: %v", err)
		return
	}

	go func() {
		var i int = 0
		for frame := range sc.FrameChan() {
			if len(frame.Pix) == 0 {
				t.Logf("Received empty frame from screencapture")
				continue
			}
			// save frame to jpeg file for testing
			filename := filepath.Join(".", fmt.Sprintf("test_frame_%v.png", i))
			i++
			// encode to PNG format
			t.Logf("Saving frame to file %s", filename)
			file, err := os.Create(filename)
			if err != nil {
				t.Fatalf("Failed to create file %s: %v", filename, err)
				return
			}
			err = png.Encode(file, frame)
			if err != nil {
				t.Fatalf("Failed to encode frame to PNG: %v", err)
				file.Close()
				return
			}
			file.Close()
		}
	}()

	go func() {
		// timer thread to stop screencapture after 5 seconds
		time.Sleep(2 * time.Second)
		t.Logf("Screencapture stopped after 2 second")
		err = sc.Stop()
		t.Logf("Screencapture stopped: %v", err)
		if err != nil {
			t.Fatalf("Failed to stop screencapture: %v", err)
			return
		}
	}()

	// wait for screencapture to finish
	t.Logf("Waiting for screencapture to finish...")
	err = sc.Wait()
	if err != nil {
		t.Fatalf("Failed to wait for screencapture: %v", err)
		return
	}
	t.Logf("Screencapture finished successfully")
}
