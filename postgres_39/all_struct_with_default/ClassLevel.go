package dto

type Class_level struct {
	Classification_system_id string  `json:"classification_system_id" default:"source"`
	Class_level_id           string  `json:"class_level_id" default:"source"`
	Active_ind               *string `json:"active_ind" default:""`
	Effective_date           *string `json:"effective_date" default:""`
	Expiry_date              *string `json:"expiry_date" default:""`
	Level_definition         *string `json:"level_definition" default:""`
	Level_name               *string `json:"level_name" default:""`
	Level_ref_num            *string `json:"level_ref_num" default:""`
	Level_seq_no             *int    `json:"level_seq_no" default:""`
	Level_type               *string `json:"level_type" default:""`
	Ppdm_guid                *string `json:"ppdm_guid" default:""`
	Remark                   *string `json:"remark" default:""`
	Retention_period         *string `json:"retention_period" default:""`
	Source                   *string `json:"source" default:""`
	Row_changed_by           *string `json:"row_changed_by" default:""`
	Row_changed_date         *string `json:"row_changed_date" default:""`
	Row_created_by           *string `json:"row_created_by" default:""`
	Row_created_date         *string `json:"row_created_date" default:""`
	Row_effective_date       *string `json:"row_effective_date" default:""`
	Row_expiry_date          *string `json:"row_expiry_date" default:""`
	Row_quality              *string `json:"row_quality" default:""`
}
