package dto

import (
	"time"
)

type Land_sale_work_bid struct {
	Jurisdiction          string     `json:"jurisdiction"`
	Land_sale_number      string     `json:"land_sale_number"`
	Land_sale_offering_id string     `json:"land_sale_offering_id"`
	Land_offering_bid_id  string     `json:"land_offering_bid_id"`
	Work_bid_id           string     `json:"work_bid_id"`
	Active_ind            *string    `json:"active_ind"`
	Critical_date         *time.Time `json:"critical_date"`
	Currency_conversion   *float64   `json:"currency_conversion"`
	Currency_ouom         *string    `json:"currency_ouom"`
	Description           *string    `json:"description"`
	Effective_date        *time.Time `json:"effective_date"`
	Expiry_date           *time.Time `json:"expiry_date"`
	Ppdm_guid             string     `json:"ppdm_guid"`
	Remark                *string    `json:"remark"`
	Source                *string    `json:"source"`
	Work_bid_count        *int       `json:"work_bid_count"`
	Work_bid_count_uom    *string    `json:"work_bid_count_uom"`
	Work_bid_type         *string    `json:"work_bid_type"`
	Work_bid_value        *float64   `json:"work_bid_value"`
	Row_changed_by        *string    `json:"row_changed_by"`
	Row_changed_date      *time.Time `json:"row_changed_date"`
	Row_created_by        *string    `json:"row_created_by"`
	Row_created_date      *time.Time `json:"row_created_date"`
	Row_effective_date    *time.Time `json:"row_effective_date"`
	Row_expiry_date       *time.Time `json:"row_expiry_date"`
	Row_quality           *string    `json:"row_quality"`
}
