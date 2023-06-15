package dto

import (
	"time"
)

type Hse_incident_response struct {
	Incident_id        string     `json:"incident_id"`
	Response_obs_no    int        `json:"response_obs_no"`
	Active_ind         *string    `json:"active_ind"`
	Effective_date     *time.Time `json:"effective_date"`
	Expiry_date        *time.Time `json:"expiry_date"`
	Ppdm_guid          string     `json:"ppdm_guid"`
	Project_id         *string    `json:"project_id"`
	Remark             *string    `json:"remark"`
	Response_result    *string    `json:"response_result"`
	Response_type      *string    `json:"response_type"`
	Source             *string    `json:"source"`
	Row_changed_by     *string    `json:"row_changed_by"`
	Row_changed_date   *time.Time `json:"row_changed_date"`
	Row_created_by     *string    `json:"row_created_by"`
	Row_created_date   *time.Time `json:"row_created_date"`
	Row_effective_date *time.Time `json:"row_effective_date"`
	Row_expiry_date    *time.Time `json:"row_expiry_date"`
	Row_quality        *string    `json:"row_quality"`
}
