package dto

import (
	"time"
)

type Land_right_instrument struct {
	Land_right_subtype      string     `json:"land_right_subtype"`
	Land_right_id           string     `json:"land_right_id"`
	Instrument_id           string     `json:"instrument_id"`
	Lr_inst_seq_no          int        `json:"lr_inst_seq_no"`
	Active_ind              *string    `json:"active_ind"`
	Discharge_date          *time.Time `json:"discharge_date"`
	Discharge_ind           *string    `json:"discharge_ind"`
	Discharge_instrument_id *string    `json:"discharge_instrument_id"`
	Discharge_num           *string    `json:"discharge_num"`
	Effective_date          *time.Time `json:"effective_date"`
	Expiry_date             *time.Time `json:"expiry_date"`
	Ppdm_guid               string     `json:"ppdm_guid"`
	Remark                  *string    `json:"remark"`
	Source                  *string    `json:"source"`
	Row_changed_by          *string    `json:"row_changed_by"`
	Row_changed_date        *time.Time `json:"row_changed_date"`
	Row_created_by          *string    `json:"row_created_by"`
	Row_created_date        *time.Time `json:"row_created_date"`
	Row_effective_date      *time.Time `json:"row_effective_date"`
	Row_expiry_date         *time.Time `json:"row_expiry_date"`
	Row_quality             *string    `json:"row_quality"`
}
