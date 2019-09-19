package model

import "gopkg.in/mgo.v2/bson"
import "time"

type MedicalStaff struct {
	Id         		bson.ObjectId `bson:"_id"`
	StaffName       string        `bson:"staff_name"`
	Subject        	string        `bson:"subject"`
	ExpertKnowledge int           `bson:"expert_knowledge"` /// 0 : omoomi , 1: maghzoAsaab , 2: cheshm  , 3:zibayii ,
	//4 :zanan , 5 : goosho halgho bini , 6 :dakheli ,7 : ortoped
	Gender         	bool   		  `bson:"gender"`
	ProfileImage  	string        `bson:"profile_image"`
	PhoneNumber   	string        `bson:"phone_number"`
	UserName 	 	string        `db:"userName" bson:"user_name"`
	Password 	 	string        `db:"password" bson:"password"`
	JoinedAt     	time.Time     `bson:"joined_at"`
	MedicalCenterId string        `bson:"medical_center_id"`
}
type MedicalStaffs []MedicalStaff