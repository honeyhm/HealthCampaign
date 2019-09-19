/*login and profile created by H.MLK 12/5/1398*/
/*edit profile & register by H.MLK 22/5/1398*/

package controller

import (
	_ "AbsharAutomation/usecase"
	"authorize/model"
	"github.com/kataras/golog"
	"github.com/kataras/iris"
	"github.com/kataras/iris/mvc"
	"gopkg.in/mgo.v2/bson"
	"time"
)

//userIDkey
const UserIDKey = "UserID"

//User Controller
type UserController struct {
	Ctx iris.Context
}

func (c *UserController) GetAdduser() mvc.Result {
	return mvc.View{
		Name:"/adduser.html",
	}
}


func (c *UserController) PostAdduser() mvc.Result {
	t1,_:=time.Parse("01-01-2001",c.Ctx.FormValue("fdate"))
	u:=model.Person{Id:bson.NewObjectId(),Name:c.Ctx.FormValue("name"),
		StudentNum:c.Ctx.FormValue("snum"),
		Ssn:c.Ctx.FormValue("ssn"),
		Gender:c.Ctx.FormValue("gender"),
		FinishDate:t1}
	_,err:=PersonUsecase.SavePerson(&u)
	golog.Info(err)
	return mvc.Response{
		Path:"adduser",
	}
}

//// //get current UserID
//// func (c *UserController) getCurrentUserID() string {
//// 	return c.Session.GetString(UserIDKey)
//// }
//// //is user logged in
//// func (c *UserController) isUserLoggedIn() bool {
//// 	return c.getCurrentUserID() != ""
//// }
//// //logout
//// func (c *UserController) logout() {
//// 	c.Session.Destroy()
//// }
//
////localhost:8080/user/login GET as index
//func (c *UserController) Get() mvc.Result {
//	//if c.isUserLoggedIn(){
//	//	c.logout()
//	//}
//	t := Sess.Start(c.Ctx).GetString(config.SessionToken)
//	if t != "" {
//		u, err := UserUsecase.GetByToken(t)
//		if err != nil {
//			return mvc.View{
//				Name: "/login/index.html",
//				Data: iris.Map{"Title": "Login", "flaghshow": "0"},
//			}
//		}
//		if u == nil {
//			return mvc.View{
//				Name: "/login/index.html",
//				Data: iris.Map{"Title": "Login", "flaghshow": "0"},
//			}
//		} else {
//			return mvc.Response{
//				//Path: "/login/profile?id=" + u.Id.Hex(), //in baiad zadeh she nadarimesh
//				Path: "/login/profile", //in baiad zadeh she nadarimesh
//			}
//		}
//	} else {
//		return mvc.View{
//			Name: "/login/index.html",
//			Data: iris.Map{"Title": "Login", "flaghshow": "0"},
//		}
//	}
//
//}
//
////localhost:8080/user/login POST
//func (c *UserController) Post() mvc.Result {
//
//	userName := c.Ctx.FormValue("username")
//	password := c.Ctx.FormValue("password")
//
//	golog.Info(userName)
//	golog.Info(password)
//
//	gtUser, gtUserErr := UserUsecase.GetByUserName(userName)
//	golog.Info(gtUserErr)
//
//	if gtUserErr == nil {
//		if gtUser.Password == password {
//			session := Sess.Start(c.Ctx)
//			tk := bson.NewObjectId()
//			session.Set(config.SessionToken, tk.Hex())
//			gtUser.UserToken = tk.Hex()
//			golog.Info("tk.Hex() :")
//			golog.Info(tk.Hex())
//			UserUsecase.UpdateUser(gtUser.Id.Hex(), gtUser)
//
//			return mvc.Response{
//				//Path: "/login/profile?id=" + gtUser.Id.Hex(), //in baiad zadeh she nadarimesh
//				Path: "/login/profile", //in baiad zadeh she nadarimesh
//			}
//		}
//	}
//	return mvc.Response{
//		Path: "login",
//	}
//}
//
//
//func (c *UserController) GetProfile() mvc.Result {
//	golog.Info("Enter get profile ")
//	//uIdtmp := c.Ctx.URLParam("id")
//	t := Sess.Start(c.Ctx).GetString(config.SessionToken)
//	golog.Error(t)
//
//
//	//uId := bson.ObjectIdHex(uIdtmp)
//	//gtUser, gtUserErr := UserUsecase.GetByID(uIdtmp)
//	gtUser, gtUserErr := UserUsecase.GetByToken(t)
//
//	golog.Info("gtUserErr")
//	golog.Info(gtUserErr)
//
//	poz, gtPozErr := PositionOrganizationZoneUsecase.GetByID(gtUser.OrgPosZoneId)
//	golog.Info("gtPozErr :")
//	golog.Info(gtPozErr)
//
//	po, gtPoErr := PositionOrganizationUsecase.GetByID(poz.PosOrgId)
//	golog.Info("gtPoErr :")
//	golog.Info(gtPoErr)
//
//	pozVM := make([]viewmodel.PosOrgZoneView, 1)
//	name := ""
//	pozVM[0].PosOrgZoneId = poz.Id.Hex()
//	tmpzone, gtzoneErr := ZoneUsecase.GetByID(poz.ZoneId)
//	golog.Info(gtzoneErr)
//	name += tmpzone.ZoneName + " "
//	tmppos, gtposErr := PositionUsecase.GetByID(po.PositionId)
//	//tmppos, gtposErr := PositionUsecase.GetByID(poz.PositionId)
//	golog.Info(gtposErr)
//	name += tmppos.PositionName + " "
//	tmporg, gtorgErr := OrganizationUsecase.GetByID(po.OrganizationId)
//	golog.Info(gtorgErr)
//	name += tmporg.OrganizationName
//	pozVM[0].Name = name
//
//	var tmp string
//	if gtUser.Gender == false {
//		tmp = "مرد"
//	}
//	if gtUser.Gender == true {
//		tmp = "زن"
//	}
//
//	if t != "" {
//		u, err := UserUsecase.GetByToken(t)
//
//		golog.Info(u)
//		if err != nil {
//			return mvc.View{
//				Name: "/login/index.html",
//				Data: iris.Map{"Title": "Login", "flaghshow": "0"},
//			}
//		}
//
//		return mvc.View{
//			Name: "/user/profile.html",
//			Data: iris.Map{"flaghshow": "1", "poz": name, "username": gtUser.UserName, "name": gtUser.FirstName, "lastname": gtUser.LastName, "gender": tmp},
//		}
//
//	} else {
//		return mvc.View{
//			Name: "/login/index.html",
//			Data: iris.Map{"Title": "Login", "flaghshow": "0"},
//		}
//	}
//
//}
//
//
//////////////////////////////////////////
////func (c *UserController) GetMenu() mvc.Result {
////	return mvc.View{
////		Name: "/login/empty.html",
////		Data: iris.Map{"Title": "Menu", "flaghshow": "1"},
////	}
////}
//
//
//func (c *UserController) GetAdduser() mvc.Result {
//	ZoneList, POZListErr := ZoneUsecase.GetAll()
//	golog.Info(POZListErr)
//
//	pozSize := len(ZoneList)
//	zoneVM := make([]viewmodel.ZoneView, pozSize)
//	for i := 0; i < pozSize; i++ {
//		zoneVM[i].ZoneId = ZoneList[i].Id.Hex()
//		zoneVM[i].ZoneName = ZoneList[i].ZoneName
//
//	}
//
//	return mvc.View{
//		Name: "/register/register.html",
//		Data: iris.Map{"Title": "Menu", "flaghshow": "1", "rng": zoneVM},
//	}
//}
//
//
//func (c *UserController) PostAdduser() mvc.Result {
//
//	golog.Info("Enter Post adduser : ")
//	userId := bson.NewObjectId()
//	fname := c.Ctx.FormValue("name")
//	golog.Info("first name is :")
//	golog.Info(fname)
//	lname := c.Ctx.FormValue("lastname")
//	golog.Info("last name is :")
//	golog.Info(lname)
//	var gender bool
//	if c.Ctx.FormValue("gender") == "1" {
//		gender = true
//	} else {
//		gender = false
//	}
//
//	username := c.Ctx.FormValue("username")
//	password := c.Ctx.FormValue("password")
//	pos := c.Ctx.FormValue("pos")
//	zone := c.Ctx.FormValue("zone")
//	organ := c.Ctx.FormValue("organ")
//	poz := Getpoz(organ, pos, zone) // get poz in controller
//
//	signfile, info, formFileErr := c.Ctx.FormFile("image")
//	golog.Info("formFileErr is :")
//	golog.Info(formFileErr)
//	//filename:=c.Ctx.FormValue("image")//sign
//	//golog.Info("file name is : ")
//	//golog.Info(filename)
//	if info != nil {
//
//		tempFile, ioutilErr := ioutil.TempFile("web/public/signs/", "sign"+"*."+info.Header.Get("Content-Type")[6:])
//		golog.Info("ioutilErr is :")
//		golog.Info(ioutilErr)
//		//if err != nil {
//		//	fmt.Println(err)
//		//}
//		defer tempFile.Close()
//		//golog.Info(err)
//		defer signfile.Close()
//		fileBytes, ioutilErr2 := ioutil.ReadAll(signfile)
//		golog.Info("ioutilErr2 is :")
//		golog.Info(ioutilErr2)
//		//if err2 != nil {
//		//	fmt.Println(err)
//		//}
//		tempFile.Write(fileBytes)
//		//golog.Info(err2)
//		io.Copy(tempFile, signfile)
//		a:=make(model.InErjaList,0)
//		user := model.User{Id: userId, FirstName: fname, LastName: lname, Gender: gender,
//			OrgPosZoneId: poz.Id.Hex(),Inbox:&a,
//			UserName: username, Password: password, SignatureImage: tempFile.Name(), CreatedAt: time.Now()}
//		_, saveUserErr := UserUsecase.SaveUser(&user)
//		golog.Info(saveUserErr)
//	} else {
//		user := model.User{Id: userId, FirstName: fname, LastName: lname, Gender: gender, OrgPosZoneId: poz.Id.Hex(),
//			UserName: username, Password: password, CreatedAt: time.Now()}
//		_, saveUserErr := UserUsecase.SaveUser(&user)
//		golog.Info("saveUserErr is :")
//		golog.Info(saveUserErr)
//	}
//
//	//////////////////
//
//	return mvc.Response{
//		//Path: "profile?id=" + userId.Hex(),
//		Path: "profile",
//	}
//
//}
//
///*
////localhost:8080/user/me GET
//func (c *UserController) GetMe() mvc.Result{
//	if !c.isUserLoggedIn(){
//		return  mvc.Response{
//			Path:"/user/login",
//		}
//	}
//
//	user , err :=c.UserUsecase.GetByID(c.getCurrentUserID())
//	if err!=nil{
//		c.logout()
//		c.GetMe()
//	}
//
//	return mvc.View{
//		Name:"user/me.html",
//		Data:iris.Map{"Title":"User Information" , "User":user},
//	}
//
//}
//
//
//
//func (c *UserController) AnyLogout(){
//	if c.isUserLoggedIn(){
//		c.logout()
//	}
//	c.Ctx.Redirect("/user/login")
//}
//
//
////package controller
///*
//import(
//	"v2/AbsharAutomation/v2.1/usecase"
//	"v2/AbsharAutomation/v2.1/model"
//)
//
//var UserUsecase usecase.UserUsecase
//var TokenUsecase  usecase.TokenUsecase
//var ZoneUsecase usecase.ZoneUsecase
//var OrganizationUsecase usecase.OrganizationUsecase
//var PositionUsecase usecase.PositionUsecase
//
//func GetUserByToken(Token string)(user *model.User , err error){
//
//	token , err := TokenUsecase.GetByID(Token)
//	if err != nil{
//		return nil , err
//	}
//
//	user , err = UserUsecase.GetByID(token.UserId.Hex())
//	if err != nil{
//		return nil , err
//	}
//
//	return user , err
//
//}
//
//*/
