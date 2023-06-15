package dto

type Sf_waste struct {
	Sf_subtype             string   `json:"sf_subtype" default:"source"`
	Waste_facility_id      string   `json:"waste_facility_id" default:"source"`
	Active_ind             *string  `json:"active_ind" default:""`
	Capacity               *float64 `json:"capacity" default:""`
	Capacity_ouom          *string  `json:"capacity_ouom" default:""`
	Current_owner          *string  `json:"current_owner" default:""`
	Effective_date         *string  `json:"effective_date" default:""`
	Expiry_date            *string  `json:"expiry_date" default:""`
	License_expiry_date    *string  `json:"license_expiry_date" default:""`
	License_num            *string  `json:"license_num" default:""`
	License_register_ba_id *string  `json:"license_register_ba_id" default:""`
	Ppdm_guid              *string  `json:"ppdm_guid" default:""`
	Remark                 *string  `json:"remark" default:""`
	Source                 *string  `json:"source" default:""`
	Waste_facility_name    *string  `json:"waste_facility_name" default:""`
	Waste_facility_type    *string  `json:"waste_facility_type" default:""`
	Row_changed_by         *string  `json:"row_changed_by" default:""`
	Row_changed_date       *string  `json:"row_changed_date" default:""`
	Row_created_by         *string  `json:"row_created_by" default:""`
	Row_created_date       *string  `json:"row_created_date" default:""`
	Row_effective_date     *string  `json:"row_effective_date" default:""`
	Row_expiry_date        *string  `json:"row_expiry_date" default:""`
	Row_quality            *string  `json:"row_quality" default:""`
}
