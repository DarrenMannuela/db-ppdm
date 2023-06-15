package dto

import (
	"time"
)

type Sf_road struct {
	Sf_subtype              string     `json:"sf_subtype"`
	Road_id                 string     `json:"road_id"`
	Active_ind              *string    `json:"active_ind"`
	Capacity                *float64   `json:"capacity"`
	Capacity_ouom           *string    `json:"capacity_ouom"`
	Communication_freq      *float64   `json:"communication_freq"`
	Communication_freq_desc *string    `json:"communication_freq_desc"`
	Control_type            *string    `json:"control_type"`
	Current_operator_ba_id  *string    `json:"current_operator_ba_id"`
	Direction               *string    `json:"direction"`
	Driving_side            *string    `json:"driving_side"`
	Effective_date          *time.Time `json:"effective_date"`
	Expiry_date             *time.Time `json:"expiry_date"`
	Ppdm_guid               string     `json:"ppdm_guid"`
	Remark                  *string    `json:"remark"`
	Road_length             *float64   `json:"road_length"`
	Road_length_ouom        *string    `json:"road_length_ouom"`
	Road_type               *string    `json:"road_type"`
	Road_width              *float64   `json:"road_width"`
	Road_width_ouom         *string    `json:"road_width_ouom"`
	Source                  *string    `json:"source"`
	Surface_type            *string    `json:"surface_type"`
	Traffic_flow_type       *string    `json:"traffic_flow_type"`
	Row_changed_by          *string    `json:"row_changed_by"`
	Row_changed_date        *time.Time `json:"row_changed_date"`
	Row_created_by          *string    `json:"row_created_by"`
	Row_created_date        *time.Time `json:"row_created_date"`
	Row_effective_date      *time.Time `json:"row_effective_date"`
	Row_expiry_date         *time.Time `json:"row_expiry_date"`
	Row_quality             *string    `json:"row_quality"`
}
