package osinfo

import "fmt"

const (
	// UnknownRelease returns when can't determinate os release.
	UnknownRelease = "unknown"
)

// OSInfo contains information of running OS.
type OSInfo struct {
	Name    string `json:"os" yaml:"os"`
	Version string `json:"version" yaml:"version"`
	err     error
}

func (oi *OSInfo) String() string {
	return fmt.Sprintf("%s %s", oi.Name, oi.Version)
}

func (oi *OSInfo) Err() error {
	return oi.err
}

// Release returns release version of current running OS. It returns a
// default string if can't find.
func Release() string {
	return New().Version
}
