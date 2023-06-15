package dto

import (
	"time"
)

type Well_log_curve_remark struct {
	Uwi                string     `json:"uwi"`
	Curve_id           string     `json:"curve_id"`
	Remark_obs_no      int        `json:"remark_obs_no"`
	Active_ind         *string    `json:"active_ind"`
	Effective_date     *time.Time `json:"effective_date"`
	Expiry_date        *time.Time `json:"expiry_date"`
	Made_by_ba_id      *string    `json:"made_by_ba_id"`
	Ppdm_guid          string     `json:"ppdm_guid"`
	Remark             *string    `json:"remark"`
	Remark_type        *string    `json:"remark_type"`
	Remark_use_type    *string    `json:"remark_use_type"`
	Source             *string    `json:"source"`
	Row_changed_by     *string    `json:"row_changed_by"`
	Row_changed_date   *time.Time `json:"row_changed_date"`
	Row_created_by     *string    `json:"row_created_by"`
	Row_created_date   *time.Time `json:"row_created_date"`
	Row_effective_date *time.Time `json:"row_effective_date"`
	Row_expiry_date    *time.Time `json:"row_expiry_date"`
	Row_quality        *string    `json:"row_quality"`
}
