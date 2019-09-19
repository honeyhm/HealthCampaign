/*created by H.Mlk*/
package model

import (
	"gopkg.in/mgo.v2/bson"
	"time"
)

type Token struct{
	Id			bson.ObjectId  		`db:"id" bson:"_id"`
	PersonId		string   	` bson:"person_id"`
	StudentNum		string				`bson:"student_num"`
	EnterTime	time.Time 				`db:"isActive" bson:"enter_time"`
}


type Tokens []Token
