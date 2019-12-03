package display

import (
	"errors"
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
			begin:    &naivetime{timestamp: 1112493660}, // 2005-04-03 02:01:00 +00:00 UTC
			duration: 60 * 60,                           // display period 1h
			location: "UTC",
		},
	},
	{
		t: "2005-04-03T11:01:00+09:00",
		want: Period{
			begin:    &naivetime{timestamp: 1112526060}, // 2005-04-03 11:01:00 +0000 UTC
			duration: 12 * 60 * 60,                      // display period 12h
			location: "Asia/Tokyo",
		},
	},
	{
		t: "2005-04-03T04:01:00+02:00",
		want: Period{
			begin:    &naivetime{timestamp: 1112500860}, // 2005-04-03 04:01:00 +0000 UTC
			duration: 24 * 60 * 60,                      // display period 24h
			location: "Europe/Berlin",
		},
	},
	{
		t: "2005-04-02T18:01:00-08:00",
		want: Period{
			begin:    &naivetime{timestamp: 1112464860}, // 2005-04-02 18:01:00 +0000 UTC
			duration: 60,                                // display period 1m
			location: "America/Los_Angeles",
		},
	},
	{
		t: "2005-04-02T15:01:00-11:00",
		want: Period{
			begin:    &naivetime{timestamp: 1112454060}, // 2005-04-02 15:01:00 +0000 UTC
			duration: 48 * 60 * 60,                      // display period 48h
			location: "Pacific/Midway",
		},
	},
}

func TestNew(t *testing.T) {
	for _, tt := range testPeriod {
		// parse the test date
		loc, err := time.LoadLocation(tt.want.location)
		if err != nil {
			t.Fatalf("error %s: %v", tt.want.location, err)
		}
		ti, err := time.ParseInLocation(time.RFC3339, tt.t, loc)
		if err != nil {
			if err.Error() != tt.e.Error() {
				t.Fatalf("error %s: %v", tt.t, err)
			}
			ti = time.Time{}
		}
		// test
		p, err := New(ti, tt.want.duration)
		if err != nil && err.Error() != tt.e.Error() {
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

type active struct {
	b bool
	e error
}

var testActive = []struct {
	d    string // test case description
	t    string // UTC-offset aware time
	l    string // IANA location name
	p    Period // location is empty; only relevant for initialization
	want active // expected output
}{
	{ // UTC
		d: "not yet active in UTC",
		t: "2005-04-03T02:00:59Z", // one second before begin
		l: "UTC",
		p: Period{
			begin:    &naivetime{timestamp: 1112493660}, // 2005-04-03 02:01:00 +00:00 UTC
			duration: 60 * 60,                           // display period 1h
		},
		want: active{
			b: false,
			e: nil,
		},
	},
	{
		d: "active in UTC",
		t: "2005-04-03T02:01:00Z", // first second of display period
		l: "UTC",
		p: Period{
			begin:    &naivetime{timestamp: 1112493660}, // 2005-04-03 02:01:00 +00:00 UTC
			duration: 60 * 60,                           // display period 1h
		},
		want: active{
			b: true,
			e: nil,
		},
	},
	{
		d: "expired in UTC",
		t: "2005-04-03T03:01:00Z", // expired by one second
		l: "UTC",
		p: Period{
			begin:    &naivetime{timestamp: 1112493660}, // 2005-04-03 02:01:00 +00:00 UTC
			duration: 60 * 60,                           // display period 1h
		},
		want: active{
			b: false,
			e: nil,
		},
	},
	{ // Asia/Tokyo
		d: "not yet active in Asia/Tokyo",
		t: "2005-04-02T13:01:00+09:00", // 12h before display period
		l: "Asia/Tokyo",
		p: Period{
			begin:    &naivetime{timestamp: 1112526060}, // 2005-04-03 11:01:00 +0000 UTC
			duration: 48 * 60 * 60,                      // display period 48h
		},
		want: active{
			b: false,
			e: nil,
		},
	},
	{
		d: "active in Asia/Tokyo",
		t: "2005-04-04T11:01:00+09:00", // 24h after begin of display period
		l: "Asia/Tokyo",
		p: Period{
			begin:    &naivetime{timestamp: 1112526060}, // 2005-04-03 11:01:00 +0000 UTC
			duration: 48 * 60 * 60,                      // display period 48h
		},
		want: active{
			b: true,
			e: nil,
		},
	},
	{
		d: "expired in Asia/Tokyo",
		t: "2005-05-03T11:01:00+09:00", // expired by one month
		l: "Asia/Tokyo",
		p: Period{
			begin:    &naivetime{timestamp: 1112526060}, // 2005-04-03 11:01:00 +0000 UTC
			duration: 48 * 60 * 60,                      // display period 48h
		},
		want: active{
			b: false,
			e: nil,
		},
	},
	{ // Pacific/Midway
		d: "not yet active in Pacific/Midway",
		t: "2004-04-02T15:01:00-11:00", // 1 year before display period
		l: "Pacific/Midway",
		p: Period{
			begin:    &naivetime{timestamp: 1112454060}, // 2005-04-02 15:01:00 +0000 UTC
			duration: 30 * 24 * 60 * 60,                 // display period 30d
		},
		want: active{
			b: false,
			e: nil,
		},
	},
	{
		d: "active in Pacific/Midway",
		t: "2005-05-02T15:00:59-11:00", // last second of display period
		l: "Pacific/Midway",
		p: Period{
			begin:    &naivetime{timestamp: 1112454060}, // 2005-04-02 15:01:00 +0000 UTC
			duration: 30 * 24 * 60 * 60,                 // display period 30d
		},
		want: active{
			b: true,
			e: nil,
		},
	},
	{ // expect error
		d: "invalid location",
		t: "2005-05-02T15:01:00-11:00",
		l: "invalid/location", // expect error
		p: Period{
			begin:    &naivetime{timestamp: 0}, // 1970-01-01 0:00:00 +0000 UTC
			duration: 30 * 24 * 60 * 60,        // display period 30d
		},
		want: active{
			b: false,
			e: errors.New("unknown time zone invalid/location"),
		},
	},
}

func TestActive(t *testing.T) {
	for _, tt := range testActive {
		// parse the test date
		ti := time.Time{}
		loc, err := time.LoadLocation(tt.l)
		if err != nil {
			if err.Error() != tt.want.e.Error() { // unexpected error
				t.Fatalf("%s: %v", tt.d, err)
			}
		} else {
			ti, err = time.ParseInLocation(time.RFC3339, tt.t, loc)
			if err != nil {
				t.Fatalf("error parsing date %s: %v", tt.t, err)
			}
		}
		// test
		active, err := tt.p.Active(ti, tt.l, false)
		if err != nil && err.Error() != tt.want.e.Error() {
			t.Fatalf("case %s: %v", tt.d, err)
		}
		if active != tt.want.b {
			t.Errorf("%s: want %v got %v", tt.d, tt.want.b, active)
		}
	}
}
