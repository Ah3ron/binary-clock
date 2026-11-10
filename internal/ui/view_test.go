package ui

import (
	"strings"
	"testing"
	"time"
)

// fixedModel returns a model pinned to 12:34:56 for deterministic rendering.
func fixedModel() Model {
	m := New()
	m.now = time.Date(2026, 1, 1, 12, 34, 56, 0, time.UTC)
	m.width = 80
	m.height = 24
	return m
}

// rowPattern returns the rendered bit row for the given bit value with
// whitespace stripped, e.g. "▒▒████▒▒▒▒██".
func rowPattern(out string, bitValue string) (string, bool) {
	for _, line := range strings.Split(out, "\n") {
		if !strings.Contains(line, "|") || !strings.Contains(line, bitValue) {
			continue
		}
		parts := strings.SplitN(line, "|", 2)
		return strings.Join(strings.Fields(parts[1]), ""), true
	}
	return "", false
}

func TestRenderBCDBitPattern(t *testing.T) {
	// 12:34:56 -> digits 1 2 3 4 5 6.
	// Rows are drawn top (bit 8) to bottom (bit 1); lit bits use "██",
	// unlit "▒▒", and columns without that bit are blank.
	want := map[string]string{
		" 8 |": "▒▒▒▒▒▒",
		" 4 |": "▒▒▒▒██████",
		" 2 |": "▒▒████▒▒▒▒██",
		" 1 |": "██▒▒██▒▒██▒▒",
	}

	out := fixedModel().renderBCD()

	for prefix, pattern := range want {
		got, ok := rowPattern(out, prefix)
		if !ok {
			t.Errorf("bit row %q not found in output:\n%s", prefix, out)
			continue
		}
		if got != pattern {
			t.Errorf("bit row %q = %s, want %s", prefix, got, pattern)
		}
	}

	if !strings.Contains(out, "1") || !strings.Contains(out, "6") {
		t.Error("digit summary row missing")
	}
}

func TestViewQuitting(t *testing.T) {
	m := fixedModel()
	m.quitting = true
	if got := m.View(); got != "" {
		t.Errorf("View() while quitting = %q, want empty", got)
	}
}

func TestViewDefaultSize(t *testing.T) {
	m := fixedModel()
	m.width, m.height = 0, 0
	if got := m.View(); got == "" {
		t.Error("View() with zero size should fall back to defaults")
	}
}
