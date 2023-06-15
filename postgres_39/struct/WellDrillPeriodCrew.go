package dto

import (
	"time"
)

type Well_drill_period_crew struct {
	Uwi                     string     `json:"uwi"`
	Period_obs_no           int        `json:"period_obs_no"`
	Crew_company_ba_id      string     `json:"crew_company_ba_id"`
	Detail_obs_no           int        `json:"detail_obs_no"`
	Active_ind              *string    `json:"active_ind"`
	Booked_time_elapsed     *float64   `json:"booked_time_elapsed"`
	Booked_time_elapsed_uom *string    `json:"booked_time_elapsed_uom"`
	Crew_member_ba_id       *string    `json:"crew_member_ba_id"`
	Crew_member_name        *string    `json:"crew_member_name"`
	Crew_member_num         *string    `json:"crew_member_num"`
	Crew_member_record_ind  *string    `json:"crew_member_record_ind"`
	Crew_position           *string    `json:"crew_position"`
	Crew_record_ind         *string    `json:"crew_record_ind"`
	Crew_reference_num      *string    `json:"crew_reference_num"`
	Effective_date          *time.Time `json:"effective_date"`
	Expiry_date             *time.Time `json:"expiry_date"`
	Injury_ind              *string    `json:"injury_ind"`
	Injury_initial_ind      *string    `json:"injury_initial_ind"`
	No_injury_ind           *string    `json:"no_injury_ind"`
	Overhead_cost           *float64   `json:"overhead_cost"`
	Overhead_cost_uom       *string    `json:"overhead_cost_uom"`
	Overhead_type           *string    `json:"overhead_type"`
	Ppdm_guid               string     `json:"ppdm_guid"`
	Remark                  *string    `json:"remark"`
	Sf_subtype              *string    `json:"sf_subtype"`
	Source                  *string    `json:"source"`
	Subsistance_ind         *string    `json:"subsistance_ind"`
	Support_facility_id     *string    `json:"support_facility_id"`
	Total_crew_count        *int       `json:"total_crew_count"`
	Total_injury_count      *int       `json:"total_injury_count"`
	Row_changed_by          *string    `json:"row_changed_by"`
	Row_changed_date        *time.Time `json:"row_changed_date"`
	Row_created_by          *string    `json:"row_created_by"`
	Row_created_date        *time.Time `json:"row_created_date"`
	Row_effective_date      *time.Time `json:"row_effective_date"`
	Row_expiry_date         *time.Time `json:"row_expiry_date"`
	Row_quality             *string    `json:"row_quality"`
}
