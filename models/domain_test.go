package models

import (
	"testing"

	. "github.com/Clowl/asio/modules"
	"labix.org/v2/mgo"
)

const (
	MongoDBURI = "mongodb://localhost"
	Database   = "asio"
)

var DB *mgo.Database

func init() {
	DB = CreateDatabaseSession(MongoDBURI, Database)
}

func TestDomainSetting(t *testing.T) {
	ds, _ := NewDomainSetting(60, 5, 10)
	if ds.Interval != 60 || ds.Retry != 5 || ds.Timeout != 10 {
		t.Fatal("Problem with DomainSetting model.")
	}
	ds_invalid, _ := NewDomainSetting(250, 250, 250)
	if ds_invalid.Interval != 0 || ds_invalid.Retry != 0 || ds_invalid.Timeout != 0 {
		t.Fatal("Problem with DomainSetting validate.")
	}
}

func TestDomain(t *testing.T) {
	ds, _ := NewDomainSetting(60, 5, 10)
	d := NewDomain("Test.asio", "https://test.asio", &ds)
	err := d.Save(DB)

	if err != nil {
		t.Fatal(err)
	}

	d_invalid := NewDomain("", "", &ds)
	err = d_invalid.Save(DB)
	if err == nil {
		t.Fatal("Failed the Domain model validate.")
	}
}
