package dto

import (
	"time"
)

type Rm_image_loc struct {
	Physical_item_id   string     `json:"physical_item_id"`
	Log_section_id     string     `json:"log_section_id"`
	Position_id        string     `json:"position_id"`
	Active_ind         *string    `json:"active_ind"`
	Effective_date     *time.Time `json:"effective_date"`
	Expiry_date        *time.Time `json:"expiry_date"`
	Log_depth          *float64   `json:"log_depth"`
	Log_depth_ouom     *string    `json:"log_depth_ouom"`
	Log_depth_uom      *string    `json:"log_depth_uom"`
	Position_type      *string    `json:"position_type"`
	Ppdm_guid          string     `json:"ppdm_guid"`
	Remark             *string    `json:"remark"`
	Source             *string    `json:"source"`
	X_position         *float64   `json:"x_position"`
	Y_position         *float64   `json:"y_position"`
	Row_changed_by     *string    `json:"row_changed_by"`
	Row_changed_date   *time.Time `json:"row_changed_date"`
	Row_created_by     *string    `json:"row_created_by"`
	Row_created_date   *time.Time `json:"row_created_date"`
	Row_effective_date *time.Time `json:"row_effective_date"`
	Row_expiry_date    *time.Time `json:"row_expiry_date"`
	Row_quality        *string    `json:"row_quality"`
}
