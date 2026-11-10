// Command binary-clock renders the current time as a binary-coded decimal
// (BCD) clock in the terminal.
package main

import (
	"fmt"
	"os"

	"github.com/Ah3ron/binary-clock/internal/ui"
	tea "github.com/charmbracelet/bubbletea"
)

func main() {
	p := tea.NewProgram(ui.New(), tea.WithAltScreen())
	if _, err := p.Run(); err != nil {
		fmt.Fprintln(os.Stderr, "error:", err)
		os.Exit(1)
	}
}
