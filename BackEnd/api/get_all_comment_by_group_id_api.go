package api

import (
	_"HealthCampaign/BackEnd/model"
	_"HealthCampaign/BackEnd/usecase"
	"github.com/kataras/golog"
	"github.com/kataras/iris"
)

func AllMessage(Ctx iris.Context){
	id:=Ctx.URLParam("id")
	p,gtAllErr:=MessageUsecase.GetAllByGroupId(id)
	golog.Info(gtAllErr)
	Ctx.JSON(&p)
}
