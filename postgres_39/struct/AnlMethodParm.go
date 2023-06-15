package dto

import (
	"time"
)

type Anl_method_parm struct {
	Method_set_id        string     `json:"method_set_id"`
	Method_id            string     `json:"method_id"`
	Parameter_type       string     `json:"parameter_type"`
	Method_parm_obs_no   int        `json:"method_parm_obs_no"`
	Abbreviation         *string    `json:"abbreviation"`
	Active_ind           *string    `json:"active_ind"`
	Anl_value_remark     *string    `json:"anl_value_remark"`
	Average_value        *float64   `json:"average_value"`
	Average_value_ouom   *string    `json:"average_value_ouom"`
	Average_value_uom    *string    `json:"average_value_uom"`
	Effective_date       *time.Time `json:"effective_date"`
	Expiry_date          *time.Time `json:"expiry_date"`
	Long_name            *string    `json:"long_name"`
	Max_date             *time.Time `json:"max_date"`
	Max_value            *float64   `json:"max_value"`
	Max_value_ouom       *string    `json:"max_value_ouom"`
	Max_value_uom        *string    `json:"max_value_uom"`
	Measurement_type     *string    `json:"measurement_type"`
	Min_date             *time.Time `json:"min_date"`
	Min_value            *float64   `json:"min_value"`
	Min_value_ouom       *string    `json:"min_value_ouom"`
	Min_value_uom        *string    `json:"min_value_uom"`
	Ppdm_guid            string     `json:"ppdm_guid"`
	Reference_value      *float64   `json:"reference_value"`
	Reference_value_ouom *string    `json:"reference_value_ouom"`
	Reference_value_type *string    `json:"reference_value_type"`
	Reference_value_uom  *string    `json:"reference_value_uom"`
	Remark               *string    `json:"remark"`
	Short_name           *string    `json:"short_name"`
	Source               *string    `json:"source"`
	Substance_id         *string    `json:"substance_id"`
	Row_changed_by       *string    `json:"row_changed_by"`
	Row_changed_date     *time.Time `json:"row_changed_date"`
	Row_created_by       *string    `json:"row_created_by"`
	Row_created_date     *time.Time `json:"row_created_date"`
	Row_effective_date   *time.Time `json:"row_effective_date"`
	Row_expiry_date      *time.Time `json:"row_expiry_date"`
	Row_quality          *string    `json:"row_quality"`
}
