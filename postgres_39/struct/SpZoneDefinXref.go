package dto

import (
	"time"
)

type Sp_zone_defin_xref struct {
	Zone_definition_id_1 string     `json:"zone_definition_id_1"`
	Zone_definition_id_2 string     `json:"zone_definition_id_2"`
	Active_ind           *string    `json:"active_ind"`
	Description          *string    `json:"description"`
	Effective_date       *time.Time `json:"effective_date"`
	Expiry_date          *time.Time `json:"expiry_date"`
	Ppdm_guid            string     `json:"ppdm_guid"`
	Remark               *string    `json:"remark"`
	Responsible_ba_id    *string    `json:"responsible_ba_id"`
	Source               *string    `json:"source"`
	Xref_reason          *string    `json:"xref_reason"`
	Row_changed_by       *string    `json:"row_changed_by"`
	Row_changed_date     *time.Time `json:"row_changed_date"`
	Row_created_by       *string    `json:"row_created_by"`
	Row_created_date     *time.Time `json:"row_created_date"`
	Row_effective_date   *time.Time `json:"row_effective_date"`
	Row_expiry_date      *time.Time `json:"row_expiry_date"`
	Row_quality          *string    `json:"row_quality"`
}
