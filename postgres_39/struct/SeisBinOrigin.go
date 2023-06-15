package dto

import (
	"time"
)

type Seis_bin_origin struct {
	Seis_set_subtype       string     `json:"seis_set_subtype"`
	Seis_set_id            string     `json:"seis_set_id"`
	Bin_grid_id            string     `json:"bin_grid_id"`
	Bin_grid_source        string     `json:"bin_grid_source"`
	Bin_origin_id          string     `json:"bin_origin_id"`
	Active_ind             *string    `json:"active_ind"`
	Coord_acquisition_id   *string    `json:"coord_acquisition_id"`
	Coord_system_id        *string    `json:"coord_system_id"`
	Corner1_lat            *float64   `json:"corner_1_lat"`
	Corner1_long           *float64   `json:"corner_1_long"`
	Corner2_lat            *float64   `json:"corner_2_lat"`
	Corner2_long           *float64   `json:"corner_2_long"`
	Corner3_lat            *float64   `json:"corner_3_lat"`
	Corner3_long           *float64   `json:"corner_3_long"`
	Effective_date         *time.Time `json:"effective_date"`
	Exclusion_ind          *string    `json:"exclusion_ind"`
	Expiry_date            *time.Time `json:"expiry_date"`
	Inclusion_ind          *string    `json:"inclusion_ind"`
	Input_bin_grid_id      *string    `json:"input_bin_grid_id"`
	Input_bin_source       *string    `json:"input_bin_source"`
	Input_seis_set_id      *string    `json:"input_seis_set_id"`
	Input_seis_set_subtype *string    `json:"input_seis_set_subtype"`
	Local_coord_system_id  *string    `json:"local_coord_system_id"`
	Ppdm_guid              string     `json:"ppdm_guid"`
	Remark                 *string    `json:"remark"`
	Source                 *string    `json:"source"`
	Row_changed_by         *string    `json:"row_changed_by"`
	Row_changed_date       *time.Time `json:"row_changed_date"`
	Row_created_by         *string    `json:"row_created_by"`
	Row_created_date       *time.Time `json:"row_created_date"`
	Row_effective_date     *time.Time `json:"row_effective_date"`
	Row_expiry_date        *time.Time `json:"row_expiry_date"`
	Row_quality            *string    `json:"row_quality"`
}
