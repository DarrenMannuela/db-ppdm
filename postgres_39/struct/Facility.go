package dto

import (
	"time"
)

type Facility struct {
	Facility_id            string     `json:"facility_id"`
	Facility_type          string     `json:"facility_type"`
	Abandoned_date         *time.Time `json:"abandoned_date"`
	Active_date            *time.Time `json:"active_date"`
	Active_ind             *string    `json:"active_ind"`
	Catalogue_equip_id     *string    `json:"catalogue_equip_id"`
	Constructed_date       *time.Time `json:"constructed_date"`
	Coord_acquisition_id   *string    `json:"coord_acquisition_id"`
	Coord_system_id        *string    `json:"coord_system_id"`
	Current_operator       *string    `json:"current_operator"`
	Depth                  *float64   `json:"depth"`
	Depth_ouom             *string    `json:"depth_ouom"`
	Description            *string    `json:"description"`
	Effective_date         *time.Time `json:"effective_date"`
	Elevation              *float64   `json:"elevation"`
	Elevation_ouom         *string    `json:"elevation_ouom"`
	Expiry_date            *time.Time `json:"expiry_date"`
	Facility_diameter      *float64   `json:"facility_diameter"`
	Facility_diameter_ouom *string    `json:"facility_diameter_ouom"`
	Facility_function      *string    `json:"facility_function"`
	Facility_long_name     *string    `json:"facility_long_name"`
	Facility_no            *int       `json:"facility_no"`
	Facility_short_name    *string    `json:"facility_short_name"`
	H2s_ind                *string    `json:"h_2_s_ind"`
	Inactive_date          *time.Time `json:"inactive_date"`
	Last_injection_date    *time.Time `json:"last_injection_date"`
	Last_production_date   *time.Time `json:"last_production_date"`
	Last_reported_date     *time.Time `json:"last_reported_date"`
	Latitude               *float64   `json:"latitude"`
	Local_coord_system_id  *string    `json:"local_coord_system_id"`
	Longitude              *float64   `json:"longitude"`
	Manufactured_by        *string    `json:"manufactured_by"`
	On_injection_date      *time.Time `json:"on_injection_date"`
	On_production_date     *time.Time `json:"on_production_date"`
	Pipeline_material      *string    `json:"pipeline_material"`
	Pipeline_type          *string    `json:"pipeline_type"`
	Pipe_cover_type        *string    `json:"pipe_cover_type"`
	Plot_name              *string    `json:"plot_name"`
	Plot_symbol            *string    `json:"plot_symbol"`
	Ppdm_guid              string     `json:"ppdm_guid"`
	Primary_field_id       *string    `json:"primary_field_id"`
	Remark                 *string    `json:"remark"`
	Source                 *string    `json:"source"`
	Unit_operated_ind      *string    `json:"unit_operated_ind"`
	Row_changed_by         *string    `json:"row_changed_by"`
	Row_changed_date       *time.Time `json:"row_changed_date"`
	Row_created_by         *string    `json:"row_created_by"`
	Row_created_date       *time.Time `json:"row_created_date"`
	Row_effective_date     *time.Time `json:"row_effective_date"`
	Row_expiry_date        *time.Time `json:"row_expiry_date"`
	Row_quality            *string    `json:"row_quality"`
}
