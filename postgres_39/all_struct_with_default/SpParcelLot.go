package dto

type Sp_parcel_lot struct {
	Parcel_lot_id          string   `json:"parcel_lot_id" default:"source"`
	Parcel_lot_type        string   `json:"parcel_lot_type" default:"source"`
	Parcel_lot_num         string   `json:"parcel_lot_num" default:"source"`
	Active_ind             *string  `json:"active_ind" default:""`
	Effective_date         *string  `json:"effective_date" default:""`
	Expiry_date            *string  `json:"expiry_date" default:""`
	Gross_size             *float64 `json:"gross_size" default:""`
	Gross_size_ouom        *string  `json:"gross_size_ouom" default:""`
	Lot_description        *string  `json:"lot_description" default:""`
	Lot_name               *string  `json:"lot_name" default:""`
	Parcel_congress_id     *string  `json:"parcel_congress_id" default:""`
	Parcel_ohio_id         *string  `json:"parcel_ohio_id" default:""`
	Parcel_pbl_id          *string  `json:"parcel_pbl_id" default:""`
	Parcel_texas_id        *string  `json:"parcel_texas_id" default:""`
	Ppdm_guid              *string  `json:"ppdm_guid" default:""`
	Remark                 *string  `json:"remark" default:""`
	Remark_type            *string  `json:"remark_type" default:""`
	Source                 *string  `json:"source" default:""`
	Spatial_description_id *string  `json:"spatial_description_id" default:""`
	Spatial_obs_no         *int     `json:"spatial_obs_no" default:""`
	Row_changed_by         *string  `json:"row_changed_by" default:""`
	Row_changed_date       *string  `json:"row_changed_date" default:""`
	Row_created_by         *string  `json:"row_created_by" default:""`
	Row_created_date       *string  `json:"row_created_date" default:""`
	Row_effective_date     *string  `json:"row_effective_date" default:""`
	Row_expiry_date        *string  `json:"row_expiry_date" default:""`
	Row_quality            *string  `json:"row_quality" default:""`
}
