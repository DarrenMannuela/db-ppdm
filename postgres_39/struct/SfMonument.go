package dto

import (
	"time"
)

type Sf_monument struct {
	Sf_subtype            string     `json:"sf_subtype"`
	Monument_id           string     `json:"monument_id"`
	Active_ind            *string    `json:"active_ind"`
	Coord_acquisition_id  *string    `json:"coord_acquisition_id"`
	Effective_date        *time.Time `json:"effective_date"`
	Expiry_date           *time.Time `json:"expiry_date"`
	Horiz_coord_system    *string    `json:"horiz_coord_system"`
	Local_coord_system_id *string    `json:"local_coord_system_id"`
	Location_type         *string    `json:"location_type"`
	Monument_elev         *float64   `json:"monument_elev"`
	Monument_elev_ouom    *string    `json:"monument_elev_ouom"`
	Monument_latitude     *float64   `json:"monument_latitude"`
	Monument_longitude    *float64   `json:"monument_longitude"`
	Monument_name         *string    `json:"monument_name"`
	Ppdm_guid             string     `json:"ppdm_guid"`
	Remark                *string    `json:"remark"`
	Source                *string    `json:"source"`
	Source_document_id    *string    `json:"source_document_id"`
	Vert_coord_system     *string    `json:"vert_coord_system"`
	Row_changed_by        *string    `json:"row_changed_by"`
	Row_changed_date      *time.Time `json:"row_changed_date"`
	Row_created_by        *string    `json:"row_created_by"`
	Row_created_date      *time.Time `json:"row_created_date"`
	Row_effective_date    *time.Time `json:"row_effective_date"`
	Row_expiry_date       *time.Time `json:"row_expiry_date"`
	Row_quality           *string    `json:"row_quality"`
}
