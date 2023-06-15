package dto

import (
	"time"
)

type Well_sieve_screen struct {
	Uwi                 string     `json:"uwi"`
	Anl_source          string     `json:"anl_source"`
	Analysis_obs_no     int        `json:"analysis_obs_no"`
	Screen_obs_no       int        `json:"screen_obs_no"`
	Active_ind          *string    `json:"active_ind"`
	Cat_equip_id        *string    `json:"cat_equip_id"`
	Effective_date      *time.Time `json:"effective_date"`
	Equipment_id        *string    `json:"equipment_id"`
	Expiry_date         *time.Time `json:"expiry_date"`
	Particle_held_count *int       `json:"particle_held_count"`
	Ppdm_guid           string     `json:"ppdm_guid"`
	Remark              *string    `json:"remark"`
	Screen_mesh_size    *float64   `json:"screen_mesh_size"`
	Row_changed_by      *string    `json:"row_changed_by"`
	Row_changed_date    *time.Time `json:"row_changed_date"`
	Row_created_by      *string    `json:"row_created_by"`
	Row_created_date    *time.Time `json:"row_created_date"`
	Row_effective_date  *time.Time `json:"row_effective_date"`
	Row_expiry_date     *time.Time `json:"row_expiry_date"`
	Row_quality         *string    `json:"row_quality"`
}
