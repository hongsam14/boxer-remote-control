package screencapture_test

import (
	"fmt"
	"os"
	"path/filepath"
	"testing"
	"time"

	"github.com/hongsam14/boxer-remote-control/server/internal/screencapture"
)

func TestNewScreencapture(t *testing.T) {
	sc, err := screencapture.NewScreencapture()
	if err != nil {
		t.Fatalf("Failed to create new screencapture: %v", err)
		return
	}
	err = sc.Start()
	if err != nil {
		t.Fatalf("Failed to start screencapture: %v", err)
		return
	}
	func() {
		var i int = 0
		for frame := range sc.FrameChan() {
			fmt.Fprintf(os.Stderr, "Saving frame %d to %s", len(frame), filepath.Join(".", fmt.Sprintf("test_frame_%v.jpg", i)))
			if len(frame) == 0 {
				t.Fatalf("Received empty frame from screencapture")
				return
			}
			// save frame to jpeg file for testing
			filename := filepath.Join(".", fmt.Sprintf("test_frame_%v.jpg", i))
			i++
			err := os.WriteFile(filename, frame, 0644)
			if err != nil {
				t.Fatalf("Failed to write frame to file %s: %v", filename, err)
				return
			}
		}
	}()

	func() {
		// timer thread to stop screencapture after 5 seconds
		time.Sleep(5 * time.Second)
		err = sc.Stop()
		if err != nil {
			t.Fatalf("Failed to stop screencapture: %v", err)
			return
		}
	}()

	_, err = sc.Wait()
	if err != nil {
		t.Fatalf("Failed to wait for screencapture: %v", err)
		return
	}
}
