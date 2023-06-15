package dto

type Anl_valid_table_result struct {
	Method_set_id      string  `json:"method_set_id" default:"source"`
	Method_id          string  `json:"method_id" default:"source"`
	Ppdm_system_id     string  `json:"ppdm_system_id" default:"source"`
	Ppdm_table_name    string  `json:"ppdm_table_name" default:"source"`
	Active_ind         *string `json:"active_ind" default:""`
	Effective_date     *string `json:"effective_date" default:""`
	Expiry_date        *string `json:"expiry_date" default:""`
	Mandatory_ind      *string `json:"mandatory_ind" default:""`
	Ppdm_guid          *string `json:"ppdm_guid" default:""`
	Remark             *string `json:"remark" default:""`
	Source             *string `json:"source" default:""`
	Row_changed_by     *string `json:"row_changed_by" default:""`
	Row_changed_date   *string `json:"row_changed_date" default:""`
	Row_created_by     *string `json:"row_created_by" default:""`
	Row_created_date   *string `json:"row_created_date" default:""`
	Row_effective_date *string `json:"row_effective_date" default:""`
	Row_expiry_date    *string `json:"row_expiry_date" default:""`
	Row_quality        *string `json:"row_quality" default:""`
}
