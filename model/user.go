package model

import "github.com/kamva/mgm/v3"

type User struct {
	mgm.DefaultModel `bson:",inline"`
	Name             string  `json:"name" bson:"name"`
	Cpf              string  `json:"cpf" bson:"cpf"`
	Debit            float64 `json:"debit" bson:"debit"`
}
