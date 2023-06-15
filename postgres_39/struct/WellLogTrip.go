package dto

import (
	"time"
)

type Well_log_trip struct {
	Uwi                         string     `json:"uwi"`
	Source                      string     `json:"source"`
	Job_id                      string     `json:"job_id"`
	Trip_obs_no                 int        `json:"trip_obs_no"`
	Active_ind                  *string    `json:"active_ind"`
	Base_depth                  *float64   `json:"base_depth"`
	Base_depth_ouom             *string    `json:"base_depth_ouom"`
	Base_strat_unit_id          *string    `json:"base_strat_unit_id"`
	Effective_date              *time.Time `json:"effective_date"`
	Expiry_date                 *time.Time `json:"expiry_date"`
	Logging_service_type        *string    `json:"logging_service_type"`
	Max_depth                   *float64   `json:"max_depth"`
	Max_depth_ouom              *string    `json:"max_depth_ouom"`
	Max_temperature             *float64   `json:"max_temperature"`
	Max_temperature_ouom        *string    `json:"max_temperature_ouom"`
	Mud_sample_id               *string    `json:"mud_sample_id"`
	Mud_sample_type             *string    `json:"mud_sample_type"`
	Mud_source                  *string    `json:"mud_source"`
	Observer                    *string    `json:"observer"`
	On_bottom_date              *time.Time `json:"on_bottom_date"`
	On_bottom_time              *time.Time `json:"on_bottom_time"`
	On_bottom_timezone          *string    `json:"on_bottom_timezone"`
	Ppdm_guid                   string     `json:"ppdm_guid"`
	Remark                      *string    `json:"remark"`
	Reported_tvd                *float64   `json:"reported_tvd"`
	Reported_tvd_ouom           *string    `json:"reported_tvd_ouom"`
	Report_apd                  *float64   `json:"report_apd"`
	Report_log_datum            *string    `json:"report_log_datum"`
	Report_log_datum_elev       *float64   `json:"report_log_datum_elev"`
	Report_log_datum_elev_ouom  *string    `json:"report_log_datum_elev_ouom"`
	Report_log_run              *string    `json:"report_log_run"`
	Report_perm_datum           *string    `json:"report_perm_datum"`
	Report_perm_datum_elev      *float64   `json:"report_perm_datum_elev"`
	Report_perm_datum_elev_ouom *string    `json:"report_perm_datum_elev_ouom"`
	Strat_name_set_id           *string    `json:"strat_name_set_id"`
	Top_depth                   *float64   `json:"top_depth"`
	Top_depth_ouom              *string    `json:"top_depth_ouom"`
	Top_strat_unit_id           *string    `json:"top_strat_unit_id"`
	Trip_date                   *time.Time `json:"trip_date"`
	Tubing_bottom_depth         *float64   `json:"tubing_bottom_depth"`
	Tubing_bottom_depth_ouom    *string    `json:"tubing_bottom_depth_ouom"`
	Row_changed_by              *string    `json:"row_changed_by"`
	Row_changed_date            *time.Time `json:"row_changed_date"`
	Row_created_by              *string    `json:"row_created_by"`
	Row_created_date            *time.Time `json:"row_created_date"`
	Row_effective_date          *time.Time `json:"row_effective_date"`
	Row_expiry_date             *time.Time `json:"row_expiry_date"`
	Row_quality                 *string    `json:"row_quality"`
}
