package apiv1

import (
	dto "github.com/DarrenMannuela/pt_gtn_bibliography/dto"
	"github.com/gofiber/fiber/v2"
	"strconv"
	"encoding/json"
)


func GetSfRig(c *fiber.Ctx) error {
	rows, err := db.Query("select * from sf_rig")
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}
	defer rows.Close()
	var result []dto.Sf_rig

	for rows.Next() {
		var sf_rig dto.Sf_rig
		if err := rows.Scan(&sf_rig.Sf_subtype, &sf_rig.Rig_id, &sf_rig.Active_ind, &sf_rig.Boiler_manufacturer, &sf_rig.Boiler_power, &sf_rig.Boiler_power_ouom, &sf_rig.Choke_manifold_press, &sf_rig.Choke_manifold_press_ouom, &sf_rig.Choke_manifold_size, &sf_rig.Choke_manifold_size_ouom, &sf_rig.Clear_work_height, &sf_rig.Clear_work_height_ouom, &sf_rig.Commission_date, &sf_rig.Decommission_date, &sf_rig.Desander_type, &sf_rig.Desilter_type, &sf_rig.Drawworks_type, &sf_rig.Effective_date, &sf_rig.Expiry_date, &sf_rig.Hookblock_rating, &sf_rig.Hookblock_rating_ouom, &sf_rig.Hookblock_type, &sf_rig.Hook_load_capacity, &sf_rig.Hook_load_capacity_ouom, &sf_rig.Kb_ground_dist, &sf_rig.Kb_ground_dist_ouom, &sf_rig.Mast_height, &sf_rig.Mast_height_ouom, &sf_rig.Mast_type, &sf_rig.Max_depth, &sf_rig.Max_depth_ouom, &sf_rig.Mud_tank_count, &sf_rig.Mud_tank_size, &sf_rig.Mud_tank_size_ouom, &sf_rig.Operator_ba_id, &sf_rig.Owner_ba_id, &sf_rig.Ppdm_guid, &sf_rig.Remark, &sf_rig.Rig_category, &sf_rig.Rig_class, &sf_rig.Rig_code, &sf_rig.Rig_name, &sf_rig.Rig_type, &sf_rig.Rotary_table_cap, &sf_rig.Rotary_table_cap_ouom, &sf_rig.Rotary_table_size, &sf_rig.Rotary_table_size_ouom, &sf_rig.Source, &sf_rig.Substructure_type, &sf_rig.Substruct_casing_cap, &sf_rig.Substruct_casing_cap_ouom, &sf_rig.Substruct_setback_cap, &sf_rig.Substruct_setback_cap_ouom, &sf_rig.Swivel_rating, &sf_rig.Swivel_rating_ouom, &sf_rig.Swivel_type, &sf_rig.Top_drive_model, &sf_rig.Top_drive_rating, &sf_rig.Top_drive_rating_ouom, &sf_rig.Water_depth_capacity, &sf_rig.Water_depth_datum, &sf_rig.Water_depth_ouom, &sf_rig.Working_pressure, &sf_rig.Working_pressure_ouom, &sf_rig.Row_changed_by, &sf_rig.Row_changed_date, &sf_rig.Row_created_by, &sf_rig.Row_created_date, &sf_rig.Row_effective_date, &sf_rig.Row_expiry_date, &sf_rig.Row_quality); err != nil {
			return err // Exit if we get an error
		}

		// Append sf_rig to result
		result = append(result, sf_rig)
	}
	// Return result in JSON format
	return c.JSON(result)
	 
}

func SetSfRig(c *fiber.Ctx) error {
	var sf_rig dto.Sf_rig

	setDefaults(&sf_rig)

	if err := json.Unmarshal(c.Body(), &sf_rig); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	stmt, err := db.Prepare("insert into sf_rig values(:1, :2, :3, :4, :5, :6, :7, :8, :9, :10, :11, :12, :13, :14, :15, :16, :17, :18, :19, :20, :21, :22, :23, :24, :25, :26, :27, :28, :29, :30, :31, :32, :33, :34, :35, :36, :37, :38, :39, :40, :41, :42, :43, :44, :45, :46, :47, :48, :49, :50, :51, :52, :53, :54, :55, :56, :57, :58, :59, :60, :61, :62, :63, :64, :65, :66, :67, :68, :69, :70, :71)")
	if err != nil {
		return err
	}
	sf_rig.Row_created_date = formatDate(sf_rig.Row_created_date)
	sf_rig.Row_changed_date = formatDate(sf_rig.Row_changed_date)
	sf_rig.Commission_date = formatDateString(sf_rig.Commission_date)
	sf_rig.Decommission_date = formatDateString(sf_rig.Decommission_date)
	sf_rig.Effective_date = formatDateString(sf_rig.Effective_date)
	sf_rig.Expiry_date = formatDateString(sf_rig.Expiry_date)
	sf_rig.Row_effective_date = formatDateString(sf_rig.Row_effective_date)
	sf_rig.Row_expiry_date = formatDateString(sf_rig.Row_expiry_date)






	rows, err := stmt.Exec(sf_rig.Sf_subtype, sf_rig.Rig_id, sf_rig.Active_ind, sf_rig.Boiler_manufacturer, sf_rig.Boiler_power, sf_rig.Boiler_power_ouom, sf_rig.Choke_manifold_press, sf_rig.Choke_manifold_press_ouom, sf_rig.Choke_manifold_size, sf_rig.Choke_manifold_size_ouom, sf_rig.Clear_work_height, sf_rig.Clear_work_height_ouom, sf_rig.Commission_date, sf_rig.Decommission_date, sf_rig.Desander_type, sf_rig.Desilter_type, sf_rig.Drawworks_type, sf_rig.Effective_date, sf_rig.Expiry_date, sf_rig.Hookblock_rating, sf_rig.Hookblock_rating_ouom, sf_rig.Hookblock_type, sf_rig.Hook_load_capacity, sf_rig.Hook_load_capacity_ouom, sf_rig.Kb_ground_dist, sf_rig.Kb_ground_dist_ouom, sf_rig.Mast_height, sf_rig.Mast_height_ouom, sf_rig.Mast_type, sf_rig.Max_depth, sf_rig.Max_depth_ouom, sf_rig.Mud_tank_count, sf_rig.Mud_tank_size, sf_rig.Mud_tank_size_ouom, sf_rig.Operator_ba_id, sf_rig.Owner_ba_id, sf_rig.Ppdm_guid, sf_rig.Remark, sf_rig.Rig_category, sf_rig.Rig_class, sf_rig.Rig_code, sf_rig.Rig_name, sf_rig.Rig_type, sf_rig.Rotary_table_cap, sf_rig.Rotary_table_cap_ouom, sf_rig.Rotary_table_size, sf_rig.Rotary_table_size_ouom, sf_rig.Source, sf_rig.Substructure_type, sf_rig.Substruct_casing_cap, sf_rig.Substruct_casing_cap_ouom, sf_rig.Substruct_setback_cap, sf_rig.Substruct_setback_cap_ouom, sf_rig.Swivel_rating, sf_rig.Swivel_rating_ouom, sf_rig.Swivel_type, sf_rig.Top_drive_model, sf_rig.Top_drive_rating, sf_rig.Top_drive_rating_ouom, sf_rig.Water_depth_capacity, sf_rig.Water_depth_datum, sf_rig.Water_depth_ouom, sf_rig.Working_pressure, sf_rig.Working_pressure_ouom, sf_rig.Row_changed_by, sf_rig.Row_changed_date, sf_rig.Row_created_by, sf_rig.Row_created_date, sf_rig.Row_effective_date, sf_rig.Row_expiry_date, sf_rig.Row_quality)
	
	if err != nil {
		return err
	}

	return c.JSON(rows)
}

func UpdateSfRig(c *fiber.Ctx) error {
	var sf_rig dto.Sf_rig

	setDefaults(&sf_rig)

	if err := json.Unmarshal(c.Body(), &sf_rig); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	id := c.Params("id")
	sf_rig.Sf_subtype = id

    if sf_rig.Row_created_date != nil {
        return c.Status(400).SendString("Cannot update row_created_date")
    }

	stmt, err := db.Prepare("update sf_rig set rig_id = :1, active_ind = :2, boiler_manufacturer = :3, boiler_power = :4, boiler_power_ouom = :5, choke_manifold_press = :6, choke_manifold_press_ouom = :7, choke_manifold_size = :8, choke_manifold_size_ouom = :9, clear_work_height = :10, clear_work_height_ouom = :11, commission_date = :12, decommission_date = :13, desander_type = :14, desilter_type = :15, drawworks_type = :16, effective_date = :17, expiry_date = :18, hookblock_rating = :19, hookblock_rating_ouom = :20, hookblock_type = :21, hook_load_capacity = :22, hook_load_capacity_ouom = :23, kb_ground_dist = :24, kb_ground_dist_ouom = :25, mast_height = :26, mast_height_ouom = :27, mast_type = :28, max_depth = :29, max_depth_ouom = :30, mud_tank_count = :31, mud_tank_size = :32, mud_tank_size_ouom = :33, operator_ba_id = :34, owner_ba_id = :35, ppdm_guid = :36, remark = :37, rig_category = :38, rig_class = :39, rig_code = :40, rig_name = :41, rig_type = :42, rotary_table_cap = :43, rotary_table_cap_ouom = :44, rotary_table_size = :45, rotary_table_size_ouom = :46, source = :47, substructure_type = :48, substruct_casing_cap = :49, substruct_casing_cap_ouom = :50, substruct_setback_cap = :51, substruct_setback_cap_ouom = :52, swivel_rating = :53, swivel_rating_ouom = :54, swivel_type = :55, top_drive_model = :56, top_drive_rating = :57, top_drive_rating_ouom = :58, water_depth_capacity = :59, water_depth_datum = :60, water_depth_ouom = :61, working_pressure = :62, working_pressure_ouom = :63, row_changed_by = :64, row_changed_date = :65, row_created_by = :66, row_effective_date = :67, row_expiry_date = :68, row_quality = :69 where sf_subtype = :71")
	if err != nil {
		return err
	}

	sf_rig.Row_changed_date = formatDate(sf_rig.Row_changed_date)
	sf_rig.Commission_date = formatDateString(sf_rig.Commission_date)
	sf_rig.Decommission_date = formatDateString(sf_rig.Decommission_date)
	sf_rig.Effective_date = formatDateString(sf_rig.Effective_date)
	sf_rig.Expiry_date = formatDateString(sf_rig.Expiry_date)
	sf_rig.Row_effective_date = formatDateString(sf_rig.Row_effective_date)
	sf_rig.Row_expiry_date = formatDateString(sf_rig.Row_expiry_date)






	rows, err := stmt.Exec(sf_rig.Rig_id, sf_rig.Active_ind, sf_rig.Boiler_manufacturer, sf_rig.Boiler_power, sf_rig.Boiler_power_ouom, sf_rig.Choke_manifold_press, sf_rig.Choke_manifold_press_ouom, sf_rig.Choke_manifold_size, sf_rig.Choke_manifold_size_ouom, sf_rig.Clear_work_height, sf_rig.Clear_work_height_ouom, sf_rig.Commission_date, sf_rig.Decommission_date, sf_rig.Desander_type, sf_rig.Desilter_type, sf_rig.Drawworks_type, sf_rig.Effective_date, sf_rig.Expiry_date, sf_rig.Hookblock_rating, sf_rig.Hookblock_rating_ouom, sf_rig.Hookblock_type, sf_rig.Hook_load_capacity, sf_rig.Hook_load_capacity_ouom, sf_rig.Kb_ground_dist, sf_rig.Kb_ground_dist_ouom, sf_rig.Mast_height, sf_rig.Mast_height_ouom, sf_rig.Mast_type, sf_rig.Max_depth, sf_rig.Max_depth_ouom, sf_rig.Mud_tank_count, sf_rig.Mud_tank_size, sf_rig.Mud_tank_size_ouom, sf_rig.Operator_ba_id, sf_rig.Owner_ba_id, sf_rig.Ppdm_guid, sf_rig.Remark, sf_rig.Rig_category, sf_rig.Rig_class, sf_rig.Rig_code, sf_rig.Rig_name, sf_rig.Rig_type, sf_rig.Rotary_table_cap, sf_rig.Rotary_table_cap_ouom, sf_rig.Rotary_table_size, sf_rig.Rotary_table_size_ouom, sf_rig.Source, sf_rig.Substructure_type, sf_rig.Substruct_casing_cap, sf_rig.Substruct_casing_cap_ouom, sf_rig.Substruct_setback_cap, sf_rig.Substruct_setback_cap_ouom, sf_rig.Swivel_rating, sf_rig.Swivel_rating_ouom, sf_rig.Swivel_type, sf_rig.Top_drive_model, sf_rig.Top_drive_rating, sf_rig.Top_drive_rating_ouom, sf_rig.Water_depth_capacity, sf_rig.Water_depth_datum, sf_rig.Water_depth_ouom, sf_rig.Working_pressure, sf_rig.Working_pressure_ouom, sf_rig.Row_changed_by, sf_rig.Row_changed_date, sf_rig.Row_created_by, sf_rig.Row_effective_date, sf_rig.Row_expiry_date, sf_rig.Row_quality, sf_rig.Sf_subtype)
	if err != nil {
		return err
	}

	if count, err := rows.RowsAffected(); err == nil {
		if count == 0 {
			return c.Status(404).SendString("No matching record found")
		}
	} else {
		return err
	}

	return c.Status(201).JSON(rows)
}

func PatchSfRig(c *fiber.Ctx) error {
	id := c.Params("id")

	updatedFields := make(map[string]interface{})

	if err := c.BodyParser(&updatedFields); err != nil {
		return err
	}
	if _, ok := updatedFields["row_created_date"]; ok {
        return c.Status(400).SendString("Cannot update row_created_date field")
    }

	query := "update sf_rig set "
	values := []interface{}{}
	i := 1
	for key, value := range updatedFields {
		query += key + " = :" + strconv.Itoa(i)
		i++
		if i <= len(updatedFields) {
			query += ", "
		}
		if key == "row_changed_date" {
			if date, ok := value.(string); ok && len(date) > 0 {
				formattedDate := formatDate(&date)
				value = formattedDate
			}
		} else if key == "commission_date" || key == "decommission_date" || key == "effective_date" || key == "expiry_date" || key == "row_effective_date" || key == "row_expiry_date"      {
			if date, ok := value.(string); ok && len(date) > 0 {
				formattedDate := formatDateString(&date)
				value = formattedDate
			}
		}
		values = append(values, value)
	}
	query += " where sf_subtype = :id"

	stmt, err := db.Prepare(query)
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}

	values = append(values, id)
	res, err := stmt.Exec(values...)
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}
	if count, err := res.RowsAffected(); err == nil {
		if count == 0 {
			return c.Status(404).SendString("No matching record found")
		}
	} else {
		return err
	}

	return c.JSON(res)
}

func DeleteSfRig(c *fiber.Ctx) error {
	id := c.Params("id")
	var sf_rig dto.Sf_rig
	sf_rig.Sf_subtype = id

	stmt, err := db.Prepare("delete from sf_rig where sf_subtype = :1")
	if err != nil {
		return err
	}

	rows, err := stmt.Exec(sf_rig.Sf_subtype)
	if err != nil {
		return err
	}

	rowsAffected, err := rows.RowsAffected()
	if err != nil {
		return err
	}

	if rowsAffected == 0 {
		return c.Status(404).SendString("No matching record found")
	}

	return c.JSON(rows)
}


