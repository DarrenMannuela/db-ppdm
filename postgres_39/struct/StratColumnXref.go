package dto

import (
	"time"
)

type Strat_column_xref struct {
	Strat_column_id_1      string     `json:"strat_column_id_1"`
	Strat_column_source_1  string     `json:"strat_column_source_1"`
	Strat_column_id_2      string     `json:"strat_column_id_2"`
	Strat_column_source_2  string     `json:"strat_column_source_2"`
	Active_ind             *string    `json:"active_ind"`
	Effective_date         *time.Time `json:"effective_date"`
	Expiry_date            *time.Time `json:"expiry_date"`
	Ppdm_guid              string     `json:"ppdm_guid"`
	Remark                 *string    `json:"remark"`
	Source                 *string    `json:"source"`
	Strat_column_xref_type *string    `json:"strat_column_xref_type"`
	Row_changed_by         *string    `json:"row_changed_by"`
	Row_changed_date       *time.Time `json:"row_changed_date"`
	Row_created_by         *string    `json:"row_created_by"`
	Row_created_date       *time.Time `json:"row_created_date"`
	Row_effective_date     *time.Time `json:"row_effective_date"`
	Row_expiry_date        *time.Time `json:"row_expiry_date"`
	Row_quality            *string    `json:"row_quality"`
}
