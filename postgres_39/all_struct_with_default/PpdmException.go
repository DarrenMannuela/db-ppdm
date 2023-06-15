package dto

type Ppdm_exception struct {
	Pe_id                string  `json:"pe_id" default:"source"`
	Active_ind           *string `json:"active_ind" default:""`
	Constraint_full_name *string `json:"constraint_full_name" default:""`
	Effective_date       *string `json:"effective_date" default:""`
	Expiry_date          *string `json:"expiry_date" default:""`
	Owner                *string `json:"owner" default:""`
	Ppdm_guid            *string `json:"ppdm_guid" default:""`
	Remark               *string `json:"remark" default:""`
	Row_id               *string `json:"row_id" default:""`
	Source               *string `json:"source" default:""`
	System_id            *string `json:"system_id" default:""`
	Table_name           *string `json:"table_name" default:""`
	Row_changed_by       *string `json:"row_changed_by" default:""`
	Row_changed_date     *string `json:"row_changed_date" default:""`
	Row_created_by       *string `json:"row_created_by" default:""`
	Row_created_date     *string `json:"row_created_date" default:""`
	Row_effective_date   *string `json:"row_effective_date" default:""`
	Row_expiry_date      *string `json:"row_expiry_date" default:""`
	Row_quality          *string `json:"row_quality" default:""`
}
