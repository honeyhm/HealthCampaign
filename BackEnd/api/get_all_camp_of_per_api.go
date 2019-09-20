package api

import (
	_ "HealthCampaign/BackEnd/modelview"
	_ "HealthCampaign/BackEnd/usecase"
	"github.com/kataras/golog"
	"github.com/kataras/iris"
)

func AllCampaignByPatientId(Ctx iris.Context){
	id:=Ctx.URLParam("id")
	golog.Info(id)
	//var p model.Campaigns
	usr,gtpaErr:=PatientUsecase.GetByID(id)
	golog.Info(gtpaErr)
	golog.Info(usr)
	Ctx.JSON(usr)

}
