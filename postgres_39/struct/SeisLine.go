package dto

import (
	"time"
)

type Seis_line struct {
	Seis_set_subtype         string     `json:"seis_set_subtype"`
	Seis_line_id             string     `json:"seis_line_id"`
	Active_ind               *string    `json:"active_ind"`
	Dimension                *string    `json:"dimension"`
	Effective_date           *time.Time `json:"effective_date"`
	Expiry_date              *time.Time `json:"expiry_date"`
	Line_name                *string    `json:"line_name"`
	Ppdm_guid                string     `json:"ppdm_guid"`
	Refraction_reflection    *string    `json:"refraction_reflection"`
	Remark                   *string    `json:"remark"`
	Reshoot_of_set_id        *string    `json:"reshoot_of_set_id"`
	Reshoot_seis_set_subtype *string    `json:"reshoot_seis_set_subtype"`
	Seis_acqtn_set_subtype   *string    `json:"seis_acqtn_set_subtype"`
	Seis_acqtn_survey_id     *string    `json:"seis_acqtn_survey_id"`
	Seis_spectrum_type       *string    `json:"seis_spectrum_type"`
	Source                   *string    `json:"source"`
	Test_experimental        *string    `json:"test_experimental"`
	Row_changed_by           *string    `json:"row_changed_by"`
	Row_changed_date         *time.Time `json:"row_changed_date"`
	Row_created_by           *string    `json:"row_created_by"`
	Row_created_date         *time.Time `json:"row_created_date"`
	Row_effective_date       *time.Time `json:"row_effective_date"`
	Row_expiry_date          *time.Time `json:"row_expiry_date"`
	Row_quality              *string    `json:"row_quality"`
}
