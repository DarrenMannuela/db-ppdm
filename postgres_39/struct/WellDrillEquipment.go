package dto

import (
	"time"
)

type Well_drill_equipment struct {
	Uwi                 string     `json:"uwi"`
	Equipment_id        string     `json:"equipment_id"`
	Equipment_obs_no    int        `json:"equipment_obs_no"`
	Active_ind          *string    `json:"active_ind"`
	Effective_date      *time.Time `json:"effective_date"`
	Expiry_date         *time.Time `json:"expiry_date"`
	Offsite_date        *time.Time `json:"offsite_date"`
	Offsite_time        *time.Time `json:"offsite_time"`
	Onsite_date         *time.Time `json:"onsite_date"`
	Onsite_time         *time.Time `json:"onsite_time"`
	Operated_by_ba_id   *string    `json:"operated_by_ba_id"`
	Parent_equipment_id *string    `json:"parent_equipment_id"`
	Period_on_well      *float64   `json:"period_on_well"`
	Period_on_well_ouom *string    `json:"period_on_well_ouom"`
	Ppdm_guid           string     `json:"ppdm_guid"`
	Reference_num       *string    `json:"reference_num"`
	Remark              *string    `json:"remark"`
	Rented_ind          *string    `json:"rented_ind"`
	Source              *string    `json:"source"`
	Timezone            *string    `json:"timezone"`
	Row_changed_by      *string    `json:"row_changed_by"`
	Row_changed_date    *time.Time `json:"row_changed_date"`
	Row_created_by      *string    `json:"row_created_by"`
	Row_created_date    *time.Time `json:"row_created_date"`
	Row_effective_date  *time.Time `json:"row_effective_date"`
	Row_expiry_date     *time.Time `json:"row_expiry_date"`
	Row_quality         *string    `json:"row_quality"`
}
