package dto

import (
	"time"
)

type Ppdm_vol_meas_conv struct {
	Volume_regime_id    string     `json:"volume_regime_id"`
	Conversion_quantity string     `json:"conversion_quantity"`
	Conversion_obs_no   int        `json:"conversion_obs_no"`
	Active_ind          *string    `json:"active_ind"`
	Conversion_factor   *float64   `json:"conversion_factor"`
	Conversion_formula  *string    `json:"conversion_formula"`
	Effective_date      *time.Time `json:"effective_date"`
	Expiry_date         *time.Time `json:"expiry_date"`
	From_uom            *string    `json:"from_uom"`
	Ppdm_guid           string     `json:"ppdm_guid"`
	Preferrred_ind      *string    `json:"preferrred_ind"`
	Pressure            *float64   `json:"pressure"`
	Pressure_uom        *string    `json:"pressure_uom"`
	Remark              *string    `json:"remark"`
	Source              *string    `json:"source"`
	Temperature         *float64   `json:"temperature"`
	Temperature_uom     *string    `json:"temperature_uom"`
	To_uom              *string    `json:"to_uom"`
	Row_changed_by      *string    `json:"row_changed_by"`
	Row_changed_date    *time.Time `json:"row_changed_date"`
	Row_created_by      *string    `json:"row_created_by"`
	Row_created_date    *time.Time `json:"row_created_date"`
	Row_effective_date  *time.Time `json:"row_effective_date"`
	Row_expiry_date     *time.Time `json:"row_expiry_date"`
	Row_quality         *string    `json:"row_quality"`
}
