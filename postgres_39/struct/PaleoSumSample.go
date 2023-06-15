package dto

import (
	"time"
)

type Paleo_sum_sample struct {
	Paleo_summary_id   string     `json:"paleo_summary_id"`
	Lith_sample_id     string     `json:"lith_sample_id"`
	Active_ind         *string    `json:"active_ind"`
	Diversity_count    *int       `json:"diversity_count"`
	Effective_date     *time.Time `json:"effective_date"`
	Expiry_date        *time.Time `json:"expiry_date"`
	First_sample_ind   *string    `json:"first_sample_ind"`
	Last_sample_ind    *string    `json:"last_sample_ind"`
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
