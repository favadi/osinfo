package osinfo

import (
	"bytes"
	"os/exec"
	"runtime"
	"strings"
)

// New returns an instance of OSInfo.
func New() (oi *OSInfo) {
	oi = &OSInfo{
		Name:    runtime.GOOS,
		Version: UnknownRelease,
	}

	// kernel vesion
	unameCmd := exec.Command("uname", "-v")
	var unameOut bytes.Buffer
	unameCmd.Stdout = &unameOut
	if err := unameCmd.Run(); err != nil {
		oi.err = err
		return
	}
	oi.Version = strings.TrimSpace(unameOut.String())

	return
}
