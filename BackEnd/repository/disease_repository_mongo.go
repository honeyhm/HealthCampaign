
package repository

import (
	"HealthCampaign/BackEnd/model"
	_ "fmt"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type DiseaseRepositoryMongo struct {
	db         *mgo.Database
	collection string
}

func NewDiseaseRepositoryMongo(db *mgo.Database, collection string) *DiseaseRepositoryMongo {
	return &DiseaseRepositoryMongo{
		db:         db,
		collection: collection,
	}
}

func (r *DiseaseRepositoryMongo) Save(Disease *model.Disease) error {
	err := r.db.C(r.collection).Insert(Disease)
	//println(err)
	return err
}

func (r *DiseaseRepositoryMongo) Update(id string, Disease *model.Disease) error {

	err := r.db.C(r.collection).UpdateId(bson.ObjectIdHex(id), Disease)
	return err

}


func (r *DiseaseRepositoryMongo) Delete(id string) error {
	err := r.db.C(r.collection).Remove(bson.M{"_id": id})
	return err
}


func (r *DiseaseRepositoryMongo) FindByID(id string) (*model.Disease, error) {
	var Disease model.Disease
	err := r.db.C(r.collection).FindId(bson.ObjectIdHex(id)).One(&Disease)

	if err != nil {
		return nil, err
	}

	return &Disease, nil
}



func (r *DiseaseRepositoryMongo) FindAll() (model.Diseases, error) {
	var Diseases model.Diseases

	err := r.db.C(r.collection).Find(bson.M{}).All(&Diseases)

	if err != nil {
		return nil, err
	}

	return Diseases, nil

}
