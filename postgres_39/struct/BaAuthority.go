package dto

import (
	"time"
)

type Ba_authority struct {
	Business_associate_id string     `json:"business_associate_id"`
	Authority_id          string     `json:"authority_id"`
	Active_ind            *string    `json:"active_ind"`
	Authority_limit       *float64   `json:"authority_limit"`
	Authority_type        *string    `json:"authority_type"`
	Authorized_by         *string    `json:"authorized_by"`
	Currency_conversion   *float64   `json:"currency_conversion"`
	Currency_ouom         *string    `json:"currency_ouom"`
	Effective_date        *time.Time `json:"effective_date"`
	Expiry_date           *time.Time `json:"expiry_date"`
	Ppdm_guid             string     `json:"ppdm_guid"`
	Remark                *string    `json:"remark"`
	Represented_ba_id     *string    `json:"represented_ba_id"`
	Source                *string    `json:"source"`
	Row_changed_by        *string    `json:"row_changed_by"`
	Row_changed_date      *time.Time `json:"row_changed_date"`
	Row_created_by        *string    `json:"row_created_by"`
	Row_created_date      *time.Time `json:"row_created_date"`
	Row_effective_date    *time.Time `json:"row_effective_date"`
	Row_expiry_date       *time.Time `json:"row_expiry_date"`
	Row_quality           *string    `json:"row_quality"`
}
