package dto

import (
	"time"
)

type Sf_platform struct {
	Sf_subtype         string     `json:"sf_subtype"`
	Platform_id        string     `json:"platform_id"`
	Active_ind         *string    `json:"active_ind"`
	Drill_slot_count   *int       `json:"drill_slot_count"`
	Effective_date     *time.Time `json:"effective_date"`
	Expiry_date        *time.Time `json:"expiry_date"`
	Installed_date     *time.Time `json:"installed_date"`
	Platform_name      *string    `json:"platform_name"`
	Platform_type      *string    `json:"platform_type"`
	Ppdm_guid          string     `json:"ppdm_guid"`
	Remark             *string    `json:"remark"`
	Removal_date       *time.Time `json:"removal_date"`
	Source             *string    `json:"source"`
	Water_depth        *float64   `json:"water_depth"`
	Water_depth_datum  *string    `json:"water_depth_datum"`
	Water_depth_ouom   *string    `json:"water_depth_ouom"`
	Row_changed_by     *string    `json:"row_changed_by"`
	Row_changed_date   *time.Time `json:"row_changed_date"`
	Row_created_by     *string    `json:"row_created_by"`
	Row_created_date   *time.Time `json:"row_created_date"`
	Row_effective_date *time.Time `json:"row_effective_date"`
	Row_expiry_date    *time.Time `json:"row_expiry_date"`
	Row_quality        *string    `json:"row_quality"`
}
