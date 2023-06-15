package dto

import (
	"time"
)

type Anl_analysis_report struct {
	Analysis_id        string     `json:"analysis_id"`
	Anl_source         string     `json:"anl_source"`
	Active_ind         *string    `json:"active_ind"`
	Analysis_date      *time.Time `json:"analysis_date"`
	Analysis_purpose   *string    `json:"analysis_purpose"`
	Analysis_quality   *string    `json:"analysis_quality"`
	Base_depth         *float64   `json:"base_depth"`
	Base_depth_ouom    *string    `json:"base_depth_ouom"`
	Base_strat_unit_id *string    `json:"base_strat_unit_id"`
	Effective_date     *time.Time `json:"effective_date"`
	End_date           *time.Time `json:"end_date"`
	Expiry_date        *time.Time `json:"expiry_date"`
	Ppdm_guid          string     `json:"ppdm_guid"`
	Received_date      *time.Time `json:"received_date"`
	Remark             *string    `json:"remark"`
	Reported_date      *time.Time `json:"reported_date"`
	Reported_tvd       *float64   `json:"reported_tvd"`
	Reported_tvd_ouom  *string    `json:"reported_tvd_ouom"`
	Sample_date        *time.Time `json:"sample_date"`
	Sample_location    *string    `json:"sample_location"`
	Start_date         *time.Time `json:"start_date"`
	Strat_name_set_id  *string    `json:"strat_name_set_id"`
	Study_type         *string    `json:"study_type"`
	Top_depth          *float64   `json:"top_depth"`
	Top_depth_ouom     *string    `json:"top_depth_ouom"`
	Top_strat_unit_id  *string    `json:"top_strat_unit_id"`
	Row_changed_by     *string    `json:"row_changed_by"`
	Row_changed_date   *time.Time `json:"row_changed_date"`
	Row_created_by     *string    `json:"row_created_by"`
	Row_created_date   *time.Time `json:"row_created_date"`
	Row_effective_date *time.Time `json:"row_effective_date"`
	Row_expiry_date    *time.Time `json:"row_expiry_date"`
	Row_quality        *string    `json:"row_quality"`
}
