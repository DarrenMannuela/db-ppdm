package dto

type Ppdm_property_column struct {
	Property_set_id        string   `json:"property_set_id" default:"source"`
	Property_obs_no        int      `json:"property_obs_no" default:"1"`
	Active_ind             *string  `json:"active_ind" default:""`
	Column_precision       *float64 `json:"column_precision" default:""`
	Column_scale           *float64 `json:"column_scale" default:""`
	Column_size            *float64 `json:"column_size" default:""`
	Data_type              *string  `json:"data_type" default:""`
	Domain                 *string  `json:"domain" default:""`
	Effective_date         *string  `json:"effective_date" default:""`
	Expiry_date            *string  `json:"expiry_date" default:""`
	Ppdm_guid              *string  `json:"ppdm_guid" default:""`
	Preferred_currency_uom *string  `json:"preferred_currency_uom" default:""`
	Preferred_uom          *string  `json:"preferred_uom" default:""`
	Quantity_type          *string  `json:"quantity_type" default:""`
	Referenced_system_id   *string  `json:"referenced_system_id" default:""`
	Referenced_table_name  *string  `json:"referenced_table_name" default:""`
	Remark                 *string  `json:"remark" default:""`
	Source                 *string  `json:"source" default:""`
	Use_column_name        *string  `json:"use_column_name" default:""`
	Use_indicator_ind      *string  `json:"use_indicator_ind" default:""`
	Use_system_id          *string  `json:"use_system_id" default:""`
	Use_table_name         *string  `json:"use_table_name" default:""`
	Row_changed_by         *string  `json:"row_changed_by" default:""`
	Row_changed_date       *string  `json:"row_changed_date" default:""`
	Row_created_by         *string  `json:"row_created_by" default:""`
	Row_created_date       *string  `json:"row_created_date" default:""`
	Row_effective_date     *string  `json:"row_effective_date" default:""`
	Row_expiry_date        *string  `json:"row_expiry_date" default:""`
	Row_quality            *string  `json:"row_quality" default:""`
}
