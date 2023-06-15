package dto

type R_color_palette struct {
	Color_palette      string  `json:"color_palette" default:"source"`
	Abbreviation       *string `json:"abbreviation" default:""`
	Active_ind         *string `json:"active_ind" default:""`
	Effective_date     *string `json:"effective_date" default:""`
	Expiry_date        *string `json:"expiry_date" default:""`
	Long_name          *string `json:"long_name" default:""`
	Palette_purpose    *string `json:"palette_purpose" default:""`
	Palette_type       *string `json:"palette_type" default:""`
	Physical_item_id   *string `json:"physical_item_id" default:""`
	Ppdm_guid          *string `json:"ppdm_guid" default:""`
	Remark             *string `json:"remark" default:""`
	Short_name         *string `json:"short_name" default:""`
	Source             *string `json:"source" default:""`
	Row_changed_by     *string `json:"row_changed_by" default:""`
	Row_changed_date   *string `json:"row_changed_date" default:""`
	Row_created_by     *string `json:"row_created_by" default:""`
	Row_created_date   *string `json:"row_created_date" default:""`
	Row_effective_date *string `json:"row_effective_date" default:""`
	Row_expiry_date    *string `json:"row_expiry_date" default:""`
	Row_quality        *string `json:"row_quality" default:""`
}
