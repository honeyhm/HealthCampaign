package model

import "gopkg.in/mgo.v2/bson"

type MedicalCenter struct {
	Id         		bson.ObjectId `bson:"_id"`
	CenterName      string        `bson:"center_name"`
	Type        	string        `bson:"type"`
	Address 		string        `bson:"address"`
	Gender         	bool   		  `bson:"gender"`
	Province  		string        `bson:"profile_image"` //ostan
	City   			string        `bson:"city"`
	Country 	 	string  	  `bson:"city"`

}
type MedicalCenters []MedicalCenter