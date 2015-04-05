package modules

import (
	"log"
	"time"

	"labix.org/v2/mgo"
)

const (
	MongoDBHost  = "example.mongo.com:35428"
	AuthDatabase = "asio"
	AuthUsername = "guest"
	AuthPassword = "youknowwhatis"
	TestDatabase = "asio"
)

func CreateDatabaseSession() {
	mongoDbDialInfo := &mgo.DialInfo{
		Addrs:    []string{MongoDBHost},
		Timeout:  60 * time.Second,
		Database: AuthDatabase,
		Username: AuthUsername,
		Password: AuthPassword,
	}
	// Create a session
	mongoSession, err := mgo.DialWithInfo(mongoDbDialInfo)
	if err != nil {
		log.Fatalf("CreateSession: %s\n", err)
	}

	// http://godoc.org/labix.org/v2/mgo#Session.SetMode
	mongoSession.SetMode(mgo.Monotonic, true)
}
