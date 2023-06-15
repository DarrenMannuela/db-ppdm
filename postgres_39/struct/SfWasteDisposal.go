package dto

import (
	"time"
)

type Sf_waste_disposal struct {
	Sf_subtype             string     `json:"sf_subtype"`
	Waste_facility_id      string     `json:"waste_facility_id"`
	Waste_id               string     `json:"waste_id"`
	Active_ind             *string    `json:"active_ind"`
	Adjust_reason          *string    `json:"adjust_reason"`
	Amount_adjustment      *float64   `json:"amount_adjustment"`
	Amount_adjustment_ouom *string    `json:"amount_adjustment_ouom"`
	Amount_adjustment_uom  *string    `json:"amount_adjustment_uom"`
	Amount_received        *float64   `json:"amount_received"`
	Amount_received_ouom   *string    `json:"amount_received_ouom"`
	Amount_received_uom    *string    `json:"amount_received_uom"`
	Amount_shipped         *float64   `json:"amount_shipped"`
	Amount_shipped_ouom    *string    `json:"amount_shipped_ouom"`
	Amount_shipped_uom     *string    `json:"amount_shipped_uom"`
	Contract_id            *string    `json:"contract_id"`
	Description            *string    `json:"description"`
	Effective_date         *time.Time `json:"effective_date"`
	End_date               *time.Time `json:"end_date"`
	Expiry_date            *time.Time `json:"expiry_date"`
	Operator_ba_id         *string    `json:"operator_ba_id"`
	Origin_facility_id     *string    `json:"origin_facility_id"`
	Origin_facility_type   *string    `json:"origin_facility_type"`
	Origin_hse_incident_id *string    `json:"origin_hse_incident_id"`
	Origin_uwi             *string    `json:"origin_uwi"`
	Ppdm_guid              string     `json:"ppdm_guid"`
	Provision_id           *string    `json:"provision_id"`
	Rate_schedule_id       *string    `json:"rate_schedule_id"`
	Receiver_ba_id         *string    `json:"receiver_ba_id"`
	Remark                 *string    `json:"remark"`
	Reporting_uom          *string    `json:"reporting_uom"`
	Service_quality        *string    `json:"service_quality"`
	Service_type           *string    `json:"service_type"`
	Shipped_date           *time.Time `json:"shipped_date"`
	Source                 *string    `json:"source"`
	Transport_ba_id        *string    `json:"transport_ba_id"`
	Waste_handling_method  *string    `json:"waste_handling_method"`
	Waste_hazard           *string    `json:"waste_hazard"`
	Waste_material         *string    `json:"waste_material"`
	Waste_origin           *string    `json:"waste_origin"`
	Row_changed_by         *string    `json:"row_changed_by"`
	Row_changed_date       *time.Time `json:"row_changed_date"`
	Row_created_by         *string    `json:"row_created_by"`
	Row_created_date       *time.Time `json:"row_created_date"`
	Row_effective_date     *time.Time `json:"row_effective_date"`
	Row_expiry_date        *time.Time `json:"row_expiry_date"`
	Row_quality            *string    `json:"row_quality"`
}
