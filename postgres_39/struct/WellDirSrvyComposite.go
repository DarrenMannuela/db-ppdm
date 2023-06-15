package dto

import (
	"time"
)

type Well_dir_srvy_composite struct {
	Composite_uwi       string     `json:"composite_uwi"`
	Composite_survey_id string     `json:"composite_survey_id"`
	Composite_source    string     `json:"composite_source"`
	Composite_obs_no    int        `json:"composite_obs_no"`
	Active_ind          *string    `json:"active_ind"`
	Depth_obs_no_from   *int       `json:"depth_obs_no_from"`
	Depth_obs_no_to     *int       `json:"depth_obs_no_to"`
	Effective_date      *time.Time `json:"effective_date"`
	Expiry_date         *time.Time `json:"expiry_date"`
	Input_source        *string    `json:"input_source"`
	Input_survey_id     *string    `json:"input_survey_id"`
	Input_uwi           *string    `json:"input_uwi"`
	Ppdm_guid           string     `json:"ppdm_guid"`
	Remark              *string    `json:"remark"`
	Source              *string    `json:"source"`
	Row_changed_by      *string    `json:"row_changed_by"`
	Row_changed_date    *time.Time `json:"row_changed_date"`
	Row_created_by      *string    `json:"row_created_by"`
	Row_created_date    *time.Time `json:"row_created_date"`
	Row_effective_date  *time.Time `json:"row_effective_date"`
	Row_expiry_date     *time.Time `json:"row_expiry_date"`
	Row_quality         *string    `json:"row_quality"`
}
