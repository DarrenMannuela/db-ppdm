package dto

import (
	"time"
)

type Seis_recvr_make struct {
	Seis_rcvr_make       string     `json:"seis_rcvr_make"`
	Abbreviation         *string    `json:"abbreviation"`
	Active_ind           *string    `json:"active_ind"`
	Base_horizontal_freq *float64   `json:"base_horizontal_freq"`
	Base_vertical_freq   *float64   `json:"base_vertical_freq"`
	Effective_date       *time.Time `json:"effective_date"`
	Expiry_date          *time.Time `json:"expiry_date"`
	Long_name            *string    `json:"long_name"`
	Ppdm_guid            string     `json:"ppdm_guid"`
	Remark               *string    `json:"remark"`
	Seis_rcvr_type       *string    `json:"seis_rcvr_type"`
	Short_name           *string    `json:"short_name"`
	Source               *string    `json:"source"`
	Row_changed_by       *string    `json:"row_changed_by"`
	Row_changed_date     *time.Time `json:"row_changed_date"`
	Row_created_by       *string    `json:"row_created_by"`
	Row_created_date     *time.Time `json:"row_created_date"`
	Row_effective_date   *time.Time `json:"row_effective_date"`
	Row_expiry_date      *time.Time `json:"row_expiry_date"`
	Row_quality          *string    `json:"row_quality"`
}
