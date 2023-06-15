package dto

import (
	"time"
)

type Sample_desc_other struct {
	Sample_id            string     `json:"sample_id"`
	Description_id       string     `json:"description_id"`
	Active_ind           *string    `json:"active_ind"`
	Average_ratio_value  *float64   `json:"average_ratio_value"`
	Average_value        *float64   `json:"average_value"`
	Average_value_ouom   *string    `json:"average_value_ouom"`
	Average_value_uom    *string    `json:"average_value_uom"`
	Calculate_method_id  *string    `json:"calculate_method_id"`
	Cost                 *float64   `json:"cost"`
	Currency_conversion  *float64   `json:"currency_conversion"`
	Currency_ouom        *string    `json:"currency_ouom"`
	Currency_uom         *string    `json:"currency_uom"`
	Date_format_desc     *time.Time `json:"date_format_desc"`
	Description          *string    `json:"description"`
	Description_code     *string    `json:"description_code"`
	Description_type     *string    `json:"description_type"`
	Effective_date       *time.Time `json:"effective_date"`
	Expiry_date          *time.Time `json:"expiry_date"`
	Max_date             *time.Time `json:"max_date"`
	Max_ratio_value      *float64   `json:"max_ratio_value"`
	Max_value            *float64   `json:"max_value"`
	Max_value_ouom       *string    `json:"max_value_ouom"`
	Max_value_uom        *string    `json:"max_value_uom"`
	Min_date             *time.Time `json:"min_date"`
	Min_ratio_value      *float64   `json:"min_ratio_value"`
	Min_value            *float64   `json:"min_value"`
	Min_value_ouom       *string    `json:"min_value_ouom"`
	Min_value_uom        *string    `json:"min_value_uom"`
	Ppdm_guid            string     `json:"ppdm_guid"`
	Reference_value      *float64   `json:"reference_value"`
	Reference_value_ouom *string    `json:"reference_value_ouom"`
	Reference_value_type *string    `json:"reference_value_type"`
	Reference_value_uom  *string    `json:"reference_value_uom"`
	Remark               *string    `json:"remark"`
	Reported_ratio       *string    `json:"reported_ratio"`
	Reported_ratio_ouom  *string    `json:"reported_ratio_ouom"`
	Source               *string    `json:"source"`
	Row_changed_by       *string    `json:"row_changed_by"`
	Row_changed_date     *time.Time `json:"row_changed_date"`
	Row_created_by       *string    `json:"row_created_by"`
	Row_created_date     *time.Time `json:"row_created_date"`
	Row_effective_date   *time.Time `json:"row_effective_date"`
	Row_expiry_date      *time.Time `json:"row_expiry_date"`
	Row_quality          *string    `json:"row_quality"`
}
