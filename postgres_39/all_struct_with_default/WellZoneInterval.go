package dto

type Well_zone_interval struct {
	Uwi                string   `json:"uwi" default:"source"`
	Source             string   `json:"source" default:"source"`
	Interval_id        string   `json:"interval_id" default:"source"`
	Zone_id            string   `json:"zone_id" default:"source"`
	Zone_source        string   `json:"zone_source" default:"source"`
	Active_ind         *string  `json:"active_ind" default:""`
	Base_md            *float64 `json:"base_md" default:""`
	Base_md_ouom       *string  `json:"base_md_ouom" default:""`
	Base_tvd           *float64 `json:"base_tvd" default:""`
	Base_tvd_ouom      *string  `json:"base_tvd_ouom" default:""`
	Effective_date     *string  `json:"effective_date" default:""`
	Expiry_date        *string  `json:"expiry_date" default:""`
	Ppdm_guid          *string  `json:"ppdm_guid" default:""`
	Remark             *string  `json:"remark" default:""`
	Top_delta_x        *float64 `json:"top_delta_x" default:""`
	Top_delta_x_ouom   *string  `json:"top_delta_x_ouom" default:""`
	Top_delta_y        *float64 `json:"top_delta_y" default:""`
	Top_delta_y_ouom   *string  `json:"top_delta_y_ouom" default:""`
	Top_md             *float64 `json:"top_md" default:""`
	Top_md_ouom        *string  `json:"top_md_ouom" default:""`
	Top_tvd            *float64 `json:"top_tvd" default:""`
	Top_tvd_ouom       *string  `json:"top_tvd_ouom" default:""`
	Row_changed_by     *string  `json:"row_changed_by" default:""`
	Row_changed_date   *string  `json:"row_changed_date" default:""`
	Row_created_by     *string  `json:"row_created_by" default:""`
	Row_created_date   *string  `json:"row_created_date" default:""`
	Row_effective_date *string  `json:"row_effective_date" default:""`
	Row_expiry_date    *string  `json:"row_expiry_date" default:""`
	Row_quality        *string  `json:"row_quality" default:""`
}
