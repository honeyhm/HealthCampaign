package usecase

import (
	"HealthCampaign/BackEnd/model"
	"HealthCampaign/BackEnd/repository"
	"github.com/kataras/golog"
)

type ProgressUsecaseImpl struct {
	ProgresstRepository repository.ProgressRepository
}

func NewProgressUsecase(ProgressRepository repository.ProgressRepository) *ProgressUsecaseImpl {
	//func NewUserUsecase(userRepository  repository.UserRepository) *UserUsecaseImpl {
	return &ProgressUsecaseImpl{ProgressRepository}
}

/////////////////////////////////////////////////////////////////////


// save Progress
func (ProgressUsecase *ProgressUsecaseImpl) SaveProgress(Progress *model.Progress) (*model.Progress, error) {
	//golog.Info("**************statr****************")
	err:=ProgressUsecase.ProgresstRepository.Save(Progress)
	//golog.Info("**************finish****************")

	if err != nil {
		return nil, err
	}
	return Progress, nil
}

//update Progress by ObjectId
func (ProgressUsecase *ProgressUsecaseImpl) UpdateProgress(id string, Progress *model.Progress) (*model.Progress, error) {
	err := ProgressUsecase.ProgresstRepository.Update(id, Progress)
	if err != nil {
		return nil, err
	}
	return Progress, nil
}

//find Progress by id
func (ProgressUsecase *ProgressUsecaseImpl) GetByID(id string) (*model.Progress, error) {
	Progress, err := ProgressUsecase.ProgresstRepository.FindByID(id)
	//err:=UserUsecase.userRepository.FindByID(id)
	if err != nil {
		return nil, err
	}
	return Progress, nil
}
//find Progress by id
//func (ProgressUsecase *ProgressUsecaseImpl) GetByNum(id string) (*model.Progress, error) {
//	Progress, err := ProgressUsecase.ProgresstRepository.FindByNum(id)
//	//err:=UserUsecase.userRepository.FindByID(id)
//	if err != nil {
//		return nil, err
//	}
//	return Progress, nil
//}



func (ProgressUsecase *ProgressUsecaseImpl) DeleteByID(id string) error {
	err :=ProgressUsecase.ProgresstRepository.Delete(id)
	if err != nil {
		return err
	}
	return nil
}


// find and get all Progresss
func (ProgressUsecase *ProgressUsecaseImpl) GetAll() (model.Progresses, error) {
	golog.Info("Enter Get All Progress Usecase :")
	Progresses, err := ProgressUsecase.ProgresstRepository.FindAll()
	golog.Info("finish !")
	if err != nil {
		return nil, err
	}
	return Progresses, nil
}