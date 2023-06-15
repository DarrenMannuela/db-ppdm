package dto

import (
	"time"
)

type Hse_incident_equip struct {
	Incident_id        string     `json:"incident_id"`
	Equip_obs_no       int        `json:"equip_obs_no"`
	Equip_role_obs_no  int        `json:"equip_role_obs_no"`
	Active_ind         *string    `json:"active_ind"`
	Crew_role          *string    `json:"crew_role"`
	Effective_date     *time.Time `json:"effective_date"`
	Equipment_group    *string    `json:"equipment_group"`
	Equipment_id       *string    `json:"equipment_id"`
	Equipment_type     *string    `json:"equipment_type"`
	Expiry_date        *time.Time `json:"expiry_date"`
	Period_obs_no      *int       `json:"period_obs_no"`
	Ppdm_guid          string     `json:"ppdm_guid"`
	Remark             *string    `json:"remark"`
	Source             *string    `json:"source"`
	Uwi                *string    `json:"uwi"`
	Row_changed_by     *string    `json:"row_changed_by"`
	Row_changed_date   *time.Time `json:"row_changed_date"`
	Row_created_by     *string    `json:"row_created_by"`
	Row_created_date   *time.Time `json:"row_created_date"`
	Row_effective_date *time.Time `json:"row_effective_date"`
	Row_expiry_date    *time.Time `json:"row_expiry_date"`
	Row_quality        *string    `json:"row_quality"`
}
