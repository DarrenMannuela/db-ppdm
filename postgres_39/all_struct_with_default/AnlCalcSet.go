package dto

type Anl_calc_set struct {
	Calc_set_id             string  `json:"calc_set_id" default:"source"`
	Active_ind              *string `json:"active_ind" default:""`
	Effective_date          *string `json:"effective_date" default:""`
	Expiry_date             *string `json:"expiry_date" default:""`
	Method_id               *string `json:"method_id" default:""`
	Method_set_id           *string `json:"method_set_id" default:""`
	Owner_ba_id             *string `json:"owner_ba_id" default:""`
	Ppdm_guid               *string `json:"ppdm_guid" default:""`
	Preferred_name          *string `json:"preferred_name" default:""`
	Remark                  *string `json:"remark" default:""`
	Software_application_id *string `json:"software_application_id" default:""`
	Source                  *string `json:"source" default:""`
	Source_document_id      *string `json:"source_document_id" default:""`
	Row_changed_by          *string `json:"row_changed_by" default:""`
	Row_changed_date        *string `json:"row_changed_date" default:""`
	Row_created_by          *string `json:"row_created_by" default:""`
	Row_created_date        *string `json:"row_created_date" default:""`
	Row_effective_date      *string `json:"row_effective_date" default:""`
	Row_expiry_date         *string `json:"row_expiry_date" default:""`
	Row_quality             *string `json:"row_quality" default:""`
}
