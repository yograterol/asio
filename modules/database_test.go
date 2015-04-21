package modules

import (
	"testing"
)

func TestDatabaseSession(t *testing.T) {
	MongoDBURI := "mongodb://localhost"
	MongoDBURI_invalid := "mongodb://localhost_no_valid"
	Database := "asio"

	DB := CreateDatabaseSession(MongoDBURI, Database)
	if DB.Name != Database {
		t.Fatal("MongoDB is not available.")
	}

	defer func() {
		if r := recover(); r != nil {
			t.Log("MongoDB with invalid URI was recovered.")
		}
	}()
	DB = CreateDatabaseSession(MongoDBURI_invalid, Database)
}
