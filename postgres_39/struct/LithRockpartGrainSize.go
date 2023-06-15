package dto

import (
	"time"
)

type Lith_rockpart_grain_size struct {
	Lithology_log_id   string     `json:"lithology_log_id"`
	Lith_log_source    string     `json:"lith_log_source"`
	Depth_obs_no       int        `json:"depth_obs_no"`
	Scale_scheme       string     `json:"scale_scheme"`
	Rock_type          string     `json:"rock_type"`
	Rock_type_obs_no   int        `json:"rock_type_obs_no"`
	Rockpart_name      string     `json:"rockpart_name"`
	Grain_size         string     `json:"grain_size"`
	Active_ind         *string    `json:"active_ind"`
	Distribution_ind   *string    `json:"distribution_ind"`
	Effective_date     *time.Time `json:"effective_date"`
	Expiry_date        *time.Time `json:"expiry_date"`
	Ppdm_guid          string     `json:"ppdm_guid"`
	Remark             *string    `json:"remark"`
	Size_type_ind      *string    `json:"size_type_ind"`
	Row_changed_by     *string    `json:"row_changed_by"`
	Row_changed_date   *time.Time `json:"row_changed_date"`
	Row_created_by     *string    `json:"row_created_by"`
	Row_created_date   *time.Time `json:"row_created_date"`
	Row_effective_date *time.Time `json:"row_effective_date"`
	Row_expiry_date    *time.Time `json:"row_expiry_date"`
	Row_quality        *string    `json:"row_quality"`
}
