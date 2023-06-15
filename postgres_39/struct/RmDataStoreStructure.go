package dto

import (
	"time"
)

type Rm_data_store_structure struct {
	Store_id                 string     `json:"store_id"`
	Structure_obs_no         int        `json:"structure_obs_no"`
	Active_ind               *string    `json:"active_ind"`
	Dim_height               *float64   `json:"dim_height"`
	Dim_height_ouom          *string    `json:"dim_height_ouom"`
	Dim_height_uom           *string    `json:"dim_height_uom"`
	Dim_length               *float64   `json:"dim_length"`
	Dim_length_ouom          *string    `json:"dim_length_ouom"`
	Dim_length_uom           *string    `json:"dim_length_uom"`
	Dim_weight               *float64   `json:"dim_weight"`
	Dim_weight_ouom          *string    `json:"dim_weight_ouom"`
	Dim_weight_uom           *string    `json:"dim_weight_uom"`
	Dim_width                *float64   `json:"dim_width"`
	Dim_width_ouom           *string    `json:"dim_width_ouom"`
	Dim_width_uom            *string    `json:"dim_width_uom"`
	Effective_date           *time.Time `json:"effective_date"`
	Expiry_date              *time.Time `json:"expiry_date"`
	Lower_contained_store_id *string    `json:"lower_contained_store_id"`
	Ppdm_guid                string     `json:"ppdm_guid"`
	Remark                   *string    `json:"remark"`
	Source                   *string    `json:"source"`
	Total_capacity           *int       `json:"total_capacity"`
	Upper_contained_store_id *string    `json:"upper_contained_store_id"`
	Used_capacity            *int       `json:"used_capacity"`
	Used_capacity_date       *time.Time `json:"used_capacity_date"`
	Row_changed_by           *string    `json:"row_changed_by"`
	Row_changed_date         *time.Time `json:"row_changed_date"`
	Row_created_by           *string    `json:"row_created_by"`
	Row_created_date         *time.Time `json:"row_created_date"`
	Row_effective_date       *time.Time `json:"row_effective_date"`
	Row_expiry_date          *time.Time `json:"row_expiry_date"`
	Row_quality              *string    `json:"row_quality"`
}
