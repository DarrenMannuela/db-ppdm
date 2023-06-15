package dto

import (
	"time"
)

type Strat_column_acqtn struct {
	Strat_column_id       string     `json:"strat_column_id"`
	Strat_column_source   string     `json:"strat_column_source"`
	Strat_name_set_id     string     `json:"strat_name_set_id"`
	Strat_unit_id         string     `json:"strat_unit_id"`
	Interp_id             string     `json:"interp_id"`
	Strat_column_acqtn_id string     `json:"strat_column_acqtn_id"`
	Active_ind            *string    `json:"active_ind"`
	Core_id               *string    `json:"core_id"`
	Core_source           *string    `json:"core_source"`
	Description           *string    `json:"description"`
	Effective_date        *time.Time `json:"effective_date"`
	Expiry_date           *time.Time `json:"expiry_date"`
	Field_interp_id       *string    `json:"field_interp_id"`
	Field_station_id      *string    `json:"field_station_id"`
	Field_strat_unit_id   *string    `json:"field_strat_unit_id"`
	Information_item_id   *string    `json:"information_item_id"`
	Info_item_subtype     *string    `json:"info_item_subtype"`
	Log_curve_id          *string    `json:"log_curve_id"`
	Ppdm_guid             string     `json:"ppdm_guid"`
	Project_id            *string    `json:"project_id"`
	Remark                *string    `json:"remark"`
	Seis_set_id           *string    `json:"seis_set_id"`
	Seis_set_subtype      *string    `json:"seis_set_subtype"`
	Source                *string    `json:"source"`
	Uwi                   *string    `json:"uwi"`
	Ws_interp_id          *string    `json:"ws_interp_id"`
	Ws_strat_unit_id      *string    `json:"ws_strat_unit_id"`
	Row_changed_by        *string    `json:"row_changed_by"`
	Row_changed_date      *time.Time `json:"row_changed_date"`
	Row_created_by        *string    `json:"row_created_by"`
	Row_created_date      *time.Time `json:"row_created_date"`
	Row_effective_date    *time.Time `json:"row_effective_date"`
	Row_expiry_date       *time.Time `json:"row_expiry_date"`
	Row_quality           *string    `json:"row_quality"`
}
