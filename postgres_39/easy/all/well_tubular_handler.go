package apiv1

import (
	dto "github.com/DarrenMannuela/pt_gtn_bibliography/dto"
	"github.com/gofiber/fiber/v2"
	"strconv"
	"encoding/json"
)


func GetWellTubular(c *fiber.Ctx) error {
	rows, err := db.Query("select * from well_tubular")
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}
	defer rows.Close()
	var result []dto.Well_tubular

	for rows.Next() {
		var well_tubular dto.Well_tubular
		if err := rows.Scan(&well_tubular.Uwi, &well_tubular.Source, &well_tubular.Tubing_type, &well_tubular.Tubing_obs_no, &well_tubular.Active_ind, &well_tubular.Base_depth, &well_tubular.Base_depth_ouom, &well_tubular.Base_strat_unit_id, &well_tubular.Cat_equip_id, &well_tubular.Collar_type, &well_tubular.Coupling_type, &well_tubular.Effective_date, &well_tubular.Equipment_id, &well_tubular.Expiry_date, &well_tubular.Hole_size, &well_tubular.Hole_size_ouom, &well_tubular.Hung_top_depth, &well_tubular.Hung_top_depth_ouom, &well_tubular.Inside_diameter, &well_tubular.Inside_diameter_ouom, &well_tubular.Joint_count, &well_tubular.Left_in_hole_length, &well_tubular.Left_in_hole_length_ouom, &well_tubular.Liner_type, &well_tubular.Manufacturer_ba_id, &well_tubular.Mixed_string_ind, &well_tubular.Observation_date, &well_tubular.Outside_diameter, &well_tubular.Outside_diameter_desc, &well_tubular.Outside_diameter_ouom, &well_tubular.Packer_set_depth, &well_tubular.Packer_set_depth_ouom, &well_tubular.Period_obs_no, &well_tubular.Ppdm_guid, &well_tubular.Pulled_length, &well_tubular.Pulled_length_ouom, &well_tubular.Remark, &well_tubular.Reported_base_tvd, &well_tubular.Reported_base_tvd_ouom, &well_tubular.Reported_top_tvd, &well_tubular.Reported_top_tvd_ouom, &well_tubular.Sea_floor_penetration, &well_tubular.Sea_floor_penetration_ouom, &well_tubular.Shoe_depth, &well_tubular.Shoe_depth_ouom, &well_tubular.Steel_spec, &well_tubular.Strat_name_set_id, &well_tubular.Top_strat_unit_id, &well_tubular.Torque, &well_tubular.Torque_ouom, &well_tubular.Tubing_density, &well_tubular.Tubing_density_ouom, &well_tubular.Tubing_grade, &well_tubular.Tubing_strength, &well_tubular.Tubing_strength_ouom, &well_tubular.Tubing_weight, &well_tubular.Tubing_weight_ouom, &well_tubular.Row_changed_by, &well_tubular.Row_changed_date, &well_tubular.Row_created_by, &well_tubular.Row_created_date, &well_tubular.Row_effective_date, &well_tubular.Row_expiry_date, &well_tubular.Row_quality); err != nil {
			return err // Exit if we get an error
		}

		// Append well_tubular to result
		result = append(result, well_tubular)
	}
	// Return result in JSON format
	return c.JSON(result)
	 
}

func SetWellTubular(c *fiber.Ctx) error {
	var well_tubular dto.Well_tubular

	setDefaults(&well_tubular)

	if err := json.Unmarshal(c.Body(), &well_tubular); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	stmt, err := db.Prepare("insert into well_tubular values(:1, :2, :3, :4, :5, :6, :7, :8, :9, :10, :11, :12, :13, :14, :15, :16, :17, :18, :19, :20, :21, :22, :23, :24, :25, :26, :27, :28, :29, :30, :31, :32, :33, :34, :35, :36, :37, :38, :39, :40, :41, :42, :43, :44, :45, :46, :47, :48, :49, :50, :51, :52, :53, :54, :55, :56, :57, :58, :59, :60, :61, :62, :63, :64)")
	if err != nil {
		return err
	}
	well_tubular.Row_created_date = formatDate(well_tubular.Row_created_date)
	well_tubular.Row_changed_date = formatDate(well_tubular.Row_changed_date)
	well_tubular.Effective_date = formatDateString(well_tubular.Effective_date)
	well_tubular.Expiry_date = formatDateString(well_tubular.Expiry_date)
	well_tubular.Observation_date = formatDateString(well_tubular.Observation_date)
	well_tubular.Row_effective_date = formatDateString(well_tubular.Row_effective_date)
	well_tubular.Row_expiry_date = formatDateString(well_tubular.Row_expiry_date)






	rows, err := stmt.Exec(well_tubular.Uwi, well_tubular.Source, well_tubular.Tubing_type, well_tubular.Tubing_obs_no, well_tubular.Active_ind, well_tubular.Base_depth, well_tubular.Base_depth_ouom, well_tubular.Base_strat_unit_id, well_tubular.Cat_equip_id, well_tubular.Collar_type, well_tubular.Coupling_type, well_tubular.Effective_date, well_tubular.Equipment_id, well_tubular.Expiry_date, well_tubular.Hole_size, well_tubular.Hole_size_ouom, well_tubular.Hung_top_depth, well_tubular.Hung_top_depth_ouom, well_tubular.Inside_diameter, well_tubular.Inside_diameter_ouom, well_tubular.Joint_count, well_tubular.Left_in_hole_length, well_tubular.Left_in_hole_length_ouom, well_tubular.Liner_type, well_tubular.Manufacturer_ba_id, well_tubular.Mixed_string_ind, well_tubular.Observation_date, well_tubular.Outside_diameter, well_tubular.Outside_diameter_desc, well_tubular.Outside_diameter_ouom, well_tubular.Packer_set_depth, well_tubular.Packer_set_depth_ouom, well_tubular.Period_obs_no, well_tubular.Ppdm_guid, well_tubular.Pulled_length, well_tubular.Pulled_length_ouom, well_tubular.Remark, well_tubular.Reported_base_tvd, well_tubular.Reported_base_tvd_ouom, well_tubular.Reported_top_tvd, well_tubular.Reported_top_tvd_ouom, well_tubular.Sea_floor_penetration, well_tubular.Sea_floor_penetration_ouom, well_tubular.Shoe_depth, well_tubular.Shoe_depth_ouom, well_tubular.Steel_spec, well_tubular.Strat_name_set_id, well_tubular.Top_strat_unit_id, well_tubular.Torque, well_tubular.Torque_ouom, well_tubular.Tubing_density, well_tubular.Tubing_density_ouom, well_tubular.Tubing_grade, well_tubular.Tubing_strength, well_tubular.Tubing_strength_ouom, well_tubular.Tubing_weight, well_tubular.Tubing_weight_ouom, well_tubular.Row_changed_by, well_tubular.Row_changed_date, well_tubular.Row_created_by, well_tubular.Row_created_date, well_tubular.Row_effective_date, well_tubular.Row_expiry_date, well_tubular.Row_quality)
	
	if err != nil {
		return err
	}

	return c.JSON(rows)
}

func UpdateWellTubular(c *fiber.Ctx) error {
	var well_tubular dto.Well_tubular

	setDefaults(&well_tubular)

	if err := json.Unmarshal(c.Body(), &well_tubular); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	id := c.Params("id")
	well_tubular.Uwi = id

    if well_tubular.Row_created_date != nil {
        return c.Status(400).SendString("Cannot update row_created_date")
    }

	stmt, err := db.Prepare("update well_tubular set source = :1, tubing_type = :2, tubing_obs_no = :3, active_ind = :4, base_depth = :5, base_depth_ouom = :6, base_strat_unit_id = :7, cat_equip_id = :8, collar_type = :9, coupling_type = :10, effective_date = :11, equipment_id = :12, expiry_date = :13, hole_size = :14, hole_size_ouom = :15, hung_top_depth = :16, hung_top_depth_ouom = :17, inside_diameter = :18, inside_diameter_ouom = :19, joint_count = :20, left_in_hole_length = :21, left_in_hole_length_ouom = :22, liner_type = :23, manufacturer_ba_id = :24, mixed_string_ind = :25, observation_date = :26, outside_diameter = :27, outside_diameter_desc = :28, outside_diameter_ouom = :29, packer_set_depth = :30, packer_set_depth_ouom = :31, period_obs_no = :32, ppdm_guid = :33, pulled_length = :34, pulled_length_ouom = :35, remark = :36, reported_base_tvd = :37, reported_base_tvd_ouom = :38, reported_top_tvd = :39, reported_top_tvd_ouom = :40, sea_floor_penetration = :41, sea_floor_penetration_ouom = :42, shoe_depth = :43, shoe_depth_ouom = :44, steel_spec = :45, strat_name_set_id = :46, top_strat_unit_id = :47, torque = :48, torque_ouom = :49, tubing_density = :50, tubing_density_ouom = :51, tubing_grade = :52, tubing_strength = :53, tubing_strength_ouom = :54, tubing_weight = :55, tubing_weight_ouom = :56, row_changed_by = :57, row_changed_date = :58, row_created_by = :59, row_effective_date = :60, row_expiry_date = :61, row_quality = :62 where uwi = :64")
	if err != nil {
		return err
	}

	well_tubular.Row_changed_date = formatDate(well_tubular.Row_changed_date)
	well_tubular.Effective_date = formatDateString(well_tubular.Effective_date)
	well_tubular.Expiry_date = formatDateString(well_tubular.Expiry_date)
	well_tubular.Observation_date = formatDateString(well_tubular.Observation_date)
	well_tubular.Row_effective_date = formatDateString(well_tubular.Row_effective_date)
	well_tubular.Row_expiry_date = formatDateString(well_tubular.Row_expiry_date)






	rows, err := stmt.Exec(well_tubular.Source, well_tubular.Tubing_type, well_tubular.Tubing_obs_no, well_tubular.Active_ind, well_tubular.Base_depth, well_tubular.Base_depth_ouom, well_tubular.Base_strat_unit_id, well_tubular.Cat_equip_id, well_tubular.Collar_type, well_tubular.Coupling_type, well_tubular.Effective_date, well_tubular.Equipment_id, well_tubular.Expiry_date, well_tubular.Hole_size, well_tubular.Hole_size_ouom, well_tubular.Hung_top_depth, well_tubular.Hung_top_depth_ouom, well_tubular.Inside_diameter, well_tubular.Inside_diameter_ouom, well_tubular.Joint_count, well_tubular.Left_in_hole_length, well_tubular.Left_in_hole_length_ouom, well_tubular.Liner_type, well_tubular.Manufacturer_ba_id, well_tubular.Mixed_string_ind, well_tubular.Observation_date, well_tubular.Outside_diameter, well_tubular.Outside_diameter_desc, well_tubular.Outside_diameter_ouom, well_tubular.Packer_set_depth, well_tubular.Packer_set_depth_ouom, well_tubular.Period_obs_no, well_tubular.Ppdm_guid, well_tubular.Pulled_length, well_tubular.Pulled_length_ouom, well_tubular.Remark, well_tubular.Reported_base_tvd, well_tubular.Reported_base_tvd_ouom, well_tubular.Reported_top_tvd, well_tubular.Reported_top_tvd_ouom, well_tubular.Sea_floor_penetration, well_tubular.Sea_floor_penetration_ouom, well_tubular.Shoe_depth, well_tubular.Shoe_depth_ouom, well_tubular.Steel_spec, well_tubular.Strat_name_set_id, well_tubular.Top_strat_unit_id, well_tubular.Torque, well_tubular.Torque_ouom, well_tubular.Tubing_density, well_tubular.Tubing_density_ouom, well_tubular.Tubing_grade, well_tubular.Tubing_strength, well_tubular.Tubing_strength_ouom, well_tubular.Tubing_weight, well_tubular.Tubing_weight_ouom, well_tubular.Row_changed_by, well_tubular.Row_changed_date, well_tubular.Row_created_by, well_tubular.Row_effective_date, well_tubular.Row_expiry_date, well_tubular.Row_quality, well_tubular.Uwi)
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

func PatchWellTubular(c *fiber.Ctx) error {
	id := c.Params("id")

	updatedFields := make(map[string]interface{})

	if err := c.BodyParser(&updatedFields); err != nil {
		return err
	}
	if _, ok := updatedFields["row_created_date"]; ok {
        return c.Status(400).SendString("Cannot update row_created_date field")
    }

	query := "update well_tubular set "
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
		} else if key == "effective_date" || key == "expiry_date" || key == "observation_date" || key == "row_effective_date" || key == "row_expiry_date"      {
			if date, ok := value.(string); ok && len(date) > 0 {
				formattedDate := formatDateString(&date)
				value = formattedDate
			}
		}
		values = append(values, value)
	}
	query += " where uwi = :id"

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

func DeleteWellTubular(c *fiber.Ctx) error {
	id := c.Params("id")
	var well_tubular dto.Well_tubular
	well_tubular.Uwi = id

	stmt, err := db.Prepare("delete from well_tubular where uwi = :1")
	if err != nil {
		return err
	}

	rows, err := stmt.Exec(well_tubular.Uwi)
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


