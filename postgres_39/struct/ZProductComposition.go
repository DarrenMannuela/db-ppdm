package dto

import (
	"time"
)

type Z_product_composition struct {
	Conversion_id          string     `json:"conversion_id"`
	Active_ind             *string    `json:"active_ind"`
	Defined_by_ba_id       *string    `json:"defined_by_ba_id"`
	Effective_date         *time.Time `json:"effective_date"`
	Expiry_date            *time.Time `json:"expiry_date"`
	Formula                *string    `json:"formula"`
	Ppdm_guid              string     `json:"ppdm_guid"`
	Product_component_type *string    `json:"product_component_type"`
	Product_type           *string    `json:"product_type"`
	Remark                 *string    `json:"remark"`
	Source                 *string    `json:"source"`
	Source_document_id     *string    `json:"source_document_id"`
	Subproduct_type        *string    `json:"subproduct_type"`
	Row_changed_by         *string    `json:"row_changed_by"`
	Row_changed_date       *time.Time `json:"row_changed_date"`
	Row_created_by         *string    `json:"row_created_by"`
	Row_created_date       *time.Time `json:"row_created_date"`
	Row_effective_date     *time.Time `json:"row_effective_date"`
	Row_expiry_date        *time.Time `json:"row_expiry_date"`
	Row_quality            *string    `json:"row_quality"`
}
