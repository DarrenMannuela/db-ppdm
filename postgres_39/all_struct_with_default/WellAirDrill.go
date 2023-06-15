package dto

type Well_air_drill struct {
	Uwi                         string   `json:"uwi" default:"source"`
	Source                      string   `json:"source" default:"source"`
	Air_drill_obs_no            int      `json:"air_drill_obs_no" default:"1"`
	Active_ind                  *string  `json:"active_ind" default:""`
	Compressor_count            *int     `json:"compressor_count" default:""`
	Contractor                  *string  `json:"contractor" default:""`
	Effective_date              *string  `json:"effective_date" default:""`
	Expiry_date                 *string  `json:"expiry_date" default:""`
	Max_compressor_cap_vol      *float64 `json:"max_compressor_cap_vol" default:""`
	Max_compressor_cap_vol_ouom *string  `json:"max_compressor_cap_vol_ouom" default:""`
	Ppdm_guid                   *string  `json:"ppdm_guid" default:""`
	Remark                      *string  `json:"remark" default:""`
	Row_changed_by              *string  `json:"row_changed_by" default:""`
	Row_changed_date            *string  `json:"row_changed_date" default:""`
	Row_created_by              *string  `json:"row_created_by" default:""`
	Row_created_date            *string  `json:"row_created_date" default:""`
	Row_effective_date          *string  `json:"row_effective_date" default:""`
	Row_expiry_date             *string  `json:"row_expiry_date" default:""`
	Row_quality                 *string  `json:"row_quality" default:""`
}
