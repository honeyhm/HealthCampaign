package api

import (
	"HealthCampaign/BackEnd/model"
	"github.com/kataras/golog"
	"github.com/kataras/iris"
	"gopkg.in/mgo.v2/bson"
)

func MedCenterSignin(Ctx iris.Context){
	var p model.MedicalCenter
	Ctx.ReadJSON(&p)//todo false
	p.Id=bson.NewObjectId()
	_,svErr:=MedicalCenterUsecase.Save(&p)
	golog.Info(svErr)
	Ctx.JSON("success")
}
