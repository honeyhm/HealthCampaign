/*created by H.Mlk*/
package repository

import (
	"authorize/model"
	_ "fmt"
	"time"

	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type UserRepositoryMongo struct {
	db         *mgo.Database
	collection string
}

func NewUserRepositoryMongo(db *mgo.Database, collection string) *UserRepositoryMongo {
	return &UserRepositoryMongo{
		db:         db,
		collection: collection,
	}
}

func (r *UserRepositoryMongo) Save(user *model.User) error {
	err := r.db.C(r.collection).Insert(user)
	//println(err)
	return err
}

//update user by ObjectId
func (r *UserRepositoryMongo) Update(id string, user *model.User) error {
	user.UpdatedAt = time.Now()
	//user.UpdateAt =time.Now()
	err := r.db.C(r.collection).UpdateId(bson.ObjectIdHex(id), user)
	//err := r.db.C(r.collection).UpdateId(bson.M{"_id": id}, user)
	return err

}

/*func(r *UserRepositoryMongo) UpdateByUserPass(username string ,password string,id string) error{
	//user.UpdatedAt=time.Now()
	//user.UpdateAt =time.Now()
	//err := r.db.C(r.collection).UpdateId(bson.ObjectIdHex(id),user)
	err := r.db.C(r.collection).Update({"id":id},{"$set":{{"username":username},{"password":password}}})
	return err

}*/

//delete user by id
func (r *UserRepositoryMongo) Delete(id string) error {
	err := r.db.C(r.collection).Remove(bson.M{"_id": id})
	return err
}

//find user by id
func (r *UserRepositoryMongo) FindByID(id string) (*model.User, error) {
	var user model.User
	//err := r.db.C(r.collection).Find(bson.M{"_id": id}).One(&user)
	err := r.db.C(r.collection).FindId(bson.ObjectIdHex(id)).One(&user)

	if err != nil {
		return nil, err
	}

	return &user, nil
}


func (r *UserRepositoryMongo) FindByToken(userToken string ) (*model.User, error) {
	var user model.User
	err := r.db.C(r.collection).Find(bson.M{"user_token": userToken }).One(&user)

	if err != nil {
		return nil, err
	}

	return &user, nil

}



func (r *UserRepositoryMongo) FindByPozId(poz string ) (*model.User, error) {
	var user model.User
	err := r.db.C(r.collection).Find(bson.M{"org_pos_zone_id": poz }).One(&user)

	if err != nil {
		return nil, err
	}

	return &user, nil

}




// find user according to first and last name
func (r *UserRepositoryMongo) FindByName(FirstName string /*, LastName string*/) (*model.User, error) {
	var user model.User
	err := r.db.C(r.collection).Find(bson.M{"first_name": FirstName /*,"last_name":LastName*/}).One(&user) //////is it and??

	if err != nil {
		return nil, err
	}

	return &user, nil

}

func (r *UserRepositoryMongo) FindByUserName(UserName string) (*model.User, error) {
	var user model.User
	err := r.db.C(r.collection).Find(bson.M{"user_name": UserName}).One(&user)

	if err != nil {
		return nil, err
	}

	return &user, nil

}

/*//find and return all users
func(r *UserRepositoryMongo) FindAll()(model.Users , error){
	var users model.Users

	//err := r.db.C(r.collection).Find(nil).All(&users)//too on yeki project injoori bood
	err := r.db.C(r.collection).Find(bson.M{}).All(&users)

	if err != nil {
		return nil ,err
	}

	return users , nil

}*/

//find and return all users
func (r *UserRepositoryMongo) FindAll() (model.Users, error) {
	var users model.Users

	//err := r.db.C(r.collection).Find(nil).All(&users)//too on yeki project injoori bood
	err := r.db.C(r.collection).Find(bson.M{}).All(&users)

	if err != nil {
		return nil, err
	}

	return users, nil

}

/*func(r *UserRepositoryMongo) CheckUser (UserName string ,Password string) (bool){

	err := r.db.C(r.collection).Find(bson.M{"username":UserName , "password":Password})

	if err != nil {
		return false
	}

	return true
	//fmt.Println(UserName)

}


*/
