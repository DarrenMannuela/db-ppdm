package dto

import (
	"time"
)

type Well_drill_remark struct {
	Uwi                 string     `json:"uwi"`
	Remark_seq_no       int        `json:"remark_seq_no"`
	Active_ind          *string    `json:"active_ind"`
	Base_depth          *float64   `json:"base_depth"`
	Base_depth_ouom     *string    `json:"base_depth_ouom"`
	Effective_date      *time.Time `json:"effective_date"`
	End_period_obs_no   *int       `json:"end_period_obs_no"`
	Expiry_date         *time.Time `json:"expiry_date"`
	Ppdm_guid           string     `json:"ppdm_guid"`
	Remark              *string    `json:"remark"`
	Remark_by_ba_id     *string    `json:"remark_by_ba_id"`
	Remark_date         *time.Time `json:"remark_date"`
	Remark_type         *string    `json:"remark_type"`
	Source              *string    `json:"source"`
	Start_period_obs_no *int       `json:"start_period_obs_no"`
	Strat_name_set_id   *string    `json:"strat_name_set_id"`
	Strat_unit_id       *string    `json:"strat_unit_id"`
	Top_depth           *float64   `json:"top_depth"`
	Top_depth_ouom      *string    `json:"top_depth_ouom"`
	Row_changed_by      *string    `json:"row_changed_by"`
	Row_changed_date    *time.Time `json:"row_changed_date"`
	Row_created_by      *string    `json:"row_created_by"`
	Row_created_date    *time.Time `json:"row_created_date"`
	Row_effective_date  *time.Time `json:"row_effective_date"`
	Row_expiry_date     *time.Time `json:"row_expiry_date"`
	Row_quality         *string    `json:"row_quality"`
}
