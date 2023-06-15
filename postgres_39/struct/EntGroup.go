package dto

import (
	"time"
)

type Ent_group struct {
	Entitlement_id     string     `json:"entitlement_id"`
	Security_group_id  string     `json:"security_group_id"`
	Group_obs_no       int        `json:"group_obs_no"`
	Access_condition   *string    `json:"access_condition"`
	Access_type        *string    `json:"access_type"`
	Active_ind         *string    `json:"active_ind"`
	Effective_date     *time.Time `json:"effective_date"`
	Expiry_date        *time.Time `json:"expiry_date"`
	Ppdm_guid          string     `json:"ppdm_guid"`
	Remark             *string    `json:"remark"`
	Restriction_desc   *string    `json:"restriction_desc"`
	Security_desc      *string    `json:"security_desc"`
	Source             *string    `json:"source"`
	Use_desc           *string    `json:"use_desc"`
	Row_changed_by     *string    `json:"row_changed_by"`
	Row_changed_date   *time.Time `json:"row_changed_date"`
	Row_created_by     *string    `json:"row_created_by"`
	Row_created_date   *time.Time `json:"row_created_date"`
	Row_effective_date *time.Time `json:"row_effective_date"`
	Row_expiry_date    *time.Time `json:"row_expiry_date"`
	Row_quality        *string    `json:"row_quality"`
}
