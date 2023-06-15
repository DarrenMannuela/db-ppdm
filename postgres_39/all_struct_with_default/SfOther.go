package dto

type Sf_other struct {
	Sf_subtype          string   `json:"sf_subtype" default:"source"`
	Support_facility_id string   `json:"support_facility_id" default:"source"`
	Active_ind          *string  `json:"active_ind" default:""`
	Area_size           *float64 `json:"area_size" default:""`
	Area_size_ouom      *string  `json:"area_size_ouom" default:""`
	Capacity            *float64 `json:"capacity" default:""`
	Capacity_ouom       *string  `json:"capacity_ouom" default:""`
	Capacity_uom        *string  `json:"capacity_uom" default:""`
	Effective_date      *string  `json:"effective_date" default:""`
	Expiry_date         *string  `json:"expiry_date" default:""`
	Height              *float64 `json:"height" default:""`
	Height_ouom         *string  `json:"height_ouom" default:""`
	Length              *float64 `json:"length" default:""`
	Length_ouom         *string  `json:"length_ouom" default:""`
	Ppdm_guid           *string  `json:"ppdm_guid" default:""`
	Remark              *string  `json:"remark" default:""`
	Source              *string  `json:"source" default:""`
	Width               *float64 `json:"width" default:""`
	Width_ouom          *string  `json:"width_ouom" default:""`
	Row_changed_by      *string  `json:"row_changed_by" default:""`
	Row_changed_date    *string  `json:"row_changed_date" default:""`
	Row_created_by      *string  `json:"row_created_by" default:""`
	Row_created_date    *string  `json:"row_created_date" default:""`
	Row_effective_date  *string  `json:"row_effective_date" default:""`
	Row_expiry_date     *string  `json:"row_expiry_date" default:""`
	Row_quality         *string  `json:"row_quality" default:""`
}
