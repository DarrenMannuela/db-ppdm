package dto

import (
	"time"
)

type Sp_line struct {
	Line_id                string     `json:"line_id"`
	Acquisition_id         *string    `json:"acquisition_id"`
	Active_ind             *string    `json:"active_ind"`
	Coord_system_id        *string    `json:"coord_system_id"`
	Datum_elev             *float64   `json:"datum_elev"`
	Datum_elev_ouom        *string    `json:"datum_elev_ouom"`
	Effective_date         *time.Time `json:"effective_date"`
	Expiry_date            *time.Time `json:"expiry_date"`
	Line_set_id            *string    `json:"line_set_id"`
	Local_coord_system_id  *string    `json:"local_coord_system_id"`
	Location_type          *string    `json:"location_type"`
	Max_plot_scale         *string    `json:"max_plot_scale"`
	Min_plot_scale         *string    `json:"min_plot_scale"`
	Ppdm_guid              string     `json:"ppdm_guid"`
	Preferred_ind          *string    `json:"preferred_ind"`
	Reference_datum        *string    `json:"reference_datum"`
	Remark                 *string    `json:"remark"`
	Source                 *string    `json:"source"`
	Spatial_description_id *string    `json:"spatial_description_id"`
	Spatial_obs_no         *int       `json:"spatial_obs_no"`
	Row_changed_by         *string    `json:"row_changed_by"`
	Row_changed_date       *time.Time `json:"row_changed_date"`
	Row_created_by         *string    `json:"row_created_by"`
	Row_created_date       *time.Time `json:"row_created_date"`
	Row_effective_date     *time.Time `json:"row_effective_date"`
	Row_expiry_date        *time.Time `json:"row_expiry_date"`
	Row_quality            *string    `json:"row_quality"`
}
