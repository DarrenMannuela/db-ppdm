package dto

type Paleo_summary struct {
	Paleo_summary_id           string   `json:"paleo_summary_id" default:"source"`
	Access_condition           *string  `json:"access_condition" default:""`
	Active_ind                 *string  `json:"active_ind" default:""`
	Analysis_end_date          *string  `json:"analysis_end_date" default:""`
	Analysis_start_date        *string  `json:"analysis_start_date" default:""`
	Confidential_ind           *string  `json:"confidential_ind" default:""`
	Confidential_reason        *string  `json:"confidential_reason" default:""`
	Confidential_release_date  *string  `json:"confidential_release_date" default:""`
	Confidential_term          *string  `json:"confidential_term" default:""`
	Confidential_type          *string  `json:"confidential_type" default:""`
	Diversity_count            *int     `json:"diversity_count" default:""`
	Effective_date             *string  `json:"effective_date" default:""`
	Expiry_date                *string  `json:"expiry_date" default:""`
	Ppdm_guid                  *string  `json:"ppdm_guid" default:""`
	Preservation_quality       *string  `json:"preservation_quality" default:""`
	Reference_num              *string  `json:"reference_num" default:""`
	Remark                     *string  `json:"remark" default:""`
	Report_date                *string  `json:"report_date" default:""`
	Report_title               *string  `json:"report_title" default:""`
	Source                     *string  `json:"source" default:""`
	Total_sample_interval      *float64 `json:"total_sample_interval" default:""`
	Total_sample_interval_ouom *string  `json:"total_sample_interval_ouom" default:""`
	Row_changed_by             *string  `json:"row_changed_by" default:""`
	Row_changed_date           *string  `json:"row_changed_date" default:""`
	Row_created_by             *string  `json:"row_created_by" default:""`
	Row_created_date           *string  `json:"row_created_date" default:""`
	Row_effective_date         *string  `json:"row_effective_date" default:""`
	Row_expiry_date            *string  `json:"row_expiry_date" default:""`
	Row_quality                *string  `json:"row_quality" default:""`
}
