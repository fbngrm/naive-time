package banner

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
// The check is performed by considering the
func (b banner) activeIn(location string) (bool, error) {
	// get time in location
	// t,err:=timeIn(location)
	return false, nil
}
