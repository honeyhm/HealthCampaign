package api

import (
	_"HealthCampaign/BackEnd/model"
	_"HealthCampaign/BackEnd/usecase"
	"github.com/kataras/golog"
	"github.com/kataras/iris"
)

func GetByCampaignId(Ctx iris.Context){
	id:=Ctx.URLParam("id")
	p,gtAllErr:=CampaignUsecase.GetByID(id)
	golog.Info(gtAllErr)
	Ctx.JSON(&p)

}
