/*created by H.MLK*/
package usecase

import (
	"authorize/model"
	"authorize/repository"
	"github.com/kataras/golog"
)

type PatientUsecaseImpl struct {
	PatienttRepository repository.PatientRepository
}

func NewPatientUsecase(PatientRepository repository.PatientRepository) *PatientUsecaseImpl {
	//func NewUserUsecase(userRepository  repository.UserRepository) *UserUsecaseImpl {
	return &PatientUsecaseImpl{PatientRepository}
}

/////////////////////////////////////////////////////////////////////


// save Patient
func (PatientUsecase *PatientUsecaseImpl) SavePatient(Patient *model.Patient) (*model.Patient, error) {
	//golog.Info("**************statr****************")
	err:=PatientUsecase.PatienttRepository.Save(Patient)
	//golog.Info("**************finish****************")

	if err != nil {
		return nil, err
	}
	return Patient, nil
}

//update Patient by ObjectId
func (PatientUsecase *PatientUsecaseImpl) UpdatePatient(id string, Patient *model.Patient) (*model.Patient, error) {
	err := PatientUsecase.PatienttRepository.Update(id, Patient)
	if err != nil {
		return nil, err
	}
	return Patient, nil
}

//find Patient by id
func (PatientUsecase *PatientUsecaseImpl) GetByID(id string) (*model.Patient, error) {
	Patient, err := PatientUsecase.PatienttRepository.FindByID(id)
	//err:=UserUsecase.userRepository.FindByID(id)
	if err != nil {
		return nil, err
	}
	return Patient, nil
}
//find Patient by id
//func (PatientUsecase *PatientUsecaseImpl) GetByNum(id string) (*model.Patient, error) {
//	Patient, err := PatientUsecase.PatienttRepository.FindByNum(id)
//	//err:=UserUsecase.userRepository.FindByID(id)
//	if err != nil {
//		return nil, err
//	}
//	return Patient, nil
//}



func (PatientUsecase *PatientUsecaseImpl) DeleteByID(id string) error {
	err :=PatientUsecase.PatienttRepository.Delete(id)
	if err != nil {
		return err
	}
	return nil
}


// find and get all Patients
func (PatientUsecase *PatientUsecaseImpl) GetAll() (model.Patients, error) {
	golog.Info("Enter Get All Patient Usecase :")
	ersons, err := PatientUsecase.PatienttRepository.FindAll()
	golog.Info("finish !")
	if err != nil {
		return nil, err
	}
	return ersons, nil
}