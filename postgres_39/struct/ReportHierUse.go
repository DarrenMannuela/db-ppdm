package dto

import (
	"time"
)

type Report_hier_use struct {
	Report_hierarchy_id  string     `json:"report_hierarchy_id"`
	Component_id         string     `json:"component_id"`
	Hierarchy_use_obs_no int        `json:"hierarchy_use_obs_no"`
	Active_ind           *string    `json:"active_ind"`
	Contribution_percent *float64   `json:"contribution_percent"`
	Effective_date       *time.Time `json:"effective_date"`
	Expiry_date          *time.Time `json:"expiry_date"`
	Pden_id              *string    `json:"pden_id"`
	Pden_source          *string    `json:"pden_source"`
	Pden_subtype         *string    `json:"pden_subtype"`
	Ppdm_guid            string     `json:"ppdm_guid"`
	Remark               *string    `json:"remark"`
	Report_ind           *string    `json:"report_ind"`
	Resent_id            *string    `json:"resent_id"`
	Source               *string    `json:"source"`
	Substance_id         *string    `json:"substance_id"`
	Row_changed_by       *string    `json:"row_changed_by"`
	Row_changed_date     *time.Time `json:"row_changed_date"`
	Row_created_by       *string    `json:"row_created_by"`
	Row_created_date     *time.Time `json:"row_created_date"`
	Row_effective_date   *time.Time `json:"row_effective_date"`
	Row_expiry_date      *time.Time `json:"row_expiry_date"`
	Row_quality          *string    `json:"row_quality"`
}
