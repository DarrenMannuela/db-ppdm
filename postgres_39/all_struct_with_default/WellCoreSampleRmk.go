package dto

type Well_core_sample_rmk struct {
	Uwi                    string  `json:"uwi" default:"source"`
	Source                 string  `json:"source" default:"source"`
	Core_id                string  `json:"core_id" default:"source"`
	Analysis_obs_no        int     `json:"analysis_obs_no" default:"1"`
	Sample_num             string  `json:"sample_num" default:"source"`
	Sample_analysis_obs_no int     `json:"sample_analysis_obs_no" default:"1"`
	Remark_obs_no          int     `json:"remark_obs_no" default:"1"`
	Active_ind             *string `json:"active_ind" default:""`
	Effective_date         *string `json:"effective_date" default:""`
	Expiry_date            *string `json:"expiry_date" default:""`
	Ppdm_guid              *string `json:"ppdm_guid" default:""`
	Remark                 *string `json:"remark" default:""`
	Remark_source          *string `json:"remark_source" default:""`
	Row_changed_by         *string `json:"row_changed_by" default:""`
	Row_changed_date       *string `json:"row_changed_date" default:""`
	Row_created_by         *string `json:"row_created_by" default:""`
	Row_created_date       *string `json:"row_created_date" default:""`
	Row_effective_date     *string `json:"row_effective_date" default:""`
	Row_expiry_date        *string `json:"row_expiry_date" default:""`
	Row_quality            *string `json:"row_quality" default:""`
}
