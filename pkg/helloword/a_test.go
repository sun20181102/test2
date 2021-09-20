package hellowword

import "testing"

func TestHello(t *testing.T) {
	want := "hello world"
	if got := HelloWorld(); got != want {
		t.Errorf("Hello()=%s, want=%s", got, want)
	}
}
