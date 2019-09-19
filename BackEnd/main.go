

package main

import (
	"AbsharAutomation/config"
	"HealthCampaign/BackEnd/api"
	"HealthCampaign/BackEnd/repository"
	"HealthCampaign/BackEnd/usecase"
	"HealthCampaign/BackEnd/web/controller"
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
	//userUsecase usecase.UserUsecase
	tokenUsecase usecase.TokenUsecase
	personUsecase usecase.PersonUsecase
	
	//actionUsecase usecase.ActionUsecase
	//letterArchiveUsecase usecase.LetterArchiveUsecase

)


func main() {

	app := iris.New()

	app.Logger().SetLevel("debug")

	app.StaticWeb("/public", "./web/public")

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

	//actionRepository := repository.NewActionRepositoryMongo(db, config.CollectionAction)
	//actionUsecase = usecase.NewActionUsecase(actionRepository)


	//userRepository := repository.NewUserRepositoryMongo(db, config.CollectionUser) // returns *UserRepositoryMongo
	//userUsecase = usecase.NewUserUsecase(userRepository)

	//
	tokenRepository := repository.NewTokenRepositoryMongo(db,config.CollectionToken)
	tokenUsecase = usecase.NewTokenUsecase(tokenRepository)


	
	//letterArchiveRepository := repository.NewFlowRepositoryMongo(db,config.CollectionLetterArchive)
	//letterArchiveUsecase = usecase.NewLetterArchiveUsecase(letterArchiveRepository)
	//

}


func InitWeb(){
	//controller.ActionUsecase = actionUsecase
	//controller.UserUsecase = userUsecase
	controller.TokenUsecase = tokenUsecase

	controller.PersonUsecase = personUsecase
	api.TokenUsecase = tokenUsecase
	api.PersonUsecase = personUsecase
}

func ApiControlller(app *iris.Application){
	app.Get("/userget",api.GetStudent)
	app.Post("/t",api.PostStudent)
}
