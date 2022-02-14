package entity

import "time"

type TransactionRecord struct {
	TxID             uint64    `json:"TransactionID,omitempty"`
	TxType           string    `json:"transaction_type,omitempty"`
	TxDate           time.Time `json:"transaction_date,omitempty"`
	AvailableBalance float64   `json:"AvailableBalance,omitempty"`
	TxAmount         float64   `json:"amount,omitempty"`
	Description      string    `json:"Description,omitempty"`
	UserId           string    `json:"UserId,omitempty"`
}
