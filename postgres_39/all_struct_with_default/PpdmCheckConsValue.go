package dto

type Ppdm_check_cons_value struct {
	System_id          string  `json:"system_id" default:"source"`
	Table_name         string  `json:"table_name" default:"source"`
	Column_name        string  `json:"column_name" default:"source"`
	Check_value        string  `json:"check_value" default:"source"`
	Active_ind         *string `json:"active_ind" default:""`
	Check_cons_name    *string `json:"check_cons_name" default:""`
	Constraint_name    *string `json:"constraint_name" default:""`
	Effective_date     *string `json:"effective_date" default:""`
	Expiry_date        *string `json:"expiry_date" default:""`
	Extension_ind      *string `json:"extension_ind" default:""`
	Ppdm_guid          *string `json:"ppdm_guid" default:""`
	Remark             *string `json:"remark" default:""`
	Source             *string `json:"source" default:""`
	Value_long_name    *string `json:"value_long_name" default:""`
	Row_changed_by     *string `json:"row_changed_by" default:""`
	Row_changed_date   *string `json:"row_changed_date" default:""`
	Row_created_by     *string `json:"row_created_by" default:""`
	Row_created_date   *string `json:"row_created_date" default:""`
	Row_effective_date *string `json:"row_effective_date" default:""`
	Row_expiry_date    *string `json:"row_expiry_date" default:""`
	Row_quality        *string `json:"row_quality" default:""`
}
