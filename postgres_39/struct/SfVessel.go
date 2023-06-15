package dto

import (
	"time"
)

type Sf_vessel struct {
	Sf_subtype                  string     `json:"sf_subtype"`
	Vessel_id                   string     `json:"vessel_id"`
	Active_ind                  *string    `json:"active_ind"`
	Area_id                     *string    `json:"area_id"`
	Area_type                   *string    `json:"area_type"`
	Backup_antenna_location     *string    `json:"backup_antenna_location"`
	Backup_antenna_type         *string    `json:"backup_antenna_type"`
	Current_operator            *string    `json:"current_operator"`
	Drill_hole_position         *string    `json:"drill_hole_position"`
	Effective_date              *time.Time `json:"effective_date"`
	Expiry_date                 *time.Time `json:"expiry_date"`
	Positioning_method          *string    `json:"positioning_method"`
	Ppdm_guid                   string     `json:"ppdm_guid"`
	Primary_antenna_location    *string    `json:"primary_antenna_location"`
	Primary_antenna_type        *string    `json:"primary_antenna_type"`
	Remark                      *string    `json:"remark"`
	Secondary_antennal_location *string    `json:"secondary_antennal_location"`
	Secondary_antenna_type      *string    `json:"secondary_antenna_type"`
	Source                      *string    `json:"source"`
	Vessel_beam                 *float64   `json:"vessel_beam"`
	Vessel_displacement         *float64   `json:"vessel_displacement"`
	Vessel_draft                *float64   `json:"vessel_draft"`
	Vessel_length               *float64   `json:"vessel_length"`
	Vessel_name                 *string    `json:"vessel_name"`
	Vessel_size                 *string    `json:"vessel_size"`
	Vessel_type                 *string    `json:"vessel_type"`
	Row_changed_by              *string    `json:"row_changed_by"`
	Row_changed_date            *time.Time `json:"row_changed_date"`
	Row_created_by              *string    `json:"row_created_by"`
	Row_created_date            *time.Time `json:"row_created_date"`
	Row_effective_date          *time.Time `json:"row_effective_date"`
	Row_expiry_date             *time.Time `json:"row_expiry_date"`
	Row_quality                 *string    `json:"row_quality"`
}
