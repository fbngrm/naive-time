package display

import (
	"testing"
	"time"
)

var naiveTests = []struct {
	ts   int64  // UTC-offset aware timestamp
	loc  string // IANA location ID
	want int64  // UTC-offset naive timestamp
	err  string // expected error message
}{
	{
		ts:   1112493660, // 2005-04-03 02:01:00 +00:00 UTC
		loc:  "UTC",
		want: 1112493660, // 2005-04-03 02:01:00 +00:00 UTC
	},
	{
		ts:   1112493660,   // 2005-04-03 11:01:00 +09:00 JST
		loc:  "Asia/Tokyo", // 32400 seconds east of UTC
		want: 1112526060,   // 2005-04-03 11:01:00 +0000 UTC
	},
	{
		ts:   1112493660,      // 2005-04-03 04:01:00 +02:00 JST
		loc:  "Europe/Berlin", // 7200 seconds east of UTC
		want: 1112500860,      // 2005-04-03 04:01:00 +0000 UTC
	},
	{
		ts:   1112493660,            // 2005-04-02 18:01:00 -0800 PST
		loc:  "America/Los_Angeles", // -28800 seconds east of UTC
		want: 1112464860,            // 2005-04-02 18:01:00 +0000 UTC
	},
	{
		ts:   1112493660,       // 2005-04-02 15:01:00 -1100 SST
		loc:  "Pacific/Midway", // -39600 seconds east of UTC; date line
		want: 1112454060,       // 2005-04-02 15:01:00 +0000 UTC
	},
	{
		ts:   1112493660,     // 2005-04-02 15:01:00 -1100 SST
		loc:  "invalid/zone", // -39600 seconds east of UTC; date line
		want: 0,              // unix epoch
		err:  "unknown time zone invalid/zone",
	},
}

func TestNaiveTime(t *testing.T) {
	for _, tt := range naiveTests {
		ti := time.Unix(tt.ts, 0).UTC()
		loc, err := time.LoadLocation(tt.loc)
		if err != nil {
			if err.Error() != tt.err {
				t.Fatalf("error loading location %s: %v", tt.loc, err)
			}
		} else {
			// only parse the location when we don't expect an invalid location
			ti = ti.In(loc)
		}
		nt, err := naiveTime(ti, tt.loc)
		if err != nil && err.Error() != tt.err {
			t.Fatalf("error %s: %v", tt.loc, err)
		}
		if tt.want != nt.timestamp {
			t.Errorf("%s: want %d have %d", tt.loc, tt.want, nt.timestamp)
		}
	}
}
