package dto

import (
	"time"
)

type Legal_nts_loc struct {
	Legal_nts_loc_id   string     `json:"legal_nts_loc_id"`
	Location_type      string     `json:"location_type"`
	Source             string     `json:"source"`
	Active_ind         *string    `json:"active_ind"`
	Block              *string    `json:"block"`
	Coord_system_id    *string    `json:"coord_system_id"`
	Effective_date     *time.Time `json:"effective_date"`
	Event_sequence     *string    `json:"event_sequence"`
	Expiry_date        *time.Time `json:"expiry_date"`
	Letter_quadrangle  *string    `json:"letter_quadrangle"`
	Nts_loc_exception  *string    `json:"nts_loc_exception"`
	Ppdm_guid          string     `json:"ppdm_guid"`
	Primary_quadrangle *float64   `json:"primary_quadrangle"`
	Quarter_unit       *string    `json:"quarter_unit"`
	Remark             *string    `json:"remark"`
	Sixteenth          *float64   `json:"sixteenth"`
	Unit               *int       `json:"unit"`
	Uwi                *string    `json:"uwi"`
	Well_node_id       *string    `json:"well_node_id"`
	Row_changed_by     *string    `json:"row_changed_by"`
	Row_changed_date   *time.Time `json:"row_changed_date"`
	Row_created_by     *string    `json:"row_created_by"`
	Row_created_date   *time.Time `json:"row_created_date"`
	Row_effective_date *time.Time `json:"row_effective_date"`
	Row_expiry_date    *time.Time `json:"row_expiry_date"`
	Row_quality        *string    `json:"row_quality"`
}
