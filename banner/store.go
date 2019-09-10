package banner

import (
	"github.com/fgrimme/mercari/datastore"
)

// store handles database operations on banners and holds an in-memory
// cache of banners. The cache is sorted ascendingly by display period.
type store struct {
	db      *datastore.DB
	banners []banner // sorted
}

// NewStore returns a store.
// This is the only type to be used in depending modules.
func NewStore(db *datastore.DB) (*store, error) {
	s := &store{db: db}
	if err := load(s); err != nil {
		return nil, err
	}
	return s, nil
}

// load initializes a store by loading banners from the database.
func load(s *store) error {
	// perform database operations to lookup banners
	// sort banners ascendingly by display time
	// assign banners to store
	// ...
	return nil
}
