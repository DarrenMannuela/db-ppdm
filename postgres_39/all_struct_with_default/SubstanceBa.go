package dto

type Substance_ba struct {
	Substance_id           string   `json:"substance_id" default:"source"`
	Anl_source             string   `json:"anl_source" default:"source"`
	Ba_obs_no              int      `json:"ba_obs_no" default:"1"`
	Active_ind             *string  `json:"active_ind" default:""`
	Ba_role_type           *string  `json:"ba_role_type" default:""`
	Cas_number             *string  `json:"cas_number" default:""`
	Effective_date         *string  `json:"effective_date" default:""`
	Expiry_date            *string  `json:"expiry_date" default:""`
	Location               *string  `json:"location" default:""`
	Manufacturer           *string  `json:"manufacturer" default:""`
	Ppdm_guid              *string  `json:"ppdm_guid" default:""`
	Price                  *float64 `json:"price" default:""`
	Price_ouom             *string  `json:"price_ouom" default:""`
	Problem_ind            *string  `json:"problem_ind" default:""`
	Purchase_quantity      *float64 `json:"purchase_quantity" default:""`
	Purchase_quantity_ouom *string  `json:"purchase_quantity_ouom" default:""`
	Purchase_quantity_type *string  `json:"purchase_quantity_type" default:""`
	Remark                 *string  `json:"remark" default:""`
	Source                 *string  `json:"source" default:""`
	Step_seq_no            *int     `json:"step_seq_no" default:""`
	Substance_ba_id        *string  `json:"substance_ba_id" default:""`
	Supplier               *string  `json:"supplier" default:""`
	Row_changed_by         *string  `json:"row_changed_by" default:""`
	Row_changed_date       *string  `json:"row_changed_date" default:""`
	Row_created_by         *string  `json:"row_created_by" default:""`
	Row_created_date       *string  `json:"row_created_date" default:""`
	Row_effective_date     *string  `json:"row_effective_date" default:""`
	Row_expiry_date        *string  `json:"row_expiry_date" default:""`
	Row_quality            *string  `json:"row_quality" default:""`
}
