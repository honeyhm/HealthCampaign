
package repository

import (
	"HealthCampaign/BackEnd/model"
	_ "fmt"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type MessageRepositoryMongo struct {
	db         *mgo.Database
	collection string
}

func NewMessageRepositoryMongo(db *mgo.Database, collection string) *MessageRepositoryMongo {
	return &MessageRepositoryMongo{
		db:         db,
		collection: collection,
	}
}

func (r *MessageRepositoryMongo) Save(Message *model.Message) error {
	err := r.db.C(r.collection).Insert(Message)
	//println(err)
	return err
}

func (r *MessageRepositoryMongo) Update(id string, Message *model.Message) error {

	err := r.db.C(r.collection).UpdateId(bson.ObjectIdHex(id), Message)
	return err

}


func (r *MessageRepositoryMongo) Delete(id string) error {
	err := r.db.C(r.collection).Remove(bson.M{"_id": id})
	return err
}


func (r *MessageRepositoryMongo) FindByID(id string) (*model.Message, error) {
	var Message model.Message
	err := r.db.C(r.collection).FindId(bson.ObjectIdHex(id)).One(&Message)

	if err != nil {
		return nil, err
	}

	return &Message, nil
}



func (r *MessageRepositoryMongo) FindAll() (model.Messages, error) {
	var Messages model.Messages

	err := r.db.C(r.collection).Find(bson.M{}).All(&Messages)

	if err != nil {
		return nil, err
	}

	return Messages, nil

}

