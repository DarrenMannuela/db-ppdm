package dto

import (
	"time"
)

type Prod_string struct {
	Uwi                   string     `json:"uwi"`
	Source                string     `json:"source"`
	String_id             string     `json:"string_id"`
	Active_ind            *string    `json:"active_ind"`
	Base_depth            *float64   `json:"base_depth"`
	Base_depth_ouom       *string    `json:"base_depth_ouom"`
	Business_associate_id *string    `json:"business_associate_id"`
	Commingled_ind        *string    `json:"commingled_ind"`
	Current_status        *string    `json:"current_status"`
	Current_status_date   *time.Time `json:"current_status_date"`
	Effective_date        *time.Time `json:"effective_date"`
	Expiry_date           *time.Time `json:"expiry_date"`
	Facility_id           *string    `json:"facility_id"`
	Facility_type         *string    `json:"facility_type"`
	Field_id              *string    `json:"field_id"`
	Government_string_id  *string    `json:"government_string_id"`
	Lease_unit_id         *string    `json:"lease_unit_id"`
	On_injection_date     *time.Time `json:"on_injection_date"`
	On_production_date    *time.Time `json:"on_production_date"`
	Plot_symbol           *string    `json:"plot_symbol"`
	Pool_id               *string    `json:"pool_id"`
	Ppdm_guid             string     `json:"ppdm_guid"`
	Prod_string_tvd       *float64   `json:"prod_string_tvd"`
	Prod_string_tvd_ouom  *string    `json:"prod_string_tvd_ouom"`
	Prod_string_type      *string    `json:"prod_string_type"`
	Profile_type          *string    `json:"profile_type"`
	Remark                *string    `json:"remark"`
	Status_type           *string    `json:"status_type"`
	Strat_name_set_id     *string    `json:"strat_name_set_id"`
	Strat_unit_id         *string    `json:"strat_unit_id"`
	Tax_credit_code       *string    `json:"tax_credit_code"`
	Top_depth             *float64   `json:"top_depth"`
	Top_depth_ouom        *string    `json:"top_depth_ouom"`
	Total_depth           *float64   `json:"total_depth"`
	Total_depth_ouom      *string    `json:"total_depth_ouom"`
	Row_changed_by        *string    `json:"row_changed_by"`
	Row_changed_date      *time.Time `json:"row_changed_date"`
	Row_created_by        *string    `json:"row_created_by"`
	Row_created_date      *time.Time `json:"row_created_date"`
	Row_effective_date    *time.Time `json:"row_effective_date"`
	Row_expiry_date       *time.Time `json:"row_expiry_date"`
	Row_quality           *string    `json:"row_quality"`
}
