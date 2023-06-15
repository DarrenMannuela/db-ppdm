package dto

type Pden_material_bal struct {
	Pden_subtype            string   `json:"pden_subtype" default:"source"`
	Pden_id                 string   `json:"pden_id" default:"source"`
	Pden_source             string   `json:"pden_source" default:"source"`
	Product_type            string   `json:"product_type" default:"source"`
	Case_id                 string   `json:"case_id" default:"source"`
	Active_ind              *string  `json:"active_ind" default:""`
	Case_name               *string  `json:"case_name" default:""`
	Close_month             *string  `json:"close_month" default:""`
	Close_year              *int     `json:"close_year" default:""`
	Co2_percent             *float64 `json:"co_2_percent" default:""`
	Critical_press          *float64 `json:"critical_press" default:""`
	Critical_press_ouom     *string  `json:"critical_press_ouom" default:""`
	Critical_temp           *float64 `json:"critical_temp" default:""`
	Critical_temp_ouom      *string  `json:"critical_temp_ouom" default:""`
	Cum_volume              *float64 `json:"cum_volume" default:""`
	Cum_volume_date         *string  `json:"cum_volume_date" default:""`
	Cum_volume_date_desc    *string  `json:"cum_volume_date_desc" default:""`
	Cum_volume_ouom         *string  `json:"cum_volume_ouom" default:""`
	Cum_volume_uom          *string  `json:"cum_volume_uom" default:""`
	Curve_fit_error         *float64 `json:"curve_fit_error" default:""`
	Curve_fit_type          *string  `json:"curve_fit_type" default:""`
	Curve_intercept         *float64 `json:"curve_intercept" default:""`
	Curve_name              *string  `json:"curve_name" default:""`
	Curve_slope             *float64 `json:"curve_slope" default:""`
	Effective_date          *string  `json:"effective_date" default:""`
	End_date                *string  `json:"end_date" default:""`
	Expiry_date             *string  `json:"expiry_date" default:""`
	Final_cum_volume        *float64 `json:"final_cum_volume" default:""`
	Final_cum_volume_ouom   *string  `json:"final_cum_volume_ouom" default:""`
	Final_cum_volume_uom    *string  `json:"final_cum_volume_uom" default:""`
	Final_press             *float64 `json:"final_press" default:""`
	Final_press_ouom        *string  `json:"final_press_ouom" default:""`
	Gas_abandon_press       *float64 `json:"gas_abandon_press" default:""`
	Gas_abandon_press_ouom  *string  `json:"gas_abandon_press_ouom" default:""`
	Gas_abandon_recover     *float64 `json:"gas_abandon_recover" default:""`
	H2s_percent             *float64 `json:"h_2_s_percent" default:""`
	Initial_cum_volume      *float64 `json:"initial_cum_volume" default:""`
	Initial_cum_volume_ouom *string  `json:"initial_cum_volume_ouom" default:""`
	Initial_cum_volume_uom  *string  `json:"initial_cum_volume_uom" default:""`
	Initial_press           *float64 `json:"initial_press" default:""`
	Initial_press_ouom      *string  `json:"initial_press_ouom" default:""`
	Initial_temp            *float64 `json:"initial_temp" default:""`
	Initial_temp_ouom       *string  `json:"initial_temp_ouom" default:""`
	Measured_press_ind      *string  `json:"measured_press_ind" default:""`
	N2_percent              *float64 `json:"n_2_percent" default:""`
	Orig_gas_in_place       *float64 `json:"orig_gas_in_place" default:""`
	Pool_datum_depth        *float64 `json:"pool_datum_depth" default:""`
	Pool_datum_depth_ouom   *string  `json:"pool_datum_depth_ouom" default:""`
	Ppdm_guid               *string  `json:"ppdm_guid" default:""`
	Project_id              *string  `json:"project_id" default:""`
	Recov_gas_in_place      *float64 `json:"recov_gas_in_place" default:""`
	Remark                  *string  `json:"remark" default:""`
	Source                  *string  `json:"source" default:""`
	Specific_gravity        *float64 `json:"specific_gravity" default:""`
	Start_date              *string  `json:"start_date" default:""`
	Surface_loss_percent    *float64 `json:"surface_loss_percent" default:""`
	Volume                  *float64 `json:"volume" default:""`
	Volume_ouom             *string  `json:"volume_ouom" default:""`
	Volume_uom              *string  `json:"volume_uom" default:""`
	Zero_press_ind          *string  `json:"zero_press_ind" default:""`
	Z_factor_method         *string  `json:"z_factor_method" default:""`
	Row_changed_by          *string  `json:"row_changed_by" default:""`
	Row_changed_date        *string  `json:"row_changed_date" default:""`
	Row_created_by          *string  `json:"row_created_by" default:""`
	Row_created_date        *string  `json:"row_created_date" default:""`
	Row_effective_date      *string  `json:"row_effective_date" default:""`
	Row_expiry_date         *string  `json:"row_expiry_date" default:""`
	Row_quality             *string  `json:"row_quality" default:""`
}
