package dto

import (
	"time"
)

type Seis_group_comp struct {
	Seis_group_set_subtype string     `json:"seis_group_set_subtype"`
	Seis_group_id          string     `json:"seis_group_id"`
	Input_seis_set_subtype string     `json:"input_seis_set_subtype"`
	Input_seis_set_id      string     `json:"input_seis_set_id"`
	Component_id           string     `json:"component_id"`
	Active_ind             *string    `json:"active_ind"`
	Corner1_lat            *float64   `json:"corner_1_lat"`
	Corner1_long           *float64   `json:"corner_1_long"`
	Corner2_lat            *float64   `json:"corner_2_lat"`
	Corner2_long           *float64   `json:"corner_2_long"`
	Corner3_lat            *float64   `json:"corner_3_lat"`
	Corner3_long           *float64   `json:"corner_3_long"`
	Description            *string    `json:"description"`
	Effective_date         *time.Time `json:"effective_date"`
	Exclusion_ind          *string    `json:"exclusion_ind"`
	Expiry_date            *time.Time `json:"expiry_date"`
	Inclusion_ind          *string    `json:"inclusion_ind"`
	Ppdm_guid              string     `json:"ppdm_guid"`
	Remark                 *string    `json:"remark"`
	Seis_group_type        *string    `json:"seis_group_type"`
	Source                 *string    `json:"source"`
	Row_changed_by         *string    `json:"row_changed_by"`
	Row_changed_date       *time.Time `json:"row_changed_date"`
	Row_created_by         *string    `json:"row_created_by"`
	Row_created_date       *time.Time `json:"row_created_date"`
	Row_effective_date     *time.Time `json:"row_effective_date"`
	Row_expiry_date        *time.Time `json:"row_expiry_date"`
	Row_quality            *string    `json:"row_quality"`
}
