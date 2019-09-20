package api

import (
	"HealthCampaign/BackEnd/model"
	_ "HealthCampaign/BackEnd/model"
	_ "HealthCampaign/BackEnd/usecase"
	"github.com/kataras/golog"
	"github.com/kataras/iris"
)

func SearchMessageByHashtag(Ctx iris.Context){
	tagName:= Ctx.URLParam("id")
	hashtag , gtHashtagErr := HashtagUsecase.GetByName(tagName)
	golog.Info("gtHashtagErr")
	golog.Info(gtHashtagErr)
	hashtagId := hashtag.Id.Hex()
	searches , gtSearchErr := SearchUsecase.GetByTagId(hashtagId)
	golog.Info("gtSearchErr")
	golog.Info(gtSearchErr)
	sSize := len(searches)
	msgs:=make([]model.Message,sSize)
	for i := 0; i<sSize ;i++ {
		if searches[i].ReferState == true{
			message , gtMessageErr := MessageUsecase.GetByID(searches[i].ReferId)
			golog.Info("gtMessageErr")
			golog.Info(gtMessageErr)
			msgs[i]=*message


		}
	}

	Ctx.JSON(msgs)
}
