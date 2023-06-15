package dto

type Sf_xref struct {
	Sf_subtype           string   `json:"sf_subtype" default:"source"`
	Support_facility_id  string   `json:"support_facility_id" default:"source"`
	Sf_subtype2          string   `json:"sf_subtype_2" default:"source"`
	Support_facility_id2 string   `json:"support_facility_id_2" default:"source"`
	Active_ind           *string  `json:"active_ind" default:""`
	Effective_date       *string  `json:"effective_date" default:""`
	Expiry_date          *string  `json:"expiry_date" default:""`
	From_distance        *float64 `json:"from_distance" default:""`
	From_distance_ouom   *string  `json:"from_distance_ouom" default:""`
	Portion_percent      *float64 `json:"portion_percent" default:""`
	Ppdm_guid            *string  `json:"ppdm_guid" default:""`
	Remark               *string  `json:"remark" default:""`
	Source               *string  `json:"source" default:""`
	To_distance          *float64 `json:"to_distance" default:""`
	To_distance_ouom     *string  `json:"to_distance_ouom" default:""`
	Xref_type            *string  `json:"xref_type" default:""`
	Row_changed_by       *string  `json:"row_changed_by" default:""`
	Row_changed_date     *string  `json:"row_changed_date" default:""`
	Row_created_by       *string  `json:"row_created_by" default:""`
	Row_created_date     *string  `json:"row_created_date" default:""`
	Row_effective_date   *string  `json:"row_effective_date" default:""`
	Row_expiry_date      *string  `json:"row_expiry_date" default:""`
	Row_quality          *string  `json:"row_quality" default:""`
}
