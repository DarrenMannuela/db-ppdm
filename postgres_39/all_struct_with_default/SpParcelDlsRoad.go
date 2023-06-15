package dto

type Sp_parcel_dls_road struct {
	Parcel_dls_id          string  `json:"parcel_dls_id" default:"source"`
	Dls_road_allowance_id  string  `json:"dls_road_allowance_id" default:"source"`
	Active_ind             *string `json:"active_ind" default:""`
	Effective_date         *string `json:"effective_date" default:""`
	Expiry_date            *string `json:"expiry_date" default:""`
	Ppdm_guid              *string `json:"ppdm_guid" default:""`
	Remark                 *string `json:"remark" default:""`
	Road_allowance_portion *string `json:"road_allowance_portion" default:""`
	Road_allow_desc        *string `json:"road_allow_desc" default:""`
	Source                 *string `json:"source" default:""`
	Row_changed_by         *string `json:"row_changed_by" default:""`
	Row_changed_date       *string `json:"row_changed_date" default:""`
	Row_created_by         *string `json:"row_created_by" default:""`
	Row_created_date       *string `json:"row_created_date" default:""`
	Row_effective_date     *string `json:"row_effective_date" default:""`
	Row_expiry_date        *string `json:"row_expiry_date" default:""`
	Row_quality            *string `json:"row_quality" default:""`
}
