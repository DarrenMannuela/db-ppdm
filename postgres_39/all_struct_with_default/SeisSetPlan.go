package dto

type Seis_set_plan struct {
	Seis_set_subtype    string   `json:"seis_set_subtype" default:"source"`
	Seis_set_id         string   `json:"seis_set_id" default:"source"`
	Active_ind          *string  `json:"active_ind" default:""`
	Effective_date      *string  `json:"effective_date" default:""`
	Expiry_date         *string  `json:"expiry_date" default:""`
	Planned_end_date    *string  `json:"planned_end_date" default:""`
	Planned_length      *float64 `json:"planned_length" default:""`
	Planned_length_ouom *string  `json:"planned_length_ouom" default:""`
	Planned_start_date  *string  `json:"planned_start_date" default:""`
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
