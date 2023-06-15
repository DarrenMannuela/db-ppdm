package dto

import (
	"time"
)

type Hse_incident_interaction struct {
	Incident_id        string     `json:"incident_id"`
	Interaction_obs_no int        `json:"interaction_obs_no"`
	Active_ind         *string    `json:"active_ind"`
	Cause_obs_no       *int       `json:"cause_obs_no"`
	Description        *string    `json:"description"`
	Detail_obs_no      *int       `json:"detail_obs_no"`
	Effective_date     *time.Time `json:"effective_date"`
	Equip_obs_no       *int       `json:"equip_obs_no"`
	Equip_role_obs_no  *int       `json:"equip_role_obs_no"`
	Expiry_date        *time.Time `json:"expiry_date"`
	Incident_substance *string    `json:"incident_substance"`
	Interaction_type   *string    `json:"interaction_type"`
	Party_obs_no       *int       `json:"party_obs_no"`
	Party_role_obs_no  *int       `json:"party_role_obs_no"`
	Ppdm_guid          string     `json:"ppdm_guid"`
	Remark             *string    `json:"remark"`
	Response_obs_no    *int       `json:"response_obs_no"`
	Source             *string    `json:"source"`
	Substance_seq_no   *int       `json:"substance_seq_no"`
	Row_changed_by     *string    `json:"row_changed_by"`
	Row_changed_date   *time.Time `json:"row_changed_date"`
	Row_created_by     *string    `json:"row_created_by"`
	Row_created_date   *time.Time `json:"row_created_date"`
	Row_effective_date *time.Time `json:"row_effective_date"`
	Row_expiry_date    *time.Time `json:"row_expiry_date"`
	Row_quality        *string    `json:"row_quality"`
}
