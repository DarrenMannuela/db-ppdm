package dto

type Pden_decline_segment struct {
	Pden_subtype       string   `json:"pden_subtype" default:"source"`
	Pden_id            string   `json:"pden_id" default:"source"`
	Pden_source        string   `json:"pden_source" default:"source"`
	Case_id            string   `json:"case_id" default:"source"`
	Segment_id         string   `json:"segment_id" default:"source"`
	Active_ind         *string  `json:"active_ind" default:""`
	Decline_curve_type *string  `json:"decline_curve_type" default:""`
	Decline_type       *string  `json:"decline_type" default:""`
	Duration           *float64 `json:"duration" default:""`
	Duration_ouom      *string  `json:"duration_ouom" default:""`
	Effective_date     *string  `json:"effective_date" default:""`
	End_date           *string  `json:"end_date" default:""`
	Expiry_date        *string  `json:"expiry_date" default:""`
	Final_decline      *float64 `json:"final_decline" default:""`
	Final_rate         *float64 `json:"final_rate" default:""`
	Initial_decline    *float64 `json:"initial_decline" default:""`
	Initial_rate       *float64 `json:"initial_rate" default:""`
	Minimum_decline    *float64 `json:"minimum_decline" default:""`
	N_factor           *float64 `json:"n_factor" default:""`
	Ppdm_guid          *string  `json:"ppdm_guid" default:""`
	Product_type       *string  `json:"product_type" default:""`
	Rate_ouom          *string  `json:"rate_ouom" default:""`
	Ratio_curve_type   *string  `json:"ratio_curve_type" default:""`
	Ratio_final_rate   *float64 `json:"ratio_final_rate" default:""`
	Ratio_fluid_type   *string  `json:"ratio_fluid_type" default:""`
	Ratio_initial_rate *float64 `json:"ratio_initial_rate" default:""`
	Ratio_rate_ouom    *string  `json:"ratio_rate_ouom" default:""`
	Ratio_volume       *float64 `json:"ratio_volume" default:""`
	Ratio_volume_ouom  *string  `json:"ratio_volume_ouom" default:""`
	Remark             *string  `json:"remark" default:""`
	Source             *string  `json:"source" default:""`
	Start_date         *string  `json:"start_date" default:""`
	Volume             *float64 `json:"volume" default:""`
	Volume_ouom        *string  `json:"volume_ouom" default:""`
	Row_changed_by     *string  `json:"row_changed_by" default:""`
	Row_changed_date   *string  `json:"row_changed_date" default:""`
	Row_created_by     *string  `json:"row_created_by" default:""`
	Row_created_date   *string  `json:"row_created_date" default:""`
	Row_effective_date *string  `json:"row_effective_date" default:""`
	Row_expiry_date    *string  `json:"row_expiry_date" default:""`
	Row_quality        *string  `json:"row_quality" default:""`
}
