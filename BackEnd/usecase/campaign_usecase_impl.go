/*created by H.MLK*/
package usecase

import (
	"HealthCampaign/BackEnd/model"
	"HealthCampaign/BackEnd/repository"
	"github.com/kataras/golog"
)

type CampaignUsecaseImpl struct {
	CampaigntRepository repository.CampaignRepository
}

func NewCampaignUsecase(CampaignRepository repository.CampaignRepository) *CampaignUsecaseImpl {
	//func NewUserUsecase(userRepository  repository.UserRepository) *UserUsecaseImpl {
	return &CampaignUsecaseImpl{CampaignRepository}
}

/////////////////////////////////////////////////////////////////////


// save Campaign
func (CampaignUsecase *CampaignUsecaseImpl) SaveCampaign(Campaign *model.Campaign) (*model.Campaign, error) {
	//golog.Info("**************statr****************")
	err:=CampaignUsecase.CampaigntRepository.Save(Campaign)
	//golog.Info("**************finish****************")

	if err != nil {
		return nil, err
	}
	return Campaign, nil
}

//update Campaign by ObjectId
func (CampaignUsecase *CampaignUsecaseImpl) UpdateCampaign(id string, Campaign *model.Campaign) (*model.Campaign, error) {
	err := CampaignUsecase.CampaigntRepository.Update(id, Campaign)
	if err != nil {
		return nil, err
	}
	return Campaign, nil
}

//find Campaign by id
func (CampaignUsecase *CampaignUsecaseImpl) GetByID(id string) (*model.Campaign, error) {
	Campaign, err := CampaignUsecase.CampaigntRepository.FindByID(id)
	//err:=UserUsecase.userRepository.FindByID(id)
	if err != nil {
		return nil, err
	}
	return Campaign, nil
}
//find Campaign by id
//func (CampaignUsecase *CampaignUsecaseImpl) GetByNum(id string) (*model.Campaign, error) {
//	Campaign, err := CampaignUsecase.CampaigntRepository.FindByNum(id)
//	//err:=UserUsecase.userRepository.FindByID(id)
//	if err != nil {
//		return nil, err
//	}
//	return Campaign, nil
//}



func (CampaignUsecase *CampaignUsecaseImpl) DeleteByID(id string) error {
	err :=CampaignUsecase.CampaigntRepository.Delete(id)
	if err != nil {
		return err
	}
	return nil
}


// find and get all Campaigns
func (CampaignUsecase *CampaignUsecaseImpl) GetAll() (model.Campaigns, error) {
	golog.Info("Enter Get All Campaign Usecase :")
	ersons, err := CampaignUsecase.CampaigntRepository.FindAll()
	golog.Info("finish !")
	if err != nil {
		return nil, err
	}
	return ersons, nil
}