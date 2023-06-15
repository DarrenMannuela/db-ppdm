package dto

type Class_system struct {
	Classification_system_id string  `json:"classification_system_id" default:"source"`
	Active_ind               *string `json:"active_ind" default:""`
	Class_dimension          *string `json:"class_dimension" default:""`
	Effective_date           *string `json:"effective_date" default:""`
	Expiry_date              *string `json:"expiry_date" default:""`
	Ppdm_guid                *string `json:"ppdm_guid" default:""`
	Remark                   *string `json:"remark" default:""`
	Source                   *string `json:"source" default:""`
	Source_document_id       *string `json:"source_document_id" default:""`
	System_definition        *string `json:"system_definition" default:""`
	System_name              *string `json:"system_name" default:""`
	System_owner             *string `json:"system_owner" default:""`
	System_ref_num           *string `json:"system_ref_num" default:""`
	System_version           *string `json:"system_version" default:""`
	Row_changed_by           *string `json:"row_changed_by" default:""`
	Row_changed_date         *string `json:"row_changed_date" default:""`
	Row_created_by           *string `json:"row_created_by" default:""`
	Row_created_date         *string `json:"row_created_date" default:""`
	Row_effective_date       *string `json:"row_effective_date" default:""`
	Row_expiry_date          *string `json:"row_expiry_date" default:""`
	Row_quality              *string `json:"row_quality" default:""`
}
