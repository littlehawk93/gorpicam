package gorpicam

import (
	"bytes"
	"errors"
	"fmt"
	"io/ioutil"
	"os/exec"
	"regexp"
	"sync"
)

const (
	errAlreadyRunning string = "Video recording is already running"
	errNotRunning     string = "Video recording isn't running"
)

// VideoRecordingCompleteCallback callback function for when video recording finishes
type VideoRecordingCompleteCallback func(r *VideoCommand)

// VideoCommand manages the underlying CMD process for recording a video
type VideoCommand struct {
	args     []string
	cmd      *exec.Cmd
	running  bool
	m        *sync.Mutex
	cmdErr   *bytes.Buffer
	Callback VideoRecordingCompleteCallback
}

// Start start recording video
func (me *VideoCommand) Start() error {

	me.m = &sync.Mutex{}

	if me.running {
		return errors.New(errAlreadyRunning)
	}

	me.cmd = exec.Command("raspivid", me.args...)

	me.cmdErr = &bytes.Buffer{}
	me.cmd.Stderr = me.cmdErr

	err := me.cmd.Start()

	if err != nil {
		return err
	}

	go func() {
		me.m.Lock()
		me.running = true
		me.cmd.Wait()

		if me.Callback != nil {
			me.Callback(me)
		}

		me.cmd.Process.Kill()
		me.running = false
		me.m.Unlock()
	}()

	return nil
}

// Wait wait for this video command to finish recording
func (me *VideoCommand) Wait() error {

	if !me.running {
		return errors.New(errNotRunning)
	}

	me.m.Lock()
	return nil
}

// Stop force command to stop recording
func (me *VideoCommand) Stop() error {

	if !me.running {
		return errors.New(errNotRunning)
	}

	err := me.cmd.Process.Kill()

	if err == nil {
		me.running = false
	}
	return err
}

// ExitStatus get the exit code for this command. If it has not finished running, returns 0
func (me *VideoCommand) ExitStatus() int {

	if me.running || me.cmd == nil {
		return 0
	}

	return me.cmd.ProcessState.ExitCode()
}

// IsRecording returns true if this command is currently recording
func (me *VideoCommand) IsRecording() bool {

	return me.running
}

// ErrorMsg get any messages written to Stderr during recording
func (me *VideoCommand) ErrorMsg() string {

	if !me.running && me.cmdErr != nil {
		bytes, _ := ioutil.ReadAll(me.cmdErr)
		return string(bytes)
	}
	return ""
}

// CommandString return the underlying CLI command and arguments used to run this video command
func (me *VideoCommand) CommandString() string {

	buffer := &bytes.Buffer{}

	if _, err := buffer.WriteString("raspivid"); err != nil {
		panic("Unable to write to string buffer")
	}

	if me.args != nil && len(me.args) > 0 {

		regex := regexp.MustCompile("\\s+")

		for _, val := range me.args {
			if regex.MatchString(val) {
				if _, err := buffer.WriteString(fmt.Sprintf(" \"%s\"", val)); err != nil {
					panic("Unable to write to string buffer")
				}
			} else {
				if _, err := buffer.WriteString(fmt.Sprintf(" %s", val)); err != nil {
					panic("Unable to write to string buffer")
				}
			}
		}
	}
	return buffer.String()
}

// NewVideo create a new raspivid command to execute
func NewVideo(output string, options VidOptions) *VideoCommand {

	args := options.CmdArgs()

	args = pushString(args, "output", output)

	return &VideoCommand{
		args:    args,
		running: false,
		cmd:     nil,
		cmdErr:  nil,
	}
}
