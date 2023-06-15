package dto

type Anl_qgf_tsf struct {
	Analysis_id              string   `json:"analysis_id" default:"source"`
	Anl_source               string   `json:"anl_source" default:"source"`
	Measurement_obs_no       int      `json:"measurement_obs_no" default:"1"`
	Active_ind               *string  `json:"active_ind" default:""`
	Effective_date           *string  `json:"effective_date" default:""`
	Emission_intensity       *float64 `json:"emission_intensity" default:""`
	Emission_intensity_uom   *string  `json:"emission_intensity_uom" default:""`
	Expiry_date              *string  `json:"expiry_date" default:""`
	Ppdm_guid                *string  `json:"ppdm_guid" default:""`
	Preferred_ind            *string  `json:"preferred_ind" default:""`
	Problem_ind              *string  `json:"problem_ind" default:""`
	Remark                   *string  `json:"remark" default:""`
	Source                   *string  `json:"source" default:""`
	Step_seq_no              *int     `json:"step_seq_no" default:""`
	Wavelength_emission      *float64 `json:"wavelength_emission" default:""`
	Wavelength_emission_uom  *string  `json:"wavelength_emission_uom" default:""`
	Wavelenth_excitation     *float64 `json:"wavelenth_excitation" default:""`
	Wavelenth_excitation_uom *string  `json:"wavelenth_excitation_uom" default:""`
	Row_changed_by           *string  `json:"row_changed_by" default:""`
	Row_changed_date         *string  `json:"row_changed_date" default:""`
	Row_created_by           *string  `json:"row_created_by" default:""`
	Row_created_date         *string  `json:"row_created_date" default:""`
	Row_effective_date       *string  `json:"row_effective_date" default:""`
	Row_expiry_date          *string  `json:"row_expiry_date" default:""`
	Row_quality              *string  `json:"row_quality" default:""`
}
