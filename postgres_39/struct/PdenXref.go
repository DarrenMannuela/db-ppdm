package dto

import (
	"time"
)

type Pden_xref struct {
	From_pden_subtype  string     `json:"from_pden_subtype"`
	From_pden_id       string     `json:"from_pden_id"`
	From_pden_source   string     `json:"from_pden_source"`
	To_pden_subtype    string     `json:"to_pden_subtype"`
	To_pden_id         string     `json:"to_pden_id"`
	To_pden_source     string     `json:"to_pden_source"`
	Xref_obs_no        int        `json:"xref_obs_no"`
	Active_ind         *string    `json:"active_ind"`
	Effective_date     *time.Time `json:"effective_date"`
	Expiry_date        *time.Time `json:"expiry_date"`
	Pden_xref_type     *string    `json:"pden_xref_type"`
	Ppdm_guid          string     `json:"ppdm_guid"`
	Remark             *string    `json:"remark"`
	Source             *string    `json:"source"`
	Row_changed_by     *string    `json:"row_changed_by"`
	Row_changed_date   *time.Time `json:"row_changed_date"`
	Row_created_by     *string    `json:"row_created_by"`
	Row_created_date   *time.Time `json:"row_created_date"`
	Row_effective_date *time.Time `json:"row_effective_date"`
	Row_expiry_date    *time.Time `json:"row_expiry_date"`
	Row_quality        *string    `json:"row_quality"`
}
