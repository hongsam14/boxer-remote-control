package exec

import (
	"context"
	"fmt"
	"io"
	"os"
	"os/exec"
	"sync/atomic"
	"syscall"
	"time"

	berror "github.com/hongsam14/boxer-remote-control/error"
	"golang.org/x/sync/errgroup"
)

// # Promise
//
// Promise calls an external program as a subprocess and manages it asynchronously.
// It provides a way to wait for the subprocess to finish or cancel it.
type Promise interface {
	PID() int
	IsExecuted() bool
	Wait() (exitCode int, err error)
	Cancel() (err error)
}

type promise struct {
	cmd     exec.Cmd
	eg      *errgroup.Group
	waitCnt int32
}

func (p *promise) PID() int {
	if p.cmd.Process == nil {
		return 0
	}
	return p.cmd.Process.Pid
}

func (p *promise) IsExecuted() bool {
	return p.cmd.Process != nil || atomic.LoadInt32(&p.waitCnt) >= 0
}

// # Run
//
// Run starts a new subprocess with the given commandline and returns a Promise.
func Run(inStream *os.File, outStream io.Writer, arg0 string, args ...string) (Promise, error) {
	prom := new(promise)
	// set conditional variable to -1
	atomic.StoreInt32(&prom.waitCnt, -1)
	// set the commandline
	prom.cmd = *exec.Command(arg0, args...)
	prom.cmd.Stdin = inStream
	prom.cmd.Stdout = outStream
	prom.cmd.Stderr = os.Stderr

	// generate process group
	prom.cmd.SysProcAttr = &syscall.SysProcAttr{Setpgid: true}

	// set the error for run commandline goroutine
	err := prom.cmd.Start()
	if err != nil {
		return nil, berror.BoxerError{
			Code:   berror.SystemError,
			Origin: err,
			Msg:    fmt.Sprintf("error while execute %s %v promise.Start()", prom.cmd.Path, prom.cmd.Args),
		}
	}
	// set the conditional variable to 0
	atomic.StoreInt32(&prom.waitCnt, 0)
	return prom, nil
}

// # Wait
//
// Wait waits for the subprocess to finish and returns an error if the subprocess is not successful.
// Wait should be called only once and it will return an error if it is called multiple times.
// Also, it will return an error if it is called before Run.
// ExitCode will be returned as 0 if the subprocess is successful.
// Also ExitCode will be returned as -1 if subprocess is killed by signal.
// So the caller should check the ExitCode and Error to determine if the subprocess is successful or not.
// Because ExitCode -1 means the subprocess is killed by signal, and it is not an error.
func (p *promise) Wait() (exitCode int, err error) {
	// if conditional variable is not set to 0, return error
	val := atomic.LoadInt32(&p.waitCnt)
	if val != 0 {
		// errorcode 1 means general error
		return 1, berror.BoxerError{
			Code:   berror.InvalidOperation,
			Msg:    fmt.Sprintf("error while execute %s %v promise.Wait()", p.cmd.Path, p.cmd.Args),
			Origin: fmt.Errorf("wait is called while promise is not initialized [%d]", atomic.LoadInt32(&p.waitCnt)),
		}
	}
	// set the conditional variable to 1
	atomic.StoreInt32(&p.waitCnt, 1)
	state, err := p.cmd.Process.Wait()
	if err != nil {
		// errorcode 130 means fatal error
		return 130, berror.BoxerError{
			Code:   berror.SystemError,
			Msg:    fmt.Sprintf("error while execute %s %v promise.Wait()", p.cmd.Path, p.cmd.Args),
			Origin: err,
		}
	}
	// set the conditional variable to 0 because the process is finished
	atomic.StoreInt32(&p.waitCnt, 0)
	exitCode = state.ExitCode()
	// release the process resource to prevent zombie process
	err = p.cmd.Process.Release()
	if exitCode < 0 && err != nil {
		return exitCode, berror.BoxerError{
			Code:   berror.SystemError,
			Msg:    fmt.Sprintf("error while execute %s %v promise.Wait() in Process.Release()", p.cmd.Path, p.cmd.Args),
			Origin: err,
		}
	}
	return exitCode, nil
}

// # Cancel
//
// Cancel sends a signal to the subprocess to kill it.
// It will return an error if the subprocess is not initialized.
// It will be blocked until the subprocess is killed and release the resource.
func (p *promise) Cancel() (err error) {
	if atomic.LoadInt32(&p.waitCnt) < 0 {
		return berror.BoxerError{
			Code:   berror.InvalidOperation,
			Msg:    fmt.Sprintf("error while execute %s %v promise.Cancel()", p.cmd.Path, p.cmd.Args),
			Origin: fmt.Errorf("cancel is called while promise is not initialized [%d]", atomic.LoadInt32(&p.waitCnt)),
		}
	}
	// create timeout context for killing subprocesskiller with timeout.
	// if the subprocesskiller is not killed in the timeout, the subprocesskiller will be killed by the parent process.
	subProcessKillerCtx, subProcessKillerCancel := context.WithDeadline(context.TODO(), time.Now().Add(5*time.Second))
	defer subProcessKillerCancel()
	p.eg, _ = errgroup.WithContext(subProcessKillerCtx)
	// send signal to subprocesskiller
	p.eg.Go(func() error {
		return p.subProcessKiller()
	})

	err = p.eg.Wait()
	if err != nil {
		return berror.BoxerError{
			Code:   berror.SystemError,
			Msg:    fmt.Sprintf("error while execute %s %v promise.Cancel()", p.cmd.Path, p.cmd.Args),
			Origin: err,
		}
	}
	return nil
}

func (p *promise) subProcessKiller() error {
	// kill proc group
	//
	// (DEPRECATED) kill process group
	//
	// pgid, err := syscall.Getpgid(p.cmd.Process.Pid)
	// if err == nil {
	// 	fmt.Fprintf(os.Stderr, "Killing process group %d\n", pgid)
	// 	syscall.Kill(-pgid, syscall.SIGKILL)
	// }
	//
	// kill process by SIGKILL signal
	//
	// p.cmd.Process.Kill()

	err := p.cmd.Process.Signal(syscall.SIGINT)
	if err != nil {
		// if the process is not killed by SIGINT, send SIGTERM
		return p.cmd.Process.Signal(syscall.SIGTERM)
	}
	// to prevent process is not killed by SIGINT, send SIGTERM
	return p.cmd.Process.Signal(syscall.SIGTERM)

	// wait to prevent zombie process (DEPRECATED)
	// p.cmd.Wait()
	// err := p.cmd.Process.Release()
	// if err != nil {
	// 	return perror.PolvoGeneralError{
	// 		Code:   perror.SystemError,
	// 		Msg:    fmt.Sprintf("error while execute %s %v in subProcessKiller", p.cmd.Path, p.cmd.Args),
	// 		Origin: err,
	// 	}
	// }
	// release the process resource
	// return nil
}
