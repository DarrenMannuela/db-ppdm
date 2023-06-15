package dto

import (
	"time"
)

type Well_mud_resistivity struct {
	Uwi                          string     `json:"uwi"`
	Source                       string     `json:"source"`
	Sample_id                    string     `json:"sample_id"`
	Resistivity_obs_no           int        `json:"resistivity_obs_no"`
	Active_ind                   *string    `json:"active_ind"`
	Effective_date               *time.Time `json:"effective_date"`
	Expiry_date                  *time.Time `json:"expiry_date"`
	Mud_resistivity              *float64   `json:"mud_resistivity"`
	Mud_resistivity_ouom         *string    `json:"mud_resistivity_ouom"`
	Ppdm_guid                    string     `json:"ppdm_guid"`
	Remark                       *string    `json:"remark"`
	Resistivity_temperature      *float64   `json:"resistivity_temperature"`
	Resistivity_temperature_ouom *string    `json:"resistivity_temperature_ouom"`
	Sample_type                  *string    `json:"sample_type"`
	Row_changed_by               *string    `json:"row_changed_by"`
	Row_changed_date             *time.Time `json:"row_changed_date"`
	Row_created_by               *string    `json:"row_created_by"`
	Row_created_date             *time.Time `json:"row_created_date"`
	Row_effective_date           *time.Time `json:"row_effective_date"`
	Row_expiry_date              *time.Time `json:"row_expiry_date"`
	Row_quality                  *string    `json:"row_quality"`
}
