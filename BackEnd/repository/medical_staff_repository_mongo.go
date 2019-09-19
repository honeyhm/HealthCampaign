
package repository

import (
	"HealthCampaign/BackEnd/model"
	_ "fmt"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type MedicalStaffRepositoryMongo struct {
	db         *mgo.Database
	collection string
}

func NewMedicalStaffRepositoryMongo(db *mgo.Database, collection string) *MedicalStaffRepositoryMongo {
	return &MedicalStaffRepositoryMongo{
		db:         db,
		collection: collection,
	}
}

func (r *MedicalStaffRepositoryMongo) Save(MedicalStaff *model.MedicalStaff) error {
	err := r.db.C(r.collection).Insert(MedicalStaff)
	//println(err)
	return err
}

func (r *MedicalStaffRepositoryMongo) Update(id string, MedicalStaff *model.MedicalStaff) error {

	err := r.db.C(r.collection).UpdateId(bson.ObjectIdHex(id), MedicalStaff)
	return err

}


func (r *MedicalStaffRepositoryMongo) Delete(id string) error {
	err := r.db.C(r.collection).Remove(bson.M{"_id": id})
	return err
}


func (r *MedicalStaffRepositoryMongo) FindByID(id string) (*model.MedicalStaff, error) {
	var MedicalStaff model.MedicalStaff
	err := r.db.C(r.collection).FindId(bson.ObjectIdHex(id)).One(&MedicalStaff)

	if err != nil {
		return nil, err
	}

	return &MedicalStaff, nil
}


func(r *MedicalStaffRepositoryMongo) FindByUserPass (UserName string ,Password string) (bool){

	err := r.db.C(r.collection).Find(bson.M{"username":UserName , "password":Password})

	if err != nil {
		return false
	}

	return true

}



func (r *MedicalStaffRepositoryMongo) FindAll() (model.MedicalStaffs, error) {
	var MedicalStaffs model.MedicalStaffs

	err := r.db.C(r.collection).Find(bson.M{}).All(&MedicalStaffs)

	if err != nil {
		return nil, err
	}

	return MedicalStaffs, nil

}
