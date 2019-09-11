package banner

import (
	"testing"
	"time"
)

type testPeriod struct {
	begin    int64  // seconds since unix epoch
	duration int64  // seconds banner should be displayed
	location string // IANA location ID
}
type testBanner struct {
	id      int64
	content string
	p       testPeriod
}

var testBanners = []testBanner{
	{
		id:      1,
		content: "FOO BANNER 1",
		p: testPeriod{
			begin:    1574647260,   // "2019-11-25T02:01:00Z"
			duration: 12 * 60 * 60, // 12h
			location: "UTC",
		},
	},
	{ // 12h offset to the expiration to banner 1; duration 48h
		id:      2,
		content: "FOO BANNER 2",
		p: testPeriod{
			begin:    1574733660,   // "2019-11-26T02:01:00Z"
			duration: 48 * 60 * 60, // 48h
			location: "UTC",
		},
	},
	{ // 1 day overlap with the banner 2; duration 48h
		id:      3,
		content: "FOO BANNER 3",
		p: testPeriod{
			begin:    1574820060, // "2019-11-27T02:01:00Z"
			duration: 48 * 60 * 60,
			location: "UTC",
		},
	},
}

var testActiveIn = []struct {
	d string       // description
	t int64        // request time
	l string       // request location
	e error        // expected error
	b []testBanner // expected active banners
}{
	{
		d: "request time equals begin of display time",
		t: 1574647260, // "2019-11-25T02:01:00Z",
		l: "UTC",
		b: []testBanner{testBanners[0]},
	},
	{
		d: "request time exceeds display time",
		t: 1574690460, // "2019-11-25T14:01:00Z",
		l: "UTC",
		b: []testBanner{},
	},
	{
		d: "request time 1h after begin of display time",
		t: 1574733660, // "2019-11-26T03:01:00Z"
		l: "UTC",
		b: []testBanner{testBanners[1]},
	},
	{
		d: "request time overlaps, choose earlier display begin",
		t: 1574820060, // "2019-11-27T02:01:00Z"
		l: "UTC",
		b: []testBanner{testBanners[1], testBanners[2]},
	},
	// TODO: test errors
}

func TestActiveIn(t *testing.T) {
	// FIXME: we pass nil since this is mocked out in coding challenge
	store, err := NewStore(nil)
	if err != nil {
		t.Fatalf("error loading store: %v", err)
	}
	for _, tt := range testActiveIn {
		banners, err := store.ActiveIn(time.Unix(tt.t, 0), tt.l)
		if err != nil && err.Error() != tt.e.Error() {
			t.Fatalf("%s: %v", tt.d, err)
		}
		if len(banners) != len(tt.b) {
			t.Fatalf("%s: want %d active banner(s) have %d", tt.d, len(tt.b), len(banners))
		}
		for i, b := range banners {
			if b.id != tt.b[i].id {
				t.Errorf("%s: want id %d have %d", tt.d, tt.b[i].id, b.id)
			}
			if b.content != tt.b[i].content {
				t.Errorf("%s: want content %s have %s", tt.d, tt.b[i].content, b.content)
			}
		}
	}
}
