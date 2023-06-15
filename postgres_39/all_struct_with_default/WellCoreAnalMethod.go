package dto

type Well_core_anal_method struct {
	Uwi                       string   `json:"uwi" default:"source"`
	Source                    string   `json:"source" default:"source"`
	Core_id                   string   `json:"core_id" default:"source"`
	Analysis_obs_no           int      `json:"analysis_obs_no" default:"1"`
	Active_ind                *string  `json:"active_ind" default:""`
	Core_solvent              *string  `json:"core_solvent" default:""`
	Cutting_fluid             *string  `json:"cutting_fluid" default:""`
	Density_analysis          *string  `json:"density_analysis" default:""`
	Drying_method             *string  `json:"drying_method" default:""`
	Drying_temperature        *float64 `json:"drying_temperature" default:""`
	Drying_temperature_ouom   *string  `json:"drying_temperature_ouom" default:""`
	Drying_time_elapsed       *float64 `json:"drying_time_elapsed" default:""`
	Drying_time_elapsed_ouom  *string  `json:"drying_time_elapsed_ouom" default:""`
	Effective_date            *string  `json:"effective_date" default:""`
	Expiry_date               *string  `json:"expiry_date" default:""`
	Extract_method            *string  `json:"extract_method" default:""`
	Extract_time_elapsed      *float64 `json:"extract_time_elapsed" default:""`
	Extract_time_elapsed_ouom *string  `json:"extract_time_elapsed_ouom" default:""`
	Fluid_sat_analysis        *string  `json:"fluid_sat_analysis" default:""`
	Permeability_analysis     *string  `json:"permeability_analysis" default:""`
	Porosity_analysis         *string  `json:"porosity_analysis" default:""`
	Ppdm_guid                 *string  `json:"ppdm_guid" default:""`
	Remark                    *string  `json:"remark" default:""`
	Row_changed_by            *string  `json:"row_changed_by" default:""`
	Row_changed_date          *string  `json:"row_changed_date" default:""`
	Row_created_by            *string  `json:"row_created_by" default:""`
	Row_created_date          *string  `json:"row_created_date" default:""`
	Row_effective_date        *string  `json:"row_effective_date" default:""`
	Row_expiry_date           *string  `json:"row_expiry_date" default:""`
	Row_quality               *string  `json:"row_quality" default:""`
}
