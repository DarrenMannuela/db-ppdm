package dto

type Well_drill_period_crew struct {
	Uwi                     string   `json:"uwi" default:"source"`
	Period_obs_no           int      `json:"period_obs_no" default:"1"`
	Crew_company_ba_id      string   `json:"crew_company_ba_id" default:"source"`
	Detail_obs_no           int      `json:"detail_obs_no" default:"1"`
	Active_ind              *string  `json:"active_ind" default:""`
	Booked_time_elapsed     *float64 `json:"booked_time_elapsed" default:""`
	Booked_time_elapsed_uom *string  `json:"booked_time_elapsed_uom" default:""`
	Crew_member_ba_id       *string  `json:"crew_member_ba_id" default:""`
	Crew_member_name        *string  `json:"crew_member_name" default:""`
	Crew_member_num         *string  `json:"crew_member_num" default:""`
	Crew_member_record_ind  *string  `json:"crew_member_record_ind" default:""`
	Crew_position           *string  `json:"crew_position" default:""`
	Crew_record_ind         *string  `json:"crew_record_ind" default:""`
	Crew_reference_num      *string  `json:"crew_reference_num" default:""`
	Effective_date          *string  `json:"effective_date" default:""`
	Expiry_date             *string  `json:"expiry_date" default:""`
	Injury_ind              *string  `json:"injury_ind" default:""`
	Injury_initial_ind      *string  `json:"injury_initial_ind" default:""`
	No_injury_ind           *string  `json:"no_injury_ind" default:""`
	Overhead_cost           *float64 `json:"overhead_cost" default:""`
	Overhead_cost_uom       *string  `json:"overhead_cost_uom" default:""`
	Overhead_type           *string  `json:"overhead_type" default:""`
	Ppdm_guid               *string  `json:"ppdm_guid" default:""`
	Remark                  *string  `json:"remark" default:""`
	Sf_subtype              *string  `json:"sf_subtype" default:""`
	Source                  *string  `json:"source" default:""`
	Subsistance_ind         *string  `json:"subsistance_ind" default:""`
	Support_facility_id     *string  `json:"support_facility_id" default:""`
	Total_crew_count        *int     `json:"total_crew_count" default:""`
	Total_injury_count      *int     `json:"total_injury_count" default:""`
	Row_changed_by          *string  `json:"row_changed_by" default:""`
	Row_changed_date        *string  `json:"row_changed_date" default:""`
	Row_created_by          *string  `json:"row_created_by" default:""`
	Row_created_date        *string  `json:"row_created_date" default:""`
	Row_effective_date      *string  `json:"row_effective_date" default:""`
	Row_expiry_date         *string  `json:"row_expiry_date" default:""`
	Row_quality             *string  `json:"row_quality" default:""`
}