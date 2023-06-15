package dto

import (
	"time"
)

type Well_activity_cause struct {
	Uwi                string     `json:"uwi"`
	Source             string     `json:"source"`
	Activity_obs_no    int        `json:"activity_obs_no"`
	Cause_type         string     `json:"cause_type"`
	Active_ind         *string    `json:"active_ind"`
	Description        *string    `json:"description"`
	Effective_date     *time.Time `json:"effective_date"`
	Equipment_id       *string    `json:"equipment_id"`
	Expiry_date        *time.Time `json:"expiry_date"`
	Period_obs_no      *int       `json:"period_obs_no"`
	Ppdm_guid          string     `json:"ppdm_guid"`
	Remark             *string    `json:"remark"`
	Row_changed_by     *string    `json:"row_changed_by"`
	Row_changed_date   *time.Time `json:"row_changed_date"`
	Row_created_by     *string    `json:"row_created_by"`
	Row_created_date   *time.Time `json:"row_created_date"`
	Row_effective_date *time.Time `json:"row_effective_date"`
	Row_expiry_date    *time.Time `json:"row_expiry_date"`
	Row_quality        *string    `json:"row_quality"`
}
