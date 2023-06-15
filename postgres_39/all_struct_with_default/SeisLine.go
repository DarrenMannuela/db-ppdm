package dto

type Seis_line struct {
	Seis_set_subtype         string  `json:"seis_set_subtype" default:"source"`
	Seis_line_id             string  `json:"seis_line_id" default:"source"`
	Active_ind               *string `json:"active_ind" default:""`
	Dimension                *string `json:"dimension" default:""`
	Effective_date           *string `json:"effective_date" default:""`
	Expiry_date              *string `json:"expiry_date" default:""`
	Line_name                *string `json:"line_name" default:""`
	Ppdm_guid                *string `json:"ppdm_guid" default:""`
	Refraction_reflection    *string `json:"refraction_reflection" default:""`
	Remark                   *string `json:"remark" default:""`
	Reshoot_of_set_id        *string `json:"reshoot_of_set_id" default:""`
	Reshoot_seis_set_subtype *string `json:"reshoot_seis_set_subtype" default:""`
	Seis_acqtn_set_subtype   *string `json:"seis_acqtn_set_subtype" default:""`
	Seis_acqtn_survey_id     *string `json:"seis_acqtn_survey_id" default:""`
	Seis_spectrum_type       *string `json:"seis_spectrum_type" default:""`
	Source                   *string `json:"source" default:""`
	Test_experimental        *string `json:"test_experimental" default:""`
	Row_changed_by           *string `json:"row_changed_by" default:""`
	Row_changed_date         *string `json:"row_changed_date" default:""`
	Row_created_by           *string `json:"row_created_by" default:""`
	Row_created_date         *string `json:"row_created_date" default:""`
	Row_effective_date       *string `json:"row_effective_date" default:""`
	Row_expiry_date          *string `json:"row_expiry_date" default:""`
	Row_quality              *string `json:"row_quality" default:""`
}
