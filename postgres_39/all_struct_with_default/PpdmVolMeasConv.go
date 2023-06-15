package dto

type Ppdm_vol_meas_conv struct {
	Volume_regime_id    string   `json:"volume_regime_id" default:"source"`
	Conversion_quantity string   `json:"conversion_quantity" default:"source"`
	Conversion_obs_no   int      `json:"conversion_obs_no" default:"1"`
	Active_ind          *string  `json:"active_ind" default:""`
	Conversion_factor   *float64 `json:"conversion_factor" default:""`
	Conversion_formula  *string  `json:"conversion_formula" default:""`
	Effective_date      *string  `json:"effective_date" default:""`
	Expiry_date         *string  `json:"expiry_date" default:""`
	From_uom            *string  `json:"from_uom" default:""`
	Ppdm_guid           *string  `json:"ppdm_guid" default:""`
	Preferrred_ind      *string  `json:"preferrred_ind" default:""`
	Pressure            *float64 `json:"pressure" default:""`
	Pressure_uom        *string  `json:"pressure_uom" default:""`
	Remark              *string  `json:"remark" default:""`
	Source              *string  `json:"source" default:""`
	Temperature         *float64 `json:"temperature" default:""`
	Temperature_uom     *string  `json:"temperature_uom" default:""`
	To_uom              *string  `json:"to_uom" default:""`
	Row_changed_by      *string  `json:"row_changed_by" default:""`
	Row_changed_date    *string  `json:"row_changed_date" default:""`
	Row_created_by      *string  `json:"row_created_by" default:""`
	Row_created_date    *string  `json:"row_created_date" default:""`
	Row_effective_date  *string  `json:"row_effective_date" default:""`
	Row_expiry_date     *string  `json:"row_expiry_date" default:""`
	Row_quality         *string  `json:"row_quality" default:""`
}
