package dto

type Ppdm_code_version_xref struct {
	System_id          string  `json:"system_id" default:"source"`
	Table_name         string  `json:"table_name" default:"source"`
	Source             string  `json:"source" default:"source"`
	Version_obs_no     int     `json:"version_obs_no" default:"1"`
	System_id2         string  `json:"system_id_2" default:"source"`
	Table_name2        string  `json:"table_name_2" default:"source"`
	Source2            string  `json:"source_2" default:"source"`
	Version_obs_no2    int     `json:"version_obs_no_2" default:"1"`
	Active_ind         *string `json:"active_ind" default:""`
	Code_xref_type     *string `json:"code_xref_type" default:""`
	Effective_date     *string `json:"effective_date" default:""`
	Equivalence_desc   *string `json:"equivalence_desc" default:""`
	Expiry_date        *string `json:"expiry_date" default:""`
	Ppdm_guid          *string `json:"ppdm_guid" default:""`
	Remark             *string `json:"remark" default:""`
	Source_document_id *string `json:"source_document_id" default:""`
	User_rule_desc     *string `json:"user_rule_desc" default:""`
	Row_changed_by     *string `json:"row_changed_by" default:""`
	Row_changed_date   *string `json:"row_changed_date" default:""`
	Row_created_by     *string `json:"row_created_by" default:""`
	Row_created_date   *string `json:"row_created_date" default:""`
	Row_effective_date *string `json:"row_effective_date" default:""`
	Row_expiry_date    *string `json:"row_expiry_date" default:""`
	Row_quality        *string `json:"row_quality" default:""`
}
