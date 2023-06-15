package dto

import (
	"time"
)

type Rm_information_item struct {
	Info_item_subtype     string     `json:"info_item_subtype"`
	Information_item_id   string     `json:"information_item_id"`
	Abstract              *string    `json:"abstract"`
	Accepted_date         *time.Time `json:"accepted_date"`
	Access_condition      *string    `json:"access_condition"`
	Active_ind            *string    `json:"active_ind"`
	Available_date        *time.Time `json:"available_date"`
	Coord_acquisition_id  *string    `json:"coord_acquisition_id"`
	Copyright_date        *time.Time `json:"copyright_date"`
	Effective_date        *time.Time `json:"effective_date"`
	Expiry_date           *time.Time `json:"expiry_date"`
	Geog_coord_system_id  *string    `json:"geog_coord_system_id"`
	Group_ind             *string    `json:"group_ind"`
	Issue_date            *time.Time `json:"issue_date"`
	Item_category         *string    `json:"item_category"`
	Item_sub_category     *string    `json:"item_sub_category"`
	Local_coord_system_id *string    `json:"local_coord_system_id"`
	Map_coord_system_id   *string    `json:"map_coord_system_id"`
	Max_latitude          *float64   `json:"max_latitude"`
	Max_longitude         *float64   `json:"max_longitude"`
	Min_latitude          *float64   `json:"min_latitude"`
	Min_longitude         *float64   `json:"min_longitude"`
	Modified_date         *time.Time `json:"modified_date"`
	Origin_date           *time.Time `json:"origin_date"`
	Ppdm_guid             string     `json:"ppdm_guid"`
	Publish_date          *time.Time `json:"publish_date"`
	Purpose               *string    `json:"purpose"`
	Reference_num         *string    `json:"reference_num"`
	Remark                *string    `json:"remark"`
	Security_desc         *string    `json:"security_desc"`
	Source                *string    `json:"source"`
	Source_document_id    *string    `json:"source_document_id"`
	Submit_date           *time.Time `json:"submit_date"`
	Time_period_desc      *string    `json:"time_period_desc"`
	Title                 *string    `json:"title"`
	Use_condition         *string    `json:"use_condition"`
	Version_num           *string    `json:"version_num"`
	Row_changed_by        *string    `json:"row_changed_by"`
	Row_changed_date      *time.Time `json:"row_changed_date"`
	Row_created_by        *string    `json:"row_created_by"`
	Row_created_date      *time.Time `json:"row_created_date"`
	Row_effective_date    *time.Time `json:"row_effective_date"`
	Row_expiry_date       *time.Time `json:"row_expiry_date"`
	Row_quality           *string    `json:"row_quality"`
}
