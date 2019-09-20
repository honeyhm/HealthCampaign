package model

import "gopkg.in/mgo.v2/bson"

type Search struct {
	Id         		bson.ObjectId `bson:"_id"`
	HashTagId       string        `bson:"hashtag_id"`
	ReferId 	    string        `bson:"refer_id"`
	ReferState   	bool          `bson:"refer_state"`
}
type Searches []Search