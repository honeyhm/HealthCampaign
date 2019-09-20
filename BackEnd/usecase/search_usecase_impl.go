package usecase

import (
	"HealthCampaign/BackEnd/model"
	"HealthCampaign/BackEnd/repository"
	"github.com/kataras/golog"
)

type SearchUsecaseImpl struct {
	SearchtRepository repository.SearchRepository
}

func NewSearchUsecase(SearchRepository repository.SearchRepository) *SearchUsecaseImpl {
	//func NewUserUsecase(userRepository  repository.UserRepository) *UserUsecaseImpl {
	return &SearchUsecaseImpl{SearchRepository}
}

/////////////////////////////////////////////////////////////////////


// save Search
func (SearchUsecase *SearchUsecaseImpl) SaveSearch(Search *model.Search) (*model.Search, error) {
	//golog.Info("**************statr****************")
	err:=SearchUsecase.SearchtRepository.Save(Search)
	//golog.Info("**************finish****************")

	if err != nil {
		return nil, err
	}
	return Search, nil
}

//update Search by ObjectId
func (SearchUsecase *SearchUsecaseImpl) UpdateSearch(id string, Search *model.Search) (*model.Search, error) {
	err := SearchUsecase.SearchtRepository.Update(id, Search)
	if err != nil {
		return nil, err
	}
	return Search, nil
}

//find Search by id
func (SearchUsecase *SearchUsecaseImpl) GetByID(id string) (*model.Search, error) {
	Search, err := SearchUsecase.SearchtRepository.FindByID(id)
	//err:=UserUsecase.userRepository.FindByID(id)
	if err != nil {
		return nil, err
	}
	return Search, nil
}


func (SearchUsecase *SearchUsecaseImpl) DeleteByID(id string) error {
	err :=SearchUsecase.SearchtRepository.Delete(id)
	if err != nil {
		return err
	}
	return nil
}


// find and get all Searchs
func (SearchUsecase *SearchUsecaseImpl) GetAll() (model.Searches, error) {
	golog.Info("Enter Get All Search Usecase :")
	Searches, err := SearchUsecase.SearchtRepository.FindAll()
	golog.Info("finish !")
	if err != nil {
		return nil, err
	}
	return Searches, nil
}