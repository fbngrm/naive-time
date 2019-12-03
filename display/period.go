package display

import (
	"time"
)

// period is an UTC-offset aware period in time. It is used as the display
// period for a banner.
// begin represents seconds of UTC time since
// Unix epoch 1970-01-01T00:00:00Z.
// Todo, use time.Duration and time.Location types
type Period struct {
	begin    *naivetime // seconds since unix epoch
	duration int64      // seconds banner should be displayed
}

// New returns a Period instantiated with the given time.
func New(t time.Time, duration int64) (*Period, error) {
	loc := t.Location().String()
	nt, err := naiveTime(t, loc)
	if err != nil {
		return nil, err
	}
	return &Period{
		begin:    nt,
		duration: duration,
		location: loc,
	}, nil
}

// active checks if the naive representation of the given time in location
// is an instant in period; if the Period is active in this location.
// If internal is true, it is ensured only that Period is not expired,
// begin is omitted.
func (p Period) Active(t time.Time, location string, internal bool) (bool, error) {
	n, err := naiveTime(t, location)
	if err != nil {
		return false, err
	}
	started := n.timestamp >= p.begin.timestamp
	expired := n.timestamp >= p.begin.timestamp+p.duration
	return (started || internal) && !expired, nil
}
