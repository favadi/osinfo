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
func New() (oi *OSInfo, err error) {
	oi = &OSInfo{
		Name: runtime.GOOS,
	}

	// kernel vesion
	unameCmd := exec.Command("uname", "--kernel-release")
	var unameOut bytes.Buffer
	unameCmd.Stdout = &unameOut
	if err = unameCmd.Run(); err != nil {
		return
	}
	kernelRelease := strings.TrimSpace(unameOut.String())

	// distribution name and release
	lsbPath, err := exec.LookPath("lsb_release")
	if err != nil { // lsb_release is not installed, returns kernel release
		oi.Version = kernelRelease
		return
	}
	lsbCmd := exec.Command(lsbPath, "--description")
	var lsbOut bytes.Buffer
	lsbCmd.Stdout = &lsbOut
	if err = lsbCmd.Run(); err != nil {
		return
	}
	desc := lsbOut.String()
	// Description:	Ubuntu 14.04.4 LTS
	dist := strings.TrimSpace(strings.TrimPrefix(desc, "Description:\t"))
	oi.Version = fmt.Sprintf("%s (%s)", dist, kernelRelease)

	return
}
