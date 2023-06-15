package dto

import (
	"time"
)

type Seis_interp_comp struct {
	Seis_set_subtype      string     `json:"seis_set_subtype"`
	Interp_set_id         string     `json:"interp_set_id"`
	Component_id          string     `json:"component_id"`
	Active_ind            *string    `json:"active_ind"`
	Area_id               *string    `json:"area_id"`
	Area_type             *string    `json:"area_type"`
	Coord_acquisition_id  *string    `json:"coord_acquisition_id"`
	Coord_system_id       *string    `json:"coord_system_id"`
	Corner1_lat           *float64   `json:"corner_1_lat"`
	Corner1_long          *float64   `json:"corner_1_long"`
	Corner2_lat           *float64   `json:"corner_2_lat"`
	Corner2_long          *float64   `json:"corner_2_long"`
	Corner3_lat           *float64   `json:"corner_3_lat"`
	Corner3_long          *float64   `json:"corner_3_long"`
	Data_sample_size      *float64   `json:"data_sample_size"`
	Effective_date        *time.Time `json:"effective_date"`
	Expiry_date           *time.Time `json:"expiry_date"`
	Information_item_id   *string    `json:"information_item_id"`
	Info_item_subtype     *string    `json:"info_item_subtype"`
	Input_ind             *string    `json:"input_ind"`
	Local_coord_system_id *string    `json:"local_coord_system_id"`
	Origin_type           *string    `json:"origin_type"`
	Output_ind            *string    `json:"output_ind"`
	Ppdm_guid             string     `json:"ppdm_guid"`
	Process_step_id       *string    `json:"process_step_id"`
	Proc_component_id     *string    `json:"proc_component_id"`
	Proc_set_id           *string    `json:"proc_set_id"`
	Proc_set_subtype      *string    `json:"proc_set_subtype"`
	Remark                *string    `json:"remark"`
	Source                *string    `json:"source"`
	Uwi                   *string    `json:"uwi"`
	Velocity_volume_id    *string    `json:"velocity_volume_id"`
	Row_changed_by        *string    `json:"row_changed_by"`
	Row_changed_date      *time.Time `json:"row_changed_date"`
	Row_created_by        *string    `json:"row_created_by"`
	Row_created_date      *time.Time `json:"row_created_date"`
	Row_effective_date    *time.Time `json:"row_effective_date"`
	Row_expiry_date       *time.Time `json:"row_expiry_date"`
	Row_quality           *string    `json:"row_quality"`
}
