package console

import (
	"testing"

	"github.com/fatih/color"
)

func TestSetColor(t *testing.T) {
	SetColor("unknown", color.New(color.FgWhite))
	_, ok := Theme["unknown"]
	if !ok {
		t.Fatal("missing theme")
	}
}

func TestColorLock(t *testing.T) {
	Lock()
	Print("") // Test for deadlocks.
	Unlock()
}
