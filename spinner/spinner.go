// Package spinner provides a terminal spinner that counts and tracks elapsed time.
package spinner

import (
	"fmt"
	"time"
)

// Spinner cycles chars on term to indicate activity.
type Spinner struct {
	Chars []string
	Start time.Time
	Count int
}

// New creates a Spinner with Chars prepopulated.
func New() *Spinner {

	return &Spinner{
		Chars: []string{`-`, `\`, `|`, `/`},
	}
}

// Spin prints the next character.
func (sp *Spinner) Spin() {

	if sp.Start.IsZero() {
		sp.Start = time.Now()
	}

	fmt.Printf(" %s \r", sp.Chars[sp.Count%len(sp.Chars)])
	sp.Count++
}

// Elapsed reports seconds since the first Spin.
func (sp *Spinner) Elapsed() (secs float64) {

	return time.Since(sp.Start).Seconds()
}
