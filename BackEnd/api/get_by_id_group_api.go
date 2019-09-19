package api

import (
	_"HealthCampaign/BackEnd/model"
	_"HealthCampaign/BackEnd/usecase"
	"github.com/kataras/golog"
	"github.com/kataras/iris"
)

func GetByGroupId(Ctx iris.Context){
	id:=Ctx.URLParam("id")
	p,gtAllErr:=GroupUsecase.GetById(id)
	golog.Info(gtAllErr)
	Ctx.JSON(&p)

}
