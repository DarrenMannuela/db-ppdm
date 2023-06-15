package dto

type Ppdm_property_set struct {
	Property_set_id    string  `json:"property_set_id" default:"source"`
	Active_ind         *string `json:"active_ind" default:""`
	Effective_date     *string `json:"effective_date" default:""`
	Expiry_date        *string `json:"expiry_date" default:""`
	Ppdm_guid          *string `json:"ppdm_guid" default:""`
	Property_set_name  *string `json:"property_set_name" default:""`
	Remark             *string `json:"remark" default:""`
	Source             *string `json:"source" default:""`
	Use_system_id      *string `json:"use_system_id" default:""`
	Use_table_name     *string `json:"use_table_name" default:""`
	Row_changed_by     *string `json:"row_changed_by" default:""`
	Row_changed_date   *string `json:"row_changed_date" default:""`
	Row_created_by     *string `json:"row_created_by" default:""`
	Row_created_date   *string `json:"row_created_date" default:""`
	Row_effective_date *string `json:"row_effective_date" default:""`
	Row_expiry_date    *string `json:"row_expiry_date" default:""`
	Row_quality        *string `json:"row_quality" default:""`
}
