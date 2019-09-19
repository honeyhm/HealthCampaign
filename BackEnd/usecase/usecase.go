
package usecase

import (
	"authorize/model"
)


type PersonUsecase interface {
	SavePerson(user *model.Person) (*model.Person, error)
	UpdatePerson(string, *model.Person) (*model.Person, error)
	GetByID(string) (*model.Person, error)
	GetAll() (model.Persons, error)
	//GetByNum(string) (*model.Person, error)
}


type DiseaseUsecase interface {
	SaveDisease(user *model.Disease) (*model.Disease, error)
	UpdateDisease(string, *model.Disease) (*model.Disease, error)
	GetByID(string) (*model.Disease, error)
	GetAll() (model.Diseases, error)
	//GetByNum(string) (*model.Disease, error)
}


type CampaignUsecase interface {
	SaveCampaign(user *model.Campaign) (*model.Campaign, error)
	UpdateCampaign(string, *model.Campaign) (*model.Campaign, error)
	GetByID(string) (*model.Campaign, error)
	GetAll() (model.Campaigns, error)
	//GetByNum(string) (*model.Campaign, error)
}



type AlterationUsecase interface {
	SaveAlteration(user *model.Alteration) (*model.Alteration, error)
	UpdateAlteration(string, *model.Alteration) (*model.Alteration, error)
	GetByID(string) (*model.Alteration, error)
	GetAll() (model.Alterations, error)
	//GetByNum(string) (*model.Alteration, error)
}


type PatientUsecase interface {
	SavePatient(user *model.Patient) (*model.Patient, error)
	UpdatePatient(string, *model.Patient) (*model.Patient, error)
	GetByID(string) (*model.Patient, error)
	GetAll() (model.Patients, error)
	//GetByNum(string) (*model.Patient, error)
}

//
//type UserUsecase interface {
//	SaveUser(user *model.User) (*model.User, error)
//	UpdateUser(string, *model.User) (*model.User, error)
//	GetByID(string) (*model.User, error)
//	GetByToken(string) (*model.User, error)
//	GetByName(string) (*model.User, error)
//	GetByUserName(string) (*model.User, error)
//	GetAll() (model.Users, error)
//	GetByPozId( string) (*model.User, error)
//}
//
//type TokenUsecase interface {
//	SaveToken(user *model.Token) (*model.Token, error)
//	UpdateToken(string, *model.Token) (*model.Token, error)
//	GetByID(string) (*model.Token, error)
//	DeleteByID(string) error
//	GetAll() (model.Tokens, error)
//}
