package dto

import (
	"time"
)

type Rm_data_store struct {
	Store_id              string     `json:"store_id"`
	Active_ind            *string    `json:"active_ind"`
	Address_obs_no        *int       `json:"address_obs_no"`
	Address_source        *string    `json:"address_source"`
	Business_associate_id *string    `json:"business_associate_id"`
	Contained_by_store_id *string    `json:"contained_by_store_id"`
	Data_store_hier_id    *string    `json:"data_store_hier_id"`
	Data_store_name       *string    `json:"data_store_name"`
	Data_store_type       *string    `json:"data_store_type"`
	Effective_date        *time.Time `json:"effective_date"`
	Expiry_date           *time.Time `json:"expiry_date"`
	Hier_level_id         *string    `json:"hier_level_id"`
	Location_id           *string    `json:"location_id"`
	Long_location         *string    `json:"long_location"`
	Physical_store_status *string    `json:"physical_store_status"`
	Ppdm_guid             string     `json:"ppdm_guid"`
	Remark                *string    `json:"remark"`
	Short_location        *string    `json:"short_location"`
	Source                *string    `json:"source"`
	Total_capacity        *int       `json:"total_capacity"`
	Used_capacity         *int       `json:"used_capacity"`
	Used_capacity_date    *time.Time `json:"used_capacity_date"`
	Row_changed_by        *string    `json:"row_changed_by"`
	Row_changed_date      *time.Time `json:"row_changed_date"`
	Row_created_by        *string    `json:"row_created_by"`
	Row_created_date      *time.Time `json:"row_created_date"`
	Row_effective_date    *time.Time `json:"row_effective_date"`
	Row_expiry_date       *time.Time `json:"row_expiry_date"`
	Row_quality           *string    `json:"row_quality"`
}
