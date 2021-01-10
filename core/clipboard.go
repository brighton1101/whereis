package core

import (
	"errors"
	"fmt"
	"os/exec"
	"runtime"
)

// Allows users to copy a string to their
// systems' clipboard
type Clipboard interface {
	Copy(string) error
}

type Shell struct {
	cmd string
}

type Powershell struct {
	cmd string
}

const (
	LinuxR    = "linux"
	MacR      = "darwin"
	WindowsR  = "windows"
	Pbcopy    = "pbcopy"
	Pscopy    = "clip"
	XclipName = "xclip"
	XclipCmd  = "xclip -selection clipboard"
	XselName  = "xsel"
	XselCmd   = "xsel --clipboard"
)

var (
	ErrUnsupportedOS   = errors.New("Unsupported Operating System for operation")
	ErrNoClipInstalled = errors.New("No clipboard application installed")
)

type LinuxCmdMapping struct {
	name string
	cmd  string
}

// Maps the name of the utility application for Linux
// OS for clipboard operations to the appropriate sh
// command to use. Currently, xsel and xclip are
// supported.
var LinuxCmdMappings = [2]LinuxCmdMapping{
	LinuxCmdMapping{XclipName, XclipCmd},
	LinuxCmdMapping{XselName, XselCmd},
}

// Gets clipboard appropriate to users' OS
func GetClipboard() (Clipboard, error) {
	var clip Clipboard
	var err error
	switch runtime.GOOS {
	case LinuxR:
		cmd, err := unixGetClipCmd()
		if err != nil {
			break
		}
		clip = &Shell{cmd: cmd}
	case WindowsR:
		clip = &Powershell{cmd: Pscopy}
	case MacR:
		clip = &Shell{cmd: Pbcopy}
	}
	return clip, err
}

// Verifies a supported Linux clipboard copier
// is installed. See `LinuxCmdMappings` for supported
func unixGetClipCmd() (string, error) {
	if runtime.GOOS != LinuxR {
		return "", ErrUnsupportedOS
	}
	for _, mapping := range LinuxCmdMappings {
		whichcmd := fmt.Sprintf("which %s", mapping.name)
		bytesout, _ := exec.Command("sh", "-c", whichcmd).Output()
		strout := string(bytesout)
		if len(strout) != 0 {
			return mapping.cmd, nil
		}
	}
	return "", ErrNoClipInstalled
}

// Copies using shell command on a unix-like machine
// ie, MacOS or Linux
func (c *Shell) Copy(text string) error {
	cmd := fmt.Sprintf("echo \"%s\" | %s", text, c.cmd)
	return exec.Command("sh", "-c", cmd).Run()
}

// Copies using powershell command on a windows machine
func (c *Powershell) Copy(text string) error {
	ps, _ := exec.LookPath("powershell.exe")
	cmd := fmt.Sprintf("echo \"%s\" | %s", text, c.cmd)
	return exec.Command(ps, cmd).Run()
}
