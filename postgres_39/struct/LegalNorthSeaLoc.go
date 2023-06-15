package dto

import (
	"time"
)

type Legal_north_sea_loc struct {
	Legal_north_sea_loc_id string     `json:"legal_north_sea_loc_id"`
	Location_type          string     `json:"location_type"`
	Source                 string     `json:"source"`
	Active_ind             *string    `json:"active_ind"`
	Block_no               *int       `json:"block_no"`
	Block_suffix           *string    `json:"block_suffix"`
	Coord_system_id        *string    `json:"coord_system_id"`
	Effective_date         *time.Time `json:"effective_date"`
	Expiry_date            *time.Time `json:"expiry_date"`
	Land_well_ind          *string    `json:"land_well_ind"`
	Ppdm_guid              string     `json:"ppdm_guid"`
	Principal_meridian     *string    `json:"principal_meridian"`
	Quadrant               *float64   `json:"quadrant"`
	Quadrant_prefix        *string    `json:"quadrant_prefix"`
	Remark                 *string    `json:"remark"`
	Uwi                    *string    `json:"uwi"`
	Well_node_id           *string    `json:"well_node_id"`
	Well_no_prefix         *string    `json:"well_no_prefix"`
	Well_no_suffix         *string    `json:"well_no_suffix"`
	Well_seq_no            *int       `json:"well_seq_no"`
	Row_changed_by         *string    `json:"row_changed_by"`
	Row_changed_date       *time.Time `json:"row_changed_date"`
	Row_created_by         *string    `json:"row_created_by"`
	Row_created_date       *time.Time `json:"row_created_date"`
	Row_effective_date     *time.Time `json:"row_effective_date"`
	Row_expiry_date        *time.Time `json:"row_expiry_date"`
	Row_quality            *string    `json:"row_quality"`
}
