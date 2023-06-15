package dto

import (
	"time"
)

type Seis_set struct {
	Seis_set_subtype           string     `json:"seis_set_subtype"`
	Seis_set_id                string     `json:"seis_set_id"`
	Acqtn_design_id            *string    `json:"acqtn_design_id"`
	Active_ind                 *string    `json:"active_ind"`
	Area_id                    *string    `json:"area_id"`
	Area_size                  *float64   `json:"area_size"`
	Area_size_ouom             *string    `json:"area_size_ouom"`
	Area_type                  *string    `json:"area_type"`
	Coord_acquisition_id       *string    `json:"coord_acquisition_id"`
	Current_seis_status        *string    `json:"current_seis_status"`
	Effective_date             *time.Time `json:"effective_date"`
	End_date                   *time.Time `json:"end_date"`
	Expiry_date                *time.Time `json:"expiry_date"`
	Finance_id                 *string    `json:"finance_id"`
	Geographic_coord_system_id *string    `json:"geographic_coord_system_id"`
	Local_coord_system_id      *string    `json:"local_coord_system_id"`
	Max_latitude               *float64   `json:"max_latitude"`
	Max_longitude              *float64   `json:"max_longitude"`
	Min_latitude               *float64   `json:"min_latitude"`
	Min_longitude              *float64   `json:"min_longitude"`
	Ppdm_guid                  string     `json:"ppdm_guid"`
	Preferred_name             *string    `json:"preferred_name"`
	Remark                     *string    `json:"remark"`
	Source                     *string    `json:"source"`
	Start_date                 *time.Time `json:"start_date"`
	Xy_coord_system_id         *string    `json:"xy_coord_system_id"`
	Row_changed_by             *string    `json:"row_changed_by"`
	Row_changed_date           *time.Time `json:"row_changed_date"`
	Row_created_by             *string    `json:"row_created_by"`
	Row_created_date           *time.Time `json:"row_created_date"`
	Row_effective_date         *time.Time `json:"row_effective_date"`
	Row_expiry_date            *time.Time `json:"row_expiry_date"`
	Row_quality                *string    `json:"row_quality"`
}
