package dto

type Equipment_maint_status struct {
	Equipment_id         string  `json:"equipment_id" default:"source"`
	Equip_maint_id       string  `json:"equip_maint_id" default:"source"`
	Status_id            string  `json:"status_id" default:"source"`
	Active_ind           *string `json:"active_ind" default:""`
	Effective_date       *string `json:"effective_date" default:""`
	Expiry_date          *string `json:"expiry_date" default:""`
	Maintain_status      *string `json:"maintain_status" default:""`
	Maintain_status_type *string `json:"maintain_status_type" default:""`
	Ppdm_guid            *string `json:"ppdm_guid" default:""`
	Remark               *string `json:"remark" default:""`
	Source               *string `json:"source" default:""`
	Row_changed_by       *string `json:"row_changed_by" default:""`
	Row_changed_date     *string `json:"row_changed_date" default:""`
	Row_created_by       *string `json:"row_created_by" default:""`
	Row_created_date     *string `json:"row_created_date" default:""`
	Row_effective_date   *string `json:"row_effective_date" default:""`
	Row_expiry_date      *string `json:"row_expiry_date" default:""`
	Row_quality          *string `json:"row_quality" default:""`
}
