package dto

import (
	"time"
)

type Sf_rig_shaker struct {
	Sf_subtype         string     `json:"sf_subtype"`
	Rig_id             string     `json:"rig_id"`
	Shaker_id          string     `json:"shaker_id"`
	Active_ind         *string    `json:"active_ind"`
	Capacity           *float64   `json:"capacity"`
	Capacity_ouom      *string    `json:"capacity_ouom"`
	Catalogue_equip_id *string    `json:"catalogue_equip_id"`
	Effective_date     *time.Time `json:"effective_date"`
	Equipment_id       *string    `json:"equipment_id"`
	Expiry_date        *time.Time `json:"expiry_date"`
	Input_type         *string    `json:"input_type"`
	Install_date       *time.Time `json:"install_date"`
	Position_desc      *string    `json:"position_desc"`
	Ppdm_guid          string     `json:"ppdm_guid"`
	Reference_num      *string    `json:"reference_num"`
	Remark             *string    `json:"remark"`
	Remove_date        *time.Time `json:"remove_date"`
	Shaker_count       *int       `json:"shaker_count"`
	Shaker_size        *float64   `json:"shaker_size"`
	Shaker_size_desc   *string    `json:"shaker_size_desc"`
	Shaker_size_ouom   *string    `json:"shaker_size_ouom"`
	Source             *string    `json:"source"`
	Row_changed_by     *string    `json:"row_changed_by"`
	Row_changed_date   *time.Time `json:"row_changed_date"`
	Row_created_by     *string    `json:"row_created_by"`
	Row_created_date   *time.Time `json:"row_created_date"`
	Row_effective_date *time.Time `json:"row_effective_date"`
	Row_expiry_date    *time.Time `json:"row_expiry_date"`
	Row_quality        *string    `json:"row_quality"`
}
