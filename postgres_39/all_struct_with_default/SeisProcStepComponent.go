package dto

type Seis_proc_step_component struct {
	Seis_set_subtype   string  `json:"seis_set_subtype" default:"source"`
	Seis_proc_set_id   string  `json:"seis_proc_set_id" default:"source"`
	Component_id       string  `json:"component_id" default:"source"`
	Process_step_id    string  `json:"process_step_id" default:"source"`
	Active_ind         *string `json:"active_ind" default:""`
	Effective_date     *string `json:"effective_date" default:""`
	Expiry_date        *string `json:"expiry_date" default:""`
	Input_ind          *string `json:"input_ind" default:""`
	Output_ind         *string `json:"output_ind" default:""`
	Ppdm_guid          *string `json:"ppdm_guid" default:""`
	Remark             *string `json:"remark" default:""`
	Source             *string `json:"source" default:""`
	Row_changed_by     *string `json:"row_changed_by" default:""`
	Row_changed_date   *string `json:"row_changed_date" default:""`
	Row_created_by     *string `json:"row_created_by" default:""`
	Row_created_date   *string `json:"row_created_date" default:""`
	Row_effective_date *string `json:"row_effective_date" default:""`
	Row_expiry_date    *string `json:"row_expiry_date" default:""`
	Row_quality        *string `json:"row_quality" default:""`
}
