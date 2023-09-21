//go:build !android
// +build !android

package osinfo

import (
	"bytes"
	"fmt"
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

	// Use kernel release for better description:
	//
	// - "uname --kernel-release": 5.14.0-1048-oem
	// - "uname -v"              : #55-Ubuntu SMP Mon Aug 8 14:58:10 UTC 2022
	kr, err := kernelRelease()
	if err != nil {
		oi.err = err
		return
	}

	// distribution name and release
	lsbPath, err := exec.LookPath("lsb_release")
	if err != nil { // lsb_release is not installed, returns kernel release
		oi.Version = fmt.Sprintf("(%s)", kr)
		oi.err = err
		return
	}
	lsbCmd := exec.Command(lsbPath, "--description")
	var lsbOut bytes.Buffer
	lsbCmd.Stdout = &lsbOut
	if err := lsbCmd.Run(); err != nil {
		oi.err = err
		return
	}
	desc := lsbOut.String()
	// Description:	Ubuntu 14.04.4 LTS
	dist := strings.TrimSpace(strings.TrimPrefix(desc, "Description:\t"))
	oi.Version = fmt.Sprintf("%s (%s)", dist, kr)

	return
}
