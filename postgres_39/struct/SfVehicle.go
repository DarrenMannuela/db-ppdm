package dto

import (
	"time"
)

type Sf_vehicle struct {
	Sf_subtype             string     `json:"sf_subtype"`
	Vehicle_id             string     `json:"vehicle_id"`
	Active_ind             *string    `json:"active_ind"`
	Current_owner          *string    `json:"current_owner"`
	Effective_date         *time.Time `json:"effective_date"`
	Engine_size            *float64   `json:"engine_size"`
	Engine_size_ouom       *string    `json:"engine_size_ouom"`
	Expiry_date            *time.Time `json:"expiry_date"`
	License_expiry_date    *time.Time `json:"license_expiry_date"`
	License_num            *string    `json:"license_num"`
	License_register_ba_id *string    `json:"license_register_ba_id"`
	Ppdm_guid              string     `json:"ppdm_guid"`
	Remark                 *string    `json:"remark"`
	Source                 *string    `json:"source"`
	Vehicle_name           *string    `json:"vehicle_name"`
	Vehicle_type           *string    `json:"vehicle_type"`
	Row_changed_by         *string    `json:"row_changed_by"`
	Row_changed_date       *time.Time `json:"row_changed_date"`
	Row_created_by         *string    `json:"row_created_by"`
	Row_created_date       *time.Time `json:"row_created_date"`
	Row_effective_date     *time.Time `json:"row_effective_date"`
	Row_expiry_date        *time.Time `json:"row_expiry_date"`
	Row_quality            *string    `json:"row_quality"`
}
