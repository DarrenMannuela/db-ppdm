package dto

type Resent_prod_property struct {
	Resent_id            string   `json:"resent_id" default:"source"`
	Reserve_class_id     string   `json:"reserve_class_id" default:"source"`
	Property_seq_no      int      `json:"property_seq_no" default:"1"`
	Active_ind           *string  `json:"active_ind" default:""`
	Effective_date       *string  `json:"effective_date" default:""`
	Expiry_date          *string  `json:"expiry_date" default:""`
	Heat_content         *float64 `json:"heat_content" default:""`
	Heat_content_ouom    *string  `json:"heat_content_ouom" default:""`
	Loss_factor          *float64 `json:"loss_factor" default:""`
	Oil_density          *float64 `json:"oil_density" default:""`
	Oil_density_ouom     *string  `json:"oil_density_ouom" default:""`
	Ppdm_guid            *string  `json:"ppdm_guid" default:""`
	Product_type         *string  `json:"product_type" default:""`
	Remark               *string  `json:"remark" default:""`
	Source               *string  `json:"source" default:""`
	Sulphur_content      *float64 `json:"sulphur_content" default:""`
	Sulphur_content_ouom *string  `json:"sulphur_content_ouom" default:""`
	Row_changed_by       *string  `json:"row_changed_by" default:""`
	Row_changed_date     *string  `json:"row_changed_date" default:""`
	Row_created_by       *string  `json:"row_created_by" default:""`
	Row_created_date     *string  `json:"row_created_date" default:""`
	Row_effective_date   *string  `json:"row_effective_date" default:""`
	Row_expiry_date      *string  `json:"row_expiry_date" default:""`
	Row_quality          *string  `json:"row_quality" default:""`
}
