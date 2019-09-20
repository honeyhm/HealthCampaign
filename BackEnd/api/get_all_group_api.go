package api

import (
	_ "HealthCampaign/BackEnd/usecase"
	"github.com/kataras/golog"
	"github.com/kataras/iris"
)

func AllGroup(Ctx iris.Context){
	//var p model.Groups
	p,gtAllErr:=GroupUsecase.GetAll()
	golog.Info(gtAllErr)
	Ctx.JSON(&p)

}
