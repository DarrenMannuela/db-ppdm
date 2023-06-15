package dto

import (
	"time"
)

type Seis_3d struct {
	Seis_set_subtype   string     `json:"seis_set_subtype"`
	Seis_3d_id         string     `json:"seis_3_d_id"`
	Active_ind         *string    `json:"active_ind"`
	Effective_date     *time.Time `json:"effective_date"`
	Expiry_date        *time.Time `json:"expiry_date"`
	Ppdm_guid          string     `json:"ppdm_guid"`
	Remark             *string    `json:"remark"`
	Seis_3d_reason     *string    `json:"seis_3_d_reason"`
	Seis_3d_type       *string    `json:"seis_3_d_type"`
	Source             *string    `json:"source"`
	Row_changed_by     *string    `json:"row_changed_by"`
	Row_changed_date   *time.Time `json:"row_changed_date"`
	Row_created_by     *string    `json:"row_created_by"`
	Row_created_date   *time.Time `json:"row_created_date"`
	Row_effective_date *time.Time `json:"row_effective_date"`
	Row_expiry_date    *time.Time `json:"row_expiry_date"`
	Row_quality        *string    `json:"row_quality"`
}
