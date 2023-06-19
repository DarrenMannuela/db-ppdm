package dto

type Prod_string struct {
	Uwi                   string   `json:"uwi" default:"source"`
	Source                string   `json:"source" default:"source"`
	String_id             string   `json:"string_id" default:"source"`
	Active_ind            *string  `json:"active_ind" default:""`
	Base_depth            *float64 `json:"base_depth" default:""`
	Base_depth_ouom       *string  `json:"base_depth_ouom" default:""`
	Business_associate_id *string  `json:"business_associate_id" default:""`
	Commingled_ind        *string  `json:"commingled_ind" default:""`
	Current_status        *string  `json:"current_status" default:""`
	Current_status_date   *string  `json:"current_status_date" default:""`
	Effective_date        *string  `json:"effective_date" default:""`
	Expiry_date           *string  `json:"expiry_date" default:""`
	Facility_id           *string  `json:"facility_id" default:""`
	Facility_type         *string  `json:"facility_type" default:""`
	Field_id              *string  `json:"field_id" default:""`
	Government_string_id  *string  `json:"government_string_id" default:""`
	Lease_unit_id         *string  `json:"lease_unit_id" default:""`
	On_injection_date     *string  `json:"on_injection_date" default:""`
	On_production_date    *string  `json:"on_production_date" default:""`
	Plot_symbol           *string  `json:"plot_symbol" default:""`
	Pool_id               *string  `json:"pool_id" default:""`
	Ppdm_guid             *string  `json:"ppdm_guid" default:""`
	Prod_string_tvd       *float64 `json:"prod_string_tvd" default:""`
	Prod_string_tvd_ouom  *string  `json:"prod_string_tvd_ouom" default:""`
	Prod_string_type      *string  `json:"prod_string_type" default:""`
	Profile_type          *string  `json:"profile_type" default:""`
	Remark                *string  `json:"remark" default:""`
	Status_type           *string  `json:"status_type" default:""`
	Strat_name_set_id     *string  `json:"strat_name_set_id" default:""`
	Strat_unit_id         *string  `json:"strat_unit_id" default:""`
	Tax_credit_code       *string  `json:"tax_credit_code" default:""`
	Top_depth             *float64 `json:"top_depth" default:""`
	Top_depth_ouom        *string  `json:"top_depth_ouom" default:""`
	Total_depth           *float64 `json:"total_depth" default:""`
	Total_depth_ouom      *string  `json:"total_depth_ouom" default:""`
	Row_changed_by        *string  `json:"row_changed_by" default:""`
	Row_changed_date      *string  `json:"row_changed_date" default:""`
	Row_created_by        *string  `json:"row_created_by" default:""`
	Row_created_date      *string  `json:"row_created_date" default:""`
	Row_effective_date    *string  `json:"row_effective_date" default:""`
	Row_expiry_date       *string  `json:"row_expiry_date" default:""`
	Row_quality           *string  `json:"row_quality" default:""`
}