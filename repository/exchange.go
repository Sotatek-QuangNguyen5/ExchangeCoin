package repository

import (
	"servercoin/errs"
	"servercoin/model"

	"gorm.io/gorm"
)


type ExchangeRepository interface {

	GetExchange(address string) ([]*model.Exchange, *errs.AppError)
	DeleteExchange(message string) (*errs.AppError)
	GetMessage(message string) (bool, *errs.AppError)
	CreateExchange(newExchange *model.Exchange) (*errs.AppError)
}

type DefaultExchangeRepository struct {

	db *gorm.DB
}

func NewExchangeRepository(db *gorm.DB) ExchangeRepository {

	return DefaultExchangeRepository{

		db: db,
	}
}

func (e DefaultExchangeRepository) GetExchange(address string) ([]*model.Exchange, *errs.AppError) {

	var exchanges []*model.Exchange
	err := e.db.Where("address = ?", address).Find(&exchanges).Error
	if err != nil {

		return nil, errs.ErrorReadData()
	}
	return exchanges,  nil
}

func (e DefaultExchangeRepository) DeleteExchange(message string) (*errs.AppError) {

	err := e.db.Where("message = ?", message).Delete(&model.Exchange{}).Error
	if err != nil {

		return errs.ErrorDeleteData()
	}
	return nil
}

func (e DefaultExchangeRepository) GetMessage(message string) (bool, *errs.AppError) {

	var exchanges []model.Exchange
	err := e.db.Where("message = ?", message).Find(&exchanges).Error
	if err != nil {

		return true, errs.ErrorReadData()
	}
	return len(exchanges) != 0, nil
}

func (e DefaultExchangeRepository) CreateExchange(newExchange *model.Exchange) (*errs.AppError) {

	err := e.db.Create(&newExchange).Error
	if err != nil {

		return errs.ErrorInsertData()
	}
	return nil
}