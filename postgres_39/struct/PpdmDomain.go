package dto

import (
	"time"
)

type Ppdm_domain struct {
	Domain               string     `json:"domain"`
	Active_ind           *string    `json:"active_ind"`
	Data_type            *string    `json:"data_type"`
	Domain_name          *string    `json:"domain_name"`
	Domain_quantity_type *string    `json:"domain_quantity_type"`
	Domain_scale         *float64   `json:"domain_scale"`
	Domain_width         *float64   `json:"domain_width"`
	Effective_date       *time.Time `json:"effective_date"`
	Expiry_date          *time.Time `json:"expiry_date"`
	Ppdm_guid            string     `json:"ppdm_guid"`
	Remark               *string    `json:"remark"`
	Source               *string    `json:"source"`
	Row_changed_by       *string    `json:"row_changed_by"`
	Row_changed_date     *time.Time `json:"row_changed_date"`
	Row_created_by       *string    `json:"row_created_by"`
	Row_created_date     *time.Time `json:"row_created_date"`
	Row_effective_date   *time.Time `json:"row_effective_date"`
	Row_expiry_date      *time.Time `json:"row_expiry_date"`
	Row_quality          *string    `json:"row_quality"`
}
