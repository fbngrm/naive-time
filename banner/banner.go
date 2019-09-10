package banner

import (
	"time"
)

// period is the display period for a banner.
// begin represents seconds of UTC time since Unix epoch 1970-01-01T00:00:00Z.
// Values must be from 0001-01-01T00:00:00Z to 9999-12-31T23:59:59Z inclusive
// to be parsable in a RFC3339 compliant layout.
// NOTE: This program does not provide a mechanism to deal with future changes
// in the IANA database rules. Therefor, when creating a period, it is not
// recommended to convert a time with an UTC-offset other than `00:00` to a
// timestamp.
type period struct {
	begin    int64 // seconds since unix epoch
	duration int64 // seconds banner should be displayed
}

// banner represents a banner associated with a display period.
// A banner is active during the display period. It is expired after the
// display period exceeded.
type banner struct {
	id      int64
	content string
	period  period
}

// activeIn checks if a banner is currently active in the given location.
func (b banner) activeIn(location string) (bool, error) {
	t, err := timeIn(location)
	if err != nil {
		return false, err
	}
	return naive(t).in(b.period), nil
}

// naivetime represents an UTC-offset-naive time.
type naivetime struct {
	timestamp int64
}

// naive normalizes the given time by removing the UTC-offset.
// Returns an Unix epoch timestamp representation.
func naive(t time.Time) naivetime {
	_, offset := t.Zone()
	return naivetime{timestamp: t.Unix() + int64(offset)}
}

// in checks if naivetime is an instant in the given period.
func (n naivetime) in(p period) bool {
	return n.timestamp >= p.begin && n.timestamp < p.begin+p.duration
}

// timeIn loads the time in the provided location. If the provided location name
// is "" or "UTC" it returns the current time in UTC. If the name is "Local" it
// returns the local time. Otherwise, the name is looked up in the systems zone
// information and is assumed to be a fully qualified name matching the IANA
// Time Zone database format, e.g. "Asia/Tokyo".
// NOTE: The location name is looked up in the directory or uncompressed zip
// file named by the ZONEINFO environment variable, if any, then looks in known
// installation locations on Unix systems, and finally looks in
// $GOROOT/lib/time/zoneinfo.zip.
func timeIn(location string) (time.Time, error) {
	loc, err := time.LoadLocation(location)
	if err != nil {
		return time.Time{}, err
	}
	return time.Now().In(loc), nil
}
