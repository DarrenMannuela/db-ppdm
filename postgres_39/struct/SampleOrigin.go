package dto

import (
	"time"
)

type Sample_origin struct {
	Sample_id             string     `json:"sample_id"`
	Collection_obs_no     int        `json:"collection_obs_no"`
	Active_ind            *string    `json:"active_ind"`
	Analysis_id           *string    `json:"analysis_id"`
	Analysis_source       *string    `json:"analysis_source"`
	Anl_step_seq_no       *int       `json:"anl_step_seq_no"`
	Area_id               *string    `json:"area_id"`
	Area_type             *string    `json:"area_type"`
	Collected_by_ba_id    *string    `json:"collected_by_ba_id"`
	Collected_for_ba_id   *string    `json:"collected_for_ba_id"`
	Collect_method        *string    `json:"collect_method"`
	Core_type             *string    `json:"core_type"`
	Effective_date        *time.Time `json:"effective_date"`
	Expiry_date           *time.Time `json:"expiry_date"`
	Field_station_id      *string    `json:"field_station_id"`
	Land_right_id         *string    `json:"land_right_id"`
	Land_right_subtype    *string    `json:"land_right_subtype"`
	Meas_section_id       *string    `json:"meas_section_id"`
	Meas_section_source   *string    `json:"meas_section_source"`
	Portion_diameter      *float64   `json:"portion_diameter"`
	Portion_diameter_ouom *string    `json:"portion_diameter_ouom"`
	Portion_length        *float64   `json:"portion_length"`
	Portion_length_ouom   *string    `json:"portion_length_ouom"`
	Portion_volume        *float64   `json:"portion_volume"`
	Portion_volume_ouom   *string    `json:"portion_volume_ouom"`
	Portion_weight        *float64   `json:"portion_weight"`
	Portion_weight_ouom   *string    `json:"portion_weight_ouom"`
	Ppdm_guid             string     `json:"ppdm_guid"`
	Prod_string_id        *string    `json:"prod_string_id"`
	Prod_string_source    *string    `json:"prod_string_source"`
	Project_id            *string    `json:"project_id"`
	Remark                *string    `json:"remark"`
	Sample_part_id        *string    `json:"sample_part_id"`
	Source                *string    `json:"source"`
	Well_core_id          *string    `json:"well_core_id"`
	Well_source           *string    `json:"well_source"`
	Well_survey_id        *string    `json:"well_survey_id"`
	Well_survey_source    *string    `json:"well_survey_source"`
	Well_uwi              *string    `json:"well_uwi"`
	Row_changed_by        *string    `json:"row_changed_by"`
	Row_changed_date      *time.Time `json:"row_changed_date"`
	Row_created_by        *string    `json:"row_created_by"`
	Row_created_date      *time.Time `json:"row_created_date"`
	Row_effective_date    *time.Time `json:"row_effective_date"`
	Row_expiry_date       *time.Time `json:"row_expiry_date"`
	Row_quality           *string    `json:"row_quality"`
}
