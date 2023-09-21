package osinfo

import (
	"fmt"
	"runtime"
)

// New returns an instance of OSInfo.
func New() (oi *OSInfo) {
	oi = &OSInfo{
		Name:    runtime.GOOS,
		Version: UnknownRelease,
	}
	kv, err := kernelRelease()
	if err != nil {
		oi.err = err
		return
	}
	oi.Version = fmt.Sprintf("(%s)", kv)

	return
}
