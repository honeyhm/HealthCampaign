
package model

import "gopkg.in/mgo.v2/bson"
import "time"

//Eshel !
type Person struct {
	Id         bson.ObjectId `bson:"_id"`
	Name       string        `bson:"name"`
	StudentNum string        `bson:"student_num"`
	Ssn        string        `bson:"ssn"`
	Gender        string        `bson:"gender"`
	//StartTime  time.Time     `bson:"start_time"`
	//FinishTime time.Time     `bson:"finish_time"`
	FinishDate time.Time     `bson:"finish_date"`
	// ActionArray				[]string		`bson:"action_array"`
}
type Persons []Person