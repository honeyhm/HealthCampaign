
package repository

import (
	"HealthCampaign/BackEnd/model"
	_ "fmt"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type ProgressRepositoryMongo struct {
	db         *mgo.Database
	collection string
}

func NewProgressRepositoryMongo(db *mgo.Database, collection string) *ProgressRepositoryMongo {
	return &ProgressRepositoryMongo{
		db:         db,
		collection: collection,
	}
}

func (r *ProgressRepositoryMongo) Save(Progress *model.Progress) error {
	err := r.db.C(r.collection).Insert(Progress)
	//println(err)
	return err
}

func (r *ProgressRepositoryMongo) Update(id string, Progress *model.Progress) error {

	err := r.db.C(r.collection).UpdateId(bson.ObjectIdHex(id), Progress)
	return err

}


func (r *ProgressRepositoryMongo) Delete(id string) error {
	err := r.db.C(r.collection).Remove(bson.M{"_id": id})
	return err
}


func (r *ProgressRepositoryMongo) FindByID(id string) (*model.Progress, error) {
	var Progress model.Progress
	err := r.db.C(r.collection).FindId(bson.ObjectIdHex(id)).One(&Progress)

	if err != nil {
		return nil, err
	}

	return &Progress, nil
}



func (r *ProgressRepositoryMongo) FindAll() (model.Progresses, error) {
	var Progresss model.Progresses

	err := r.db.C(r.collection).Find(bson.M{}).All(&Progresss)

	if err != nil {
		return nil, err
	}

	return Progresss, nil

}
