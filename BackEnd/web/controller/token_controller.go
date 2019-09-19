package controller

import (
	"github.com/kataras/iris"
	"github.com/kataras/iris/mvc"
)

type TokenController struct {
	Ctx iris.Context
}

func (c *TokenController) Get() mvc.Result {
	t,_:=TokenUsecase.GetAll()
	return mvc.View{
		Name:"log.html",
		Data:iris.Map{"rng":t} ,
	}
}