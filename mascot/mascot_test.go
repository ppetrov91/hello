package mascot_test

import (
	"testing"

	"github.com/ppetrov91/hello/mascot"
)

func TestMascot(t *testing.T) {
	if mascot.BestMascot() != "Go Gopher" {
		t.Fatal("Wrong mascot")
	}
}
