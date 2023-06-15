package apiv1

import (
	dto "github.com/DarrenMannuela/pt_gtn_bibliography/dto"
	"github.com/gofiber/fiber/v2"
	"strconv"
	"encoding/json"
)


func GetWellTestEquipment(c *fiber.Ctx) error {
	rows, err := db.Query("select * from well_test_equipment")
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}
	defer rows.Close()
	var result []dto.Well_test_equipment

	for rows.Next() {
		var well_test_equipment dto.Well_test_equipment
		if err := rows.Scan(&well_test_equipment.Uwi, &well_test_equipment.Source, &well_test_equipment.Test_type, &well_test_equipment.Run_num, &well_test_equipment.Test_num, &well_test_equipment.Equipment_type, &well_test_equipment.Equip_obs_no, &well_test_equipment.Active_ind, &well_test_equipment.Effective_date, &well_test_equipment.Equipment_id, &well_test_equipment.Equip_length, &well_test_equipment.Equip_length_ouom, &well_test_equipment.Equip_weight, &well_test_equipment.Equip_weight_ouom, &well_test_equipment.Expiry_date, &well_test_equipment.Inside_diameter, &well_test_equipment.Inside_diameter_ouom, &well_test_equipment.Outside_diameter, &well_test_equipment.Outside_diameter_ouom, &well_test_equipment.Ppdm_guid, &well_test_equipment.Remark, &well_test_equipment.Top_depth, &well_test_equipment.Top_depth_ouom, &well_test_equipment.Row_changed_by, &well_test_equipment.Row_changed_date, &well_test_equipment.Row_created_by, &well_test_equipment.Row_created_date, &well_test_equipment.Row_effective_date, &well_test_equipment.Row_expiry_date, &well_test_equipment.Row_quality); err != nil {
			return err // Exit if we get an error
		}

		// Append well_test_equipment to result
		result = append(result, well_test_equipment)
	}
	// Return result in JSON format
	return c.JSON(result)
	 
}

func SetWellTestEquipment(c *fiber.Ctx) error {
	var well_test_equipment dto.Well_test_equipment

	setDefaults(&well_test_equipment)

	if err := json.Unmarshal(c.Body(), &well_test_equipment); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	stmt, err := db.Prepare("insert into well_test_equipment values(:1, :2, :3, :4, :5, :6, :7, :8, :9, :10, :11, :12, :13, :14, :15, :16, :17, :18, :19, :20, :21, :22, :23, :24, :25, :26, :27, :28, :29, :30)")
	if err != nil {
		return err
	}
	well_test_equipment.Row_created_date = formatDate(well_test_equipment.Row_created_date)
	well_test_equipment.Row_changed_date = formatDate(well_test_equipment.Row_changed_date)
	well_test_equipment.Effective_date = formatDateString(well_test_equipment.Effective_date)
	well_test_equipment.Expiry_date = formatDateString(well_test_equipment.Expiry_date)
	well_test_equipment.Row_effective_date = formatDateString(well_test_equipment.Row_effective_date)
	well_test_equipment.Row_expiry_date = formatDateString(well_test_equipment.Row_expiry_date)






	rows, err := stmt.Exec(well_test_equipment.Uwi, well_test_equipment.Source, well_test_equipment.Test_type, well_test_equipment.Run_num, well_test_equipment.Test_num, well_test_equipment.Equipment_type, well_test_equipment.Equip_obs_no, well_test_equipment.Active_ind, well_test_equipment.Effective_date, well_test_equipment.Equipment_id, well_test_equipment.Equip_length, well_test_equipment.Equip_length_ouom, well_test_equipment.Equip_weight, well_test_equipment.Equip_weight_ouom, well_test_equipment.Expiry_date, well_test_equipment.Inside_diameter, well_test_equipment.Inside_diameter_ouom, well_test_equipment.Outside_diameter, well_test_equipment.Outside_diameter_ouom, well_test_equipment.Ppdm_guid, well_test_equipment.Remark, well_test_equipment.Top_depth, well_test_equipment.Top_depth_ouom, well_test_equipment.Row_changed_by, well_test_equipment.Row_changed_date, well_test_equipment.Row_created_by, well_test_equipment.Row_created_date, well_test_equipment.Row_effective_date, well_test_equipment.Row_expiry_date, well_test_equipment.Row_quality)
	
	if err != nil {
		return err
	}

	return c.JSON(rows)
}

func UpdateWellTestEquipment(c *fiber.Ctx) error {
	var well_test_equipment dto.Well_test_equipment

	setDefaults(&well_test_equipment)

	if err := json.Unmarshal(c.Body(), &well_test_equipment); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	id := c.Params("id")
	well_test_equipment.Uwi = id

    if well_test_equipment.Row_created_date != nil {
        return c.Status(400).SendString("Cannot update row_created_date")
    }

	stmt, err := db.Prepare("update well_test_equipment set source = :1, test_type = :2, run_num = :3, test_num = :4, equipment_type = :5, equip_obs_no = :6, active_ind = :7, effective_date = :8, equipment_id = :9, equip_length = :10, equip_length_ouom = :11, equip_weight = :12, equip_weight_ouom = :13, expiry_date = :14, inside_diameter = :15, inside_diameter_ouom = :16, outside_diameter = :17, outside_diameter_ouom = :18, ppdm_guid = :19, remark = :20, top_depth = :21, top_depth_ouom = :22, row_changed_by = :23, row_changed_date = :24, row_created_by = :25, row_effective_date = :26, row_expiry_date = :27, row_quality = :28 where uwi = :30")
	if err != nil {
		return err
	}

	well_test_equipment.Row_changed_date = formatDate(well_test_equipment.Row_changed_date)
	well_test_equipment.Effective_date = formatDateString(well_test_equipment.Effective_date)
	well_test_equipment.Expiry_date = formatDateString(well_test_equipment.Expiry_date)
	well_test_equipment.Row_effective_date = formatDateString(well_test_equipment.Row_effective_date)
	well_test_equipment.Row_expiry_date = formatDateString(well_test_equipment.Row_expiry_date)






	rows, err := stmt.Exec(well_test_equipment.Source, well_test_equipment.Test_type, well_test_equipment.Run_num, well_test_equipment.Test_num, well_test_equipment.Equipment_type, well_test_equipment.Equip_obs_no, well_test_equipment.Active_ind, well_test_equipment.Effective_date, well_test_equipment.Equipment_id, well_test_equipment.Equip_length, well_test_equipment.Equip_length_ouom, well_test_equipment.Equip_weight, well_test_equipment.Equip_weight_ouom, well_test_equipment.Expiry_date, well_test_equipment.Inside_diameter, well_test_equipment.Inside_diameter_ouom, well_test_equipment.Outside_diameter, well_test_equipment.Outside_diameter_ouom, well_test_equipment.Ppdm_guid, well_test_equipment.Remark, well_test_equipment.Top_depth, well_test_equipment.Top_depth_ouom, well_test_equipment.Row_changed_by, well_test_equipment.Row_changed_date, well_test_equipment.Row_created_by, well_test_equipment.Row_effective_date, well_test_equipment.Row_expiry_date, well_test_equipment.Row_quality, well_test_equipment.Uwi)
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

func PatchWellTestEquipment(c *fiber.Ctx) error {
	id := c.Params("id")

	updatedFields := make(map[string]interface{})

	if err := c.BodyParser(&updatedFields); err != nil {
		return err
	}
	if _, ok := updatedFields["row_created_date"]; ok {
        return c.Status(400).SendString("Cannot update row_created_date field")
    }

	query := "update well_test_equipment set "
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
		} else if key == "effective_date" || key == "expiry_date" || key == "row_effective_date" || key == "row_expiry_date"      {
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

func DeleteWellTestEquipment(c *fiber.Ctx) error {
	id := c.Params("id")
	var well_test_equipment dto.Well_test_equipment
	well_test_equipment.Uwi = id

	stmt, err := db.Prepare("delete from well_test_equipment where uwi = :1")
	if err != nil {
		return err
	}

	rows, err := stmt.Exec(well_test_equipment.Uwi)
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


