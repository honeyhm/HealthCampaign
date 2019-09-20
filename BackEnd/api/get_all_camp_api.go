package api

import (
	"HealthCampaign/BackEnd/model"
	_"HealthCampaign/BackEnd/usecase"
	"github.com/kataras/golog"
	"github.com/kataras/iris"
)

func AllCampaign(Ctx iris.Context){
	var p model.Campaigns
	p,gtAllErr:=CampaignUsecase.GetAll()
	golog.Info(gtAllErr)
	Ctx.JSON(&p)

}
