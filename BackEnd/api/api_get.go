package api

import (
	_"encoding/json"
	"github.com/kataras/golog"
	"github.com/kataras/iris"
	"gopkg.in/mgo.v2/bson"
	"time"
)

func GetStudent(Ctx iris.Context) {
	oid:=Ctx.URLParam("q")
	// zid:=c.Ctx.URLParam("zid")
	golog.Info("q :")
	golog.Info(oid)
	// golog.Info("zid :")
	// golog.Info(zid)
	po , gtPoErr := PersonUsecase.GetByNum(oid)
	golog.Info(po)
	golog.Info("gtPoErr :")
	golog.Info(gtPoErr)

	if gtPoErr != nil {
		golog.Info("11111")
		Ctx.JSON("")
		//res,_:=json.Marshal("")
		//golog.Info(res)
		//Ctx.ResponseWriter().Header().Set("Content-Type","application/json")
		//Ctx.ResponseWriter().Write(res)
		/*} else if time.Now().After(po.FinishDate){
		golog.Info("2222222")
		golog.Info(time.Now())
		golog.Info(po.FinishDate)
		golog.Info(time.Now().After(po.FinishDate))
		res,_:=json.Marshal('0')
		golog.Info(res)
		c.Ctx.ResponseWriter().Header().Set("Content-Type","application/json")
		c.Ctx.ResponseWriter().Write(res)*/
	}else {
		golog.Info("3333333333")

		token := model.Token{Id: bson.NewObjectId(), PersonId: po.Id.Hex(), StudentNum: po.StudentNum, EnterTime: time.Now()}
		_, err := TokenUsecase.SaveToken(&token)
		golog.Info(err)
		Ctx.JSON(po)
		//res, _ := json.Marshal(po)
		//golog.Info(res)
		//Ctx.ResponseWriter().Header().Set("Content-Type", "application/json")
		//Ctx.ResponseWriter().Write(res)
	}
}


//
//func PostStudent(Ctx iris.Context) {
//	var s model.Person
//	Ctx.ReadJSON(&s)
//	_ , gtPoErr := PersonUsecase.SavePerson(&s)
//	golog.Info(gtPoErr)
//	Ctx.JSON("succsess")
//
//}
