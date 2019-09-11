package display

import (
	"testing"
	"time"
)

var testPeriod = []struct {
	t    string // UTC-offset aware time
	d    int64  // duration in seconds
	e    error  // expected error
	want Period
}{
	{
		t: "2005-04-03T02:01:00Z",
		want: Period{
			begin:    naivetime{timestamp: 1112493660}, // 2005-04-03 02:01:00 +00:00 UTC
			duration: 60 * 60,                          // display period 1h
			location: "UTC",
		},
	},
	{
		t: "2005-04-03T11:01:00+09:00",
		want: Period{
			begin:    naivetime{timestamp: 1112526060}, // 2005-04-03 11:01:00 +0000 UTC
			duration: 12 * 60 * 60,                     // display period 12h
			location: "Asia/Tokyo",
		},
	},
	{
		t: "2005-04-03T04:01:00+02:00",
		want: Period{
			begin:    naivetime{timestamp: 1112500860}, // 2005-04-03 04:01:00 +0000 UTC
			duration: 24 * 60 * 60,                     // display period 24h
			location: "Europe/Berlin",
		},
	},
	{
		t: "2005-04-02T18:01:00-08:00",
		want: Period{
			begin:    naivetime{timestamp: 1112464860}, // 2005-04-02 18:01:00 +0000 UTC
			duration: 60,                               // display period 1m
			location: "America/Los_Angeles",
		},
	},
	{
		t: "2005-04-02T15:01:00-11:00",
		want: Period{
			begin:    naivetime{timestamp: 1112454060}, // 2005-04-02 15:01:00 +0000 UTC
			duration: 48 * 60 * 60,                     // display period 48h
			location: "Pacific/Midway",
		},
	},
	// TODO: test invalid time, location and duration
}

func TestNew(t *testing.T) {
	for _, tt := range testPeriod {
		loc, err := time.LoadLocation(tt.want.location)
		if err != nil {
			t.Fatalf("error loading location %s: %v", tt.want.location, err)
		}
		ti, err := time.ParseInLocation(time.RFC3339, tt.t, loc)
		if err != nil {
			t.Fatalf("error parsing date %s: %v", tt.t, err)
		}
		p, err := New(ti, tt.want.duration)
		if err != nil && err.Error() != tt.err.Error() {
			t.Fatalf("error %s: %v", tt.want.location, err)
		}
		if tt.want.begin != p.begin {
			t.Errorf("%s begin: want %d have %d", tt.want.location, tt.want.begin, p.begin)
		}
		if tt.want.duration != p.duration {
			t.Errorf("%s duration: want %d have %d", tt.want.location, tt.want.duration, p.duration)
		}
		if tt.want.location != p.location {
			t.Errorf("%s location: want %s have %s", tt.want.location, tt.want.location, p.location)
		}
	}
}
