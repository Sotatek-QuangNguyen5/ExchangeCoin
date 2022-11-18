package service

import (
	"fmt"
	"log"
	"servercoin/config"
	"servercoin/dto"
	"servercoin/errs"
	"servercoin/repository"
)

type ExchangeService interface {

	GetExchange(address string) ([]*dto.Exchange, *errs.AppError)
	DeleteExchange(message string) (*errs.AppError)
	GetMessage(message string) (bool, *errs.AppError)
	CreateExchange(newExchange *dto.Exchange) (*errs.AppError)
	UpdateUseSignature(newExchange *dto.Exchange) (*errs.AppError)
	SaveEventReceiveMoney(newExchange *dto.Exchange)
	SaveEventWithdrawMoney(newExchange *dto.Exchange)
}

var (

	Serv ExchangeService
)

func Init() {

	for {

		if config.DB != nil {

			break
		}
	}
	Serv = NewExchangeService(repository.NewExchangeRepository(config.DB))
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

func (e DefaultExchangeService) UpdateUseSignature(newExchange *dto.Exchange) (*errs.AppError) {

	err := e.repo.UpdateUseSignature(dto.DExchangeToMExchange(newExchange))
	return err
}

func (e DefaultExchangeService) SaveEventReceiveMoney(newExchange *dto.Exchange) {

    err := e.CreateExchange(newExchange)
    if err != nil {

        fmt.Println()
        log.Println(err.Message)
        return
    }
    fmt.Println()
    fmt.Printf("Saved Transaction Receive Money Address %v to Database.\n", newExchange.Address)
    fmt.Println()
}

func (e DefaultExchangeService) SaveEventWithdrawMoney(newExchange *dto.Exchange) {

    err := e.UpdateUseSignature(newExchange)
    if err != nil {

        fmt.Println()
        log.Println(err.Message)
        return
    }
    fmt.Println()
    fmt.Printf("Saved Transaction Withdraw Money of Address %v to Database.\n", newExchange.Address)
    fmt.Println()
}