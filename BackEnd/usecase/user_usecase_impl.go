/*created by H.Mlk*/
package usecase

import (
	"AbsharAutomation/model"
	"AbsharAutomation/repository"
)

type UserUsecaseImpl struct {
	userRepository repository.UserRepository
}

func NewUserUsecase(userRepository repository.UserRepository) *UserUsecaseImpl {
	//func NewUserUsecase(userRepository  repository.UserRepository) *UserUsecaseImpl {
	return &UserUsecaseImpl{userRepository}
}

/////////////////////////////////////////////////////////////////////

func (UserUsecase *UserUsecaseImpl) SaveUser(user *model.User) (*model.User, error) {
	err := UserUsecase.userRepository.Save(user)
	if err != nil {
		return nil, err
	}
	return user, nil
}

//update user by ObjectId
func (UserUsecase *UserUsecaseImpl) UpdateUser(id string, user *model.User) (*model.User, error) {
	err := UserUsecase.userRepository.Update(id, user)
	if err != nil {
		return nil, err
	}
	return user, nil
}

//find user by id
func (UserUsecase *UserUsecaseImpl) GetByID(id string) (*model.User, error) {
	user, err := UserUsecase.userRepository.FindByID(id)
	//err:=UserUsecase.userRepository.FindByID(id)
	if err != nil {
		return nil, err
	}
	return user, nil
}

//find user by id
func (UserUsecase *UserUsecaseImpl) GetByPozId(id string) (*model.User, error) {
	user, err := UserUsecase.userRepository.FindByPozId(id)
	//err:=UserUsecase.userRepository.FindByID(id)
	if err != nil {
		return nil, err
	}
	return user, nil
}



func (UserUsecase *UserUsecaseImpl) GetByToken(token string) (*model.User, error) {
	user, err := UserUsecase.userRepository.FindByToken(token)
	//err:=UserUsecase.userRepository.FindByID(id)
	if err != nil {
		return nil, err
	}
	return user, nil
}

// find user according to first and last name
func (UserUsecase *UserUsecaseImpl) GetByName(FirstName string /*, LastName string*/) (*model.User, error) {
	user, err := UserUsecase.userRepository.FindByName(FirstName /*,LastName*/)
	//err:=UserUsecase.userRepository.FindByName(FirstName,LastName)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (UserUsecase *UserUsecaseImpl) GetByUserName(UserName string) (*model.User, error) {
	user, err := UserUsecase.userRepository.FindByUserName(UserName)
	//err:=UserUsecase.userRepository.FindByUserName(UserName)
	if err != nil {
		return nil, err
	}
	return user, nil

}

func (UserUsecase *UserUsecaseImpl) GetAll() (model.Users, error) {
	users, err := UserUsecase.userRepository.FindAll()
	if err != nil {
		return nil, err
	}
	return users, nil
}
