package dto

type Sp_desc_xref struct {
	Spatial_description_id   string   `json:"spatial_description_id" default:"source"`
	Spatial_obs_no           int      `json:"spatial_obs_no" default:"1"`
	Spatial_description_id_2 string   `json:"spatial_description_id_2" default:"source"`
	Spatial_obs_no_2         int      `json:"spatial_obs_no_2" default:"1"`
	Active_ind               *string  `json:"active_ind" default:""`
	Effective_date           *string  `json:"effective_date" default:""`
	Expiry_date              *string  `json:"expiry_date" default:""`
	Overlap_size             *float64 `json:"overlap_size" default:""`
	Overlap_size_ouom        *string  `json:"overlap_size_ouom" default:""`
	Ppdm_guid                *string  `json:"ppdm_guid" default:""`
	Remark                   *string  `json:"remark" default:""`
	Source                   *string  `json:"source" default:""`
	Spatial_xref_type        *string  `json:"spatial_xref_type" default:""`
	Row_changed_by           *string  `json:"row_changed_by" default:""`
	Row_changed_date         *string  `json:"row_changed_date" default:""`
	Row_created_by           *string  `json:"row_created_by" default:""`
	Row_created_date         *string  `json:"row_created_date" default:""`
	Row_effective_date       *string  `json:"row_effective_date" default:""`
	Row_expiry_date          *string  `json:"row_expiry_date" default:""`
	Row_quality              *string  `json:"row_quality" default:""`
}
