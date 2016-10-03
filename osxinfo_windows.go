package osinfo

import (
	"fmt"
	"runtime"
	"syscall"
)

// New returns a instance of OSInfo if the current running OS is Windows.
func New() (oi *OSInfo, err error) {
	oi = &OSInfo{
		Name: runtime.GOOS,
	}

	dll := syscall.MustLoadDLL("kernel32.dll")
	p, err := dll.FindProc("GetVersion")
	if err != nil {
		return
	}
	v, _, err := p.Call()
	// err is always non-nil, if v is zero, it is an error.
	version := uint32(v)
	if version == 0 {
		return
	} else {
		err = nil
	}
	oi.Version = fmt.Sprintf("%d.%d", byte(v), uint8(v>>8))
	return
}
