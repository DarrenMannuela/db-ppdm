package dto

type Well_test_flow struct {
	Uwi                       string   `json:"uwi" default:"source"`
	Source                    string   `json:"source" default:"source"`
	Test_type                 string   `json:"test_type" default:"source"`
	Run_num                   string   `json:"run_num" default:"source"`
	Test_num                  string   `json:"test_num" default:"source"`
	Period_type               string   `json:"period_type" default:"source"`
	Period_obs_no             int      `json:"period_obs_no" default:"1"`
	Fluid_type                string   `json:"fluid_type" default:"source"`
	Active_ind                *string  `json:"active_ind" default:""`
	Basic_sediment_ind        *string  `json:"basic_sediment_ind" default:""`
	Blow_desc                 *string  `json:"blow_desc" default:""`
	Effective_date            *string  `json:"effective_date" default:""`
	Expiry_date               *string  `json:"expiry_date" default:""`
	Gas_riser_diameter        *float64 `json:"gas_riser_diameter" default:""`
	Gas_riser_diameter_ouom   *string  `json:"gas_riser_diameter_ouom" default:""`
	Max_fluid_rate            *float64 `json:"max_fluid_rate" default:""`
	Max_fluid_rate_ouom       *string  `json:"max_fluid_rate_ouom" default:""`
	Max_fluid_rate_uom        *string  `json:"max_fluid_rate_uom" default:""`
	Max_surface_pressure      *float64 `json:"max_surface_pressure" default:""`
	Max_surface_pressure_ouom *string  `json:"max_surface_pressure_ouom" default:""`
	Measurement_technique     *string  `json:"measurement_technique" default:""`
	Ppdm_guid                 *string  `json:"ppdm_guid" default:""`
	Remark                    *string  `json:"remark" default:""`
	Shakeout_percent          *float64 `json:"shakeout_percent" default:""`
	Tts_time_elapsed          *float64 `json:"tts_time_elapsed" default:""`
	Tts_time_elapsed_ouom     *string  `json:"tts_time_elapsed_ouom" default:""`
	Row_changed_by            *string  `json:"row_changed_by" default:""`
	Row_changed_date          *string  `json:"row_changed_date" default:""`
	Row_created_by            *string  `json:"row_created_by" default:""`
	Row_created_date          *string  `json:"row_created_date" default:""`
	Row_effective_date        *string  `json:"row_effective_date" default:""`
	Row_expiry_date           *string  `json:"row_expiry_date" default:""`
	Row_quality               *string  `json:"row_quality" default:""`
}
