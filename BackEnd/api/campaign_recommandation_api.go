package api

import (
	"HealthCampaign/BackEnd/model"
	"github.com/kataras/golog"
	"github.com/kataras/iris"
)

func CampRec(Ctx iris.Context){
	userId := Ctx.URLParam("id")
	patient , gtPErr:= PatientUsecase.GetByID(userId)
	golog.Info(gtPErr)
	PCamps:=make([]model.Campaign,len(patient.CampaignId))
	for i:=0;i<len(patient.CampaignId)  ;i++  {
		PCamp , gtPCampErr := CampaignUsecase.GetByID(patient.CampaignId[i])
		golog.Info(gtPCampErr)
		PCamps[i]=*PCamp ////////patients camps
	}

	//var cams []model.Campaign
	GC:=make([]string,0)
	for i:=0;i<len(patient.GroupId) ;i++  {
		group , gtGErr := GroupUsecase.GetByID(patient.GroupId[i])
		golog.Info(gtGErr)
		for k:=0 ;k<len(PCamps) ;k++  {
			for z:=0;z<len(group.CampaignId) ;z++  {
				if group.CampaignId[z] != PCamps[k].Id.Hex(){
					GC = append(GC, group.CampaignId[z])
				}
			}
		}
	}


	///////////
	rcamps:=make([]model.Campaign,len(GC))
	for i:=0;i<len(GC) ;i++  {
		res,err:=CampaignUsecase.GetByID(GC[i])
		golog.Info(err)
		rcamps[i]=*res
	}

	Ctx.JSON(rcamps)

}
