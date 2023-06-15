package dto

import (
	"time"
)

type Well_drill_period_vessel struct {
	Uwi                string     `json:"uwi"`
	Period_obs_no      int        `json:"period_obs_no"`
	Sf_subtype         string     `json:"sf_subtype"`
	Vessel_id          string     `json:"vessel_id"`
	Active_ind         *string    `json:"active_ind"`
	Effective_date     *time.Time `json:"effective_date"`
	Expiry_date        *time.Time `json:"expiry_date"`
	Heading            *float64   `json:"heading"`
	Heading_north_type *string    `json:"heading_north_type"`
	Passengers_off     *int       `json:"passengers_off"`
	Passengers_on      *int       `json:"passengers_on"`
	Ppdm_guid          string     `json:"ppdm_guid"`
	Remark             *string    `json:"remark"`
	Riser_angle        *float64   `json:"riser_angle"`
	Riser_tension      *float64   `json:"riser_tension"`
	Riser_tension_ouom *string    `json:"riser_tension_ouom"`
	Riser_tension_uom  *string    `json:"riser_tension_uom"`
	Source             *string    `json:"source"`
	Vessel_role        *string    `json:"vessel_role"`
	Row_changed_by     *string    `json:"row_changed_by"`
	Row_changed_date   *time.Time `json:"row_changed_date"`
	Row_created_by     *string    `json:"row_created_by"`
	Row_created_date   *time.Time `json:"row_created_date"`
	Row_effective_date *time.Time `json:"row_effective_date"`
	Row_expiry_date    *time.Time `json:"row_expiry_date"`
	Row_quality        *string    `json:"row_quality"`
}
