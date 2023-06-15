package dto

import (
	"time"
)

type Seis_well struct {
	Seis_set_subtype        string     `json:"seis_set_subtype"`
	Seis_set_id             string     `json:"seis_set_id"`
	Active_ind              *string    `json:"active_ind"`
	Checkshot_survey_type   *string    `json:"checkshot_survey_type"`
	Checkshot_velocity      *float64   `json:"checkshot_velocity"`
	Checkshot_velocity_ouom *string    `json:"checkshot_velocity_ouom"`
	Depth_datum             *string    `json:"depth_datum"`
	Depth_datum_elev        *float64   `json:"depth_datum_elev"`
	Depth_datum_elev_ouom   *string    `json:"depth_datum_elev_ouom"`
	Dir_survey_id           *string    `json:"dir_survey_id"`
	Dir_survey_source       *string    `json:"dir_survey_source"`
	Dir_survey_uwi          *string    `json:"dir_survey_uwi"`
	Effective_date          *time.Time `json:"effective_date"`
	Expiry_date             *time.Time `json:"expiry_date"`
	Ppdm_guid               string     `json:"ppdm_guid"`
	Receiver_uwi            *string    `json:"receiver_uwi"`
	Remark                  *string    `json:"remark"`
	Seismic_path            *string    `json:"seismic_path"`
	Source                  *string    `json:"source"`
	Source_uwi              *string    `json:"source_uwi"`
	Survey_id               *string    `json:"survey_id"`
	Survey_run_num          *string    `json:"survey_run_num"`
	Vsp_type                *string    `json:"vsp_type"`
	Row_changed_by          *string    `json:"row_changed_by"`
	Row_changed_date        *time.Time `json:"row_changed_date"`
	Row_created_by          *string    `json:"row_created_by"`
	Row_created_date        *time.Time `json:"row_created_date"`
	Row_effective_date      *time.Time `json:"row_effective_date"`
	Row_expiry_date         *time.Time `json:"row_expiry_date"`
	Row_quality             *string    `json:"row_quality"`
}
