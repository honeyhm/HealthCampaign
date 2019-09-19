
package usecase

import (
	"HealthCampaign/BackEnd/model"
	"HealthCampaign/BackEnd/repository"
	"github.com/kataras/golog"
)

type MedicalStaffUsecaseImpl struct {
	MedicalStafftRepository repository.MedicalStaffRepository
}


func NewMedicalStaffUsecase(MedicalStaffRepository repository.MedicalStaffRepository) *MedicalStaffUsecaseImpl {
	//func NewUserUsecase(userRepository  repository.UserRepository) *UserUsecaseImpl {
	return &MedicalStaffUsecaseImpl{MedicalStaffRepository}
}

/////////////////////////////////////////////////////////////////////


// save MedicalStaff
func (MedicalStaffUsecase *MedicalStaffUsecaseImpl) SaveMedicalStaff(MedicalStaff *model.MedicalStaff) (*model.MedicalStaff, error) {
	//golog.Info("**************statr****************")
	err:=MedicalStaffUsecase.MedicalStafftRepository.Save(MedicalStaff)
	//golog.Info("**************finish****************")

	if err != nil {
		return nil, err
	}
	return MedicalStaff, nil
}

//update MedicalStaff by ObjectId
func (MedicalStaffUsecase *MedicalStaffUsecaseImpl) UpdateMedicalStaff(id string, MedicalStaff *model.MedicalStaff) (*model.MedicalStaff, error) {
	err := MedicalStaffUsecase.MedicalStafftRepository.Update(id, MedicalStaff)
	if err != nil {
		return nil, err
	}
	return MedicalStaff, nil
}

//find MedicalStaff by id
func (MedicalStaffUsecase *MedicalStaffUsecaseImpl) GetByID(id string) (*model.MedicalStaff, error) {
	MedicalStaff, err := MedicalStaffUsecase.MedicalStafftRepository.FindByID(id)
	//err:=UserUsecase.userRepository.FindByID(id)
	if err != nil {
		return nil, err
	}
	return MedicalStaff, nil
}
//find MedicalStaff by id
//func (MedicalStaffUsecase *MedicalStaffUsecaseImpl) GetByNum(id string) (*model.MedicalStaff, error) {
//	MedicalStaff, err := MedicalStaffUsecase.MedicalStafftRepository.FindByNum(id)
//	//err:=UserUsecase.userRepository.FindByID(id)
//	if err != nil {
//		return nil, err
//	}
//	return MedicalStaff, nil
//}



func (MedicalStaffUsecase *MedicalStaffUsecaseImpl) DeleteByID(id string) error {
	err :=MedicalStaffUsecase.MedicalStafftRepository.Delete(id)
	if err != nil {
		return err
	}
	return nil
}


func (MedicalStaffUsecase *MedicalStaffUsecaseImpl) GetByUserPass(u string , p string) bool {
	err := MedicalStaffUsecase.MedicalStafftRepository.FindByUserPass(u , p)
	return err
}


// find and get all MedicalStaffs
func (MedicalStaffUsecase *MedicalStaffUsecaseImpl) GetAll() (model.MedicalStaffs, error) {
	golog.Info("Enter Get All MedicalStaff Usecase :")
	ersons, err := MedicalStaffUsecase.MedicalStafftRepository.FindAll()
	golog.Info("finish !")
	if err != nil {
		return nil, err
	}
	return ersons, nil
}