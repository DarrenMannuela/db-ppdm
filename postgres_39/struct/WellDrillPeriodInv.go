package dto

import (
	"time"
)

type Well_drill_period_inv struct {
	Uwi                     string     `json:"uwi"`
	Period_obs_no           int        `json:"period_obs_no"`
	Inventory_seq_no        int        `json:"inventory_seq_no"`
	Active_ind              *string    `json:"active_ind"`
	Cat_additive_id         *string    `json:"cat_additive_id"`
	Cat_equip_id            *string    `json:"cat_equip_id"`
	Effective_date          *time.Time `json:"effective_date"`
	Equipment_id            *string    `json:"equipment_id"`
	Expiry_date             *time.Time `json:"expiry_date"`
	Inventory_material_type *string    `json:"inventory_material_type"`
	Ppdm_guid               string     `json:"ppdm_guid"`
	Quantity_on_hand        *float64   `json:"quantity_on_hand"`
	Quantity_on_hand_ouom   *string    `json:"quantity_on_hand_ouom"`
	Quantity_on_hand_uom    *string    `json:"quantity_on_hand_uom"`
	Quantity_ordered        *float64   `json:"quantity_ordered"`
	Quantity_ordered_ouom   *string    `json:"quantity_ordered_ouom"`
	Quantity_ordered_uom    *string    `json:"quantity_ordered_uom"`
	Remark                  *string    `json:"remark"`
	Report_time             *time.Time `json:"report_time"`
	Report_timezone         *string    `json:"report_timezone"`
	Report_time_type        *string    `json:"report_time_type"`
	Source                  *string    `json:"source"`
	Row_changed_by          *string    `json:"row_changed_by"`
	Row_changed_date        *time.Time `json:"row_changed_date"`
	Row_created_by          *string    `json:"row_created_by"`
	Row_created_date        *time.Time `json:"row_created_date"`
	Row_effective_date      *time.Time `json:"row_effective_date"`
	Row_expiry_date         *time.Time `json:"row_expiry_date"`
	Row_quality             *string    `json:"row_quality"`
}
