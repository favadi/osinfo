package osinfo

import "fmt"

// OSInfo contains information of running OS.
type OSInfo struct {
	Name    string `json:"os" yaml:"os"`
	Version string `json:"version" yaml:"version"`
}

func (oi *OSInfo) String() string {
	return fmt.Sprintf("%s %s", oi.Name, oi.Version)
}
