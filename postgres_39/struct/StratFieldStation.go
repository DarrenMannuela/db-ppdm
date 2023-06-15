package dto

import (
	"time"
)

type Strat_field_station struct {
	Field_station_id   string     `json:"field_station_id"`
	Active_ind         *string    `json:"active_ind"`
	Air_photo_num      *string    `json:"air_photo_num"`
	Area_id            *string    `json:"area_id"`
	Area_type          *string    `json:"area_type"`
	Complete_date      *time.Time `json:"complete_date"`
	Description        *string    `json:"description"`
	Effective_date     *time.Time `json:"effective_date"`
	Expiry_date        *time.Time `json:"expiry_date"`
	Field_station_type *string    `json:"field_station_type"`
	Ppdm_guid          string     `json:"ppdm_guid"`
	Remark             *string    `json:"remark"`
	Source             *string    `json:"source"`
	Source_document_id *string    `json:"source_document_id"`
	Start_date         *time.Time `json:"start_date"`
	Row_changed_by     *string    `json:"row_changed_by"`
	Row_changed_date   *time.Time `json:"row_changed_date"`
	Row_created_by     *string    `json:"row_created_by"`
	Row_created_date   *time.Time `json:"row_created_date"`
	Row_effective_date *time.Time `json:"row_effective_date"`
	Row_expiry_date    *time.Time `json:"row_expiry_date"`
	Row_quality        *string    `json:"row_quality"`
}
