package dto

import (
	"time"
)

type Land_right_well struct {
	Uwi                    string     `json:"uwi"`
	Land_right_subtype     string     `json:"land_right_subtype"`
	Land_right_id          string     `json:"land_right_id"`
	Lr_well_seq_no         int        `json:"lr_well_seq_no"`
	Active_ind             *string    `json:"active_ind"`
	Completion_obs_no      *int       `json:"completion_obs_no"`
	Effective_date         *time.Time `json:"effective_date"`
	Expiry_date            *time.Time `json:"expiry_date"`
	Gas_percent_psu        *float64   `json:"gas_percent_psu"`
	Key_well_ind           *string    `json:"key_well_ind"`
	Oil_percent_psu        *float64   `json:"oil_percent_psu"`
	Ppdm_guid              string     `json:"ppdm_guid"`
	Pr_str_form_obs_no     *int       `json:"pr_str_form_obs_no"`
	Pr_str_source          *string    `json:"pr_str_source"`
	Remark                 *string    `json:"remark"`
	Source                 *string    `json:"source"`
	Spacing_complete_ind   *string    `json:"spacing_complete_ind"`
	Spacing_unit_id        *string    `json:"spacing_unit_id"`
	String_id              *string    `json:"string_id"`
	Well_in_tract_ind      *string    `json:"well_in_tract_ind"`
	Well_relationship_type *string    `json:"well_relationship_type"`
	Row_changed_by         *string    `json:"row_changed_by"`
	Row_changed_date       *time.Time `json:"row_changed_date"`
	Row_created_by         *string    `json:"row_created_by"`
	Row_created_date       *time.Time `json:"row_created_date"`
	Row_effective_date     *time.Time `json:"row_effective_date"`
	Row_expiry_date        *time.Time `json:"row_expiry_date"`
	Row_quality            *string    `json:"row_quality"`
}
