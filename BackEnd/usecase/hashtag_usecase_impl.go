package usecase

import (
	"HealthCampaign/BackEnd/model"
	"HealthCampaign/BackEnd/repository"
	"github.com/kataras/golog"
)

type HashtagUsecaseImpl struct {
	HashtagtRepository repository.HashtagRepository
}

func NewHashtagUsecase(HashtagRepository repository.HashtagRepository) *HashtagUsecaseImpl {
	//func NewUserUsecase(userRepository  repository.UserRepository) *UserUsecaseImpl {
	return &HashtagUsecaseImpl{HashtagRepository}
}

/////////////////////////////////////////////////////////////////////


// save Hashtag
func (HashtagUsecase *HashtagUsecaseImpl) SaveHashtag(Hashtag *model.Hashtag) (*model.Hashtag, error) {
	//golog.Info("**************statr****************")
	err:=HashtagUsecase.HashtagtRepository.Save(Hashtag)
	//golog.Info("**************finish****************")

	if err != nil {
		return nil, err
	}
	return Hashtag, nil
}

//update Hashtag by ObjectId
func (HashtagUsecase *HashtagUsecaseImpl) UpdateHashtag(id string, Hashtag *model.Hashtag) (*model.Hashtag, error) {
	err := HashtagUsecase.HashtagtRepository.Update(id, Hashtag)
	if err != nil {
		return nil, err
	}
	return Hashtag, nil
}

//find Hashtag by id
func (HashtagUsecase *HashtagUsecaseImpl) GetByID(id string) (*model.Hashtag, error) {
	Hashtag, err := HashtagUsecase.HashtagtRepository.FindByID(id)
	//err:=UserUsecase.userRepository.FindByID(id)
	if err != nil {
		return nil, err
	}
	return Hashtag, nil
}


func (HashtagUsecase *HashtagUsecaseImpl) GetByName(n string) (*model.Hashtag, error) {
	Hashtag, err := HashtagUsecase.HashtagtRepository.FindByName(n)
	if err != nil {
		return nil, err
	}
	return Hashtag, nil
}



func (HashtagUsecase *HashtagUsecaseImpl) DeleteByID(id string) error {
	err :=HashtagUsecase.HashtagtRepository.Delete(id)
	if err != nil {
		return err
	}
	return nil
}


// find and get all Hashtags
func (HashtagUsecase *HashtagUsecaseImpl) GetAll() (model.Hashtags, error) {
	golog.Info("Enter Get All Hashtag Usecase :")
	Hashtags, err := HashtagUsecase.HashtagtRepository.FindAll()
	golog.Info("finish !")
	if err != nil {
		return nil, err
	}
	return Hashtags, nil
}