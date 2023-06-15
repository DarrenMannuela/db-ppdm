package dto

import (
	"time"
)

type Well_core_sample_desc struct {
	Uwi                    string     `json:"uwi"`
	Source                 string     `json:"source"`
	Core_id                string     `json:"core_id"`
	Analysis_obs_no        int        `json:"analysis_obs_no"`
	Sample_num             string     `json:"sample_num"`
	Sample_analysis_obs_no int        `json:"sample_analysis_obs_no"`
	Sample_desc_obs_no     int        `json:"sample_desc_obs_no"`
	Active_ind             *string    `json:"active_ind"`
	Base_depth             *float64   `json:"base_depth"`
	Base_depth_ouom        *string    `json:"base_depth_ouom"`
	Description            *string    `json:"description"`
	Desc_source            *string    `json:"desc_source"`
	Dip_angle              *float64   `json:"dip_angle"`
	Effective_date         *time.Time `json:"effective_date"`
	Expiry_date            *time.Time `json:"expiry_date"`
	Lithology_desc         *string    `json:"lithology_desc"`
	Porosity_length        *float64   `json:"porosity_length"`
	Porosity_length_ouom   *string    `json:"porosity_length_ouom"`
	Porosity_quality       *string    `json:"porosity_quality"`
	Porosity_type          *string    `json:"porosity_type"`
	Ppdm_guid              string     `json:"ppdm_guid"`
	Recovered_amount       *float64   `json:"recovered_amount"`
	Recovered_amount_ouom  *string    `json:"recovered_amount_ouom"`
	Remark                 *string    `json:"remark"`
	Sample_type            *string    `json:"sample_type"`
	Show_length            *float64   `json:"show_length"`
	Show_length_ouom       *string    `json:"show_length_ouom"`
	Show_quality           *string    `json:"show_quality"`
	Show_type              *string    `json:"show_type"`
	Strat_name_set_id      *string    `json:"strat_name_set_id"`
	Strat_unit_id          *string    `json:"strat_unit_id"`
	Swc_recovery_type      *string    `json:"swc_recovery_type"`
	Top_depth              *float64   `json:"top_depth"`
	Top_depth_ouom         *string    `json:"top_depth_ouom"`
	Row_changed_by         *string    `json:"row_changed_by"`
	Row_changed_date       *time.Time `json:"row_changed_date"`
	Row_created_by         *string    `json:"row_created_by"`
	Row_created_date       *time.Time `json:"row_created_date"`
	Row_effective_date     *time.Time `json:"row_effective_date"`
	Row_expiry_date        *time.Time `json:"row_expiry_date"`
	Row_quality            *string    `json:"row_quality"`
}
