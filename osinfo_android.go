package osinfo

import (
	"bufio"
	"fmt"
	"os"
	"runtime"
	"strings"
)

const (
	buildPropPath = "/system/build.prop"
	releasePrefix = " ro.build.version.release="
)

// New returns an instance of OSInfo.
func New() (oi *OSInfo, err error) {
	oi = &OSInfo{
		Name: runtime.GOOS,
	}

	f, err := os.Open(buildPropPath)
	if err != nil {
		return
	}

	sc := bufio.NewScanner(f)
	for sc.Scan() {
		line := sc.Text()
		if strings.HasPrefix(line, releasePrefix) {
			oi.Version = strings.TrimSpace(strings.TrimPrefix(releasePrefix, line))
			break
		}
	}

	if len(oi.Version) == 0 {
		err = fmt.Errorf("property not found: %s", releasePrefix)
	}
	return
}
