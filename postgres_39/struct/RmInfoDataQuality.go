package dto

import (
	"time"
)

type Rm_info_data_quality struct {
	Info_item_subtype        string     `json:"info_item_subtype"`
	Information_item_id      string     `json:"information_item_id"`
	Quality_id               string     `json:"quality_id"`
	Active_ind               *string    `json:"active_ind"`
	Attribute_accuracy       *string    `json:"attribute_accuracy"`
	Attribute_accuracy_desc  *string    `json:"attribute_accuracy_desc"`
	Completeness_desc        *string    `json:"completeness_desc"`
	Corrected_date           *time.Time `json:"corrected_date"`
	Deficiency_ind           *string    `json:"deficiency_ind"`
	Deficiency_type          *string    `json:"deficiency_type"`
	Described_by_ba_id       *string    `json:"described_by_ba_id"`
	Description              *string    `json:"description"`
	Effective_date           *time.Time `json:"effective_date"`
	Expiry_date              *time.Time `json:"expiry_date"`
	Information_quality      *string    `json:"information_quality"`
	Information_quality_desc *string    `json:"information_quality_desc"`
	Ppdm_guid                string     `json:"ppdm_guid"`
	Quality_date             *time.Time `json:"quality_date"`
	Remark                   *string    `json:"remark"`
	Source                   *string    `json:"source"`
	Row_changed_by           *string    `json:"row_changed_by"`
	Row_changed_date         *time.Time `json:"row_changed_date"`
	Row_created_by           *string    `json:"row_created_by"`
	Row_created_date         *time.Time `json:"row_created_date"`
	Row_effective_date       *time.Time `json:"row_effective_date"`
	Row_expiry_date          *time.Time `json:"row_expiry_date"`
	Row_quality              *string    `json:"row_quality"`
}
