# totango

Library for the Totango [server integration API](http://help.totango.com/installing-totango/quick-start-http-api-server-side-integration/)

## Example

This example is available at http://play.golang.org/p/nxXPMIIzDo

```go
package main

import (
	"github.com/BenjaminRH/totango"
)

func main() {
	tracker := totango.NewTracker("SP-XXXX-YY")
	
	// Track a new event
	resp, err := tracker.Track("account123", "userbob@example.com", "User Bob", "Some Activity", "A Module")
	
	// Now update an attribute
	resp, err = tracker.TrackAttribute("account123", "userbob@example.com", "Some Attribute", "The Value")
	
	// How about multiple attributes?
	resp, err = tracker.TrackAttributes("account123", "userbob@example.com", map[string]string{
		"An Attribute": "The value",
		"Another one": "value",
		"Maybe a third": "value",
	})
}
```
