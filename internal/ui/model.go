// Package ui implements the Bubble Tea terminal UI for the binary clock.
package ui

import (
	"time"

	"github.com/charmbracelet/bubbles/key"
	tea "github.com/charmbracelet/bubbletea"
)

// Model is the Bubble Tea model for the binary clock.
type Model struct {
	now      time.Time
	keys     keymap
	quitting bool
	styles   styles
	width    int
	height   int
}

type keymap struct{ quit key.Binding }

func newKeymap() keymap {
	return keymap{
		quit: key.NewBinding(
			key.WithKeys("q", "ctrl+c"),
			key.WithHelp("q", "quit"),
		),
	}
}

// New returns the initial clock model.
func New() Model {
	return Model{
		now:    time.Now(),
		keys:   newKeymap(),
		styles: newStyles(),
		width:  80,
		height: 24,
	}
}

/* ticking */

type tickMsg time.Time

func tickCmd() tea.Cmd {
	return tea.Tick(time.Second, func(t time.Time) tea.Msg { return tickMsg(t) })
}

// Init starts the one-second tick loop.
func (m Model) Init() tea.Cmd { return tickCmd() }

// Update handles ticks, key presses and terminal resizes.
func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tickMsg:
		m.now = time.Time(msg)
		return m, tickCmd()

	case tea.KeyMsg:
		if key.Matches(msg, m.keys.quit) {
			m.quitting = true
			return m, tea.Quit
		}
		return m, nil

	case tea.WindowSizeMsg:
		m.width = msg.Width
		m.height = msg.Height
		return m, nil
	}

	return m, nil
}
