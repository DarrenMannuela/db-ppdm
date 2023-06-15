package dto

import (
	"time"
)

type Strat_acqtn_method struct {
	Strat_acqtn_method_id string     `json:"strat_acqtn_method_id"`
	Acqtn_date            *time.Time `json:"acqtn_date"`
	Acqtn_method_type     *string    `json:"acqtn_method_type"`
	Active_ind            *string    `json:"active_ind"`
	Effective_date        *time.Time `json:"effective_date"`
	Expiry_date           *time.Time `json:"expiry_date"`
	Field_station_id      *string    `json:"field_station_id"`
	Interp_id             *string    `json:"interp_id"`
	Ppdm_guid             string     `json:"ppdm_guid"`
	Remark                *string    `json:"remark"`
	Source                *string    `json:"source"`
	Strat_column_id       *string    `json:"strat_column_id"`
	Strat_column_source   *string    `json:"strat_column_source"`
	Strat_name_set_id     *string    `json:"strat_name_set_id"`
	Strat_unit_id         *string    `json:"strat_unit_id"`
	Uwi                   *string    `json:"uwi"`
	Row_changed_by        *string    `json:"row_changed_by"`
	Row_changed_date      *time.Time `json:"row_changed_date"`
	Row_created_by        *string    `json:"row_created_by"`
	Row_created_date      *time.Time `json:"row_created_date"`
	Row_effective_date    *time.Time `json:"row_effective_date"`
	Row_expiry_date       *time.Time `json:"row_expiry_date"`
	Row_quality           *string    `json:"row_quality"`
}
