package dto

import (
	"time"
)

type Equipment_ba struct {
	Equip_id           string     `json:"equip_id"`
	Ba_obs_no          int        `json:"ba_obs_no"`
	Acquire_date       *time.Time `json:"acquire_date"`
	Active_ind         *string    `json:"active_ind"`
	Effective_date     *time.Time `json:"effective_date"`
	Equip_ba_id        *string    `json:"equip_ba_id"`
	Equip_lease_ind    *string    `json:"equip_lease_ind"`
	Equip_owned_ind    *string    `json:"equip_owned_ind"`
	Equip_rent_ind     *string    `json:"equip_rent_ind"`
	Expiry_date        *time.Time `json:"expiry_date"`
	Finance_id         *string    `json:"finance_id"`
	Ppdm_guid          string     `json:"ppdm_guid"`
	Preferred_name     *string    `json:"preferred_name"`
	Purchase_date      *time.Time `json:"purchase_date"`
	Release_date       *time.Time `json:"release_date"`
	Remark             *string    `json:"remark"`
	Role_type          *string    `json:"role_type"`
	Source             *string    `json:"source"`
	Row_changed_by     *string    `json:"row_changed_by"`
	Row_changed_date   *time.Time `json:"row_changed_date"`
	Row_created_by     *string    `json:"row_created_by"`
	Row_created_date   *time.Time `json:"row_created_date"`
	Row_effective_date *time.Time `json:"row_effective_date"`
	Row_expiry_date    *time.Time `json:"row_expiry_date"`
	Row_quality        *string    `json:"row_quality"`
}
