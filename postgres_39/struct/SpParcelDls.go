package dto

import (
	"time"
)

type Sp_parcel_dls struct {
	Parcel_dls_id               string     `json:"parcel_dls_id"`
	Active_ind                  *string    `json:"active_ind"`
	Area_id                     *string    `json:"area_id"`
	Area_type                   *string    `json:"area_type"`
	Coord_system_id             *string    `json:"coord_system_id"`
	Description                 *string    `json:"description"`
	Dls_legal_subdivision       *float64   `json:"dls_legal_subdivision"`
	Dls_meridian                *float64   `json:"dls_meridian"`
	Dls_meridian_direction      *string    `json:"dls_meridian_direction"`
	Dls_quarter_section         *string    `json:"dls_quarter_section"`
	Dls_quarter_section_quarter *string    `json:"dls_quarter_section_quarter"`
	Dls_range                   *float64   `json:"dls_range"`
	Dls_range_modifier          *string    `json:"dls_range_modifier"`
	Dls_section                 *float64   `json:"dls_section"`
	Dls_township                *float64   `json:"dls_township"`
	Dls_township_modifier       *string    `json:"dls_township_modifier"`
	Effective_date              *time.Time `json:"effective_date"`
	Expiry_date                 *time.Time `json:"expiry_date"`
	Local_coord_system_id       *string    `json:"local_coord_system_id"`
	Ppdm_guid                   string     `json:"ppdm_guid"`
	Principal_meridian          *string    `json:"principal_meridian"`
	Reference_plan_num          *string    `json:"reference_plan_num"`
	Remark                      *string    `json:"remark"`
	Source                      *string    `json:"source"`
	Spatial_description_id      *string    `json:"spatial_description_id"`
	Spatial_obs_no              *int       `json:"spatial_obs_no"`
	Row_changed_by              *string    `json:"row_changed_by"`
	Row_changed_date            *time.Time `json:"row_changed_date"`
	Row_created_by              *string    `json:"row_created_by"`
	Row_created_date            *time.Time `json:"row_created_date"`
	Row_effective_date          *time.Time `json:"row_effective_date"`
	Row_expiry_date             *time.Time `json:"row_expiry_date"`
	Row_quality                 *string    `json:"row_quality"`
}
