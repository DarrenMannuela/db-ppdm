package dto

import (
	"time"
)

type Well_license_violation struct {
	Uwi                string     `json:"uwi"`
	Well_lic_source    string     `json:"well_lic_source"`
	License_id         string     `json:"license_id"`
	Violation_id       string     `json:"violation_id"`
	Active_ind         *string    `json:"active_ind"`
	Condition_id       *string    `json:"condition_id"`
	Effective_date     *time.Time `json:"effective_date"`
	Expiry_date        *time.Time `json:"expiry_date"`
	Ppdm_guid          string     `json:"ppdm_guid"`
	Remark             *string    `json:"remark"`
	Resolve_date       *time.Time `json:"resolve_date"`
	Resolve_desc       *string    `json:"resolve_desc"`
	Resolve_type       *string    `json:"resolve_type"`
	Source             *string    `json:"source"`
	Violation_date     *time.Time `json:"violation_date"`
	Violation_desc     *string    `json:"violation_desc"`
	Violation_type     *string    `json:"violation_type"`
	Row_changed_by     *string    `json:"row_changed_by"`
	Row_changed_date   *time.Time `json:"row_changed_date"`
	Row_created_by     *string    `json:"row_created_by"`
	Row_created_date   *time.Time `json:"row_created_date"`
	Row_effective_date *time.Time `json:"row_effective_date"`
	Row_expiry_date    *time.Time `json:"row_expiry_date"`
	Row_quality        *string    `json:"row_quality"`
}
