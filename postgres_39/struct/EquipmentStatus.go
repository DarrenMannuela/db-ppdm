package dto

import (
	"time"
)

type Equipment_status struct {
	Equipment_id       string     `json:"equipment_id"`
	Equip_status_id    string     `json:"equip_status_id"`
	Active_ind         *string    `json:"active_ind"`
	Effective_date     *time.Time `json:"effective_date"`
	End_time           *time.Time `json:"end_time"`
	Equip_status       *string    `json:"equip_status"`
	Equip_status_type  *string    `json:"equip_status_type"`
	Expiry_date        *time.Time `json:"expiry_date"`
	Percent_capability *float64   `json:"percent_capability"`
	Ppdm_guid          string     `json:"ppdm_guid"`
	Preferred_ind      *string    `json:"preferred_ind"`
	Remark             *string    `json:"remark"`
	Source             *string    `json:"source"`
	Start_time         *time.Time `json:"start_time"`
	Timezone           *string    `json:"timezone"`
	Row_changed_by     *string    `json:"row_changed_by"`
	Row_changed_date   *time.Time `json:"row_changed_date"`
	Row_created_by     *string    `json:"row_created_by"`
	Row_created_date   *time.Time `json:"row_created_date"`
	Row_effective_date *time.Time `json:"row_effective_date"`
	Row_expiry_date    *time.Time `json:"row_expiry_date"`
	Row_quality        *string    `json:"row_quality"`
}
