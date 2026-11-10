// Package bcd encodes time as binary-coded decimal digits and exposes the
// bit layout used by the clock rendering.
package bcd

import "time"

// Columns is the number of BCD digit columns: HH MM SS.
const Columns = 6

// BitsPerColumn holds the number of bits shown for each digit column,
// from most significant (hour tens) to least significant (second units).
var BitsPerColumn = [Columns]int{2, 4, 3, 4, 3, 4}

// ColumnLabels are the short column captions, upper-case for tens
// and lower-case for units.
var ColumnLabels = [Columns]string{"H", "h", "M", "m", "S", "s"}

// Digits splits t into six BCD digits: hour tens, hour units,
// minute tens, minute units, second tens, second units.
func Digits(t time.Time) [Columns]int {
	h := t.Hour()
	return [Columns]int{
		h / 10, h % 10,
		t.Minute() / 10, t.Minute() % 10,
		t.Second() / 10, t.Second() % 10,
	}
}

// BitAt returns bit p of digit d (0 is the least significant bit).
func BitAt(d, p int) int { return (d >> p) & 1 }
