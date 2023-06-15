package dto

import (
	"time"
)

type Well_test_comput_anal struct {
	Uwi                         string     `json:"uwi"`
	Source                      string     `json:"source"`
	Test_type                   string     `json:"test_type"`
	Run_num                     string     `json:"run_num"`
	Test_num                    string     `json:"test_num"`
	Report_no                   int        `json:"report_no"`
	Active_ind                  *string    `json:"active_ind"`
	Analysis_company            *string    `json:"analysis_company"`
	Computed_permeability       *float64   `json:"computed_permeability"`
	Computed_permeability_ouom  *string    `json:"computed_permeability_ouom"`
	Confidence_limit            *float64   `json:"confidence_limit"`
	Confidence_limit_ouom       *string    `json:"confidence_limit_ouom"`
	Effective_date              *time.Time `json:"effective_date"`
	Est_damage_ratio            *float64   `json:"est_damage_ratio"`
	Expiry_date                 *time.Time `json:"expiry_date"`
	Extrap_pressure_percent     *float64   `json:"extrap_pressure_percent"`
	Final_reservoir_pressure    *float64   `json:"final_reservoir_pressure"`
	Final_reservoir_press_ouom  *string    `json:"final_reservoir_press_ouom"`
	Gauge_depth                 *float64   `json:"gauge_depth"`
	Gauge_depth_ouom            *string    `json:"gauge_depth_ouom"`
	Investigation_radius        *float64   `json:"investigation_radius"`
	Investigation_radius_ouom   *string    `json:"investigation_radius_ouom"`
	Max_reservoir_pressure      *float64   `json:"max_reservoir_pressure"`
	Max_reservoir_pressure_ouom *string    `json:"max_reservoir_pressure_ouom"`
	Potmtrc_surf_height         *float64   `json:"potmtrc_surf_height"`
	Potmtrc_surf_height_ouom    *string    `json:"potmtrc_surf_height_ouom"`
	Ppdm_guid                   string     `json:"ppdm_guid"`
	Production_index_rate       *float64   `json:"production_index_rate"`
	Production_index_rate_ouom  *string    `json:"production_index_rate_ouom"`
	Recorder_id                 *string    `json:"recorder_id"`
	Remark                      *string    `json:"remark"`
	Row_changed_by              *string    `json:"row_changed_by"`
	Row_changed_date            *time.Time `json:"row_changed_date"`
	Row_created_by              *string    `json:"row_created_by"`
	Row_created_date            *time.Time `json:"row_created_date"`
	Row_effective_date          *time.Time `json:"row_effective_date"`
	Row_expiry_date             *time.Time `json:"row_expiry_date"`
	Row_quality                 *string    `json:"row_quality"`
}
