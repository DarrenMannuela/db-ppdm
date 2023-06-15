package dto

import (
	"time"
)

type Ecozone struct {
	Ecozone_id         string     `json:"ecozone_id"`
	Active_ind         *string    `json:"active_ind"`
	Base_depth         *float64   `json:"base_depth"`
	Depth_ouom         *string    `json:"depth_ouom"`
	Ecozone_type       *string    `json:"ecozone_type"`
	Effective_date     *time.Time `json:"effective_date"`
	Expiry_date        *time.Time `json:"expiry_date"`
	Ppdm_guid          string     `json:"ppdm_guid"`
	Preferred_name     *string    `json:"preferred_name"`
	Remark             *string    `json:"remark"`
	Source             *string    `json:"source"`
	Top_depth          *float64   `json:"top_depth"`
	Row_changed_by     *string    `json:"row_changed_by"`
	Row_changed_date   *time.Time `json:"row_changed_date"`
	Row_created_by     *string    `json:"row_created_by"`
	Row_created_date   *time.Time `json:"row_created_date"`
	Row_effective_date *time.Time `json:"row_effective_date"`
	Row_expiry_date    *time.Time `json:"row_expiry_date"`
	Row_quality        *string    `json:"row_quality"`
}
