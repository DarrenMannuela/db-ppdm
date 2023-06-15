package dto

import (
	"time"
)

type Field struct {
	Field_id           string     `json:"field_id"`
	Active_ind         *string    `json:"active_ind"`
	Area_id            *string    `json:"area_id"`
	Area_type          *string    `json:"area_type"`
	Discovery_date     *time.Time `json:"discovery_date"`
	Effective_date     *time.Time `json:"effective_date"`
	Expiry_date        *time.Time `json:"expiry_date"`
	Field_abbreviation *string    `json:"field_abbreviation"`
	Field_area_obs_no  *int       `json:"field_area_obs_no"`
	Field_name         *string    `json:"field_name"`
	Field_type         *string    `json:"field_type"`
	Group_field_id     *string    `json:"group_field_id"`
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
