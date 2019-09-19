package controller

import (
	"authorize/model"
	"encoding/json"
	"github.com/kataras/golog"
	"github.com/kataras/iris"
	"gopkg.in/mgo.v2/bson"
	"time"
)

type SearchController struct {
	Ctx iris.Context
}

func (c *SearchController) Get()  {
	// golog.Info("Enter get OrganZoneid")
	// idAction := "10"
	// token := Sess.Start(c.Ctx).GetString(config.SessionToken)
	// access := Authorize(token ,idAction)
	// //access=true
	// if access ==true{
		oid:=c.Ctx.URLParam("q")
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
			res,_:=json.Marshal("")
			golog.Info(res)
			c.Ctx.ResponseWriter().Header().Set("Content-Type","application/json")
			c.Ctx.ResponseWriter().Write(res)
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
			res, _ := json.Marshal(po)
			golog.Info(res)
			c.Ctx.ResponseWriter().Header().Set("Content-Type", "application/json")
			c.Ctx.ResponseWriter().Write(res)
		}

	// }

}