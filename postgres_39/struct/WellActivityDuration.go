package dto

import (
	"time"
)

type Well_activity_duration struct {
	Uwi                    string     `json:"uwi"`
	Source                 string     `json:"source"`
	Activity_obs_no        int        `json:"activity_obs_no"`
	Duration_obs_no        int        `json:"duration_obs_no"`
	Active_ind             *string    `json:"active_ind"`
	Activity_duration      *float64   `json:"activity_duration"`
	Activity_duration_ouom *string    `json:"activity_duration_ouom"`
	Downtime_type          *string    `json:"downtime_type"`
	Effective_date         *time.Time `json:"effective_date"`
	End_time               *time.Time `json:"end_time"`
	End_timezone           *string    `json:"end_timezone"`
	Event_date             *time.Time `json:"event_date"`
	Expiry_date            *time.Time `json:"expiry_date"`
	Period_obs_no          *int       `json:"period_obs_no"`
	Ppdm_guid              string     `json:"ppdm_guid"`
	Remark                 *string    `json:"remark"`
	Start_time             *time.Time `json:"start_time"`
	Start_timezone         *string    `json:"start_timezone"`
	Row_changed_by         *string    `json:"row_changed_by"`
	Row_changed_date       *time.Time `json:"row_changed_date"`
	Row_created_by         *string    `json:"row_created_by"`
	Row_created_date       *time.Time `json:"row_created_date"`
	Row_effective_date     *time.Time `json:"row_effective_date"`
	Row_expiry_date        *time.Time `json:"row_expiry_date"`
	Row_quality            *string    `json:"row_quality"`
}
