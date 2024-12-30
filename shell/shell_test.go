package shell

import (
	"testing"
)

func TestRunGETPrices(t *testing.T) {
	want := "BTC"
	got := GETPrices("BTC")

	assertStatus(t, want, got)
}

func assertStatus(t *testing.T, want string, got string) {
	if want != got {
		t.Errorf("got %s expected %s", want, got)
	}
}
