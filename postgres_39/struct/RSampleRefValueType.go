package dto

import (
	"time"
)

type R_sample_ref_value_type struct {
	Reference_value_type string     `json:"reference_value_type"`
	Abbreviation         *string    `json:"abbreviation"`
	Active_ind           *string    `json:"active_ind"`
	Effective_date       *time.Time `json:"effective_date"`
	Expiry_date          *time.Time `json:"expiry_date"`
	Long_name            *string    `json:"long_name"`
	Ppdm_guid            string     `json:"ppdm_guid"`
	Property_set_id      *string    `json:"property_set_id"`
	Remark               *string    `json:"remark"`
	Short_name           *string    `json:"short_name"`
	Source               *string    `json:"source"`
	Row_changed_by       *string    `json:"row_changed_by"`
	Row_changed_date     *time.Time `json:"row_changed_date"`
	Row_created_by       *string    `json:"row_created_by"`
	Row_created_date     *time.Time `json:"row_created_date"`
	Row_effective_date   *time.Time `json:"row_effective_date"`
	Row_expiry_date      *time.Time `json:"row_expiry_date"`
	Row_quality          *string    `json:"row_quality"`
}
