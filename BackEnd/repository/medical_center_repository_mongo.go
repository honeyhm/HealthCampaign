
package repository

import (
	"HealthCampaign/BackEnd/model"
	_ "fmt"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type MedicalCenterRepositoryMongo struct {
	db         *mgo.Database
	collection string
}

func NewMedicalCenterRepositoryMongo(db *mgo.Database, collection string) *MedicalCenterRepositoryMongo {
	return &MedicalCenterRepositoryMongo{
		db:         db,
		collection: collection,
	}
}

func (r *MedicalCenterRepositoryMongo) Save(MedicalCenter *model.MedicalCenter) error {
	err := r.db.C(r.collection).Insert(MedicalCenter)
	//println(err)
	return err
}

func (r *MedicalCenterRepositoryMongo) Update(id string, MedicalCenter *model.MedicalCenter) error {

	err := r.db.C(r.collection).UpdateId(bson.ObjectIdHex(id), MedicalCenter)
	return err

}


func (r *MedicalCenterRepositoryMongo) Delete(id string) error {
	err := r.db.C(r.collection).Remove(bson.M{"_id": id})
	return err
}


func (r *MedicalCenterRepositoryMongo) FindByID(id string) (*model.MedicalCenter, error) {
	var MedicalCenter model.MedicalCenter
	err := r.db.C(r.collection).FindId(bson.ObjectIdHex(id)).One(&MedicalCenter)

	if err != nil {
		return nil, err
	}

	return &MedicalCenter, nil
}



func (r *MedicalCenterRepositoryMongo) FindAll() (model.MedicalCenters, error) {
	var MedicalCenters model.MedicalCenters

	err := r.db.C(r.collection).Find(bson.M{}).All(&MedicalCenters)

	if err != nil {
		return nil, err
	}

	return MedicalCenters, nil

}
