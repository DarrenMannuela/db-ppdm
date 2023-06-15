package dto

import (
	"time"
)

type Well_air_drill_period struct {
	Uwi                string     `json:"uwi"`
	Interval_source    string     `json:"interval_source"`
	Air_drill_obs_no   int        `json:"air_drill_obs_no"`
	Depth_obs_no       int        `json:"depth_obs_no"`
	Period_obs_no      int        `json:"period_obs_no"`
	Active_ind         *string    `json:"active_ind"`
	Base_depth         *float64   `json:"base_depth"`
	Base_depth_ouom    *string    `json:"base_depth_ouom"`
	Effective_date     *time.Time `json:"effective_date"`
	Expiry_date        *time.Time `json:"expiry_date"`
	Ppdm_guid          string     `json:"ppdm_guid"`
	Remark             *string    `json:"remark"`
	Source             *string    `json:"source"`
	Top_depth          *float64   `json:"top_depth"`
	Top_depth_ouom     *string    `json:"top_depth_ouom"`
	Row_changed_by     *string    `json:"row_changed_by"`
	Row_changed_date   *time.Time `json:"row_changed_date"`
	Row_created_by     *string    `json:"row_created_by"`
	Row_created_date   *time.Time `json:"row_created_date"`
	Row_effective_date *time.Time `json:"row_effective_date"`
	Row_expiry_date    *time.Time `json:"row_expiry_date"`
	Row_quality        *string    `json:"row_quality"`
}
