/*created by H.Mlk*/
package function

import (
	"fmt"
	"AbsharAutomation/usecase"
)


//check first login of the system
func SystemLogin(username string , password string) error{  /// add token as output

	////check if there is any user with this username or password in UserCollection
	var usecase1 usecase.UserUsecase
	user,err1 := usecase1.GetByUserName(username)
	//user,err1 :=usecase.UserUsecaseImpl.GetByUserName(username)
	//user,err :=repository.UserRepositoryMongo.FindByUserName(username)
	if err1 != nil {
		return err1
	}

	////check for username==admin and password==admin in database in admin irst login
	if user.UserName =="admin" && user.Password =="admin"{

		newUsername , NewPassword ,e1 := GetNewUserPass()

		if e1 != nil {
			return  e1
		}

		user.UserName=newUsername
		user.Password=NewPassword

		////update admin with new username and new password
		_ , e := usecase1.UpdateUser(user.Id.String() ,user)

		if e != nil {
			return  e
		}

		////show user new created username and password
		ShowChangedUserPass()

		////make new token for new user
		MakeToken()

	}


	/////???????????????????????????????????????????????? not sureeee
	var usecase2 usecase.TokenUsecase
	token,err2 :=usecase2.GetByID(user.Id.String())
	if err2 != nil {
		return err2
	}

	if token.IsActive == false{
		MakeToken()
	}

	return nil
}



func GetNewUserPass()(string , string ,error){

	//ask for new user name and password from user
	newUser ,newPass ,err := AskForNewUserPass()
	if err != nil {
		return "","" ,err
	}

	return newUser,newPass,nil

}


//ask for new user name and password from user
func AskForNewUserPass()(string , string ,error){
	var s1  , s2 string
	fmt.Println("نام کاربری جدید را وارد کنید:")
	_, err1 := fmt.Scanf("%s", &s1)
	fmt.Println("رمز عبور را وارد کنید:")
	_, err2 := fmt.Scanf("%s", &s2)


	if err1 != nil {
		return "" , "" ,err1
	}

	if err2 != nil {
		return "" , "" ,err2
	}

	return s1 , s2 , nil

}



func MakeToken(){
	//model.Tokens{}
}


func ShowChangedUserPass(){

}