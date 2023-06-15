package dto

import (
	"time"
)

type Sf_rig struct {
	Sf_subtype                 string     `json:"sf_subtype"`
	Rig_id                     string     `json:"rig_id"`
	Active_ind                 *string    `json:"active_ind"`
	Boiler_manufacturer        *string    `json:"boiler_manufacturer"`
	Boiler_power               *float64   `json:"boiler_power"`
	Boiler_power_ouom          *string    `json:"boiler_power_ouom"`
	Choke_manifold_press       *float64   `json:"choke_manifold_press"`
	Choke_manifold_press_ouom  *string    `json:"choke_manifold_press_ouom"`
	Choke_manifold_size        *float64   `json:"choke_manifold_size"`
	Choke_manifold_size_ouom   *string    `json:"choke_manifold_size_ouom"`
	Clear_work_height          *float64   `json:"clear_work_height"`
	Clear_work_height_ouom     *string    `json:"clear_work_height_ouom"`
	Commission_date            *time.Time `json:"commission_date"`
	Decommission_date          *time.Time `json:"decommission_date"`
	Desander_type              *string    `json:"desander_type"`
	Desilter_type              *string    `json:"desilter_type"`
	Drawworks_type             *string    `json:"drawworks_type"`
	Effective_date             *time.Time `json:"effective_date"`
	Expiry_date                *time.Time `json:"expiry_date"`
	Hookblock_rating           *float64   `json:"hookblock_rating"`
	Hookblock_rating_ouom      *string    `json:"hookblock_rating_ouom"`
	Hookblock_type             *string    `json:"hookblock_type"`
	Hook_load_capacity         *float64   `json:"hook_load_capacity"`
	Hook_load_capacity_ouom    *string    `json:"hook_load_capacity_ouom"`
	Kb_ground_dist             *float64   `json:"kb_ground_dist"`
	Kb_ground_dist_ouom        *string    `json:"kb_ground_dist_ouom"`
	Mast_height                *float64   `json:"mast_height"`
	Mast_height_ouom           *string    `json:"mast_height_ouom"`
	Mast_type                  *string    `json:"mast_type"`
	Max_depth                  *float64   `json:"max_depth"`
	Max_depth_ouom             *string    `json:"max_depth_ouom"`
	Mud_tank_count             *int       `json:"mud_tank_count"`
	Mud_tank_size              *float64   `json:"mud_tank_size"`
	Mud_tank_size_ouom         *string    `json:"mud_tank_size_ouom"`
	Operator_ba_id             *string    `json:"operator_ba_id"`
	Owner_ba_id                *string    `json:"owner_ba_id"`
	Ppdm_guid                  string     `json:"ppdm_guid"`
	Remark                     *string    `json:"remark"`
	Rig_category               *string    `json:"rig_category"`
	Rig_class                  *string    `json:"rig_class"`
	Rig_code                   *string    `json:"rig_code"`
	Rig_name                   *string    `json:"rig_name"`
	Rig_type                   *string    `json:"rig_type"`
	Rotary_table_cap           *float64   `json:"rotary_table_cap"`
	Rotary_table_cap_ouom      *string    `json:"rotary_table_cap_ouom"`
	Rotary_table_size          *float64   `json:"rotary_table_size"`
	Rotary_table_size_ouom     *string    `json:"rotary_table_size_ouom"`
	Source                     *string    `json:"source"`
	Substructure_type          *string    `json:"substructure_type"`
	Substruct_casing_cap       *float64   `json:"substruct_casing_cap"`
	Substruct_casing_cap_ouom  *string    `json:"substruct_casing_cap_ouom"`
	Substruct_setback_cap      *float64   `json:"substruct_setback_cap"`
	Substruct_setback_cap_ouom *string    `json:"substruct_setback_cap_ouom"`
	Swivel_rating              *float64   `json:"swivel_rating"`
	Swivel_rating_ouom         *string    `json:"swivel_rating_ouom"`
	Swivel_type                *string    `json:"swivel_type"`
	Top_drive_model            *string    `json:"top_drive_model"`
	Top_drive_rating           *float64   `json:"top_drive_rating"`
	Top_drive_rating_ouom      *string    `json:"top_drive_rating_ouom"`
	Water_depth_capacity       *float64   `json:"water_depth_capacity"`
	Water_depth_datum          *string    `json:"water_depth_datum"`
	Water_depth_ouom           *string    `json:"water_depth_ouom"`
	Working_pressure           *float64   `json:"working_pressure"`
	Working_pressure_ouom      *string    `json:"working_pressure_ouom"`
	Row_changed_by             *string    `json:"row_changed_by"`
	Row_changed_date           *time.Time `json:"row_changed_date"`
	Row_created_by             *string    `json:"row_created_by"`
	Row_created_date           *time.Time `json:"row_created_date"`
	Row_effective_date         *time.Time `json:"row_effective_date"`
	Row_expiry_date            *time.Time `json:"row_expiry_date"`
	Row_quality                *string    `json:"row_quality"`
}
