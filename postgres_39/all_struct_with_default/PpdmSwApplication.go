package dto

type Ppdm_sw_application struct {
	Sw_application_id       string  `json:"sw_application_id" default:"source"`
	Abbreviation            *string `json:"abbreviation" default:""`
	Active_ind              *string `json:"active_ind" default:""`
	Application_version     *string `json:"application_version" default:""`
	Effective_date          *string `json:"effective_date" default:""`
	Expiry_date             *string `json:"expiry_date" default:""`
	Function_description    *string `json:"function_description" default:""`
	Long_name               *string `json:"long_name" default:""`
	Owner_ba_id             *string `json:"owner_ba_id" default:""`
	Ppdm_guid               *string `json:"ppdm_guid" default:""`
	Preferred_uom_system_id *string `json:"preferred_uom_system_id" default:""`
	Remark                  *string `json:"remark" default:""`
	Short_name              *string `json:"short_name" default:""`
	Source                  *string `json:"source" default:""`
	Row_changed_by          *string `json:"row_changed_by" default:""`
	Row_changed_date        *string `json:"row_changed_date" default:""`
	Row_created_by          *string `json:"row_created_by" default:""`
	Row_created_date        *string `json:"row_created_date" default:""`
	Row_effective_date      *string `json:"row_effective_date" default:""`
	Row_expiry_date         *string `json:"row_expiry_date" default:""`
	Row_quality             *string `json:"row_quality" default:""`
}
