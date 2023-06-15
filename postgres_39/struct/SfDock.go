package dto

import (
	"time"
)

type Sf_dock struct {
	Sf_subtype         string     `json:"sf_subtype"`
	Dock_id            string     `json:"dock_id"`
	Active_ind         *string    `json:"active_ind"`
	Dock_length        *float64   `json:"dock_length"`
	Dock_length_ouom   *string    `json:"dock_length_ouom"`
	Dock_type          *string    `json:"dock_type"`
	Effective_date     *time.Time `json:"effective_date"`
	Expiry_date        *time.Time `json:"expiry_date"`
	Ppdm_guid          string     `json:"ppdm_guid"`
	Remark             *string    `json:"remark"`
	Source             *string    `json:"source"`
	Surface_type       *string    `json:"surface_type"`
	Water_depth        *float64   `json:"water_depth"`
	Water_depth_ouom   *string    `json:"water_depth_ouom"`
	Row_changed_by     *string    `json:"row_changed_by"`
	Row_changed_date   *time.Time `json:"row_changed_date"`
	Row_created_by     *string    `json:"row_created_by"`
	Row_created_date   *time.Time `json:"row_created_date"`
	Row_effective_date *time.Time `json:"row_effective_date"`
	Row_expiry_date    *time.Time `json:"row_expiry_date"`
	Row_quality        *string    `json:"row_quality"`
}
