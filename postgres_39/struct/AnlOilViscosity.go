package dto

import (
	"time"
)

type Anl_oil_viscosity struct {
	Analysis_id                 string     `json:"analysis_id"`
	Anl_source                  string     `json:"anl_source"`
	Viscosity_obs_no            int        `json:"viscosity_obs_no"`
	Active_ind                  *string    `json:"active_ind"`
	Effective_date              *time.Time `json:"effective_date"`
	Expiry_date                 *time.Time `json:"expiry_date"`
	Oil_temperature             *float64   `json:"oil_temperature"`
	Oil_temperature_ouom        *string    `json:"oil_temperature_ouom"`
	Oil_type                    *string    `json:"oil_type"`
	Oil_viscosity               *float64   `json:"oil_viscosity"`
	Oil_viscosity_ouom          *string    `json:"oil_viscosity_ouom"`
	Oil_viscosity_uom           *string    `json:"oil_viscosity_uom"`
	Pour_point_temperature      *float64   `json:"pour_point_temperature"`
	Pour_point_temperature_ouom *string    `json:"pour_point_temperature_ouom"`
	Ppdm_guid                   string     `json:"ppdm_guid"`
	Preferred_ind               *string    `json:"preferred_ind"`
	Remark                      *string    `json:"remark"`
	Source                      *string    `json:"source"`
	Step_seq_no                 *int       `json:"step_seq_no"`
	Viscosity_press             *float64   `json:"viscosity_press"`
	Viscosity_press_ouom        *string    `json:"viscosity_press_ouom"`
	Viscosity_temp              *float64   `json:"viscosity_temp"`
	Viscosity_temp_ouom         *string    `json:"viscosity_temp_ouom"`
	Viscosity_value             *float64   `json:"viscosity_value"`
	Viscosity_value_ouom        *string    `json:"viscosity_value_ouom"`
	Viscosity_value_uom         *string    `json:"viscosity_value_uom"`
	Row_changed_by              *string    `json:"row_changed_by"`
	Row_changed_date            *time.Time `json:"row_changed_date"`
	Row_created_by              *string    `json:"row_created_by"`
	Row_created_date            *time.Time `json:"row_created_date"`
	Row_effective_date          *time.Time `json:"row_effective_date"`
	Row_expiry_date             *time.Time `json:"row_expiry_date"`
	Row_quality                 *string    `json:"row_quality"`
}
