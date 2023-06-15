package dto

import (
	"time"
)

type Ppdm_system struct {
	System_id             string     `json:"system_id"`
	Active_ind            *string    `json:"active_ind"`
	Business_owner_ba_id  *string    `json:"business_owner_ba_id"`
	Connect_string        *string    `json:"connect_string"`
	Creator_ba_id         *string    `json:"creator_ba_id"`
	Effective_date        *time.Time `json:"effective_date"`
	Expiry_date           *time.Time `json:"expiry_date"`
	Operating_system      *string    `json:"operating_system"`
	Ppdm_guid             string     `json:"ppdm_guid"`
	Ppdm_system_type      *string    `json:"ppdm_system_type"`
	Rdbms_id              *string    `json:"rdbms_id"`
	Remark                *string    `json:"remark"`
	Source                *string    `json:"source"`
	System_long_name      *string    `json:"system_long_name"`
	Technical_owner_ba_id *string    `json:"technical_owner_ba_id"`
	Version_num           *string    `json:"version_num"`
	Row_changed_by        *string    `json:"row_changed_by"`
	Row_changed_date      *time.Time `json:"row_changed_date"`
	Row_created_by        *string    `json:"row_created_by"`
	Row_created_date      *time.Time `json:"row_created_date"`
	Row_effective_date    *time.Time `json:"row_effective_date"`
	Row_expiry_date       *time.Time `json:"row_expiry_date"`
	Row_quality           *string    `json:"row_quality"`
}
