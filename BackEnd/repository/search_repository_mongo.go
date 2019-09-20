
package repository

import (
	"HealthCampaign/BackEnd/model"
	_ "fmt"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type SearchRepositoryMongo struct {
	db         *mgo.Database
	collection string
}

func NewSearchRepositoryMongo(db *mgo.Database, collection string) *SearchRepositoryMongo {
	return &SearchRepositoryMongo{
		db:         db,
		collection: collection,
	}
}

func (r *SearchRepositoryMongo) Save(Search *model.Search) error {
	err := r.db.C(r.collection).Insert(Search)
	//println(err)
	return err
}

func (r *SearchRepositoryMongo) Update(id string, Search *model.Search) error {

	err := r.db.C(r.collection).UpdateId(bson.ObjectIdHex(id), Search)
	return err

}


func (r *SearchRepositoryMongo) Delete(id string) error {
	err := r.db.C(r.collection).Remove(bson.M{"_id": id})
	return err
}


func (r *SearchRepositoryMongo) FindByID(id string) (*model.Search, error) {
	var Search model.Search
	err := r.db.C(r.collection).FindId(bson.ObjectIdHex(id)).One(&Search)

	if err != nil {
		return nil, err
	}

	return &Search, nil
}



func (r *SearchRepositoryMongo) FindAll() (model.Searches, error) {
	var Searches model.Searches

	err := r.db.C(r.collection).Find(bson.M{}).All(&Searches)

	if err != nil {
		return nil, err
	}

	return Searches, nil

}

