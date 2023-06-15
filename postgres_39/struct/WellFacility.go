package dto

import (
	"time"
)

type Well_facility struct {
	Uwi                string     `json:"uwi"`
	Facility_id        string     `json:"facility_id"`
	Facility_type      string     `json:"facility_type"`
	Active_obs_no      int        `json:"active_obs_no"`
	Active_ind         *string    `json:"active_ind"`
	Effective_date     *time.Time `json:"effective_date"`
	Expiry_date        *time.Time `json:"expiry_date"`
	Facility_use_type  *string    `json:"facility_use_type"`
	Ppdm_guid          string     `json:"ppdm_guid"`
	Remark             *string    `json:"remark"`
	Single_source_ind  *string    `json:"single_source_ind"`
	Source             *string    `json:"source"`
	String_id          *string    `json:"string_id"`
	String_source      *string    `json:"string_source"`
	Row_changed_by     *string    `json:"row_changed_by"`
	Row_changed_date   *time.Time `json:"row_changed_date"`
	Row_created_by     *string    `json:"row_created_by"`
	Row_created_date   *time.Time `json:"row_created_date"`
	Row_effective_date *time.Time `json:"row_effective_date"`
	Row_expiry_date    *time.Time `json:"row_expiry_date"`
	Row_quality        *string    `json:"row_quality"`
}
