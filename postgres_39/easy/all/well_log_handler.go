package apiv1

import (
	dto "github.com/DarrenMannuela/pt_gtn_bibliography/dto"
	"github.com/gofiber/fiber/v2"
	"strconv"
	"encoding/json"
)


func GetWellLog(c *fiber.Ctx) error {
	rows, err := db.Query("select * from well_log")
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}
	defer rows.Close()
	var result []dto.Well_log

	for rows.Next() {
		var well_log dto.Well_log
		if err := rows.Scan(&well_log.Uwi, &well_log.Well_log_id, &well_log.Source, &well_log.Acquired_for_ba_id, &well_log.Active_ind, &well_log.Base_depth, &well_log.Base_depth_ouom, &well_log.Bypass_ind, &well_log.Cased_hole_ind, &well_log.Composite_ind, &well_log.Depth_type, &well_log.Dictionary_id, &well_log.Effective_date, &well_log.Expiry_date, &well_log.Log_job_id, &well_log.Log_job_source, &well_log.Log_ref_num, &well_log.Log_title, &well_log.Log_tool_pass_no, &well_log.Mwd_ind, &well_log.Ppdm_guid, &well_log.Remark, &well_log.Top_depth, &well_log.Top_depth_ouom, &well_log.Trip_obs_no, &well_log.Row_changed_by, &well_log.Row_changed_date, &well_log.Row_created_by, &well_log.Row_created_date, &well_log.Row_effective_date, &well_log.Row_expiry_date, &well_log.Row_quality); err != nil {
			return err // Exit if we get an error
		}

		// Append well_log to result
		result = append(result, well_log)
	}
	// Return result in JSON format
	return c.JSON(result)
	 
}

func SetWellLog(c *fiber.Ctx) error {
	var well_log dto.Well_log

	setDefaults(&well_log)

	if err := json.Unmarshal(c.Body(), &well_log); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	stmt, err := db.Prepare("insert into well_log values(:1, :2, :3, :4, :5, :6, :7, :8, :9, :10, :11, :12, :13, :14, :15, :16, :17, :18, :19, :20, :21, :22, :23, :24, :25, :26, :27, :28, :29, :30, :31, :32)")
	if err != nil {
		return err
	}
	well_log.Row_created_date = formatDate(well_log.Row_created_date)
	well_log.Row_changed_date = formatDate(well_log.Row_changed_date)
	well_log.Effective_date = formatDateString(well_log.Effective_date)
	well_log.Expiry_date = formatDateString(well_log.Expiry_date)
	well_log.Row_effective_date = formatDateString(well_log.Row_effective_date)
	well_log.Row_expiry_date = formatDateString(well_log.Row_expiry_date)






	rows, err := stmt.Exec(well_log.Uwi, well_log.Well_log_id, well_log.Source, well_log.Acquired_for_ba_id, well_log.Active_ind, well_log.Base_depth, well_log.Base_depth_ouom, well_log.Bypass_ind, well_log.Cased_hole_ind, well_log.Composite_ind, well_log.Depth_type, well_log.Dictionary_id, well_log.Effective_date, well_log.Expiry_date, well_log.Log_job_id, well_log.Log_job_source, well_log.Log_ref_num, well_log.Log_title, well_log.Log_tool_pass_no, well_log.Mwd_ind, well_log.Ppdm_guid, well_log.Remark, well_log.Top_depth, well_log.Top_depth_ouom, well_log.Trip_obs_no, well_log.Row_changed_by, well_log.Row_changed_date, well_log.Row_created_by, well_log.Row_created_date, well_log.Row_effective_date, well_log.Row_expiry_date, well_log.Row_quality)
	
	if err != nil {
		return err
	}

	return c.JSON(rows)
}

func UpdateWellLog(c *fiber.Ctx) error {
	var well_log dto.Well_log

	setDefaults(&well_log)

	if err := json.Unmarshal(c.Body(), &well_log); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	id := c.Params("id")
	well_log.Uwi = id

    if well_log.Row_created_date != nil {
        return c.Status(400).SendString("Cannot update row_created_date")
    }

	stmt, err := db.Prepare("update well_log set well_log_id = :1, source = :2, acquired_for_ba_id = :3, active_ind = :4, base_depth = :5, base_depth_ouom = :6, bypass_ind = :7, cased_hole_ind = :8, composite_ind = :9, depth_type = :10, dictionary_id = :11, effective_date = :12, expiry_date = :13, log_job_id = :14, log_job_source = :15, log_ref_num = :16, log_title = :17, log_tool_pass_no = :18, mwd_ind = :19, ppdm_guid = :20, remark = :21, top_depth = :22, top_depth_ouom = :23, trip_obs_no = :24, row_changed_by = :25, row_changed_date = :26, row_created_by = :27, row_effective_date = :28, row_expiry_date = :29, row_quality = :30 where uwi = :32")
	if err != nil {
		return err
	}

	well_log.Row_changed_date = formatDate(well_log.Row_changed_date)
	well_log.Effective_date = formatDateString(well_log.Effective_date)
	well_log.Expiry_date = formatDateString(well_log.Expiry_date)
	well_log.Row_effective_date = formatDateString(well_log.Row_effective_date)
	well_log.Row_expiry_date = formatDateString(well_log.Row_expiry_date)






	rows, err := stmt.Exec(well_log.Well_log_id, well_log.Source, well_log.Acquired_for_ba_id, well_log.Active_ind, well_log.Base_depth, well_log.Base_depth_ouom, well_log.Bypass_ind, well_log.Cased_hole_ind, well_log.Composite_ind, well_log.Depth_type, well_log.Dictionary_id, well_log.Effective_date, well_log.Expiry_date, well_log.Log_job_id, well_log.Log_job_source, well_log.Log_ref_num, well_log.Log_title, well_log.Log_tool_pass_no, well_log.Mwd_ind, well_log.Ppdm_guid, well_log.Remark, well_log.Top_depth, well_log.Top_depth_ouom, well_log.Trip_obs_no, well_log.Row_changed_by, well_log.Row_changed_date, well_log.Row_created_by, well_log.Row_effective_date, well_log.Row_expiry_date, well_log.Row_quality, well_log.Uwi)
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

func PatchWellLog(c *fiber.Ctx) error {
	id := c.Params("id")

	updatedFields := make(map[string]interface{})

	if err := c.BodyParser(&updatedFields); err != nil {
		return err
	}
	if _, ok := updatedFields["row_created_date"]; ok {
        return c.Status(400).SendString("Cannot update row_created_date field")
    }

	query := "update well_log set "
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

func DeleteWellLog(c *fiber.Ctx) error {
	id := c.Params("id")
	var well_log dto.Well_log
	well_log.Uwi = id

	stmt, err := db.Prepare("delete from well_log where uwi = :1")
	if err != nil {
		return err
	}

	rows, err := stmt.Exec(well_log.Uwi)
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


