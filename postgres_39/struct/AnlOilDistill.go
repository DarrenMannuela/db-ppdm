package dto

import (
	"time"
)

type Anl_oil_distill struct {
	Analysis_id             string     `json:"analysis_id"`
	Anl_source              string     `json:"anl_source"`
	Dstl_summary_obs_no     int        `json:"dstl_summary_obs_no"`
	Active_ind              *string    `json:"active_ind"`
	Atmosp_dstl_press       *float64   `json:"atmosp_dstl_press"`
	Atmosp_dstl_press_ouom  *string    `json:"atmosp_dstl_press_ouom"`
	Atmosp_dstl_temp        *float64   `json:"atmosp_dstl_temp"`
	Atmosp_dstl_temp_ouom   *string    `json:"atmosp_dstl_temp_ouom"`
	Calculated_ind          *string    `json:"calculated_ind"`
	Calculate_method_id     *string    `json:"calculate_method_id"`
	Effective_date          *time.Time `json:"effective_date"`
	Expiry_date             *time.Time `json:"expiry_date"`
	Final_boil_pt_temp      *float64   `json:"final_boil_pt_temp"`
	Final_boil_pt_temp_ouom *string    `json:"final_boil_pt_temp_ouom"`
	Measurement_temp        *float64   `json:"measurement_temp"`
	Measurement_temp_ouom   *string    `json:"measurement_temp_ouom"`
	Ppdm_guid               string     `json:"ppdm_guid"`
	Preferred_ind           *string    `json:"preferred_ind"`
	Problem_ind             *string    `json:"problem_ind"`
	Remark                  *string    `json:"remark"`
	Reported_ind            *string    `json:"reported_ind"`
	Source                  *string    `json:"source"`
	Start_boil_pt_temp      *float64   `json:"start_boil_pt_temp"`
	Start_boil_pt_temp_ouom *string    `json:"start_boil_pt_temp_ouom"`
	Step_seq_no             *int       `json:"step_seq_no"`
	Substance_id            *string    `json:"substance_id"`
	Volume_fraction_type    *string    `json:"volume_fraction_type"`
	Vol_fraction_temp       *float64   `json:"vol_fraction_temp"`
	Vol_fraction_temp_ouom  *string    `json:"vol_fraction_temp_ouom"`
	Vol_fraction_value      *float64   `json:"vol_fraction_value"`
	Vol_fraction_value_ouom *string    `json:"vol_fraction_value_ouom"`
	Weight_cut              *float64   `json:"weight_cut"`
	Weight_cut_ouom         *string    `json:"weight_cut_ouom"`
	Row_changed_by          *string    `json:"row_changed_by"`
	Row_changed_date        *time.Time `json:"row_changed_date"`
	Row_created_by          *string    `json:"row_created_by"`
	Row_created_date        *time.Time `json:"row_created_date"`
	Row_effective_date      *time.Time `json:"row_effective_date"`
	Row_expiry_date         *time.Time `json:"row_expiry_date"`
	Row_quality             *string    `json:"row_quality"`
}
