package dto

import (
	"time"
)

type Sf_bridge struct {
	Sf_subtype           string     `json:"sf_subtype"`
	Bridge_id            string     `json:"bridge_id"`
	Active_ind           *string    `json:"active_ind"`
	Bridge_capacity      *float64   `json:"bridge_capacity"`
	Bridge_capacity_ouom *string    `json:"bridge_capacity_ouom"`
	Bridge_height        *float64   `json:"bridge_height"`
	Bridge_height_ouom   *string    `json:"bridge_height_ouom"`
	Bridge_length        *float64   `json:"bridge_length"`
	Bridge_length_ouom   *string    `json:"bridge_length_ouom"`
	Bridge_type          *string    `json:"bridge_type"`
	Bridge_width         *float64   `json:"bridge_width"`
	Bridge_width_ouom    *string    `json:"bridge_width_ouom"`
	Effective_date       *time.Time `json:"effective_date"`
	Expiry_date          *time.Time `json:"expiry_date"`
	Ppdm_guid            string     `json:"ppdm_guid"`
	Remark               *string    `json:"remark"`
	Source               *string    `json:"source"`
	Surface_type         *string    `json:"surface_type"`
	Row_changed_by       *string    `json:"row_changed_by"`
	Row_changed_date     *time.Time `json:"row_changed_date"`
	Row_created_by       *string    `json:"row_created_by"`
	Row_created_date     *time.Time `json:"row_created_date"`
	Row_effective_date   *time.Time `json:"row_effective_date"`
	Row_expiry_date      *time.Time `json:"row_expiry_date"`
	Row_quality          *string    `json:"row_quality"`
}
