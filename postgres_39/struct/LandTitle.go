package dto

import (
	"time"
)

type Land_title struct {
	Land_right_subtype  string     `json:"land_right_subtype"`
	Land_right_id       string     `json:"land_right_id"`
	Active_ind          *string    `json:"active_ind"`
	Certified_desc      *string    `json:"certified_desc"`
	Effective_date      *time.Time `json:"effective_date"`
	Expiry_date         *time.Time `json:"expiry_date"`
	Ppdm_guid           string     `json:"ppdm_guid"`
	Registration_date   *time.Time `json:"registration_date"`
	Remark              *string    `json:"remark"`
	Source              *string    `json:"source"`
	Title_change_reason *string    `json:"title_change_reason"`
	Title_holder        *string    `json:"title_holder"`
	Title_number        *string    `json:"title_number"`
	Title_reference_num *string    `json:"title_reference_num"`
	Title_type          *string    `json:"title_type"`
	Row_changed_by      *string    `json:"row_changed_by"`
	Row_changed_date    *time.Time `json:"row_changed_date"`
	Row_created_by      *string    `json:"row_created_by"`
	Row_created_date    *time.Time `json:"row_created_date"`
	Row_effective_date  *time.Time `json:"row_effective_date"`
	Row_expiry_date     *time.Time `json:"row_expiry_date"`
	Row_quality         *string    `json:"row_quality"`
}
