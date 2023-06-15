package dto

import (
	"time"
)

type Well_core struct {
	Uwi                        string     `json:"uwi"`
	Source                     string     `json:"source"`
	Core_id                    string     `json:"core_id"`
	Active_ind                 *string    `json:"active_ind"`
	Base_depth                 *float64   `json:"base_depth"`
	Base_depth_ouom            *string    `json:"base_depth_ouom"`
	Contractor                 *string    `json:"contractor"`
	Core_barrel_size           *float64   `json:"core_barrel_size"`
	Core_barrel_size_ouom      *string    `json:"core_barrel_size_ouom"`
	Core_diameter              *float64   `json:"core_diameter"`
	Core_diameter_ouom         *string    `json:"core_diameter_ouom"`
	Core_handling_type         *string    `json:"core_handling_type"`
	Core_oriented_ind          *string    `json:"core_oriented_ind"`
	Core_show_type             *string    `json:"core_show_type"`
	Core_type                  *string    `json:"core_type"`
	Coring_fluid               *string    `json:"coring_fluid"`
	Digit_avail_ind            *string    `json:"digit_avail_ind"`
	Effective_date             *time.Time `json:"effective_date"`
	Expiry_date                *time.Time `json:"expiry_date"`
	Gamma_correlation_ind      *string    `json:"gamma_correlation_ind"`
	Operation_seq_no           *int       `json:"operation_seq_no"`
	Ppdm_guid                  string     `json:"ppdm_guid"`
	Primary_core_strat_unit_id *string    `json:"primary_core_strat_unit_id"`
	Recovered_amount           *float64   `json:"recovered_amount"`
	Recovered_amount_ouom      *string    `json:"recovered_amount_ouom"`
	Recovered_amount_uom       *string    `json:"recovered_amount_uom"`
	Recovery_date              *time.Time `json:"recovery_date"`
	Remark                     *string    `json:"remark"`
	Reported_core_num          *string    `json:"reported_core_num"`
	Run_num                    *string    `json:"run_num"`
	Shot_recovered_count       *int       `json:"shot_recovered_count"`
	Sidewall_ind               *string    `json:"sidewall_ind"`
	Strat_name_set_id          *string    `json:"strat_name_set_id"`
	Top_depth                  *float64   `json:"top_depth"`
	Top_depth_ouom             *string    `json:"top_depth_ouom"`
	Total_shot_count           *int       `json:"total_shot_count"`
	Row_changed_by             *string    `json:"row_changed_by"`
	Row_changed_date           *time.Time `json:"row_changed_date"`
	Row_created_by             *string    `json:"row_created_by"`
	Row_created_date           *time.Time `json:"row_created_date"`
	Row_effective_date         *time.Time `json:"row_effective_date"`
	Row_expiry_date            *time.Time `json:"row_expiry_date"`
	Row_quality                *string    `json:"row_quality"`
}
