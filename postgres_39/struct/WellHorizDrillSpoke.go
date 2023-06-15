package dto

import (
	"time"
)

type Well_horiz_drill_spoke struct {
	Uwi                  string     `json:"uwi"`
	Source               string     `json:"source"`
	Kickoff_point_obs_no int        `json:"kickoff_point_obs_no"`
	Spoke_obs_no         int        `json:"spoke_obs_no"`
	Active_ind           *string    `json:"active_ind"`
	Effective_date       *time.Time `json:"effective_date"`
	Expiry_date          *time.Time `json:"expiry_date"`
	Lateral_hole_id      *string    `json:"lateral_hole_id"`
	Node_id              *string    `json:"node_id"`
	Ppdm_guid            string     `json:"ppdm_guid"`
	Remark               *string    `json:"remark"`
	Spoke_length         *float64   `json:"spoke_length"`
	Spoke_length_ouom    *string    `json:"spoke_length_ouom"`
	Spoke_md             *float64   `json:"spoke_md"`
	Spoke_md_ouom        *string    `json:"spoke_md_ouom"`
	Spoke_tvd            *float64   `json:"spoke_tvd"`
	Spoke_tvd_ouom       *string    `json:"spoke_tvd_ouom"`
	Row_changed_by       *string    `json:"row_changed_by"`
	Row_changed_date     *time.Time `json:"row_changed_date"`
	Row_created_by       *string    `json:"row_created_by"`
	Row_created_date     *time.Time `json:"row_created_date"`
	Row_effective_date   *time.Time `json:"row_effective_date"`
	Row_expiry_date      *time.Time `json:"row_expiry_date"`
	Row_quality          *string    `json:"row_quality"`
}
