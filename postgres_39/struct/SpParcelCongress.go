package dto

import (
	"time"
)

type Sp_parcel_congress struct {
	Parcel_congress_id      string     `json:"parcel_congress_id"`
	Active_ind              *string    `json:"active_ind"`
	Area_id                 *string    `json:"area_id"`
	Area_type               *string    `json:"area_type"`
	Congress_range          *float64   `json:"congress_range"`
	Congress_section        *float64   `json:"congress_section"`
	Congress_section_suffix *string    `json:"congress_section_suffix"`
	Congress_township       *float64   `json:"congress_township"`
	Congress_twp_name       *string    `json:"congress_twp_name"`
	Coord_system_id         *string    `json:"coord_system_id"`
	Description             *string    `json:"description"`
	Effective_date          *time.Time `json:"effective_date"`
	Ew_direction            *string    `json:"ew_direction"`
	Expiry_date             *time.Time `json:"expiry_date"`
	Ns_direction            *string    `json:"ns_direction"`
	Ppdm_guid               string     `json:"ppdm_guid"`
	Principal_meridian      *string    `json:"principal_meridian"`
	Reference_plan_num      *string    `json:"reference_plan_num"`
	Remark                  *string    `json:"remark"`
	Section_type            *string    `json:"section_type"`
	Source                  *string    `json:"source"`
	Spatial_description_id  *string    `json:"spatial_description_id"`
	Spatial_obs_no          *int       `json:"spatial_obs_no"`
	Spot_code               *string    `json:"spot_code"`
	Row_changed_by          *string    `json:"row_changed_by"`
	Row_changed_date        *time.Time `json:"row_changed_date"`
	Row_created_by          *string    `json:"row_created_by"`
	Row_created_date        *time.Time `json:"row_created_date"`
	Row_effective_date      *time.Time `json:"row_effective_date"`
	Row_expiry_date         *time.Time `json:"row_expiry_date"`
	Row_quality             *string    `json:"row_quality"`
}
