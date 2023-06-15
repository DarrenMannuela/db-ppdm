package dto

type Well_drill_statistic struct {
	Uwi                  string   `json:"uwi" default:"source"`
	Period_obs_no        int      `json:"period_obs_no" default:"1"`
	Statistic_type       string   `json:"statistic_type" default:"source"`
	Statistic_obs_no     int      `json:"statistic_obs_no" default:"1"`
	Active_ind           *string  `json:"active_ind" default:""`
	Currency_conversion  *float64 `json:"currency_conversion" default:""`
	Date_format_desc     *string  `json:"date_format_desc" default:""`
	Description          *string  `json:"description" default:""`
	Effective_date       *string  `json:"effective_date" default:""`
	Expiry_date          *string  `json:"expiry_date" default:""`
	Max_value            *float64 `json:"max_value" default:""`
	Max_value_ouom       *string  `json:"max_value_ouom" default:""`
	Max_value_uom        *string  `json:"max_value_uom" default:""`
	Min_value            *float64 `json:"min_value" default:""`
	Min_value_ouom       *string  `json:"min_value_ouom" default:""`
	Min_value_uom        *string  `json:"min_value_uom" default:""`
	Ppdm_guid            *string  `json:"ppdm_guid" default:""`
	Remark               *string  `json:"remark" default:""`
	Source               *string  `json:"source" default:""`
	Statistic_code       *string  `json:"statistic_code" default:""`
	Statistic_date       *string  `json:"statistic_date" default:""`
	Statistic_time       *string  `json:"statistic_time" default:""`
	Statistic_timezone   *string  `json:"statistic_timezone" default:""`
	Statistic_value      *float64 `json:"statistic_value" default:""`
	Statistic_value_ouom *string  `json:"statistic_value_ouom" default:""`
	Statistic_value_uom  *string  `json:"statistic_value_uom" default:""`
	Row_changed_by       *string  `json:"row_changed_by" default:""`
	Row_changed_date     *string  `json:"row_changed_date" default:""`
	Row_created_by       *string  `json:"row_created_by" default:""`
	Row_created_date     *string  `json:"row_created_date" default:""`
	Row_effective_date   *string  `json:"row_effective_date" default:""`
	Row_expiry_date      *string  `json:"row_expiry_date" default:""`
	Row_quality          *string  `json:"row_quality" default:""`
}
