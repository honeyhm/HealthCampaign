/*created by H.MLK*/
package usecase

import (
	"authorize/model"
	"authorize/repository"
	"github.com/kataras/golog"
)

type PersonUsecaseImpl struct {
	PersontRepository repository.PersonRepository
}

func NewPersonUsecase(PersonRepository repository.PersonRepository) *PersonUsecaseImpl {
	//func NewUserUsecase(userRepository  repository.UserRepository) *UserUsecaseImpl {
	return &PersonUsecaseImpl{PersonRepository}
}

/////////////////////////////////////////////////////////////////////


// save Person
func (PersonUsecase *PersonUsecaseImpl) SavePerson(Person *model.Person) (*model.Person, error) {
	//golog.Info("**************statr****************")
	err:=PersonUsecase.PersontRepository.Save(Person)
	//golog.Info("**************finish****************")

	if err != nil {
		return nil, err
	}
	return Person, nil
}

//update Person by ObjectId
func (PersonUsecase *PersonUsecaseImpl) UpdatePerson(id string, Person *model.Person) (*model.Person, error) {
	err := PersonUsecase.PersontRepository.Update(id, Person)
	if err != nil {
		return nil, err
	}
	return Person, nil
}

//find Person by id
func (PersonUsecase *PersonUsecaseImpl) GetByID(id string) (*model.Person, error) {
	Person, err := PersonUsecase.PersontRepository.FindByID(id)
	//err:=UserUsecase.userRepository.FindByID(id)
	if err != nil {
		return nil, err
	}
	return Person, nil
}
//find Person by id
func (PersonUsecase *PersonUsecaseImpl) GetByNum(id string) (*model.Person, error) {
	Person, err := PersonUsecase.PersontRepository.FindByNum(id)
	//err:=UserUsecase.userRepository.FindByID(id)
	if err != nil {
		return nil, err
	}
	return Person, nil
}



func (PersonUsecase *PersonUsecaseImpl) DeleteByID(id string) error {
	err :=PersonUsecase.PersontRepository.Delete(id)
	if err != nil {
		return err
	}
	return nil
}


// find and get all Persons
func (PersonUsecase *PersonUsecaseImpl) GetAll() (model.Persons, error) {
	golog.Info("Enter Get All Person Usecase :")
	ersons, err := PersonUsecase.PersontRepository.FindAll()
	golog.Info("finish !")
	if err != nil {
		return nil, err
	}
	return ersons, nil
}