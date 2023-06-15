package dto

type Class_level_xref struct {
	Classification_system_id  string  `json:"classification_system_id" default:"source"`
	Class_level_id            string  `json:"class_level_id" default:"source"`
	Classification_system_id2 string  `json:"classification_system_id_2" default:"source"`
	Class_level_id2           string  `json:"class_level_id_2" default:"source"`
	Xref_id                   string  `json:"xref_id" default:"source"`
	Active_ind                *string `json:"active_ind" default:""`
	Effective_date            *string `json:"effective_date" default:""`
	Expiry_date               *string `json:"expiry_date" default:""`
	Level_xref_type           *string `json:"level_xref_type" default:""`
	Ppdm_guid                 *string `json:"ppdm_guid" default:""`
	Remark                    *string `json:"remark" default:""`
	Source                    *string `json:"source" default:""`
	Source_document_id        *string `json:"source_document_id" default:""`
	Row_changed_by            *string `json:"row_changed_by" default:""`
	Row_changed_date          *string `json:"row_changed_date" default:""`
	Row_created_by            *string `json:"row_created_by" default:""`
	Row_created_date          *string `json:"row_created_date" default:""`
	Row_effective_date        *string `json:"row_effective_date" default:""`
	Row_expiry_date           *string `json:"row_expiry_date" default:""`
	Row_quality               *string `json:"row_quality" default:""`
}
