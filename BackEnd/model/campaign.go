
package model

import "gopkg.in/mgo.v2/bson"
import "time"

type Campaign struct {
	Id         			bson.ObjectId `bson:"_id"`
	CampaignName       	string        `bson:"campaign_name"`
	Description 		string        `bson:"description"`
	CreationDate 		time.Time     `bson:"creation_date"`
	MedicalStaffId 		[]string      `bson:"medical_staff_id"`
	MemberNumber 		time.Time     `bson:"member_number"`
	DailyThreshold 		int     	  `bson:"daily_threshold"`
	FullThreshold 		int    		  `bson:"full_threshold"`
	Privilege 		    int    		  `bson:"privilege"`   ////Emtiaz


}

type Campaigns []Campaign