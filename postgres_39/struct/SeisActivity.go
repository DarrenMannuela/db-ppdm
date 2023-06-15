package dto

import (
	"time"
)

type Seis_activity struct {
	Seis_set_subtype        string     `json:"seis_set_subtype"`
	Seis_set_id             string     `json:"seis_set_id"`
	Activity_obs_no         int        `json:"activity_obs_no"`
	Active_ind              *string    `json:"active_ind"`
	Activity_duration       *float64   `json:"activity_duration"`
	Activity_duration_ouom  *string    `json:"activity_duration_ouom"`
	Area_size               *float64   `json:"area_size"`
	Area_size_ouom          *string    `json:"area_size_ouom"`
	Crew_company_ba_id      *string    `json:"crew_company_ba_id"`
	Crew_id                 *string    `json:"crew_id"`
	Effective_date          *time.Time `json:"effective_date"`
	End_date                *time.Time `json:"end_date"`
	End_time                *time.Time `json:"end_time"`
	End_timezone            *string    `json:"end_timezone"`
	Expiry_date             *time.Time `json:"expiry_date"`
	First_nline_no          *int       `json:"first_nline_no"`
	First_seis_point_id     *string    `json:"first_seis_point_id"`
	First_xline_no          *int       `json:"first_xline_no"`
	Last_nline_no           *int       `json:"last_nline_no"`
	Last_seis_point_id      *string    `json:"last_seis_point_id"`
	Last_xline_no           *int       `json:"last_xline_no"`
	Line_length             *float64   `json:"line_length"`
	Line_length_ouom        *string    `json:"line_length_ouom"`
	Owner_ba_id             *string    `json:"owner_ba_id"`
	Ppdm_guid               string     `json:"ppdm_guid"`
	Remark                  *string    `json:"remark"`
	Seis_activity_class     *string    `json:"seis_activity_class"`
	Seis_activity_type      *string    `json:"seis_activity_type"`
	Source                  *string    `json:"source"`
	Start_date              *time.Time `json:"start_date"`
	Start_time              *time.Time `json:"start_time"`
	Start_timezone          *string    `json:"start_timezone"`
	Total_time_elapsed      *float64   `json:"total_time_elapsed"`
	Total_time_elapsed_ouom *string    `json:"total_time_elapsed_ouom"`
	Row_changed_by          *string    `json:"row_changed_by"`
	Row_changed_date        *time.Time `json:"row_changed_date"`
	Row_created_by          *string    `json:"row_created_by"`
	Row_created_date        *time.Time `json:"row_created_date"`
	Row_effective_date      *time.Time `json:"row_effective_date"`
	Row_expiry_date         *time.Time `json:"row_expiry_date"`
	Row_quality             *string    `json:"row_quality"`
}
