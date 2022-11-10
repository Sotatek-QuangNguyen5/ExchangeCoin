package service

import (

	"servercoin/dto"
	"servercoin/errs"
	"servercoin/repository"
)

type ExchangeService interface {

	GetExchange(address string) ([]*dto.Exchange, *errs.AppError)
	DeleteExchange(message string) (*errs.AppError)
	GetMessage(message string) (bool, *errs.AppError)
	CreateExchange(newExchange *dto.Exchange) (*errs.AppError)
}

type DefaultExchangeService struct {

	repo repository.ExchangeRepository
}

func NewExchangeService(repo repository.ExchangeRepository) ExchangeService {

	return DefaultExchangeService{

		repo: repo,
	}
}

func (e DefaultExchangeService) GetExchange(address string) ([]*dto.Exchange, *errs.AppError) {

	res, err := e.repo.GetExchange(address)
	if err != nil {

		return nil, err
	}
	return dto.MExchangesToDExchanges(res), nil
}

func (e DefaultExchangeService) DeleteExchange(message string) (*errs.AppError) {

	err := e.repo.DeleteExchange(message)
	return err
}

func (e DefaultExchangeService) GetMessage(message string) (bool, *errs.AppError) {

	res, err := e.repo.GetMessage(message)
	return res, err
}

func (e DefaultExchangeService) CreateExchange(newExchange *dto.Exchange) (*errs.AppError) {

	err := e.repo.CreateExchange(dto.DExchangeToMExchange(newExchange))
	return err
}