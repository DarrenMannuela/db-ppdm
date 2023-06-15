package dto

import (
	"time"
)

type Cs_coord_transform struct {
	Transform_id           string     `json:"transform_id"`
	Active_ind             *string    `json:"active_ind"`
	Effective_date         *time.Time `json:"effective_date"`
	Expiry_date            *time.Time `json:"expiry_date"`
	Ppdm_guid              string     `json:"ppdm_guid"`
	Preferred_ind          *string    `json:"preferred_ind"`
	Remark                 *string    `json:"remark"`
	Source                 *string    `json:"source"`
	Source_coord_system_id *string    `json:"source_coord_system_id"`
	Source_document_id     *string    `json:"source_document_id"`
	Target_coord_system_id *string    `json:"target_coord_system_id"`
	Transform_name         *string    `json:"transform_name"`
	Transform_type         *string    `json:"transform_type"`
	Row_changed_by         *string    `json:"row_changed_by"`
	Row_changed_date       *time.Time `json:"row_changed_date"`
	Row_created_by         *string    `json:"row_created_by"`
	Row_created_date       *time.Time `json:"row_created_date"`
	Row_effective_date     *time.Time `json:"row_effective_date"`
	Row_expiry_date        *time.Time `json:"row_expiry_date"`
	Row_quality            *string    `json:"row_quality"`
}
