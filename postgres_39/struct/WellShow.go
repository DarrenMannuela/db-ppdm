package dto

import (
	"time"
)

type Well_show struct {
	Uwi                string     `json:"uwi"`
	Source             string     `json:"source"`
	Show_type          string     `json:"show_type"`
	Show_obs_no        int        `json:"show_obs_no"`
	Active_ind         *string    `json:"active_ind"`
	Base_depth         *float64   `json:"base_depth"`
	Base_depth_ouom    *string    `json:"base_depth_ouom"`
	Base_strat_unit_id *string    `json:"base_strat_unit_id"`
	Effective_date     *time.Time `json:"effective_date"`
	Expiry_date        *time.Time `json:"expiry_date"`
	Lithology_desc     *string    `json:"lithology_desc"`
	Ppdm_guid          string     `json:"ppdm_guid"`
	Remark             *string    `json:"remark"`
	Reservoir_ind      *string    `json:"reservoir_ind"`
	Sample_type        *string    `json:"sample_type"`
	Show_symbol        *string    `json:"show_symbol"`
	Source_document_id *string    `json:"source_document_id"`
	Strat_name_set_id  *string    `json:"strat_name_set_id"`
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
