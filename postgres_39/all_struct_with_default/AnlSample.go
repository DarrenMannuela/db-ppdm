package dto

type Anl_sample struct {
	Analysis_id         string  `json:"analysis_id" default:"source"`
	Anl_source          string  `json:"anl_source" default:"source"`
	Sample_id           string  `json:"sample_id" default:"source"`
	Active_ind          *string `json:"active_ind" default:""`
	Batch_id            *string `json:"batch_id" default:""`
	Created_by_step_ind *string `json:"created_by_step_ind" default:""`
	Effective_date      *string `json:"effective_date" default:""`
	End_date            *string `json:"end_date" default:""`
	Expiry_date         *string `json:"expiry_date" default:""`
	Ppdm_guid           *string `json:"ppdm_guid" default:""`
	Remark              *string `json:"remark" default:""`
	Sample_description  *string `json:"sample_description" default:""`
	Source              *string `json:"source" default:""`
	Standard_sample_ind *string `json:"standard_sample_ind" default:""`
	Step_seq_no         *int    `json:"step_seq_no" default:""`
	Used_by_step_ind    *string `json:"used_by_step_ind" default:""`
	Row_changed_by      *string `json:"row_changed_by" default:""`
	Row_changed_date    *string `json:"row_changed_date" default:""`
	Row_created_by      *string `json:"row_created_by" default:""`
	Row_created_date    *string `json:"row_created_date" default:""`
	Row_effective_date  *string `json:"row_effective_date" default:""`
	Row_expiry_date     *string `json:"row_expiry_date" default:""`
	Row_quality         *string `json:"row_quality" default:""`
}
