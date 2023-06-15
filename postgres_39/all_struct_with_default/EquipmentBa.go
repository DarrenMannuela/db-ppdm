package dto

type Equipment_ba struct {
	Equip_id           string  `json:"equip_id" default:"source"`
	Ba_obs_no          int     `json:"ba_obs_no" default:"1"`
	Acquire_date       *string `json:"acquire_date" default:""`
	Active_ind         *string `json:"active_ind" default:""`
	Effective_date     *string `json:"effective_date" default:""`
	Equip_ba_id        *string `json:"equip_ba_id" default:""`
	Equip_lease_ind    *string `json:"equip_lease_ind" default:""`
	Equip_owned_ind    *string `json:"equip_owned_ind" default:""`
	Equip_rent_ind     *string `json:"equip_rent_ind" default:""`
	Expiry_date        *string `json:"expiry_date" default:""`
	Finance_id         *string `json:"finance_id" default:""`
	Ppdm_guid          *string `json:"ppdm_guid" default:""`
	Preferred_name     *string `json:"preferred_name" default:""`
	Purchase_date      *string `json:"purchase_date" default:""`
	Release_date       *string `json:"release_date" default:""`
	Remark             *string `json:"remark" default:""`
	Role_type          *string `json:"role_type" default:""`
	Source             *string `json:"source" default:""`
	Row_changed_by     *string `json:"row_changed_by" default:""`
	Row_changed_date   *string `json:"row_changed_date" default:""`
	Row_created_by     *string `json:"row_created_by" default:""`
	Row_created_date   *string `json:"row_created_date" default:""`
	Row_effective_date *string `json:"row_effective_date" default:""`
	Row_expiry_date    *string `json:"row_expiry_date" default:""`
	Row_quality        *string `json:"row_quality" default:""`
}
