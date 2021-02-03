package usecase

import (
	"errors"

	"github.com/diogojorgebasso/imersao-fullstack-fullcycle/domain/model"
)


type PixUseCase struct {
	PixKeyRepository model.PixKeyRepositoryInterface
}

func(p*PixUseCase) RegisterKey(key string, kind string, accountId string) (*model.PixKey, error){
	account, err := p.PixKeyRepository.FindAccount(accountId)
	if err !=nil{
		return nil, err
	}
	pixKey, err := model.NewPixKey(kind, account,key)
	if err != nil {
		return nil, err
	}
	p.PixKeyRepository.RegisterKey(pixKey)
	if pixKey.ID =="" {
		return nil, errors.New("unable to create PixKey")
	}
	return pixKey, nil
}

//FindKey : in the pix repository
func (p *PixUseCase) FindKey(key string, kind string) (*model.PixKey, error) {
	pixKey, err := p.PixKeyRepository.FindKeyByKind(key, kind)
	if err != nil {
		return nil, err
	}
	return pixKey, nil
}