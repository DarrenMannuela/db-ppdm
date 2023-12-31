package dto

type Anl_water_analysis struct {
	Analysis_id                string   `json:"analysis_id" default:"source"`
	Anl_source                 string   `json:"anl_source" default:"source"`
	Water_analysis_obs_no      int      `json:"water_analysis_obs_no" default:"1"`
	Active_ind                 *string  `json:"active_ind" default:""`
	Barium                     *float64 `json:"barium" default:""`
	Barium_ouom                *string  `json:"barium_ouom" default:""`
	Bicarbonate                *float64 `json:"bicarbonate" default:""`
	Bicarbonate_ouom           *string  `json:"bicarbonate_ouom" default:""`
	Boron                      *float64 `json:"boron" default:""`
	Boron_ouom                 *string  `json:"boron_ouom" default:""`
	Bromine                    *float64 `json:"bromine" default:""`
	Bromine_ouom               *string  `json:"bromine_ouom" default:""`
	Calcium                    *float64 `json:"calcium" default:""`
	Calcium_ouom               *string  `json:"calcium_ouom" default:""`
	Carbonate                  *float64 `json:"carbonate" default:""`
	Carbonate_ouom             *string  `json:"carbonate_ouom" default:""`
	Chlorine                   *float64 `json:"chlorine" default:""`
	Chlorine_ouom              *string  `json:"chlorine_ouom" default:""`
	Conductivity               *float64 `json:"conductivity" default:""`
	Conductivity_ouom          *string  `json:"conductivity_ouom" default:""`
	Effective_date             *string  `json:"effective_date" default:""`
	Expiry_date                *string  `json:"expiry_date" default:""`
	Hydrogen_sulphide          *float64 `json:"hydrogen_sulphide" default:""`
	Hydrogen_sulphide_ouom     *string  `json:"hydrogen_sulphide_ouom" default:""`
	Hydroxide                  *float64 `json:"hydroxide" default:""`
	Hydroxide_ouom             *string  `json:"hydroxide_ouom" default:""`
	Iodine                     *float64 `json:"iodine" default:""`
	Iodine_ouom                *string  `json:"iodine_ouom" default:""`
	Ion_ouom                   *string  `json:"ion_ouom" default:""`
	Ion_uom                    *string  `json:"ion_uom" default:""`
	Iron                       *float64 `json:"iron" default:""`
	Iron_ouom                  *string  `json:"iron_ouom" default:""`
	Magnesium                  *float64 `json:"magnesium" default:""`
	Magnesium_ouom             *string  `json:"magnesium_ouom" default:""`
	Organics_desc              *string  `json:"organics_desc" default:""`
	Ph_observed                *float64 `json:"ph_observed" default:""`
	Ph_temp                    *float64 `json:"ph_temp" default:""`
	Ph_temp_ouom               *string  `json:"ph_temp_ouom" default:""`
	Potassium                  *float64 `json:"potassium" default:""`
	Potassium_ouom             *string  `json:"potassium_ouom" default:""`
	Ppdm_guid                  *string  `json:"ppdm_guid" default:""`
	Preferred_ind              *string  `json:"preferred_ind" default:""`
	Problem_ind                *string  `json:"problem_ind" default:""`
	Recovery_desc              *string  `json:"recovery_desc" default:""`
	Refractive_index           *float64 `json:"refractive_index" default:""`
	Refractive_index_temp      *float64 `json:"refractive_index_temp" default:""`
	Refractive_index_temp_ouom *string  `json:"refractive_index_temp_ouom" default:""`
	Remark                     *string  `json:"remark" default:""`
	Reservoir_temperature      *float64 `json:"reservoir_temperature" default:""`
	Reservoir_temperature_ouom *string  `json:"reservoir_temperature_ouom" default:""`
	Rw_temperature             *float64 `json:"rw_temperature" default:""`
	Rw_temperature_ouom        *string  `json:"rw_temperature_ouom" default:""`
	Sample_water_type          *string  `json:"sample_water_type" default:""`
	Sodium                     *float64 `json:"sodium" default:""`
	Sodium_ouom                *string  `json:"sodium_ouom" default:""`
	Sodium_potassium           *float64 `json:"sodium_potassium" default:""`
	Sodium_potassium_ouom      *string  `json:"sodium_potassium_ouom" default:""`
	Source                     *string  `json:"source" default:""`
	Step_seq_no                *int     `json:"step_seq_no" default:""`
	Strontium                  *float64 `json:"strontium" default:""`
	Strontium_ouom             *string  `json:"strontium_ouom" default:""`
	Sulphate                   *float64 `json:"sulphate" default:""`
	Sulphate_ouom              *string  `json:"sulphate_ouom" default:""`
	Water_density              *float64 `json:"water_density" default:""`
	Water_density_ouom         *string  `json:"water_density_ouom" default:""`
	Water_ph                   *float64 `json:"water_ph" default:""`
	Water_resistivity          *float64 `json:"water_resistivity" default:""`
	Water_resistivity_ouom     *string  `json:"water_resistivity_ouom" default:""`
	Water_type                 *string  `json:"water_type" default:""`
	Row_changed_by             *string  `json:"row_changed_by" default:""`
	Row_changed_date           *string  `json:"row_changed_date" default:""`
	Row_created_by             *string  `json:"row_created_by" default:""`
	Row_created_date           *string  `json:"row_created_date" default:""`
	Row_effective_date         *string  `json:"row_effective_date" default:""`
	Row_expiry_date            *string  `json:"row_expiry_date" default:""`
	Row_quality                *string  `json:"row_quality" default:""`
}
