package osinfo

import (
	"testing"
)

func TestOSInfo(t *testing.T) {
	oi := New()
	if err := oi.Err(); err != nil {
		t.Fatal(err)
	}
	t.Log(oi)
}
