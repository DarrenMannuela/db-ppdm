package apiv1

import (
	dto "github.com/DarrenMannuela/pt_gtn_bibliography/dto"
	"github.com/gofiber/fiber/v2"
	"strconv"
	"encoding/json"
)


func GetProjectStepTime(c *fiber.Ctx) error {
	rows, err := db.Query("select * from project_step_time")
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}
	defer rows.Close()
	var result []dto.Project_step_time

	for rows.Next() {
		var project_step_time dto.Project_step_time
		if err := rows.Scan(&project_step_time.Project_id, &project_step_time.Step_id, &project_step_time.Time_obs_no, &project_step_time.Active_ind, &project_step_time.Business_associate_id, &project_step_time.Effective_date, &project_step_time.End_date, &project_step_time.End_time, &project_step_time.End_timezone, &project_step_time.Expiry_date, &project_step_time.Plan_ind, &project_step_time.Ppdm_guid, &project_step_time.Remark, &project_step_time.Source, &project_step_time.Start_date, &project_step_time.Start_time, &project_step_time.Start_timezone, &project_step_time.Total_time_elapsed, &project_step_time.Total_time_elapsed_ouom, &project_step_time.Total_time_elapsed_uom, &project_step_time.Row_changed_by, &project_step_time.Row_changed_date, &project_step_time.Row_created_by, &project_step_time.Row_created_date, &project_step_time.Row_effective_date, &project_step_time.Row_expiry_date, &project_step_time.Row_quality); err != nil {
			return err // Exit if we get an error
		}

		// Append project_step_time to result
		result = append(result, project_step_time)
	}
	// Return result in JSON format
	return c.JSON(result)
	 
}

func SetProjectStepTime(c *fiber.Ctx) error {
	var project_step_time dto.Project_step_time

	setDefaults(&project_step_time)

	if err := json.Unmarshal(c.Body(), &project_step_time); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	stmt, err := db.Prepare("insert into project_step_time values(:1, :2, :3, :4, :5, :6, :7, :8, :9, :10, :11, :12, :13, :14, :15, :16, :17, :18, :19, :20, :21, :22, :23, :24, :25, :26, :27)")
	if err != nil {
		return err
	}
	project_step_time.Row_created_date = formatDate(project_step_time.Row_created_date)
	project_step_time.Row_changed_date = formatDate(project_step_time.Row_changed_date)
	project_step_time.Effective_date = formatDateString(project_step_time.Effective_date)
	project_step_time.End_date = formatDateString(project_step_time.End_date)
	project_step_time.End_time = formatDateString(project_step_time.End_time)
	project_step_time.Expiry_date = formatDateString(project_step_time.Expiry_date)
	project_step_time.Start_date = formatDateString(project_step_time.Start_date)
	project_step_time.Start_time = formatDateString(project_step_time.Start_time)
	project_step_time.Row_effective_date = formatDateString(project_step_time.Row_effective_date)
	project_step_time.Row_expiry_date = formatDateString(project_step_time.Row_expiry_date)






	rows, err := stmt.Exec(project_step_time.Project_id, project_step_time.Step_id, project_step_time.Time_obs_no, project_step_time.Active_ind, project_step_time.Business_associate_id, project_step_time.Effective_date, project_step_time.End_date, project_step_time.End_time, project_step_time.End_timezone, project_step_time.Expiry_date, project_step_time.Plan_ind, project_step_time.Ppdm_guid, project_step_time.Remark, project_step_time.Source, project_step_time.Start_date, project_step_time.Start_time, project_step_time.Start_timezone, project_step_time.Total_time_elapsed, project_step_time.Total_time_elapsed_ouom, project_step_time.Total_time_elapsed_uom, project_step_time.Row_changed_by, project_step_time.Row_changed_date, project_step_time.Row_created_by, project_step_time.Row_created_date, project_step_time.Row_effective_date, project_step_time.Row_expiry_date, project_step_time.Row_quality)
	
	if err != nil {
		return err
	}

	return c.JSON(rows)
}

func UpdateProjectStepTime(c *fiber.Ctx) error {
	var project_step_time dto.Project_step_time

	setDefaults(&project_step_time)

	if err := json.Unmarshal(c.Body(), &project_step_time); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	id := c.Params("id")
	project_step_time.Project_id = id

    if project_step_time.Row_created_date != nil {
        return c.Status(400).SendString("Cannot update row_created_date")
    }

	stmt, err := db.Prepare("update project_step_time set step_id = :1, time_obs_no = :2, active_ind = :3, business_associate_id = :4, effective_date = :5, end_date = :6, end_time = :7, end_timezone = :8, expiry_date = :9, plan_ind = :10, ppdm_guid = :11, remark = :12, source = :13, start_date = :14, start_time = :15, start_timezone = :16, total_time_elapsed = :17, total_time_elapsed_ouom = :18, total_time_elapsed_uom = :19, row_changed_by = :20, row_changed_date = :21, row_created_by = :22, row_effective_date = :23, row_expiry_date = :24, row_quality = :25 where project_id = :27")
	if err != nil {
		return err
	}

	project_step_time.Row_changed_date = formatDate(project_step_time.Row_changed_date)
	project_step_time.Effective_date = formatDateString(project_step_time.Effective_date)
	project_step_time.End_date = formatDateString(project_step_time.End_date)
	project_step_time.End_time = formatDateString(project_step_time.End_time)
	project_step_time.Expiry_date = formatDateString(project_step_time.Expiry_date)
	project_step_time.Start_date = formatDateString(project_step_time.Start_date)
	project_step_time.Start_time = formatDateString(project_step_time.Start_time)
	project_step_time.Row_effective_date = formatDateString(project_step_time.Row_effective_date)
	project_step_time.Row_expiry_date = formatDateString(project_step_time.Row_expiry_date)






	rows, err := stmt.Exec(project_step_time.Step_id, project_step_time.Time_obs_no, project_step_time.Active_ind, project_step_time.Business_associate_id, project_step_time.Effective_date, project_step_time.End_date, project_step_time.End_time, project_step_time.End_timezone, project_step_time.Expiry_date, project_step_time.Plan_ind, project_step_time.Ppdm_guid, project_step_time.Remark, project_step_time.Source, project_step_time.Start_date, project_step_time.Start_time, project_step_time.Start_timezone, project_step_time.Total_time_elapsed, project_step_time.Total_time_elapsed_ouom, project_step_time.Total_time_elapsed_uom, project_step_time.Row_changed_by, project_step_time.Row_changed_date, project_step_time.Row_created_by, project_step_time.Row_effective_date, project_step_time.Row_expiry_date, project_step_time.Row_quality, project_step_time.Project_id)
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

func PatchProjectStepTime(c *fiber.Ctx) error {
	id := c.Params("id")

	updatedFields := make(map[string]interface{})

	if err := c.BodyParser(&updatedFields); err != nil {
		return err
	}
	if _, ok := updatedFields["row_created_date"]; ok {
        return c.Status(400).SendString("Cannot update row_created_date field")
    }

	query := "update project_step_time set "
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
		} else if key == "effective_date" || key == "end_date" || key == "end_time" || key == "expiry_date" || key == "start_date" || key == "start_time" || key == "row_effective_date" || key == "row_expiry_date"      {
			if date, ok := value.(string); ok && len(date) > 0 {
				formattedDate := formatDateString(&date)
				value = formattedDate
			}
		}
		values = append(values, value)
	}
	query += " where project_id = :id"

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

func DeleteProjectStepTime(c *fiber.Ctx) error {
	id := c.Params("id")
	var project_step_time dto.Project_step_time
	project_step_time.Project_id = id

	stmt, err := db.Prepare("delete from project_step_time where project_id = :1")
	if err != nil {
		return err
	}

	rows, err := stmt.Exec(project_step_time.Project_id)
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


