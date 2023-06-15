package dto

type Anl_step_xref struct {
	Analysis_id        string  `json:"analysis_id" default:"source"`
	Anl_source         string  `json:"anl_source" default:"source"`
	Step_seq_no1       int     `json:"step_seq_no_1" default:"1"`
	Step_seq_no2       int     `json:"step_seq_no_2" default:"1"`
	Active_ind         *string `json:"active_ind" default:""`
	Effective_date     *string `json:"effective_date" default:""`
	Expiry_date        *string `json:"expiry_date" default:""`
	Ppdm_guid          *string `json:"ppdm_guid" default:""`
	Remark             *string `json:"remark" default:""`
	Source             *string `json:"source" default:""`
	Step_xref_reason   *string `json:"step_xref_reason" default:""`
	Row_changed_by     *string `json:"row_changed_by" default:""`
	Row_changed_date   *string `json:"row_changed_date" default:""`
	Row_created_by     *string `json:"row_created_by" default:""`
	Row_created_date   *string `json:"row_created_date" default:""`
	Row_effective_date *string `json:"row_effective_date" default:""`
	Row_expiry_date    *string `json:"row_expiry_date" default:""`
	Row_quality        *string `json:"row_quality" default:""`
}
