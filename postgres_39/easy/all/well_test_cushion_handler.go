package apiv1

import (
	dto "github.com/DarrenMannuela/pt_gtn_bibliography/dto"
	"github.com/gofiber/fiber/v2"
	"strconv"
	"encoding/json"
)


func GetWellTestCushion(c *fiber.Ctx) error {
	rows, err := db.Query("select * from well_test_cushion")
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}
	defer rows.Close()
	var result []dto.Well_test_cushion

	for rows.Next() {
		var well_test_cushion dto.Well_test_cushion
		if err := rows.Scan(&well_test_cushion.Uwi, &well_test_cushion.Source, &well_test_cushion.Test_type, &well_test_cushion.Run_num, &well_test_cushion.Test_num, &well_test_cushion.Cushion_obs_no, &well_test_cushion.Active_ind, &well_test_cushion.Cushion_gas_pressure, &well_test_cushion.Cushion_gas_pressure_ouom, &well_test_cushion.Cushion_inhibitor_volume, &well_test_cushion.Cushion_inhibitor_vol_ouom, &well_test_cushion.Cushion_length, &well_test_cushion.Cushion_length_ouom, &well_test_cushion.Cushion_type, &well_test_cushion.Cushion_volume, &well_test_cushion.Cushion_volume_ouom, &well_test_cushion.Effective_date, &well_test_cushion.Expiry_date, &well_test_cushion.Ppdm_guid, &well_test_cushion.Remark, &well_test_cushion.Row_changed_by, &well_test_cushion.Row_changed_date, &well_test_cushion.Row_created_by, &well_test_cushion.Row_created_date, &well_test_cushion.Row_effective_date, &well_test_cushion.Row_expiry_date, &well_test_cushion.Row_quality); err != nil {
			return err // Exit if we get an error
		}

		// Append well_test_cushion to result
		result = append(result, well_test_cushion)
	}
	// Return result in JSON format
	return c.JSON(result)
	 
}

func SetWellTestCushion(c *fiber.Ctx) error {
	var well_test_cushion dto.Well_test_cushion

	setDefaults(&well_test_cushion)

	if err := json.Unmarshal(c.Body(), &well_test_cushion); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	stmt, err := db.Prepare("insert into well_test_cushion values(:1, :2, :3, :4, :5, :6, :7, :8, :9, :10, :11, :12, :13, :14, :15, :16, :17, :18, :19, :20, :21, :22, :23, :24, :25, :26, :27)")
	if err != nil {
		return err
	}
	well_test_cushion.Row_created_date = formatDate(well_test_cushion.Row_created_date)
	well_test_cushion.Row_changed_date = formatDate(well_test_cushion.Row_changed_date)
	well_test_cushion.Effective_date = formatDateString(well_test_cushion.Effective_date)
	well_test_cushion.Expiry_date = formatDateString(well_test_cushion.Expiry_date)
	well_test_cushion.Row_effective_date = formatDateString(well_test_cushion.Row_effective_date)
	well_test_cushion.Row_expiry_date = formatDateString(well_test_cushion.Row_expiry_date)






	rows, err := stmt.Exec(well_test_cushion.Uwi, well_test_cushion.Source, well_test_cushion.Test_type, well_test_cushion.Run_num, well_test_cushion.Test_num, well_test_cushion.Cushion_obs_no, well_test_cushion.Active_ind, well_test_cushion.Cushion_gas_pressure, well_test_cushion.Cushion_gas_pressure_ouom, well_test_cushion.Cushion_inhibitor_volume, well_test_cushion.Cushion_inhibitor_vol_ouom, well_test_cushion.Cushion_length, well_test_cushion.Cushion_length_ouom, well_test_cushion.Cushion_type, well_test_cushion.Cushion_volume, well_test_cushion.Cushion_volume_ouom, well_test_cushion.Effective_date, well_test_cushion.Expiry_date, well_test_cushion.Ppdm_guid, well_test_cushion.Remark, well_test_cushion.Row_changed_by, well_test_cushion.Row_changed_date, well_test_cushion.Row_created_by, well_test_cushion.Row_created_date, well_test_cushion.Row_effective_date, well_test_cushion.Row_expiry_date, well_test_cushion.Row_quality)
	
	if err != nil {
		return err
	}

	return c.JSON(rows)
}

func UpdateWellTestCushion(c *fiber.Ctx) error {
	var well_test_cushion dto.Well_test_cushion

	setDefaults(&well_test_cushion)

	if err := json.Unmarshal(c.Body(), &well_test_cushion); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	id := c.Params("id")
	well_test_cushion.Uwi = id

    if well_test_cushion.Row_created_date != nil {
        return c.Status(400).SendString("Cannot update row_created_date")
    }

	stmt, err := db.Prepare("update well_test_cushion set source = :1, test_type = :2, run_num = :3, test_num = :4, cushion_obs_no = :5, active_ind = :6, cushion_gas_pressure = :7, cushion_gas_pressure_ouom = :8, cushion_inhibitor_volume = :9, cushion_inhibitor_vol_ouom = :10, cushion_length = :11, cushion_length_ouom = :12, cushion_type = :13, cushion_volume = :14, cushion_volume_ouom = :15, effective_date = :16, expiry_date = :17, ppdm_guid = :18, remark = :19, row_changed_by = :20, row_changed_date = :21, row_created_by = :22, row_effective_date = :23, row_expiry_date = :24, row_quality = :25 where uwi = :27")
	if err != nil {
		return err
	}

	well_test_cushion.Row_changed_date = formatDate(well_test_cushion.Row_changed_date)
	well_test_cushion.Effective_date = formatDateString(well_test_cushion.Effective_date)
	well_test_cushion.Expiry_date = formatDateString(well_test_cushion.Expiry_date)
	well_test_cushion.Row_effective_date = formatDateString(well_test_cushion.Row_effective_date)
	well_test_cushion.Row_expiry_date = formatDateString(well_test_cushion.Row_expiry_date)






	rows, err := stmt.Exec(well_test_cushion.Source, well_test_cushion.Test_type, well_test_cushion.Run_num, well_test_cushion.Test_num, well_test_cushion.Cushion_obs_no, well_test_cushion.Active_ind, well_test_cushion.Cushion_gas_pressure, well_test_cushion.Cushion_gas_pressure_ouom, well_test_cushion.Cushion_inhibitor_volume, well_test_cushion.Cushion_inhibitor_vol_ouom, well_test_cushion.Cushion_length, well_test_cushion.Cushion_length_ouom, well_test_cushion.Cushion_type, well_test_cushion.Cushion_volume, well_test_cushion.Cushion_volume_ouom, well_test_cushion.Effective_date, well_test_cushion.Expiry_date, well_test_cushion.Ppdm_guid, well_test_cushion.Remark, well_test_cushion.Row_changed_by, well_test_cushion.Row_changed_date, well_test_cushion.Row_created_by, well_test_cushion.Row_effective_date, well_test_cushion.Row_expiry_date, well_test_cushion.Row_quality, well_test_cushion.Uwi)
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

func PatchWellTestCushion(c *fiber.Ctx) error {
	id := c.Params("id")

	updatedFields := make(map[string]interface{})

	if err := c.BodyParser(&updatedFields); err != nil {
		return err
	}
	if _, ok := updatedFields["row_created_date"]; ok {
        return c.Status(400).SendString("Cannot update row_created_date field")
    }

	query := "update well_test_cushion set "
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

func DeleteWellTestCushion(c *fiber.Ctx) error {
	id := c.Params("id")
	var well_test_cushion dto.Well_test_cushion
	well_test_cushion.Uwi = id

	stmt, err := db.Prepare("delete from well_test_cushion where uwi = :1")
	if err != nil {
		return err
	}

	rows, err := stmt.Exec(well_test_cushion.Uwi)
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


