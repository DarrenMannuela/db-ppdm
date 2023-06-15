package dto

type Seis_point_summary struct {
	Seis_set_subtype    string   `json:"seis_set_subtype" default:"source"`
	Seis_set_id         string   `json:"seis_set_id" default:"source"`
	Seis_summary_obs_no int      `json:"seis_summary_obs_no" default:"1"`
	Active_ind          *string  `json:"active_ind" default:""`
	Area_size           *float64 `json:"area_size" default:""`
	Area_size_ouom      *string  `json:"area_size_ouom" default:""`
	Cdp_coverage        *float64 `json:"cdp_coverage" default:""`
	Effective_date      *string  `json:"effective_date" default:""`
	Expiry_date         *string  `json:"expiry_date" default:""`
	First_nline_no      *int     `json:"first_nline_no" default:""`
	First_seis_point_id *string  `json:"first_seis_point_id" default:""`
	First_xline_no      *int     `json:"first_xline_no" default:""`
	Last_nline_no       *int     `json:"last_nline_no" default:""`
	Last_seis_point_id  *string  `json:"last_seis_point_id" default:""`
	Last_xline_no       *int     `json:"last_xline_no" default:""`
	Line_length         *float64 `json:"line_length" default:""`
	Line_length_ouom    *string  `json:"line_length_ouom" default:""`
	Ppdm_guid           *string  `json:"ppdm_guid" default:""`
	Remark              *string  `json:"remark" default:""`
	Seis_station_type   *string  `json:"seis_station_type" default:""`
	Source              *string  `json:"source" default:""`
	Summary_type        *string  `json:"summary_type" default:""`
	Row_changed_by      *string  `json:"row_changed_by" default:""`
	Row_changed_date    *string  `json:"row_changed_date" default:""`
	Row_created_by      *string  `json:"row_created_by" default:""`
	Row_created_date    *string  `json:"row_created_date" default:""`
	Row_effective_date  *string  `json:"row_effective_date" default:""`
	Row_expiry_date     *string  `json:"row_expiry_date" default:""`
	Row_quality         *string  `json:"row_quality" default:""`
}
