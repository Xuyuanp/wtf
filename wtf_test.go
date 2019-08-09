package wtf

import "testing"

func TestWTF2(t *testing.T) {
	if want, got := "WTF2", WTF2(); want != got {
		t.Fatalf("want: %v, but got: %v", want, got)
	}
}
