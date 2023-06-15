package dto

import (
	"time"
)

type Pden_flow_measurement struct {
	Pden_id                string     `json:"pden_id"`
	Pden_subtype           string     `json:"pden_subtype"`
	Pden_source            string     `json:"pden_source"`
	Product_type           string     `json:"product_type"`
	Amendment_seq_no       int        `json:"amendment_seq_no"`
	Period_type            string     `json:"period_type"`
	Measurement_obs_no     int        `json:"measurement_obs_no"`
	Active_ind             *string    `json:"active_ind"`
	Amend_reason           *string    `json:"amend_reason"`
	Casing_pressure        *float64   `json:"casing_pressure"`
	Casing_pressure_ouom   *string    `json:"casing_pressure_ouom"`
	Choke_position         *float64   `json:"choke_position"`
	Choke_size             *float64   `json:"choke_size"`
	Choke_size_ouom        *string    `json:"choke_size_ouom"`
	Differential_pressure  *float64   `json:"differential_pressure"`
	Diff_pressure_ouom     *string    `json:"diff_pressure_ouom"`
	Effective_date         *time.Time `json:"effective_date"`
	Expiry_date            *time.Time `json:"expiry_date"`
	Flow_rate              *float64   `json:"flow_rate"`
	Flow_rate_ouom         *string    `json:"flow_rate_ouom"`
	Measurement_date       *time.Time `json:"measurement_date"`
	Measurement_date_desc  *string    `json:"measurement_date_desc"`
	Measurement_time       *float64   `json:"measurement_time"`
	Measurement_timezone   *string    `json:"measurement_timezone"`
	Measurement_type       *string    `json:"measurement_type"`
	Meas_temperature       *float64   `json:"meas_temperature"`
	Meas_temperature_ouom  *string    `json:"meas_temperature_ouom"`
	Posted_date            *time.Time `json:"posted_date"`
	Ppdm_guid              string     `json:"ppdm_guid"`
	Production_volume      *float64   `json:"production_volume"`
	Production_volume_ouom *string    `json:"production_volume_ouom"`
	Production_volume_uom  *string    `json:"production_volume_uom"`
	Prod_time_elapsed      *float64   `json:"prod_time_elapsed"`
	Prod_time_elapsed_ouom *string    `json:"prod_time_elapsed_ouom"`
	Remark                 *string    `json:"remark"`
	Source                 *string    `json:"source"`
	Static_pressure        *float64   `json:"static_pressure"`
	Static_pressure_ouom   *string    `json:"static_pressure_ouom"`
	Tubing_pressure        *float64   `json:"tubing_pressure"`
	Tubing_pressure_ouom   *string    `json:"tubing_pressure_ouom"`
	Row_changed_by         *string    `json:"row_changed_by"`
	Row_changed_date       *time.Time `json:"row_changed_date"`
	Row_created_by         *string    `json:"row_created_by"`
	Row_created_date       *time.Time `json:"row_created_date"`
	Row_effective_date     *time.Time `json:"row_effective_date"`
	Row_expiry_date        *time.Time `json:"row_expiry_date"`
	Row_quality            *string    `json:"row_quality"`
}
