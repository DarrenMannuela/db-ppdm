package dto

import (
	"time"
)

type Well_perf_remark struct {
	Uwi                string     `json:"uwi"`
	Perforation_source string     `json:"perforation_source"`
	Perforation_obs_no int        `json:"perforation_obs_no"`
	Remark_obs_no      int        `json:"remark_obs_no"`
	Active_ind         *string    `json:"active_ind"`
	Effective_date     *time.Time `json:"effective_date"`
	Expiry_date        *time.Time `json:"expiry_date"`
	Ppdm_guid          string     `json:"ppdm_guid"`
	Remark             *string    `json:"remark"`
	Remark_ba_id       *string    `json:"remark_ba_id"`
	Remark_date        *time.Time `json:"remark_date"`
	Source             *string    `json:"source"`
	Row_changed_by     *string    `json:"row_changed_by"`
	Row_changed_date   *time.Time `json:"row_changed_date"`
	Row_created_by     *string    `json:"row_created_by"`
	Row_created_date   *time.Time `json:"row_created_date"`
	Row_effective_date *time.Time `json:"row_effective_date"`
	Row_expiry_date    *time.Time `json:"row_expiry_date"`
	Row_quality        *string    `json:"row_quality"`
}
