package dto

type Ppdm_metric struct {
	Metric_id             string  `json:"metric_id" default:"source"`
	Active_ind            *string `json:"active_ind" default:""`
	Effective_date        *string `json:"effective_date" default:""`
	End_date              *string `json:"end_date" default:""`
	Expiry_date           *string `json:"expiry_date" default:""`
	Metric_name           *string `json:"metric_name" default:""`
	Metric_procedure      *string `json:"metric_procedure" default:""`
	Metric_type           *string `json:"metric_type" default:""`
	Organization_id       *string `json:"organization_id" default:""`
	Organization_seq_no   *int    `json:"organization_seq_no" default:""`
	Owner_ba_id           *string `json:"owner_ba_id" default:""`
	Ppdm_guid             *string `json:"ppdm_guid" default:""`
	Procedure_id          *string `json:"procedure_id" default:""`
	Procedure_system_id   *string `json:"procedure_system_id" default:""`
	Projected_final_count *int    `json:"projected_final_count" default:""`
	Purpose_desc          *string `json:"purpose_desc" default:""`
	Remark                *string `json:"remark" default:""`
	Source                *string `json:"source" default:""`
	Start_date            *string `json:"start_date" default:""`
	Row_changed_by        *string `json:"row_changed_by" default:""`
	Row_changed_date      *string `json:"row_changed_date" default:""`
	Row_created_by        *string `json:"row_created_by" default:""`
	Row_created_date      *string `json:"row_created_date" default:""`
	Row_effective_date    *string `json:"row_effective_date" default:""`
	Row_expiry_date       *string `json:"row_expiry_date" default:""`
	Row_quality           *string `json:"row_quality" default:""`
}
