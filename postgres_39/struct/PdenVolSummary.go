package dto

import (
	"time"
)

type Pden_vol_summary struct {
	Pden_subtype              string     `json:"pden_subtype"`
	Pden_id                   string     `json:"pden_id"`
	Period_id                 string     `json:"period_id"`
	Pden_source               string     `json:"pden_source"`
	Volume_method             string     `json:"volume_method"`
	Activity_type             string     `json:"activity_type"`
	Period_type               string     `json:"period_type"`
	Amendment_seq_no          int        `json:"amendment_seq_no"`
	Active_ind                *string    `json:"active_ind"`
	Amend_reason              *string    `json:"amend_reason"`
	Boe_cum_volume            *float64   `json:"boe_cum_volume"`
	Boe_volume                *float64   `json:"boe_volume"`
	Boe_volume_ouom           *string    `json:"boe_volume_ouom"`
	Boe_ytd_volume            *float64   `json:"boe_ytd_volume"`
	Co2_cum_volume            *float64   `json:"co_2_cum_volume"`
	Co2_volume                *float64   `json:"co_2_volume"`
	Co2_volume_ouom           *string    `json:"co_2_volume_ouom"`
	Co2_ytd_volume            *float64   `json:"co_2_ytd_volume"`
	Effective_date            *time.Time `json:"effective_date"`
	Expiry_date               *time.Time `json:"expiry_date"`
	Gas_cum_volume            *float64   `json:"gas_cum_volume"`
	Gas_energy                *float64   `json:"gas_energy"`
	Gas_energy_ouom           *string    `json:"gas_energy_ouom"`
	Gas_quality               *float64   `json:"gas_quality"`
	Gas_quality_ouom          *string    `json:"gas_quality_ouom"`
	Gas_volume                *float64   `json:"gas_volume"`
	Gas_volume_ouom           *string    `json:"gas_volume_ouom"`
	Gas_ytd_volume            *float64   `json:"gas_ytd_volume"`
	Injection_cycle           *float64   `json:"injection_cycle"`
	Injection_pressure        *float64   `json:"injection_pressure"`
	Injection_pressure_ouom   *string    `json:"injection_pressure_ouom"`
	Inventory_close_balance   *float64   `json:"inventory_close_balance"`
	Inventory_open_balance    *float64   `json:"inventory_open_balance"`
	Inventory_product         *string    `json:"inventory_product"`
	Invent_close_bal_ouom     *string    `json:"invent_close_bal_ouom"`
	Invent_open_bal_ouom      *string    `json:"invent_open_bal_ouom"`
	Ngl_cum_volume            *float64   `json:"ngl_cum_volume"`
	Ngl_volume                *float64   `json:"ngl_volume"`
	Ngl_volume_ouom           *string    `json:"ngl_volume_ouom"`
	Ngl_ytd_volume            *float64   `json:"ngl_ytd_volume"`
	Nitrogen_cum_volume       *float64   `json:"nitrogen_cum_volume"`
	Nitrogen_volume           *float64   `json:"nitrogen_volume"`
	Nitrogen_volume_ouom      *string    `json:"nitrogen_volume_ouom"`
	Nitrogen_ytd_volume       *float64   `json:"nitrogen_ytd_volume"`
	No_of_gas_wells           *float64   `json:"no_of_gas_wells"`
	No_of_injection_wells     *float64   `json:"no_of_injection_wells"`
	No_of_oil_wells           *float64   `json:"no_of_oil_wells"`
	Oil_cum_volume            *float64   `json:"oil_cum_volume"`
	Oil_quality               *float64   `json:"oil_quality"`
	Oil_quality_ouom          *string    `json:"oil_quality_ouom"`
	Oil_volume                *float64   `json:"oil_volume"`
	Oil_volume_ouom           *string    `json:"oil_volume_ouom"`
	Oil_ytd_volume            *float64   `json:"oil_ytd_volume"`
	Period_on_injection       *float64   `json:"period_on_injection"`
	Period_on_injection_ouom  *string    `json:"period_on_injection_ouom"`
	Period_on_production      *float64   `json:"period_on_production"`
	Period_on_production_ouom *string    `json:"period_on_production_ouom"`
	Posted_date               *time.Time `json:"posted_date"`
	Ppdm_guid                 string     `json:"ppdm_guid"`
	Primary_allowable         *float64   `json:"primary_allowable"`
	Primary_allowable_ouom    *string    `json:"primary_allowable_ouom"`
	Primary_product           *string    `json:"primary_product"`
	Project_id                *string    `json:"project_id"`
	Remark                    *string    `json:"remark"`
	Report_ind                *string    `json:"report_ind"`
	Source                    *string    `json:"source"`
	Sulphur_cum_volume        *float64   `json:"sulphur_cum_volume"`
	Sulphur_volume            *float64   `json:"sulphur_volume"`
	Sulphur_volume_ouom       *string    `json:"sulphur_volume_ouom"`
	Sulphur_ytd_volume        *float64   `json:"sulphur_ytd_volume"`
	Volume_date               *time.Time `json:"volume_date"`
	Volume_date_desc          *string    `json:"volume_date_desc"`
	Volume_period             *float64   `json:"volume_period"`
	Volume_period_ouom        *string    `json:"volume_period_ouom"`
	Water_cum_volume          *float64   `json:"water_cum_volume"`
	Water_volume              *float64   `json:"water_volume"`
	Water_volume_ouom         *string    `json:"water_volume_ouom"`
	Water_ytd_volume          *float64   `json:"water_ytd_volume"`
	Row_changed_by            *string    `json:"row_changed_by"`
	Row_changed_date          *time.Time `json:"row_changed_date"`
	Row_created_by            *string    `json:"row_created_by"`
	Row_created_date          *time.Time `json:"row_created_date"`
	Row_effective_date        *time.Time `json:"row_effective_date"`
	Row_expiry_date           *time.Time `json:"row_expiry_date"`
	Row_quality               *string    `json:"row_quality"`
}