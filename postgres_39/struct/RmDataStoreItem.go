package dto

import (
	"time"
)

type Rm_data_store_item struct {
	Store_id           string     `json:"store_id"`
	Structure_obs_no   int        `json:"structure_obs_no"`
	Item_obs_no        int        `json:"item_obs_no"`
	Active_ind         *string    `json:"active_ind"`
	Effective_date     *time.Time `json:"effective_date"`
	Expiry_date        *time.Time `json:"expiry_date"`
	Item_category      *string    `json:"item_category"`
	Item_sub_category  *string    `json:"item_sub_category"`
	Ppdm_guid          string     `json:"ppdm_guid"`
	Remark             *string    `json:"remark"`
	Source             *string    `json:"source"`
	Total_capacity     *int       `json:"total_capacity"`
	Used_capacity      *int       `json:"used_capacity"`
	Used_capacity_date *time.Time `json:"used_capacity_date"`
	Row_changed_by     *string    `json:"row_changed_by"`
	Row_changed_date   *time.Time `json:"row_changed_date"`
	Row_created_by     *string    `json:"row_created_by"`
	Row_created_date   *time.Time `json:"row_created_date"`
	Row_effective_date *time.Time `json:"row_effective_date"`
	Row_expiry_date    *time.Time `json:"row_expiry_date"`
	Row_quality        *string    `json:"row_quality"`
}
