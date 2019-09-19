/*created by H.MLK*/
package usecase

import (
	"authorize/model"
	"authorize/repository"
	"github.com/kataras/golog"
)

type MedicalCenterUsecaseImpl struct {
	MedicalCentertRepository repository.MedicalCenterRepository
}

func NewMedicalCenterUsecase(MedicalCenterRepository repository.MedicalCenterRepository) *MedicalCenterUsecaseImpl {
	//func NewUserUsecase(userRepository  repository.UserRepository) *UserUsecaseImpl {
	return &MedicalCenterUsecaseImpl{MedicalCenterRepository}
}

/////////////////////////////////////////////////////////////////////


// save MedicalCenter
func (MedicalCenterUsecase *MedicalCenterUsecaseImpl) SaveMedicalCenter(MedicalCenter *model.MedicalCenter) (*model.MedicalCenter, error) {
	//golog.Info("**************statr****************")
	err:=MedicalCenterUsecase.MedicalCentertRepository.Save(MedicalCenter)
	//golog.Info("**************finish****************")

	if err != nil {
		return nil, err
	}
	return MedicalCenter, nil
}

//update MedicalCenter by ObjectId
func (MedicalCenterUsecase *MedicalCenterUsecaseImpl) UpdateMedicalCenter(id string, MedicalCenter *model.MedicalCenter) (*model.MedicalCenter, error) {
	err := MedicalCenterUsecase.MedicalCentertRepository.Update(id, MedicalCenter)
	if err != nil {
		return nil, err
	}
	return MedicalCenter, nil
}

//find MedicalCenter by id
func (MedicalCenterUsecase *MedicalCenterUsecaseImpl) GetByID(id string) (*model.MedicalCenter, error) {
	MedicalCenter, err := MedicalCenterUsecase.MedicalCentertRepository.FindByID(id)
	//err:=UserUsecase.userRepository.FindByID(id)
	if err != nil {
		return nil, err
	}
	return MedicalCenter, nil
}
//find MedicalCenter by id
//func (MedicalCenterUsecase *MedicalCenterUsecaseImpl) GetByNum(id string) (*model.MedicalCenter, error) {
//	MedicalCenter, err := MedicalCenterUsecase.MedicalCentertRepository.FindByNum(id)
//	//err:=UserUsecase.userRepository.FindByID(id)
//	if err != nil {
//		return nil, err
//	}
//	return MedicalCenter, nil
//}



func (MedicalCenterUsecase *MedicalCenterUsecaseImpl) DeleteByID(id string) error {
	err :=MedicalCenterUsecase.MedicalCentertRepository.Delete(id)
	if err != nil {
		return err
	}
	return nil
}


// find and get all MedicalCenters
func (MedicalCenterUsecase *MedicalCenterUsecaseImpl) GetAll() (model.MedicalCenters, error) {
	golog.Info("Enter Get All MedicalCenter Usecase :")
	ersons, err := MedicalCenterUsecase.MedicalCentertRepository.FindAll()
	golog.Info("finish !")
	if err != nil {
		return nil, err
	}
	return ersons, nil
}