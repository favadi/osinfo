//go:build unix

package osinfo

import (
	"strings"

	"golang.org/x/sys/unix"
)

func kernelRelease() (string, error) {
	var un unix.Utsname
	if err := unix.Uname(&un); err != nil {
		return "", err
	}
	return unameFieldToString(un.Release[:]), nil
}

// TODO(cuonglm): use unsafe.String when drop support for go1.19
func unameFieldToString(sl []byte) string {
	var sb strings.Builder
	for _, b := range sl {
		if b == 0 {
			break
		}
		sb.WriteByte(b)
	}
	return strings.TrimSpace(sb.String())
}
