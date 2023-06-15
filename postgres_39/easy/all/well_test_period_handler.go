package apiv1

import (
	dto "github.com/DarrenMannuela/pt_gtn_bibliography/dto"
	"github.com/gofiber/fiber/v2"
	"strconv"
	"encoding/json"
)


func GetWellTestPeriod(c *fiber.Ctx) error {
	rows, err := db.Query("select * from well_test_period")
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}
	defer rows.Close()
	var result []dto.Well_test_period

	for rows.Next() {
		var well_test_period dto.Well_test_period
		if err := rows.Scan(&well_test_period.Uwi, &well_test_period.Source, &well_test_period.Test_type, &well_test_period.Run_num, &well_test_period.Test_num, &well_test_period.Period_type, &well_test_period.Period_obs_no, &well_test_period.Active_ind, &well_test_period.Casing_pressure, &well_test_period.Casing_pressure_ouom, &well_test_period.Effective_date, &well_test_period.Expiry_date, &well_test_period.Period_duration, &well_test_period.Period_duration_ouom, &well_test_period.Ppdm_guid, &well_test_period.Remark, &well_test_period.Tubing_pressure, &well_test_period.Tubing_pressure_ouom, &well_test_period.Row_changed_by, &well_test_period.Row_changed_date, &well_test_period.Row_created_by, &well_test_period.Row_created_date, &well_test_period.Row_effective_date, &well_test_period.Row_expiry_date, &well_test_period.Row_quality); err != nil {
			return err // Exit if we get an error
		}

		// Append well_test_period to result
		result = append(result, well_test_period)
	}
	// Return result in JSON format
	return c.JSON(result)
	 
}

func SetWellTestPeriod(c *fiber.Ctx) error {
	var well_test_period dto.Well_test_period

	setDefaults(&well_test_period)

	if err := json.Unmarshal(c.Body(), &well_test_period); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	stmt, err := db.Prepare("insert into well_test_period values(:1, :2, :3, :4, :5, :6, :7, :8, :9, :10, :11, :12, :13, :14, :15, :16, :17, :18, :19, :20, :21, :22, :23, :24, :25)")
	if err != nil {
		return err
	}
	well_test_period.Row_created_date = formatDate(well_test_period.Row_created_date)
	well_test_period.Row_changed_date = formatDate(well_test_period.Row_changed_date)
	well_test_period.Effective_date = formatDateString(well_test_period.Effective_date)
	well_test_period.Expiry_date = formatDateString(well_test_period.Expiry_date)
	well_test_period.Row_effective_date = formatDateString(well_test_period.Row_effective_date)
	well_test_period.Row_expiry_date = formatDateString(well_test_period.Row_expiry_date)






	rows, err := stmt.Exec(well_test_period.Uwi, well_test_period.Source, well_test_period.Test_type, well_test_period.Run_num, well_test_period.Test_num, well_test_period.Period_type, well_test_period.Period_obs_no, well_test_period.Active_ind, well_test_period.Casing_pressure, well_test_period.Casing_pressure_ouom, well_test_period.Effective_date, well_test_period.Expiry_date, well_test_period.Period_duration, well_test_period.Period_duration_ouom, well_test_period.Ppdm_guid, well_test_period.Remark, well_test_period.Tubing_pressure, well_test_period.Tubing_pressure_ouom, well_test_period.Row_changed_by, well_test_period.Row_changed_date, well_test_period.Row_created_by, well_test_period.Row_created_date, well_test_period.Row_effective_date, well_test_period.Row_expiry_date, well_test_period.Row_quality)
	
	if err != nil {
		return err
	}

	return c.JSON(rows)
}

func UpdateWellTestPeriod(c *fiber.Ctx) error {
	var well_test_period dto.Well_test_period

	setDefaults(&well_test_period)

	if err := json.Unmarshal(c.Body(), &well_test_period); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	id := c.Params("id")
	well_test_period.Uwi = id

    if well_test_period.Row_created_date != nil {
        return c.Status(400).SendString("Cannot update row_created_date")
    }

	stmt, err := db.Prepare("update well_test_period set source = :1, test_type = :2, run_num = :3, test_num = :4, period_type = :5, period_obs_no = :6, active_ind = :7, casing_pressure = :8, casing_pressure_ouom = :9, effective_date = :10, expiry_date = :11, period_duration = :12, period_duration_ouom = :13, ppdm_guid = :14, remark = :15, tubing_pressure = :16, tubing_pressure_ouom = :17, row_changed_by = :18, row_changed_date = :19, row_created_by = :20, row_effective_date = :21, row_expiry_date = :22, row_quality = :23 where uwi = :25")
	if err != nil {
		return err
	}

	well_test_period.Row_changed_date = formatDate(well_test_period.Row_changed_date)
	well_test_period.Effective_date = formatDateString(well_test_period.Effective_date)
	well_test_period.Expiry_date = formatDateString(well_test_period.Expiry_date)
	well_test_period.Row_effective_date = formatDateString(well_test_period.Row_effective_date)
	well_test_period.Row_expiry_date = formatDateString(well_test_period.Row_expiry_date)






	rows, err := stmt.Exec(well_test_period.Source, well_test_period.Test_type, well_test_period.Run_num, well_test_period.Test_num, well_test_period.Period_type, well_test_period.Period_obs_no, well_test_period.Active_ind, well_test_period.Casing_pressure, well_test_period.Casing_pressure_ouom, well_test_period.Effective_date, well_test_period.Expiry_date, well_test_period.Period_duration, well_test_period.Period_duration_ouom, well_test_period.Ppdm_guid, well_test_period.Remark, well_test_period.Tubing_pressure, well_test_period.Tubing_pressure_ouom, well_test_period.Row_changed_by, well_test_period.Row_changed_date, well_test_period.Row_created_by, well_test_period.Row_effective_date, well_test_period.Row_expiry_date, well_test_period.Row_quality, well_test_period.Uwi)
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

func PatchWellTestPeriod(c *fiber.Ctx) error {
	id := c.Params("id")

	updatedFields := make(map[string]interface{})

	if err := c.BodyParser(&updatedFields); err != nil {
		return err
	}
	if _, ok := updatedFields["row_created_date"]; ok {
        return c.Status(400).SendString("Cannot update row_created_date field")
    }

	query := "update well_test_period set "
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

func DeleteWellTestPeriod(c *fiber.Ctx) error {
	id := c.Params("id")
	var well_test_period dto.Well_test_period
	well_test_period.Uwi = id

	stmt, err := db.Prepare("delete from well_test_period where uwi = :1")
	if err != nil {
		return err
	}

	rows, err := stmt.Exec(well_test_period.Uwi)
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


