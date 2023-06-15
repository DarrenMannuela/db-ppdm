package dto

import (
	"time"
)

type Rest_remark struct {
	Restriction_id            string     `json:"restriction_id"`
	Restriction_version       int        `json:"restriction_version"`
	Restriction_remark_id     string     `json:"restriction_remark_id"`
	Restriction_remark_seq_no int        `json:"restriction_remark_seq_no"`
	Active_ind                *string    `json:"active_ind"`
	Business_associate_id     *string    `json:"business_associate_id"`
	Effective_date            *time.Time `json:"effective_date"`
	Expiry_date               *time.Time `json:"expiry_date"`
	Ppdm_guid                 string     `json:"ppdm_guid"`
	Remark                    *string    `json:"remark"`
	Remark_code               *string    `json:"remark_code"`
	Remark_date               *time.Time `json:"remark_date"`
	Source                    *string    `json:"source"`
	Row_changed_by            *string    `json:"row_changed_by"`
	Row_changed_date          *time.Time `json:"row_changed_date"`
	Row_created_by            *string    `json:"row_created_by"`
	Row_created_date          *time.Time `json:"row_created_date"`
	Row_effective_date        *time.Time `json:"row_effective_date"`
	Row_expiry_date           *time.Time `json:"row_expiry_date"`
	Row_quality               *string    `json:"row_quality"`
}
