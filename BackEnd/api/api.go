package api

import (
	"HealthCampaign/BackEnd/model"
	"HealthCampaign/BackEnd/config"
	"HealthCampaign/BackEnd/usecase"
	"github.com/kataras/golog"
	"github.com/kataras/iris/mvc"
	"gopkg.in/mgo.v2/bson"
	"strings"
	_"strings"
	"time"
	_"unicode"

	_"github.com/kataras/golog"
	_ "github.com/kataras/iris/mvc"
	"github.com/kataras/iris/sessions"
	_"gopkg.in/mgo.v2/bson"
)

var AlterationUsecase usecase.AlterationUsecase
var CampaignUsecase usecase.CampaignUsecase
var DiseaseUsecase usecase.DiseaseUsecase
var GroupUsecase usecase.GroupUsecase
var MedicalCenterUsecase usecase.MedicalCenterUsecase
var MedicalStaffUsecase usecase.MedicalStaffUsecase
var MessageUsecase usecase.MessageUsecase
var PatientUsecase usecase.PatientUsecase
var ProgressUsecase usecase.ProgressUsecase
var HashtagUsecase usecase.HashtagUsecase
var SearchUsecase usecase.SearchUsecase


var Sess = sessions.New(sessions.Config{Cookie: config.CookieNameForSessionID, AllowReclaim: true})

func TagTokenize( messageText string) (*model.Hashtag , error) {

	Num := strings.Count(messageText, "#")
	var err error
	var hashtag model.Hashtag
	for i := 1; i < Num+1; i++ {
		newObjID := bson.NewObjectId()
		i1 := strings.IndexAny(messageText, "#")
		i2 := strings.IndexAny(messageText, " ")
		messageText =messageText[:i2]+string("*")+messageText[i2+1:]
		if i2 > i1{
			tag := messageText[i1+1:i2]
			hashtag := model.Hashtag{Id:newObjID , TagName:tag}
			_ , err := HashtagUsecase.SaveHashtag(&hashtag)
			golog.Info(err)
		}else{
				i--
		}
	}
	if err != nil {
		return nil ,err
	}
	return &hashtag , nil
}





//func SaveSearch(ReferId string ,ReferState bool) error {
//	newObjID := bson.NewObjectId()
//	newObjID2 := bson.NewObjectId()
//	search := model.Search{Id:newObjID , HashTagId:newObjID2.Hex() ,ReferId:ReferId ,ReferState:ReferState }
//	/////////////////tokeniz text to find field\\\\\\\\\\\
//	tokenErr:= LetterTokenize(content , LetterId)
//