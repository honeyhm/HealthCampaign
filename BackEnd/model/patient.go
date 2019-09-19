
package model

import "gopkg.in/mgo.v2/bson"
import "time"

type Patient struct {
	Id         		bson.ObjectId `bson:"_id"`
	PatientName       string        `bson:"patient_name"`
	Country        string        `bson:"country"`
	BirthYear	string     `bson:"birth_year"`
	//UserId         string `db:"userId" bson:"user_id" json:"Token"`
	Gender         bool   `bson:"gender"`
	ProfileImage  string `bson:"profile_image"`
	PhoneNumber   string `bson:"phone_number"`
	CampaignId   []string `bson:"campaign_id"`
	UserName 	 string `db:"userName" bson:"user_name"`
	Password 	 string `db:"password" bson:"password"`
	JoinedAt     time.Time  `bson:"joined_at"`
}
type Patients []Patient