package dto

type Well_test_press_meas struct {
	Uwi                           string   `json:"uwi" default:"source"`
	Source                        string   `json:"source" default:"source"`
	Test_type                     string   `json:"test_type" default:"source"`
	Run_num                       string   `json:"run_num" default:"source"`
	Test_num                      string   `json:"test_num" default:"source"`
	Measurement_obs_no            int      `json:"measurement_obs_no" default:"1"`
	Active_ind                    *string  `json:"active_ind" default:""`
	Effective_date                *string  `json:"effective_date" default:""`
	Expiry_date                   *string  `json:"expiry_date" default:""`
	Measurement_pressure          *float64 `json:"measurement_pressure" default:""`
	Measurement_pressure_ouom     *string  `json:"measurement_pressure_ouom" default:""`
	Measurement_temperature       *float64 `json:"measurement_temperature" default:""`
	Measurement_temp_ouom         *string  `json:"measurement_temp_ouom" default:""`
	Measurement_time_elapsed      *float64 `json:"measurement_time_elapsed" default:""`
	Measurement_time_elapsed_ouom *string  `json:"measurement_time_elapsed_ouom" default:""`
	Period_obs_no                 *int     `json:"period_obs_no" default:""`
	Period_type                   *string  `json:"period_type" default:""`
	Ppdm_guid                     *string  `json:"ppdm_guid" default:""`
	Recorder_id                   *string  `json:"recorder_id" default:""`
	Remark                        *string  `json:"remark" default:""`
	Row_changed_by                *string  `json:"row_changed_by" default:""`
	Row_changed_date              *string  `json:"row_changed_date" default:""`
	Row_created_by                *string  `json:"row_created_by" default:""`
	Row_created_date              *string  `json:"row_created_date" default:""`
	Row_effective_date            *string  `json:"row_effective_date" default:""`
	Row_expiry_date               *string  `json:"row_expiry_date" default:""`
	Row_quality                   *string  `json:"row_quality" default:""`
}
