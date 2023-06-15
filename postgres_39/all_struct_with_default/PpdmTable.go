package dto

type Ppdm_table struct {
	System_id          string   `json:"system_id" default:"source"`
	Table_name         string   `json:"table_name" default:"source"`
	Active_ind         *string  `json:"active_ind" default:""`
	Effective_date     *string  `json:"effective_date" default:""`
	Expiry_date        *string  `json:"expiry_date" default:""`
	Extension_ind      *string  `json:"extension_ind" default:""`
	Ppdm_guid          *string  `json:"ppdm_guid" default:""`
	Primary_key_name   *string  `json:"primary_key_name" default:""`
	Remark             *string  `json:"remark" default:""`
	Row_count          *float64 `json:"row_count" default:""`
	Row_count_date     *string  `json:"row_count_date" default:""`
	Source             *string  `json:"source" default:""`
	Table_comment      *string  `json:"table_comment" default:""`
	Table_type         *string  `json:"table_type" default:""`
	Row_changed_by     *string  `json:"row_changed_by" default:""`
	Row_changed_date   *string  `json:"row_changed_date" default:""`
	Row_created_by     *string  `json:"row_created_by" default:""`
	Row_created_date   *string  `json:"row_created_date" default:""`
	Row_effective_date *string  `json:"row_effective_date" default:""`
	Row_expiry_date    *string  `json:"row_expiry_date" default:""`
	Row_quality        *string  `json:"row_quality" default:""`
}
