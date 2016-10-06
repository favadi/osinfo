// +build !android

package osinfo

import (
	"bytes"
	"os/exec"
	"runtime"
	"strings"
)

// New returns an instance of OSInfo.
func New() (oi *OSInfo, err error) {
	oi = &OSInfo{
		Name: runtime.GOOS,
	}

	cmd := exec.Command("/usr/bin/lsb_release", "--description")
	var out bytes.Buffer
	cmd.Stdout = &out
	if err = cmd.Run(); err != nil {
		return
	}
	desc := out.String()
	// Description:	Ubuntu 14.04.4 LTS
	oi.Version = strings.TrimPrefix(desc, "Description:\t")
	return
}
