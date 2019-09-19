/*created by H.Mlk*/

package repository

import (
	_ "fmt"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"authorize/model"
)

type PersonRepositoryMongo struct{
	db	*mgo.Database
	collection string
}


func NewPersonRepositoryMongo(db *mgo.Database, collection string) *PersonRepositoryMongo {
	return &PersonRepositoryMongo{
		db:	db,
		collection:	collection,
	}
}


//save Person
func(r *PersonRepositoryMongo) Save(Person *model.Person) error{
	err := r.db.C(r.collection).Insert(Person)
	return err
}

//update Person by ObjectId
func(r *PersonRepositoryMongo) Update(id string , Person *model.Person) error{
	err := r.db.C(r.collection).UpdateId(bson.ObjectIdHex(id),Person)
	//err := r.db.C(r.collection).UpdateId(bson.M{"_id":id},agreement)
	return err

}

//delete Person by id
func(r *PersonRepositoryMongo) Delete(id string) error{
	err := r.db.C(r.collection).Remove(bson.M{"_id": id})
	return err
}

//find Person by id
func(r *PersonRepositoryMongo) FindByID(id string) (*model.Person ,error){
	var Person model.Person
	//err := r.db.C(r.collection).Find(bson.M{"_id": id}).One(&agreement)
	err := r.db.C(r.collection).FindId(bson.ObjectIdHex(id)).One(&Person)

	if err != nil {
		return nil ,err
	}

	return &Person,nil
}

//find Person by id
func(r *PersonRepositoryMongo) FindByNum(num string) (*model.Person ,error){
	var Person model.Person
	//err := r.db.C(r.collection).Find(bson.M{"_id": id}).One(&agreement)
	err := r.db.C(r.collection).Find(bson.M{"student_num": num}).One(&Person)

	if err != nil {
		return nil ,err
	}

	return &Person,nil
}







//find and return all Persons
func(r *PersonRepositoryMongo) FindAll()(model.Persons , error){
	var Persons model.Persons

	//err := r.db.C(r.collection).Find(nil).All(&users)//too on yeki project injoori bood
	err := r.db.C(r.collection).Find(bson.M{}).All(&Persons)

	if err != nil {
		return nil ,err
	}

	return Persons , nil

}



