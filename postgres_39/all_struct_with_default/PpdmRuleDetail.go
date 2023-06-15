package dto

type Ppdm_rule_detail struct {
	Rule_id                string   `json:"rule_id" default:"source"`
	Detail_seq_no          int      `json:"detail_seq_no" default:"1"`
	Active_ind             *string  `json:"active_ind" default:""`
	Average_value          *float64 `json:"average_value" default:""`
	Average_value_ouom     *string  `json:"average_value_ouom" default:""`
	Average_value_uom      *string  `json:"average_value_uom" default:""`
	Boolean_rule           *string  `json:"boolean_rule" default:""`
	Business_rule          *string  `json:"business_rule" default:""`
	Date_format_desc       *string  `json:"date_format_desc" default:""`
	Effective_date         *string  `json:"effective_date" default:""`
	Expiry_date            *string  `json:"expiry_date" default:""`
	Max_date               *string  `json:"max_date" default:""`
	Max_value              *float64 `json:"max_value" default:""`
	Max_value_ouom         *string  `json:"max_value_ouom" default:""`
	Max_value_uom          *string  `json:"max_value_uom" default:""`
	Min_date               *string  `json:"min_date" default:""`
	Min_value              *float64 `json:"min_value" default:""`
	Min_value_ouom         *string  `json:"min_value_ouom" default:""`
	Min_value_uom          *string  `json:"min_value_uom" default:""`
	Ppdm_guid              *string  `json:"ppdm_guid" default:""`
	Referenced_column_name *string  `json:"referenced_column_name" default:""`
	Reference_column_name2 *string  `json:"reference_column_name_2" default:""`
	Reference_system_id    *string  `json:"reference_system_id" default:""`
	Reference_table_name   *string  `json:"reference_table_name" default:""`
	Reference_table_name2  *string  `json:"reference_table_name_2" default:""`
	Reference_value        *float64 `json:"reference_value" default:""`
	Reference_value_ouom   *string  `json:"reference_value_ouom" default:""`
	Reference_value_type   *string  `json:"reference_value_type" default:""`
	Reference_value_uom    *string  `json:"reference_value_uom" default:""`
	Remark                 *string  `json:"remark" default:""`
	Rule_desc              *string  `json:"rule_desc" default:""`
	Rule_detail_type       *string  `json:"rule_detail_type" default:""`
	Source                 *string  `json:"source" default:""`
	Row_changed_by         *string  `json:"row_changed_by" default:""`
	Row_changed_date       *string  `json:"row_changed_date" default:""`
	Row_created_by         *string  `json:"row_created_by" default:""`
	Row_created_date       *string  `json:"row_created_date" default:""`
	Row_effective_date     *string  `json:"row_effective_date" default:""`
	Row_expiry_date        *string  `json:"row_expiry_date" default:""`
	Row_quality            *string  `json:"row_quality" default:""`
}
