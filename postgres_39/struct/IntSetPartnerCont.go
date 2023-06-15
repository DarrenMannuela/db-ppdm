package dto

import (
	"time"
)

type Int_set_partner_cont struct {
	Interest_set_id             string     `json:"interest_set_id"`
	Interest_set_seq_no         int        `json:"interest_set_seq_no"`
	Partner_ba_id               string     `json:"partner_ba_id"`
	Partner_obs_no              int        `json:"partner_obs_no"`
	Contact_ba_id               string     `json:"contact_ba_id"`
	Contact_obs_no              int        `json:"contact_obs_no"`
	Contact_id                  string     `json:"contact_id"`
	Active_ind                  *string    `json:"active_ind"`
	Address_obs_no              *int       `json:"address_obs_no"`
	Address_source              *string    `json:"address_source"`
	Contact_org_ba_id           *string    `json:"contact_org_ba_id"`
	Contact_org_ba_seq_no       *int       `json:"contact_org_ba_seq_no"`
	Contact_org_organization_id *string    `json:"contact_org_organization_id"`
	Contact_role                *string    `json:"contact_role"`
	Contract_id                 *string    `json:"contract_id"`
	Effective_date              *time.Time `json:"effective_date"`
	Expiry_date                 *time.Time `json:"expiry_date"`
	Inactive_date               *time.Time `json:"inactive_date"`
	Ppdm_guid                   string     `json:"ppdm_guid"`
	Provision_id                *string    `json:"provision_id"`
	Remark                      *string    `json:"remark"`
	Source                      *string    `json:"source"`
	Row_changed_by              *string    `json:"row_changed_by"`
	Row_changed_date            *time.Time `json:"row_changed_date"`
	Row_created_by              *string    `json:"row_created_by"`
	Row_created_date            *time.Time `json:"row_created_date"`
	Row_effective_date          *time.Time `json:"row_effective_date"`
	Row_expiry_date             *time.Time `json:"row_expiry_date"`
	Row_quality                 *string    `json:"row_quality"`
}
