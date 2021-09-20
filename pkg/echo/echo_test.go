package echo

import "testing"

func TestEcho(t *testing.T) {
	if msg := Echo("test"); msg != "test" {
		t.Errorf("Echo(): %s, want=%s", msg, "test")
	}
}
