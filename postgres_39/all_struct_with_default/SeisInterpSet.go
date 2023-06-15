package dto

type Seis_interp_set struct {
	Seis_set_subtype   string  `json:"seis_set_subtype" default:"source"`
	Interp_set_id      string  `json:"interp_set_id" default:"source"`
	Active_ind         *string `json:"active_ind" default:""`
	Effective_date     *string `json:"effective_date" default:""`
	Expiry_date        *string `json:"expiry_date" default:""`
	Interpreter        *string `json:"interpreter" default:""`
	Interp_date        *string `json:"interp_date" default:""`
	Interp_objective   *string `json:"interp_objective" default:""`
	Interp_set_name    *string `json:"interp_set_name" default:""`
	Interp_type        *string `json:"interp_type" default:""`
	Pick_method        *string `json:"pick_method" default:""`
	Ppdm_guid          *string `json:"ppdm_guid" default:""`
	Remark             *string `json:"remark" default:""`
	Source             *string `json:"source" default:""`
	Sw_application_id  *string `json:"sw_application_id" default:""`
	Trace_position     *string `json:"trace_position" default:""`
	Row_changed_by     *string `json:"row_changed_by" default:""`
	Row_changed_date   *string `json:"row_changed_date" default:""`
	Row_created_by     *string `json:"row_created_by" default:""`
	Row_created_date   *string `json:"row_created_date" default:""`
	Row_effective_date *string `json:"row_effective_date" default:""`
	Row_expiry_date    *string `json:"row_expiry_date" default:""`
	Row_quality        *string `json:"row_quality" default:""`
}
