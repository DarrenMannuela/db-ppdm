package dto

import (
	"time"
)

type Well_pressure_aof_4pt struct {
	Uwi                        string     `json:"uwi"`
	Source                     string     `json:"source"`
	Pressure_obs_no            int        `json:"pressure_obs_no"`
	Aof_obs_no                 int        `json:"aof_obs_no"`
	Point_obs_no               int        `json:"point_obs_no"`
	Active_ind                 *string    `json:"active_ind"`
	Bhfp                       *float64   `json:"bhfp"`
	Bhfp_ouom                  *string    `json:"bhfp_ouom"`
	Choke_size_desc            *string    `json:"choke_size_desc"`
	Condensate_flow_rate       *float64   `json:"condensate_flow_rate"`
	Condensate_flow_rate_ouom  *string    `json:"condensate_flow_rate_ouom"`
	Effective_date             *time.Time `json:"effective_date"`
	Expiry_date                *time.Time `json:"expiry_date"`
	Final_shutin_pressure      *float64   `json:"final_shutin_pressure"`
	Final_shutin_pressure_ouom *string    `json:"final_shutin_pressure_ouom"`
	Flow_duration              *float64   `json:"flow_duration"`
	Flow_duration_ouom         *string    `json:"flow_duration_ouom"`
	Flow_pressure              *float64   `json:"flow_pressure"`
	Flow_pressure_ouom         *string    `json:"flow_pressure_ouom"`
	Flow_rate                  *float64   `json:"flow_rate"`
	Flow_rate_ouom             *string    `json:"flow_rate_ouom"`
	Ppdm_guid                  string     `json:"ppdm_guid"`
	Remark                     *string    `json:"remark"`
	Water_flow_rate            *float64   `json:"water_flow_rate"`
	Water_flow_rate_ouom       *string    `json:"water_flow_rate_ouom"`
	Row_changed_by             *string    `json:"row_changed_by"`
	Row_changed_date           *time.Time `json:"row_changed_date"`
	Row_created_by             *string    `json:"row_created_by"`
	Row_created_date           *time.Time `json:"row_created_date"`
	Row_effective_date         *time.Time `json:"row_effective_date"`
	Row_expiry_date            *time.Time `json:"row_expiry_date"`
	Row_quality                *string    `json:"row_quality"`
}
