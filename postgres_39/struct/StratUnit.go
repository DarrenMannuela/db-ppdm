package dto

import (
	"time"
)

type Strat_unit struct {
	Strat_name_set_id      string     `json:"strat_name_set_id"`
	Strat_unit_id          string     `json:"strat_unit_id"`
	Abbreviation           *string    `json:"abbreviation"`
	Active_ind             *string    `json:"active_ind"`
	Area_id                *string    `json:"area_id"`
	Area_type              *string    `json:"area_type"`
	Business_associate_id  *string    `json:"business_associate_id"`
	Confidence_id          *string    `json:"confidence_id"`
	Current_status_date    *time.Time `json:"current_status_date"`
	Description            *string    `json:"description"`
	Effective_date         *time.Time `json:"effective_date"`
	Expiry_date            *time.Time `json:"expiry_date"`
	Fault_type             *string    `json:"fault_type"`
	Form_code              *string    `json:"form_code"`
	Group_code             *string    `json:"group_code"`
	Long_name              *string    `json:"long_name"`
	Ordinal_age_code       *float64   `json:"ordinal_age_code"`
	Ppdm_guid              string     `json:"ppdm_guid"`
	Preferred_ind          *string    `json:"preferred_ind"`
	Remark                 *string    `json:"remark"`
	Short_name             *string    `json:"short_name"`
	Source                 *string    `json:"source"`
	Strat_interpret_method *string    `json:"strat_interpret_method"`
	Strat_status           *string    `json:"strat_status"`
	Strat_type             *string    `json:"strat_type"`
	Strat_unit_type        *string    `json:"strat_unit_type"`
	Row_changed_by         *string    `json:"row_changed_by"`
	Row_changed_date       *time.Time `json:"row_changed_date"`
	Row_created_by         *string    `json:"row_created_by"`
	Row_created_date       *time.Time `json:"row_created_date"`
	Row_effective_date     *time.Time `json:"row_effective_date"`
	Row_expiry_date        *time.Time `json:"row_expiry_date"`
	Row_quality            *string    `json:"row_quality"`
}
