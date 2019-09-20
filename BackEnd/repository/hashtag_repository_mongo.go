
package repository

import (
	"HealthCampaign/BackEnd/model"
	_ "fmt"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type HashtagRepositoryMongo struct {
	db         *mgo.Database
	collection string
}

func NewHashtagRepositoryMongo(db *mgo.Database, collection string) *HashtagRepositoryMongo {
	return &HashtagRepositoryMongo{
		db:         db,
		collection: collection,
	}
}

func (r *HashtagRepositoryMongo) Save(Hashtag *model.Hashtag) error {
	err := r.db.C(r.collection).Insert(Hashtag)
	//println(err)
	return err
}

func (r *HashtagRepositoryMongo) Update(id string, Hashtag *model.Hashtag) error {

	err := r.db.C(r.collection).UpdateId(bson.ObjectIdHex(id), Hashtag)
	return err

}


func (r *HashtagRepositoryMongo) Delete(id string) error {
	err := r.db.C(r.collection).Remove(bson.M{"_id": id})
	return err
}


func (r *HashtagRepositoryMongo) FindByID(id string) (*model.Hashtag, error) {
	var Hashtag model.Hashtag
	err := r.db.C(r.collection).FindId(bson.ObjectIdHex(id)).One(&Hashtag)

	if err != nil {
		return nil, err
	}

	return &Hashtag, nil
}



func (r *HashtagRepositoryMongo) FindAll() (model.Hashtags, error) {
	var Hashtags model.Hashtags

	err := r.db.C(r.collection).Find(bson.M{}).All(&Hashtags)

	if err != nil {
		return nil, err
	}

	return Hashtags, nil

}

