package dto

type Ppdm_code_version_use struct {
	System_id          string  `json:"system_id" default:"source"`
	Table_name         string  `json:"table_name" default:"source"`
	Source             string  `json:"source" default:"source"`
	Version_obs_no     int     `json:"version_obs_no" default:"1"`
	Use_obs_no         int     `json:"use_obs_no" default:"1"`
	Active_ind         *string `json:"active_ind" default:""`
	Effective_date     *string `json:"effective_date" default:""`
	Expiry_date        *string `json:"expiry_date" default:""`
	Group_id           *string `json:"group_id" default:""`
	Ppdm_guid          *string `json:"ppdm_guid" default:""`
	Preferred_ind      *string `json:"preferred_ind" default:""`
	Procedure_id       *string `json:"procedure_id" default:""`
	Remark             *string `json:"remark" default:""`
	Source_document_id *string `json:"source_document_id" default:""`
	Sw_application_id  *string `json:"sw_application_id" default:""`
	Use_owner_ba_id    *string `json:"use_owner_ba_id" default:""`
	Use_rule_desc      *string `json:"use_rule_desc" default:""`
	Row_changed_by     *string `json:"row_changed_by" default:""`
	Row_changed_date   *string `json:"row_changed_date" default:""`
	Row_created_by     *string `json:"row_created_by" default:""`
	Row_created_date   *string `json:"row_created_date" default:""`
	Row_effective_date *string `json:"row_effective_date" default:""`
	Row_expiry_date    *string `json:"row_expiry_date" default:""`
	Row_quality        *string `json:"row_quality" default:""`
}
