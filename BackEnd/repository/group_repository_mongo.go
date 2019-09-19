
package repository

import (
	"HealthCampaign/BackEnd/model"
	_ "fmt"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type GroupRepositoryMongo struct {
	db         *mgo.Database
	collection string
}

func NewGroupRepositoryMongo(db *mgo.Database, collection string) *GroupRepositoryMongo {
	return &GroupRepositoryMongo{
		db:         db,
		collection: collection,
	}
}

func (r *GroupRepositoryMongo) Save(Group *model.Group) error {
	err := r.db.C(r.collection).Insert(Group)
	//println(err)
	return err
}

func (r *GroupRepositoryMongo) Update(id string, Group *model.Group) error {

	err := r.db.C(r.collection).UpdateId(bson.ObjectIdHex(id), Group)
	return err

}


func (r *GroupRepositoryMongo) Delete(id string) error {
	err := r.db.C(r.collection).Remove(bson.M{"_id": id})
	return err
}


func (r *GroupRepositoryMongo) FindByID(id string) (*model.Group, error) {
	var Group model.Group
	err := r.db.C(r.collection).FindId(bson.ObjectIdHex(id)).One(&Group)

	if err != nil {
		return nil, err
	}

	return &Group, nil
}



func (r *GroupRepositoryMongo) FindAll() (model.Groups, error) {
	var Groups model.Groups

	err := r.db.C(r.collection).Find(bson.M{}).All(&Groups)

	if err != nil {
		return nil, err
	}

	return Groups, nil

}

