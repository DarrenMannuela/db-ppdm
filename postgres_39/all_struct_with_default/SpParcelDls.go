package dto

type Sp_parcel_dls struct {
	Parcel_dls_id               string   `json:"parcel_dls_id" default:"source"`
	Active_ind                  *string  `json:"active_ind" default:""`
	Area_id                     *string  `json:"area_id" default:""`
	Area_type                   *string  `json:"area_type" default:""`
	Coord_system_id             *string  `json:"coord_system_id" default:""`
	Description                 *string  `json:"description" default:""`
	Dls_legal_subdivision       *float64 `json:"dls_legal_subdivision" default:""`
	Dls_meridian                *float64 `json:"dls_meridian" default:""`
	Dls_meridian_direction      *string  `json:"dls_meridian_direction" default:""`
	Dls_quarter_section         *string  `json:"dls_quarter_section" default:""`
	Dls_quarter_section_quarter *string  `json:"dls_quarter_section_quarter" default:""`
	Dls_range                   *float64 `json:"dls_range" default:""`
	Dls_range_modifier          *string  `json:"dls_range_modifier" default:""`
	Dls_section                 *float64 `json:"dls_section" default:""`
	Dls_township                *float64 `json:"dls_township" default:""`
	Dls_township_modifier       *string  `json:"dls_township_modifier" default:""`
	Effective_date              *string  `json:"effective_date" default:""`
	Expiry_date                 *string  `json:"expiry_date" default:""`
	Local_coord_system_id       *string  `json:"local_coord_system_id" default:""`
	Ppdm_guid                   *string  `json:"ppdm_guid" default:""`
	Principal_meridian          *string  `json:"principal_meridian" default:""`
	Reference_plan_num          *string  `json:"reference_plan_num" default:""`
	Remark                      *string  `json:"remark" default:""`
	Source                      *string  `json:"source" default:""`
	Spatial_description_id      *string  `json:"spatial_description_id" default:""`
	Spatial_obs_no              *int     `json:"spatial_obs_no" default:""`
	Row_changed_by              *string  `json:"row_changed_by" default:""`
	Row_changed_date            *string  `json:"row_changed_date" default:""`
	Row_created_by              *string  `json:"row_created_by" default:""`
	Row_created_date            *string  `json:"row_created_date" default:""`
	Row_effective_date          *string  `json:"row_effective_date" default:""`
	Row_expiry_date             *string  `json:"row_expiry_date" default:""`
	Row_quality                 *string  `json:"row_quality" default:""`
}
