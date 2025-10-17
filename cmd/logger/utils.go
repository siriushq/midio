package logger

import (
	"fmt"
	"regexp"
	"runtime"

	"github.com/siriushq/midio/pkg/color"
)

var ansiRE = regexp.MustCompile("(\x1b[^m]*m)")

// Print ANSI Control escape
func ansiEscape(format string, args ...interface{}) {
	var Esc = "\x1b"
	fmt.Printf("%s%s", Esc, fmt.Sprintf(format, args...))
}

func ansiMoveRight(n int) {
	if runtime.GOOS == "windows" {
		return
	}
	if color.IsTerminal() {
		ansiEscape("[%dC", n)
	}
}

func ansiSaveAttributes() {
	if runtime.GOOS == "windows" {
		return
	}
	if color.IsTerminal() {
		ansiEscape("7")
	}
}

func ansiRestoreAttributes() {
	if runtime.GOOS == "windows" {
		return
	}
	if color.IsTerminal() {
		ansiEscape("8")
	}

}
