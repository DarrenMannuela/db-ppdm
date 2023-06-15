package dto

type Int_set_partner_cont struct {
	Interest_set_id             string  `json:"interest_set_id" default:"source"`
	Interest_set_seq_no         int     `json:"interest_set_seq_no" default:"1"`
	Partner_ba_id               string  `json:"partner_ba_id" default:"source"`
	Partner_obs_no              int     `json:"partner_obs_no" default:"1"`
	Contact_ba_id               string  `json:"contact_ba_id" default:"source"`
	Contact_obs_no              int     `json:"contact_obs_no" default:"1"`
	Contact_id                  string  `json:"contact_id" default:"source"`
	Active_ind                  *string `json:"active_ind" default:""`
	Address_obs_no              *int    `json:"address_obs_no" default:""`
	Address_source              *string `json:"address_source" default:""`
	Contact_org_ba_id           *string `json:"contact_org_ba_id" default:""`
	Contact_org_ba_seq_no       *int    `json:"contact_org_ba_seq_no" default:""`
	Contact_org_organization_id *string `json:"contact_org_organization_id" default:""`
	Contact_role                *string `json:"contact_role" default:""`
	Contract_id                 *string `json:"contract_id" default:""`
	Effective_date              *string `json:"effective_date" default:""`
	Expiry_date                 *string `json:"expiry_date" default:""`
	Inactive_date               *string `json:"inactive_date" default:""`
	Ppdm_guid                   *string `json:"ppdm_guid" default:""`
	Provision_id                *string `json:"provision_id" default:""`
	Remark                      *string `json:"remark" default:""`
	Source                      *string `json:"source" default:""`
	Row_changed_by              *string `json:"row_changed_by" default:""`
	Row_changed_date            *string `json:"row_changed_date" default:""`
	Row_created_by              *string `json:"row_created_by" default:""`
	Row_created_date            *string `json:"row_created_date" default:""`
	Row_effective_date          *string `json:"row_effective_date" default:""`
	Row_expiry_date             *string `json:"row_expiry_date" default:""`
	Row_quality                 *string `json:"row_quality" default:""`
}
