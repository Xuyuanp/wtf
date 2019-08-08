package wtf

import "testing"

func TestWTF(t *testing.T) {
	if want, got := "WTF", WTF(); want != got {
		t.Fatalf("want: %v, but got: %v", want, got)
	}
}
