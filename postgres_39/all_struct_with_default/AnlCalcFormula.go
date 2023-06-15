package dto

type Anl_calc_formula struct {
	Calculate_method_id string  `json:"calculate_method_id" default:"source"`
	Formula_part_obs_no int     `json:"formula_part_obs_no" default:"1"`
	Active_ind          *string `json:"active_ind" default:""`
	Assigned_variable   *string `json:"assigned_variable" default:""`
	Column_name         *string `json:"column_name" default:""`
	Effective_date      *string `json:"effective_date" default:""`
	Expiry_date         *string `json:"expiry_date" default:""`
	Ppdm_guid           *string `json:"ppdm_guid" default:""`
	Remark              *string `json:"remark" default:""`
	Source              *string `json:"source" default:""`
	Substance_id        *string `json:"substance_id" default:""`
	System_id           *string `json:"system_id" default:""`
	Table_name          *string `json:"table_name" default:""`
	Row_changed_by      *string `json:"row_changed_by" default:""`
	Row_changed_date    *string `json:"row_changed_date" default:""`
	Row_created_by      *string `json:"row_created_by" default:""`
	Row_created_date    *string `json:"row_created_date" default:""`
	Row_effective_date  *string `json:"row_effective_date" default:""`
	Row_expiry_date     *string `json:"row_expiry_date" default:""`
	Row_quality         *string `json:"row_quality" default:""`
}
