package dto

type Cat_additive struct {
	Catalogue_additive_id  string   `json:"catalogue_additive_id" default:"source"`
	Active_ind             *string  `json:"active_ind" default:""`
	Additive_group         *string  `json:"additive_group" default:""`
	Additive_name          *string  `json:"additive_name" default:""`
	Additive_type          *string  `json:"additive_type" default:""`
	Effective_date         *string  `json:"effective_date" default:""`
	Expiry_date            *string  `json:"expiry_date" default:""`
	Manufacturer           *string  `json:"manufacturer" default:""`
	Ppdm_guid              *string  `json:"ppdm_guid" default:""`
	Purchase_quantity      *float64 `json:"purchase_quantity" default:""`
	Purchase_quantity_type *string  `json:"purchase_quantity_type" default:""`
	Purchase_quantity_uom  *string  `json:"purchase_quantity_uom" default:""`
	Remark                 *string  `json:"remark" default:""`
	Source                 *string  `json:"source" default:""`
	Supplier               *string  `json:"supplier" default:""`
	Upc_code               *string  `json:"upc_code" default:""`
	Row_changed_by         *string  `json:"row_changed_by" default:""`
	Row_changed_date       *string  `json:"row_changed_date" default:""`
	Row_created_by         *string  `json:"row_created_by" default:""`
	Row_created_date       *string  `json:"row_created_date" default:""`
	Row_effective_date     *string  `json:"row_effective_date" default:""`
	Row_expiry_date        *string  `json:"row_expiry_date" default:""`
	Row_quality            *string  `json:"row_quality" default:""`
}
