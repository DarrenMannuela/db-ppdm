package dto

import (
	"time"
)

type Well_payzone struct {
	Uwi                        string     `json:"uwi"`
	Source                     string     `json:"source"`
	Zone_id                    string     `json:"zone_id"`
	Zone_source                string     `json:"zone_source"`
	Payzone_type               string     `json:"payzone_type"`
	Fluid_type                 string     `json:"fluid_type"`
	Active_ind                 *string    `json:"active_ind"`
	Base_depth                 *float64   `json:"base_depth"`
	Base_depth_ouom            *string    `json:"base_depth_ouom"`
	Base_strat_unit_id         *string    `json:"base_strat_unit_id"`
	Effective_date             *time.Time `json:"effective_date"`
	Expiry_date                *time.Time `json:"expiry_date"`
	Gas_oil_contact_depth      *float64   `json:"gas_oil_contact_depth"`
	Gas_oil_contact_depth_ouom *string    `json:"gas_oil_contact_depth_ouom"`
	Gas_wtr_contact_depth      *float64   `json:"gas_wtr_contact_depth"`
	Gas_wtr_contact_depth_ouom *string    `json:"gas_wtr_contact_depth_ouom"`
	Gross_pay                  *float64   `json:"gross_pay"`
	Gross_pay_ouom             *string    `json:"gross_pay_ouom"`
	Net_pay                    *float64   `json:"net_pay"`
	Net_pay_ouom               *string    `json:"net_pay_ouom"`
	Oil_wtr_contact_depth      *float64   `json:"oil_wtr_contact_depth"`
	Oil_wtr_contact_depth_ouom *string    `json:"oil_wtr_contact_depth_ouom"`
	Ppdm_guid                  string     `json:"ppdm_guid"`
	Remark                     *string    `json:"remark"`
	Strat_name_set_id          *string    `json:"strat_name_set_id"`
	Top_depth                  *float64   `json:"top_depth"`
	Top_depth_ouom             *string    `json:"top_depth_ouom"`
	Top_strat_unit_id          *string    `json:"top_strat_unit_id"`
	Row_changed_by             *string    `json:"row_changed_by"`
	Row_changed_date           *time.Time `json:"row_changed_date"`
	Row_created_by             *string    `json:"row_created_by"`
	Row_created_date           *time.Time `json:"row_created_date"`
	Row_effective_date         *time.Time `json:"row_effective_date"`
	Row_expiry_date            *time.Time `json:"row_expiry_date"`
	Row_quality                *string    `json:"row_quality"`
}
