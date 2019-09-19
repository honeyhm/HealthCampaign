
package repository

import (
	"HealthCampaign/BackEnd/model"
	_ "fmt"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type PatientRepositoryMongo struct {
	db         *mgo.Database
	collection string
}

func NewPatientRepositoryMongo(db *mgo.Database, collection string) *PatientRepositoryMongo {
	return &PatientRepositoryMongo{
		db:         db,
		collection: collection,
	}
}

func (r *PatientRepositoryMongo) Save(Patient *model.Patient) error {
	err := r.db.C(r.collection).Insert(Patient)
	//println(err)
	return err
}

func (r *PatientRepositoryMongo) Update(id string, Patient *model.Patient) error {

	err := r.db.C(r.collection).UpdateId(bson.ObjectIdHex(id), Patient)
	return err

}


func (r *PatientRepositoryMongo) Delete(id string) error {
	err := r.db.C(r.collection).Remove(bson.M{"_id": id})
	return err
}


func (r *PatientRepositoryMongo) FindByID(id string) (*model.Patient, error) {
	var Patient model.Patient
	err := r.db.C(r.collection).FindId(bson.ObjectIdHex(id)).One(&Patient)

	if err != nil {
		return nil, err
	}

	return &Patient, nil
}


func(r *PatientRepositoryMongo) FindByUserPass (UserName string ,Password string) (bool){

	err := r.db.C(r.collection).Find(bson.M{"username":UserName , "password":Password})

	if err != nil {
		return false
	}

	return true

}



func (r *PatientRepositoryMongo) FindAll() (model.Patients, error) {
	var Patients model.Patients

	err := r.db.C(r.collection).Find(bson.M{}).All(&Patients)

	if err != nil {
		return nil, err
	}

	return Patients, nil

}
