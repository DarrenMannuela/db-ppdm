package dto

import (
	"time"
)

type Ppdm_sw_app_function struct {
	Sw_application_id    string     `json:"sw_application_id"`
	Sw_app_function      string     `json:"sw_app_function"`
	Function_obs_no      int        `json:"function_obs_no"`
	Abbreviation         *string    `json:"abbreviation"`
	Active_ind           *string    `json:"active_ind"`
	Effective_date       *time.Time `json:"effective_date"`
	Expiry_date          *time.Time `json:"expiry_date"`
	Long_name            *string    `json:"long_name"`
	Ppdm_guid            string     `json:"ppdm_guid"`
	Remark               *string    `json:"remark"`
	Short_name           *string    `json:"short_name"`
	Source               *string    `json:"source"`
	Sw_app_function_type *string    `json:"sw_app_function_type"`
	Row_changed_by       *string    `json:"row_changed_by"`
	Row_changed_date     *time.Time `json:"row_changed_date"`
	Row_created_by       *string    `json:"row_created_by"`
	Row_created_date     *time.Time `json:"row_created_date"`
	Row_effective_date   *time.Time `json:"row_effective_date"`
	Row_expiry_date      *time.Time `json:"row_expiry_date"`
	Row_quality          *string    `json:"row_quality"`
}
