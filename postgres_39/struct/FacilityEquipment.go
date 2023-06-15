package dto

import (
	"time"
)

type Facility_equipment struct {
	Facility_id        string     `json:"facility_id"`
	Facility_type      string     `json:"facility_type"`
	Equipment_id       string     `json:"equipment_id"`
	Install_obs_no     int        `json:"install_obs_no"`
	Active_ind         *string    `json:"active_ind"`
	Effective_date     *time.Time `json:"effective_date"`
	Expiry_date        *time.Time `json:"expiry_date"`
	Install_date       *time.Time `json:"install_date"`
	Ppdm_guid          string     `json:"ppdm_guid"`
	Remark             *string    `json:"remark"`
	Remove_date        *time.Time `json:"remove_date"`
	Remove_reason      *string    `json:"remove_reason"`
	Source             *string    `json:"source"`
	Use_description    *string    `json:"use_description"`
	Row_changed_by     *string    `json:"row_changed_by"`
	Row_changed_date   *time.Time `json:"row_changed_date"`
	Row_created_by     *string    `json:"row_created_by"`
	Row_created_date   *time.Time `json:"row_created_date"`
	Row_effective_date *time.Time `json:"row_effective_date"`
	Row_expiry_date    *time.Time `json:"row_expiry_date"`
	Row_quality        *string    `json:"row_quality"`
}
