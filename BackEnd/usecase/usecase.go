
package usecase

import (
	"authorize/model"
)


type PersonUsecase interface {
	SavePerson(user *model.Person) (*model.Person, error)
	UpdatePerson(string, *model.Person) (*model.Person, error)
	GetByID(string) (*model.Person, error)
	GetAll() (model.Persons, error)
	GetByNum(string) (*model.Person, error)
}


type UserUsecase interface {
	SaveUser(user *model.User) (*model.User, error)
	UpdateUser(string, *model.User) (*model.User, error)
	GetByID(string) (*model.User, error)
	GetByToken(string) (*model.User, error)
	GetByName(string) (*model.User, error)
	GetByUserName(string) (*model.User, error)
	GetAll() (model.Users, error)
	GetByPozId( string) (*model.User, error)
}

type TokenUsecase interface {
	SaveToken(user *model.Token) (*model.Token, error)
	UpdateToken(string, *model.Token) (*model.Token, error)
	GetByID(string) (*model.Token, error)
	DeleteByID(string) error
	GetAll() (model.Tokens, error)
}
