package banner

import "time"

// naivetime represents an UTC-offset naive time.
type naivetime struct {
	timestamp int64
}

// naiveTime normalizes the time at the given location by removing the UTC-offset.
// Returns an Unix epoch timestamp representation of the normalized time.
func naiveTime(location string) (naivetime, error) {
	// local wall clock time in the provided location (UTC-offset aware).
	t, err := timeIn(location)
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
func timeIn(location string) (time.Time, error) {
	loc, err := time.LoadLocation(location)
	if err != nil {
		return time.Time{}, err
	}
	return time.Now().In(loc), nil
}
