package dto

type Seis_proc_plan_parm struct {
	Seis_proc_plan_id  string   `json:"seis_proc_plan_id" default:"source"`
	Proc_plan_step_id  string   `json:"proc_plan_step_id" default:"source"`
	Parameter_id       string   `json:"parameter_id" default:"source"`
	Active_ind         *string  `json:"active_ind" default:""`
	Date_format_desc   *string  `json:"date_format_desc" default:""`
	Effective_date     *string  `json:"effective_date" default:""`
	Expiry_date        *string  `json:"expiry_date" default:""`
	Parameter_desc     *string  `json:"parameter_desc" default:""`
	Parameter_origin   *string  `json:"parameter_origin" default:""`
	Parameter_type     *string  `json:"parameter_type" default:""`
	Parameter_uom      *string  `json:"parameter_uom" default:""`
	Parameter_value_1  *float64 `json:"parameter_value_1" default:""`
	Parameter_value_2  *float64 `json:"parameter_value_2" default:""`
	Ppdm_guid          *string  `json:"ppdm_guid" default:""`
	Remark             *string  `json:"remark" default:""`
	Source             *string  `json:"source" default:""`
	Row_changed_by     *string  `json:"row_changed_by" default:""`
	Row_changed_date   *string  `json:"row_changed_date" default:""`
	Row_created_by     *string  `json:"row_created_by" default:""`
	Row_created_date   *string  `json:"row_created_date" default:""`
	Row_effective_date *string  `json:"row_effective_date" default:""`
	Row_expiry_date    *string  `json:"row_expiry_date" default:""`
	Row_quality        *string  `json:"row_quality" default:""`
}
