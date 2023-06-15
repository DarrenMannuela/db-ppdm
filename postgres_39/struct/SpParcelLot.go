package dto

import (
	"time"
)

type Sp_parcel_lot struct {
	Parcel_lot_id          string     `json:"parcel_lot_id"`
	Parcel_lot_type        string     `json:"parcel_lot_type"`
	Parcel_lot_num         string     `json:"parcel_lot_num"`
	Active_ind             *string    `json:"active_ind"`
	Effective_date         *time.Time `json:"effective_date"`
	Expiry_date            *time.Time `json:"expiry_date"`
	Gross_size             *float64   `json:"gross_size"`
	Gross_size_ouom        *string    `json:"gross_size_ouom"`
	Lot_description        *string    `json:"lot_description"`
	Lot_name               *string    `json:"lot_name"`
	Parcel_congress_id     *string    `json:"parcel_congress_id"`
	Parcel_ohio_id         *string    `json:"parcel_ohio_id"`
	Parcel_pbl_id          *string    `json:"parcel_pbl_id"`
	Parcel_texas_id        *string    `json:"parcel_texas_id"`
	Ppdm_guid              string     `json:"ppdm_guid"`
	Remark                 *string    `json:"remark"`
	Remark_type            *string    `json:"remark_type"`
	Source                 *string    `json:"source"`
	Spatial_description_id *string    `json:"spatial_description_id"`
	Spatial_obs_no         *int       `json:"spatial_obs_no"`
	Row_changed_by         *string    `json:"row_changed_by"`
	Row_changed_date       *time.Time `json:"row_changed_date"`
	Row_created_by         *string    `json:"row_created_by"`
	Row_created_date       *time.Time `json:"row_created_date"`
	Row_effective_date     *time.Time `json:"row_effective_date"`
	Row_expiry_date        *time.Time `json:"row_expiry_date"`
	Row_quality            *string    `json:"row_quality"`
}
