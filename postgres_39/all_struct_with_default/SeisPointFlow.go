package dto

type Seis_point_flow struct {
	Seis_set_subtype   string   `json:"seis_set_subtype" default:"source"`
	Seis_set_id        string   `json:"seis_set_id" default:"source"`
	Seis_point_id      string   `json:"seis_point_id" default:"source"`
	Flow_id            string   `json:"flow_id" default:"source"`
	Active_ind         *string  `json:"active_ind" default:""`
	Effective_date     *string  `json:"effective_date" default:""`
	Expiry_date        *string  `json:"expiry_date" default:""`
	Flow_depth         *float64 `json:"flow_depth" default:""`
	Flow_depth_ouom    *string  `json:"flow_depth_ouom" default:""`
	Flow_duration      *float64 `json:"flow_duration" default:""`
	Flow_duration_uom  *string  `json:"flow_duration_uom" default:""`
	Flow_volume        *float64 `json:"flow_volume" default:""`
	Flow_volume_ouom   *string  `json:"flow_volume_ouom" default:""`
	Ppdm_guid          *string  `json:"ppdm_guid" default:""`
	Remark             *string  `json:"remark" default:""`
	Remedial_ba_id     *string  `json:"remedial_ba_id" default:""`
	Remedial_date      *string  `json:"remedial_date" default:""`
	Remedial_ind       *string  `json:"remedial_ind" default:""`
	Report_date        *string  `json:"report_date" default:""`
	Source             *string  `json:"source" default:""`
	Row_changed_by     *string  `json:"row_changed_by" default:""`
	Row_changed_date   *string  `json:"row_changed_date" default:""`
	Row_created_by     *string  `json:"row_created_by" default:""`
	Row_created_date   *string  `json:"row_created_date" default:""`
	Row_effective_date *string  `json:"row_effective_date" default:""`
	Row_expiry_date    *string  `json:"row_expiry_date" default:""`
	Row_quality        *string  `json:"row_quality" default:""`
}
