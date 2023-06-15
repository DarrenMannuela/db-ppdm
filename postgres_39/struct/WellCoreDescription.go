package dto

import (
	"time"
)

type Well_core_description struct {
	Uwi                      string     `json:"uwi"`
	Source                   string     `json:"source"`
	Core_id                  string     `json:"core_id"`
	Description_obs_no       int        `json:"description_obs_no"`
	Active_ind               *string    `json:"active_ind"`
	Base_depth               *float64   `json:"base_depth"`
	Base_depth_ouom          *string    `json:"base_depth_ouom"`
	Core_description_company *string    `json:"core_description_company"`
	Description_date         *time.Time `json:"description_date"`
	Dip_angle                *float64   `json:"dip_angle"`
	Effective_date           *time.Time `json:"effective_date"`
	Expiry_date              *time.Time `json:"expiry_date"`
	Interval_thickness       *float64   `json:"interval_thickness"`
	Interval_thickness_ouom  *string    `json:"interval_thickness_ouom"`
	Lithology_desc           *string    `json:"lithology_desc"`
	Porosity_length          *float64   `json:"porosity_length"`
	Porosity_length_ouom     *string    `json:"porosity_length_ouom"`
	Porosity_quality         *string    `json:"porosity_quality"`
	Porosity_type            *string    `json:"porosity_type"`
	Ppdm_guid                string     `json:"ppdm_guid"`
	Remark                   *string    `json:"remark"`
	Show_length              *float64   `json:"show_length"`
	Show_length_ouom         *string    `json:"show_length_ouom"`
	Show_type                *string    `json:"show_type"`
	Top_depth                *float64   `json:"top_depth"`
	Top_depth_ouom           *string    `json:"top_depth_ouom"`
	Row_changed_by           *string    `json:"row_changed_by"`
	Row_changed_date         *time.Time `json:"row_changed_date"`
	Row_created_by           *string    `json:"row_created_by"`
	Row_created_date         *time.Time `json:"row_created_date"`
	Row_effective_date       *time.Time `json:"row_effective_date"`
	Row_expiry_date          *time.Time `json:"row_expiry_date"`
	Row_quality              *string    `json:"row_quality"`
}
