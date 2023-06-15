package dto

import (
	"time"
)

type Anl_valid_measure struct {
	Method_set_id          string     `json:"method_set_id"`
	Method_id              string     `json:"method_id"`
	Measurement_type       string     `json:"measurement_type"`
	Valid_value_obs_no     int        `json:"valid_value_obs_no"`
	Accuracy_type          *string    `json:"accuracy_type"`
	Active_ind             *string    `json:"active_ind"`
	Average_ratio_value    *float64   `json:"average_ratio_value"`
	Average_value          *float64   `json:"average_value"`
	Average_value_ouom     *string    `json:"average_value_ouom"`
	Average_value_uom      *string    `json:"average_value_uom"`
	Calculate_method_id    *string    `json:"calculate_method_id"`
	Effective_date         *time.Time `json:"effective_date"`
	Expiry_date            *time.Time `json:"expiry_date"`
	Maximum_ratio_value    *float64   `json:"maximum_ratio_value"`
	Max_value              *float64   `json:"max_value"`
	Max_value_ouom         *string    `json:"max_value_ouom"`
	Max_value_uom          *string    `json:"max_value_uom"`
	Minimum_ratio_value    *float64   `json:"minimum_ratio_value"`
	Min_value              *float64   `json:"min_value"`
	Min_value_ouom         *string    `json:"min_value_ouom"`
	Min_value_uom          *string    `json:"min_value_uom"`
	Missing_representation *string    `json:"missing_representation"`
	Null_representation    *string    `json:"null_representation"`
	Ppdm_guid              string     `json:"ppdm_guid"`
	Reference_value        *float64   `json:"reference_value"`
	Reference_value_ouom   *string    `json:"reference_value_ouom"`
	Reference_value_uom    *string    `json:"reference_value_uom"`
	Remark                 *string    `json:"remark"`
	Source                 *string    `json:"source"`
	Substance_id           *string    `json:"substance_id"`
	Valid_desc             *string    `json:"valid_desc"`
	Valid_value_text       *string    `json:"valid_value_text"`
	Row_changed_by         *string    `json:"row_changed_by"`
	Row_changed_date       *time.Time `json:"row_changed_date"`
	Row_created_by         *string    `json:"row_created_by"`
	Row_created_date       *time.Time `json:"row_created_date"`
	Row_effective_date     *time.Time `json:"row_effective_date"`
	Row_expiry_date        *time.Time `json:"row_expiry_date"`
	Row_quality            *string    `json:"row_quality"`
}
