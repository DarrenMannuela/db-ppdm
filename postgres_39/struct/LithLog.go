package dto

import (
	"time"
)

type Lith_log struct {
	Lithology_log_id    string     `json:"lithology_log_id"`
	Lith_log_source     string     `json:"lith_log_source"`
	Active_ind          *string    `json:"active_ind"`
	Base_depth          *float64   `json:"base_depth"`
	Base_depth_ouom     *string    `json:"base_depth_ouom"`
	Effective_date      *time.Time `json:"effective_date"`
	Expiry_date         *time.Time `json:"expiry_date"`
	Geologist           *string    `json:"geologist"`
	Lith_log_class      *string    `json:"lith_log_class"`
	Lith_log_type       *string    `json:"lith_log_type"`
	Logged_date         *time.Time `json:"logged_date"`
	Meas_section_id     *string    `json:"meas_section_id"`
	Meas_section_source *string    `json:"meas_section_source"`
	Ppdm_guid           string     `json:"ppdm_guid"`
	Remark              *string    `json:"remark"`
	Top_depth           *float64   `json:"top_depth"`
	Top_depth_ouom      *string    `json:"top_depth_ouom"`
	Uwi                 *string    `json:"uwi"`
	Row_changed_by      *string    `json:"row_changed_by"`
	Row_changed_date    *time.Time `json:"row_changed_date"`
	Row_created_by      *string    `json:"row_created_by"`
	Row_created_date    *time.Time `json:"row_created_date"`
	Row_effective_date  *time.Time `json:"row_effective_date"`
	Row_expiry_date     *time.Time `json:"row_expiry_date"`
	Row_quality         *string    `json:"row_quality"`
}
