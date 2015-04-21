// domain model
package models

import (
	"errors"

	"labix.org/v2/mgo"
	"labix.org/v2/mgo/bson"
)

type Domain struct {
	ID      bson.ObjectId  `bson:"_id,omitempty" json:"id"`
	Name    string         `bson:"name"          json:"name"`
	URL     string         `bson:"url"           json:"url"`
	Setting *DomainSetting `bson:"settings"      json:"settigns"`
	// Owner
}

func (d *Domain) Validate() bool {
	return len(d.Name) > 0 && len(d.URL) > 0
}

func (d *Domain) Save(DB *mgo.Database) error {
	collection := DB.C("domain")
	if d.Validate() {
		return collection.Insert(&d)
	}
	return errors.New("The document is not valid. (Domain model).")
}

type DomainSetting struct {
	Interval uint8 `bson:"interval"      json:"interval"` // Interval for each ping in seconds. min: 60s max: 180s
	Retry    uint8 `bson:"retry"         json:"retry"`    // Amount times for retry a ping to domain. min: 1 max: 5
	Timeout  uint8 `bson:"timeout"       json:"timeout"`  // Timeout for each retry ping. min: 5s max: 20s
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
