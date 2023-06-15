package dto

import (
	"time"
)

type Well_drill_assembly_comp struct {
	Uwi                       string     `json:"uwi"`
	Assembly_id               string     `json:"assembly_id"`
	Component_id              string     `json:"component_id"`
	Active_ind                *string    `json:"active_ind"`
	Catalogue_equip_id        *string    `json:"catalogue_equip_id"`
	Component_count           *int       `json:"component_count"`
	Component_desc            *string    `json:"component_desc"`
	Component_length          *float64   `json:"component_length"`
	Component_length_ouom     *string    `json:"component_length_ouom"`
	Component_seq_no          *int       `json:"component_seq_no"`
	Component_type            *string    `json:"component_type"`
	Component_weight          *float64   `json:"component_weight"`
	Component_weight_ouom     *string    `json:"component_weight_ouom"`
	Effective_date            *time.Time `json:"effective_date"`
	Expiry_date               *time.Time `json:"expiry_date"`
	Max_inside_diameter       *float64   `json:"max_inside_diameter"`
	Max_inside_diameter_ouom  *string    `json:"max_inside_diameter_ouom"`
	Max_outside_diameter      *float64   `json:"max_outside_diameter"`
	Max_outside_diameter_ouom *string    `json:"max_outside_diameter_ouom"`
	Min_inside_diameter       *float64   `json:"min_inside_diameter"`
	Min_inside_diameter_ouom  *string    `json:"min_inside_diameter_ouom"`
	Min_outside_diameter      *float64   `json:"min_outside_diameter"`
	Min_outside_diameter_ouom *string    `json:"min_outside_diameter_ouom"`
	Ppdm_guid                 string     `json:"ppdm_guid"`
	Remark                    *string    `json:"remark"`
	Serial_number             *string    `json:"serial_number"`
	Source                    *string    `json:"source"`
	Row_changed_by            *string    `json:"row_changed_by"`
	Row_changed_date          *time.Time `json:"row_changed_date"`
	Row_created_by            *string    `json:"row_created_by"`
	Row_created_date          *time.Time `json:"row_created_date"`
	Row_effective_date        *time.Time `json:"row_effective_date"`
	Row_expiry_date           *time.Time `json:"row_expiry_date"`
	Row_quality               *string    `json:"row_quality"`
}
