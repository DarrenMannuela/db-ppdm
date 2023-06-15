package dto

import (
	"time"
)

type Well_mud_property struct {
	Sample_analysis_id     string     `json:"sample_analysis_id"`
	Sample_analysis_source string     `json:"sample_analysis_source"`
	Property_obs_no        int        `json:"property_obs_no"`
	Active_ind             *string    `json:"active_ind"`
	Average_value          *float64   `json:"average_value"`
	Average_value_ouom     *string    `json:"average_value_ouom"`
	Average_value_uom      *string    `json:"average_value_uom"`
	Date_format_desc       *time.Time `json:"date_format_desc"`
	Effective_date         *time.Time `json:"effective_date"`
	Expiry_date            *time.Time `json:"expiry_date"`
	Max_value              *float64   `json:"max_value"`
	Max_value_ouom         *string    `json:"max_value_ouom"`
	Max_value_uom          *string    `json:"max_value_uom"`
	Min_value              *float64   `json:"min_value"`
	Min_value_ouom         *string    `json:"min_value_ouom"`
	Min_value_uom          *string    `json:"min_value_uom"`
	Mud_anl_obs_no         *int       `json:"mud_anl_obs_no"`
	Mud_property           *string    `json:"mud_property"`
	Mud_property_code      *string    `json:"mud_property_code"`
	Mud_property_desc      *string    `json:"mud_property_desc"`
	Mud_property_ref       *string    `json:"mud_property_ref"`
	Ppdm_guid              string     `json:"ppdm_guid"`
	Reference_value        *float64   `json:"reference_value"`
	Reference_value_ouom   *string    `json:"reference_value_ouom"`
	Reference_value_uom    *string    `json:"reference_value_uom"`
	Remark                 *string    `json:"remark"`
	Sample_type            *string    `json:"sample_type"`
	Source                 *string    `json:"source"`
	Step_seq_no            *int       `json:"step_seq_no"`
	Uwi                    *string    `json:"uwi"`
	Row_changed_by         *string    `json:"row_changed_by"`
	Row_changed_date       *time.Time `json:"row_changed_date"`
	Row_created_by         *string    `json:"row_created_by"`
	Row_created_date       *time.Time `json:"row_created_date"`
	Row_effective_date     *time.Time `json:"row_effective_date"`
	Row_expiry_date        *time.Time `json:"row_expiry_date"`
	Row_quality            *string    `json:"row_quality"`
}
