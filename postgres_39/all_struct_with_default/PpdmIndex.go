package dto

type Ppdm_index struct {
	System_id          string  `json:"system_id" default:"source"`
	Table_name         string  `json:"table_name" default:"source"`
	Index_id           string  `json:"index_id" default:"N"`
	Active_ind         *string `json:"active_ind" default:""`
	Effective_date     *string `json:"effective_date" default:""`
	Expiry_date        *string `json:"expiry_date" default:""`
	Extension_ind      *string `json:"extension_ind" default:""`
	Index_category     *string `json:"index_category" default:""`
	Index_name         *string `json:"index_name" default:""`
	Ppdm_guid          *string `json:"ppdm_guid" default:""`
	Remark             *string `json:"remark" default:""`
	Source             *string `json:"source" default:""`
	Unique_ind         *string `json:"unique_ind" default:""`
	Row_changed_by     *string `json:"row_changed_by" default:""`
	Row_changed_date   *string `json:"row_changed_date" default:""`
	Row_created_by     *string `json:"row_created_by" default:""`
	Row_created_date   *string `json:"row_created_date" default:""`
	Row_effective_date *string `json:"row_effective_date" default:""`
	Row_expiry_date    *string `json:"row_expiry_date" default:""`
	Row_quality        *string `json:"row_quality" default:""`
}
