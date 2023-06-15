package apiv1

import (
	dto "github.com/DarrenMannuela/pt_gtn_bibliography/dto"
	"github.com/gofiber/fiber/v2"
	"strconv"
	"encoding/json"
)


func GetWellTestRecorder(c *fiber.Ctx) error {
	rows, err := db.Query("select * from well_test_recorder")
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}
	defer rows.Close()
	var result []dto.Well_test_recorder

	for rows.Next() {
		var well_test_recorder dto.Well_test_recorder
		if err := rows.Scan(&well_test_recorder.Uwi, &well_test_recorder.Source, &well_test_recorder.Test_type, &well_test_recorder.Run_num, &well_test_recorder.Test_num, &well_test_recorder.Recorder_id, &well_test_recorder.Active_ind, &well_test_recorder.Effective_date, &well_test_recorder.Expiry_date, &well_test_recorder.Max_capacity_pressure, &well_test_recorder.Max_capacity_pressure_ouom, &well_test_recorder.Max_capacity_temp, &well_test_recorder.Max_capacity_temp_ouom, &well_test_recorder.Performance_quality, &well_test_recorder.Ppdm_guid, &well_test_recorder.Recorder_depth, &well_test_recorder.Recorder_depth_ouom, &well_test_recorder.Recorder_inside_ind, &well_test_recorder.Recorder_max_temp, &well_test_recorder.Recorder_max_temp_ouom, &well_test_recorder.Recorder_min_temp, &well_test_recorder.Recorder_min_temp_ouom, &well_test_recorder.Recorder_position, &well_test_recorder.Recorder_resolution, &well_test_recorder.Recorder_resolution_ouom, &well_test_recorder.Recorder_type, &well_test_recorder.Recorder_used_ind, &well_test_recorder.Remark, &well_test_recorder.Row_changed_by, &well_test_recorder.Row_changed_date, &well_test_recorder.Row_created_by, &well_test_recorder.Row_created_date, &well_test_recorder.Row_effective_date, &well_test_recorder.Row_expiry_date, &well_test_recorder.Row_quality); err != nil {
			return err // Exit if we get an error
		}

		// Append well_test_recorder to result
		result = append(result, well_test_recorder)
	}
	// Return result in JSON format
	return c.JSON(result)
	 
}

func SetWellTestRecorder(c *fiber.Ctx) error {
	var well_test_recorder dto.Well_test_recorder

	setDefaults(&well_test_recorder)

	if err := json.Unmarshal(c.Body(), &well_test_recorder); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	stmt, err := db.Prepare("insert into well_test_recorder values(:1, :2, :3, :4, :5, :6, :7, :8, :9, :10, :11, :12, :13, :14, :15, :16, :17, :18, :19, :20, :21, :22, :23, :24, :25, :26, :27, :28, :29, :30, :31, :32, :33, :34, :35)")
	if err != nil {
		return err
	}
	well_test_recorder.Row_created_date = formatDate(well_test_recorder.Row_created_date)
	well_test_recorder.Row_changed_date = formatDate(well_test_recorder.Row_changed_date)
	well_test_recorder.Effective_date = formatDateString(well_test_recorder.Effective_date)
	well_test_recorder.Expiry_date = formatDateString(well_test_recorder.Expiry_date)
	well_test_recorder.Row_effective_date = formatDateString(well_test_recorder.Row_effective_date)
	well_test_recorder.Row_expiry_date = formatDateString(well_test_recorder.Row_expiry_date)






	rows, err := stmt.Exec(well_test_recorder.Uwi, well_test_recorder.Source, well_test_recorder.Test_type, well_test_recorder.Run_num, well_test_recorder.Test_num, well_test_recorder.Recorder_id, well_test_recorder.Active_ind, well_test_recorder.Effective_date, well_test_recorder.Expiry_date, well_test_recorder.Max_capacity_pressure, well_test_recorder.Max_capacity_pressure_ouom, well_test_recorder.Max_capacity_temp, well_test_recorder.Max_capacity_temp_ouom, well_test_recorder.Performance_quality, well_test_recorder.Ppdm_guid, well_test_recorder.Recorder_depth, well_test_recorder.Recorder_depth_ouom, well_test_recorder.Recorder_inside_ind, well_test_recorder.Recorder_max_temp, well_test_recorder.Recorder_max_temp_ouom, well_test_recorder.Recorder_min_temp, well_test_recorder.Recorder_min_temp_ouom, well_test_recorder.Recorder_position, well_test_recorder.Recorder_resolution, well_test_recorder.Recorder_resolution_ouom, well_test_recorder.Recorder_type, well_test_recorder.Recorder_used_ind, well_test_recorder.Remark, well_test_recorder.Row_changed_by, well_test_recorder.Row_changed_date, well_test_recorder.Row_created_by, well_test_recorder.Row_created_date, well_test_recorder.Row_effective_date, well_test_recorder.Row_expiry_date, well_test_recorder.Row_quality)
	
	if err != nil {
		return err
	}

	return c.JSON(rows)
}

func UpdateWellTestRecorder(c *fiber.Ctx) error {
	var well_test_recorder dto.Well_test_recorder

	setDefaults(&well_test_recorder)

	if err := json.Unmarshal(c.Body(), &well_test_recorder); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	id := c.Params("id")
	well_test_recorder.Uwi = id

    if well_test_recorder.Row_created_date != nil {
        return c.Status(400).SendString("Cannot update row_created_date")
    }

	stmt, err := db.Prepare("update well_test_recorder set source = :1, test_type = :2, run_num = :3, test_num = :4, recorder_id = :5, active_ind = :6, effective_date = :7, expiry_date = :8, max_capacity_pressure = :9, max_capacity_pressure_ouom = :10, max_capacity_temp = :11, max_capacity_temp_ouom = :12, performance_quality = :13, ppdm_guid = :14, recorder_depth = :15, recorder_depth_ouom = :16, recorder_inside_ind = :17, recorder_max_temp = :18, recorder_max_temp_ouom = :19, recorder_min_temp = :20, recorder_min_temp_ouom = :21, recorder_position = :22, recorder_resolution = :23, recorder_resolution_ouom = :24, recorder_type = :25, recorder_used_ind = :26, remark = :27, row_changed_by = :28, row_changed_date = :29, row_created_by = :30, row_effective_date = :31, row_expiry_date = :32, row_quality = :33 where uwi = :35")
	if err != nil {
		return err
	}

	well_test_recorder.Row_changed_date = formatDate(well_test_recorder.Row_changed_date)
	well_test_recorder.Effective_date = formatDateString(well_test_recorder.Effective_date)
	well_test_recorder.Expiry_date = formatDateString(well_test_recorder.Expiry_date)
	well_test_recorder.Row_effective_date = formatDateString(well_test_recorder.Row_effective_date)
	well_test_recorder.Row_expiry_date = formatDateString(well_test_recorder.Row_expiry_date)






	rows, err := stmt.Exec(well_test_recorder.Source, well_test_recorder.Test_type, well_test_recorder.Run_num, well_test_recorder.Test_num, well_test_recorder.Recorder_id, well_test_recorder.Active_ind, well_test_recorder.Effective_date, well_test_recorder.Expiry_date, well_test_recorder.Max_capacity_pressure, well_test_recorder.Max_capacity_pressure_ouom, well_test_recorder.Max_capacity_temp, well_test_recorder.Max_capacity_temp_ouom, well_test_recorder.Performance_quality, well_test_recorder.Ppdm_guid, well_test_recorder.Recorder_depth, well_test_recorder.Recorder_depth_ouom, well_test_recorder.Recorder_inside_ind, well_test_recorder.Recorder_max_temp, well_test_recorder.Recorder_max_temp_ouom, well_test_recorder.Recorder_min_temp, well_test_recorder.Recorder_min_temp_ouom, well_test_recorder.Recorder_position, well_test_recorder.Recorder_resolution, well_test_recorder.Recorder_resolution_ouom, well_test_recorder.Recorder_type, well_test_recorder.Recorder_used_ind, well_test_recorder.Remark, well_test_recorder.Row_changed_by, well_test_recorder.Row_changed_date, well_test_recorder.Row_created_by, well_test_recorder.Row_effective_date, well_test_recorder.Row_expiry_date, well_test_recorder.Row_quality, well_test_recorder.Uwi)
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

func PatchWellTestRecorder(c *fiber.Ctx) error {
	id := c.Params("id")

	updatedFields := make(map[string]interface{})

	if err := c.BodyParser(&updatedFields); err != nil {
		return err
	}
	if _, ok := updatedFields["row_created_date"]; ok {
        return c.Status(400).SendString("Cannot update row_created_date field")
    }

	query := "update well_test_recorder set "
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

func DeleteWellTestRecorder(c *fiber.Ctx) error {
	id := c.Params("id")
	var well_test_recorder dto.Well_test_recorder
	well_test_recorder.Uwi = id

	stmt, err := db.Prepare("delete from well_test_recorder where uwi = :1")
	if err != nil {
		return err
	}

	rows, err := stmt.Exec(well_test_recorder.Uwi)
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


