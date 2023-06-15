package dto

import (
	"time"
)

type Sf_rig_overhead_equip struct {
	Sf_subtype           string     `json:"sf_subtype"`
	Rig_id               string     `json:"rig_id"`
	Equip_id             string     `json:"equip_id"`
	Active_ind           *string    `json:"active_ind"`
	Capacity             *float64   `json:"capacity"`
	Capacity_ouom        *string    `json:"capacity_ouom"`
	Capacity_type        *string    `json:"capacity_type"`
	Capacity_uom         *string    `json:"capacity_uom"`
	Catalogue_equip_id   *string    `json:"catalogue_equip_id"`
	Effective_date       *time.Time `json:"effective_date"`
	Equipment_id         *string    `json:"equipment_id"`
	Expiry_date          *time.Time `json:"expiry_date"`
	Input_type           *string    `json:"input_type"`
	Inside_diameter      *float64   `json:"inside_diameter"`
	Inside_diameter_ouom *string    `json:"inside_diameter_ouom"`
	Install_date         *time.Time `json:"install_date"`
	Overhead_count       *int       `json:"overhead_count"`
	Overhead_equip_type  *string    `json:"overhead_equip_type"`
	Ppdm_guid            string     `json:"ppdm_guid"`
	Remark               *string    `json:"remark"`
	Remove_date          *time.Time `json:"remove_date"`
	Source               *string    `json:"source"`
	Row_changed_by       *string    `json:"row_changed_by"`
	Row_changed_date     *time.Time `json:"row_changed_date"`
	Row_created_by       *string    `json:"row_created_by"`
	Row_created_date     *time.Time `json:"row_created_date"`
	Row_effective_date   *time.Time `json:"row_effective_date"`
	Row_expiry_date      *time.Time `json:"row_expiry_date"`
	Row_quality          *string    `json:"row_quality"`
}
