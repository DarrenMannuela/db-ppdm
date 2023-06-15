package apiv1

import (
	dto "github.com/DarrenMannuela/pt_gtn_bibliography/dto"
	"github.com/gofiber/fiber/v2"
	"strconv"
	"encoding/json"
)


func GetProjectStep(c *fiber.Ctx) error {
	rows, err := db.Query("select * from project_step")
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}
	defer rows.Close()
	var result []dto.Project_step

	for rows.Next() {
		var project_step dto.Project_step
		if err := rows.Scan(&project_step.Project_id, &project_step.Step_id, &project_step.Active_ind, &project_step.Actual_end_date, &project_step.Actual_start_date, &project_step.Actual_time_elapsed, &project_step.Actual_time_elapsed_ouom, &project_step.Actual_time_elapsed_uom, &project_step.Critical_date, &project_step.Description, &project_step.Due_date, &project_step.Effective_date, &project_step.Estimated_time_elapsed, &project_step.Estimated_time_elapsed_ouom, &project_step.Estimated_time_elapsed_uom, &project_step.Expiry_date, &project_step.Plan_end_date, &project_step.Plan_start_date, &project_step.Plan_step_id, &project_step.Ppdm_guid, &project_step.Project_plan_id, &project_step.Remark, &project_step.Source, &project_step.Step_name, &project_step.Step_seq_no, &project_step.Step_type, &project_step.Where_completed, &project_step.Row_changed_by, &project_step.Row_changed_date, &project_step.Row_created_by, &project_step.Row_created_date, &project_step.Row_effective_date, &project_step.Row_expiry_date, &project_step.Row_quality); err != nil {
			return err // Exit if we get an error
		}

		// Append project_step to result
		result = append(result, project_step)
	}
	// Return result in JSON format
	return c.JSON(result)
	 
}

func SetProjectStep(c *fiber.Ctx) error {
	var project_step dto.Project_step

	setDefaults(&project_step)

	if err := json.Unmarshal(c.Body(), &project_step); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	stmt, err := db.Prepare("insert into project_step values(:1, :2, :3, :4, :5, :6, :7, :8, :9, :10, :11, :12, :13, :14, :15, :16, :17, :18, :19, :20, :21, :22, :23, :24, :25, :26, :27, :28, :29, :30, :31, :32, :33, :34)")
	if err != nil {
		return err
	}
	project_step.Row_created_date = formatDate(project_step.Row_created_date)
	project_step.Row_changed_date = formatDate(project_step.Row_changed_date)
	project_step.Actual_end_date = formatDateString(project_step.Actual_end_date)
	project_step.Actual_start_date = formatDateString(project_step.Actual_start_date)
	project_step.Critical_date = formatDateString(project_step.Critical_date)
	project_step.Due_date = formatDateString(project_step.Due_date)
	project_step.Effective_date = formatDateString(project_step.Effective_date)
	project_step.Expiry_date = formatDateString(project_step.Expiry_date)
	project_step.Plan_end_date = formatDateString(project_step.Plan_end_date)
	project_step.Plan_start_date = formatDateString(project_step.Plan_start_date)
	project_step.Row_effective_date = formatDateString(project_step.Row_effective_date)
	project_step.Row_expiry_date = formatDateString(project_step.Row_expiry_date)






	rows, err := stmt.Exec(project_step.Project_id, project_step.Step_id, project_step.Active_ind, project_step.Actual_end_date, project_step.Actual_start_date, project_step.Actual_time_elapsed, project_step.Actual_time_elapsed_ouom, project_step.Actual_time_elapsed_uom, project_step.Critical_date, project_step.Description, project_step.Due_date, project_step.Effective_date, project_step.Estimated_time_elapsed, project_step.Estimated_time_elapsed_ouom, project_step.Estimated_time_elapsed_uom, project_step.Expiry_date, project_step.Plan_end_date, project_step.Plan_start_date, project_step.Plan_step_id, project_step.Ppdm_guid, project_step.Project_plan_id, project_step.Remark, project_step.Source, project_step.Step_name, project_step.Step_seq_no, project_step.Step_type, project_step.Where_completed, project_step.Row_changed_by, project_step.Row_changed_date, project_step.Row_created_by, project_step.Row_created_date, project_step.Row_effective_date, project_step.Row_expiry_date, project_step.Row_quality)
	
	if err != nil {
		return err
	}

	return c.JSON(rows)
}

func UpdateProjectStep(c *fiber.Ctx) error {
	var project_step dto.Project_step

	setDefaults(&project_step)

	if err := json.Unmarshal(c.Body(), &project_step); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	id := c.Params("id")
	project_step.Project_id = id

    if project_step.Row_created_date != nil {
        return c.Status(400).SendString("Cannot update row_created_date")
    }

	stmt, err := db.Prepare("update project_step set step_id = :1, active_ind = :2, actual_end_date = :3, actual_start_date = :4, actual_time_elapsed = :5, actual_time_elapsed_ouom = :6, actual_time_elapsed_uom = :7, critical_date = :8, description = :9, due_date = :10, effective_date = :11, estimated_time_elapsed = :12, estimated_time_elapsed_ouom = :13, estimated_time_elapsed_uom = :14, expiry_date = :15, plan_end_date = :16, plan_start_date = :17, plan_step_id = :18, ppdm_guid = :19, project_plan_id = :20, remark = :21, source = :22, step_name = :23, step_seq_no = :24, step_type = :25, where_completed = :26, row_changed_by = :27, row_changed_date = :28, row_created_by = :29, row_effective_date = :30, row_expiry_date = :31, row_quality = :32 where project_id = :34")
	if err != nil {
		return err
	}

	project_step.Row_changed_date = formatDate(project_step.Row_changed_date)
	project_step.Actual_end_date = formatDateString(project_step.Actual_end_date)
	project_step.Actual_start_date = formatDateString(project_step.Actual_start_date)
	project_step.Critical_date = formatDateString(project_step.Critical_date)
	project_step.Due_date = formatDateString(project_step.Due_date)
	project_step.Effective_date = formatDateString(project_step.Effective_date)
	project_step.Expiry_date = formatDateString(project_step.Expiry_date)
	project_step.Plan_end_date = formatDateString(project_step.Plan_end_date)
	project_step.Plan_start_date = formatDateString(project_step.Plan_start_date)
	project_step.Row_effective_date = formatDateString(project_step.Row_effective_date)
	project_step.Row_expiry_date = formatDateString(project_step.Row_expiry_date)






	rows, err := stmt.Exec(project_step.Step_id, project_step.Active_ind, project_step.Actual_end_date, project_step.Actual_start_date, project_step.Actual_time_elapsed, project_step.Actual_time_elapsed_ouom, project_step.Actual_time_elapsed_uom, project_step.Critical_date, project_step.Description, project_step.Due_date, project_step.Effective_date, project_step.Estimated_time_elapsed, project_step.Estimated_time_elapsed_ouom, project_step.Estimated_time_elapsed_uom, project_step.Expiry_date, project_step.Plan_end_date, project_step.Plan_start_date, project_step.Plan_step_id, project_step.Ppdm_guid, project_step.Project_plan_id, project_step.Remark, project_step.Source, project_step.Step_name, project_step.Step_seq_no, project_step.Step_type, project_step.Where_completed, project_step.Row_changed_by, project_step.Row_changed_date, project_step.Row_created_by, project_step.Row_effective_date, project_step.Row_expiry_date, project_step.Row_quality, project_step.Project_id)
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

func PatchProjectStep(c *fiber.Ctx) error {
	id := c.Params("id")

	updatedFields := make(map[string]interface{})

	if err := c.BodyParser(&updatedFields); err != nil {
		return err
	}
	if _, ok := updatedFields["row_created_date"]; ok {
        return c.Status(400).SendString("Cannot update row_created_date field")
    }

	query := "update project_step set "
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
		} else if key == "actual_end_date" || key == "actual_start_date" || key == "critical_date" || key == "due_date" || key == "effective_date" || key == "expiry_date" || key == "plan_end_date" || key == "plan_start_date" || key == "row_effective_date" || key == "row_expiry_date"      {
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

func DeleteProjectStep(c *fiber.Ctx) error {
	id := c.Params("id")
	var project_step dto.Project_step
	project_step.Project_id = id

	stmt, err := db.Prepare("delete from project_step where project_id = :1")
	if err != nil {
		return err
	}

	rows, err := stmt.Exec(project_step.Project_id)
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


