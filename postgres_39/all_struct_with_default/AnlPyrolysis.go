package dto

type Anl_pyrolysis struct {
	Analysis_id                 string   `json:"analysis_id" default:"source"`
	Anl_source                  string   `json:"anl_source" default:"source"`
	Pyrolysis_anl_obs_no        int      `json:"pyrolysis_anl_obs_no" default:"1"`
	Active_ind                  *string  `json:"active_ind" default:""`
	Calculated_ind              *string  `json:"calculated_ind" default:""`
	Calculate_method_id         *string  `json:"calculate_method_id" default:""`
	Effective_date              *string  `json:"effective_date" default:""`
	Expiry_date                 *string  `json:"expiry_date" default:""`
	Max_temperature             *float64 `json:"max_temperature" default:""`
	Max_temperature_ouom        *string  `json:"max_temperature_ouom" default:""`
	Max_temperature_uom         *string  `json:"max_temperature_uom" default:""`
	Peak_temperature            *float64 `json:"peak_temperature" default:""`
	Peak_temperature_ouom       *string  `json:"peak_temperature_ouom" default:""`
	Peak_temperature_uom        *string  `json:"peak_temperature_uom" default:""`
	Ppdm_guid                   *string  `json:"ppdm_guid" default:""`
	Preferred_ind               *string  `json:"preferred_ind" default:""`
	Problem_ind                 *string  `json:"problem_ind" default:""`
	Remark                      *string  `json:"remark" default:""`
	Reported_ind                *string  `json:"reported_ind" default:""`
	Rep_bitumen_index           *float64 `json:"rep_bitumen_index" default:""`
	Rep_bitumen_index_ouom      *string  `json:"rep_bitumen_index_ouom" default:""`
	Rep_bitumen_index_uom       *string  `json:"rep_bitumen_index_uom" default:""`
	Rep_hydrogen_index          *float64 `json:"rep_hydrogen_index" default:""`
	Rep_hydrogen_index_ouom     *string  `json:"rep_hydrogen_index_ouom" default:""`
	Rep_hydrogen_index_uom      *string  `json:"rep_hydrogen_index_uom" default:""`
	Rep_oxygen_index            *float64 `json:"rep_oxygen_index" default:""`
	Rep_oxygen_index_ouom       *string  `json:"rep_oxygen_index_ouom" default:""`
	Rep_oxygen_index_uom        *string  `json:"rep_oxygen_index_uom" default:""`
	Rep_production_index        *float64 `json:"rep_production_index" default:""`
	Rep_pyrolisable_carbon      *float64 `json:"rep_pyrolisable_carbon" default:""`
	Rep_pyrolisable_carbon_ouom *string  `json:"rep_pyrolisable_carbon_ouom" default:""`
	Rep_pyrolisable_carbon_uom  *string  `json:"rep_pyrolisable_carbon_uom" default:""`
	Source                      *string  `json:"source" default:""`
	Step_seq_no                 *int     `json:"step_seq_no" default:""`
	S_0                         *float64 `json:"s_0" default:""`
	S_0_ouom                    *string  `json:"s_0_ouom" default:""`
	S_1                         *float64 `json:"s_1" default:""`
	S_1_ouom                    *string  `json:"s_1_ouom" default:""`
	S_2                         *float64 `json:"s_2" default:""`
	S_2_ouom                    *string  `json:"s_2_ouom" default:""`
	S_3                         *float64 `json:"s_3" default:""`
	S_3_ouom                    *string  `json:"s_3_ouom" default:""`
	S_4                         *float64 `json:"s_4" default:""`
	S_4_ouom                    *string  `json:"s_4_ouom" default:""`
	S_5                         *float64 `json:"s_5" default:""`
	S_5_ouom                    *string  `json:"s_5_ouom" default:""`
	Total_organic_carbon        *float64 `json:"total_organic_carbon" default:""`
	Total_organic_carbon_ouom   *string  `json:"total_organic_carbon_ouom" default:""`
	Total_organic_carbon_uom    *string  `json:"total_organic_carbon_uom" default:""`
	Valid_anl_ind               *string  `json:"valid_anl_ind" default:""`
	Row_changed_by              *string  `json:"row_changed_by" default:""`
	Row_changed_date            *string  `json:"row_changed_date" default:""`
	Row_created_by              *string  `json:"row_created_by" default:""`
	Row_created_date            *string  `json:"row_created_date" default:""`
	Row_effective_date          *string  `json:"row_effective_date" default:""`
	Row_expiry_date             *string  `json:"row_expiry_date" default:""`
	Row_quality                 *string  `json:"row_quality" default:""`
}
