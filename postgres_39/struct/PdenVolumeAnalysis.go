package dto

import (
	"time"
)

type Pden_volume_analysis struct {
	Pden_subtype                string     `json:"pden_subtype"`
	Pden_id                     string     `json:"pden_id"`
	Pden_source                 string     `json:"pden_source"`
	Product_type                string     `json:"product_type"`
	Case_id                     string     `json:"case_id"`
	Active_ind                  *string    `json:"active_ind"`
	Area_size                   *float64   `json:"area_size"`
	Area_size_ouom              *string    `json:"area_size_ouom"`
	Case_name                   *string    `json:"case_name"`
	Date_format_desc            *time.Time `json:"date_format_desc"`
	Effective_date              *time.Time `json:"effective_date"`
	End_date                    *time.Time `json:"end_date"`
	Expiry_date                 *time.Time `json:"expiry_date"`
	Gas_abandon_compress        *float64   `json:"gas_abandon_compress"`
	Gas_abandon_press           *float64   `json:"gas_abandon_press"`
	Gas_abandon_press_ouom      *string    `json:"gas_abandon_press_ouom"`
	Gas_init_compress           *float64   `json:"gas_init_compress"`
	Gas_init_pressure           *float64   `json:"gas_init_pressure"`
	Gas_in_place                *float64   `json:"gas_in_place"`
	Gas_in_place_ouom           *string    `json:"gas_in_place_ouom"`
	Gas_original_in_place       *float64   `json:"gas_original_in_place"`
	Gas_original_in_place_ouom  *string    `json:"gas_original_in_place_ouom"`
	Gas_ratio_bgi               *float64   `json:"gas_ratio_bgi"`
	Gas_recovery                *float64   `json:"gas_recovery"`
	Gross_net_pay_ratio         *float64   `json:"gross_net_pay_ratio"`
	Gross_pay                   *float64   `json:"gross_pay"`
	Gross_pay_ouom              *string    `json:"gross_pay_ouom"`
	Init_res_temp               *float64   `json:"init_res_temp"`
	Init_res_temp_ouom          *string    `json:"init_res_temp_ouom"`
	Net_pay                     *float64   `json:"net_pay"`
	Net_pay_ouom                *string    `json:"net_pay_ouom"`
	Oil_in_place                *float64   `json:"oil_in_place"`
	Oil_in_place_ouom           *string    `json:"oil_in_place_ouom"`
	Oil_original_in_place       *float64   `json:"oil_original_in_place"`
	Oil_original_in_place_ouom  *string    `json:"oil_original_in_place_ouom"`
	Oil_recovery_primary        *float64   `json:"oil_recovery_primary"`
	Oil_recovery_secondary      *float64   `json:"oil_recovery_secondary"`
	Oil_recovery_total          *float64   `json:"oil_recovery_total"`
	Oil_residual_sat            *float64   `json:"oil_residual_sat"`
	Oil_shrinkage               *float64   `json:"oil_shrinkage"`
	Original_gor                *float64   `json:"original_gor"`
	Original_gor_ouom           *string    `json:"original_gor_ouom"`
	Orig_sol_gas_in_place       *float64   `json:"orig_sol_gas_in_place"`
	Orig_sol_gas_in_place_ouom  *string    `json:"orig_sol_gas_in_place_ouom"`
	Permeability                *float64   `json:"permeability"`
	Permeability_ouom           *string    `json:"permeability_ouom"`
	Porosity                    *float64   `json:"porosity"`
	Ppdm_guid                   string     `json:"ppdm_guid"`
	Project_id                  *string    `json:"project_id"`
	Recov_gor                   *float64   `json:"recov_gor"`
	Recov_sol_gas_in_place      *float64   `json:"recov_sol_gas_in_place"`
	Recov_sol_gas_in_place_ouom *string    `json:"recov_sol_gas_in_place_ouom"`
	Remark                      *string    `json:"remark"`
	Sol_gas_recovery            *float64   `json:"sol_gas_recovery"`
	Source                      *string    `json:"source"`
	Start_date                  *time.Time `json:"start_date"`
	Volume                      *float64   `json:"volume"`
	Volume_ouom                 *string    `json:"volume_ouom"`
	Volume_uom                  *string    `json:"volume_uom"`
	Water_saturation            *float64   `json:"water_saturation"`
	Row_changed_by              *string    `json:"row_changed_by"`
	Row_changed_date            *time.Time `json:"row_changed_date"`
	Row_created_by              *string    `json:"row_created_by"`
	Row_created_date            *time.Time `json:"row_created_date"`
	Row_effective_date          *time.Time `json:"row_effective_date"`
	Row_expiry_date             *time.Time `json:"row_expiry_date"`
	Row_quality                 *string    `json:"row_quality"`
}
