package dto

type Ppdm_constraint struct {
	System_id                  string  `json:"system_id" default:"source"`
	Table_name                 string  `json:"table_name" default:"source"`
	Constraint_name            string  `json:"constraint_name" default:"source"`
	Active_ind                 *string `json:"active_ind" default:""`
	Constraint_full_name       *string `json:"constraint_full_name" default:""`
	Constraint_type            *string `json:"constraint_type" default:""`
	Effective_date             *string `json:"effective_date" default:""`
	Expiry_date                *string `json:"expiry_date" default:""`
	Extension_ind              *string `json:"extension_ind" default:""`
	Ppdm_guid                  *string `json:"ppdm_guid" default:""`
	Referenced_constraint_name *string `json:"referenced_constraint_name" default:""`
	Referenced_system_id       *string `json:"referenced_system_id" default:""`
	Referenced_table_name      *string `json:"referenced_table_name" default:""`
	Remark                     *string `json:"remark" default:""`
	Source                     *string `json:"source" default:""`
	Row_changed_by             *string `json:"row_changed_by" default:""`
	Row_changed_date           *string `json:"row_changed_date" default:""`
	Row_created_by             *string `json:"row_created_by" default:""`
	Row_created_date           *string `json:"row_created_date" default:""`
	Row_effective_date         *string `json:"row_effective_date" default:""`
	Row_expiry_date            *string `json:"row_expiry_date" default:""`
	Row_quality                *string `json:"row_quality" default:""`
}
