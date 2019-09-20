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

func TagTokenize( messageText string, referId string ) error {

	Num := strings.Count(messageText, "#")
	var err error
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
		return err
	}
	return nil
}





func SaveSearch() error {
	subject := c.Ctx.FormValue("subject")
	content := c.Ctx.FormValue("content-editor")
	LetterId := c.Ctx.FormValue("AI")
	attachment:=c.Ctx.FormValue("attachment")
	golog.Info("attach")
	golog.Info(attachment)
	dt := time.Now() //.Format("01-02-2006")
	lttrObjID := bson.ObjectIdHex(LetterId)
	var lttr model.Letter
	lttr.Id = lttrObjID
	lttr.LetterSubject = subject
	lttr.LetterText = content
	lttr.CreatedAt = dt
	/////////////////tokeniz text to find field\\\\\\\\\\\
	tokenErr:= LetterTokenize(content , LetterId)
	if tokenErr!= nil{
		return mvc.Response{
			Path: "addletter",//ارور خوردیم حالا باید بریم یه جایی که نمیدونم فعلا
		}
	}
	////////////////////////////////end tokenize\\\\\\\\\\\\\\\\\
	_ , err := LetterUsecase.SaveLetter(&lttr)
	golog.Info(err)
	if err == nil {
		//tokenize
		golog.Info(attachment);
		if attachment=="1" {
			return mvc.Response{
				Path: "/letter/addattachment?id=" + LetterId,
			}
		} else {
			return mvc.Response{
				Path: "/letter/showletter?id=" + LetterId,
			}
		}
	}
	return mvc.Response{
		Path: "/letter/addletter",
	}
}