package dto

import (
	"time"
)

type Seis_set_authorize struct {
	Seis_set_subtype    string     `json:"seis_set_subtype"`
	Seis_set_id         string     `json:"seis_set_id"`
	Authorize_id        string     `json:"authorize_id"`
	Active_ind          *string    `json:"active_ind"`
	Authorized_by       *string    `json:"authorized_by"`
	Authorize_type      *string    `json:"authorize_type"`
	Currency_conversion *float64   `json:"currency_conversion"`
	Currency_ouom       *string    `json:"currency_ouom"`
	Effective_date      *time.Time `json:"effective_date"`
	Expiry_date         *time.Time `json:"expiry_date"`
	Limit_desc          *string    `json:"limit_desc"`
	Limit_type          *string    `json:"limit_type"`
	Partner_ba_id       *string    `json:"partner_ba_id"`
	Ppdm_guid           string     `json:"ppdm_guid"`
	Price_per_length    *float64   `json:"price_per_length"`
	Reason_type         *string    `json:"reason_type"`
	Remark              *string    `json:"remark"`
	Source              *string    `json:"source"`
	Row_changed_by      *string    `json:"row_changed_by"`
	Row_changed_date    *time.Time `json:"row_changed_date"`
	Row_created_by      *string    `json:"row_created_by"`
	Row_created_date    *time.Time `json:"row_created_date"`
	Row_effective_date  *time.Time `json:"row_effective_date"`
	Row_expiry_date     *time.Time `json:"row_expiry_date"`
	Row_quality         *string    `json:"row_quality"`
}
