
package repository

import (
	"HealthCampaign/BackEnd/model"
	_ "fmt"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type CampaignRepositoryMongo struct {
	db         *mgo.Database
	collection string
}

func NewCampaignRepositoryMongo(db *mgo.Database, collection string) *CampaignRepositoryMongo {
	return &CampaignRepositoryMongo{
		db:         db,
		collection: collection,
	}
}

func (r *CampaignRepositoryMongo) Save(Campaign *model.Campaign) error {
	err := r.db.C(r.collection).Insert(Campaign)
	//println(err)
	return err
}

func (r *CampaignRepositoryMongo) Update(id string, Campaign *model.Campaign) error {

	err := r.db.C(r.collection).UpdateId(bson.ObjectIdHex(id), Campaign)
	return err

}


func (r *CampaignRepositoryMongo) Delete(id string) error {
	err := r.db.C(r.collection).Remove(bson.M{"_id": id})
	return err
}


func (r *CampaignRepositoryMongo) FindByID(id string) (*model.Campaign, error) {
	var Campaign model.Campaign
	err := r.db.C(r.collection).FindId(bson.ObjectIdHex(id)).One(&Campaign)

	if err != nil {
		return nil, err
	}

	return &Campaign, nil
}



func (r *CampaignRepositoryMongo) FindAll() (model.Campaigns, error) {
	var Campaigns model.Campaigns

	err := r.db.C(r.collection).Find(bson.M{}).All(&Campaigns)

	if err != nil {
		return nil, err
	}

	return Campaigns, nil

}
