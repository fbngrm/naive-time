package display

import (
	"fmt"
	"time"
)

// period is an UTC-offset aware period in time. It is used as the display
// period for a banner.
// begin represents seconds of UTC time since Unix epoch 1970-01-01T00:00:00Z.
// Values must be from 0001-01-01T00:00:00Z to 9999-12-31T23:59:59Z inclusive
// to be parsable in a RFC3339 compliant layout.
// NOTE: This program does not provide a mechanism to deal with future changes
// in the IANA database rules. Therefor, when creating a period, it is not
// recommended to convert a time with an UTC-offset other than `00:00` to a
// timestamp.
type Period struct {
	begin    int64 // seconds since unix epoch
	duration int64 // seconds banner should be displayed
}

// active checks if the naive time in location is an instant in period.
func (p Period) Active(t time.Time, location string) (bool, error) {
	n, err := naiveTime(t, location)
	if err != nil {
		return false, err
	}
	return n.timestamp >= p.begin && n.timestamp < p.begin+p.duration, nil
}

// String returns a human readable representation of the period.
func (p Period) String() string {
	begin := fmt.Sprintf("begin: %v", p.begin)
	duration := fmt.Sprintf("duration: %d seconds", p.duration)
	expire := fmt.Sprintf("expire: %v\n", time.Unix(p.begin+p.duration, 0).UTC())
	return fmt.Sprintf("%s\n%s\n%s\n", begin, duration, expire)
}
