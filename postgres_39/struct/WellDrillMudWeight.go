package dto

import (
	"time"
)

type Well_drill_mud_weight struct {
	Uwi                string     `json:"uwi"`
	Source             string     `json:"source"`
	Depth_obs_no       int        `json:"depth_obs_no"`
	Media_obs_no       int        `json:"media_obs_no"`
	Active_ind         *string    `json:"active_ind"`
	Effective_date     *time.Time `json:"effective_date"`
	End_date           *time.Time `json:"end_date"`
	Expiry_date        *time.Time `json:"expiry_date"`
	Mud_depth          *float64   `json:"mud_depth"`
	Mud_depth_ouom     *string    `json:"mud_depth_ouom"`
	Mud_weight         *float64   `json:"mud_weight"`
	Mud_weight_ouom    *string    `json:"mud_weight_ouom"`
	Mud_weight_uom     *string    `json:"mud_weight_uom"`
	Ppdm_guid          string     `json:"ppdm_guid"`
	Remark             *string    `json:"remark"`
	Reported_tvd       *float64   `json:"reported_tvd"`
	Reported_tvd_ouom  *string    `json:"reported_tvd_ouom"`
	Start_date         *time.Time `json:"start_date"`
	Row_changed_by     *string    `json:"row_changed_by"`
	Row_changed_date   *time.Time `json:"row_changed_date"`
	Row_created_by     *string    `json:"row_created_by"`
	Row_created_date   *time.Time `json:"row_created_date"`
	Row_effective_date *time.Time `json:"row_effective_date"`
	Row_expiry_date    *time.Time `json:"row_expiry_date"`
	Row_quality        *string    `json:"row_quality"`
}
