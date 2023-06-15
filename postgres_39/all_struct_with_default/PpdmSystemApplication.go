package dto

type Ppdm_system_application struct {
	System_id                 string  `json:"system_id" default:"source"`
	Sw_application_id         string  `json:"sw_application_id" default:"source"`
	Application_seq_no        int     `json:"application_seq_no" default:"1"`
	Active_ind                *string `json:"active_ind" default:""`
	Application_function_desc *string `json:"application_function_desc" default:""`
	Effective_date            *string `json:"effective_date" default:""`
	Expiry_date               *string `json:"expiry_date" default:""`
	Language_id               *string `json:"language_id" default:""`
	Ppdm_guid                 *string `json:"ppdm_guid" default:""`
	Remark                    *string `json:"remark" default:""`
	Source                    *string `json:"source" default:""`
	Row_changed_by            *string `json:"row_changed_by" default:""`
	Row_changed_date          *string `json:"row_changed_date" default:""`
	Row_created_by            *string `json:"row_created_by" default:""`
	Row_created_date          *string `json:"row_created_date" default:""`
	Row_effective_date        *string `json:"row_effective_date" default:""`
	Row_expiry_date           *string `json:"row_expiry_date" default:""`
	Row_quality               *string `json:"row_quality" default:""`
}
