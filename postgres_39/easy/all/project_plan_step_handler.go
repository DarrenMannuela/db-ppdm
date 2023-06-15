package apiv1

import (
	dto "github.com/DarrenMannuela/pt_gtn_bibliography/dto"
	"github.com/gofiber/fiber/v2"
	"strconv"
	"encoding/json"
)


func GetProjectPlanStep(c *fiber.Ctx) error {
	rows, err := db.Query("select * from project_plan_step")
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}
	defer rows.Close()
	var result []dto.Project_plan_step

	for rows.Next() {
		var project_plan_step dto.Project_plan_step
		if err := rows.Scan(&project_plan_step.Project_plan_id, &project_plan_step.Plan_step_id, &project_plan_step.Active_ind, &project_plan_step.Description, &project_plan_step.Effective_date, &project_plan_step.Expiry_date, &project_plan_step.Ppdm_guid, &project_plan_step.Remark, &project_plan_step.Source, &project_plan_step.Step_name, &project_plan_step.Step_rule, &project_plan_step.Step_seq_no, &project_plan_step.Step_type, &project_plan_step.Row_changed_by, &project_plan_step.Row_changed_date, &project_plan_step.Row_created_by, &project_plan_step.Row_created_date, &project_plan_step.Row_effective_date, &project_plan_step.Row_expiry_date, &project_plan_step.Row_quality); err != nil {
			return err // Exit if we get an error
		}

		// Append project_plan_step to result
		result = append(result, project_plan_step)
	}
	// Return result in JSON format
	return c.JSON(result)
	 
}

func SetProjectPlanStep(c *fiber.Ctx) error {
	var project_plan_step dto.Project_plan_step

	setDefaults(&project_plan_step)

	if err := json.Unmarshal(c.Body(), &project_plan_step); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	stmt, err := db.Prepare("insert into project_plan_step values(:1, :2, :3, :4, :5, :6, :7, :8, :9, :10, :11, :12, :13, :14, :15, :16, :17, :18, :19, :20)")
	if err != nil {
		return err
	}
	project_plan_step.Row_created_date = formatDate(project_plan_step.Row_created_date)
	project_plan_step.Row_changed_date = formatDate(project_plan_step.Row_changed_date)
	project_plan_step.Effective_date = formatDateString(project_plan_step.Effective_date)
	project_plan_step.Expiry_date = formatDateString(project_plan_step.Expiry_date)
	project_plan_step.Row_effective_date = formatDateString(project_plan_step.Row_effective_date)
	project_plan_step.Row_expiry_date = formatDateString(project_plan_step.Row_expiry_date)






	rows, err := stmt.Exec(project_plan_step.Project_plan_id, project_plan_step.Plan_step_id, project_plan_step.Active_ind, project_plan_step.Description, project_plan_step.Effective_date, project_plan_step.Expiry_date, project_plan_step.Ppdm_guid, project_plan_step.Remark, project_plan_step.Source, project_plan_step.Step_name, project_plan_step.Step_rule, project_plan_step.Step_seq_no, project_plan_step.Step_type, project_plan_step.Row_changed_by, project_plan_step.Row_changed_date, project_plan_step.Row_created_by, project_plan_step.Row_created_date, project_plan_step.Row_effective_date, project_plan_step.Row_expiry_date, project_plan_step.Row_quality)
	
	if err != nil {
		return err
	}

	return c.JSON(rows)
}

func UpdateProjectPlanStep(c *fiber.Ctx) error {
	var project_plan_step dto.Project_plan_step

	setDefaults(&project_plan_step)

	if err := json.Unmarshal(c.Body(), &project_plan_step); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	id := c.Params("id")
	project_plan_step.Project_plan_id = id

    if project_plan_step.Row_created_date != nil {
        return c.Status(400).SendString("Cannot update row_created_date")
    }

	stmt, err := db.Prepare("update project_plan_step set plan_step_id = :1, active_ind = :2, description = :3, effective_date = :4, expiry_date = :5, ppdm_guid = :6, remark = :7, source = :8, step_name = :9, step_rule = :10, step_seq_no = :11, step_type = :12, row_changed_by = :13, row_changed_date = :14, row_created_by = :15, row_effective_date = :16, row_expiry_date = :17, row_quality = :18 where project_plan_id = :20")
	if err != nil {
		return err
	}

	project_plan_step.Row_changed_date = formatDate(project_plan_step.Row_changed_date)
	project_plan_step.Effective_date = formatDateString(project_plan_step.Effective_date)
	project_plan_step.Expiry_date = formatDateString(project_plan_step.Expiry_date)
	project_plan_step.Row_effective_date = formatDateString(project_plan_step.Row_effective_date)
	project_plan_step.Row_expiry_date = formatDateString(project_plan_step.Row_expiry_date)






	rows, err := stmt.Exec(project_plan_step.Plan_step_id, project_plan_step.Active_ind, project_plan_step.Description, project_plan_step.Effective_date, project_plan_step.Expiry_date, project_plan_step.Ppdm_guid, project_plan_step.Remark, project_plan_step.Source, project_plan_step.Step_name, project_plan_step.Step_rule, project_plan_step.Step_seq_no, project_plan_step.Step_type, project_plan_step.Row_changed_by, project_plan_step.Row_changed_date, project_plan_step.Row_created_by, project_plan_step.Row_effective_date, project_plan_step.Row_expiry_date, project_plan_step.Row_quality, project_plan_step.Project_plan_id)
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

func PatchProjectPlanStep(c *fiber.Ctx) error {
	id := c.Params("id")

	updatedFields := make(map[string]interface{})

	if err := c.BodyParser(&updatedFields); err != nil {
		return err
	}
	if _, ok := updatedFields["row_created_date"]; ok {
        return c.Status(400).SendString("Cannot update row_created_date field")
    }

	query := "update project_plan_step set "
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
	query += " where project_plan_id = :id"

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

func DeleteProjectPlanStep(c *fiber.Ctx) error {
	id := c.Params("id")
	var project_plan_step dto.Project_plan_step
	project_plan_step.Project_plan_id = id

	stmt, err := db.Prepare("delete from project_plan_step where project_plan_id = :1")
	if err != nil {
		return err
	}

	rows, err := stmt.Exec(project_plan_step.Project_plan_id)
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


