package model

import "gopkg.in/mgo.v2/bson"

type Progress struct {
	Id         		  bson.ObjectId `bson:"_id"`
	CampaignId        string        `bson:"campaign_id"`
	PatientName		  string        `bson:"patient_name"`
	Progress		  int           `bson:"organ_name"`

}
type Progresses []Progress
