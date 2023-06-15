package dto

type Class_level_type struct {
	Classification_system_id string  `json:"classification_system_id" default:"source"`
	Level_type               string  `json:"level_type" default:"source"`
	Abbreviation             *string `json:"abbreviation" default:""`
	Active_ind               *string `json:"active_ind" default:""`
	Effective_date           *string `json:"effective_date" default:""`
	Expiry_date              *string `json:"expiry_date" default:""`
	Level_order_seq_no       *int    `json:"level_order_seq_no" default:""`
	Long_name                *string `json:"long_name" default:""`
	Ppdm_guid                *string `json:"ppdm_guid" default:""`
	Ppdm_system_id           *string `json:"ppdm_system_id" default:""`
	Ppdm_table_name          *string `json:"ppdm_table_name" default:""`
	Remark                   *string `json:"remark" default:""`
	Selection_criteria       *string `json:"selection_criteria" default:""`
	Short_name               *string `json:"short_name" default:""`
	Source                   *string `json:"source" default:""`
	Row_changed_by           *string `json:"row_changed_by" default:""`
	Row_changed_date         *string `json:"row_changed_date" default:""`
	Row_created_by           *string `json:"row_created_by" default:""`
	Row_created_date         *string `json:"row_created_date" default:""`
	Row_effective_date       *string `json:"row_effective_date" default:""`
	Row_expiry_date          *string `json:"row_expiry_date" default:""`
	Row_quality              *string `json:"row_quality" default:""`
}
