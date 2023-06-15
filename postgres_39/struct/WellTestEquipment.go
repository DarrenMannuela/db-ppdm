package dto

import (
	"time"
)

type Well_test_equipment struct {
	Uwi                   string     `json:"uwi"`
	Source                string     `json:"source"`
	Test_type             string     `json:"test_type"`
	Run_num               string     `json:"run_num"`
	Test_num              string     `json:"test_num"`
	Equipment_type        string     `json:"equipment_type"`
	Equip_obs_no          int        `json:"equip_obs_no"`
	Active_ind            *string    `json:"active_ind"`
	Effective_date        *time.Time `json:"effective_date"`
	Equipment_id          *string    `json:"equipment_id"`
	Equip_length          *float64   `json:"equip_length"`
	Equip_length_ouom     *string    `json:"equip_length_ouom"`
	Equip_weight          *float64   `json:"equip_weight"`
	Equip_weight_ouom     *string    `json:"equip_weight_ouom"`
	Expiry_date           *time.Time `json:"expiry_date"`
	Inside_diameter       *float64   `json:"inside_diameter"`
	Inside_diameter_ouom  *string    `json:"inside_diameter_ouom"`
	Outside_diameter      *float64   `json:"outside_diameter"`
	Outside_diameter_ouom *string    `json:"outside_diameter_ouom"`
	Ppdm_guid             string     `json:"ppdm_guid"`
	Remark                *string    `json:"remark"`
	Top_depth             *float64   `json:"top_depth"`
	Top_depth_ouom        *string    `json:"top_depth_ouom"`
	Row_changed_by        *string    `json:"row_changed_by"`
	Row_changed_date      *time.Time `json:"row_changed_date"`
	Row_created_by        *string    `json:"row_created_by"`
	Row_created_date      *time.Time `json:"row_created_date"`
	Row_effective_date    *time.Time `json:"row_effective_date"`
	Row_expiry_date       *time.Time `json:"row_expiry_date"`
	Row_quality           *string    `json:"row_quality"`
}
