package dto

import (
	"time"
)

type Hse_incident_ba struct {
	Incident_id          string     `json:"incident_id"`
	Party_obs_no         int        `json:"party_obs_no"`
	Party_role_obs_no    int        `json:"party_role_obs_no"`
	Active_ind           *string    `json:"active_ind"`
	Company_ba_id        *string    `json:"company_ba_id"`
	Detail_obs_no        *int       `json:"detail_obs_no"`
	Effective_date       *time.Time `json:"effective_date"`
	Expiry_date          *time.Time `json:"expiry_date"`
	Involved_ba_role     *string    `json:"involved_ba_role"`
	Involved_ba_status   *string    `json:"involved_ba_status"`
	Involved_party_ba_id *string    `json:"involved_party_ba_id"`
	Period_obs_no        *int       `json:"period_obs_no"`
	Ppdm_guid            string     `json:"ppdm_guid"`
	Remark               *string    `json:"remark"`
	Source               *string    `json:"source"`
	Uwi                  *string    `json:"uwi"`
	Row_changed_by       *string    `json:"row_changed_by"`
	Row_changed_date     *time.Time `json:"row_changed_date"`
	Row_created_by       *string    `json:"row_created_by"`
	Row_created_date     *time.Time `json:"row_created_date"`
	Row_effective_date   *time.Time `json:"row_effective_date"`
	Row_expiry_date      *time.Time `json:"row_expiry_date"`
	Row_quality          *string    `json:"row_quality"`
}
