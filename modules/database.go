package modules

import (
	"log"
	"time"

	"labix.org/v2/mgo"
)

const (
	MongoDBHost  = "localhost:27017"
	AuthDatabase = "asio"
	AuthUserName = ""
	AuthPassword = ""
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
<<<<<<< HEAD

	// http://godoc.org/labix.org/v2/mgo#Session.SetMode
	mongoSession.SetMode(mgo.Monotonic, true)
}
=======

	return mongoSession
}
>>>>>>> d9b99e8c410611eb8c22cefb3627a0111ab541cf
