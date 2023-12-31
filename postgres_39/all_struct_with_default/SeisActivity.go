package dto

type Seis_activity struct {
	Seis_set_subtype        string   `json:"seis_set_subtype" default:"source"`
	Seis_set_id             string   `json:"seis_set_id" default:"source"`
	Activity_obs_no         int      `json:"activity_obs_no" default:"1"`
	Active_ind              *string  `json:"active_ind" default:""`
	Activity_duration       *float64 `json:"activity_duration" default:""`
	Activity_duration_ouom  *string  `json:"activity_duration_ouom" default:""`
	Area_size               *float64 `json:"area_size" default:""`
	Area_size_ouom          *string  `json:"area_size_ouom" default:""`
	Crew_company_ba_id      *string  `json:"crew_company_ba_id" default:""`
	Crew_id                 *string  `json:"crew_id" default:""`
	Effective_date          *string  `json:"effective_date" default:""`
	End_date                *string  `json:"end_date" default:""`
	End_time                *string  `json:"end_time" default:""`
	End_timezone            *string  `json:"end_timezone" default:""`
	Expiry_date             *string  `json:"expiry_date" default:""`
	First_nline_no          *int     `json:"first_nline_no" default:""`
	First_seis_point_id     *string  `json:"first_seis_point_id" default:""`
	First_xline_no          *int     `json:"first_xline_no" default:""`
	Last_nline_no           *int     `json:"last_nline_no" default:""`
	Last_seis_point_id      *string  `json:"last_seis_point_id" default:""`
	Last_xline_no           *int     `json:"last_xline_no" default:""`
	Line_length             *float64 `json:"line_length" default:""`
	Line_length_ouom        *string  `json:"line_length_ouom" default:""`
	Owner_ba_id             *string  `json:"owner_ba_id" default:""`
	Ppdm_guid               *string  `json:"ppdm_guid" default:""`
	Remark                  *string  `json:"remark" default:""`
	Seis_activity_class     *string  `json:"seis_activity_class" default:""`
	Seis_activity_type      *string  `json:"seis_activity_type" default:""`
	Source                  *string  `json:"source" default:""`
	Start_date              *string  `json:"start_date" default:""`
	Start_time              *string  `json:"start_time" default:""`
	Start_timezone          *string  `json:"start_timezone" default:""`
	Total_time_elapsed      *float64 `json:"total_time_elapsed" default:""`
	Total_time_elapsed_ouom *string  `json:"total_time_elapsed_ouom" default:""`
	Row_changed_by          *string  `json:"row_changed_by" default:""`
	Row_changed_date        *string  `json:"row_changed_date" default:""`
	Row_created_by          *string  `json:"row_created_by" default:""`
	Row_created_date        *string  `json:"row_created_date" default:""`
	Row_effective_date      *string  `json:"row_effective_date" default:""`
	Row_expiry_date         *string  `json:"row_expiry_date" default:""`
	Row_quality             *string  `json:"row_quality" default:""`
}
