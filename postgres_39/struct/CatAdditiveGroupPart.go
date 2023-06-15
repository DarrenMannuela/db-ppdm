package dto

import (
	"time"
)

type Cat_additive_group_part struct {
	Additive_group_id     string     `json:"additive_group_id"`
	Additive_type_id      string     `json:"additive_type_id"`
	Additive_part_obs_no  int        `json:"additive_part_obs_no"`
	Active_ind            *string    `json:"active_ind"`
	Catalogue_additive_id *string    `json:"catalogue_additive_id"`
	Effective_date        *time.Time `json:"effective_date"`
	Expiry_date           *time.Time `json:"expiry_date"`
	Long_name             *string    `json:"long_name"`
	Ppdm_guid             string     `json:"ppdm_guid"`
	Remark                *string    `json:"remark"`
	Short_name            *string    `json:"short_name"`
	Source                *string    `json:"source"`
	Source_document_id    *string    `json:"source_document_id"`
	Substance_id          *string    `json:"substance_id"`
	Row_changed_by        *string    `json:"row_changed_by"`
	Row_changed_date      *time.Time `json:"row_changed_date"`
	Row_created_by        *string    `json:"row_created_by"`
	Row_created_date      *time.Time `json:"row_created_date"`
	Row_effective_date    *time.Time `json:"row_effective_date"`
	Row_expiry_date       *time.Time `json:"row_expiry_date"`
	Row_quality           *string    `json:"row_quality"`
}
