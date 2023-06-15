package dto

import (
	"time"
)

type Facility_xref struct {
	Facility_id        string     `json:"facility_id"`
	Facility_type      string     `json:"facility_type"`
	Facility_id_2      string     `json:"facility_id_2"`
	Facility_type_2    string     `json:"facility_type_2"`
	Xref_obs_no        int        `json:"xref_obs_no"`
	Active_ind         *string    `json:"active_ind"`
	Effective_date     *time.Time `json:"effective_date"`
	Expiry_date        *time.Time `json:"expiry_date"`
	Facility_xref_type *string    `json:"facility_xref_type"`
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
