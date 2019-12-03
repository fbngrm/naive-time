package main

import (
	"fmt"
	"time"

	"github.com/fgrimme/naive-time/banner"
)

func main() {
	// Initialize a database or an abstraction to pass to the banner store.
	// db, err := sql.Open(driverName, dsn)
	// if err != nil {
	//  panic(err)
	// }

	// Pass the database abstraction instead of nil.
	store, err := banner.NewStore(nil)
	if err != nil {
		panic(err)
	}

	// matches the mock data; normally time.Now() with the request timezone
	// would be used.
	t, err := time.Parse(time.RFC3339, "2019-11-26T02:01:00Z")
	if err != nil {
		panic(err)
	}

	// Request with external IP address. The IP restriction/allowance
	// checks would be performed in an earlier stage/middleware and would
	// be passed in a request context or flag.
	internal := false

	// An external request from timezone "Asia/Tokyo" with the mocked time.
	// We check if there are active banners for the current local wall clock
	// time in this timezone. The result is a slice of all active banners,
	// sorted by earliest display period.
	bExternal, err := store.ActiveIn(t, "Asia/Tokyo", internal)
	if err != nil {
		panic(err)
	}
	fmt.Println(bExternal) // TODO: add human readable representation of period

	internal = true
	// An internal request from timezone "Asia/Tokyo" with the mocked time.
	// We check if there are active banners in the future. The result is a
	// slice of all non-expired banners, sorted by earliest expiration date.
	bInternal, err := store.ActiveIn(t, "Asia/Tokyo", internal)
	if err != nil {
		panic(err)
	}
	fmt.Println(bInternal)
}
