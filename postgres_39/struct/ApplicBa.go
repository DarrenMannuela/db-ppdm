package dto

import (
	"time"
)

type Applic_ba struct {
	Application_id        string     `json:"application_id"`
	Application_ba_id     string     `json:"application_ba_id"`
	Application_ba_obs_no int        `json:"application_ba_obs_no"`
	Active_ind            *string    `json:"active_ind"`
	Application_ba_role   *string    `json:"application_ba_role"`
	Contact_ba_id         *string    `json:"contact_ba_id"`
	Description           *string    `json:"description"`
	Effective_date        *time.Time `json:"effective_date"`
	Expiry_date           *time.Time `json:"expiry_date"`
	Ppdm_guid             string     `json:"ppdm_guid"`
	Remark                *string    `json:"remark"`
	Source                *string    `json:"source"`
	Row_changed_by        *string    `json:"row_changed_by"`
	Row_changed_date      *time.Time `json:"row_changed_date"`
	Row_created_by        *string    `json:"row_created_by"`
	Row_created_date      *time.Time `json:"row_created_date"`
	Row_effective_date    *time.Time `json:"row_effective_date"`
	Row_expiry_date       *time.Time `json:"row_expiry_date"`
	Row_quality           *string    `json:"row_quality"`
}
