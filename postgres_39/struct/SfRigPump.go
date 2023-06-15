package dto

import (
	"time"
)

type Sf_rig_pump struct {
	Sf_subtype         string     `json:"sf_subtype"`
	Rig_id             string     `json:"rig_id"`
	Pump_id            string     `json:"pump_id"`
	Active_ind         *string    `json:"active_ind"`
	Catalogue_equip_id *string    `json:"catalogue_equip_id"`
	Effective_date     *time.Time `json:"effective_date"`
	Equipment_id       *string    `json:"equipment_id"`
	Expiry_date        *time.Time `json:"expiry_date"`
	Input_type         *string    `json:"input_type"`
	Install_date       *time.Time `json:"install_date"`
	Ppdm_guid          string     `json:"ppdm_guid"`
	Pump_count         *int       `json:"pump_count"`
	Pump_function      *string    `json:"pump_function"`
	Pump_hp            *float64   `json:"pump_hp"`
	Pump_hp_ouom       *string    `json:"pump_hp_ouom"`
	Pump_rating        *float64   `json:"pump_rating"`
	Pump_rating_ouom   *string    `json:"pump_rating_ouom"`
	Pump_type          *string    `json:"pump_type"`
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
