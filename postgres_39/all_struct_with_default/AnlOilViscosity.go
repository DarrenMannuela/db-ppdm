package dto

type Anl_oil_viscosity struct {
	Analysis_id                 string   `json:"analysis_id" default:"source"`
	Anl_source                  string   `json:"anl_source" default:"source"`
	Viscosity_obs_no            int      `json:"viscosity_obs_no" default:"1"`
	Active_ind                  *string  `json:"active_ind" default:""`
	Effective_date              *string  `json:"effective_date" default:""`
	Expiry_date                 *string  `json:"expiry_date" default:""`
	Oil_temperature             *float64 `json:"oil_temperature" default:""`
	Oil_temperature_ouom        *string  `json:"oil_temperature_ouom" default:""`
	Oil_type                    *string  `json:"oil_type" default:""`
	Oil_viscosity               *float64 `json:"oil_viscosity" default:""`
	Oil_viscosity_ouom          *string  `json:"oil_viscosity_ouom" default:""`
	Oil_viscosity_uom           *string  `json:"oil_viscosity_uom" default:""`
	Pour_point_temperature      *float64 `json:"pour_point_temperature" default:""`
	Pour_point_temperature_ouom *string  `json:"pour_point_temperature_ouom" default:""`
	Ppdm_guid                   *string  `json:"ppdm_guid" default:""`
	Preferred_ind               *string  `json:"preferred_ind" default:""`
	Remark                      *string  `json:"remark" default:""`
	Source                      *string  `json:"source" default:""`
	Step_seq_no                 *int     `json:"step_seq_no" default:""`
	Viscosity_press             *float64 `json:"viscosity_press" default:""`
	Viscosity_press_ouom        *string  `json:"viscosity_press_ouom" default:""`
	Viscosity_temp              *float64 `json:"viscosity_temp" default:""`
	Viscosity_temp_ouom         *string  `json:"viscosity_temp_ouom" default:""`
	Viscosity_value             *float64 `json:"viscosity_value" default:""`
	Viscosity_value_ouom        *string  `json:"viscosity_value_ouom" default:""`
	Viscosity_value_uom         *string  `json:"viscosity_value_uom" default:""`
	Row_changed_by              *string  `json:"row_changed_by" default:""`
	Row_changed_date            *string  `json:"row_changed_date" default:""`
	Row_created_by              *string  `json:"row_created_by" default:""`
	Row_created_date            *string  `json:"row_created_date" default:""`
	Row_effective_date          *string  `json:"row_effective_date" default:""`
	Row_expiry_date             *string  `json:"row_expiry_date" default:""`
	Row_quality                 *string  `json:"row_quality" default:""`
}
