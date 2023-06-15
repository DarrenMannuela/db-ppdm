package dto

import (
	"time"
)

type Well_horiz_drill struct {
	Uwi                      string     `json:"uwi"`
	Source                   string     `json:"source"`
	Active_ind               *string    `json:"active_ind"`
	Buildup_radius_type      *string    `json:"buildup_radius_type"`
	Buildup_rate_degree      *float64   `json:"buildup_rate_degree"`
	Buildup_rate_length      *float64   `json:"buildup_rate_length"`
	Buildup_rate_length_ouom *string    `json:"buildup_rate_length_ouom"`
	Contractor               *string    `json:"contractor"`
	Effective_date           *time.Time `json:"effective_date"`
	Expiry_date              *time.Time `json:"expiry_date"`
	Horiz_displacement       *float64   `json:"horiz_displacement"`
	Horiz_displacement_ouom  *string    `json:"horiz_displacement_ouom"`
	Horiz_drilling_reason    *string    `json:"horiz_drilling_reason"`
	Horiz_drilling_type      *string    `json:"horiz_drilling_type"`
	Horiz_strat_unit_id      *string    `json:"horiz_strat_unit_id"`
	Lateral_hole_id          *string    `json:"lateral_hole_id"`
	Lateral_hole_length      *float64   `json:"lateral_hole_length"`
	Lateral_hole_length_ouom *string    `json:"lateral_hole_length_ouom"`
	Max_deviation_angle      *float64   `json:"max_deviation_angle"`
	Pay_length               *float64   `json:"pay_length"`
	Pay_length_ouom          *string    `json:"pay_length_ouom"`
	Ppdm_guid                string     `json:"ppdm_guid"`
	Rat_hole_depth           *float64   `json:"rat_hole_depth"`
	Rat_hole_depth_ouom      *string    `json:"rat_hole_depth_ouom"`
	Rat_hole_depth_type      *string    `json:"rat_hole_depth_type"`
	Remark                   *string    `json:"remark"`
	Reservoir                *string    `json:"reservoir"`
	Strat_name_set_id        *string    `json:"strat_name_set_id"`
	Wb_length_in_form        *float64   `json:"wb_length_in_form"`
	Wb_length_in_form_ouom   *string    `json:"wb_length_in_form_ouom"`
	Row_changed_by           *string    `json:"row_changed_by"`
	Row_changed_date         *time.Time `json:"row_changed_date"`
	Row_created_by           *string    `json:"row_created_by"`
	Row_created_date         *time.Time `json:"row_created_date"`
	Row_effective_date       *time.Time `json:"row_effective_date"`
	Row_expiry_date          *time.Time `json:"row_expiry_date"`
	Row_quality              *string    `json:"row_quality"`
}
