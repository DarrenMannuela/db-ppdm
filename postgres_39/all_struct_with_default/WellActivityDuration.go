package dto

type Well_activity_duration struct {
	Uwi                    string   `json:"uwi" default:"source"`
	Source                 string   `json:"source" default:"source"`
	Activity_obs_no        int      `json:"activity_obs_no" default:"1"`
	Duration_obs_no        int      `json:"duration_obs_no" default:"1"`
	Active_ind             *string  `json:"active_ind" default:""`
	Activity_duration      *float64 `json:"activity_duration" default:""`
	Activity_duration_ouom *string  `json:"activity_duration_ouom" default:""`
	Downtime_type          *string  `json:"downtime_type" default:""`
	Effective_date         *string  `json:"effective_date" default:""`
	End_time               *string  `json:"end_time" default:""`
	End_timezone           *string  `json:"end_timezone" default:""`
	Event_date             *string  `json:"event_date" default:""`
	Expiry_date            *string  `json:"expiry_date" default:""`
	Period_obs_no          *int     `json:"period_obs_no" default:""`
	Ppdm_guid              *string  `json:"ppdm_guid" default:""`
	Remark                 *string  `json:"remark" default:""`
	Start_time             *string  `json:"start_time" default:""`
	Start_timezone         *string  `json:"start_timezone" default:""`
	Row_changed_by         *string  `json:"row_changed_by" default:""`
	Row_changed_date       *string  `json:"row_changed_date" default:""`
	Row_created_by         *string  `json:"row_created_by" default:""`
	Row_created_date       *string  `json:"row_created_date" default:""`
	Row_effective_date     *string  `json:"row_effective_date" default:""`
	Row_expiry_date        *string  `json:"row_expiry_date" default:""`
	Row_quality            *string  `json:"row_quality" default:""`
}
