/*created by H.Mlk*/
/*last modified : */
package model

import (
	"time"

	"gopkg.in/mgo.v2/bson"
)

// the sign picture should be added
type User struct { //// User model must be modified later!
	Id bson.ObjectId `db:"id" bson:"_id"` //json:"Token"`
	//Role			string
	UserId         string `db:"userId" bson:"user_id" json:"Token"`
	FirstName      string `db:"firstName" bson:"first_name"`
	LastName       string `db:"firstName" bson:"last_name"`
	Gender         bool   `bson:"gender"`
	SignatureImage string `bson:"signature_image"`
	//ZoneCode		string
	UserName string `db:"userName" bson:"user_name"`
	Password string `db:"password" bson:"password"`
	//PositionId		string
	//ZoneId			string
	OrgPosZoneId string     `bson:"org_pos_zone_id"`
	CreatedAt    time.Time  `bson:"created_at"`
	UpdatedAt    time.Time  `bson:"updated_at"`
	UserToken    string     `bson:"user_token"`

}

type Users []User

func (u *User) IsValidPassword(password string) bool {
	return u.Password == password
}
