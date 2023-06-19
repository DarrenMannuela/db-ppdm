package dto

type Sf_aircraft struct {
	Sf_subtype          string   `json:"sf_subtype" default:"source"`
	Support_facility_id string   `json:"support_facility_id" default:"source"`
	Active_ind          *string  `json:"active_ind" default:""`
	Aircraft_type       *string  `json:"aircraft_type" default:""`
	Effective_date      *string  `json:"effective_date" default:""`
	Expiry_date         *string  `json:"expiry_date" default:""`
	Length              *float64 `json:"length" default:""`
	Length_ouom         *string  `json:"length_ouom" default:""`
	Ppdm_guid           *string  `json:"ppdm_guid" default:""`
	Remark              *string  `json:"remark" default:""`
	Source              *string  `json:"source" default:""`
	Row_changed_by      *string  `json:"row_changed_by" default:""`
	Row_changed_date    *string  `json:"row_changed_date" default:""`
	Row_created_by      *string  `json:"row_created_by" default:""`
	Row_created_date    *string  `json:"row_created_date" default:""`
	Row_effective_date  *string  `json:"row_effective_date" default:""`
	Row_expiry_date     *string  `json:"row_expiry_date" default:""`
	Row_quality         *string  `json:"row_quality" default:""`
}