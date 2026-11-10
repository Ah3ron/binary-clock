package ui

import "github.com/charmbracelet/lipgloss"

type styles struct {
	root     lipgloss.Style
	colLabel lipgloss.Style
	digit    lipgloss.Style
	onDot    lipgloss.Style
	offDot   lipgloss.Style
	spacer   lipgloss.Style
	bitValue lipgloss.Style
}

func newStyles() styles {
	purple := lipgloss.AdaptiveColor{Light: "#7c4dff", Dark: "#b388ff"}
	green := lipgloss.AdaptiveColor{Light: "#2e7d32", Dark: "#66bb6a"}
	gray := lipgloss.AdaptiveColor{Light: "#6b7280", Dark: "#9ca3af"}

	return styles{
		root:     lipgloss.NewStyle().Padding(0, 1),
		colLabel: lipgloss.NewStyle().Foreground(purple).Bold(true).Width(3).Align(lipgloss.Center),
		digit:    lipgloss.NewStyle().Foreground(gray).Width(3).Align(lipgloss.Center),
		onDot:    lipgloss.NewStyle().Foreground(green).Bold(true),
		offDot:   lipgloss.NewStyle().Foreground(gray),
		spacer:   lipgloss.NewStyle().Width(1),
		bitValue: lipgloss.NewStyle().Foreground(gray).Width(4),
	}
}
