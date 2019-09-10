package banner

import "github.com/fgrimme/mercari/datastore"

// store handles database operations on banners and holds an in-memory
// cache of banners.
type store struct {
	db      *datastore.DB
	banners []banner
}

// NewStore returns a store.
// It is the only type to be used in depending modules.
func NewStore(db *datastore.DB) *store {
	return &store{db: db}
}

// init initializes a store by loading banners from the database.
func (s *store) init() error {
	// perform some database operations
	// ...
	return nil
}

func (s *store) refresh() error           {}
func (s *store) create() (*banner, error) {}
func (s *store) get() (*banner, error)    {}
