package dto

type Prod_lease_unit struct {
	Lease_unit_id            string   `json:"lease_unit_id" default:"source"`
	Acreage                  *float64 `json:"acreage" default:""`
	Acreage_ouom             *string  `json:"acreage_ouom" default:""`
	Active_ind               *string  `json:"active_ind" default:""`
	Area_id                  *string  `json:"area_id" default:""`
	Area_type                *string  `json:"area_type" default:""`
	Current_operator         *string  `json:"current_operator" default:""`
	Current_status_date      *string  `json:"current_status_date" default:""`
	Effective_date           *string  `json:"effective_date" default:""`
	Expiry_date              *string  `json:"expiry_date" default:""`
	Field_id                 *string  `json:"field_id" default:""`
	Government_lease_unit_id *string  `json:"government_lease_unit_id" default:""`
	Land_right_id            *string  `json:"land_right_id" default:""`
	Land_right_subtype       *string  `json:"land_right_subtype" default:""`
	Lease_unit_long_name     *string  `json:"lease_unit_long_name" default:""`
	Lease_unit_short_name    *string  `json:"lease_unit_short_name" default:""`
	Lease_unit_status        *string  `json:"lease_unit_status" default:""`
	Lease_unit_type          *string  `json:"lease_unit_type" default:""`
	Pool_id                  *string  `json:"pool_id" default:""`
	Ppdm_guid                *string  `json:"ppdm_guid" default:""`
	Remark                   *string  `json:"remark" default:""`
	Source                   *string  `json:"source" default:""`
	Strat_name_set_id        *string  `json:"strat_name_set_id" default:""`
	Strat_unit_id            *string  `json:"strat_unit_id" default:""`
	Row_changed_by           *string  `json:"row_changed_by" default:""`
	Row_changed_date         *string  `json:"row_changed_date" default:""`
	Row_created_by           *string  `json:"row_created_by" default:""`
	Row_created_date         *string  `json:"row_created_date" default:""`
	Row_effective_date       *string  `json:"row_effective_date" default:""`
	Row_expiry_date          *string  `json:"row_expiry_date" default:""`
	Row_quality              *string  `json:"row_quality" default:""`
}
