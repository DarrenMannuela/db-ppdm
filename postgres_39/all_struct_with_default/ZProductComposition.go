package dto

type Z_product_composition struct {
	Conversion_id          string  `json:"conversion_id" default:"source"`
	Active_ind             *string `json:"active_ind" default:""`
	Defined_by_ba_id       *string `json:"defined_by_ba_id" default:""`
	Effective_date         *string `json:"effective_date" default:""`
	Expiry_date            *string `json:"expiry_date" default:""`
	Formula                *string `json:"formula" default:""`
	Ppdm_guid              *string `json:"ppdm_guid" default:""`
	Product_component_type *string `json:"product_component_type" default:""`
	Product_type           *string `json:"product_type" default:""`
	Remark                 *string `json:"remark" default:""`
	Source                 *string `json:"source" default:""`
	Source_document_id     *string `json:"source_document_id" default:""`
	Subproduct_type        *string `json:"subproduct_type" default:""`
	Row_changed_by         *string `json:"row_changed_by" default:""`
	Row_changed_date       *string `json:"row_changed_date" default:""`
	Row_created_by         *string `json:"row_created_by" default:""`
	Row_created_date       *string `json:"row_created_date" default:""`
	Row_effective_date     *string `json:"row_effective_date" default:""`
	Row_expiry_date        *string `json:"row_expiry_date" default:""`
	Row_quality            *string `json:"row_quality" default:""`
}
