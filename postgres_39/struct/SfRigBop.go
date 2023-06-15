package dto

import (
	"time"
)

type Sf_rig_bop struct {
	Sf_subtype               string     `json:"sf_subtype"`
	Rig_id                   string     `json:"rig_id"`
	Bop_id                   string     `json:"bop_id"`
	Active_ind               *string    `json:"active_ind"`
	Annular_count            *int       `json:"annular_count"`
	Bop_count                *int       `json:"bop_count"`
	Bop_diameter             *float64   `json:"bop_diameter"`
	Bop_diameter_ouom        *string    `json:"bop_diameter_ouom"`
	Bop_nace_certified_ind   *string    `json:"bop_nace_certified_ind"`
	Bop_position_desc        *string    `json:"bop_position_desc"`
	Bop_pressure_rating      *float64   `json:"bop_pressure_rating"`
	Bop_pressure_rating_ouom *string    `json:"bop_pressure_rating_ouom"`
	Bop_type                 *string    `json:"bop_type"`
	Capacity                 *float64   `json:"capacity"`
	Capacity_ouom            *string    `json:"capacity_ouom"`
	Catalogue_equip_id       *string    `json:"catalogue_equip_id"`
	Double_count             *int       `json:"double_count"`
	Effective_date           *time.Time `json:"effective_date"`
	Equipment_id             *string    `json:"equipment_id"`
	Expiry_date              *time.Time `json:"expiry_date"`
	Input_type               *string    `json:"input_type"`
	Install_date             *time.Time `json:"install_date"`
	Ppdm_guid                string     `json:"ppdm_guid"`
	Reference_num            *string    `json:"reference_num"`
	Remark                   *string    `json:"remark"`
	Remove_date              *time.Time `json:"remove_date"`
	Single_count             *int       `json:"single_count"`
	Source                   *string    `json:"source"`
	Row_changed_by           *string    `json:"row_changed_by"`
	Row_changed_date         *time.Time `json:"row_changed_date"`
	Row_created_by           *string    `json:"row_created_by"`
	Row_created_date         *time.Time `json:"row_created_date"`
	Row_effective_date       *time.Time `json:"row_effective_date"`
	Row_expiry_date          *time.Time `json:"row_expiry_date"`
	Row_quality              *string    `json:"row_quality"`
}
