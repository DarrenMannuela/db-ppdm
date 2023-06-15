package dto

import (
	"time"
)

type Seis_transaction struct {
	Seis_transaction_id string     `json:"seis_transaction_id"`
	Transaction_type    string     `json:"transaction_type"`
	Active_ind          *string    `json:"active_ind"`
	Currency_conversion *float64   `json:"currency_conversion"`
	Currency_ouom       *string    `json:"currency_ouom"`
	Effective_date      *time.Time `json:"effective_date"`
	Expiry_date         *time.Time `json:"expiry_date"`
	Length              *float64   `json:"length"`
	Length_ouom         *string    `json:"length_ouom"`
	Ppdm_guid           string     `json:"ppdm_guid"`
	Price_per_length    *float64   `json:"price_per_length"`
	Reference_num       *string    `json:"reference_num"`
	Remark              *string    `json:"remark"`
	Source              *string    `json:"source"`
	Total_price         *float64   `json:"total_price"`
	Transaction_date    *time.Time `json:"transaction_date"`
	Transaction_status  *string    `json:"transaction_status"`
	Row_changed_by      *string    `json:"row_changed_by"`
	Row_changed_date    *time.Time `json:"row_changed_date"`
	Row_created_by      *string    `json:"row_created_by"`
	Row_created_date    *time.Time `json:"row_created_date"`
	Row_effective_date  *time.Time `json:"row_effective_date"`
	Row_expiry_date     *time.Time `json:"row_expiry_date"`
	Row_quality         *string    `json:"row_quality"`
}
