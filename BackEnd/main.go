

package main

import (
	"HealthCampaign/BackEnd/config"
	"HealthCampaign/BackEnd/api"
	"HealthCampaign/BackEnd/repository"
	"HealthCampaign/BackEnd/usecase"
	_ "github.com/kataras/golog"
	"github.com/kataras/iris"

	_ "github.com/kataras/iris/sessions"
	"gopkg.in/mgo.v2"
	"os"
	_ "time"
)


var (
	db  *mgo.Database
	err error
	patientUsecase usecase.PatientUsecase
	progressUsecase usecase.ProgressUsecase
	alterationUsecase usecase.AlterationUsecase
	campaignUsecase usecase.CampaignUsecase
	diseaseUsecase usecase.DiseaseUsecase
	groupUsecase usecase.GroupUsecase
	messageUsecase usecase.MessageUsecase
	medicalCenterUsecase usecase.MedicalCenterUsecase
	medicalStaffUsecase usecase.MedicalStaffUsecase
)


func main() {

	app := iris.New()

	app.Logger().SetLevel("debug")

	app.StaticWeb("/public", "./web/public")
	//app.HandleDir("/public", "./web/public")

	db, err := config.GetMongoDB()

	if err != nil {
		//golog.error("db connection error accured")
		os.Exit(2)
	}
	app.RegisterView(iris.HTML("web/view", ".html").Layout("layout.html"))
	//app.RegisterView(iris.HTML("web/view", ".html").Layout("layout.html"))//.Layout("layout.html"))//.Reload(true))
	//mvc.New(app.Party("authorize/web")).Handle(new(controller.ZoneController))
	//mvc.New(app.Party("/search")).Handle(new(controller.SearchController))
	//mvc.New(app.Party("/user")).Handle(new(controller.UserController))
	//mvc.New(app.Party("/log")).Handle(new(controller.TokenController))

	//	mvc.New(app.Party("/addzone")).Handle(new(controller.ZoneController))
	//mvc.New(app.Party(("/zone/editezone/{id}")).Handle(new(controller.AgreementController))
	InitApp(db)
	InitWeb()
	ApiControlller(app)
	////run application at port 8080 and set char to UTF-8
	app.Run(iris.Addr(":"+config.AppPort), iris.WithCharset("UTF-8"))

	/////////////////////////////////////////////////////

	//app.Get("/", func(ctx iris.Context) {
	//	ctx.ViewData("Message", "text for test")
	//	ctx.View("index.html") /////hamoon login ehtemalan
	//})
	//sessionManager := sessions.New(sessions.Config{
	//	Cookie:  "cookiename",
	//	Expires: time.Minute * 10,
	//})
	//////user route todo
	//user := mvc.New(app.Party("/web"))
	//user.Register(userUsecase, sessionManager.Start)
	//user.Handle(new(controller.UserController))

}


func InitApp(db *mgo.Database){

	alterationRepository := repository.NewAlterationRepositoryMongo(db, config.CollectionAlteration)
	alterationUsecase = usecase.NewAlterationUsecase(alterationRepository)

	campaignRepository := repository.NewCampaignRepositoryMongo(db, config.CollectionCampaign)
	campaignUsecase = usecase.NewCampaignUsecase(campaignRepository)

	diseaseRepository := repository.NewDiseaseRepositoryMongo(db,config.CollectionDisease)
	diseaseUsecase = usecase.NewDiseaseUsecase(diseaseRepository)

	groupRepository := repository.NewGroupRepositoryMongo(db,config.CollectionGroup)
	groupUsecase = usecase.NewGroupUsecase(groupRepository)

	medicalCenterRepository := repository.NewMedicalCenterRepositoryMongo(db,config.CollectionMedicalCenter)
	medicalCenterUsecase = usecase.NewMedicalCenterUsecase(medicalCenterRepository)

	medicalStaffRepository := repository.NewMedicalStaffRepositoryMongo(db,config.CollectionMedicalStaff)
	medicalStaffUsecase = usecase.NewMedicalStaffUsecase(medicalStaffRepository)

	messageRepository := repository.NewMessageRepositoryMongo(db,config.CollectionMessage)
	messageUsecase = usecase.NewMessageUsecase(messageRepository)

	patientRepository := repository.NewPatientRepositoryMongo(db,config.CollectionPatient)
	patientUsecase = usecase.NewPatientUsecase(patientRepository)

	progressRepository := repository.NewProgressRepositoryMongo(db,config.CollectionProgress)
	progressUsecase = usecase.NewProgressUsecase(progressRepository)

}


func InitWeb(){
	api.AlterationUsecase = alterationUsecase
	api.CampaignUsecase = campaignUsecase
	api.DiseaseUsecase = diseaseUsecase
	api.GroupUsecase = groupUsecase
	api.MedicalCenterUsecase = medicalCenterUsecase
	api.MedicalStaffUsecase = medicalStaffUsecase
	api.MessageUsecase = messageUsecase
	api.PatientUsecase = patientUsecase
	api.ProgressUsecase = progressUsecase
}

func ApiControlller(app *iris.Application){
	app.Get("/all-campaign",api.AllCampaign)
	app.Get("/all-message",api.AllMessage)
	app.Get("/all-group",api.AllGroup)
	app.Get("/campaign",api.GetByCampaignId)
	app.Get("/group",api.GetByGroupId)
	app.Get("/get-camp-patient",api.AllCampaignByPatientId)
	app.Post("/save-message",api.SaveMessage)
}
