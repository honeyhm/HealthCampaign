package model

import "gopkg.in/mgo.v2/bson"

type Disease struct {
	Id         		  bson.ObjectId `bson:"_id"`
	DiseaseName       string        `bson:"disease_name"`
	OrganName		  string        `bson:"organ_name"`

}
type Diseases []Disease
