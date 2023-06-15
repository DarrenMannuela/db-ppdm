package dto

type Ppdm_quality_control struct {
	System_id                  string   `json:"system_id" default:"source"`
	Table_name                 string   `json:"table_name" default:"source"`
	Qc_seq_no                  int      `json:"qc_seq_no" default:"1"`
	Active_ind                 *string  `json:"active_ind" default:""`
	Checked_by_ba_id           *string  `json:"checked_by_ba_id" default:""`
	Column_name                *string  `json:"column_name" default:""`
	Current_date_value         *string  `json:"current_date_value" default:""`
	Current_numeric_value      *float64 `json:"current_numeric_value" default:""`
	Current_numeric_value_ouom *string  `json:"current_numeric_value_ouom" default:""`
	Current_numeric_value_uom  *string  `json:"current_numeric_value_uom" default:""`
	Current_row_guid           *string  `json:"current_row_guid" default:""`
	Current_text_value         *string  `json:"current_text_value" default:""`
	Data_type                  *string  `json:"data_type" default:""`
	Done_by_ba_id              *string  `json:"done_by_ba_id" default:""`
	Effective_date             *string  `json:"effective_date" default:""`
	Expiry_date                *string  `json:"expiry_date" default:""`
	Null_description           *string  `json:"null_description" default:""`
	Ppdm_guid                  *string  `json:"ppdm_guid" default:""`
	Qc_datetime                *string  `json:"qc_datetime" default:""`
	Qc_desc                    *string  `json:"qc_desc" default:""`
	Qc_status                  *string  `json:"qc_status" default:""`
	Qc_type                    *string  `json:"qc_type" default:""`
	Quality_type               *string  `json:"quality_type" default:""`
	Remark                     *string  `json:"remark" default:""`
	Retention_period           *string  `json:"retention_period" default:""`
	Source                     *string  `json:"source" default:""`
	Row_changed_by             *string  `json:"row_changed_by" default:""`
	Row_changed_date           *string  `json:"row_changed_date" default:""`
	Row_created_by             *string  `json:"row_created_by" default:""`
	Row_created_date           *string  `json:"row_created_date" default:""`
	Row_effective_date         *string  `json:"row_effective_date" default:""`
	Row_expiry_date            *string  `json:"row_expiry_date" default:""`
	Row_quality                *string  `json:"row_quality" default:""`
}
