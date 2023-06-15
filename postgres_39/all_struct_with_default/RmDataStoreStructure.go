package dto

type Rm_data_store_structure struct {
	Store_id                 string   `json:"store_id" default:"source"`
	Structure_obs_no         int      `json:"structure_obs_no" default:"1"`
	Active_ind               *string  `json:"active_ind" default:""`
	Dim_height               *float64 `json:"dim_height" default:""`
	Dim_height_ouom          *string  `json:"dim_height_ouom" default:""`
	Dim_height_uom           *string  `json:"dim_height_uom" default:""`
	Dim_length               *float64 `json:"dim_length" default:""`
	Dim_length_ouom          *string  `json:"dim_length_ouom" default:""`
	Dim_length_uom           *string  `json:"dim_length_uom" default:""`
	Dim_weight               *float64 `json:"dim_weight" default:""`
	Dim_weight_ouom          *string  `json:"dim_weight_ouom" default:""`
	Dim_weight_uom           *string  `json:"dim_weight_uom" default:""`
	Dim_width                *float64 `json:"dim_width" default:""`
	Dim_width_ouom           *string  `json:"dim_width_ouom" default:""`
	Dim_width_uom            *string  `json:"dim_width_uom" default:""`
	Effective_date           *string  `json:"effective_date" default:""`
	Expiry_date              *string  `json:"expiry_date" default:""`
	Lower_contained_store_id *string  `json:"lower_contained_store_id" default:""`
	Ppdm_guid                *string  `json:"ppdm_guid" default:""`
	Remark                   *string  `json:"remark" default:""`
	Source                   *string  `json:"source" default:""`
	Total_capacity           *int     `json:"total_capacity" default:""`
	Upper_contained_store_id *string  `json:"upper_contained_store_id" default:""`
	Used_capacity            *int     `json:"used_capacity" default:""`
	Used_capacity_date       *string  `json:"used_capacity_date" default:""`
	Row_changed_by           *string  `json:"row_changed_by" default:""`
	Row_changed_date         *string  `json:"row_changed_date" default:""`
	Row_created_by           *string  `json:"row_created_by" default:""`
	Row_created_date         *string  `json:"row_created_date" default:""`
	Row_effective_date       *string  `json:"row_effective_date" default:""`
	Row_expiry_date          *string  `json:"row_expiry_date" default:""`
	Row_quality              *string  `json:"row_quality" default:""`
}
