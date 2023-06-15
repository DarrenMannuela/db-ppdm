package dto

import (
	"time"
)

type Lith_diagenesis struct {
	Lithology_log_id         string     `json:"lithology_log_id"`
	Lith_log_source          string     `json:"lith_log_source"`
	Depth_obs_no             int        `json:"depth_obs_no"`
	Rock_type                string     `json:"rock_type"`
	Rock_type_obs_no         int        `json:"rock_type_obs_no"`
	Diagenesis_type          string     `json:"diagenesis_type"`
	Active_ind               *string    `json:"active_ind"`
	Diagenesis_percent       *float64   `json:"diagenesis_percent"`
	Diagenesis_rel_abundance *string    `json:"diagenesis_rel_abundance"`
	Effective_date           *time.Time `json:"effective_date"`
	Expiry_date              *time.Time `json:"expiry_date"`
	Ppdm_guid                string     `json:"ppdm_guid"`
	Remark                   *string    `json:"remark"`
	Row_changed_by           *string    `json:"row_changed_by"`
	Row_changed_date         *time.Time `json:"row_changed_date"`
	Row_created_by           *string    `json:"row_created_by"`
	Row_created_date         *time.Time `json:"row_created_date"`
	Row_effective_date       *time.Time `json:"row_effective_date"`
	Row_expiry_date          *time.Time `json:"row_expiry_date"`
	Row_quality              *string    `json:"row_quality"`
}
