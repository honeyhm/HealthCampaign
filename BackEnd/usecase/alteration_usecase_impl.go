/*created by H.MLK*/
package usecase

import (
	"authorize/model"
	"authorize/repository"
	"github.com/kataras/golog"
)

type AlterationUsecaseImpl struct {
	AlterationtRepository repository.AlterationRepository
}

func NewAlterationUsecase(AlterationRepository repository.AlterationRepository) *AlterationUsecaseImpl {
	//func NewUserUsecase(userRepository  repository.UserRepository) *UserUsecaseImpl {
	return &AlterationUsecaseImpl{AlterationRepository}
}

/////////////////////////////////////////////////////////////////////


// save Alteration
func (AlterationUsecase *AlterationUsecaseImpl) SaveAlteration(Alteration *model.Alteration) (*model.Alteration, error) {
	//golog.Info("**************statr****************")
	err:=AlterationUsecase.AlterationtRepository.Save(Alteration)
	//golog.Info("**************finish****************")

	if err != nil {
		return nil, err
	}
	return Alteration, nil
}

//update Alteration by ObjectId
func (AlterationUsecase *AlterationUsecaseImpl) UpdateAlteration(id string, Alteration *model.Alteration) (*model.Alteration, error) {
	err := AlterationUsecase.AlterationtRepository.Update(id, Alteration)
	if err != nil {
		return nil, err
	}
	return Alteration, nil
}

//find Alteration by id
func (AlterationUsecase *AlterationUsecaseImpl) GetByID(id string) (*model.Alteration, error) {
	Alteration, err := AlterationUsecase.AlterationtRepository.FindByID(id)
	//err:=UserUsecase.userRepository.FindByID(id)
	if err != nil {
		return nil, err
	}
	return Alteration, nil
}
//find Alteration by id
//func (AlterationUsecase *AlterationUsecaseImpl) GetByNum(id string) (*model.Alteration, error) {
//	Alteration, err := AlterationUsecase.AlterationtRepository.FindByNum(id)
//	//err:=UserUsecase.userRepository.FindByID(id)
//	if err != nil {
//		return nil, err
//	}
//	return Alteration, nil
//}



func (AlterationUsecase *AlterationUsecaseImpl) DeleteByID(id string) error {
	err :=AlterationUsecase.AlterationtRepository.Delete(id)
	if err != nil {
		return err
	}
	return nil
}


// find and get all Alterations
func (AlterationUsecase *AlterationUsecaseImpl) GetAll() (model.Alterations, error) {
	golog.Info("Enter Get All Alteration Usecase :")
	ersons, err := AlterationUsecase.AlterationtRepository.FindAll()
	golog.Info("finish !")
	if err != nil {
		return nil, err
	}
	return ersons, nil
}