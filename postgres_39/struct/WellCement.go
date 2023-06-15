package dto

import (
	"time"
)

type Well_cement struct {
	Uwi                         string     `json:"uwi"`
	Well_tube_source            string     `json:"well_tube_source"`
	Tubing_type                 string     `json:"tubing_type"`
	Tubing_obs_no               int        `json:"tubing_obs_no"`
	Cement_obs_no               int        `json:"cement_obs_no"`
	Active_ind                  *string    `json:"active_ind"`
	Cement_amount               *float64   `json:"cement_amount"`
	Cement_amount_ouom          *string    `json:"cement_amount_ouom"`
	Cement_amount_uom           *string    `json:"cement_amount_uom"`
	Cement_ba_id                *string    `json:"cement_ba_id"`
	Cement_type                 *string    `json:"cement_type"`
	Effective_date              *time.Time `json:"effective_date"`
	Expiry_date                 *time.Time `json:"expiry_date"`
	Perforation_base_depth      *float64   `json:"perforation_base_depth"`
	Perforation_base_depth_ouom *string    `json:"perforation_base_depth_ouom"`
	Perforation_count           *int       `json:"perforation_count"`
	Perforation_per_uom         *string    `json:"perforation_per_uom"`
	Perforation_top_depth       *float64   `json:"perforation_top_depth"`
	Perforation_top_depth_ouom  *string    `json:"perforation_top_depth_ouom"`
	Ppdm_guid                   string     `json:"ppdm_guid"`
	Recement_ind                *string    `json:"recement_ind"`
	Remark                      *string    `json:"remark"`
	Source                      *string    `json:"source"`
	Stage_no                    *int       `json:"stage_no"`
	Row_changed_by              *string    `json:"row_changed_by"`
	Row_changed_date            *time.Time `json:"row_changed_date"`
	Row_created_by              *string    `json:"row_created_by"`
	Row_created_date            *time.Time `json:"row_created_date"`
	Row_effective_date          *time.Time `json:"row_effective_date"`
	Row_expiry_date             *time.Time `json:"row_expiry_date"`
	Row_quality                 *string    `json:"row_quality"`
}
