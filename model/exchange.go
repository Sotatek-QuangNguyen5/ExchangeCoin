package model

import "gorm.io/gorm"

type Exchange struct {

	gorm.Model
	Signature	string
	Address		string
	Message		string
	Amount		int64
	Nonce		int64
}