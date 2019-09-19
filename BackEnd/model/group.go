package model

import "gopkg.in/mgo.v2/bson"

type Group struct {
	Id         		bson.ObjectId `bson:"_id"`
	GroupName       string        `bson:"group_name"`
	GroupImage      string        `bson:"group_image"`
	MedicalStaffId  []string      `bson:"medical_staff_id"`
	GroupId      	[]string      `bson:"group_id"`
	Creator 	    string        `bson:"creator"`
	MemberNumber 	int        	  `bson:"member_number"`
}
type Groups []Group