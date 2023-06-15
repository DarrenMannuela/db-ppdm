package dto

import (
	"time"
)

type Cs_coordinate_system struct {
	Coord_system_id           string     `json:"coord_system_id"`
	Active_ind                *string    `json:"active_ind"`
	Coordinate_system_type    *string    `json:"coordinate_system_type"`
	Coord_system_abbreviation *string    `json:"coord_system_abbreviation"`
	Coord_system_area         *string    `json:"coord_system_area"`
	Coord_system_long_name    *string    `json:"coord_system_long_name"`
	Coord_system_short_name   *string    `json:"coord_system_short_name"`
	Coord_system_uom          *string    `json:"coord_system_uom"`
	Datum_ouom                *string    `json:"datum_ouom"`
	Effective_date            *time.Time `json:"effective_date"`
	Expiry_date               *time.Time `json:"expiry_date"`
	Geodetic_datum            *string    `json:"geodetic_datum"`
	N_value                   *float64   `json:"n_value"`
	N_value_ouom              *string    `json:"n_value_ouom"`
	Owner_ba_id               *string    `json:"owner_ba_id"`
	Parent_coord_system_id    *string    `json:"parent_coord_system_id"`
	Perspective_height        *float64   `json:"perspective_height"`
	Ppdm_guid                 string     `json:"ppdm_guid"`
	Prime_meridian            *string    `json:"prime_meridian"`
	Principal_meridian        *string    `json:"principal_meridian"`
	Projection_type           *string    `json:"projection_type"`
	Reference_elev            *float64   `json:"reference_elev"`
	Reference_elev_ouom       *string    `json:"reference_elev_ouom"`
	Remark                    *string    `json:"remark"`
	Rotation_ind              *string    `json:"rotation_ind"`
	Source                    *string    `json:"source"`
	Source_document_id        *string    `json:"source_document_id"`
	Vertical_datum_type       *string    `json:"vertical_datum_type"`
	Row_changed_by            *string    `json:"row_changed_by"`
	Row_changed_date          *time.Time `json:"row_changed_date"`
	Row_created_by            *string    `json:"row_created_by"`
	Row_created_date          *time.Time `json:"row_created_date"`
	Row_effective_date        *time.Time `json:"row_effective_date"`
	Row_expiry_date           *time.Time `json:"row_expiry_date"`
	Row_quality               *string    `json:"row_quality"`
}
