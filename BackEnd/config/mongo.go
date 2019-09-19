package config

import(
	_"os"
	mgo "gopkg.in/mgo.v2"
)

//connect to mongo
func GetMongoDB()(*mgo.Database ,error) {

	host :=MongoHost
	//host :="localhost" //MongoHost  //os.Gretenv("MONGO_HOST")
	dbName :=MongoDbName
	//dbName :="temp"//MongoDbName  //os.Gretenv("MONGO_HDB_Name")

	//connections:
	//mongo connection
	session , err := mgo.Dial(host)

	if err != nil {
		return nil ,err
	}

	//db connection
	db := session.DB(dbName)
	return db , nil



}