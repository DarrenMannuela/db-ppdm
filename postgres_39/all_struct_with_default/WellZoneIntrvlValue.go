package dto

type Well_zone_intrvl_value struct {
	Uwi                string   `json:"uwi" default:"source"`
	Interval_source    string   `json:"interval_source" default:"source"`
	Interval_id        string   `json:"interval_id" default:"source"`
	Zone_id            string   `json:"zone_id" default:"source"`
	Zone_source        string   `json:"zone_source" default:"source"`
	Zone_value_obs_no  int      `json:"zone_value_obs_no" default:"1"`
	Active_ind         *string  `json:"active_ind" default:""`
	Date_format_desc   *string  `json:"date_format_desc" default:""`
	Effective_date     *string  `json:"effective_date" default:""`
	Expiry_date        *string  `json:"expiry_date" default:""`
	Parameter          *string  `json:"parameter" default:""`
	Ppdm_guid          *string  `json:"ppdm_guid" default:""`
	Remark             *string  `json:"remark" default:""`
	Source             *string  `json:"source" default:""`
	Value_type         *string  `json:"value_type" default:""`
	Zone_value         *float64 `json:"zone_value" default:""`
	Zone_value_ouom    *string  `json:"zone_value_ouom" default:""`
	Zone_value_uom     *string  `json:"zone_value_uom" default:""`
	Row_changed_by     *string  `json:"row_changed_by" default:""`
	Row_changed_date   *string  `json:"row_changed_date" default:""`
	Row_created_by     *string  `json:"row_created_by" default:""`
	Row_created_date   *string  `json:"row_created_date" default:""`
	Row_effective_date *string  `json:"row_effective_date" default:""`
	Row_expiry_date    *string  `json:"row_expiry_date" default:""`
	Row_quality        *string  `json:"row_quality" default:""`
}
