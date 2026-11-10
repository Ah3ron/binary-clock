package bcd

import (
	"testing"
	"time"
)

func TestDigits(t *testing.T) {
	tests := []struct {
		name string
		t    time.Time
		want [Columns]int
	}{
		{"midnight", time.Date(2026, 1, 1, 0, 0, 0, 0, time.UTC), [Columns]int{0, 0, 0, 0, 0, 0}},
		{"mixed", time.Date(2026, 1, 1, 12, 34, 56, 0, time.UTC), [Columns]int{1, 2, 3, 4, 5, 6}},
		{"end of day", time.Date(2026, 1, 1, 23, 59, 59, 0, time.UTC), [Columns]int{2, 3, 5, 9, 5, 9}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Digits(tt.t); got != tt.want {
				t.Errorf("Digits(%v) = %v, want %v", tt.t, got, tt.want)
			}
		})
	}
}

func TestBitAt(t *testing.T) {
	tests := []struct {
		d, p, want int
	}{
		{0, 0, 0},
		{1, 0, 1},
		{1, 1, 0},
		{5, 0, 1},
		{5, 1, 0},
		{5, 2, 1},
		{9, 3, 1},
	}
	for _, tt := range tests {
		if got := BitAt(tt.d, tt.p); got != tt.want {
			t.Errorf("BitAt(%d, %d) = %d, want %d", tt.d, tt.p, got, tt.want)
		}
	}
}

func TestBitsPerColumnFitsDigits(t *testing.T) {
	// Each column must hold its maximum possible digit:
	// hours 23, minutes 59, seconds 59.
	maxDigit := [Columns]int{2, 9, 5, 9, 5, 9}
	for c, bits := range BitsPerColumn {
		if max := 1<<bits - 1; max < maxDigit[c] {
			t.Errorf("column %d: %d bits hold max %d, need %d", c, bits, max, maxDigit[c])
		}
	}
}
