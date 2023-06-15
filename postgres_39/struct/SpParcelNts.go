package dto

import (
	"time"
)

type Sp_parcel_nts struct {
	Parcel_nts_id          string     `json:"parcel_nts_id"`
	Active_ind             *string    `json:"active_ind"`
	Area_id                *string    `json:"area_id"`
	Area_type              *string    `json:"area_type"`
	Block                  *string    `json:"block"`
	Coord_system_id        *string    `json:"coord_system_id"`
	Description            *string    `json:"description"`
	Effective_date         *time.Time `json:"effective_date"`
	Expiry_date            *time.Time `json:"expiry_date"`
	Letter_quadrangle      *string    `json:"letter_quadrangle"`
	Ppdm_guid              string     `json:"ppdm_guid"`
	Primary_quadrangle     *float64   `json:"primary_quadrangle"`
	Quarter_unit           *string    `json:"quarter_unit"`
	Reference_plan_num     *string    `json:"reference_plan_num"`
	Remark                 *string    `json:"remark"`
	Sixteenth              *float64   `json:"sixteenth"`
	Source                 *string    `json:"source"`
	Spatial_description_id *string    `json:"spatial_description_id"`
	Spatial_obs_no         *int       `json:"spatial_obs_no"`
	Unit                   *int       `json:"unit"`
	Row_changed_by         *string    `json:"row_changed_by"`
	Row_changed_date       *time.Time `json:"row_changed_date"`
	Row_created_by         *string    `json:"row_created_by"`
	Row_created_date       *time.Time `json:"row_created_date"`
	Row_effective_date     *time.Time `json:"row_effective_date"`
	Row_expiry_date        *time.Time `json:"row_expiry_date"`
	Row_quality            *string    `json:"row_quality"`
}
