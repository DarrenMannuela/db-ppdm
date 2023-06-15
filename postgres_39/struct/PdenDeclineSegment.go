package dto

import (
	"time"
)

type Pden_decline_segment struct {
	Pden_subtype       string     `json:"pden_subtype"`
	Pden_id            string     `json:"pden_id"`
	Pden_source        string     `json:"pden_source"`
	Case_id            string     `json:"case_id"`
	Segment_id         string     `json:"segment_id"`
	Active_ind         *string    `json:"active_ind"`
	Decline_curve_type *string    `json:"decline_curve_type"`
	Decline_type       *string    `json:"decline_type"`
	Duration           *float64   `json:"duration"`
	Duration_ouom      *string    `json:"duration_ouom"`
	Effective_date     *time.Time `json:"effective_date"`
	End_date           *time.Time `json:"end_date"`
	Expiry_date        *time.Time `json:"expiry_date"`
	Final_decline      *float64   `json:"final_decline"`
	Final_rate         *float64   `json:"final_rate"`
	Initial_decline    *float64   `json:"initial_decline"`
	Initial_rate       *float64   `json:"initial_rate"`
	Minimum_decline    *float64   `json:"minimum_decline"`
	N_factor           *float64   `json:"n_factor"`
	Ppdm_guid          string     `json:"ppdm_guid"`
	Product_type       *string    `json:"product_type"`
	Rate_ouom          *string    `json:"rate_ouom"`
	Ratio_curve_type   *string    `json:"ratio_curve_type"`
	Ratio_final_rate   *float64   `json:"ratio_final_rate"`
	Ratio_fluid_type   *string    `json:"ratio_fluid_type"`
	Ratio_initial_rate *float64   `json:"ratio_initial_rate"`
	Ratio_rate_ouom    *string    `json:"ratio_rate_ouom"`
	Ratio_volume       *float64   `json:"ratio_volume"`
	Ratio_volume_ouom  *string    `json:"ratio_volume_ouom"`
	Remark             *string    `json:"remark"`
	Source             *string    `json:"source"`
	Start_date         *time.Time `json:"start_date"`
	Volume             *float64   `json:"volume"`
	Volume_ouom        *string    `json:"volume_ouom"`
	Row_changed_by     *string    `json:"row_changed_by"`
	Row_changed_date   *time.Time `json:"row_changed_date"`
	Row_created_by     *string    `json:"row_created_by"`
	Row_created_date   *time.Time `json:"row_created_date"`
	Row_effective_date *time.Time `json:"row_effective_date"`
	Row_expiry_date    *time.Time `json:"row_expiry_date"`
	Row_quality        *string    `json:"row_quality"`
}
