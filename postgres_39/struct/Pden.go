package dto

import (
	"time"
)

type Pden struct {
	Pden_subtype            string     `json:"pden_subtype"`
	Pden_id                 string     `json:"pden_id"`
	Source                  string     `json:"source"`
	Active_ind              *string    `json:"active_ind"`
	Area_id                 *string    `json:"area_id"`
	Area_type               *string    `json:"area_type"`
	Current_operator        *string    `json:"current_operator"`
	Current_prod_str_name   *string    `json:"current_prod_str_name"`
	Current_status_date     *time.Time `json:"current_status_date"`
	Current_well_str_number *string    `json:"current_well_str_number"`
	Effective_date          *time.Time `json:"effective_date"`
	Enhanced_recovery_type  *string    `json:"enhanced_recovery_type"`
	Expiry_date             *time.Time `json:"expiry_date"`
	Field_id                *string    `json:"field_id"`
	Last_injection_date     *time.Time `json:"last_injection_date"`
	Last_production_date    *time.Time `json:"last_production_date"`
	Last_reported_date      *time.Time `json:"last_reported_date"`
	Location_desc           *string    `json:"location_desc"`
	Location_desc_type      *string    `json:"location_desc_type"`
	On_injection_date       *time.Time `json:"on_injection_date"`
	On_production_date      *time.Time `json:"on_production_date"`
	Pden_long_name          *string    `json:"pden_long_name"`
	Pden_short_name         *string    `json:"pden_short_name"`
	Pden_status             *string    `json:"pden_status"`
	Pden_status_type        *string    `json:"pden_status_type"`
	Plot_name               *string    `json:"plot_name"`
	Plot_symbol             *string    `json:"plot_symbol"`
	Pool_id                 *string    `json:"pool_id"`
	Ppdm_guid               string     `json:"ppdm_guid"`
	Primary_product         *string    `json:"primary_product"`
	Production_method       *string    `json:"production_method"`
	Proprietary_ind         *string    `json:"proprietary_ind"`
	Remark                  *string    `json:"remark"`
	State_or_federal_waters *string    `json:"state_or_federal_waters"`
	Strat_name_set_id       *string    `json:"strat_name_set_id"`
	Strat_unit_id           *string    `json:"strat_unit_id"`
	String_serial_number    *string    `json:"string_serial_number"`
	Row_changed_by          *string    `json:"row_changed_by"`
	Row_changed_date        *time.Time `json:"row_changed_date"`
	Row_created_by          *string    `json:"row_created_by"`
	Row_created_date        *time.Time `json:"row_created_date"`
	Row_effective_date      *time.Time `json:"row_effective_date"`
	Row_expiry_date         *time.Time `json:"row_expiry_date"`
	Row_quality             *string    `json:"row_quality"`
}
