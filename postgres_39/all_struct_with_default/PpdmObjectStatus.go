package dto

type Ppdm_object_status struct {
	Status_id           string  `json:"status_id" default:"source"`
	Status_obs_no       int     `json:"status_obs_no" default:"1"`
	Active_ind          *string `json:"active_ind" default:""`
	Code_version_obs_no *int    `json:"code_version_obs_no" default:""`
	Code_version_source *string `json:"code_version_source" default:""`
	Column_name         *string `json:"column_name" default:""`
	Constraint_name     *string `json:"constraint_name" default:""`
	Effective_date      *string `json:"effective_date" default:""`
	Expiry_date         *string `json:"expiry_date" default:""`
	Index_id            *string `json:"index_id" default:""`
	Object_name         *string `json:"object_name" default:""`
	Object_status       *string `json:"object_status" default:""`
	Object_type         *string `json:"object_type" default:""`
	Ppdm_guid           *string `json:"ppdm_guid" default:""`
	Procedure_id        *string `json:"procedure_id" default:""`
	Property_set_id     *string `json:"property_set_id" default:""`
	Remark              *string `json:"remark" default:""`
	Rule_id             *string `json:"rule_id" default:""`
	Source              *string `json:"source" default:""`
	System_id           *string `json:"system_id" default:""`
	Table_name          *string `json:"table_name" default:""`
	Row_changed_by      *string `json:"row_changed_by" default:""`
	Row_changed_date    *string `json:"row_changed_date" default:""`
	Row_created_by      *string `json:"row_created_by" default:""`
	Row_created_date    *string `json:"row_created_date" default:""`
	Row_effective_date  *string `json:"row_effective_date" default:""`
	Row_expiry_date     *string `json:"row_expiry_date" default:""`
	Row_quality         *string `json:"row_quality" default:""`
}
