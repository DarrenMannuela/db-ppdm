package dto

import (
	"time"
)

type Anl_fluor struct {
	Analysis_id              string     `json:"analysis_id"`
	Anl_source               string     `json:"anl_source"`
	Fluor_obs_no             int        `json:"fluor_obs_no"`
	Active_ind               *string    `json:"active_ind"`
	Color_remark             *string    `json:"color_remark"`
	Effective_date           *time.Time `json:"effective_date"`
	Expiry_date              *time.Time `json:"expiry_date"`
	Fluor_percent            *float64   `json:"fluor_percent"`
	Fluor_remark             *string    `json:"fluor_remark"`
	From_color               *string    `json:"from_color"`
	From_color_frequency     *float64   `json:"from_color_frequency"`
	From_color_frequency_uom *string    `json:"from_color_frequency_uom"`
	From_color_intensity     *float64   `json:"from_color_intensity"`
	From_color_intensity_uom *string    `json:"from_color_intensity_uom"`
	From_mobility_type       *string    `json:"from_mobility_type"`
	Hydrocarbon_show_ind     *string    `json:"hydrocarbon_show_ind"`
	Ppdm_guid                string     `json:"ppdm_guid"`
	Preferred_ind            *string    `json:"preferred_ind"`
	Problem_ind              *string    `json:"problem_ind"`
	Remark                   *string    `json:"remark"`
	Source                   *string    `json:"source"`
	Step_seq_no              *int       `json:"step_seq_no"`
	To_color                 *string    `json:"to_color"`
	To_color_frequency       *float64   `json:"to_color_frequency"`
	To_color_frequency_uom   *string    `json:"to_color_frequency_uom"`
	To_color_intensity       *float64   `json:"to_color_intensity"`
	To_color_intensity_uom   *string    `json:"to_color_intensity_uom"`
	To_mobility_type         *string    `json:"to_mobility_type"`
	Row_changed_by           *string    `json:"row_changed_by"`
	Row_changed_date         *time.Time `json:"row_changed_date"`
	Row_created_by           *string    `json:"row_created_by"`
	Row_created_date         *time.Time `json:"row_created_date"`
	Row_effective_date       *time.Time `json:"row_effective_date"`
	Row_expiry_date          *time.Time `json:"row_expiry_date"`
	Row_quality              *string    `json:"row_quality"`
}
