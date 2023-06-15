package dto

import (
	"time"
)

type Equipment_maint_status struct {
	Equipment_id         string     `json:"equipment_id"`
	Equip_maint_id       string     `json:"equip_maint_id"`
	Status_id            string     `json:"status_id"`
	Active_ind           *string    `json:"active_ind"`
	Effective_date       *time.Time `json:"effective_date"`
	Expiry_date          *time.Time `json:"expiry_date"`
	Maintain_status      *string    `json:"maintain_status"`
	Maintain_status_type *string    `json:"maintain_status_type"`
	Ppdm_guid            string     `json:"ppdm_guid"`
	Remark               *string    `json:"remark"`
	Source               *string    `json:"source"`
	Row_changed_by       *string    `json:"row_changed_by"`
	Row_changed_date     *time.Time `json:"row_changed_date"`
	Row_created_by       *string    `json:"row_created_by"`
	Row_created_date     *time.Time `json:"row_created_date"`
	Row_effective_date   *time.Time `json:"row_effective_date"`
	Row_expiry_date      *time.Time `json:"row_expiry_date"`
	Row_quality          *string    `json:"row_quality"`
}
