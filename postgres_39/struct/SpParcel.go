package dto

import (
	"time"
)

type Sp_parcel struct {
	Spatial_description_id string     `json:"spatial_description_id"`
	Spatial_obs_no         int        `json:"spatial_obs_no"`
	Parcel_id              string     `json:"parcel_id"`
	Active_ind             *string    `json:"active_ind"`
	Description            *string    `json:"description"`
	Dls_road_allowance_id  *string    `json:"dls_road_allowance_id"`
	Effective_date         *time.Time `json:"effective_date"`
	Expiry_date            *time.Time `json:"expiry_date"`
	Gross_size             *float64   `json:"gross_size"`
	Gross_size_ouom        *string    `json:"gross_size_ouom"`
	Inactivation_date      *time.Time `json:"inactivation_date"`
	Local_coord_system_id  *string    `json:"local_coord_system_id"`
	Parcel_carter_id       *string    `json:"parcel_carter_id"`
	Parcel_congress_id     *string    `json:"parcel_congress_id"`
	Parcel_dls_id          *string    `json:"parcel_dls_id"`
	Parcel_fps_id          *string    `json:"parcel_fps_id"`
	Parcel_libya_id        *string    `json:"parcel_libya_id"`
	Parcel_lot_id          *string    `json:"parcel_lot_id"`
	Parcel_lot_num         *string    `json:"parcel_lot_num"`
	Parcel_lot_type        *string    `json:"parcel_lot_type"`
	Parcel_ne_loc_id       *string    `json:"parcel_ne_loc_id"`
	Parcel_north_sea_id    *string    `json:"parcel_north_sea_id"`
	Parcel_nts_id          *string    `json:"parcel_nts_id"`
	Parcel_offshore_id     *string    `json:"parcel_offshore_id"`
	Parcel_ohio_id         *string    `json:"parcel_ohio_id"`
	Parcel_pbl_id          *string    `json:"parcel_pbl_id"`
	Parcel_texas_id        *string    `json:"parcel_texas_id"`
	Percent_ownership      *float64   `json:"percent_ownership"`
	Ppdm_guid              string     `json:"ppdm_guid"`
	Remark                 *string    `json:"remark"`
	Source                 *string    `json:"source"`
	Row_changed_by         *string    `json:"row_changed_by"`
	Row_changed_date       *time.Time `json:"row_changed_date"`
	Row_created_by         *string    `json:"row_created_by"`
	Row_created_date       *time.Time `json:"row_created_date"`
	Row_effective_date     *time.Time `json:"row_effective_date"`
	Row_expiry_date        *time.Time `json:"row_expiry_date"`
	Row_quality            *string    `json:"row_quality"`
}
