package dto

import (
	"time"
)

type Ba_crew_member struct {
	Crew_company_ba_id string     `json:"crew_company_ba_id"`
	Crew_id            string     `json:"crew_id"`
	Member_obs_no      int        `json:"member_obs_no"`
	Crew_member_ba_id  string     `json:"crew_member_ba_id"`
	Active_ind         *string    `json:"active_ind"`
	Crew_position      *string    `json:"crew_position"`
	Effective_date     *time.Time `json:"effective_date"`
	Expiry_date        *time.Time `json:"expiry_date"`
	Ppdm_guid          string     `json:"ppdm_guid"`
	Remark             *string    `json:"remark"`
	Source             *string    `json:"source"`
	Row_changed_by     *string    `json:"row_changed_by"`
	Row_changed_date   *time.Time `json:"row_changed_date"`
	Row_created_by     *string    `json:"row_created_by"`
	Row_created_date   *time.Time `json:"row_created_date"`
	Row_effective_date *time.Time `json:"row_effective_date"`
	Row_expiry_date    *time.Time `json:"row_expiry_date"`
	Row_quality        *string    `json:"row_quality"`
}
