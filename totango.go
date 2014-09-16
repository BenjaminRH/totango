// Package totango provides bindings for the Totango server side integration API
package totango

import (
	"errors"
	"net/http"
	"net/url"
)

const (
	// The base URL for all Totango API requests
	baseURL = "https://sdr.totango.com/pixel.gif/?sdr_s="
)

// The data for an API request
type request struct {
	accountID   string
	accountName string
	userName    string
	activity    string
	module      string
	attributes  map[string]string
}

func encode(s string) string {
	return url.QueryEscape(s)
}

// Construct a URL query param string from the provided fields
func (r *request) String() string {
	var url string

	switch {
	case r.accountID != "":
		url += "&sdr_o=" + encode(r.accountID)
		fallthrough
	case r.accountName != "":
		url += "&sdr_odn=" + encode(r.accountName)
		fallthrough
	case r.userName != "":
		url += "&sdr_u=" + encode(r.userName)
		fallthrough
	case r.activity != "":
		url += "&sdr_a=" + encode(r.activity)
		fallthrough
	case r.module != "":
		url += "&sdr_m=" + encode(r.module)
		fallthrough
	case len(r.attributes) > 0:
		for name, val := range r.attributes {
			url += "&sdr_o." + encode(name) + "=" + encode(val)
		}
	}

	return url
}

type Tracker struct {
	serviceID string
}

// Construct a Totango API request from a request type
func (t *Tracker) getURL(r *request) string {
	return baseURL + t.serviceID + r.String()
}

func NewTracker(serviceID string) (*Tracker, error) {
	if serviceID == "" {
		return nil, errors.New("Tracker requires a valid Totango Service ID")
	}

	return &Tracker{serviceID: serviceID}, nil
}

func (t *Tracker) Track(accountID, accountName, userName, activity, module string) (*http.Response, error) {
	r := &request{
		accountID:   accountID,
		accountName: accountName,
		userName:    userName,
		activity:    activity,
		module:      module,
	}

	return http.Get(t.getURL(r))
}

func (t *Tracker) TrackAttribute(accountID, userName, name, value string) (*http.Response, error) {
	if accountID == "" {
		return nil, errors.New("An account ID is required to track an attribute")
	}

	r := &request{
		accountID:  accountID,
		userName:   userName,
		attributes: map[string]string{name: value},
	}

	return http.Get(t.getURL(r))
}

func (t *Tracker) TrackAttributes(accountID, userName string, attributes map[string]string) (*http.Response, error) {
	if accountID == "" {
		return nil, errors.New("An account ID is required to track attributes")
	}

	r := &request{
		accountID:  accountID,
		userName:   userName,
		attributes: attributes,
	}

	return http.Get(t.getURL(r))
}
