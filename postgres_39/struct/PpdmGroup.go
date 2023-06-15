package dto

import (
	"time"
)

type Ppdm_group struct {
	Group_id               string     `json:"group_id"`
	Active_ind             *string    `json:"active_ind"`
	Default_unit_system_id *string    `json:"default_unit_system_id"`
	Effective_date         *time.Time `json:"effective_date"`
	Expiry_date            *time.Time `json:"expiry_date"`
	Group_name             *string    `json:"group_name"`
	Group_type             *string    `json:"group_type"`
	Ppdm_guid              string     `json:"ppdm_guid"`
	Remark                 *string    `json:"remark"`
	Report_heading1        *string    `json:"report_heading_1"`
	Report_heading2        *string    `json:"report_heading_2"`
	Screen_heading1        *string    `json:"screen_heading_1"`
	Screen_heading2        *string    `json:"screen_heading_2"`
	Source                 *string    `json:"source"`
	Sw_application_id      *string    `json:"sw_application_id"`
	Row_changed_by         *string    `json:"row_changed_by"`
	Row_changed_date       *time.Time `json:"row_changed_date"`
	Row_created_by         *string    `json:"row_created_by"`
	Row_created_date       *time.Time `json:"row_created_date"`
	Row_effective_date     *time.Time `json:"row_effective_date"`
	Row_expiry_date        *time.Time `json:"row_expiry_date"`
	Row_quality            *string    `json:"row_quality"`
}
