package dto

import (
	"time"
)

type Paleo_fossil_obs struct {
	Paleo_summary_id   string     `json:"paleo_summary_id"`
	Fossil_detail_id   string     `json:"fossil_detail_id"`
	Fossil_id          string     `json:"fossil_id"`
	Observation_obs_no int        `json:"observation_obs_no"`
	Active_ind         *string    `json:"active_ind"`
	Effective_date     *time.Time `json:"effective_date"`
	Expiry_date        *time.Time `json:"expiry_date"`
	Fossil_obs_type    *string    `json:"fossil_obs_type"`
	Observation        *string    `json:"observation"`
	Ppdm_guid          string     `json:"ppdm_guid"`
	Remark             *string    `json:"remark"`
	Source             *string    `json:"source"`
	Row_changed_by     *string    `json:"row_changed_by"`
	Row_changed_date   *time.Time `json:"row_changed_date"`
	Row_created_by     *string    `json:"row_created_by"`
	Row_created_date   *time.Time `json:"row_created_date"`
	Row_effective_date *time.Time `json:"row_effective_date"`
	Row_expiry_date    *time.Time `json:"row_expiry_date"`
	Row_quality        *string    `json:"row_quality"`
}
