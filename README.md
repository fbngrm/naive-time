## General assumptions

 - Calendrical calculations assume Gregorian calendar, with no leap
   seconds.
 - Relies on the IANA timezone database or equivalents to be
   accessible on the system executing the program. 
 - Definition of
   timezone is, banners should be displayed in all timezones at the same
   local wall clock time of the timezone. E.g. at 6pm - 7pm local time
   in each timezone.


## Terminology
The following terminology is used consistently throughout the program:

Wall Clock

 - Subject of changes for clock synchronization, e.g. DST.
 - Used for telling time.

Monotonic Clock
 - *not* subject of changes for clock synchronization.
 - Used for time calculations.

Offset-aware Time

 - Does have an UTC-offset.

Offset-naive Time

 - Does *not* have an UTC-offset.


## Architecture
The program is designed by the inversion of control/dependency injection principle. All dependencies a component needs, are supplied as parameters at creation time. Thus, components can be used in a modular way, have loose coupling and are more easy to test by providing mock data. Also, it is possible to ensure that we either have a ready to use component after creation or return an error.

The `banner.Store` component module is exported/exposed to the outside and provides the API for depending modules. All time calculations are encapsulated in the `display` module which is designed as a library to abstract timezone aware time periods. It is capable to handle potential future changes in the timezone conversion rules in the IANA database.

## Location
The program relies on an input location in the IANA zone database format, e.g.

 - Asia/Tokyo
 - Europe/Berlin


A locations time is the local wall clock time at the location

## Tests
Tests are more fine grained in the `display` module since this encapsulates all time related operations. Therefore, all functions are tested with different timezones and offsets. In the `banner` module, only exported functions that are not mocked out are tested. To avoid boilerplate code in loading testdata locations, only UTC is used in this module.

Run tests from project root directory:


    go test ./...


## API Example
Basic usage of the API is demonstrated in main.go. This makes used of mock-data for banners as well as inputs. To run the example from the projects root directory:


    go run main.go


## Storing time values of future events
It is often stated that storing time values in UTC is a fail proof and convenient approach. This may be true for storing and handling past time values. It is not true for instants in the future when there is a timezone conversion involved. The IANA(or other) timezone database rules for conversion may change which can lead to an undesired shift when converting UTC times back to a another timezone. This program is capable to deal with changes by using offset-naive representations combined with a timezone information.
