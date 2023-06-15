package dto

import (
	"time"
)

type Well_log_curve_value struct {
	Uwi                string     `json:"uwi"`
	Curve_id           string     `json:"curve_id"`
	Sample_id          string     `json:"sample_id"`
	Active_ind         *string    `json:"active_ind"`
	Effective_date     *time.Time `json:"effective_date"`
	Expiry_date        *time.Time `json:"expiry_date"`
	Index_type         *string    `json:"index_type"`
	Index_value        *float64   `json:"index_value"`
	Index_value_uom    *string    `json:"index_value_uom"`
	Measured_value     *float64   `json:"measured_value"`
	Measured_value_uom *string    `json:"measured_value_uom"`
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
