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

	cmd := exec.Command("/usr/bin/sw_vers", "-productVersion")
	var out bytes.Buffer
	cmd.Stdout = &out
	if err := cmd.Run(); err != nil {
		oi.err = err
		return
	}
	oi.Version = strings.TrimSpace(out.String())
	return
}
