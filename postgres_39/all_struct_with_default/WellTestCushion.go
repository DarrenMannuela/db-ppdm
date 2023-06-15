package dto

type Well_test_cushion struct {
	Uwi                        string   `json:"uwi" default:"source"`
	Source                     string   `json:"source" default:"source"`
	Test_type                  string   `json:"test_type" default:"source"`
	Run_num                    string   `json:"run_num" default:"source"`
	Test_num                   string   `json:"test_num" default:"source"`
	Cushion_obs_no             int      `json:"cushion_obs_no" default:"1"`
	Active_ind                 *string  `json:"active_ind" default:""`
	Cushion_gas_pressure       *float64 `json:"cushion_gas_pressure" default:""`
	Cushion_gas_pressure_ouom  *string  `json:"cushion_gas_pressure_ouom" default:""`
	Cushion_inhibitor_volume   *float64 `json:"cushion_inhibitor_volume" default:""`
	Cushion_inhibitor_vol_ouom *string  `json:"cushion_inhibitor_vol_ouom" default:""`
	Cushion_length             *float64 `json:"cushion_length" default:""`
	Cushion_length_ouom        *string  `json:"cushion_length_ouom" default:""`
	Cushion_type               *string  `json:"cushion_type" default:""`
	Cushion_volume             *float64 `json:"cushion_volume" default:""`
	Cushion_volume_ouom        *string  `json:"cushion_volume_ouom" default:""`
	Effective_date             *string  `json:"effective_date" default:""`
	Expiry_date                *string  `json:"expiry_date" default:""`
	Ppdm_guid                  *string  `json:"ppdm_guid" default:""`
	Remark                     *string  `json:"remark" default:""`
	Row_changed_by             *string  `json:"row_changed_by" default:""`
	Row_changed_date           *string  `json:"row_changed_date" default:""`
	Row_created_by             *string  `json:"row_created_by" default:""`
	Row_created_date           *string  `json:"row_created_date" default:""`
	Row_effective_date         *string  `json:"row_effective_date" default:""`
	Row_expiry_date            *string  `json:"row_expiry_date" default:""`
	Row_quality                *string  `json:"row_quality" default:""`
}
