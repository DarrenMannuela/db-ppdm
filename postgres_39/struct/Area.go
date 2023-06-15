package dto

import (
	"time"
)

type Area struct {
	Area_id               string     `json:"area_id"`
	Area_type             string     `json:"area_type"`
	Active_ind            *string    `json:"active_ind"`
	Area_max_latitude     *float64   `json:"area_max_latitude"`
	Area_max_longitude    *float64   `json:"area_max_longitude"`
	Area_min_latitude     *float64   `json:"area_min_latitude"`
	Area_min_longitude    *float64   `json:"area_min_longitude"`
	Coord_acquisition_id  *string    `json:"coord_acquisition_id"`
	Coord_system_id       *string    `json:"coord_system_id"`
	Effective_date        *time.Time `json:"effective_date"`
	Expiry_date           *time.Time `json:"expiry_date"`
	Local_coord_system_id *string    `json:"local_coord_system_id"`
	Ppdm_guid             string     `json:"ppdm_guid"`
	Preferred_name        *string    `json:"preferred_name"`
	Remark                *string    `json:"remark"`
	Source                *string    `json:"source"`
	Source_document_id    *string    `json:"source_document_id"`
	Row_changed_by        *string    `json:"row_changed_by"`
	Row_changed_date      *time.Time `json:"row_changed_date"`
	Row_created_by        *string    `json:"row_created_by"`
	Row_created_date      *time.Time `json:"row_created_date"`
	Row_effective_date    *time.Time `json:"row_effective_date"`
	Row_expiry_date       *time.Time `json:"row_expiry_date"`
	Row_quality           *string    `json:"row_quality"`
}
