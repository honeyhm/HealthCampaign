package usecase

import (
	"HealthCampaign/BackEnd/model"
	"HealthCampaign/BackEnd/repository"
	"github.com/kataras/golog"
)

type GroupUsecaseImpl struct {
	GrouptRepository repository.GroupRepository
}

func NewGroupUsecase(GroupRepository repository.GroupRepository) *GroupUsecaseImpl {
	//func NewUserUsecase(userRepository  repository.UserRepository) *UserUsecaseImpl {
	return &GroupUsecaseImpl{GroupRepository}
}

/////////////////////////////////////////////////////////////////////


// save Group
func (GroupUsecase *GroupUsecaseImpl) SaveGroup(Group *model.Group) (*model.Group, error) {
	//golog.Info("**************statr****************")
	err:=GroupUsecase.GrouptRepository.Save(Group)
	//golog.Info("**************finish****************")

	if err != nil {
		return nil, err
	}
	return Group, nil
}

//update Group by ObjectId
func (GroupUsecase *GroupUsecaseImpl) UpdateGroup(id string, Group *model.Group) (*model.Group, error) {
	err := GroupUsecase.GrouptRepository.Update(id, Group)
	if err != nil {
		return nil, err
	}
	return Group, nil
}

//find Group by id
func (GroupUsecase *GroupUsecaseImpl) GetByID(id string) (*model.Group, error) {
	Group, err := GroupUsecase.GrouptRepository.FindByID(id)
	//err:=UserUsecase.userRepository.FindByID(id)
	if err != nil {
		return nil, err
	}
	return Group, nil
}
//find Group by id
//func (GroupUsecase *GroupUsecaseImpl) GetByNum(id string) (*model.Group, error) {
//	Group, err := GroupUsecase.GrouptRepository.FindByNum(id)
//	//err:=UserUsecase.userRepository.FindByID(id)
//	if err != nil {
//		return nil, err
//	}
//	return Group, nil
//}



func (GroupUsecase *GroupUsecaseImpl) DeleteByID(id string) error {
	err :=GroupUsecase.GrouptRepository.Delete(id)
	if err != nil {
		return err
	}
	return nil
}


// find and get all Groups
func (GroupUsecase *GroupUsecaseImpl) GetAll() (model.Groups, error) {
	golog.Info("Enter Get All Group Usecase :")
	Groups, err := GroupUsecase.GrouptRepository.FindAll()
	golog.Info("finish !")
	if err != nil {
		return nil, err
	}
	return Groups, nil
}