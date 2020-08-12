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
	releasePrefix = "ro.build.version.release="
)

// New returns an instance of OSInfo.
func New() (oi *OSInfo) {
	oi = &OSInfo{
		Name:    runtime.GOOS,
		Version: UnknownRelease,
	}

	f, err := os.Open(buildPropPath)
	if err != nil {
		oi.err = err
		return
	}
	defer f.Close()

	sc := bufio.NewScanner(f)
	for sc.Scan() {
		line := sc.Text()
		if strings.HasPrefix(line, releasePrefix) {
			oi.Version = strings.TrimSpace(strings.TrimPrefix(line, releasePrefix))
			break
		}
	}

	if len(oi.Version) == 0 {
		oi.err = fmt.Errorf("property not found: %s", releasePrefix)
	}
	return
}
