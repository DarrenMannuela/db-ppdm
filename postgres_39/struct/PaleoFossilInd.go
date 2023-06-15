package dto

import (
	"time"
)

type Paleo_fossil_ind struct {
	Paleo_summary_id   string     `json:"paleo_summary_id"`
	Fossil_detail_id   string     `json:"fossil_detail_id"`
	Fossil_id          string     `json:"fossil_id"`
	Indicator_id       string     `json:"indicator_id"`
	Active_ind         *string    `json:"active_ind"`
	Effective_date     *time.Time `json:"effective_date"`
	Expiry_date        *time.Time `json:"expiry_date"`
	Indicator_desc     *string    `json:"indicator_desc"`
	Indicator_ind      *string    `json:"indicator_ind"`
	Indicator_type     *string    `json:"indicator_type"`
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
