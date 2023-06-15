package dto

import (
	"time"
)

type Pden_material_bal struct {
	Pden_subtype            string     `json:"pden_subtype"`
	Pden_id                 string     `json:"pden_id"`
	Pden_source             string     `json:"pden_source"`
	Product_type            string     `json:"product_type"`
	Case_id                 string     `json:"case_id"`
	Active_ind              *string    `json:"active_ind"`
	Case_name               *string    `json:"case_name"`
	Close_month             *string    `json:"close_month"`
	Close_year              *int       `json:"close_year"`
	Co2_percent             *float64   `json:"co_2_percent"`
	Critical_press          *float64   `json:"critical_press"`
	Critical_press_ouom     *string    `json:"critical_press_ouom"`
	Critical_temp           *float64   `json:"critical_temp"`
	Critical_temp_ouom      *string    `json:"critical_temp_ouom"`
	Cum_volume              *float64   `json:"cum_volume"`
	Cum_volume_date         *time.Time `json:"cum_volume_date"`
	Cum_volume_date_desc    *string    `json:"cum_volume_date_desc"`
	Cum_volume_ouom         *string    `json:"cum_volume_ouom"`
	Cum_volume_uom          *string    `json:"cum_volume_uom"`
	Curve_fit_error         *float64   `json:"curve_fit_error"`
	Curve_fit_type          *string    `json:"curve_fit_type"`
	Curve_intercept         *float64   `json:"curve_intercept"`
	Curve_name              *string    `json:"curve_name"`
	Curve_slope             *float64   `json:"curve_slope"`
	Effective_date          *time.Time `json:"effective_date"`
	End_date                *time.Time `json:"end_date"`
	Expiry_date             *time.Time `json:"expiry_date"`
	Final_cum_volume        *float64   `json:"final_cum_volume"`
	Final_cum_volume_ouom   *string    `json:"final_cum_volume_ouom"`
	Final_cum_volume_uom    *string    `json:"final_cum_volume_uom"`
	Final_press             *float64   `json:"final_press"`
	Final_press_ouom        *string    `json:"final_press_ouom"`
	Gas_abandon_press       *float64   `json:"gas_abandon_press"`
	Gas_abandon_press_ouom  *string    `json:"gas_abandon_press_ouom"`
	Gas_abandon_recover     *float64   `json:"gas_abandon_recover"`
	H2s_percent             *float64   `json:"h_2_s_percent"`
	Initial_cum_volume      *float64   `json:"initial_cum_volume"`
	Initial_cum_volume_ouom *string    `json:"initial_cum_volume_ouom"`
	Initial_cum_volume_uom  *string    `json:"initial_cum_volume_uom"`
	Initial_press           *float64   `json:"initial_press"`
	Initial_press_ouom      *string    `json:"initial_press_ouom"`
	Initial_temp            *float64   `json:"initial_temp"`
	Initial_temp_ouom       *string    `json:"initial_temp_ouom"`
	Measured_press_ind      *string    `json:"measured_press_ind"`
	N2_percent              *float64   `json:"n_2_percent"`
	Orig_gas_in_place       *float64   `json:"orig_gas_in_place"`
	Pool_datum_depth        *float64   `json:"pool_datum_depth"`
	Pool_datum_depth_ouom   *string    `json:"pool_datum_depth_ouom"`
	Ppdm_guid               string     `json:"ppdm_guid"`
	Project_id              *string    `json:"project_id"`
	Recov_gas_in_place      *float64   `json:"recov_gas_in_place"`
	Remark                  *string    `json:"remark"`
	Source                  *string    `json:"source"`
	Specific_gravity        *float64   `json:"specific_gravity"`
	Start_date              *time.Time `json:"start_date"`
	Surface_loss_percent    *float64   `json:"surface_loss_percent"`
	Volume                  *float64   `json:"volume"`
	Volume_ouom             *string    `json:"volume_ouom"`
	Volume_uom              *string    `json:"volume_uom"`
	Zero_press_ind          *string    `json:"zero_press_ind"`
	Z_factor_method         *string    `json:"z_factor_method"`
	Row_changed_by          *string    `json:"row_changed_by"`
	Row_changed_date        *time.Time `json:"row_changed_date"`
	Row_created_by          *string    `json:"row_created_by"`
	Row_created_date        *time.Time `json:"row_created_date"`
	Row_effective_date      *time.Time `json:"row_effective_date"`
	Row_expiry_date         *time.Time `json:"row_expiry_date"`
	Row_quality             *string    `json:"row_quality"`
}
