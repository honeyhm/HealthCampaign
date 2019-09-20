package model

import "gopkg.in/mgo.v2/bson"

type Message struct {
	Id         		bson.ObjectId `bson:"_id"`
	MessageText     string        `bson:"group_name"`
	Sender      	string        `bson:"group_image"`
	ParentId 	    string        `bson:"parent_id"`
	messageId 	    string        `bson:"message_id"`
}
type Messages []Message