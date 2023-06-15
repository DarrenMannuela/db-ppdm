package dto

type Anl_remark struct {
	Analysis_id            string  `json:"analysis_id" default:"source"`
	Anl_source             string  `json:"anl_source" default:"source"`
	Remark_seq_no          int     `json:"remark_seq_no" default:"1"`
	Active_ind             *string `json:"active_ind" default:""`
	Effective_date         *string `json:"effective_date" default:""`
	Expiry_date            *string `json:"expiry_date" default:""`
	Ppdm_guid              *string `json:"ppdm_guid" default:""`
	Referenced_column_name *string `json:"referenced_column_name" default:""`
	Referenced_ppdm_guid   *string `json:"referenced_ppdm_guid" default:""`
	Referenced_system_id   *string `json:"referenced_system_id" default:""`
	Referenced_table_name  *string `json:"referenced_table_name" default:""`
	Remark                 *string `json:"remark" default:""`
	Remark_ba_id           *string `json:"remark_ba_id" default:""`
	Remark_type            *string `json:"remark_type" default:""`
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
