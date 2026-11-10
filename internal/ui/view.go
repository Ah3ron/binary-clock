package ui

import (
	"fmt"

	"github.com/Ah3ron/binary-clock/internal/bcd"
	"github.com/charmbracelet/lipgloss"
)

// maxBits is the widest digit column; it defines the number of bit rows.
var maxBits = maxOf(bcd.BitsPerColumn[:])

func maxOf(vals []int) int {
	max := 0
	for _, v := range vals {
		if v > max {
			max = v
		}
	}
	return max
}

// View renders the clock, or nothing while quitting.
func (m Model) View() string {
	if m.quitting {
		return ""
	}
	return m.renderBCD()
}

func (m Model) renderBCD() string {
	s := m.styles
	d := bcd.Digits(m.now)

	// column labels
	label := "    "
	for i, name := range bcd.ColumnLabels {
		label += s.colLabel.Render(name)
		if i != len(bcd.ColumnLabels)-1 {
			label += s.spacer.Render(" ")
		}
	}
	label += "\n\n"

	// bit rows (top -> bottom)
	rows := ""
	for r := maxBits - 1; r >= 0; r-- {
		bitValue := 1 << r
		row := s.bitValue.Render(fmt.Sprintf("%2d |", bitValue)) + " "

		for c := 0; c < len(d); c++ {
			if r < bcd.BitsPerColumn[c] {
				cell := s.offDot.Render("▒▒")
				if bcd.BitAt(d[c], r) == 1 {
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
