package dto

type Ppdm_map_load_error struct {
	Map_id                string  `json:"map_id" default:"source"`
	Load_process_id       string  `json:"load_process_id" default:"source"`
	Error_id              string  `json:"error_id" default:"source"`
	Active_ind            *string `json:"active_ind" default:""`
	Delete_error_ind      *string `json:"delete_error_ind" default:""`
	Effective_date        *string `json:"effective_date" default:""`
	Error_cause_desc      *string `json:"error_cause_desc" default:""`
	Error_cause_type      *string `json:"error_cause_type" default:""`
	Error_code            *string `json:"error_code" default:""`
	Error_date            *string `json:"error_date" default:""`
	Error_message         *string `json:"error_message" default:""`
	Expiry_date           *string `json:"expiry_date" default:""`
	Insert_error_ind      *string `json:"insert_error_ind" default:""`
	Ppdm_guid             *string `json:"ppdm_guid" default:""`
	Ref_map_detail_obs_no *int    `json:"ref_map_detail_obs_no" default:""`
	Remark                *string `json:"remark" default:""`
	Resolution_desc       *string `json:"resolution_desc" default:""`
	Resolution_type       *string `json:"resolution_type" default:""`
	Source                *string `json:"source" default:""`
	Update_error_ind      *string `json:"update_error_ind" default:""`
	Row_changed_by        *string `json:"row_changed_by" default:""`
	Row_changed_date      *string `json:"row_changed_date" default:""`
	Row_created_by        *string `json:"row_created_by" default:""`
	Row_created_date      *string `json:"row_created_date" default:""`
	Row_effective_date    *string `json:"row_effective_date" default:""`
	Row_expiry_date       *string `json:"row_expiry_date" default:""`
	Row_quality           *string `json:"row_quality" default:""`
}
