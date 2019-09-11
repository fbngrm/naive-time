package display

import "time"

// naivetime is a timestamp representation of an UTC-offset naive time(offset removed).
type naivetime struct {
	timestamp int64
}

// naiveTime normalizes the time at the given location by removing the UTC-offset.
// Returns an Unix epoch timestamp representation of the normalized time.
func naiveTime(t time.Time, location string) (naivetime, error) {
	// local wall clock time in the provided location (UTC-offset aware).
	t, err := timeIn(t, location)
	if err != nil {
		return naivetime{}, err
	}
	// remove zone-offset
	_, offset := t.Zone()
	return naivetime{timestamp: t.Unix() + int64(offset)}, nil
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
func timeIn(t time.Time, location string) (time.Time, error) {
	loc, err := time.LoadLocation(location)
	if err != nil {
		return time.Time{}, err
	}
	return t.In(loc), nil
}
