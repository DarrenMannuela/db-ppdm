package dto

type Land_sale struct {
	Jurisdiction       string   `json:"jurisdiction" default:"source"`
	Land_sale_number   string   `json:"land_sale_number" default:"source"`
	Active_ind         *string  `json:"active_ind" default:""`
	Close_date         *string  `json:"close_date" default:""`
	Description        *string  `json:"description" default:""`
	Effective_date     *string  `json:"effective_date" default:""`
	Expiry_date        *string  `json:"expiry_date" default:""`
	Land_sale_name     *string  `json:"land_sale_name" default:""`
	Open_date          *string  `json:"open_date" default:""`
	Ppdm_guid          *string  `json:"ppdm_guid" default:""`
	Publish_date       *string  `json:"publish_date" default:""`
	Remark             *string  `json:"remark" default:""`
	Sale_date          *string  `json:"sale_date" default:""`
	Sold_size          *float64 `json:"sold_size" default:""`
	Sold_size_ouom     *string  `json:"sold_size_ouom" default:""`
	Source             *string  `json:"source" default:""`
	Total_bonus        *float64 `json:"total_bonus" default:""`
	Total_bonus_ouom   *string  `json:"total_bonus_ouom" default:""`
	Total_size         *float64 `json:"total_size" default:""`
	Total_size_ouom    *string  `json:"total_size_ouom" default:""`
	Row_changed_by     *string  `json:"row_changed_by" default:""`
	Row_changed_date   *string  `json:"row_changed_date" default:""`
	Row_created_by     *string  `json:"row_created_by" default:""`
	Row_created_date   *string  `json:"row_created_date" default:""`
	Row_effective_date *string  `json:"row_effective_date" default:""`
	Row_expiry_date    *string  `json:"row_expiry_date" default:""`
	Row_quality        *string  `json:"row_quality" default:""`
}
