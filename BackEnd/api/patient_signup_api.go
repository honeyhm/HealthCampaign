package api

import (
	"HealthCampaign/BackEnd/model"
	"github.com/kataras/golog"
	"github.com/kataras/iris"
	"gopkg.in/mgo.v2/bson"
)

func PatientSignup(Ctx iris.Context){
	var p model.Patient
	Ctx.ReadJSON(&p)
	p.Id=bson.NewObjectId()
	_,svErr:=PatientUsecase.Save(&p)
	golog.Info(svErr)
	Ctx.JSON("success")
}
