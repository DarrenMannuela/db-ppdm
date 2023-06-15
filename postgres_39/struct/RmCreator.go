package dto

import (
	"time"
)

type Rm_creator struct {
	Info_item_subtype   string     `json:"info_item_subtype"`
	Information_item_id string     `json:"information_item_id"`
	Creator_id          string     `json:"creator_id"`
	Active_ind          *string    `json:"active_ind"`
	Creator_ba_id       *string    `json:"creator_ba_id"`
	Creator_ba_type     *string    `json:"creator_ba_type"`
	Creator_name        *string    `json:"creator_name"`
	Creator_type        *string    `json:"creator_type"`
	Effective_date      *time.Time `json:"effective_date"`
	Expiry_date         *time.Time `json:"expiry_date"`
	First_name          *string    `json:"first_name"`
	Last_name           *string    `json:"last_name"`
	Middle_name         *string    `json:"middle_name"`
	Ppdm_guid           string     `json:"ppdm_guid"`
	Remark              *string    `json:"remark"`
	Source              *string    `json:"source"`
	Row_changed_by      *string    `json:"row_changed_by"`
	Row_changed_date    *time.Time `json:"row_changed_date"`
	Row_created_by      *string    `json:"row_created_by"`
	Row_created_date    *time.Time `json:"row_created_date"`
	Row_effective_date  *time.Time `json:"row_effective_date"`
	Row_expiry_date     *time.Time `json:"row_expiry_date"`
	Row_quality         *string    `json:"row_quality"`
}
