package dto

type Well_sieve_analysis struct {
	Uwi                  string   `json:"uwi" default:"source"`
	Source               string   `json:"source" default:"source"`
	Analysis_obs_no      int      `json:"analysis_obs_no" default:"1"`
	Active_ind           *string  `json:"active_ind" default:""`
	Analysis_date        *string  `json:"analysis_date" default:""`
	Base_depth           *float64 `json:"base_depth" default:""`
	Base_depth_ouom      *string  `json:"base_depth_ouom" default:""`
	Effective_date       *string  `json:"effective_date" default:""`
	Expiry_date          *string  `json:"expiry_date" default:""`
	Interval_depth       *float64 `json:"interval_depth" default:""`
	Interval_depth_ouom  *string  `json:"interval_depth_ouom" default:""`
	Interval_length      *float64 `json:"interval_length" default:""`
	Interval_length_ouom *string  `json:"interval_length_ouom" default:""`
	Laboratory           *string  `json:"laboratory" default:""`
	Lab_file_num         *string  `json:"lab_file_num" default:""`
	Ppdm_guid            *string  `json:"ppdm_guid" default:""`
	Remark               *string  `json:"remark" default:""`
	Top_depth            *float64 `json:"top_depth" default:""`
	Top_depth_ouom       *string  `json:"top_depth_ouom" default:""`
	Row_changed_by       *string  `json:"row_changed_by" default:""`
	Row_changed_date     *string  `json:"row_changed_date" default:""`
	Row_created_by       *string  `json:"row_created_by" default:""`
	Row_created_date     *string  `json:"row_created_date" default:""`
	Row_effective_date   *string  `json:"row_effective_date" default:""`
	Row_expiry_date      *string  `json:"row_expiry_date" default:""`
	Row_quality          *string  `json:"row_quality" default:""`
}
