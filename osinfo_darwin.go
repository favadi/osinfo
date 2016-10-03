package osinfo

import (
	"bytes"
	"os/exec"
	"runtime"
)

// New returns a instance of OSInfo if the current running OS is OS X.
func New() (oi *OSInfo, err error) {
	oi = &OSInfo{
		Name: runtime.GOOS,
	}

	cmd := exec.Command("/usr/bin/sw_vers", "-productVersion")
	var out bytes.Buffer
	cmd.Stdout = &out
	if err = cmd.Run(); err != nil {
		return
	}
	oi.Version = out.String()
	return
}
