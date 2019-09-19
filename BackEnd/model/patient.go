
package model

import "gopkg.in/mgo.v2/bson"
import "time"

type Patient struct {
	Id         bson.ObjectId `bson:"_id"`
	PatientName       string        `bson:"patientName"`
	Country        string        `bson:"country"`
	//StartTime  time.Time     `bson:"start_time"`
	//FinishTime time.Time     `bson:"finish_time"`
	FinishDate time.Time     `bson:"finish_date"`
	Id bson.ObjectId `db:"id" bson:"_id"` //json:"Token"`
	UserId         string `db:"userId" bson:"user_id" json:"Token"`
	FirstName      string `db:"firstName" bson:"first_name"`
	LastName       string `db:"firstName" bson:"last_name"`
	Gender         bool   `bson:"gender"`
	SignatureImage string `bson:"signature_image"`
	UserName string `db:"userName" bson:"user_name"`
	Password string `db:"password" bson:"password"`
	CreatedAt    time.Time  `bson:"created_at"`
	UpdatedAt    time.Time  `bson:"updated_at"`
	UserToken    string     `bson:"user_token"`
}
type Patients []Patient