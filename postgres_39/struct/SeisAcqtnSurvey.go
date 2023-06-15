package dto

import (
	"time"
)

type Seis_acqtn_survey struct {
	Seis_set_subtype     string     `json:"seis_set_subtype"`
	Seis_acqtn_survey_id string     `json:"seis_acqtn_survey_id"`
	Acqtn_survey_name    *string    `json:"acqtn_survey_name"`
	Active_ind           *string    `json:"active_ind"`
	Area_id              *string    `json:"area_id"`
	Area_type            *string    `json:"area_type"`
	Completed_date       *time.Time `json:"completed_date"`
	Completed_date_desc  *string    `json:"completed_date_desc"`
	Effective_date       *time.Time `json:"effective_date"`
	Expiry_date          *time.Time `json:"expiry_date"`
	Ppdm_guid            string     `json:"ppdm_guid"`
	Remark               *string    `json:"remark"`
	Seis_dimension       *string    `json:"seis_dimension"`
	Shot_for             *string    `json:"shot_for"`
	Source               *string    `json:"source"`
	Start_date           *time.Time `json:"start_date"`
	Start_date_desc      *string    `json:"start_date_desc"`
	Row_changed_by       *string    `json:"row_changed_by"`
	Row_changed_date     *time.Time `json:"row_changed_date"`
	Row_created_by       *string    `json:"row_created_by"`
	Row_created_date     *time.Time `json:"row_created_date"`
	Row_effective_date   *time.Time `json:"row_effective_date"`
	Row_expiry_date      *time.Time `json:"row_expiry_date"`
	Row_quality          *string    `json:"row_quality"`
}
