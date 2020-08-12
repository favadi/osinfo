package osinfo

import (
	"fmt"
	"runtime"
	"syscall"
)

// New returns an instance of OSInfo.
func New() (oi *OSInfo) {
	oi = &OSInfo{
		Name:    runtime.GOOS,
		Version: UnknownRelease,
	}

	dll := syscall.MustLoadDLL("kernel32.dll")
	p, err := dll.FindProc("GetVersion")
	if err != nil {
		oi.err = err
		return
	}
	v, _, err := p.Call()
	// err is always non-nil, if v is zero, it is an error.
	version := uint32(v)
	if version == 0 {
		oi.err = err
		return
	}

	oi.Version = fmt.Sprintf("%d.%d", byte(v), uint8(v>>8))
	return
}
