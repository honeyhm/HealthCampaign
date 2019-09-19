package api

import (
	"HealthCampaign/BackEnd/model"
	_"HealthCampaign/BackEnd/model"
	_"HealthCampaign/BackEnd/usecase"
	"github.com/kataras/golog"
	"github.com/kataras/iris"
	"gopkg.in/mgo.v2/bson"
)

func SaveMessage(Ctx iris.Context){
	var p model.Message
	Ctx.ReadJSON(&p)
	p.Id=bson.NewObjectId()
	p,gtErr:=MessageUsecase.Save(&p)
	golog.Info(gtErr)
	if gtErr==nil{
		Ctx.JSON("success")
		return
	}
	Ctx.JSON("fail")
}
