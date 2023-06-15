package dto

import (
	"time"
)

type Sp_desc_text struct {
	Spatial_description_id string     `json:"spatial_description_id"`
	Spatial_obs_no         int        `json:"spatial_obs_no"`
	Text_id                string     `json:"text_id"`
	Description_seq_no     int        `json:"description_seq_no"`
	Active_ind             *string    `json:"active_ind"`
	Description            *string    `json:"description"`
	Effective_date         *time.Time `json:"effective_date"`
	Expiry_date            *time.Time `json:"expiry_date"`
	Ppdm_guid              string     `json:"ppdm_guid"`
	Remark                 *string    `json:"remark"`
	Source                 *string    `json:"source"`
	Row_changed_by         *string    `json:"row_changed_by"`
	Row_changed_date       *time.Time `json:"row_changed_date"`
	Row_created_by         *string    `json:"row_created_by"`
	Row_created_date       *time.Time `json:"row_created_date"`
	Row_effective_date     *time.Time `json:"row_effective_date"`
	Row_expiry_date        *time.Time `json:"row_expiry_date"`
	Row_quality            *string    `json:"row_quality"`
}
