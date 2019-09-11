package banner

import (
	"database/sql"
	"time"

	"github.com/fgrimme/0ca17a7468266cd599c376f2c522790404ed878f/display"
)

// store handles database operations on banners and holds an in-memory
// cache of banners. The cache is sorted.
type store struct {
	db      *sql.DB
	banners []banner // sorted
}

// NewStore returns a store.
// This is the only type to be used in depending modules.
func NewStore(db *sql.DB) (*store, error) {
	s := &store{db: db}
	if err := load(s); err != nil {
		return nil, err
	}
	return s, nil
}

// load initializes a store by loading banners from the database.
// mocked since not part of the coding challenge.
// FIXME:
// - perform database operations to lookup banners
// - sort banners ascendingly by begin of display period
func load(s *store) error {
	t, err := time.Parse(time.RFC3339, "2019-11-25T02:01:00Z")
	if err != nil {
		return err
	}
	// first banner; duration 12h
	p1, err := display.New(t, 12*60*60)
	b1 := banner{
		id:      1,
		content: "FOO BANNER 1",
		period:  p1,
	}
	// 12h offset to the expiration of first banner; duration 48h
	p2, err := display.New(t.Add(24*time.Hour), 48*60*60)
	b2 := banner{
		id:      2,
		content: "FOO BANNER 2",
		period:  p2,
	}
	// 1 day overlap with the second banner; duration 48h
	p3, err := display.New(t.Add(48*time.Hour), 48*60*60)
	b3 := banner{
		id:      3,
		content: "FOO BANNER 3",
		period:  p3,
	}
	s.banners = []banner{b1, b2, b3}
	return nil
}

// ActiveIn returns a slice of banners which are active in the location.
// Sorted by the earliest starting display period.
func (s *store) ActiveIn(t time.Time, location string, internal bool) ([]banner, error) {
	var active bool
	var err error
	banners := make([]banner, 0)
	for _, b := range s.banners {
		active, err = b.activeIn(t, location, internal)
		if err != nil {
			return nil, err
		}
		if active {
			banners = append(banners, b)
		}
	}
	if internal {
		// sort by epiration time for internal requests
		return s.sortByExpiration(banners), nil
	}
	return banners, nil
}

// FIXME: not implemented due to time restriction
func (s *store) sortByExpiration(b []banner) []banner { return b }

// Create creates a banner for the given period and stores it in the database.
// Returns the banner if successful.
func (s *store) Create(content string, begin time.Time, duration int64) (banner, error) {
	// mocked out but since no data layer is implemented
	return banner{}, nil
}

// Update updates a banner in the database and returns an updated version.
func (s *store) Update(id, duration int64, content string, begin time.Time) (banner, error) {
	// mocked out but since no data layer is implemented
	return banner{}, nil
}

// Delete deletes a banner.
func (s *store) Delete(id int64) (bool, error) {
	// mocked out but since no data layer is implemented
	return false, nil
}
