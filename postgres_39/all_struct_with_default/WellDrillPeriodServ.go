package dto

type Well_drill_period_serv struct {
	Uwi                string   `json:"uwi" default:"source"`
	Period_obs_no      int      `json:"period_obs_no" default:"1"`
	Provided_by        string   `json:"provided_by" default:"source"`
	Service_seq_no     int      `json:"service_seq_no" default:"1"`
	Segment_obs_no     int      `json:"segment_obs_no" default:"1"`
	Active_ind         *string  `json:"active_ind" default:""`
	Description        *string  `json:"description" default:""`
	Effective_date     *string  `json:"effective_date" default:""`
	End_time           *string  `json:"end_time" default:""`
	End_timezone       *string  `json:"end_timezone" default:""`
	Expiry_date        *string  `json:"expiry_date" default:""`
	Metric_code        *string  `json:"metric_code" default:""`
	Metric_max_value   *float64 `json:"metric_max_value" default:""`
	Metric_min_value   *float64 `json:"metric_min_value" default:""`
	Metric_value       *float64 `json:"metric_value" default:""`
	Metric_value_ouom  *string  `json:"metric_value_ouom" default:""`
	Metric_value_uom   *string  `json:"metric_value_uom" default:""`
	Ppdm_guid          *string  `json:"ppdm_guid" default:""`
	Remark             *string  `json:"remark" default:""`
	Service_desc       *string  `json:"service_desc" default:""`
	Service_metric     *string  `json:"service_metric" default:""`
	Source             *string  `json:"source" default:""`
	Start_time         *string  `json:"start_time" default:""`
	Start_timezone     *string  `json:"start_timezone" default:""`
	Row_changed_by     *string  `json:"row_changed_by" default:""`
	Row_changed_date   *string  `json:"row_changed_date" default:""`
	Row_created_by     *string  `json:"row_created_by" default:""`
	Row_created_date   *string  `json:"row_created_date" default:""`
	Row_effective_date *string  `json:"row_effective_date" default:""`
	Row_expiry_date    *string  `json:"row_expiry_date" default:""`
	Row_quality        *string  `json:"row_quality" default:""`
}
