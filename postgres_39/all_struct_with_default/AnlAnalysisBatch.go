package dto

type Anl_analysis_batch struct {
	Batch_id              string  `json:"batch_id" default:"source"`
	Active_ind            *string `json:"active_ind" default:""`
	Batch_desc            *string `json:"batch_desc" default:""`
	Batch_name            *string `json:"batch_name" default:""`
	Batch_ref_num         *string `json:"batch_ref_num" default:""`
	Create_date           *string `json:"create_date" default:""`
	Effective_date        *string `json:"effective_date" default:""`
	Expiry_date           *string `json:"expiry_date" default:""`
	Ppdm_guid             *string `json:"ppdm_guid" default:""`
	Remark                *string `json:"remark" default:""`
	Sample_count          *int    `json:"sample_count" default:""`
	Source                *string `json:"source" default:""`
	Standard_sample_count *int    `json:"standard_sample_count" default:""`
	Row_changed_by        *string `json:"row_changed_by" default:""`
	Row_changed_date      *string `json:"row_changed_date" default:""`
	Row_created_by        *string `json:"row_created_by" default:""`
	Row_created_date      *string `json:"row_created_date" default:""`
	Row_effective_date    *string `json:"row_effective_date" default:""`
	Row_expiry_date       *string `json:"row_expiry_date" default:""`
	Row_quality           *string `json:"row_quality" default:""`
}
