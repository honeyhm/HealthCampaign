
package model

import "gopkg.in/mgo.v2/bson"

type Alteration struct { //// patient height and weight log

	Id         			 bson.ObjectId `bson:"_id"`
	PatientId       	 string        `bson:"patient_id"`
	PatientHeight        string        `bson:"patient_height"`
	PatientWeight        string        `bson:"patient_weight"`
}

type Alterations []Alteration