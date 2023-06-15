package dto

import (
	"time"
)

type Cs_geodetic_datum struct {
	Geodetic_datum           string     `json:"geodetic_datum"`
	Active_ind               *string    `json:"active_ind"`
	Datum_origin             *string    `json:"datum_origin"`
	Effective_date           *time.Time `json:"effective_date"`
	Ellipsoid_id             *string    `json:"ellipsoid_id"`
	Expiry_date              *time.Time `json:"expiry_date"`
	Geodetic_datum_area_id   *string    `json:"geodetic_datum_area_id"`
	Geodetic_datum_area_type *string    `json:"geodetic_datum_area_type"`
	Geodetic_datum_name      *string    `json:"geodetic_datum_name"`
	Origin_angular_ouom      *string    `json:"origin_angular_ouom"`
	Origin_latitude          *float64   `json:"origin_latitude"`
	Origin_longitude         *float64   `json:"origin_longitude"`
	Origin_name              *string    `json:"origin_name"`
	Ppdm_guid                string     `json:"ppdm_guid"`
	Remark                   *string    `json:"remark"`
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
