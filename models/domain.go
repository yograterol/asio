// domain model
package models

import (
	"errors"
)

type Domain struct {
	Name string
	URL  string
	// Owner
}

type DomainSetting struct {
	Domain   *Domain
	Interval uint8 // Interval for each ping in seconds. min: 60s max: 180s
	Retry    uint8 // Amount times for retry a ping to domain. min: 1 max: 5
	Timeout  uint8 // Timeout for each retry ping. min: 5s max: 20s
}

func (ds *DomainSetting) Validate() bool {
	return (ds.Timeout >= 5 && ds.Timeout <= 20) &&
		(ds.Retry >= 1 && ds.Retry <= 5) &&
		(ds.Interval >= 60 && ds.Interval <= 180)
}

func NewDomain(Name, URL string) Domain {
	return Domain{Name: Name, URL: URL}
}

func NewDomainSetting(Domain *Domain, Interval, Retry, Timeout uint8) (DomainSetting, error) {
	ds := DomainSetting{Domain: Domain, Interval: Interval, Retry: Retry, Timeout: Timeout}
	if ds.Validate() {
		return ds, nil
	}
	return DomainSetting{}, errors.New("The settings for the domain is not valid.")
}
