package dto

import (
	"time"
)

type Sf_port struct {
	Sf_subtype           string     `json:"sf_subtype"`
	Support_facility_id  string     `json:"support_facility_id"`
	Active_ind           *string    `json:"active_ind"`
	Area_size            *float64   `json:"area_size"`
	Area_size_ouom       *string    `json:"area_size_ouom"`
	Capacity             *float64   `json:"capacity"`
	Capacity_ouom        *string    `json:"capacity_ouom"`
	Capacity_uom         *string    `json:"capacity_uom"`
	Effective_date       *time.Time `json:"effective_date"`
	Expiry_date          *time.Time `json:"expiry_date"`
	Min_water_depth      *float64   `json:"min_water_depth"`
	Min_water_depth_ouom *string    `json:"min_water_depth_ouom"`
	Ppdm_guid            string     `json:"ppdm_guid"`
	Remark               *string    `json:"remark"`
	Source               *string    `json:"source"`
	Row_changed_by       *string    `json:"row_changed_by"`
	Row_changed_date     *time.Time `json:"row_changed_date"`
	Row_created_by       *string    `json:"row_created_by"`
	Row_created_date     *time.Time `json:"row_created_date"`
	Row_effective_date   *time.Time `json:"row_effective_date"`
	Row_expiry_date      *time.Time `json:"row_expiry_date"`
	Row_quality          *string    `json:"row_quality"`
}
