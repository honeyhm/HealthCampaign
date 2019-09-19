
package repository

import (
	"HealthCampaign/BackEnd/model"
	_ "fmt"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type AlterationRepositoryMongo struct {
	db         *mgo.Database
	collection string
}

func NewAlterationRepositoryMongo(db *mgo.Database, collection string) *AlterationRepositoryMongo {
	return &AlterationRepositoryMongo{
		db:         db,
		collection: collection,
	}
}

func (r *AlterationRepositoryMongo) Save(Alteration *model.Alteration) error {
	err := r.db.C(r.collection).Insert(Alteration)
	//println(err)
	return err
}

func (r *AlterationRepositoryMongo) Update(id string, Alteration *model.Alteration) error {

	err := r.db.C(r.collection).UpdateId(bson.ObjectIdHex(id), Alteration)
	return err

}


func (r *AlterationRepositoryMongo) Delete(id string) error {
	err := r.db.C(r.collection).Remove(bson.M{"_id": id})
	return err
}


func (r *AlterationRepositoryMongo) FindByID(id string) (*model.Alteration, error) {
	var Alteration model.Alteration
	err := r.db.C(r.collection).FindId(bson.ObjectIdHex(id)).One(&Alteration)

	if err != nil {
		return nil, err
	}

	return &Alteration, nil
}



func (r *AlterationRepositoryMongo) FindAll() (model.Alterations, error) {
	var Alterations model.Alterations

	err := r.db.C(r.collection).Find(bson.M{}).All(&Alterations)

	if err != nil {
		return nil, err
	}

	return Alterations, nil

}
