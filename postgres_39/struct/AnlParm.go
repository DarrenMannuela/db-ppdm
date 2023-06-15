package dto

import (
	"time"
)

type Anl_parm struct {
	Analysis_id          string     `json:"analysis_id"`
	Anl_source           string     `json:"anl_source"`
	Parm_obs_no          int        `json:"parm_obs_no"`
	Active_ind           *string    `json:"active_ind"`
	Anl_value_remark     *string    `json:"anl_value_remark"`
	Average_value        *float64   `json:"average_value"`
	Average_value_ouom   *string    `json:"average_value_ouom"`
	Average_value_uom    *string    `json:"average_value_uom"`
	Cat_equip_id         *string    `json:"cat_equip_id"`
	Effective_date       *time.Time `json:"effective_date"`
	Equip_id             *string    `json:"equip_id"`
	Equip_obs_no         *int       `json:"equip_obs_no"`
	Expiry_date          *time.Time `json:"expiry_date"`
	Max_date             *time.Time `json:"max_date"`
	Max_value            *float64   `json:"max_value"`
	Max_value_ouom       *string    `json:"max_value_ouom"`
	Max_value_uom        *string    `json:"max_value_uom"`
	Method_id            *string    `json:"method_id"`
	Method_parm_obs_no   *int       `json:"method_parm_obs_no"`
	Method_set_id        *string    `json:"method_set_id"`
	Min_date             *time.Time `json:"min_date"`
	Min_value            *float64   `json:"min_value"`
	Min_value_ouom       *string    `json:"min_value_ouom"`
	Min_value_uom        *string    `json:"min_value_uom"`
	Parameter_ba_id      *string    `json:"parameter_ba_id"`
	Parameter_type       *string    `json:"parameter_type"`
	Ppdm_guid            string     `json:"ppdm_guid"`
	Reference_value      *float64   `json:"reference_value"`
	Reference_value_ouom *string    `json:"reference_value_ouom"`
	Reference_value_type *string    `json:"reference_value_type"`
	Reference_value_uom  *string    `json:"reference_value_uom"`
	Remark               *string    `json:"remark"`
	Source               *string    `json:"source"`
	Step_seq_no          *int       `json:"step_seq_no"`
	Substance_id         *string    `json:"substance_id"`
	Row_changed_by       *string    `json:"row_changed_by"`
	Row_changed_date     *time.Time `json:"row_changed_date"`
	Row_created_by       *string    `json:"row_created_by"`
	Row_created_date     *time.Time `json:"row_created_date"`
	Row_effective_date   *time.Time `json:"row_effective_date"`
	Row_expiry_date      *time.Time `json:"row_expiry_date"`
	Row_quality          *string    `json:"row_quality"`
}
