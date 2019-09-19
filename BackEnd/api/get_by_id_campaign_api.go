package api

import (
	"HealthCampaign/BackEnd/model"
	_"HealthCampaign/BackEnd/usecase"
	"github.com/kataras/golog"
	"github.com/kataras/iris"
)

func GetByCampaignId(Ctx iris.Context){
	id:=Ctx.URLParam("id")
	p,gtAllErr:=CampaignUsecase.GetById(id)
	golog.Info(gtAllErr)
	Ctx.JSON(&p)

}
