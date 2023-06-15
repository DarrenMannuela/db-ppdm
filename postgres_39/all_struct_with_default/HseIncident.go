package dto

type Hse_incident struct {
	Incident_id            string   `json:"incident_id" default:"source"`
	Active_ind             *string  `json:"active_ind" default:""`
	Effective_date         *string  `json:"effective_date" default:""`
	Expiry_date            *string  `json:"expiry_date" default:""`
	Incident_class_id      *string  `json:"incident_class_id" default:""`
	Incident_date          *string  `json:"incident_date" default:""`
	Incident_duration      *float64 `json:"incident_duration" default:""`
	Incident_duration_ouom *string  `json:"incident_duration_ouom" default:""`
	Incident_duration_uom  *string  `json:"incident_duration_uom" default:""`
	Lost_time_ind          *string  `json:"lost_time_ind" default:""`
	Ppdm_guid              *string  `json:"ppdm_guid" default:""`
	Recorded_time          *string  `json:"recorded_time" default:""`
	Recorded_timezone      *string  `json:"recorded_timezone" default:""`
	Remark                 *string  `json:"remark" default:""`
	Reported_by_ba_id      *string  `json:"reported_by_ba_id" default:""`
	Reported_by_name       *string  `json:"reported_by_name" default:""`
	Reported_ind           *string  `json:"reported_ind" default:""`
	Source                 *string  `json:"source" default:""`
	Work_related_ind       *string  `json:"work_related_ind" default:""`
	Row_changed_by         *string  `json:"row_changed_by" default:""`
	Row_changed_date       *string  `json:"row_changed_date" default:""`
	Row_created_by         *string  `json:"row_created_by" default:""`
	Row_created_date       *string  `json:"row_created_date" default:""`
	Row_effective_date     *string  `json:"row_effective_date" default:""`
	Row_expiry_date        *string  `json:"row_expiry_date" default:""`
	Row_quality            *string  `json:"row_quality" default:""`
}
