package modules

import "labix.org/v2/mgo"

func CreateDatabaseSession(MongoDBURI, Database string) *mgo.Database {
	mongoSession, err := mgo.Dial(MongoDBURI)
	if err != nil {
		panic(err)
	}
	return mongoSession.DB(Database)
}
