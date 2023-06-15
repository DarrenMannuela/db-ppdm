package dto

type Sp_parcel_congress struct {
	Parcel_congress_id      string   `json:"parcel_congress_id" default:"source"`
	Active_ind              *string  `json:"active_ind" default:""`
	Area_id                 *string  `json:"area_id" default:""`
	Area_type               *string  `json:"area_type" default:""`
	Congress_range          *float64 `json:"congress_range" default:""`
	Congress_section        *float64 `json:"congress_section" default:""`
	Congress_section_suffix *string  `json:"congress_section_suffix" default:""`
	Congress_township       *float64 `json:"congress_township" default:""`
	Congress_twp_name       *string  `json:"congress_twp_name" default:""`
	Coord_system_id         *string  `json:"coord_system_id" default:""`
	Description             *string  `json:"description" default:""`
	Effective_date          *string  `json:"effective_date" default:""`
	Ew_direction            *string  `json:"ew_direction" default:""`
	Expiry_date             *string  `json:"expiry_date" default:""`
	Ns_direction            *string  `json:"ns_direction" default:""`
	Ppdm_guid               *string  `json:"ppdm_guid" default:""`
	Principal_meridian      *string  `json:"principal_meridian" default:""`
	Reference_plan_num      *string  `json:"reference_plan_num" default:""`
	Remark                  *string  `json:"remark" default:""`
	Section_type            *string  `json:"section_type" default:""`
	Source                  *string  `json:"source" default:""`
	Spatial_description_id  *string  `json:"spatial_description_id" default:""`
	Spatial_obs_no          *int     `json:"spatial_obs_no" default:""`
	Spot_code               *string  `json:"spot_code" default:""`
	Row_changed_by          *string  `json:"row_changed_by" default:""`
	Row_changed_date        *string  `json:"row_changed_date" default:""`
	Row_created_by          *string  `json:"row_created_by" default:""`
	Row_created_date        *string  `json:"row_created_date" default:""`
	Row_effective_date      *string  `json:"row_effective_date" default:""`
	Row_expiry_date         *string  `json:"row_expiry_date" default:""`
	Row_quality             *string  `json:"row_quality" default:""`
}
