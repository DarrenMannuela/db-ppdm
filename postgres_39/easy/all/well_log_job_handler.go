package apiv1

import (
	dto "github.com/DarrenMannuela/pt_gtn_bibliography/dto"
	"github.com/gofiber/fiber/v2"
	"strconv"
	"encoding/json"
)


func GetWellLogJob(c *fiber.Ctx) error {
	rows, err := db.Query("select * from well_log_job")
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}
	defer rows.Close()
	var result []dto.Well_log_job

	for rows.Next() {
		var well_log_job dto.Well_log_job
		if err := rows.Scan(&well_log_job.Uwi, &well_log_job.Source, &well_log_job.Job_id, &well_log_job.Active_ind, &well_log_job.Area_id, &well_log_job.Area_type, &well_log_job.Casing_shoe_depth, &well_log_job.Casing_shoe_depth_ouom, &well_log_job.Drilling_md, &well_log_job.Drilling_md_ouom, &well_log_job.Effective_date, &well_log_job.End_date, &well_log_job.Engineer, &well_log_job.Expiry_date, &well_log_job.Logging_company, &well_log_job.Logging_unit, &well_log_job.Observer, &well_log_job.Ppdm_guid, &well_log_job.Remark, &well_log_job.Start_date, &well_log_job.Row_changed_by, &well_log_job.Row_changed_date, &well_log_job.Row_created_by, &well_log_job.Row_created_date, &well_log_job.Row_effective_date, &well_log_job.Row_expiry_date, &well_log_job.Row_quality); err != nil {
			return err // Exit if we get an error
		}

		// Append well_log_job to result
		result = append(result, well_log_job)
	}
	// Return result in JSON format
	return c.JSON(result)
	 
}

func SetWellLogJob(c *fiber.Ctx) error {
	var well_log_job dto.Well_log_job

	setDefaults(&well_log_job)

	if err := json.Unmarshal(c.Body(), &well_log_job); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	stmt, err := db.Prepare("insert into well_log_job values(:1, :2, :3, :4, :5, :6, :7, :8, :9, :10, :11, :12, :13, :14, :15, :16, :17, :18, :19, :20, :21, :22, :23, :24, :25, :26, :27)")
	if err != nil {
		return err
	}
	well_log_job.Row_created_date = formatDate(well_log_job.Row_created_date)
	well_log_job.Row_changed_date = formatDate(well_log_job.Row_changed_date)
	well_log_job.Effective_date = formatDateString(well_log_job.Effective_date)
	well_log_job.End_date = formatDateString(well_log_job.End_date)
	well_log_job.Expiry_date = formatDateString(well_log_job.Expiry_date)
	well_log_job.Start_date = formatDateString(well_log_job.Start_date)
	well_log_job.Row_effective_date = formatDateString(well_log_job.Row_effective_date)
	well_log_job.Row_expiry_date = formatDateString(well_log_job.Row_expiry_date)






	rows, err := stmt.Exec(well_log_job.Uwi, well_log_job.Source, well_log_job.Job_id, well_log_job.Active_ind, well_log_job.Area_id, well_log_job.Area_type, well_log_job.Casing_shoe_depth, well_log_job.Casing_shoe_depth_ouom, well_log_job.Drilling_md, well_log_job.Drilling_md_ouom, well_log_job.Effective_date, well_log_job.End_date, well_log_job.Engineer, well_log_job.Expiry_date, well_log_job.Logging_company, well_log_job.Logging_unit, well_log_job.Observer, well_log_job.Ppdm_guid, well_log_job.Remark, well_log_job.Start_date, well_log_job.Row_changed_by, well_log_job.Row_changed_date, well_log_job.Row_created_by, well_log_job.Row_created_date, well_log_job.Row_effective_date, well_log_job.Row_expiry_date, well_log_job.Row_quality)
	
	if err != nil {
		return err
	}

	return c.JSON(rows)
}

func UpdateWellLogJob(c *fiber.Ctx) error {
	var well_log_job dto.Well_log_job

	setDefaults(&well_log_job)

	if err := json.Unmarshal(c.Body(), &well_log_job); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	id := c.Params("id")
	well_log_job.Uwi = id

    if well_log_job.Row_created_date != nil {
        return c.Status(400).SendString("Cannot update row_created_date")
    }

	stmt, err := db.Prepare("update well_log_job set source = :1, job_id = :2, active_ind = :3, area_id = :4, area_type = :5, casing_shoe_depth = :6, casing_shoe_depth_ouom = :7, drilling_md = :8, drilling_md_ouom = :9, effective_date = :10, end_date = :11, engineer = :12, expiry_date = :13, logging_company = :14, logging_unit = :15, observer = :16, ppdm_guid = :17, remark = :18, start_date = :19, row_changed_by = :20, row_changed_date = :21, row_created_by = :22, row_effective_date = :23, row_expiry_date = :24, row_quality = :25 where uwi = :27")
	if err != nil {
		return err
	}

	well_log_job.Row_changed_date = formatDate(well_log_job.Row_changed_date)
	well_log_job.Effective_date = formatDateString(well_log_job.Effective_date)
	well_log_job.End_date = formatDateString(well_log_job.End_date)
	well_log_job.Expiry_date = formatDateString(well_log_job.Expiry_date)
	well_log_job.Start_date = formatDateString(well_log_job.Start_date)
	well_log_job.Row_effective_date = formatDateString(well_log_job.Row_effective_date)
	well_log_job.Row_expiry_date = formatDateString(well_log_job.Row_expiry_date)






	rows, err := stmt.Exec(well_log_job.Source, well_log_job.Job_id, well_log_job.Active_ind, well_log_job.Area_id, well_log_job.Area_type, well_log_job.Casing_shoe_depth, well_log_job.Casing_shoe_depth_ouom, well_log_job.Drilling_md, well_log_job.Drilling_md_ouom, well_log_job.Effective_date, well_log_job.End_date, well_log_job.Engineer, well_log_job.Expiry_date, well_log_job.Logging_company, well_log_job.Logging_unit, well_log_job.Observer, well_log_job.Ppdm_guid, well_log_job.Remark, well_log_job.Start_date, well_log_job.Row_changed_by, well_log_job.Row_changed_date, well_log_job.Row_created_by, well_log_job.Row_effective_date, well_log_job.Row_expiry_date, well_log_job.Row_quality, well_log_job.Uwi)
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

func PatchWellLogJob(c *fiber.Ctx) error {
	id := c.Params("id")

	updatedFields := make(map[string]interface{})

	if err := c.BodyParser(&updatedFields); err != nil {
		return err
	}
	if _, ok := updatedFields["row_created_date"]; ok {
        return c.Status(400).SendString("Cannot update row_created_date field")
    }

	query := "update well_log_job set "
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
		} else if key == "effective_date" || key == "end_date" || key == "expiry_date" || key == "start_date" || key == "row_effective_date" || key == "row_expiry_date"      {
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

func DeleteWellLogJob(c *fiber.Ctx) error {
	id := c.Params("id")
	var well_log_job dto.Well_log_job
	well_log_job.Uwi = id

	stmt, err := db.Prepare("delete from well_log_job where uwi = :1")
	if err != nil {
		return err
	}

	rows, err := stmt.Exec(well_log_job.Uwi)
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


