package dto

import (
	"time"
)

type Equipment_xref struct {
	Equipment_id          string     `json:"equipment_id"`
	Equipment_part_id     string     `json:"equipment_part_id"`
	Equipment_xref_obs_no int        `json:"equipment_xref_obs_no"`
	Active_ind            *string    `json:"active_ind"`
	Commission_date       *time.Time `json:"commission_date"`
	Decommission_date     *time.Time `json:"decommission_date"`
	Description           *string    `json:"description"`
	Effective_date        *time.Time `json:"effective_date"`
	Equip_xref_type       *string    `json:"equip_xref_type"`
	Expiry_date           *time.Time `json:"expiry_date"`
	Interchangable_ind    *string    `json:"interchangable_ind"`
	Ppdm_guid             string     `json:"ppdm_guid"`
	Remark                *string    `json:"remark"`
	Source                *string    `json:"source"`
	Row_changed_by        *string    `json:"row_changed_by"`
	Row_changed_date      *time.Time `json:"row_changed_date"`
	Row_created_by        *string    `json:"row_created_by"`
	Row_created_date      *time.Time `json:"row_created_date"`
	Row_effective_date    *time.Time `json:"row_effective_date"`
	Row_expiry_date       *time.Time `json:"row_expiry_date"`
	Row_quality           *string    `json:"row_quality"`
}
