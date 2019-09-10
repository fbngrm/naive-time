package banner

import "time"

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
	// local wall clock time in the provided location (UTC-offset aware).
	t, err := timeIn(location)
	if err != nil {
		return false, err
	}
	// check UTC-offset naive time against the display period.
	return naive(t).in(b.period), nil
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
