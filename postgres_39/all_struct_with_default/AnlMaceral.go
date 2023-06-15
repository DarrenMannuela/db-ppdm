package dto

type Anl_maceral struct {
	Analysis_id                 string   `json:"analysis_id" default:"source"`
	Anl_source                  string   `json:"anl_source" default:"source"`
	Maceral_anl_obs_no          int      `json:"maceral_anl_obs_no" default:"1"`
	Active_ind                  *string  `json:"active_ind" default:""`
	Effective_date              *string  `json:"effective_date" default:""`
	Expiry_date                 *string  `json:"expiry_date" default:""`
	Lithology_desc              *string  `json:"lithology_desc" default:""`
	Maceral_amount_type         *string  `json:"maceral_amount_type" default:""`
	Ppdm_guid                   *string  `json:"ppdm_guid" default:""`
	Preferred_ind               *string  `json:"preferred_ind" default:""`
	Problem_ind                 *string  `json:"problem_ind" default:""`
	Remark                      *string  `json:"remark" default:""`
	Sample_tot_maceral_val      *float64 `json:"sample_tot_maceral_val" default:""`
	Sample_tot_maceral_val_ouom *string  `json:"sample_tot_maceral_val_ouom" default:""`
	Source                      *string  `json:"source" default:""`
	Step_seq_no                 *int     `json:"step_seq_no" default:""`
	Substance_id                *string  `json:"substance_id" default:""`
	Total_maceral_val           *float64 `json:"total_maceral_val" default:""`
	Total_maceral_val_ouom      *string  `json:"total_maceral_val_ouom" default:""`
	Row_changed_by              *string  `json:"row_changed_by" default:""`
	Row_changed_date            *string  `json:"row_changed_date" default:""`
	Row_created_by              *string  `json:"row_created_by" default:""`
	Row_created_date            *string  `json:"row_created_date" default:""`
	Row_effective_date          *string  `json:"row_effective_date" default:""`
	Row_expiry_date             *string  `json:"row_expiry_date" default:""`
	Row_quality                 *string  `json:"row_quality" default:""`
}
