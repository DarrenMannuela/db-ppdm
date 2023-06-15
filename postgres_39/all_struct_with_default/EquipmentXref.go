package dto

type Equipment_xref struct {
	Equipment_id          string  `json:"equipment_id" default:"source"`
	Equipment_part_id     string  `json:"equipment_part_id" default:"source"`
	Equipment_xref_obs_no int     `json:"equipment_xref_obs_no" default:"1"`
	Active_ind            *string `json:"active_ind" default:""`
	Commission_date       *string `json:"commission_date" default:""`
	Decommission_date     *string `json:"decommission_date" default:""`
	Description           *string `json:"description" default:""`
	Effective_date        *string `json:"effective_date" default:""`
	Equip_xref_type       *string `json:"equip_xref_type" default:""`
	Expiry_date           *string `json:"expiry_date" default:""`
	Interchangable_ind    *string `json:"interchangable_ind" default:""`
	Ppdm_guid             *string `json:"ppdm_guid" default:""`
	Remark                *string `json:"remark" default:""`
	Source                *string `json:"source" default:""`
	Row_changed_by        *string `json:"row_changed_by" default:""`
	Row_changed_date      *string `json:"row_changed_date" default:""`
	Row_created_by        *string `json:"row_created_by" default:""`
	Row_created_date      *string `json:"row_created_date" default:""`
	Row_effective_date    *string `json:"row_effective_date" default:""`
	Row_expiry_date       *string `json:"row_expiry_date" default:""`
	Row_quality           *string `json:"row_quality" default:""`
}
