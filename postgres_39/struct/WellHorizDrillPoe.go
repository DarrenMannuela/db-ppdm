package dto

import (
	"time"
)

type Well_horiz_drill_poe struct {
	Uwi                       string     `json:"uwi"`
	Source                    string     `json:"source"`
	Point_of_entry_obs_no     int        `json:"point_of_entry_obs_no"`
	Active_ind                *string    `json:"active_ind"`
	Effective_date            *time.Time `json:"effective_date"`
	Expiry_date               *time.Time `json:"expiry_date"`
	Lateral_hole_id           *string    `json:"lateral_hole_id"`
	Node_id                   *string    `json:"node_id"`
	Point_of_entry_md         *float64   `json:"point_of_entry_md"`
	Point_of_entry_md_ouom    *string    `json:"point_of_entry_md_ouom"`
	Point_of_entry_strat_unit *string    `json:"point_of_entry_strat_unit"`
	Point_of_entry_tvd        *float64   `json:"point_of_entry_tvd"`
	Point_of_entry_tvd_ouom   *string    `json:"point_of_entry_tvd_ouom"`
	Ppdm_guid                 string     `json:"ppdm_guid"`
	Remark                    *string    `json:"remark"`
	Strat_name_set_id         *string    `json:"strat_name_set_id"`
	Row_changed_by            *string    `json:"row_changed_by"`
	Row_changed_date          *time.Time `json:"row_changed_date"`
	Row_created_by            *string    `json:"row_created_by"`
	Row_created_date          *time.Time `json:"row_created_date"`
	Row_effective_date        *time.Time `json:"row_effective_date"`
	Row_expiry_date           *time.Time `json:"row_expiry_date"`
	Row_quality               *string    `json:"row_quality"`
}
