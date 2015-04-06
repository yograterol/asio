// domain model
package models

import (
	"errors"
)

type Domain struct {
	Name    string         `bson:Name`
	URL     string         `bson:URL`
	Setting *DomainSetting `bson:settings`
	// Owner
}

type DomainSetting struct {
	Interval uint8 `bson:interval` // Interval for each ping in seconds. min: 60s max: 180s
	Retry    uint8 `bson:retry`    // Amount times for retry a ping to domain. min: 1 max: 5
	Timeout  uint8 `bson:timeout`  // Timeout for each retry ping. min: 5s max: 20s
}

// Check if DomainSetting instance accomplish all restrictions.
func (ds *DomainSetting) Validate() bool {
	return (ds.Timeout >= 5 && ds.Timeout <= 20) &&
		(ds.Retry >= 1 && ds.Retry <= 5) &&
		(ds.Interval >= 60 && ds.Interval <= 180)
}

// Create a new Domain instance
func NewDomain(Name, URL string, Setting *DomainSetting) Domain {
	return Domain{Name: Name, URL: URL, Setting: Setting}
}

// Create a new DomainSetting instance
func NewDomainSetting(Interval, Retry, Timeout uint8) (DomainSetting, error) {
	ds := DomainSetting{Interval: Interval, Retry: Retry, Timeout: Timeout}
	if ds.Validate() {
		return ds, nil
	}
	
	return DomainSetting{}, errors.New("The settings for the domain is not valid.")
}