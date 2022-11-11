package dto

import "servercoin/model"

type Exchange struct {

	Message		string		`json:"message"`
	Address		string		`json:"address"`
	Signature	string		`json:"signature"`
	Amount		string		`json:"amount"`
}

func MExchangeToDExchange(exchange *model.Exchange) *Exchange {

	return &Exchange{

		Message: exchange.Message,
		Address: exchange.Address,
		Signature: exchange.Signature,
		Amount: exchange.Amount,
	}
}

func MExchangesToDExchanges(exchanges []*model.Exchange) []*Exchange {

	var Dexchanges []*Exchange
	for _, v := range exchanges {

		Dexchanges = append(Dexchanges, MExchangeToDExchange(v))
	}
	return Dexchanges
}

func DExchangeToMExchange(exchange *Exchange) *model.Exchange {

	return &model.Exchange{

		Message: exchange.Message,
		Address: exchange.Address,
		Signature: exchange.Signature,
		Amount: exchange.Amount,
	}
}

func DExchangesToMExchanges(exchanges []*Exchange) []*model.Exchange {

	var Mexchanges []*model.Exchange
	for _, v := range exchanges {

		Mexchanges = append(Mexchanges, DExchangeToMExchange(v))
	}
	return Mexchanges
}