package dto

import (
	"time"
)

type Anl_pyrolysis struct {
	Analysis_id                 string     `json:"analysis_id"`
	Anl_source                  string     `json:"anl_source"`
	Pyrolysis_anl_obs_no        int        `json:"pyrolysis_anl_obs_no"`
	Active_ind                  *string    `json:"active_ind"`
	Calculated_ind              *string    `json:"calculated_ind"`
	Calculate_method_id         *string    `json:"calculate_method_id"`
	Effective_date              *time.Time `json:"effective_date"`
	Expiry_date                 *time.Time `json:"expiry_date"`
	Max_temperature             *float64   `json:"max_temperature"`
	Max_temperature_ouom        *string    `json:"max_temperature_ouom"`
	Max_temperature_uom         *string    `json:"max_temperature_uom"`
	Peak_temperature            *float64   `json:"peak_temperature"`
	Peak_temperature_ouom       *string    `json:"peak_temperature_ouom"`
	Peak_temperature_uom        *string    `json:"peak_temperature_uom"`
	Ppdm_guid                   string     `json:"ppdm_guid"`
	Preferred_ind               *string    `json:"preferred_ind"`
	Problem_ind                 *string    `json:"problem_ind"`
	Remark                      *string    `json:"remark"`
	Reported_ind                *string    `json:"reported_ind"`
	Rep_bitumen_index           *float64   `json:"rep_bitumen_index"`
	Rep_bitumen_index_ouom      *string    `json:"rep_bitumen_index_ouom"`
	Rep_bitumen_index_uom       *string    `json:"rep_bitumen_index_uom"`
	Rep_hydrogen_index          *float64   `json:"rep_hydrogen_index"`
	Rep_hydrogen_index_ouom     *string    `json:"rep_hydrogen_index_ouom"`
	Rep_hydrogen_index_uom      *string    `json:"rep_hydrogen_index_uom"`
	Rep_oxygen_index            *float64   `json:"rep_oxygen_index"`
	Rep_oxygen_index_ouom       *string    `json:"rep_oxygen_index_ouom"`
	Rep_oxygen_index_uom        *string    `json:"rep_oxygen_index_uom"`
	Rep_production_index        *float64   `json:"rep_production_index"`
	Rep_pyrolisable_carbon      *float64   `json:"rep_pyrolisable_carbon"`
	Rep_pyrolisable_carbon_ouom *string    `json:"rep_pyrolisable_carbon_ouom"`
	Rep_pyrolisable_carbon_uom  *string    `json:"rep_pyrolisable_carbon_uom"`
	Source                      *string    `json:"source"`
	Step_seq_no                 *int       `json:"step_seq_no"`
	S_0                         *float64   `json:"s_0"`
	S_0_ouom                    *string    `json:"s_0_ouom"`
	S_1                         *float64   `json:"s_1"`
	S_1_ouom                    *string    `json:"s_1_ouom"`
	S_2                         *float64   `json:"s_2"`
	S_2_ouom                    *string    `json:"s_2_ouom"`
	S_3                         *float64   `json:"s_3"`
	S_3_ouom                    *string    `json:"s_3_ouom"`
	S_4                         *float64   `json:"s_4"`
	S_4_ouom                    *string    `json:"s_4_ouom"`
	S_5                         *float64   `json:"s_5"`
	S_5_ouom                    *string    `json:"s_5_ouom"`
	Total_organic_carbon        *float64   `json:"total_organic_carbon"`
	Total_organic_carbon_ouom   *string    `json:"total_organic_carbon_ouom"`
	Total_organic_carbon_uom    *string    `json:"total_organic_carbon_uom"`
	Valid_anl_ind               *string    `json:"valid_anl_ind"`
	Row_changed_by              *string    `json:"row_changed_by"`
	Row_changed_date            *time.Time `json:"row_changed_date"`
	Row_created_by              *string    `json:"row_created_by"`
	Row_created_date            *time.Time `json:"row_created_date"`
	Row_effective_date          *time.Time `json:"row_effective_date"`
	Row_expiry_date             *time.Time `json:"row_expiry_date"`
	Row_quality                 *string    `json:"row_quality"`
}
