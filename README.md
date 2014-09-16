# totango

Library for the Totango [server integration API](http://help.totango.com/installing-totango/quick-start-http-api-server-side-integration/)

## Example

This example is available at http://play.golang.org/p/pgc709-CnQ

```go
package main

import (
	"github.com/BenjaminRH/totango"
)

func main() {
	tracker, _ := totango.NewTracker("SP-XXXX-YY")
	
	// Track a new event
	tracker.Track("account123", "userbob@example.com", "User Bob", "Some Activity", "A Module")
	
	// Now update an account attribute
	tracker.TrackAttribute("account123", "", "Some Attribute", "The Value")
	
	// Or update a user attribute
	tracker.TrackAttribute("account123", "userbob@example.com", "Foo", "Bar")
	
	// How about multiple attributes?
	tracker.TrackAttributes("account123", "userbob@example.com", map[string]string{
		"An Attribute": "The value",
		"Another one": "value",
		"Maybe a third": "value",
	})
}
```
