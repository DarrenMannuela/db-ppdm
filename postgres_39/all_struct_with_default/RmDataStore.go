package dto

type Rm_data_store struct {
	Store_id              string  `json:"store_id" default:"source"`
	Active_ind            *string `json:"active_ind" default:""`
	Address_obs_no        *int    `json:"address_obs_no" default:""`
	Address_source        *string `json:"address_source" default:""`
	Business_associate_id *string `json:"business_associate_id" default:""`
	Contained_by_store_id *string `json:"contained_by_store_id" default:""`
	Data_store_hier_id    *string `json:"data_store_hier_id" default:""`
	Data_store_name       *string `json:"data_store_name" default:""`
	Data_store_type       *string `json:"data_store_type" default:""`
	Effective_date        *string `json:"effective_date" default:""`
	Expiry_date           *string `json:"expiry_date" default:""`
	Hier_level_id         *string `json:"hier_level_id" default:""`
	Location_id           *string `json:"location_id" default:""`
	Long_location         *string `json:"long_location" default:""`
	Physical_store_status *string `json:"physical_store_status" default:""`
	Ppdm_guid             *string `json:"ppdm_guid" default:""`
	Remark                *string `json:"remark" default:""`
	Short_location        *string `json:"short_location" default:""`
	Source                *string `json:"source" default:""`
	Total_capacity        *int    `json:"total_capacity" default:""`
	Used_capacity         *int    `json:"used_capacity" default:""`
	Used_capacity_date    *string `json:"used_capacity_date" default:""`
	Row_changed_by        *string `json:"row_changed_by" default:""`
	Row_changed_date      *string `json:"row_changed_date" default:""`
	Row_created_by        *string `json:"row_created_by" default:""`
	Row_created_date      *string `json:"row_created_date" default:""`
	Row_effective_date    *string `json:"row_effective_date" default:""`
	Row_expiry_date       *string `json:"row_expiry_date" default:""`
	Row_quality           *string `json:"row_quality" default:""`
}
