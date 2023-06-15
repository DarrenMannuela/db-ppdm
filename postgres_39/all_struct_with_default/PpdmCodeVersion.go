package dto

type Ppdm_code_version struct {
	System_id            string  `json:"system_id" default:"source"`
	Table_name           string  `json:"table_name" default:"source"`
	Source               string  `json:"source" default:"source"`
	Version_obs_no       int     `json:"version_obs_no" default:"1"`
	Active_ind           *string `json:"active_ind" default:""`
	Definition           *string `json:"definition" default:""`
	Effective_date       *string `json:"effective_date" default:""`
	Expiry_date          *string `json:"expiry_date" default:""`
	Language             *string `json:"language" default:""`
	Ppdm_guid            *string `json:"ppdm_guid" default:""`
	Preferred_ind        *string `json:"preferred_ind" default:""`
	Reference_table_guid *string `json:"reference_table_guid" default:""`
	Reference_table_ind  *string `json:"reference_table_ind" default:""`
	Remark               *string `json:"remark" default:""`
	Source_document_id   *string `json:"source_document_id" default:""`
	Use_rule_desc        *string `json:"use_rule_desc" default:""`
	Version_owner_ba_id  *string `json:"version_owner_ba_id" default:""`
	Row_changed_by       *string `json:"row_changed_by" default:""`
	Row_changed_date     *string `json:"row_changed_date" default:""`
	Row_created_by       *string `json:"row_created_by" default:""`
	Row_created_date     *string `json:"row_created_date" default:""`
	Row_effective_date   *string `json:"row_effective_date" default:""`
	Row_expiry_date      *string `json:"row_expiry_date" default:""`
	Row_quality          *string `json:"row_quality" default:""`
}
