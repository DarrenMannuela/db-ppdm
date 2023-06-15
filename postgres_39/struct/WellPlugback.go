package dto

import (
	"time"
)

type Well_plugback struct {
	Uwi                  string     `json:"uwi"`
	Source               string     `json:"source"`
	Plugback_obs_no      int        `json:"plugback_obs_no"`
	Active_ind           *string    `json:"active_ind"`
	Base_depth           *float64   `json:"base_depth"`
	Base_depth_ouom      *string    `json:"base_depth_ouom"`
	Base_strat_unit_id   *string    `json:"base_strat_unit_id"`
	Cement_amount        *float64   `json:"cement_amount"`
	Cement_amount_ouom   *string    `json:"cement_amount_ouom"`
	Cement_amount_uom    *string    `json:"cement_amount_uom"`
	Completion_obs_no    *int       `json:"completion_obs_no"`
	Completion_source    *string    `json:"completion_source"`
	Effective_date       *time.Time `json:"effective_date"`
	Expiry_date          *time.Time `json:"expiry_date"`
	Plugback_ba_id       *string    `json:"plugback_ba_id"`
	Plugback_date        *time.Time `json:"plugback_date"`
	Plug_type            *string    `json:"plug_type"`
	Ppdm_guid            string     `json:"ppdm_guid"`
	Remark               *string    `json:"remark"`
	Strat_name_set_id    *string    `json:"strat_name_set_id"`
	Surface_abandon_date *time.Time `json:"surface_abandon_date"`
	Top_depth            *float64   `json:"top_depth"`
	Top_depth_ouom       *string    `json:"top_depth_ouom"`
	Top_found_depth      *float64   `json:"top_found_depth"`
	Top_found_depth_ouom *string    `json:"top_found_depth_ouom"`
	Top_strat_unit_id    *string    `json:"top_strat_unit_id"`
	Row_changed_by       *string    `json:"row_changed_by"`
	Row_changed_date     *time.Time `json:"row_changed_date"`
	Row_created_by       *string    `json:"row_created_by"`
	Row_created_date     *time.Time `json:"row_created_date"`
	Row_effective_date   *time.Time `json:"row_effective_date"`
	Row_expiry_date      *time.Time `json:"row_expiry_date"`
	Row_quality          *string    `json:"row_quality"`
}
