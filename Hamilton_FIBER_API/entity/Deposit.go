package entity

import "time"

type Deposit struct {
	DepositType   string    `json:"transaction_type,omitempty"`
	DepositDate   time.Time `json:"transaction_date,omitempty"`
	DepositAmount float64   `json:"amount,omitempty"`
	Description   string    `json:"Description,omitempty"`
	UserId        string    `json:"UserId,omitempty"`
}
