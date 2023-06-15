package dto

import (
	"time"
)

type Well_node_strat_unit struct {
	Uwi                  string     `json:"uwi"`
	Strat_name_set_id    string     `json:"strat_name_set_id"`
	Strat_unit_id        string     `json:"strat_unit_id"`
	Interp_id            string     `json:"interp_id"`
	Active_ind           *string    `json:"active_ind"`
	Base_offset_tvd      *float64   `json:"base_offset_tvd"`
	Base_offset_tvd_ouom *string    `json:"base_offset_tvd_ouom"`
	Effective_date       *time.Time `json:"effective_date"`
	Expiry_date          *time.Time `json:"expiry_date"`
	Interpreter_ba_id    *string    `json:"interpreter_ba_id"`
	Location_type        *string    `json:"location_type"`
	Node_id              *string    `json:"node_id"`
	Node_position        *string    `json:"node_position"`
	Pick_date            *time.Time `json:"pick_date"`
	Pick_depth           *float64   `json:"pick_depth"`
	Pick_depth_ouom      *string    `json:"pick_depth_ouom"`
	Pick_location        *string    `json:"pick_location"`
	Pick_method          *string    `json:"pick_method"`
	Pick_method_desc     *string    `json:"pick_method_desc"`
	Pick_quality         *string    `json:"pick_quality"`
	Ppdm_guid            string     `json:"ppdm_guid"`
	Preferred_ind        *string    `json:"preferred_ind"`
	Remark               *string    `json:"remark"`
	Source               *string    `json:"source"`
	Top_offset_tvd       *float64   `json:"top_offset_tvd"`
	Top_offset_tvd_ouom  *string    `json:"top_offset_tvd_ouom"`
	Row_changed_by       *string    `json:"row_changed_by"`
	Row_changed_date     *time.Time `json:"row_changed_date"`
	Row_created_by       *string    `json:"row_created_by"`
	Row_created_date     *time.Time `json:"row_created_date"`
	Row_effective_date   *time.Time `json:"row_effective_date"`
	Row_expiry_date      *time.Time `json:"row_expiry_date"`
	Row_quality          *string    `json:"row_quality"`
}
