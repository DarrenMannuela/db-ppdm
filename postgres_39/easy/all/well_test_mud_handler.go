package apiv1

import (
	dto "github.com/DarrenMannuela/pt_gtn_bibliography/dto"
	"github.com/gofiber/fiber/v2"
	"strconv"
	"encoding/json"
)


func GetWellTestMud(c *fiber.Ctx) error {
	rows, err := db.Query("select * from well_test_mud")
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}
	defer rows.Close()
	var result []dto.Well_test_mud

	for rows.Next() {
		var well_test_mud dto.Well_test_mud
		if err := rows.Scan(&well_test_mud.Uwi, &well_test_mud.Source, &well_test_mud.Test_type, &well_test_mud.Run_num, &well_test_mud.Test_num, &well_test_mud.Mud_type, &well_test_mud.Active_ind, &well_test_mud.Effective_date, &well_test_mud.Expiry_date, &well_test_mud.Filtrate_resistivity, &well_test_mud.Filtrate_resistivity_ouom, &well_test_mud.Filtrate_salinity, &well_test_mud.Filtrate_salinity_ouom, &well_test_mud.Filtrate_salinity_uom, &well_test_mud.Filtrate_temperature, &well_test_mud.Filtrate_temperature_ouom, &well_test_mud.Mud_ph, &well_test_mud.Mud_resistivity, &well_test_mud.Mud_resistivity_ouom, &well_test_mud.Mud_salinity, &well_test_mud.Mud_salinity_ouom, &well_test_mud.Mud_salinity_uom, &well_test_mud.Mud_sample_type, &well_test_mud.Mud_temperature, &well_test_mud.Mud_temperature_ouom, &well_test_mud.Mud_viscosity, &well_test_mud.Mud_viscosity_ouom, &well_test_mud.Mud_weight, &well_test_mud.Mud_weight_ouom, &well_test_mud.Mud_weight_uom, &well_test_mud.Ppdm_guid, &well_test_mud.Remark, &well_test_mud.Row_changed_by, &well_test_mud.Row_changed_date, &well_test_mud.Row_created_by, &well_test_mud.Row_created_date, &well_test_mud.Row_effective_date, &well_test_mud.Row_expiry_date, &well_test_mud.Row_quality); err != nil {
			return err // Exit if we get an error
		}

		// Append well_test_mud to result
		result = append(result, well_test_mud)
	}
	// Return result in JSON format
	return c.JSON(result)
	 
}

func SetWellTestMud(c *fiber.Ctx) error {
	var well_test_mud dto.Well_test_mud

	setDefaults(&well_test_mud)

	if err := json.Unmarshal(c.Body(), &well_test_mud); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	stmt, err := db.Prepare("insert into well_test_mud values(:1, :2, :3, :4, :5, :6, :7, :8, :9, :10, :11, :12, :13, :14, :15, :16, :17, :18, :19, :20, :21, :22, :23, :24, :25, :26, :27, :28, :29, :30, :31, :32, :33, :34, :35, :36, :37, :38, :39)")
	if err != nil {
		return err
	}
	well_test_mud.Row_created_date = formatDate(well_test_mud.Row_created_date)
	well_test_mud.Row_changed_date = formatDate(well_test_mud.Row_changed_date)
	well_test_mud.Effective_date = formatDateString(well_test_mud.Effective_date)
	well_test_mud.Expiry_date = formatDateString(well_test_mud.Expiry_date)
	well_test_mud.Row_effective_date = formatDateString(well_test_mud.Row_effective_date)
	well_test_mud.Row_expiry_date = formatDateString(well_test_mud.Row_expiry_date)






	rows, err := stmt.Exec(well_test_mud.Uwi, well_test_mud.Source, well_test_mud.Test_type, well_test_mud.Run_num, well_test_mud.Test_num, well_test_mud.Mud_type, well_test_mud.Active_ind, well_test_mud.Effective_date, well_test_mud.Expiry_date, well_test_mud.Filtrate_resistivity, well_test_mud.Filtrate_resistivity_ouom, well_test_mud.Filtrate_salinity, well_test_mud.Filtrate_salinity_ouom, well_test_mud.Filtrate_salinity_uom, well_test_mud.Filtrate_temperature, well_test_mud.Filtrate_temperature_ouom, well_test_mud.Mud_ph, well_test_mud.Mud_resistivity, well_test_mud.Mud_resistivity_ouom, well_test_mud.Mud_salinity, well_test_mud.Mud_salinity_ouom, well_test_mud.Mud_salinity_uom, well_test_mud.Mud_sample_type, well_test_mud.Mud_temperature, well_test_mud.Mud_temperature_ouom, well_test_mud.Mud_viscosity, well_test_mud.Mud_viscosity_ouom, well_test_mud.Mud_weight, well_test_mud.Mud_weight_ouom, well_test_mud.Mud_weight_uom, well_test_mud.Ppdm_guid, well_test_mud.Remark, well_test_mud.Row_changed_by, well_test_mud.Row_changed_date, well_test_mud.Row_created_by, well_test_mud.Row_created_date, well_test_mud.Row_effective_date, well_test_mud.Row_expiry_date, well_test_mud.Row_quality)
	
	if err != nil {
		return err
	}

	return c.JSON(rows)
}

func UpdateWellTestMud(c *fiber.Ctx) error {
	var well_test_mud dto.Well_test_mud

	setDefaults(&well_test_mud)

	if err := json.Unmarshal(c.Body(), &well_test_mud); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	id := c.Params("id")
	well_test_mud.Uwi = id

    if well_test_mud.Row_created_date != nil {
        return c.Status(400).SendString("Cannot update row_created_date")
    }

	stmt, err := db.Prepare("update well_test_mud set source = :1, test_type = :2, run_num = :3, test_num = :4, mud_type = :5, active_ind = :6, effective_date = :7, expiry_date = :8, filtrate_resistivity = :9, filtrate_resistivity_ouom = :10, filtrate_salinity = :11, filtrate_salinity_ouom = :12, filtrate_salinity_uom = :13, filtrate_temperature = :14, filtrate_temperature_ouom = :15, mud_ph = :16, mud_resistivity = :17, mud_resistivity_ouom = :18, mud_salinity = :19, mud_salinity_ouom = :20, mud_salinity_uom = :21, mud_sample_type = :22, mud_temperature = :23, mud_temperature_ouom = :24, mud_viscosity = :25, mud_viscosity_ouom = :26, mud_weight = :27, mud_weight_ouom = :28, mud_weight_uom = :29, ppdm_guid = :30, remark = :31, row_changed_by = :32, row_changed_date = :33, row_created_by = :34, row_effective_date = :35, row_expiry_date = :36, row_quality = :37 where uwi = :39")
	if err != nil {
		return err
	}

	well_test_mud.Row_changed_date = formatDate(well_test_mud.Row_changed_date)
	well_test_mud.Effective_date = formatDateString(well_test_mud.Effective_date)
	well_test_mud.Expiry_date = formatDateString(well_test_mud.Expiry_date)
	well_test_mud.Row_effective_date = formatDateString(well_test_mud.Row_effective_date)
	well_test_mud.Row_expiry_date = formatDateString(well_test_mud.Row_expiry_date)






	rows, err := stmt.Exec(well_test_mud.Source, well_test_mud.Test_type, well_test_mud.Run_num, well_test_mud.Test_num, well_test_mud.Mud_type, well_test_mud.Active_ind, well_test_mud.Effective_date, well_test_mud.Expiry_date, well_test_mud.Filtrate_resistivity, well_test_mud.Filtrate_resistivity_ouom, well_test_mud.Filtrate_salinity, well_test_mud.Filtrate_salinity_ouom, well_test_mud.Filtrate_salinity_uom, well_test_mud.Filtrate_temperature, well_test_mud.Filtrate_temperature_ouom, well_test_mud.Mud_ph, well_test_mud.Mud_resistivity, well_test_mud.Mud_resistivity_ouom, well_test_mud.Mud_salinity, well_test_mud.Mud_salinity_ouom, well_test_mud.Mud_salinity_uom, well_test_mud.Mud_sample_type, well_test_mud.Mud_temperature, well_test_mud.Mud_temperature_ouom, well_test_mud.Mud_viscosity, well_test_mud.Mud_viscosity_ouom, well_test_mud.Mud_weight, well_test_mud.Mud_weight_ouom, well_test_mud.Mud_weight_uom, well_test_mud.Ppdm_guid, well_test_mud.Remark, well_test_mud.Row_changed_by, well_test_mud.Row_changed_date, well_test_mud.Row_created_by, well_test_mud.Row_effective_date, well_test_mud.Row_expiry_date, well_test_mud.Row_quality, well_test_mud.Uwi)
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

func PatchWellTestMud(c *fiber.Ctx) error {
	id := c.Params("id")

	updatedFields := make(map[string]interface{})

	if err := c.BodyParser(&updatedFields); err != nil {
		return err
	}
	if _, ok := updatedFields["row_created_date"]; ok {
        return c.Status(400).SendString("Cannot update row_created_date field")
    }

	query := "update well_test_mud set "
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

func DeleteWellTestMud(c *fiber.Ctx) error {
	id := c.Params("id")
	var well_test_mud dto.Well_test_mud
	well_test_mud.Uwi = id

	stmt, err := db.Prepare("delete from well_test_mud where uwi = :1")
	if err != nil {
		return err
	}

	rows, err := stmt.Exec(well_test_mud.Uwi)
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


