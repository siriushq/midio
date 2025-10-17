package console

import "github.com/fatih/color"

var (
	// Theme contains default color mapping.
	Theme = map[string]*color.Color{
		"Debug":  color.New(color.FgWhite, color.Faint, color.Italic),
		"Fatal":  color.New(color.FgRed, color.Italic, color.Bold),
		"Error":  color.New(color.FgYellow, color.Italic),
		"Info":   color.New(color.FgGreen, color.Bold),
		"Print":  color.New(),
		"PrintB": color.New(color.FgBlue, color.Bold),
		"PrintC": color.New(color.FgGreen, color.Bold),
	}
)

// SetColorOff disables coloring for the entire session.
func SetColorOff() {
	privateMutex.Lock()
	defer privateMutex.Unlock()
	color.NoColor = true
}

// SetColorOn enables coloring for the entire session.
func SetColorOn() {
	privateMutex.Lock()
	defer privateMutex.Unlock()
	color.NoColor = false
}

// SetColor sets a color for a particular tag.
func SetColor(tag string, cl *color.Color) {
	privateMutex.Lock()
	defer privateMutex.Unlock()
	// add new theme
	Theme[tag] = cl
}
