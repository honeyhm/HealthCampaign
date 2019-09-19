package api

import (
	"HealthCampaign/BackEnd/config"
	"HealthCampaign/BackEnd/usecase"
	_"strings"

	_"github.com/kataras/golog"
	_ "github.com/kataras/iris/mvc"
	"github.com/kataras/iris/sessions"
	_"gopkg.in/mgo.v2/bson"
)

var AlterationUsecase usecase.AlterationUsecase
var CampaignUsecase usecase.CampaignUsecase
var DiseaseUsecase usecase.DiseaseUsecase
var GroupUsecase usecase.GroupUsecase
var MedicalCenterUsecase usecase.MedicalCenterUsecase
var MedicalStaffUsecase usecase.MedicalStaffUsecase
var MessageUsecase usecase.MessageUsecase
var PatientUsecase usecase.PatientUsecase
var ProgressUsecase usecase.ProgressUsecase


var Sess = sessions.New(sessions.Config{Cookie: config.CookieNameForSessionID, AllowReclaim: true})

