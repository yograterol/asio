<<<<<<< HEAD
package modules
=======
package modules
>>>>>>> 1b81e1dbb9c7e44c48bb68048eb9f950584f589a


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
<<<<<<< HEAD
		Password: AuthPassword,
=======
		Password: AuthPassword,
>>>>>>> 1b81e1dbb9c7e44c48bb68048eb9f950584f589a
	}
	// Create a session
	mongoSession, err := mgo.DialWithInfo(mongoDbDialInfo)
	if err != nil {
		log.Fatalf("CreateSession: %s\n", err)
	}
<<<<<<< HEAD

	// http://godoc.org/labix.org/v2/mgo#Session.SetMode
	mongoSession.SetMode(mgo.Monotonic, true)

	// Create a wait group to manage the goroutines.
	var waitGroup sync.WaitGroup

=======

	// http://godoc.org/labix.org/v2/mgo#Session.SetMode
	mongoSession.SetMode(mgo.Monotonic, true)

	// Create a wait group to manage the goroutines.
	var waitGroup sync.WaitGroup

>>>>>>> 1b81e1dbb9c7e44c48bb68048eb9f950584f589a
	// Perform 10 concurrent queries against the database.
	waitGroup.Add(10)
	for query := 0; query < 10; query++ {
		go RunQuery(query, &waitGroup, mongoSession)
	}
<<<<<<< HEAD

	// Wait for all the queries to complete.
	waitGroup.Wait()
	log.Println("All Queries Completed")
}
=======

	// Wait for all the queries to complete.
	waitGroup.Wait()
	log.Println("All Queries Completed")
}
>>>>>>> 1b81e1dbb9c7e44c48bb68048eb9f950584f589a
