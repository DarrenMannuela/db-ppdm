package dto

import (
	"time"
)

type Cs_coord_trans_value struct {
	Transform_id       string     `json:"transform_id"`
	Transform_type     string     `json:"transform_type"`
	Parameter_id       string     `json:"parameter_id"`
	Value_id           string     `json:"value_id"`
	Active_ind         *string    `json:"active_ind"`
	Effective_date     *time.Time `json:"effective_date"`
	Expiry_date        *time.Time `json:"expiry_date"`
	Ppdm_guid          string     `json:"ppdm_guid"`
	Remark             *string    `json:"remark"`
	Source             *string    `json:"source"`
	Source_document_id *string    `json:"source_document_id"`
	Text_value         *string    `json:"text_value"`
	Trans_value        *float64   `json:"trans_value"`
	Trans_value_uom    *string    `json:"trans_value_uom"`
	Row_changed_by     *string    `json:"row_changed_by"`
	Row_changed_date   *time.Time `json:"row_changed_date"`
	Row_created_by     *string    `json:"row_created_by"`
	Row_created_date   *time.Time `json:"row_created_date"`
	Row_effective_date *time.Time `json:"row_effective_date"`
	Row_expiry_date    *time.Time `json:"row_expiry_date"`
	Row_quality        *string    `json:"row_quality"`
}
