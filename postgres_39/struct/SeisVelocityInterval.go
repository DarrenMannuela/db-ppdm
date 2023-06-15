package dto

import (
	"time"
)

type Seis_velocity_interval struct {
	Interval_id          string     `json:"interval_id"`
	Active_ind           *string    `json:"active_ind"`
	Base_depth           *float64   `json:"base_depth"`
	Base_depth_ouom      *string    `json:"base_depth_ouom"`
	Base_strat_unit_id   *string    `json:"base_strat_unit_id"`
	Compute_method       *string    `json:"compute_method"`
	Effective_date       *time.Time `json:"effective_date"`
	Expiry_date          *time.Time `json:"expiry_date"`
	Last_update          *time.Time `json:"last_update"`
	Ppdm_guid            string     `json:"ppdm_guid"`
	Remark               *string    `json:"remark"`
	Reported_tvd         *float64   `json:"reported_tvd"`
	Reported_tvd_ouom    *string    `json:"reported_tvd_ouom"`
	Seis_set_id          *string    `json:"seis_set_id"`
	Seis_set_subtype     *string    `json:"seis_set_subtype"`
	Seis_time_datum      *float64   `json:"seis_time_datum"`
	Seis_time_datum_ouom *string    `json:"seis_time_datum_ouom"`
	Source               *string    `json:"source"`
	Strat_name_set_id    *string    `json:"strat_name_set_id"`
	Top_depth            *float64   `json:"top_depth"`
	Top_depth_ouom       *string    `json:"top_depth_ouom"`
	Top_strat_unit_id    *string    `json:"top_strat_unit_id"`
	Uwi                  *string    `json:"uwi"`
	Velocity_quality     *string    `json:"velocity_quality"`
	Velocity_type        *string    `json:"velocity_type"`
	Velocity_value       *float64   `json:"velocity_value"`
	Velocity_value_ouom  *string    `json:"velocity_value_ouom"`
	Row_changed_by       *string    `json:"row_changed_by"`
	Row_changed_date     *time.Time `json:"row_changed_date"`
	Row_created_by       *string    `json:"row_created_by"`
	Row_created_date     *time.Time `json:"row_created_date"`
	Row_effective_date   *time.Time `json:"row_effective_date"`
	Row_expiry_date      *time.Time `json:"row_expiry_date"`
	Row_quality          *string    `json:"row_quality"`
}
