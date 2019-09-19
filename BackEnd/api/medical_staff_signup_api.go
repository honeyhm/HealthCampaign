package api

import (
	"HealthCampaign/BackEnd/model"
	"github.com/kataras/golog"
	"github.com/kataras/iris"
	"gopkg.in/mgo.v2/bson"
)

func MedStaffSignup(Ctx iris.Context){
	var p model.MedicalStaff
	Ctx.ReadJSON(&p)
	p.Id=bson.NewObjectId()
	_,svErr:=MedicalStaffUsecase.Save(&p)
	golog.Info(svErr)
	Ctx.JSON("success")
}
