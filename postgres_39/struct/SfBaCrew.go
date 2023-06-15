package dto

import (
	"time"
)

type Sf_ba_crew struct {
	Sf_subtype          string     `json:"sf_subtype"`
	Support_facility_id string     `json:"support_facility_id"`
	Crew_company_ba_id  string     `json:"crew_company_ba_id"`
	Crew_id             string     `json:"crew_id"`
	Sf_ba_crew_obs_no   int        `json:"sf_ba_crew_obs_no"`
	Active_ind          *string    `json:"active_ind"`
	Effective_date      *time.Time `json:"effective_date"`
	Expiry_date         *time.Time `json:"expiry_date"`
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
