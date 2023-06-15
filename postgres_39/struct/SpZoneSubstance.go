package dto

import (
	"time"
)

type Sp_zone_substance struct {
	Spatial_description_id string     `json:"spatial_description_id"`
	Spatial_obs_no         int        `json:"spatial_obs_no"`
	Mineral_zone_id        string     `json:"mineral_zone_id"`
	Substance_id           string     `json:"substance_id"`
	Active_ind             *string    `json:"active_ind"`
	Base_depth             *float64   `json:"base_depth"`
	Base_depth_ouom        *string    `json:"base_depth_ouom"`
	Base_qualifier         *string    `json:"base_qualifier"`
	Base_strat_unit_id     *string    `json:"base_strat_unit_id"`
	Effective_date         *time.Time `json:"effective_date"`
	Excluded_ind           *string    `json:"excluded_ind"`
	Expiry_date            *time.Time `json:"expiry_date"`
	Included_ind           *string    `json:"included_ind"`
	Ppdm_guid              string     `json:"ppdm_guid"`
	Remark                 *string    `json:"remark"`
	Source                 *string    `json:"source"`
	Strat_name_set_id      *string    `json:"strat_name_set_id"`
	Top_depth              *float64   `json:"top_depth"`
	Top_depth_ouom         *string    `json:"top_depth_ouom"`
	Top_qualifier          *string    `json:"top_qualifier"`
	Top_strat_unit_id      *string    `json:"top_strat_unit_id"`
	Row_changed_by         *string    `json:"row_changed_by"`
	Row_changed_date       *time.Time `json:"row_changed_date"`
	Row_created_by         *string    `json:"row_created_by"`
	Row_created_date       *time.Time `json:"row_created_date"`
	Row_effective_date     *time.Time `json:"row_effective_date"`
	Row_expiry_date        *time.Time `json:"row_expiry_date"`
	Row_quality            *string    `json:"row_quality"`
}
