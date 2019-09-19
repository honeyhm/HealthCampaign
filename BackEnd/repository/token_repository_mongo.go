
package repository

import (
	"authorize/model"
	_ "fmt"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type TokenRepositoryMongo struct {
	db         *mgo.Database
	collection string
}

func NewTokenRepositoryMongo(db *mgo.Database, collection string) *TokenRepositoryMongo {
	return &TokenRepositoryMongo{
		db:         db,
		collection: collection,
	}
}

func (r *TokenRepositoryMongo) Save(token *model.Token) error {
	err := r.db.C(r.collection).Insert(token)
	//println(err)
	return err
}

//update token by ObjectId
func (r *TokenRepositoryMongo) Update(id string, token *model.Token) error {
	//token.UpdatedAt = time.Now()
	//token.UpdateAt =time.Now()
	err := r.db.C(r.collection).UpdateId(bson.ObjectIdHex(id), token)
	//err := r.db.C(r.collection).UpdateId(bson.M{"_id": id}, token)
	return err

}

/*func(r *TokenRepositoryMongo) UpdateByTokenPass(tokenname string ,password string,id string) error{
	//token.UpdatedAt=time.Now()
	//token.UpdateAt =time.Now()
	//err := r.db.C(r.collection).UpdateId(bson.ObjectIdHex(id),token)
	err := r.db.C(r.collection).Update({"id":id},{"$set":{{"tokenname":tokenname},{"password":password}}})
	return err

}*/

//delete token by id
func (r *TokenRepositoryMongo) Delete(id string) error {
	err := r.db.C(r.collection).Remove(bson.M{"_id": id})
	return err
}

//find token by id
func (r *TokenRepositoryMongo) FindByID(id string) (*model.Token, error) {
	var token model.Token
	//err := r.db.C(r.collection).Find(bson.M{"_id": id}).One(&token)
	err := r.db.C(r.collection).FindId(bson.ObjectIdHex(id)).One(&token)

	if err != nil {
		return nil, err
	}

	return &token, nil
}


func (r *TokenRepositoryMongo) FindByToken(tokenToken string ) (*model.Token, error) {
	var token model.Token
	err := r.db.C(r.collection).Find(bson.M{"token_token": tokenToken }).One(&token)

	if err != nil {
		return nil, err
	}

	return &token, nil

}



func (r *TokenRepositoryMongo) FindByPozId(poz string ) (*model.Token, error) {
	var token model.Token
	err := r.db.C(r.collection).Find(bson.M{"org_pos_zone_id": poz }).One(&token)

	if err != nil {
		return nil, err
	}

	return &token, nil

}




// find token according to first and last name
func (r *TokenRepositoryMongo) FindByName(FirstName string /*, LastName string*/) (*model.Token, error) {
	var token model.Token
	err := r.db.C(r.collection).Find(bson.M{"first_name": FirstName /*,"last_name":LastName*/}).One(&token) //////is it and??

	if err != nil {
		return nil, err
	}

	return &token, nil

}

func (r *TokenRepositoryMongo) FindByTokenName(TokenName string) (*model.Token, error) {
	var token model.Token
	err := r.db.C(r.collection).Find(bson.M{"token_name": TokenName}).One(&token)

	if err != nil {
		return nil, err
	}

	return &token, nil

}

/*//find and return all tokens
func(r *TokenRepositoryMongo) FindAll()(model.Tokens , error){
	var tokens model.Tokens

	//err := r.db.C(r.collection).Find(nil).All(&tokens)//too on yeki project injoori bood
	err := r.db.C(r.collection).Find(bson.M{}).All(&tokens)

	if err != nil {
		return nil ,err
	}

	return tokens , nil

}*/

//find and return all tokens
func (r *TokenRepositoryMongo) FindAll() (model.Tokens, error) {
	var tokens model.Tokens

	//err := r.db.C(r.collection).Find(nil).All(&tokens)//too on yeki project injoori bood
	err := r.db.C(r.collection).Find(bson.M{}).All(&tokens)

	if err != nil {
		return nil, err
	}

	return tokens, nil

}

/*func(r *TokenRepositoryMongo) CheckToken (TokenName string ,Password string) (bool){

	err := r.db.C(r.collection).Find(bson.M{"tokenname":TokenName , "password":Password})

	if err != nil {
		return false
	}

	return true
	//fmt.Println(TokenName)

}


*/
