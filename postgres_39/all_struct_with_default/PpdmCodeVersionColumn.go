package dto

type Ppdm_code_version_column struct {
	System_id          string  `json:"system_id" default:"source"`
	Table_name         string  `json:"table_name" default:"source"`
	Source             string  `json:"source" default:"source"`
	Version_obs_no     int     `json:"version_obs_no" default:"1"`
	Column_name        string  `json:"column_name" default:"source"`
	Abbreviation       *string `json:"abbreviation" default:""`
	Active_ind         *string `json:"active_ind" default:""`
	Definition         *string `json:"definition" default:""`
	Effective_date     *string `json:"effective_date" default:""`
	Expiry_date        *string `json:"expiry_date" default:""`
	Language           *string `json:"language" default:""`
	Long_name          *string `json:"long_name" default:""`
	Ppdm_guid          *string `json:"ppdm_guid" default:""`
	Primary_key        *string `json:"primary_key" default:""`
	Remark             *string `json:"remark" default:""`
	Short_name         *string `json:"short_name" default:""`
	Row_changed_by     *string `json:"row_changed_by" default:""`
	Row_changed_date   *string `json:"row_changed_date" default:""`
	Row_created_by     *string `json:"row_created_by" default:""`
	Row_created_date   *string `json:"row_created_date" default:""`
	Row_effective_date *string `json:"row_effective_date" default:""`
	Row_expiry_date    *string `json:"row_expiry_date" default:""`
	Row_quality        *string `json:"row_quality" default:""`
}
