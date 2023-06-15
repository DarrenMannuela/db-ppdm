package dto

import (
	"time"
)

type Well_core_sample_rmk struct {
	Uwi                    string     `json:"uwi"`
	Source                 string     `json:"source"`
	Core_id                string     `json:"core_id"`
	Analysis_obs_no        int        `json:"analysis_obs_no"`
	Sample_num             string     `json:"sample_num"`
	Sample_analysis_obs_no int        `json:"sample_analysis_obs_no"`
	Remark_obs_no          int        `json:"remark_obs_no"`
	Active_ind             *string    `json:"active_ind"`
	Effective_date         *time.Time `json:"effective_date"`
	Expiry_date            *time.Time `json:"expiry_date"`
	Ppdm_guid              string     `json:"ppdm_guid"`
	Remark                 *string    `json:"remark"`
	Remark_source          *string    `json:"remark_source"`
	Row_changed_by         *string    `json:"row_changed_by"`
	Row_changed_date       *time.Time `json:"row_changed_date"`
	Row_created_by         *string    `json:"row_created_by"`
	Row_created_date       *time.Time `json:"row_created_date"`
	Row_effective_date     *time.Time `json:"row_effective_date"`
	Row_expiry_date        *time.Time `json:"row_expiry_date"`
	Row_quality            *string    `json:"row_quality"`
}
