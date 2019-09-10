package banner

import "time"

// naivetime represents an UTC-offset-naive time.
type naivetime struct {
	timestamp int64
}

// naive normalizes the given time by removing the UTC-offset.
// Returns an Unix epoch timestamp representation of the normalized time.
func naive(t time.Time) naivetime {
	_, offset := t.Zone()
	return naivetime{timestamp: t.Unix() + int64(offset)}
}

// in checks if naivetime is an instant in the given period.
func (n naivetime) in(p period) bool {
	return n.timestamp >= p.begin && n.timestamp < p.begin+p.duration
}
