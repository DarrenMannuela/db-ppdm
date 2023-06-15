package dto

import (
	"time"
)

type Resent_revision_cat struct {
	Revision_category_id string     `json:"revision_category_id"`
	Active_ind           *string    `json:"active_ind"`
	Category_name        *string    `json:"category_name"`
	Category_type        *string    `json:"category_type"`
	Effective_date       *time.Time `json:"effective_date"`
	Expiry_date          *time.Time `json:"expiry_date"`
	Gross_ind            *string    `json:"gross_ind"`
	Net_ind              *string    `json:"net_ind"`
	Part_of_category_id  *string    `json:"part_of_category_id"`
	Ppdm_guid            string     `json:"ppdm_guid"`
	Remark               *string    `json:"remark"`
	Source               *string    `json:"source"`
	Row_changed_by       *string    `json:"row_changed_by"`
	Row_changed_date     *time.Time `json:"row_changed_date"`
	Row_created_by       *string    `json:"row_created_by"`
	Row_created_date     *time.Time `json:"row_created_date"`
	Row_effective_date   *time.Time `json:"row_effective_date"`
	Row_expiry_date      *time.Time `json:"row_expiry_date"`
	Row_quality          *string    `json:"row_quality"`
}
