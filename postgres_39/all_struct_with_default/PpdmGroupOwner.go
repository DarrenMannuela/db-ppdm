package dto

type Ppdm_group_owner struct {
	Group_id               string  `json:"group_id" default:"source"`
	Owner_obs_no           int     `json:"owner_obs_no" default:"1"`
	Active_ind             *string `json:"active_ind" default:""`
	Default_unit_system_id *string `json:"default_unit_system_id" default:""`
	Effective_date         *string `json:"effective_date" default:""`
	Expiry_date            *string `json:"expiry_date" default:""`
	Organization_id        *string `json:"organization_id" default:""`
	Organization_seq_no    *int    `json:"organization_seq_no" default:""`
	Owner_ba_id            *string `json:"owner_ba_id" default:""`
	Owner_role             *string `json:"owner_role" default:""`
	Ppdm_guid              *string `json:"ppdm_guid" default:""`
	Remark                 *string `json:"remark" default:""`
	Source                 *string `json:"source" default:""`
	Sw_application_id      *string `json:"sw_application_id" default:""`
	Row_changed_by         *string `json:"row_changed_by" default:""`
	Row_changed_date       *string `json:"row_changed_date" default:""`
	Row_created_by         *string `json:"row_created_by" default:""`
	Row_created_date       *string `json:"row_created_date" default:""`
	Row_effective_date     *string `json:"row_effective_date" default:""`
	Row_expiry_date        *string `json:"row_expiry_date" default:""`
	Row_quality            *string `json:"row_quality" default:""`
}
