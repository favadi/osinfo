package osinfo

import "testing"

func TestOSInfo(t *testing.T) {
	oi, err := New()
	if err != nil {
		t.Fatal(err)
	}
	t.Log(oi)
}
