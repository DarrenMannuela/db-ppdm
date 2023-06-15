package dto

type Sp_parcel_nts struct {
	Parcel_nts_id          string   `json:"parcel_nts_id" default:"source"`
	Active_ind             *string  `json:"active_ind" default:""`
	Area_id                *string  `json:"area_id" default:""`
	Area_type              *string  `json:"area_type" default:""`
	Block                  *string  `json:"block" default:""`
	Coord_system_id        *string  `json:"coord_system_id" default:""`
	Description            *string  `json:"description" default:""`
	Effective_date         *string  `json:"effective_date" default:""`
	Expiry_date            *string  `json:"expiry_date" default:""`
	Letter_quadrangle      *string  `json:"letter_quadrangle" default:""`
	Ppdm_guid              *string  `json:"ppdm_guid" default:""`
	Primary_quadrangle     *float64 `json:"primary_quadrangle" default:""`
	Quarter_unit           *string  `json:"quarter_unit" default:""`
	Reference_plan_num     *string  `json:"reference_plan_num" default:""`
	Remark                 *string  `json:"remark" default:""`
	Sixteenth              *float64 `json:"sixteenth" default:""`
	Source                 *string  `json:"source" default:""`
	Spatial_description_id *string  `json:"spatial_description_id" default:""`
	Spatial_obs_no         *int     `json:"spatial_obs_no" default:""`
	Unit                   *int     `json:"unit" default:""`
	Row_changed_by         *string  `json:"row_changed_by" default:""`
	Row_changed_date       *string  `json:"row_changed_date" default:""`
	Row_created_by         *string  `json:"row_created_by" default:""`
	Row_created_date       *string  `json:"row_created_date" default:""`
	Row_effective_date     *string  `json:"row_effective_date" default:""`
	Row_expiry_date        *string  `json:"row_expiry_date" default:""`
	Row_quality            *string  `json:"row_quality" default:""`
}
