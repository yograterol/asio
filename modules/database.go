package modules 


import (
	"errors"
	"labix.org/v2/mgo"
	"time"
	"log"
}


const (
	MongoDBHost  = "localhost:27017"
	AuthDatabase = "asio"
	AuthUserName = ""
	AuthPassword = ""
	TestDatabase = "asio"
)

func CreateDatabaseSession() {
	mongoDBDialInfo := &mgo.DialInfo {
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
	
	return mongoSession
}