package dto

type Ppdm_system struct {
	System_id             string  `json:"system_id" default:"source"`
	Active_ind            *string `json:"active_ind" default:""`
	Business_owner_ba_id  *string `json:"business_owner_ba_id" default:""`
	Connect_string        *string `json:"connect_string" default:""`
	Creator_ba_id         *string `json:"creator_ba_id" default:""`
	Effective_date        *string `json:"effective_date" default:""`
	Expiry_date           *string `json:"expiry_date" default:""`
	Operating_system      *string `json:"operating_system" default:""`
	Ppdm_guid             *string `json:"ppdm_guid" default:""`
	Ppdm_system_type      *string `json:"ppdm_system_type" default:""`
	Rdbms_id              *string `json:"rdbms_id" default:""`
	Remark                *string `json:"remark" default:""`
	Source                *string `json:"source" default:""`
	System_long_name      *string `json:"system_long_name" default:""`
	Technical_owner_ba_id *string `json:"technical_owner_ba_id" default:""`
	Version_num           *string `json:"version_num" default:""`
	Row_changed_by        *string `json:"row_changed_by" default:""`
	Row_changed_date      *string `json:"row_changed_date" default:""`
	Row_created_by        *string `json:"row_created_by" default:""`
	Row_created_date      *string `json:"row_created_date" default:""`
	Row_effective_date    *string `json:"row_effective_date" default:""`
	Row_expiry_date       *string `json:"row_expiry_date" default:""`
	Row_quality           *string `json:"row_quality" default:""`
}
