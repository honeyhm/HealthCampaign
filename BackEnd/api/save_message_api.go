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
	_,gtErr:=MessageUsecase.SaveMessage(&p)
	golog.Info(gtErr)
	hashtag , _ :=TagTokenize(p.MessageText)
	s := model.Search{ReferState :true , ReferId :p.Id.Hex() , HashTagId:hashtag.Id.Hex() , Id:bson.NewObjectId()}
	_ , gtSearchErr :=SearchUsecase.SaveSearch(&s)
	golog.Info(gtSearchErr)
	if gtErr==nil{
		Ctx.JSON("success")
		return
	}
	Ctx.JSON("fail")
}
