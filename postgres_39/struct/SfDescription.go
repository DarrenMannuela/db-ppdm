package dto

import (
	"time"
)

type Sf_description struct {
	Sf_subtype          string     `json:"sf_subtype"`
	Support_facility_id string     `json:"support_facility_id"`
	Description_obs_no  int        `json:"description_obs_no"`
	Active_ind          *string    `json:"active_ind"`
	Cost                *float64   `json:"cost"`
	Cost_ouom           *string    `json:"cost_ouom"`
	Cost_uom            *string    `json:"cost_uom"`
	Currency_conversion *float64   `json:"currency_conversion"`
	Date_format_desc    *time.Time `json:"date_format_desc"`
	Description         *string    `json:"description"`
	Desc_type           *string    `json:"desc_type"`
	Desc_value          *float64   `json:"desc_value"`
	Desc_value_code     *string    `json:"desc_value_code"`
	Desc_value_ouom     *string    `json:"desc_value_ouom"`
	Desc_value_uom      *string    `json:"desc_value_uom"`
	Effective_date      *time.Time `json:"effective_date"`
	Expiry_date         *time.Time `json:"expiry_date"`
	Max_value           *float64   `json:"max_value"`
	Max_value_ouom      *string    `json:"max_value_ouom"`
	Max_value_uom       *string    `json:"max_value_uom"`
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
