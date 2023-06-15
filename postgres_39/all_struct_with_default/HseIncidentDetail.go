package dto

type Hse_incident_detail struct {
	Incident_id            string   `json:"incident_id" default:"source"`
	Detail_obs_no          int      `json:"detail_obs_no" default:"1"`
	Active_ind             *string  `json:"active_ind" default:""`
	Detail_type            *string  `json:"detail_type" default:""`
	Effective_date         *string  `json:"effective_date" default:""`
	Expiry_date            *string  `json:"expiry_date" default:""`
	Incident_date          *string  `json:"incident_date" default:""`
	Incident_duration      *float64 `json:"incident_duration" default:""`
	Incident_duration_ouom *string  `json:"incident_duration_ouom" default:""`
	Incident_duration_uom  *string  `json:"incident_duration_uom" default:""`
	Incident_set_id        *string  `json:"incident_set_id" default:""`
	Incident_severity_id   *string  `json:"incident_severity_id" default:""`
	Incident_type_id       *string  `json:"incident_type_id" default:""`
	Ppdm_guid              *string  `json:"ppdm_guid" default:""`
	Remark                 *string  `json:"remark" default:""`
	Source                 *string  `json:"source" default:""`
	Row_changed_by         *string  `json:"row_changed_by" default:""`
	Row_changed_date       *string  `json:"row_changed_date" default:""`
	Row_created_by         *string  `json:"row_created_by" default:""`
	Row_created_date       *string  `json:"row_created_date" default:""`
	Row_effective_date     *string  `json:"row_effective_date" default:""`
	Row_expiry_date        *string  `json:"row_expiry_date" default:""`
	Row_quality            *string  `json:"row_quality" default:""`
}
