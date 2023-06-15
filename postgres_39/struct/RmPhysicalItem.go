package dto

import (
	"time"
)

type Rm_physical_item struct {
	Physical_item_id        string     `json:"physical_item_id"`
	Active_ind              *string    `json:"active_ind"`
	Bar_code                *string    `json:"bar_code"`
	Catalogue_by_name       *string    `json:"catalogue_by_name"`
	Catalogue_date          *time.Time `json:"catalogue_date"`
	Certified_true_copy_ind *string    `json:"certified_true_copy_ind"`
	Circulation_allowed_ind *string    `json:"circulation_allowed_ind"`
	Circulation_out_ind     *string    `json:"circulation_out_ind"`
	Color_format            *string    `json:"color_format"`
	Create_date             *time.Time `json:"create_date"`
	Digital_format          *string    `json:"digital_format"`
	Digital_size            *float64   `json:"digital_size"`
	Digital_size_uom        *string    `json:"digital_size_uom"`
	Dim_height              *float64   `json:"dim_height"`
	Dim_height_uom          *string    `json:"dim_height_uom"`
	Dim_length              *float64   `json:"dim_length"`
	Dim_length_uom          *string    `json:"dim_length_uom"`
	Dim_weight              *float64   `json:"dim_weight"`
	Dim_weight_uom          *string    `json:"dim_weight_uom"`
	Dim_width               *float64   `json:"dim_width"`
	Dim_width_uom           *string    `json:"dim_width_uom"`
	Effective_date          *time.Time `json:"effective_date"`
	Expiry_date             *time.Time `json:"expiry_date"`
	File_count              *int       `json:"file_count"`
	Group_ind               *string    `json:"group_ind"`
	Image_resolution_uom    *string    `json:"image_resolution_uom"`
	Image_x_resolution      *float64   `json:"image_x_resolution"`
	Image_y_resolution      *float64   `json:"image_y_resolution"`
	Item_category           *string    `json:"item_category"`
	Item_sub_category       *string    `json:"item_sub_category"`
	Label                   *string    `json:"label"`
	Language                *string    `json:"language"`
	Last_condition          *string    `json:"last_condition"`
	Location_reference      *string    `json:"location_reference"`
	Media_type              *string    `json:"media_type"`
	Original_file_name      *string    `json:"original_file_name"`
	Original_ind            *string    `json:"original_ind"`
	Output_file_name        *string    `json:"output_file_name"`
	Page_count              *float64   `json:"page_count"`
	Physical_item_status    *string    `json:"physical_item_status"`
	Ppdm_guid               string     `json:"ppdm_guid"`
	Preferred_ind           *string    `json:"preferred_ind"`
	Qc_by_name              *string    `json:"qc_by_name"`
	Remark                  *string    `json:"remark"`
	Retention_period        *string    `json:"retention_period"`
	Sale_allowed_ind        *string    `json:"sale_allowed_ind"`
	Source                  *string    `json:"source"`
	Store_id                *string    `json:"store_id"`
	Row_changed_by          *string    `json:"row_changed_by"`
	Row_changed_date        *time.Time `json:"row_changed_date"`
	Row_created_by          *string    `json:"row_created_by"`
	Row_created_date        *time.Time `json:"row_created_date"`
	Row_effective_date      *time.Time `json:"row_effective_date"`
	Row_expiry_date         *time.Time `json:"row_expiry_date"`
	Row_quality             *string    `json:"row_quality"`
}
