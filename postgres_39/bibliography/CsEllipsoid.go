package dto

import (
	"time"
)

type Cs_ellipsoid struct {
	Ellipsoid_id             string     `json:"ellipsoid_id"`
	Active_ind               *string    `json:"active_ind"`
	Axis_ouom                *string    `json:"axis_ouom"`
	Eccentricity             *float64   `json:"eccentricity"`
	Effective_date           *time.Time `json:"effective_date"`
	Ellipsoid_name           *string    `json:"ellipsoid_name"`
	Expiry_date              *time.Time `json:"expiry_date"`
	Inverse_flattening       *float64   `json:"inverse_flattening"`
	Ppdm_guid                string     `json:"ppdm_guid"`
	Remark                   *string    `json:"remark"`
	Semi_major_axis          *float64   `json:"semi_major_axis"`
	Semi_major_axis_accuracy *int       `json:"semi_major_axis_accuracy"`
	Semi_minor_axis          *float64   `json:"semi_minor_axis"`
	Semi_minor_axis_accuracy *int       `json:"semi_minor_axis_accuracy"`
	Source                   *string    `json:"source"`
	Source_document_id       *string    `json:"source_document_id"`
	Row_changed_by           *string    `json:"row_changed_by"`
	Row_changed_date         *time.Time `json:"row_changed_date"`
	Row_created_by           *string    `json:"row_created_by"`
	Row_created_date         *time.Time `json:"row_created_date"`
	Row_effective_date       *time.Time `json:"row_effective_date"`
	Row_expiry_date          *time.Time `json:"row_expiry_date"`
	Row_quality              *string    `json:"row_quality"`
}
