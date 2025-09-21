package main

import (
	"fmt"
	"os"
	"time"

	"github.com/charmbracelet/bubbles/help"
	"github.com/charmbracelet/bubbles/key"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type model struct {
	now      time.Time
	help     help.Model
	keys     keymap
	quitting bool
	styles   styles
	width    int
	height   int
}

type keymap struct{ quit key.Binding }

type styles struct {
	root      lipgloss.Style
	header    lipgloss.Style
	clockBox  lipgloss.Style
	rowLabel  lipgloss.Style
	colLabel  lipgloss.Style
	digit     lipgloss.Style
	onDot     lipgloss.Style
	offDot    lipgloss.Style
	spacer    lipgloss.Style
	helpStyle lipgloss.Style
	footer    lipgloss.Style
	bitValue  lipgloss.Style
}

func newStyles() styles {
	purple := lipgloss.AdaptiveColor{Light: "#7c4dff", Dark: "#b388ff"}
	green := lipgloss.AdaptiveColor{Light: "#2e7d32", Dark: "#66bb6a"}
	gray := lipgloss.AdaptiveColor{Light: "#6b7280", Dark: "#9ca3af"}

	return styles{
		root:     lipgloss.NewStyle().Padding(0, 1),
		header:   lipgloss.NewStyle().Foreground(purple).Bold(true),
		clockBox: lipgloss.NewStyle(),
		rowLabel: lipgloss.NewStyle().Width(4).Foreground(gray),
		colLabel: lipgloss.NewStyle().Foreground(purple).Bold(true).Width(3).Align(lipgloss.Center),
		digit:    lipgloss.NewStyle().Foreground(gray).Width(3).Align(lipgloss.Center),
		onDot:    lipgloss.NewStyle().Foreground(green).Bold(true),
		offDot:   lipgloss.NewStyle().Foreground(gray),
		spacer:   lipgloss.NewStyle().Width(1),
		bitValue: lipgloss.NewStyle().Foreground(gray).Width(4),
	}
}

func newKeymap() keymap {
	return keymap{
		quit: key.NewBinding(
			key.WithKeys("q", "ctrl+c"),
			key.WithHelp("q", "quit"),
		),
	}
}

func initialModel() model {
	return model{
		now:    time.Now(),
		help:   help.New(),
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

func (m model) Init() tea.Cmd { return tea.Batch(tickCmd(), tea.EnterAltScreen) }

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
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

/* BCD helpers */

func digits(t time.Time) [6]int {
	h := t.Hour()
	return [6]int{
		h / 10, h % 10,
		t.Minute() / 10, t.Minute() % 10,
		t.Second() / 10, t.Second() % 10,
	}
}

// per-column number of bits for typical BCD clock
var colBits = []int{2, 4, 3, 4, 3, 4}

func bitAt(d, p int) int { return (d >> p) & 1 }

/* rendering */

const maxRows = 4

func (m model) renderBCD() string {
	s := m.styles
	d := digits(m.now)

	// column labels
	colNames := []string{"H", "h", "M", "m", "S", "s"}
	label := s.clockBox.Render("    ")
	for i := range colNames {
		label += s.colLabel.Render(colNames[i])
		if i != len(colNames)-1 {
			label += s.spacer.Render(" ")
		}
	}
	label = label + "\n\n"

	// bit rows (top -> bottom)
	rows := ""
	for r := maxRows - 1; r >= 0; r-- {
		bitValue := 1 << r
		row := s.bitValue.Render(fmt.Sprintf("%2d |", bitValue)) + " "

		for c := 0; c < len(d); c++ {
			if r < colBits[c] {
				cell := s.offDot.Render("▒▒")
				if bitAt(d[c], r) == 1 {
					cell = s.onDot.Render("██")
				}
				row += s.digit.Render(cell)
			} else {
				row += s.digit.Render("   ")
			}
			if c != len(d)-1 {
				row += s.spacer.Render(" ")
			}
		}
		rows += row + "\n\n"
	}

	// digit summary
	digitsRow := "     "
	for i := range d {
		digitsRow += s.digit.Render(fmt.Sprintf("%d  ", d[i]))
		if i != len(d)-1 {
			digitsRow += s.spacer.Render(" ")
		}
	}

	body := label + rows + digitsRow

	w := m.width
	h := m.height
	if w <= 0 {
		w = 80
	}
	if h <= 0 {
		h = 24
	}

	block := lipgloss.NewStyle().Width(lipgloss.Width(body)).Render(body)
	centered := lipgloss.Place(w, h, lipgloss.Center, lipgloss.Center, block)

	return s.root.Render(centered)
}

func (m model) View() string {
	if m.quitting {
		return ""
	}
	return m.renderBCD()
}

func main() {
	p := tea.NewProgram(initialModel(), tea.WithAltScreen(), tea.WithMouseCellMotion())
	if _, err := p.Run(); err != nil {
		fmt.Fprintln(os.Stderr, "error:", err)
		os.Exit(1)
	}
}
