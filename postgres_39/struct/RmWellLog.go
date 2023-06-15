package dto

import (
	"time"
)

type Rm_well_log struct {
	Info_item_subtype        string     `json:"info_item_subtype"`
	Information_item_id      string     `json:"information_item_id"`
	Active_ind               *string    `json:"active_ind"`
	Display_interval         *float64   `json:"display_interval"`
	Display_interval_uom     *string    `json:"display_interval_uom"`
	Display_scale            *float64   `json:"display_scale"`
	Display_scale_uom        *string    `json:"display_scale_uom"`
	Effective_date           *time.Time `json:"effective_date"`
	Expiry_date              *time.Time `json:"expiry_date"`
	Interpreted_ind          *string    `json:"interpreted_ind"`
	Ppdm_guid                string     `json:"ppdm_guid"`
	Remark                   *string    `json:"remark"`
	Reported_base_depth      *float64   `json:"reported_base_depth"`
	Reported_base_depth_ouom *string    `json:"reported_base_depth_ouom"`
	Reported_top_depth       *float64   `json:"reported_top_depth"`
	Reported_top_depth_ouom  *string    `json:"reported_top_depth_ouom"`
	Source                   *string    `json:"source"`
	Row_changed_by           *string    `json:"row_changed_by"`
	Row_changed_date         *time.Time `json:"row_changed_date"`
	Row_created_by           *string    `json:"row_created_by"`
	Row_created_date         *time.Time `json:"row_created_date"`
	Row_effective_date       *time.Time `json:"row_effective_date"`
	Row_expiry_date          *time.Time `json:"row_expiry_date"`
	Row_quality              *string    `json:"row_quality"`
}
