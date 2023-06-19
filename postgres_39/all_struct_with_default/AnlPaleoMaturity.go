package dto

type Anl_paleo_maturity struct {
	Analysis_id                string   `json:"analysis_id" default:"source"`
	Anl_source                 string   `json:"anl_source" default:"source"`
	Maturation_obs_no          int      `json:"maturation_obs_no" default:"1"`
	Active_ind                 *string  `json:"active_ind" default:""`
	Effective_date             *string  `json:"effective_date" default:""`
	Expiry_date                *string  `json:"expiry_date" default:""`
	Fluor_color                *string  `json:"fluor_color" default:""`
	Fluor_intensity_desc       *string  `json:"fluor_intensity_desc" default:""`
	Fluor_intensity_value      *float64 `json:"fluor_intensity_value" default:""`
	Fluor_intensity_value_ouom *string  `json:"fluor_intensity_value_ouom" default:""`
	Fluor_intensity_value_uom  *string  `json:"fluor_intensity_value_uom" default:""`
	Maturity_method_type       *string  `json:"maturity_method_type" default:""`
	Ppdm_guid                  *string  `json:"ppdm_guid" default:""`
	Preferred_ind              *string  `json:"preferred_ind" default:""`
	Problem_ind                *string  `json:"problem_ind" default:""`
	Remark                     *string  `json:"remark" default:""`
	Sci_color                  *string  `json:"sci_color" default:""`
	Sci_max                    *float64 `json:"sci_max" default:""`
	Sci_max_ouom               *string  `json:"sci_max_ouom" default:""`
	Sci_max_uom                *string  `json:"sci_max_uom" default:""`
	Sci_min                    *float64 `json:"sci_min" default:""`
	Sci_min_ouom               *string  `json:"sci_min_ouom" default:""`
	Sci_min_uom                *string  `json:"sci_min_uom" default:""`
	Source                     *string  `json:"source" default:""`
	Step_seq_no                *int     `json:"step_seq_no" default:""`
	Substance_id               *string  `json:"substance_id" default:""`
	Tai_color                  *string  `json:"tai_color" default:""`
	Tai_max                    *float64 `json:"tai_max" default:""`
	Tai_max_ouom               *string  `json:"tai_max_ouom" default:""`
	Tai_max_uom                *string  `json:"tai_max_uom" default:""`
	Tai_min                    *float64 `json:"tai_min" default:""`
	Tai_min_ouom               *string  `json:"tai_min_ouom" default:""`
	Tai_min_uom                *string  `json:"tai_min_uom" default:""`
	Row_changed_by             *string  `json:"row_changed_by" default:""`
	Row_changed_date           *string  `json:"row_changed_date" default:""`
	Row_created_by             *string  `json:"row_created_by" default:""`
	Row_created_date           *string  `json:"row_created_date" default:""`
	Row_effective_date         *string  `json:"row_effective_date" default:""`
	Row_expiry_date            *string  `json:"row_expiry_date" default:""`
	Row_quality                *string  `json:"row_quality" default:""`
}