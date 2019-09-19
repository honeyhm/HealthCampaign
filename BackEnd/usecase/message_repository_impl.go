package usecase

import (
	"HealthCampaign/BackEnd/model"
	"HealthCampaign/BackEnd/repository"
	"github.com/kataras/golog"
)

type MessageUsecaseImpl struct {
	MessageRepository repository.MessageRepository
}

func NewMessageUsecase(MessageRepository repository.MessageRepository) *MessageUsecaseImpl {
	//func NewUserUsecase(userRepository  repository.UserRepository) *UserUsecaseImpl {
	return &MessageUsecaseImpl{MessageRepository}
}

/////////////////////////////////////////////////////////////////////


// save Message
func (MessageUsecase *MessageUsecaseImpl) SaveMessage(Message *model.Message) (*model.Message, error) {
	//golog.Info("**************statr****************")
	err:=MessageUsecase.MessageRepository.Save(Message)
	//golog.Info("**************finish****************")

	if err != nil {
		return nil, err
	}
	return Message, nil
}

//update Message by ObjectId
func (MessageUsecase *MessageUsecaseImpl) UpdateMessage(id string, Message *model.Message) (*model.Message, error) {
	err := MessageUsecase.MessageRepository.Update(id, Message)
	if err != nil {
		return nil, err
	}
	return Message, nil
}

//find Message by id
func (MessageUsecase *MessageUsecaseImpl) GetByID(id string) (*model.Message, error) {
	Message, err := MessageUsecase.MessageRepository.FindByID(id)
	//err:=UserUsecase.userRepository.FindByID(id)
	if err != nil {
		return nil, err
	}
	return Message, nil
}
//find Message by id
//func (MessageUsecase *MessageUsecaseImpl) GetByNum(id string) (*model.Message, error) {
//	Message, err := MessageUsecase.MessagetRepository.FindByNum(id)
//	//err:=UserUsecase.userRepository.FindByID(id)
//	if err != nil {
//		return nil, err
//	}
//	return Message, nil
//}



func (MessageUsecase *MessageUsecaseImpl) DeleteByID(id string) error {
	err :=MessageUsecase.MessageRepository.Delete(id)
	if err != nil {
		return err
	}
	return nil
}


// find and get all Messages
func (MessageUsecase *MessageUsecaseImpl) GetAll() (model.Messages, error) {
	golog.Info("Enter Get All Message Usecase :")
	Messages, err := MessageUsecase.MessageRepository.FindAll()
	golog.Info("finish !")
	if err != nil {
		return nil, err
	}
	return Messages, nil
}