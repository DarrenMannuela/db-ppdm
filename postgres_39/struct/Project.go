package dto

import (
	"time"
)

type Project struct {
	Project_id                string     `json:"project_id"`
	Active_ind                *string    `json:"active_ind"`
	Complete_date             *time.Time `json:"complete_date"`
	Confidential_ind          *string    `json:"confidential_ind"`
	Confidential_release_date *time.Time `json:"confidential_release_date"`
	Description               *string    `json:"description"`
	Effective_date            *time.Time `json:"effective_date"`
	Expiry_date               *time.Time `json:"expiry_date"`
	Field_station_ind         *string    `json:"field_station_ind"`
	Land_right_ind            *string    `json:"land_right_ind"`
	Pden_ind                  *string    `json:"pden_ind"`
	Ppdm_guid                 string     `json:"ppdm_guid"`
	Project_name              *string    `json:"project_name"`
	Project_num               *string    `json:"project_num"`
	Project_type              *string    `json:"project_type"`
	Remark                    *string    `json:"remark"`
	Seis_set_ind              *string    `json:"seis_set_ind"`
	Source                    *string    `json:"source"`
	Start_date                *time.Time `json:"start_date"`
	Strat_column_ind          *string    `json:"strat_column_ind"`
	Strat_interpretation_ind  *string    `json:"strat_interpretation_ind"`
	Well_ind                  *string    `json:"well_ind"`
	Row_changed_by            *string    `json:"row_changed_by"`
	Row_changed_date          *time.Time `json:"row_changed_date"`
	Row_created_by            *string    `json:"row_created_by"`
	Row_created_date          *time.Time `json:"row_created_date"`
	Row_effective_date        *time.Time `json:"row_effective_date"`
	Row_expiry_date           *time.Time `json:"row_expiry_date"`
	Row_quality               *string    `json:"row_quality"`
}
