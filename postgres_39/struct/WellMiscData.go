package dto

import (
	"time"
)

type Well_misc_data struct {
	Uwi                  string     `json:"uwi"`
	Source               string     `json:"source"`
	Misc_data_type       string     `json:"misc_data_type"`
	Misc_data_obs_no     int        `json:"misc_data_obs_no"`
	Active_ind           *string    `json:"active_ind"`
	Cost                 *float64   `json:"cost"`
	Currency_conversion  *float64   `json:"currency_conversion"`
	Currency_ouom        *string    `json:"currency_ouom"`
	Currency_uom         *string    `json:"currency_uom"`
	Data_value           *float64   `json:"data_value"`
	Data_value_ouom      *string    `json:"data_value_ouom"`
	Data_value_uom       *string    `json:"data_value_uom"`
	Date_format_desc     *time.Time `json:"date_format_desc"`
	Effective_date       *time.Time `json:"effective_date"`
	Expiry_date          *time.Time `json:"expiry_date"`
	Max_date             *time.Time `json:"max_date"`
	Max_value            *float64   `json:"max_value"`
	Max_value_ouom       *string    `json:"max_value_ouom"`
	Max_value_uom        *string    `json:"max_value_uom"`
	Min_date             *time.Time `json:"min_date"`
	Min_value            *float64   `json:"min_value"`
	Min_value_ouom       *string    `json:"min_value_ouom"`
	Min_value_uom        *string    `json:"min_value_uom"`
	Misc_data_code       *string    `json:"misc_data_code"`
	Misc_data_desc       *string    `json:"misc_data_desc"`
	Ppdm_guid            string     `json:"ppdm_guid"`
	Reference_value      *float64   `json:"reference_value"`
	Reference_value_ouom *string    `json:"reference_value_ouom"`
	Reference_value_type *string    `json:"reference_value_type"`
	Reference_value_uom  *string    `json:"reference_value_uom"`
	Remark               *string    `json:"remark"`
	Row_changed_by       *string    `json:"row_changed_by"`
	Row_changed_date     *time.Time `json:"row_changed_date"`
	Row_created_by       *string    `json:"row_created_by"`
	Row_created_date     *time.Time `json:"row_created_date"`
	Row_effective_date   *time.Time `json:"row_effective_date"`
	Row_expiry_date      *time.Time `json:"row_expiry_date"`
	Row_quality          *string    `json:"row_quality"`
}
