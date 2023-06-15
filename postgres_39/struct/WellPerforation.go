package dto

import (
	"time"
)

type Well_perforation struct {
	Uwi                       string     `json:"uwi"`
	Source                    string     `json:"source"`
	Perforation_obs_no        int        `json:"perforation_obs_no"`
	Active_ind                *string    `json:"active_ind"`
	Base_depth                *float64   `json:"base_depth"`
	Base_depth_ouom           *string    `json:"base_depth_ouom"`
	Base_strat_unit_id        *string    `json:"base_strat_unit_id"`
	Completion_obs_no         *int       `json:"completion_obs_no"`
	Completion_source         *string    `json:"completion_source"`
	Completion_status         *string    `json:"completion_status"`
	Completion_status_type    *string    `json:"completion_status_type"`
	Current_status_date       *time.Time `json:"current_status_date"`
	Effective_date            *time.Time `json:"effective_date"`
	Expiry_date               *time.Time `json:"expiry_date"`
	Perforation_angle         *float64   `json:"perforation_angle"`
	Perforation_ba_id         *string    `json:"perforation_ba_id"`
	Perforation_count         *int       `json:"perforation_count"`
	Perforation_date          *time.Time `json:"perforation_date"`
	Perforation_density       *float64   `json:"perforation_density"`
	Perforation_diameter      *float64   `json:"perforation_diameter"`
	Perforation_diameter_ouom *string    `json:"perforation_diameter_ouom"`
	Perforation_method        *string    `json:"perforation_method"`
	Perforation_per_uom       *string    `json:"perforation_per_uom"`
	Perforation_phase         *string    `json:"perforation_phase"`
	Perforation_type          *string    `json:"perforation_type"`
	Ppdm_guid                 string     `json:"ppdm_guid"`
	Remark                    *string    `json:"remark"`
	Strat_name_set_id         *string    `json:"strat_name_set_id"`
	Top_depth                 *float64   `json:"top_depth"`
	Top_depth_ouom            *string    `json:"top_depth_ouom"`
	Top_strat_unit_id         *string    `json:"top_strat_unit_id"`
	Row_changed_by            *string    `json:"row_changed_by"`
	Row_changed_date          *time.Time `json:"row_changed_date"`
	Row_created_by            *string    `json:"row_created_by"`
	Row_created_date          *time.Time `json:"row_created_date"`
	Row_effective_date        *time.Time `json:"row_effective_date"`
	Row_expiry_date           *time.Time `json:"row_expiry_date"`
	Row_quality               *string    `json:"row_quality"`
}
