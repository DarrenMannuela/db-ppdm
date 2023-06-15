package dto

type Seis_interp_load struct {
	Seis_set_subtype   string  `json:"seis_set_subtype" default:"source"`
	Interp_set_id      string  `json:"interp_set_id" default:"source"`
	Process_step_id    string  `json:"process_step_id" default:"source"`
	Active_ind         *string `json:"active_ind" default:""`
	Description        *string `json:"description" default:""`
	Effective_date     *string `json:"effective_date" default:""`
	End_date           *string `json:"end_date" default:""`
	Expiry_date        *string `json:"expiry_date" default:""`
	Interp_origin_id   *string `json:"interp_origin_id" default:""`
	Ppdm_guid          *string `json:"ppdm_guid" default:""`
	Process_status     *string `json:"process_status" default:""`
	Remark             *string `json:"remark" default:""`
	Source             *string `json:"source" default:""`
	Start_date         *string `json:"start_date" default:""`
	Step_name          *string `json:"step_name" default:""`
	Step_seq_no        *int    `json:"step_seq_no" default:""`
	Row_changed_by     *string `json:"row_changed_by" default:""`
	Row_changed_date   *string `json:"row_changed_date" default:""`
	Row_created_by     *string `json:"row_created_by" default:""`
	Row_created_date   *string `json:"row_created_date" default:""`
	Row_effective_date *string `json:"row_effective_date" default:""`
	Row_expiry_date    *string `json:"row_expiry_date" default:""`
	Row_quality        *string `json:"row_quality" default:""`
}
