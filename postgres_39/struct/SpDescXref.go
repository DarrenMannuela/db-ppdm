package dto

import (
	"time"
)

type Sp_desc_xref struct {
	Spatial_description_id   string     `json:"spatial_description_id"`
	Spatial_obs_no           int        `json:"spatial_obs_no"`
	Spatial_description_id_2 string     `json:"spatial_description_id_2"`
	Spatial_obs_no_2         int        `json:"spatial_obs_no_2"`
	Active_ind               *string    `json:"active_ind"`
	Effective_date           *time.Time `json:"effective_date"`
	Expiry_date              *time.Time `json:"expiry_date"`
	Overlap_size             *float64   `json:"overlap_size"`
	Overlap_size_ouom        *string    `json:"overlap_size_ouom"`
	Ppdm_guid                string     `json:"ppdm_guid"`
	Remark                   *string    `json:"remark"`
	Source                   *string    `json:"source"`
	Spatial_xref_type        *string    `json:"spatial_xref_type"`
	Row_changed_by           *string    `json:"row_changed_by"`
	Row_changed_date         *time.Time `json:"row_changed_date"`
	Row_created_by           *string    `json:"row_created_by"`
	Row_created_date         *time.Time `json:"row_created_date"`
	Row_effective_date       *time.Time `json:"row_effective_date"`
	Row_expiry_date          *time.Time `json:"row_expiry_date"`
	Row_quality              *string    `json:"row_quality"`
}
