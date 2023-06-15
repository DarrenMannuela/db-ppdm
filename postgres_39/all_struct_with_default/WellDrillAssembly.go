package dto

type Well_drill_assembly struct {
	Uwi                 string   `json:"uwi" default:"source"`
	Assembly_id         string   `json:"assembly_id" default:"source"`
	Active_ind          *string  `json:"active_ind" default:""`
	Assembly_ref_number *string  `json:"assembly_ref_number" default:""`
	Calculated_length   *float64 `json:"calculated_length" default:""`
	Calculated_weight   *float64 `json:"calculated_weight" default:""`
	Effective_date      *string  `json:"effective_date" default:""`
	End_time            *string  `json:"end_time" default:""`
	End_timezone        *string  `json:"end_timezone" default:""`
	Expiry_date         *string  `json:"expiry_date" default:""`
	Ppdm_guid           *string  `json:"ppdm_guid" default:""`
	Remark              *string  `json:"remark" default:""`
	Source              *string  `json:"source" default:""`
	Start_time          *string  `json:"start_time" default:""`
	Start_timezone      *string  `json:"start_timezone" default:""`
	Row_changed_by      *string  `json:"row_changed_by" default:""`
	Row_changed_date    *string  `json:"row_changed_date" default:""`
	Row_created_by      *string  `json:"row_created_by" default:""`
	Row_created_date    *string  `json:"row_created_date" default:""`
	Row_effective_date  *string  `json:"row_effective_date" default:""`
	Row_expiry_date     *string  `json:"row_expiry_date" default:""`
	Row_quality         *string  `json:"row_quality" default:""`
}
