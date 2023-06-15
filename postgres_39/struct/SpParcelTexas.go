package dto

import (
	"time"
)

type Sp_parcel_texas struct {
	Parcel_texas_id        string     `json:"parcel_texas_id"`
	Abstract_num           *string    `json:"abstract_num"`
	Active_ind             *string    `json:"active_ind"`
	Area_id                *string    `json:"area_id"`
	Area_type              *string    `json:"area_type"`
	Block_fraction         *string    `json:"block_fraction"`
	Coord_system_id        *string    `json:"coord_system_id"`
	Description            *string    `json:"description"`
	Effective_date         *time.Time `json:"effective_date"`
	Expiry_date            *time.Time `json:"expiry_date"`
	Labor                  *string    `json:"labor"`
	League                 *string    `json:"league"`
	Ns_direction           *string    `json:"ns_direction"`
	Porcion_num            *string    `json:"porcion_num"`
	Porcion_survey_name    *string    `json:"porcion_survey_name"`
	Ppdm_guid              string     `json:"ppdm_guid"`
	Reference_plan_num     *string    `json:"reference_plan_num"`
	Remark                 *string    `json:"remark"`
	Section_fraction       *string    `json:"section_fraction"`
	Source                 *string    `json:"source"`
	Spatial_description_id *string    `json:"spatial_description_id"`
	Spatial_obs_no         *int       `json:"spatial_obs_no"`
	Spot_code              *string    `json:"spot_code"`
	Texas_block            *string    `json:"texas_block"`
	Texas_lot              *string    `json:"texas_lot"`
	Texas_section          *string    `json:"texas_section"`
	Texas_share            *string    `json:"texas_share"`
	Texas_subdivision      *string    `json:"texas_subdivision"`
	Texas_survey           *string    `json:"texas_survey"`
	Texas_township         *float64   `json:"texas_township"`
	Row_changed_by         *string    `json:"row_changed_by"`
	Row_changed_date       *time.Time `json:"row_changed_date"`
	Row_created_by         *string    `json:"row_created_by"`
	Row_created_date       *time.Time `json:"row_created_date"`
	Row_effective_date     *time.Time `json:"row_effective_date"`
	Row_expiry_date        *time.Time `json:"row_expiry_date"`
	Row_quality            *string    `json:"row_quality"`
}
