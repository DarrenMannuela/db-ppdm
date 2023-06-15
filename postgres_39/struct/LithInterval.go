package dto

import (
	"time"
)

type Lith_interval struct {
	Lithology_log_id    string     `json:"lithology_log_id"`
	Lith_log_source     string     `json:"lith_log_source"`
	Depth_obs_no        int        `json:"depth_obs_no"`
	Active_ind          *string    `json:"active_ind"`
	Base_depth          *float64   `json:"base_depth"`
	Base_depth_ouom     *string    `json:"base_depth_ouom"`
	Cycle_bed           *string    `json:"cycle_bed"`
	Effective_date      *time.Time `json:"effective_date"`
	Expiry_date         *time.Time `json:"expiry_date"`
	No_data_desc        *string    `json:"no_data_desc"`
	Parent_depth_obs_no *int       `json:"parent_depth_obs_no"`
	Ppdm_guid           string     `json:"ppdm_guid"`
	Preferred_ind       *string    `json:"preferred_ind"`
	Remark              *string    `json:"remark"`
	Sample_id           *string    `json:"sample_id"`
	Sample_quality      *string    `json:"sample_quality"`
	Structure_type      *string    `json:"structure_type"`
	Thin_bed_ind        *string    `json:"thin_bed_ind"`
	Top_depth           *float64   `json:"top_depth"`
	Top_depth_ouom      *string    `json:"top_depth_ouom"`
	Row_changed_by      *string    `json:"row_changed_by"`
	Row_changed_date    *time.Time `json:"row_changed_date"`
	Row_created_by      *string    `json:"row_created_by"`
	Row_created_date    *time.Time `json:"row_created_date"`
	Row_effective_date  *time.Time `json:"row_effective_date"`
	Row_expiry_date     *time.Time `json:"row_expiry_date"`
	Row_quality         *string    `json:"row_quality"`
}
