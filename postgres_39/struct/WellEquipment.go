package dto

import (
	"time"
)

type Well_equipment struct {
	Uwi                     string     `json:"uwi"`
	Source                  string     `json:"source"`
	Equipment_id            string     `json:"equipment_id"`
	Equip_obs_no            int        `json:"equip_obs_no"`
	Active_ind              *string    `json:"active_ind"`
	Effective_date          *time.Time `json:"effective_date"`
	Expiry_date             *time.Time `json:"expiry_date"`
	Install_base_depth      *float64   `json:"install_base_depth"`
	Install_base_depth_ouom *string    `json:"install_base_depth_ouom"`
	Install_date            *time.Time `json:"install_date"`
	Install_seq_no          *int       `json:"install_seq_no"`
	Install_top_depth       *float64   `json:"install_top_depth"`
	Install_top_depth_ouom  *string    `json:"install_top_depth_ouom"`
	Parent_equipment_id     *string    `json:"parent_equipment_id"`
	Ppdm_guid               string     `json:"ppdm_guid"`
	Remark                  *string    `json:"remark"`
	Removal_date            *time.Time `json:"removal_date"`
	String_id               *string    `json:"string_id"`
	String_source           *string    `json:"string_source"`
	Row_changed_by          *string    `json:"row_changed_by"`
	Row_changed_date        *time.Time `json:"row_changed_date"`
	Row_created_by          *string    `json:"row_created_by"`
	Row_created_date        *time.Time `json:"row_created_date"`
	Row_effective_date      *time.Time `json:"row_effective_date"`
	Row_expiry_date         *time.Time `json:"row_expiry_date"`
	Row_quality             *string    `json:"row_quality"`
}
