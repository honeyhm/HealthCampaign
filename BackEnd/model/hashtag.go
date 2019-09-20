package model

import "gopkg.in/mgo.v2/bson"

type Hashtag struct {
	Id         		bson.ObjectId `bson:"_id"`
	TagName       string        `bson:"tag_name"`
}
type Hashtags []Hashtag