package dto

import (
	"time"
)

type Seis_proc_component struct {
	Seis_set_subtype          string     `json:"seis_set_subtype"`
	Seis_proc_set_id          string     `json:"seis_proc_set_id"`
	Component_id              string     `json:"component_id"`
	Active_ind                *string    `json:"active_ind"`
	Bin_grid_id               *string    `json:"bin_grid_id"`
	Bin_grid_source           *string    `json:"bin_grid_source"`
	Comp_seis_set_id          *string    `json:"comp_seis_set_id"`
	Comp_seis_set_subtype     *string    `json:"comp_seis_set_subtype"`
	Effective_date            *time.Time `json:"effective_date"`
	Expiry_date               *time.Time `json:"expiry_date"`
	Information_item_id       *string    `json:"information_item_id"`
	Info_item_subtype         *string    `json:"info_item_subtype"`
	Input_ind                 *string    `json:"input_ind"`
	Output_ind                *string    `json:"output_ind"`
	Ppdm_guid                 string     `json:"ppdm_guid"`
	Processing_component_type *string    `json:"processing_component_type"`
	Remark                    *string    `json:"remark"`
	Seis_proc_plan_id         *string    `json:"seis_proc_plan_id"`
	Source                    *string    `json:"source"`
	Uwi                       *string    `json:"uwi"`
	Velocity_volume_id        *string    `json:"velocity_volume_id"`
	Well_log_curve_id         *string    `json:"well_log_curve_id"`
	Row_changed_by            *string    `json:"row_changed_by"`
	Row_changed_date          *time.Time `json:"row_changed_date"`
	Row_created_by            *string    `json:"row_created_by"`
	Row_created_date          *time.Time `json:"row_created_date"`
	Row_effective_date        *time.Time `json:"row_effective_date"`
	Row_expiry_date           *time.Time `json:"row_expiry_date"`
	Row_quality               *string    `json:"row_quality"`
}
