
package usecase

import (
	"authorize/model"
	"authorize/repository"
)

type TokenUsecaseImpl struct {
	tokenRepository repository.TokenRepository
}

func NewTokenUsecase(tokenRepository repository.TokenRepository) *TokenUsecaseImpl {
	//func NewTokenUsecase(tokenRepository  repository.TokenRepository) *TokenUsecaseImpl {
	return &TokenUsecaseImpl{tokenRepository}
}

/////////////////////////////////////////////////////////////////////

func (TokenUsecase *TokenUsecaseImpl) SaveToken(token *model.Token) (*model.Token, error) {
	err := TokenUsecase.tokenRepository.Save(token)
	if err != nil {
		return nil, err
	}
	return token, nil
}

//update token by ObjectId
func (TokenUsecase *TokenUsecaseImpl) UpdateToken(id string, token *model.Token) (*model.Token, error) {
	err := TokenUsecase.tokenRepository.Update(id, token)
	if err != nil {
		return nil, err
	}
	return token, nil
}

//find token by id
func (TokenUsecase *TokenUsecaseImpl) GetByID(id string) (*model.Token, error) {
	token, err := TokenUsecase.tokenRepository.FindByID(id)
	//err:=TokenUsecase.tokenRepository.FindByID(id)
	if err != nil {
		return nil, err
	}
	return token, nil
}

func (TokenUsecase *TokenUsecaseImpl) GetAll() (model.Tokens, error) {
	tokens, err := TokenUsecase.tokenRepository.FindAll()
	if err != nil {
		return nil, err
	}
	return tokens, nil
}




func (TokenUsecase *TokenUsecaseImpl) DeleteByID(id string) error {
	err :=TokenUsecase.tokenRepository.Delete(id)
	if err != nil {
		return err
	}
	return nil
}