package banner

import (
	"database/sql"
	"errors"
	"time"

	"github.com/fgrimme/0ca17a7468266cd599c376f2c522790404ed878f/display"
)

// store handles database operations on banners and holds an in-memory
// cache of banners. The cache is sorted ascendingly by display period.
type store struct {
	db      *sql.DB
	banners []banner // sorted
}

// NewStore returns a store.
// This is the only type to be used in depending modules.
// TODO: the db parameter would either be an interface or a more flexible
// custom type to manage datastore access.
func NewStore(db *sql.DB) (*store, error) {
	s := &store{db: db}
	if err := load(s); err != nil {
		return nil, err
	}
	return s, nil
}

// load initializes a store by loading banners from the database.
// TODO: implement - mocked since this is not part of the coding challenge.
func load(s *store) error {
	// perform database operations to lookup banners
	// sort banners ascendingly by display time
	// assign banners to store
	// ...
	p := []display.Period{
		display.Period{},
	}
	b := []banner{
		banner{
			id:      1,
			content: "FOO BANNER",
			period:  p[0],
		},
	}
	s.banners = b
	return nil
}

// ActiveIn returns the first banner which is active in the location.
// Returns an error if no banner is found.
func (s *store) ActiveIn(location string) (banner, error) {
	var active bool
	var err error
	for _, b := range s.banners {
		active, err = b.activeIn(time.Now(), location)
		if active {
			return b, nil
		}
		if err != nil {
			return banner{}, err
		}
	}
	return banner{}, errors.New("no active banner in this location")
}
