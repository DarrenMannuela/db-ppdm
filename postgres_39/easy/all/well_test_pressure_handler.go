package apiv1

import (
	dto "github.com/DarrenMannuela/pt_gtn_bibliography/dto"
	"github.com/gofiber/fiber/v2"
	"strconv"
	"encoding/json"
)


func GetWellTestPressure(c *fiber.Ctx) error {
	rows, err := db.Query("select * from well_test_pressure")
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}
	defer rows.Close()
	var result []dto.Well_test_pressure

	for rows.Next() {
		var well_test_pressure dto.Well_test_pressure
		if err := rows.Scan(&well_test_pressure.Uwi, &well_test_pressure.Source, &well_test_pressure.Test_type, &well_test_pressure.Run_num, &well_test_pressure.Test_num, &well_test_pressure.Period_type, &well_test_pressure.Period_obs_no, &well_test_pressure.Active_ind, &well_test_pressure.Effective_date, &well_test_pressure.End_pressure, &well_test_pressure.End_pressure_ouom, &well_test_pressure.Expiry_date, &well_test_pressure.Ppdm_guid, &well_test_pressure.Quality_code, &well_test_pressure.Recorder_id, &well_test_pressure.Remark, &well_test_pressure.Start_pressure, &well_test_pressure.Start_pressure_ouom, &well_test_pressure.Summary_ind, &well_test_pressure.Row_changed_by, &well_test_pressure.Row_changed_date, &well_test_pressure.Row_created_by, &well_test_pressure.Row_created_date, &well_test_pressure.Row_effective_date, &well_test_pressure.Row_expiry_date, &well_test_pressure.Row_quality); err != nil {
			return err // Exit if we get an error
		}

		// Append well_test_pressure to result
		result = append(result, well_test_pressure)
	}
	// Return result in JSON format
	return c.JSON(result)
	 
}

func SetWellTestPressure(c *fiber.Ctx) error {
	var well_test_pressure dto.Well_test_pressure

	setDefaults(&well_test_pressure)

	if err := json.Unmarshal(c.Body(), &well_test_pressure); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	stmt, err := db.Prepare("insert into well_test_pressure values(:1, :2, :3, :4, :5, :6, :7, :8, :9, :10, :11, :12, :13, :14, :15, :16, :17, :18, :19, :20, :21, :22, :23, :24, :25, :26)")
	if err != nil {
		return err
	}
	well_test_pressure.Row_created_date = formatDate(well_test_pressure.Row_created_date)
	well_test_pressure.Row_changed_date = formatDate(well_test_pressure.Row_changed_date)
	well_test_pressure.Effective_date = formatDateString(well_test_pressure.Effective_date)
	well_test_pressure.Expiry_date = formatDateString(well_test_pressure.Expiry_date)
	well_test_pressure.Row_effective_date = formatDateString(well_test_pressure.Row_effective_date)
	well_test_pressure.Row_expiry_date = formatDateString(well_test_pressure.Row_expiry_date)






	rows, err := stmt.Exec(well_test_pressure.Uwi, well_test_pressure.Source, well_test_pressure.Test_type, well_test_pressure.Run_num, well_test_pressure.Test_num, well_test_pressure.Period_type, well_test_pressure.Period_obs_no, well_test_pressure.Active_ind, well_test_pressure.Effective_date, well_test_pressure.End_pressure, well_test_pressure.End_pressure_ouom, well_test_pressure.Expiry_date, well_test_pressure.Ppdm_guid, well_test_pressure.Quality_code, well_test_pressure.Recorder_id, well_test_pressure.Remark, well_test_pressure.Start_pressure, well_test_pressure.Start_pressure_ouom, well_test_pressure.Summary_ind, well_test_pressure.Row_changed_by, well_test_pressure.Row_changed_date, well_test_pressure.Row_created_by, well_test_pressure.Row_created_date, well_test_pressure.Row_effective_date, well_test_pressure.Row_expiry_date, well_test_pressure.Row_quality)
	
	if err != nil {
		return err
	}

	return c.JSON(rows)
}

func UpdateWellTestPressure(c *fiber.Ctx) error {
	var well_test_pressure dto.Well_test_pressure

	setDefaults(&well_test_pressure)

	if err := json.Unmarshal(c.Body(), &well_test_pressure); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	id := c.Params("id")
	well_test_pressure.Uwi = id

    if well_test_pressure.Row_created_date != nil {
        return c.Status(400).SendString("Cannot update row_created_date")
    }

	stmt, err := db.Prepare("update well_test_pressure set source = :1, test_type = :2, run_num = :3, test_num = :4, period_type = :5, period_obs_no = :6, active_ind = :7, effective_date = :8, end_pressure = :9, end_pressure_ouom = :10, expiry_date = :11, ppdm_guid = :12, quality_code = :13, recorder_id = :14, remark = :15, start_pressure = :16, start_pressure_ouom = :17, summary_ind = :18, row_changed_by = :19, row_changed_date = :20, row_created_by = :21, row_effective_date = :22, row_expiry_date = :23, row_quality = :24 where uwi = :26")
	if err != nil {
		return err
	}

	well_test_pressure.Row_changed_date = formatDate(well_test_pressure.Row_changed_date)
	well_test_pressure.Effective_date = formatDateString(well_test_pressure.Effective_date)
	well_test_pressure.Expiry_date = formatDateString(well_test_pressure.Expiry_date)
	well_test_pressure.Row_effective_date = formatDateString(well_test_pressure.Row_effective_date)
	well_test_pressure.Row_expiry_date = formatDateString(well_test_pressure.Row_expiry_date)






	rows, err := stmt.Exec(well_test_pressure.Source, well_test_pressure.Test_type, well_test_pressure.Run_num, well_test_pressure.Test_num, well_test_pressure.Period_type, well_test_pressure.Period_obs_no, well_test_pressure.Active_ind, well_test_pressure.Effective_date, well_test_pressure.End_pressure, well_test_pressure.End_pressure_ouom, well_test_pressure.Expiry_date, well_test_pressure.Ppdm_guid, well_test_pressure.Quality_code, well_test_pressure.Recorder_id, well_test_pressure.Remark, well_test_pressure.Start_pressure, well_test_pressure.Start_pressure_ouom, well_test_pressure.Summary_ind, well_test_pressure.Row_changed_by, well_test_pressure.Row_changed_date, well_test_pressure.Row_created_by, well_test_pressure.Row_effective_date, well_test_pressure.Row_expiry_date, well_test_pressure.Row_quality, well_test_pressure.Uwi)
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

func PatchWellTestPressure(c *fiber.Ctx) error {
	id := c.Params("id")

	updatedFields := make(map[string]interface{})

	if err := c.BodyParser(&updatedFields); err != nil {
		return err
	}
	if _, ok := updatedFields["row_created_date"]; ok {
        return c.Status(400).SendString("Cannot update row_created_date field")
    }

	query := "update well_test_pressure set "
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

func DeleteWellTestPressure(c *fiber.Ctx) error {
	id := c.Params("id")
	var well_test_pressure dto.Well_test_pressure
	well_test_pressure.Uwi = id

	stmt, err := db.Prepare("delete from well_test_pressure where uwi = :1")
	if err != nil {
		return err
	}

	rows, err := stmt.Exec(well_test_pressure.Uwi)
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


