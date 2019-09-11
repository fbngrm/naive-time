package banner

import (
	"time"

	"github.com/m-rec/0ca17a7468266cd599c376f2c522790404ed878f/display"
)

// banner represents a banner associated with a display period.
// A banner is active during the display period. It is expired after the
// display period exceeded.
type banner struct {
	id      int64
	content string
	period  display.Period
}

// activeIn checks if a banner is active in the given location at a given time.
func (b banner) activeIn(t time.Time, location string, internal bool) (bool, error) {
	// check UTC-offset naive time in location against the display period.
	return b.period.Active(t, location, internal)
}
