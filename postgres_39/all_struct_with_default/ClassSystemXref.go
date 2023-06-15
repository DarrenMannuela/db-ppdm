package dto

type Class_system_xref struct {
	Classification_system_id  string  `json:"classification_system_id" default:"source"`
	Classification_system_id2 string  `json:"classification_system_id_2" default:"source"`
	Equiv_id                  string  `json:"equiv_id" default:"source"`
	Active_ind                *string `json:"active_ind" default:""`
	Effective_date            *string `json:"effective_date" default:""`
	Expiry_date               *string `json:"expiry_date" default:""`
	Ppdm_guid                 *string `json:"ppdm_guid" default:""`
	Remark                    *string `json:"remark" default:""`
	Source                    *string `json:"source" default:""`
	Source_document           *string `json:"source_document" default:""`
	System_xref_type          *string `json:"system_xref_type" default:""`
	Row_changed_by            *string `json:"row_changed_by" default:""`
	Row_changed_date          *string `json:"row_changed_date" default:""`
	Row_created_by            *string `json:"row_created_by" default:""`
	Row_created_date          *string `json:"row_created_date" default:""`
	Row_effective_date        *string `json:"row_effective_date" default:""`
	Row_expiry_date           *string `json:"row_expiry_date" default:""`
	Row_quality               *string `json:"row_quality" default:""`
}
