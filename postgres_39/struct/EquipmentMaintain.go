package dto

import (
	"time"
)

type Equipment_maintain struct {
	Equipment_id               string     `json:"equipment_id"`
	Equip_maint_id             string     `json:"equip_maint_id"`
	Active_ind                 *string    `json:"active_ind"`
	Actual_end_date            *time.Time `json:"actual_end_date"`
	Actual_start_date          *time.Time `json:"actual_start_date"`
	Catalogue_equip_id         *string    `json:"catalogue_equip_id"`
	Completed_by_ba_id         *string    `json:"completed_by_ba_id"`
	Effective_date             *time.Time `json:"effective_date"`
	End_date                   *time.Time `json:"end_date"`
	Expiry_date                *time.Time `json:"expiry_date"`
	Failure_ind                *string    `json:"failure_ind"`
	Location_ba_address_obs_no *int       `json:"location_ba_address_obs_no"`
	Location_ba_id             *string    `json:"location_ba_id"`
	Location_ba_source         *string    `json:"location_ba_source"`
	Maint_location_type        *string    `json:"maint_location_type"`
	Maint_reason               *string    `json:"maint_reason"`
	Maint_type                 *string    `json:"maint_type"`
	Ppdm_guid                  string     `json:"ppdm_guid"`
	Project_id                 *string    `json:"project_id"`
	Remark                     *string    `json:"remark"`
	Scheduled_date             *time.Time `json:"scheduled_date"`
	Scheduled_ind              *string    `json:"scheduled_ind"`
	Source                     *string    `json:"source"`
	Start_date                 *time.Time `json:"start_date"`
	System_condition           *string    `json:"system_condition"`
	Row_changed_by             *string    `json:"row_changed_by"`
	Row_changed_date           *time.Time `json:"row_changed_date"`
	Row_created_by             *string    `json:"row_created_by"`
	Row_created_date           *time.Time `json:"row_created_date"`
	Row_effective_date         *time.Time `json:"row_effective_date"`
	Row_expiry_date            *time.Time `json:"row_expiry_date"`
	Row_quality                *string    `json:"row_quality"`
}
