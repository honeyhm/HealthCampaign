package controller

import (
	"authorize/config"
	_"authorize/model"
	"authorize/usecase"

	_"strings"

	_"github.com/kataras/golog"
	_ "github.com/kataras/iris/mvc"
	"github.com/kataras/iris/sessions"
	_"gopkg.in/mgo.v2/bson"
)

var UserUsecase usecase.UserUsecase
var TokenUsecase usecase.TokenUsecase
var PersonUsecase usecase.PersonUsecase

//var LetterArchiveUsecase usecase.LetterArchiveUsecase

var Sess = sessions.New(sessions.Config{Cookie: config.CookieNameForSessionID, AllowReclaim: true})
