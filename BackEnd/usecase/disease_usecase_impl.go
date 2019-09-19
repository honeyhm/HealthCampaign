/*created by H.MLK*/
package usecase

import (
	"authorize/model"
	"authorize/repository"
	"github.com/kataras/golog"
)

type DiseaseUsecaseImpl struct {
	DiseasetRepository repository.DiseaseRepository
}

func NewDiseaseUsecase(DiseaseRepository repository.DiseaseRepository) *DiseaseUsecaseImpl {
	//func NewUserUsecase(userRepository  repository.UserRepository) *UserUsecaseImpl {
	return &DiseaseUsecaseImpl{DiseaseRepository}
}

/////////////////////////////////////////////////////////////////////


// save Disease
func (DiseaseUsecase *DiseaseUsecaseImpl) SaveDisease(Disease *model.Disease) (*model.Disease, error) {
	//golog.Info("**************statr****************")
	err:=DiseaseUsecase.DiseasetRepository.Save(Disease)
	//golog.Info("**************finish****************")

	if err != nil {
		return nil, err
	}
	return Disease, nil
}

//update Disease by ObjectId
func (DiseaseUsecase *DiseaseUsecaseImpl) UpdateDisease(id string, Disease *model.Disease) (*model.Disease, error) {
	err := DiseaseUsecase.DiseasetRepository.Update(id, Disease)
	if err != nil {
		return nil, err
	}
	return Disease, nil
}

//find Disease by id
func (DiseaseUsecase *DiseaseUsecaseImpl) GetByID(id string) (*model.Disease, error) {
	Disease, err := DiseaseUsecase.DiseasetRepository.FindByID(id)
	//err:=UserUsecase.userRepository.FindByID(id)
	if err != nil {
		return nil, err
	}
	return Disease, nil
}
//find Disease by id
//func (DiseaseUsecase *DiseaseUsecaseImpl) GetByNum(id string) (*model.Disease, error) {
//	Disease, err := DiseaseUsecase.DiseasetRepository.FindByNum(id)
//	//err:=UserUsecase.userRepository.FindByID(id)
//	if err != nil {
//		return nil, err
//	}
//	return Disease, nil
//}



func (DiseaseUsecase *DiseaseUsecaseImpl) DeleteByID(id string) error {
	err :=DiseaseUsecase.DiseasetRepository.Delete(id)
	if err != nil {
		return err
	}
	return nil
}


// find and get all Diseases
func (DiseaseUsecase *DiseaseUsecaseImpl) GetAll() (model.Diseases, error) {
	golog.Info("Enter Get All Disease Usecase :")
	ersons, err := DiseaseUsecase.DiseasetRepository.FindAll()
	golog.Info("finish !")
	if err != nil {
		return nil, err
	}
	return ersons, nil
}