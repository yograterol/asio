package modules


import (
	"errors"
	"labix.org/v2/mgo"
	"time"
	"log"
)

const {
	MongoDBHost  = "example.mongo.com:35428"
    AuthDatabase = "asio"
	AuthUserName = "guest"
	AuthPassword = "youknowwhatis"
	TestDatabase = "asio"
}

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

	// http://godoc.org/labix.org/v2/mgo#Session.SetMode
	mongoSession.SetMode(mgo.Monotonic, true)

	// Create a wait group to manage the goroutines.
	var waitGroup sync.WaitGroup

	// Perform 10 concurrent queries against the database.
	waitGroup.Add(10)
	for query := 0; query < 10; query++ {
		go RunQuery(query, &waitGroup, mongoSession)
	}

	// Wait for all the queries to complete.
	waitGroup.Wait()
	log.Println("All Queries Completed")
}
