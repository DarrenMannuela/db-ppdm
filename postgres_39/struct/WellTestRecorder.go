package dto

import (
	"time"
)

type Well_test_recorder struct {
	Uwi                        string     `json:"uwi"`
	Source                     string     `json:"source"`
	Test_type                  string     `json:"test_type"`
	Run_num                    string     `json:"run_num"`
	Test_num                   string     `json:"test_num"`
	Recorder_id                string     `json:"recorder_id"`
	Active_ind                 *string    `json:"active_ind"`
	Effective_date             *time.Time `json:"effective_date"`
	Expiry_date                *time.Time `json:"expiry_date"`
	Max_capacity_pressure      *float64   `json:"max_capacity_pressure"`
	Max_capacity_pressure_ouom *string    `json:"max_capacity_pressure_ouom"`
	Max_capacity_temp          *float64   `json:"max_capacity_temp"`
	Max_capacity_temp_ouom     *string    `json:"max_capacity_temp_ouom"`
	Performance_quality        *string    `json:"performance_quality"`
	Ppdm_guid                  string     `json:"ppdm_guid"`
	Recorder_depth             *float64   `json:"recorder_depth"`
	Recorder_depth_ouom        *string    `json:"recorder_depth_ouom"`
	Recorder_inside_ind        *string    `json:"recorder_inside_ind"`
	Recorder_max_temp          *float64   `json:"recorder_max_temp"`
	Recorder_max_temp_ouom     *string    `json:"recorder_max_temp_ouom"`
	Recorder_min_temp          *float64   `json:"recorder_min_temp"`
	Recorder_min_temp_ouom     *string    `json:"recorder_min_temp_ouom"`
	Recorder_position          *string    `json:"recorder_position"`
	Recorder_resolution        *float64   `json:"recorder_resolution"`
	Recorder_resolution_ouom   *string    `json:"recorder_resolution_ouom"`
	Recorder_type              *string    `json:"recorder_type"`
	Recorder_used_ind          *string    `json:"recorder_used_ind"`
	Remark                     *string    `json:"remark"`
	Row_changed_by             *string    `json:"row_changed_by"`
	Row_changed_date           *time.Time `json:"row_changed_date"`
	Row_created_by             *string    `json:"row_created_by"`
	Row_created_date           *time.Time `json:"row_created_date"`
	Row_effective_date         *time.Time `json:"row_effective_date"`
	Row_expiry_date            *time.Time `json:"row_expiry_date"`
	Row_quality                *string    `json:"row_quality"`
}
