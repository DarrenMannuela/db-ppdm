package dto

import (
	"time"
)

type Land_xref struct {
	Parent_land_right_subtype string     `json:"parent_land_right_subtype"`
	Child_land_right_subtype  string     `json:"child_land_right_subtype"`
	Parent_land_right_id      string     `json:"parent_land_right_id"`
	Child_land_right_id       string     `json:"child_land_right_id"`
	Lr_xref_seq_no            int        `json:"lr_xref_seq_no"`
	Active_ind                *string    `json:"active_ind"`
	Effective_date            *time.Time `json:"effective_date"`
	Expiry_date               *time.Time `json:"expiry_date"`
	Land_xref_type            *string    `json:"land_xref_type"`
	Percent_allocation        *float64   `json:"percent_allocation"`
	Ppdm_guid                 string     `json:"ppdm_guid"`
	Remark                    *string    `json:"remark"`
	Source                    *string    `json:"source"`
	Row_changed_by            *string    `json:"row_changed_by"`
	Row_changed_date          *time.Time `json:"row_changed_date"`
	Row_created_by            *string    `json:"row_created_by"`
	Row_created_date          *time.Time `json:"row_created_date"`
	Row_effective_date        *time.Time `json:"row_effective_date"`
	Row_expiry_date           *time.Time `json:"row_expiry_date"`
	Row_quality               *string    `json:"row_quality"`
}
