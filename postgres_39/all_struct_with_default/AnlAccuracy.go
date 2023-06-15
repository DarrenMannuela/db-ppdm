package dto

type Anl_accuracy struct {
	Analysis_id            string  `json:"analysis_id" default:"source"`
	Anl_source             string  `json:"anl_source" default:"source"`
	Accuracy_obs_no        int     `json:"accuracy_obs_no" default:"1"`
	Accuracy_desc          *string `json:"accuracy_desc" default:""`
	Accuracy_type          *string `json:"accuracy_type" default:""`
	Active_ind             *string `json:"active_ind" default:""`
	Calculated_ind         *string `json:"calculated_ind" default:""`
	Calculate_method_id    *string `json:"calculate_method_id" default:""`
	Confidence_type        *string `json:"confidence_type" default:""`
	Effective_date         *string `json:"effective_date" default:""`
	Expiry_date            *string `json:"expiry_date" default:""`
	Measured_value_ind     *string `json:"measured_value_ind" default:""`
	Ppdm_guid              *string `json:"ppdm_guid" default:""`
	Raw_value_ind          *string `json:"raw_value_ind" default:""`
	Referenced_column_name *string `json:"referenced_column_name" default:""`
	Referenced_ppdm_guid   *string `json:"referenced_ppdm_guid" default:""`
	Referenced_system_id   *string `json:"referenced_system_id" default:""`
	Referenced_table_name  *string `json:"referenced_table_name" default:""`
	Remark                 *string `json:"remark" default:""`
	Repeatability_type     *string `json:"repeatability_type" default:""`
	Reported_ind           *string `json:"reported_ind" default:""`
	Source                 *string `json:"source" default:""`
	Step_seq_no            *int    `json:"step_seq_no" default:""`
	Row_changed_by         *string `json:"row_changed_by" default:""`
	Row_changed_date       *string `json:"row_changed_date" default:""`
	Row_created_by         *string `json:"row_created_by" default:""`
	Row_created_date       *string `json:"row_created_date" default:""`
	Row_effective_date     *string `json:"row_effective_date" default:""`
	Row_expiry_date        *string `json:"row_expiry_date" default:""`
	Row_quality            *string `json:"row_quality" default:""`
}
