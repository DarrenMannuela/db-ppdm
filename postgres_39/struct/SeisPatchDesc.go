package dto

import (
	"time"
)

type Seis_patch_desc struct {
	Patch_id           string     `json:"patch_id"`
	Patch_desc_id      string     `json:"patch_desc_id"`
	Active_ind         *string    `json:"active_ind"`
	Effective_date     *time.Time `json:"effective_date"`
	End_channel        *float64   `json:"end_channel"`
	End_station_id     *string    `json:"end_station_id"`
	End_x_offset       *float64   `json:"end_x_offset"`
	End_y_offset       *float64   `json:"end_y_offset"`
	End_z_offset       *float64   `json:"end_z_offset"`
	Expiry_date        *time.Time `json:"expiry_date"`
	Offset_ouom        *string    `json:"offset_ouom"`
	Ppdm_guid          string     `json:"ppdm_guid"`
	Recorded_line      *string    `json:"recorded_line"`
	Remark             *string    `json:"remark"`
	Source             *string    `json:"source"`
	Start_channel      *float64   `json:"start_channel"`
	Start_station_id   *string    `json:"start_station_id"`
	Start_x_offset     *float64   `json:"start_x_offset"`
	Start_y_offset     *float64   `json:"start_y_offset"`
	Start_z_offset     *float64   `json:"start_z_offset"`
	Row_changed_by     *string    `json:"row_changed_by"`
	Row_changed_date   *time.Time `json:"row_changed_date"`
	Row_created_by     *string    `json:"row_created_by"`
	Row_created_date   *time.Time `json:"row_created_date"`
	Row_effective_date *time.Time `json:"row_effective_date"`
	Row_expiry_date    *time.Time `json:"row_expiry_date"`
	Row_quality        *string    `json:"row_quality"`
}
