package dto

type Sp_parcel_libya struct {
	Parcel_lybia_id        string  `json:"parcel_lybia_id" default:"source"`
	Active_ind             *string `json:"active_ind" default:""`
	Area_id                *string `json:"area_id" default:""`
	Area_type              *string `json:"area_type" default:""`
	Coord_system_id        *string `json:"coord_system_id" default:""`
	Description            *string `json:"description" default:""`
	Effective_date         *string `json:"effective_date" default:""`
	Ew_direction           *string `json:"ew_direction" default:""`
	Expiry_date            *string `json:"expiry_date" default:""`
	Libya_area             *string `json:"libya_area" default:""`
	Libya_block            *string `json:"libya_block" default:""`
	Libya_section          *string `json:"libya_section" default:""`
	Libya_subsection       *string `json:"libya_subsection" default:""`
	Ppdm_guid              *string `json:"ppdm_guid" default:""`
	Remark                 *string `json:"remark" default:""`
	Source                 *string `json:"source" default:""`
	Spatial_description_id *string `json:"spatial_description_id" default:""`
	Spatial_obs_no         *int    `json:"spatial_obs_no" default:""`
	Row_changed_by         *string `json:"row_changed_by" default:""`
	Row_changed_date       *string `json:"row_changed_date" default:""`
	Row_created_by         *string `json:"row_created_by" default:""`
	Row_created_date       *string `json:"row_created_date" default:""`
	Row_effective_date     *string `json:"row_effective_date" default:""`
	Row_expiry_date        *string `json:"row_expiry_date" default:""`
	Row_quality            *string `json:"row_quality" default:""`
}
