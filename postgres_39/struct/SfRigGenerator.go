package dto

import (
	"time"
)

type Sf_rig_generator struct {
	Sf_subtype         string     `json:"sf_subtype"`
	Rig_id             string     `json:"rig_id"`
	Generator_id       string     `json:"generator_id"`
	Active_ind         *string    `json:"active_ind"`
	Catalogue_equip_id *string    `json:"catalogue_equip_id"`
	Effective_date     *time.Time `json:"effective_date"`
	Equipment_id       *string    `json:"equipment_id"`
	Expiry_date        *time.Time `json:"expiry_date"`
	Generator_count    *int       `json:"generator_count"`
	Generator_type     *string    `json:"generator_type"`
	Input_type         *string    `json:"input_type"`
	Install_date       *time.Time `json:"install_date"`
	Power              *float64   `json:"power"`
	Power_ouom         *string    `json:"power_ouom"`
	Power_rating       *float64   `json:"power_rating"`
	Power_rating_ouom  *string    `json:"power_rating_ouom"`
	Ppdm_guid          string     `json:"ppdm_guid"`
	Reference_num      *string    `json:"reference_num"`
	Remark             *string    `json:"remark"`
	Remove_date        *time.Time `json:"remove_date"`
	Source             *string    `json:"source"`
	Row_changed_by     *string    `json:"row_changed_by"`
	Row_changed_date   *time.Time `json:"row_changed_date"`
	Row_created_by     *string    `json:"row_created_by"`
	Row_created_date   *time.Time `json:"row_created_date"`
	Row_effective_date *time.Time `json:"row_effective_date"`
	Row_expiry_date    *time.Time `json:"row_expiry_date"`
	Row_quality        *string    `json:"row_quality"`
}
