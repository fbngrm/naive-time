package banner

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
	// check UTC-offset naive time in location against the display period.
	return b.period.active(location)
}
