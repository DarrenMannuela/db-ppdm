package dto

type Anl_parm struct {
	Analysis_id          string   `json:"analysis_id" default:"source"`
	Anl_source           string   `json:"anl_source" default:"source"`
	Parm_obs_no          int      `json:"parm_obs_no" default:"1"`
	Active_ind           *string  `json:"active_ind" default:""`
	Anl_value_remark     *string  `json:"anl_value_remark" default:""`
	Average_value        *float64 `json:"average_value" default:""`
	Average_value_ouom   *string  `json:"average_value_ouom" default:""`
	Average_value_uom    *string  `json:"average_value_uom" default:""`
	Cat_equip_id         *string  `json:"cat_equip_id" default:""`
	Effective_date       *string  `json:"effective_date" default:""`
	Equip_id             *string  `json:"equip_id" default:""`
	Equip_obs_no         *int     `json:"equip_obs_no" default:""`
	Expiry_date          *string  `json:"expiry_date" default:""`
	Max_date             *string  `json:"max_date" default:""`
	Max_value            *float64 `json:"max_value" default:""`
	Max_value_ouom       *string  `json:"max_value_ouom" default:""`
	Max_value_uom        *string  `json:"max_value_uom" default:""`
	Method_id            *string  `json:"method_id" default:""`
	Method_parm_obs_no   *int     `json:"method_parm_obs_no" default:""`
	Method_set_id        *string  `json:"method_set_id" default:""`
	Min_date             *string  `json:"min_date" default:""`
	Min_value            *float64 `json:"min_value" default:""`
	Min_value_ouom       *string  `json:"min_value_ouom" default:""`
	Min_value_uom        *string  `json:"min_value_uom" default:""`
	Parameter_ba_id      *string  `json:"parameter_ba_id" default:""`
	Parameter_type       *string  `json:"parameter_type" default:""`
	Ppdm_guid            *string  `json:"ppdm_guid" default:""`
	Reference_value      *float64 `json:"reference_value" default:""`
	Reference_value_ouom *string  `json:"reference_value_ouom" default:""`
	Reference_value_type *string  `json:"reference_value_type" default:""`
	Reference_value_uom  *string  `json:"reference_value_uom" default:""`
	Remark               *string  `json:"remark" default:""`
	Source               *string  `json:"source" default:""`
	Step_seq_no          *int     `json:"step_seq_no" default:""`
	Substance_id         *string  `json:"substance_id" default:""`
	Row_changed_by       *string  `json:"row_changed_by" default:""`
	Row_changed_date     *string  `json:"row_changed_date" default:""`
	Row_created_by       *string  `json:"row_created_by" default:""`
	Row_created_date     *string  `json:"row_created_date" default:""`
	Row_effective_date   *string  `json:"row_effective_date" default:""`
	Row_expiry_date      *string  `json:"row_expiry_date" default:""`
	Row_quality          *string  `json:"row_quality" default:""`
}
