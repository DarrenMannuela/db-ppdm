package dto

import (
	"time"
)

type Well_test_mud struct {
	Uwi                       string     `json:"uwi"`
	Source                    string     `json:"source"`
	Test_type                 string     `json:"test_type"`
	Run_num                   string     `json:"run_num"`
	Test_num                  string     `json:"test_num"`
	Mud_type                  string     `json:"mud_type"`
	Active_ind                *string    `json:"active_ind"`
	Effective_date            *time.Time `json:"effective_date"`
	Expiry_date               *time.Time `json:"expiry_date"`
	Filtrate_resistivity      *float64   `json:"filtrate_resistivity"`
	Filtrate_resistivity_ouom *string    `json:"filtrate_resistivity_ouom"`
	Filtrate_salinity         *float64   `json:"filtrate_salinity"`
	Filtrate_salinity_ouom    *string    `json:"filtrate_salinity_ouom"`
	Filtrate_salinity_uom     *string    `json:"filtrate_salinity_uom"`
	Filtrate_temperature      *float64   `json:"filtrate_temperature"`
	Filtrate_temperature_ouom *string    `json:"filtrate_temperature_ouom"`
	Mud_ph                    *float64   `json:"mud_ph"`
	Mud_resistivity           *float64   `json:"mud_resistivity"`
	Mud_resistivity_ouom      *string    `json:"mud_resistivity_ouom"`
	Mud_salinity              *float64   `json:"mud_salinity"`
	Mud_salinity_ouom         *string    `json:"mud_salinity_ouom"`
	Mud_salinity_uom          *string    `json:"mud_salinity_uom"`
	Mud_sample_type           *string    `json:"mud_sample_type"`
	Mud_temperature           *float64   `json:"mud_temperature"`
	Mud_temperature_ouom      *string    `json:"mud_temperature_ouom"`
	Mud_viscosity             *float64   `json:"mud_viscosity"`
	Mud_viscosity_ouom        *string    `json:"mud_viscosity_ouom"`
	Mud_weight                *float64   `json:"mud_weight"`
	Mud_weight_ouom           *string    `json:"mud_weight_ouom"`
	Mud_weight_uom            *string    `json:"mud_weight_uom"`
	Ppdm_guid                 string     `json:"ppdm_guid"`
	Remark                    *string    `json:"remark"`
	Row_changed_by            *string    `json:"row_changed_by"`
	Row_changed_date          *time.Time `json:"row_changed_date"`
	Row_created_by            *string    `json:"row_created_by"`
	Row_created_date          *time.Time `json:"row_created_date"`
	Row_effective_date        *time.Time `json:"row_effective_date"`
	Row_expiry_date           *time.Time `json:"row_expiry_date"`
	Row_quality               *string    `json:"row_quality"`
}
