package dto

import (
	"time"
)

type Applic_desc struct {
	Application_id      string     `json:"application_id"`
	Description_id      string     `json:"description_id"`
	Active_ind          *string    `json:"active_ind"`
	Application_type    *string    `json:"application_type"`
	Applic_desc_type    *string    `json:"applic_desc_type"`
	Contact_ba_id       *string    `json:"contact_ba_id"`
	Currency_conversion *float64   `json:"currency_conversion"`
	Currency_ouom       *string    `json:"currency_ouom"`
	Currency_uom        *string    `json:"currency_uom"`
	Date_format_desc    *time.Time `json:"date_format_desc"`
	Description         *string    `json:"description"`
	Desc_date           *time.Time `json:"desc_date"`
	Desc_value          *float64   `json:"desc_value"`
	Desc_value_ouom     *string    `json:"desc_value_ouom"`
	Desc_value_uom      *string    `json:"desc_value_uom"`
	Effective_date      *time.Time `json:"effective_date"`
	Expiry_date         *time.Time `json:"expiry_date"`
	Max_cost            *float64   `json:"max_cost"`
	Max_date            *time.Time `json:"max_date"`
	Max_value           *float64   `json:"max_value"`
	Max_value_ouom      *string    `json:"max_value_ouom"`
	Max_value_uom       *string    `json:"max_value_uom"`
	Min_cost            *float64   `json:"min_cost"`
	Min_date            *time.Time `json:"min_date"`
	Min_value           *float64   `json:"min_value"`
	Min_value_ouom      *string    `json:"min_value_ouom"`
	Min_value_uom       *string    `json:"min_value_uom"`
	Ppdm_guid           string     `json:"ppdm_guid"`
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
