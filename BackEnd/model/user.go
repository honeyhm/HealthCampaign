package model

import (
	"time"

	"gopkg.in/mgo.v2/bson"
)

type User struct {
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

type Users []User

func (u *User) IsValidPassword(password string) bool {
	return u.Password == password
}