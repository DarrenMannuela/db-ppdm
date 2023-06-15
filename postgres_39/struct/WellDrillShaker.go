package dto

import (
	"time"
)

type Well_drill_shaker struct {
	Uwi                string     `json:"uwi"`
	Period_obs_no      int        `json:"period_obs_no"`
	Shaker_id          string     `json:"shaker_id"`
	Screen_obs_no      int        `json:"screen_obs_no"`
	Active_ind         *string    `json:"active_ind"`
	Catalogue_equip_id *string    `json:"catalogue_equip_id"`
	Effective_date     *time.Time `json:"effective_date"`
	Equipment_id       *string    `json:"equipment_id"`
	Expiry_date        *time.Time `json:"expiry_date"`
	New_screen_ind     *string    `json:"new_screen_ind"`
	Ppdm_guid          string     `json:"ppdm_guid"`
	Remark             *string    `json:"remark"`
	Screen_change_ind  *string    `json:"screen_change_ind"`
	Screen_location    *string    `json:"screen_location"`
	Screen_mesh_desc   *string    `json:"screen_mesh_desc"`
	Source             *string    `json:"source"`
	Row_changed_by     *string    `json:"row_changed_by"`
	Row_changed_date   *time.Time `json:"row_changed_date"`
	Row_created_by     *string    `json:"row_created_by"`
	Row_created_date   *time.Time `json:"row_created_date"`
	Row_effective_date *time.Time `json:"row_effective_date"`
	Row_expiry_date    *time.Time `json:"row_expiry_date"`
	Row_quality        *string    `json:"row_quality"`
}
