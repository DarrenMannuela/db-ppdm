package dto

import (
	"time"
)

type Well_ba_service struct {
	Uwi                      string     `json:"uwi"`
	Provided_by              string     `json:"provided_by"`
	Service_seq_no           int        `json:"service_seq_no"`
	Active_ind               *string    `json:"active_ind"`
	Contact_ba_id            *string    `json:"contact_ba_id"`
	Contract_id              *string    `json:"contract_id"`
	Description              *string    `json:"description"`
	Effective_date           *time.Time `json:"effective_date"`
	End_date                 *time.Time `json:"end_date"`
	Expiry_date              *time.Time `json:"expiry_date"`
	Ppdm_guid                string     `json:"ppdm_guid"`
	Provided_for             *string    `json:"provided_for"`
	Provision_id             *string    `json:"provision_id"`
	Rate_schedule_id         *string    `json:"rate_schedule_id"`
	Remark                   *string    `json:"remark"`
	Represented_ba_id        *string    `json:"represented_ba_id"`
	Service_date             *time.Time `json:"service_date"`
	Service_period           *float64   `json:"service_period"`
	Service_period_uom       *string    `json:"service_period_uom"`
	Service_quality          *string    `json:"service_quality"`
	Service_time             *time.Time `json:"service_time"`
	Service_timezone         *string    `json:"service_timezone"`
	Service_time_desc        *string    `json:"service_time_desc"`
	Service_type             *string    `json:"service_type"`
	Source                   *string    `json:"source"`
	Start_date               *time.Time `json:"start_date"`
	Well_activity_obs_no     *int       `json:"well_activity_obs_no"`
	Well_activity_source     *string    `json:"well_activity_source"`
	Well_drill_period_obs_no *int       `json:"well_drill_period_obs_no"`
	Row_changed_by           *string    `json:"row_changed_by"`
	Row_changed_date         *time.Time `json:"row_changed_date"`
	Row_created_by           *string    `json:"row_created_by"`
	Row_created_date         *time.Time `json:"row_created_date"`
	Row_effective_date       *time.Time `json:"row_effective_date"`
	Row_expiry_date          *time.Time `json:"row_expiry_date"`
	Row_quality              *string    `json:"row_quality"`
}
