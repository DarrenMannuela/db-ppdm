package dto

type Ppdm_audit_history struct {
	System_id                   string   `json:"system_id" default:"source"`
	Table_name                  string   `json:"table_name" default:"source"`
	Column_name                 string   `json:"column_name" default:"source"`
	Audit_row_guid              string   `json:"audit_row_guid" default:"source"`
	Audit_seq_no                int      `json:"audit_seq_no" default:"1"`
	Active_ind                  *string  `json:"active_ind" default:""`
	Audit_authorized_by_ba_id   *string  `json:"audit_authorized_by_ba_id" default:""`
	Audit_created_by_ba_id      *string  `json:"audit_created_by_ba_id" default:""`
	Audit_datetime              *string  `json:"audit_datetime" default:""`
	Audit_desc                  *string  `json:"audit_desc" default:""`
	Audit_reason                *string  `json:"audit_reason" default:""`
	Audit_source                *string  `json:"audit_source" default:""`
	Audit_type                  *string  `json:"audit_type" default:""`
	Audit_verified_by_ba_id     *string  `json:"audit_verified_by_ba_id" default:""`
	Data_type                   *string  `json:"data_type" default:""`
	Effective_date              *string  `json:"effective_date" default:""`
	Expiry_date                 *string  `json:"expiry_date" default:""`
	New_date_value              *string  `json:"new_date_value" default:""`
	New_numeric_value           *float64 `json:"new_numeric_value" default:""`
	New_numeric_value_ouom      *string  `json:"new_numeric_value_ouom" default:""`
	New_numeric_value_uom       *string  `json:"new_numeric_value_uom" default:""`
	New_text_value              *string  `json:"new_text_value" default:""`
	Null_description            *string  `json:"null_description" default:""`
	Original_date_value         *string  `json:"original_date_value" default:""`
	Original_numeric_value      *float64 `json:"original_numeric_value" default:""`
	Original_numeric_value_ouom *string  `json:"original_numeric_value_ouom" default:""`
	Original_numeric_value_uom  *string  `json:"original_numeric_value_uom" default:""`
	Original_text_value         *string  `json:"original_text_value" default:""`
	Ppdm_guid                   *string  `json:"ppdm_guid" default:""`
	Remark                      *string  `json:"remark" default:""`
	Retention_period            *string  `json:"retention_period" default:""`
	Source                      *string  `json:"source" default:""`
	Row_changed_by              *string  `json:"row_changed_by" default:""`
	Row_changed_date            *string  `json:"row_changed_date" default:""`
	Row_created_by              *string  `json:"row_created_by" default:""`
	Row_created_date            *string  `json:"row_created_date" default:""`
	Row_effective_date          *string  `json:"row_effective_date" default:""`
	Row_expiry_date             *string  `json:"row_expiry_date" default:""`
	Row_quality                 *string  `json:"row_quality" default:""`
}