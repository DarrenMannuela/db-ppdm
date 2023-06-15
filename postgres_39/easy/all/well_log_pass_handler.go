package apiv1

import (
	dto "github.com/DarrenMannuela/pt_gtn_bibliography/dto"
	"github.com/gofiber/fiber/v2"
	"strconv"
	"encoding/json"
)


func GetWellLogPass(c *fiber.Ctx) error {
	rows, err := db.Query("select * from well_log_pass")
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}
	defer rows.Close()
	var result []dto.Well_log_pass

	for rows.Next() {
		var well_log_pass dto.Well_log_pass
		if err := rows.Scan(&well_log_pass.Uwi, &well_log_pass.Source, &well_log_pass.Job_id, &well_log_pass.Trip_obs_no, &well_log_pass.Log_tool_pass_no, &well_log_pass.Active_ind, &well_log_pass.Base_depth, &well_log_pass.Base_depth_ouom, &well_log_pass.Base_log_ind, &well_log_pass.Effective_date, &well_log_pass.End_time, &well_log_pass.End_timezone, &well_log_pass.Expiry_date, &well_log_pass.Ppdm_guid, &well_log_pass.Remark, &well_log_pass.Start_time, &well_log_pass.Start_timezone, &well_log_pass.Top_depth, &well_log_pass.Top_depth_ouom, &well_log_pass.Row_changed_by, &well_log_pass.Row_changed_date, &well_log_pass.Row_created_by, &well_log_pass.Row_created_date, &well_log_pass.Row_effective_date, &well_log_pass.Row_expiry_date, &well_log_pass.Row_quality); err != nil {
			return err // Exit if we get an error
		}

		// Append well_log_pass to result
		result = append(result, well_log_pass)
	}
	// Return result in JSON format
	return c.JSON(result)
	 
}

func SetWellLogPass(c *fiber.Ctx) error {
	var well_log_pass dto.Well_log_pass

	setDefaults(&well_log_pass)

	if err := json.Unmarshal(c.Body(), &well_log_pass); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	stmt, err := db.Prepare("insert into well_log_pass values(:1, :2, :3, :4, :5, :6, :7, :8, :9, :10, :11, :12, :13, :14, :15, :16, :17, :18, :19, :20, :21, :22, :23, :24, :25, :26)")
	if err != nil {
		return err
	}
	well_log_pass.Row_created_date = formatDate(well_log_pass.Row_created_date)
	well_log_pass.Row_changed_date = formatDate(well_log_pass.Row_changed_date)
	well_log_pass.Effective_date = formatDateString(well_log_pass.Effective_date)
	well_log_pass.End_time = formatDateString(well_log_pass.End_time)
	well_log_pass.Expiry_date = formatDateString(well_log_pass.Expiry_date)
	well_log_pass.Start_time = formatDateString(well_log_pass.Start_time)
	well_log_pass.Row_effective_date = formatDateString(well_log_pass.Row_effective_date)
	well_log_pass.Row_expiry_date = formatDateString(well_log_pass.Row_expiry_date)






	rows, err := stmt.Exec(well_log_pass.Uwi, well_log_pass.Source, well_log_pass.Job_id, well_log_pass.Trip_obs_no, well_log_pass.Log_tool_pass_no, well_log_pass.Active_ind, well_log_pass.Base_depth, well_log_pass.Base_depth_ouom, well_log_pass.Base_log_ind, well_log_pass.Effective_date, well_log_pass.End_time, well_log_pass.End_timezone, well_log_pass.Expiry_date, well_log_pass.Ppdm_guid, well_log_pass.Remark, well_log_pass.Start_time, well_log_pass.Start_timezone, well_log_pass.Top_depth, well_log_pass.Top_depth_ouom, well_log_pass.Row_changed_by, well_log_pass.Row_changed_date, well_log_pass.Row_created_by, well_log_pass.Row_created_date, well_log_pass.Row_effective_date, well_log_pass.Row_expiry_date, well_log_pass.Row_quality)
	
	if err != nil {
		return err
	}

	return c.JSON(rows)
}

func UpdateWellLogPass(c *fiber.Ctx) error {
	var well_log_pass dto.Well_log_pass

	setDefaults(&well_log_pass)

	if err := json.Unmarshal(c.Body(), &well_log_pass); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	id := c.Params("id")
	well_log_pass.Uwi = id

    if well_log_pass.Row_created_date != nil {
        return c.Status(400).SendString("Cannot update row_created_date")
    }

	stmt, err := db.Prepare("update well_log_pass set source = :1, job_id = :2, trip_obs_no = :3, log_tool_pass_no = :4, active_ind = :5, base_depth = :6, base_depth_ouom = :7, base_log_ind = :8, effective_date = :9, end_time = :10, end_timezone = :11, expiry_date = :12, ppdm_guid = :13, remark = :14, start_time = :15, start_timezone = :16, top_depth = :17, top_depth_ouom = :18, row_changed_by = :19, row_changed_date = :20, row_created_by = :21, row_effective_date = :22, row_expiry_date = :23, row_quality = :24 where uwi = :26")
	if err != nil {
		return err
	}

	well_log_pass.Row_changed_date = formatDate(well_log_pass.Row_changed_date)
	well_log_pass.Effective_date = formatDateString(well_log_pass.Effective_date)
	well_log_pass.End_time = formatDateString(well_log_pass.End_time)
	well_log_pass.Expiry_date = formatDateString(well_log_pass.Expiry_date)
	well_log_pass.Start_time = formatDateString(well_log_pass.Start_time)
	well_log_pass.Row_effective_date = formatDateString(well_log_pass.Row_effective_date)
	well_log_pass.Row_expiry_date = formatDateString(well_log_pass.Row_expiry_date)






	rows, err := stmt.Exec(well_log_pass.Source, well_log_pass.Job_id, well_log_pass.Trip_obs_no, well_log_pass.Log_tool_pass_no, well_log_pass.Active_ind, well_log_pass.Base_depth, well_log_pass.Base_depth_ouom, well_log_pass.Base_log_ind, well_log_pass.Effective_date, well_log_pass.End_time, well_log_pass.End_timezone, well_log_pass.Expiry_date, well_log_pass.Ppdm_guid, well_log_pass.Remark, well_log_pass.Start_time, well_log_pass.Start_timezone, well_log_pass.Top_depth, well_log_pass.Top_depth_ouom, well_log_pass.Row_changed_by, well_log_pass.Row_changed_date, well_log_pass.Row_created_by, well_log_pass.Row_effective_date, well_log_pass.Row_expiry_date, well_log_pass.Row_quality, well_log_pass.Uwi)
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

func PatchWellLogPass(c *fiber.Ctx) error {
	id := c.Params("id")

	updatedFields := make(map[string]interface{})

	if err := c.BodyParser(&updatedFields); err != nil {
		return err
	}
	if _, ok := updatedFields["row_created_date"]; ok {
        return c.Status(400).SendString("Cannot update row_created_date field")
    }

	query := "update well_log_pass set "
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
		} else if key == "effective_date" || key == "end_time" || key == "expiry_date" || key == "start_time" || key == "row_effective_date" || key == "row_expiry_date"      {
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

func DeleteWellLogPass(c *fiber.Ctx) error {
	id := c.Params("id")
	var well_log_pass dto.Well_log_pass
	well_log_pass.Uwi = id

	stmt, err := db.Prepare("delete from well_log_pass where uwi = :1")
	if err != nil {
		return err
	}

	rows, err := stmt.Exec(well_log_pass.Uwi)
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


