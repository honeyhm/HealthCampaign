package api

import (
	"HealthCampaign/BackEnd/model"
	_"HealthCampaign/BackEnd/model"
	_"HealthCampaign/BackEnd/usecase"
	"github.com/kataras/golog"
	"github.com/kataras/iris"
)

func AllMessage(Ctx iris.Context){
	id:=Ctx.URLParam("id")
	p,gtAllErr:=MessageUsecase.GetAllByGroupId(id)
	t:=make([]model.Message,0)
	for i:=0;i<len(p) ;i++  {
		if p[i].ParentId == ""{
			t= append(t, p[i])
		}
	}
	golog.Info(gtAllErr)
	Ctx.JSON(&t)
}
