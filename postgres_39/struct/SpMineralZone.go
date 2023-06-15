package dto

import (
	"time"
)

type Sp_mineral_zone struct {
	Spatial_description_id   string     `json:"spatial_description_id"`
	Spatial_obs_no           int        `json:"spatial_obs_no"`
	Mineral_zone_id          string     `json:"mineral_zone_id"`
	Active_ind               *string    `json:"active_ind"`
	Base_qualifier           *string    `json:"base_qualifier"`
	Base_zone_definition_id  *string    `json:"base_zone_definition_id"`
	Deep_right_reversion_ind *string    `json:"deep_right_reversion_ind"`
	Effective_date           *time.Time `json:"effective_date"`
	Expiry_date              *time.Time `json:"expiry_date"`
	Inactivation_date        *time.Time `json:"inactivation_date"`
	Metes_and_bounds         *string    `json:"metes_and_bounds"`
	Ppdm_guid                string     `json:"ppdm_guid"`
	Remark                   *string    `json:"remark"`
	Source                   *string    `json:"source"`
	Top_qualifier            *string    `json:"top_qualifier"`
	Top_zone_definition_id   *string    `json:"top_zone_definition_id"`
	Row_changed_by           *string    `json:"row_changed_by"`
	Row_changed_date         *time.Time `json:"row_changed_date"`
	Row_created_by           *string    `json:"row_created_by"`
	Row_created_date         *time.Time `json:"row_created_date"`
	Row_effective_date       *time.Time `json:"row_effective_date"`
	Row_expiry_date          *time.Time `json:"row_expiry_date"`
	Row_quality              *string    `json:"row_quality"`
}
