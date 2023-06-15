package dto

import (
	"time"
)

type Area_description struct {
	Area_id               string     `json:"area_id"`
	Area_type             string     `json:"area_type"`
	Description_obs_no    int        `json:"description_obs_no"`
	Active_ind            *string    `json:"active_ind"`
	Area_description_date *time.Time `json:"area_description_date"`
	Area_desc_code        *string    `json:"area_desc_code"`
	Area_desc_type        *string    `json:"area_desc_type"`
	Average_value         *float64   `json:"average_value"`
	Average_value_ouom    *string    `json:"average_value_ouom"`
	Average_value_uom     *string    `json:"average_value_uom"`
	Date_format_desc      *time.Time `json:"date_format_desc"`
	Description           *string    `json:"description"`
	Effective_date        *time.Time `json:"effective_date"`
	Expiry_date           *time.Time `json:"expiry_date"`
	Max_value             *float64   `json:"max_value"`
	Max_value_ouom        *string    `json:"max_value_ouom"`
	Max_value_uom         *string    `json:"max_value_uom"`
	Min_value             *float64   `json:"min_value"`
	Min_value_ouom        *string    `json:"min_value_ouom"`
	Min_value_uom         *string    `json:"min_value_uom"`
	Ppdm_guid             string     `json:"ppdm_guid"`
	Remark                *string    `json:"remark"`
	Source                *string    `json:"source"`
	Source_document_id    *string    `json:"source_document_id"`
	Row_changed_by        *string    `json:"row_changed_by"`
	Row_changed_date      *time.Time `json:"row_changed_date"`
	Row_created_by        *string    `json:"row_created_by"`
	Row_created_date      *time.Time `json:"row_created_date"`
	Row_effective_date    *time.Time `json:"row_effective_date"`
	Row_expiry_date       *time.Time `json:"row_expiry_date"`
	Row_quality           *string    `json:"row_quality"`
}
