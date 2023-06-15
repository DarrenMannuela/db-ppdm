package dto

import (
	"time"
)

type Well_porous_interval struct {
	Uwi                     string     `json:"uwi"`
	Source                  string     `json:"source"`
	Porous_interval_id      string     `json:"porous_interval_id"`
	Active_ind              *string    `json:"active_ind"`
	Average_porosity        *float64   `json:"average_porosity"`
	Base_depth              *float64   `json:"base_depth"`
	Base_depth_ouom         *string    `json:"base_depth_ouom"`
	Base_tvd                *float64   `json:"base_tvd"`
	Base_tvd_ouom           *string    `json:"base_tvd_ouom"`
	Core_sample_length      *float64   `json:"core_sample_length"`
	Core_sample_length_ouom *string    `json:"core_sample_length_ouom"`
	Effective_date          *time.Time `json:"effective_date"`
	Expiry_date             *time.Time `json:"expiry_date"`
	Gas_show_ind            *string    `json:"gas_show_ind"`
	Gross_thickness         *float64   `json:"gross_thickness"`
	Gross_thickness_ouom    *string    `json:"gross_thickness_ouom"`
	Net_thickness           *float64   `json:"net_thickness"`
	Net_thickness_ouom      *string    `json:"net_thickness_ouom"`
	Oil_show_ind            *string    `json:"oil_show_ind"`
	Porosity_cutoff         *float64   `json:"porosity_cutoff"`
	Porosity_quality        *string    `json:"porosity_quality"`
	Porosity_type           *string    `json:"porosity_type"`
	Porous_form_age         *string    `json:"porous_form_age"`
	Porous_strat_unit_id    *string    `json:"porous_strat_unit_id"`
	Ppdm_guid               string     `json:"ppdm_guid"`
	Remark                  *string    `json:"remark"`
	Strat_age               *string    `json:"strat_age"`
	Strat_age_name_set_id   *string    `json:"strat_age_name_set_id"`
	Strat_name_set_id       *string    `json:"strat_name_set_id"`
	Top_depth               *float64   `json:"top_depth"`
	Top_depth_ouom          *string    `json:"top_depth_ouom"`
	Top_tvd                 *float64   `json:"top_tvd"`
	Top_tvd_ouom            *string    `json:"top_tvd_ouom"`
	Row_changed_by          *string    `json:"row_changed_by"`
	Row_changed_date        *time.Time `json:"row_changed_date"`
	Row_created_by          *string    `json:"row_created_by"`
	Row_created_date        *time.Time `json:"row_created_date"`
	Row_effective_date      *time.Time `json:"row_effective_date"`
	Row_expiry_date         *time.Time `json:"row_expiry_date"`
	Row_quality             *string    `json:"row_quality"`
}
