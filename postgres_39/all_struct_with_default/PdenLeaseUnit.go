package dto

type Pden_lease_unit struct {
	Pden_subtype          string   `json:"pden_subtype" default:"source"`
	Pden_id               string   `json:"pden_id" default:"source"`
	Pden_source           string   `json:"pden_source" default:"source"`
	Active_ind            *string  `json:"active_ind" default:""`
	Effective_date        *string  `json:"effective_date" default:""`
	Expiry_date           *string  `json:"expiry_date" default:""`
	Facility_id           *string  `json:"facility_id" default:""`
	Facility_type         *string  `json:"facility_type" default:""`
	Lease_unit_id         *string  `json:"lease_unit_id" default:""`
	No_of_gas_wells       *float64 `json:"no_of_gas_wells" default:""`
	No_of_injection_wells *float64 `json:"no_of_injection_wells" default:""`
	No_of_oil_wells       *float64 `json:"no_of_oil_wells" default:""`
	Ppdm_guid             *string  `json:"ppdm_guid" default:""`
	Remark                *string  `json:"remark" default:""`
	Row_changed_by        *string  `json:"row_changed_by" default:""`
	Row_changed_date      *string  `json:"row_changed_date" default:""`
	Row_created_by        *string  `json:"row_created_by" default:""`
	Row_created_date      *string  `json:"row_created_date" default:""`
	Row_effective_date    *string  `json:"row_effective_date" default:""`
	Row_expiry_date       *string  `json:"row_expiry_date" default:""`
	Row_quality           *string  `json:"row_quality" default:""`
}
