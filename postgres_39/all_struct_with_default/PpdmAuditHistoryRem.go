package dto

type Ppdm_audit_history_rem struct {
	System_id          string  `json:"system_id" default:"source"`
	Table_name         string  `json:"table_name" default:"source"`
	Column_name        string  `json:"column_name" default:"source"`
	Audit_row_guid     string  `json:"audit_row_guid" default:"source"`
	Audit_seq_no       int     `json:"audit_seq_no" default:"1"`
	Remark_seq_no      int     `json:"remark_seq_no" default:"1"`
	Active_ind         *string `json:"active_ind" default:""`
	Audit_datetime     *string `json:"audit_datetime" default:""`
	Effective_date     *string `json:"effective_date" default:""`
	Expiry_date        *string `json:"expiry_date" default:""`
	Ppdm_guid          *string `json:"ppdm_guid" default:""`
	Remark             *string `json:"remark" default:""`
	Remark_about_ba_id *string `json:"remark_about_ba_id" default:""`
	Remark_by_ba_id    *string `json:"remark_by_ba_id" default:""`
	Remark_for_ba_id   *string `json:"remark_for_ba_id" default:""`
	Remark_type        *string `json:"remark_type" default:""`
	Remark_use_type    *string `json:"remark_use_type" default:""`
	Retention_period   *string `json:"retention_period" default:""`
	Source             *string `json:"source" default:""`
	Row_changed_by     *string `json:"row_changed_by" default:""`
	Row_changed_date   *string `json:"row_changed_date" default:""`
	Row_created_by     *string `json:"row_created_by" default:""`
	Row_created_date   *string `json:"row_created_date" default:""`
	Row_effective_date *string `json:"row_effective_date" default:""`
	Row_expiry_date    *string `json:"row_expiry_date" default:""`
	Row_quality        *string `json:"row_quality" default:""`
}
