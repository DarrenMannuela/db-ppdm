package dto

type Project_step_xref struct {
	Project_id         string  `json:"project_id" default:"source"`
	Step_id1           string  `json:"step_id_1" default:"source"`
	Step_id2           string  `json:"step_id_2" default:"source"`
	Xref_obs_no        int     `json:"xref_obs_no" default:"1"`
	Active_ind         *string `json:"active_ind" default:""`
	Description        *string `json:"description" default:""`
	Effective_date     *string `json:"effective_date" default:""`
	Expiry_date        *string `json:"expiry_date" default:""`
	Ppdm_guid          *string `json:"ppdm_guid" default:""`
	Remark             *string `json:"remark" default:""`
	Source             *string `json:"source" default:""`
	Step_seq_no        *int    `json:"step_seq_no" default:""`
	Step_xref_type     *string `json:"step_xref_type" default:""`
	Use_rule           *string `json:"use_rule" default:""`
	Row_changed_by     *string `json:"row_changed_by" default:""`
	Row_changed_date   *string `json:"row_changed_date" default:""`
	Row_created_by     *string `json:"row_created_by" default:""`
	Row_created_date   *string `json:"row_created_date" default:""`
	Row_effective_date *string `json:"row_effective_date" default:""`
	Row_expiry_date    *string `json:"row_expiry_date" default:""`
	Row_quality        *string `json:"row_quality" default:""`
}
