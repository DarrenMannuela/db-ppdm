package apiv1

import (
	dto "github.com/DarrenMannuela/pt_gtn_bibliography/dto"
	"github.com/gofiber/fiber/v2"
	"strconv"
	"encoding/json"
)


func GetProjectPlan(c *fiber.Ctx) error {
	rows, err := db.Query("select * from project_plan")
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}
	defer rows.Close()
	var result []dto.Project_plan

	for rows.Next() {
		var project_plan dto.Project_plan
		if err := rows.Scan(&project_plan.Project_plan_id, &project_plan.Active_ind, &project_plan.Description, &project_plan.Effective_date, &project_plan.Expiry_date, &project_plan.Plan_name, &project_plan.Ppdm_guid, &project_plan.Project_type, &project_plan.Remark, &project_plan.Source, &project_plan.Row_changed_by, &project_plan.Row_changed_date, &project_plan.Row_created_by, &project_plan.Row_created_date, &project_plan.Row_effective_date, &project_plan.Row_expiry_date, &project_plan.Row_quality); err != nil {
			return err // Exit if we get an error
		}

		// Append project_plan to result
		result = append(result, project_plan)
	}
	// Return result in JSON format
	return c.JSON(result)
	 
}

func SetProjectPlan(c *fiber.Ctx) error {
	var project_plan dto.Project_plan

	setDefaults(&project_plan)

	if err := json.Unmarshal(c.Body(), &project_plan); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	stmt, err := db.Prepare("insert into project_plan values(:1, :2, :3, :4, :5, :6, :7, :8, :9, :10, :11, :12, :13, :14, :15, :16, :17)")
	if err != nil {
		return err
	}
	project_plan.Row_created_date = formatDate(project_plan.Row_created_date)
	project_plan.Row_changed_date = formatDate(project_plan.Row_changed_date)
	project_plan.Effective_date = formatDateString(project_plan.Effective_date)
	project_plan.Expiry_date = formatDateString(project_plan.Expiry_date)
	project_plan.Row_effective_date = formatDateString(project_plan.Row_effective_date)
	project_plan.Row_expiry_date = formatDateString(project_plan.Row_expiry_date)






	rows, err := stmt.Exec(project_plan.Project_plan_id, project_plan.Active_ind, project_plan.Description, project_plan.Effective_date, project_plan.Expiry_date, project_plan.Plan_name, project_plan.Ppdm_guid, project_plan.Project_type, project_plan.Remark, project_plan.Source, project_plan.Row_changed_by, project_plan.Row_changed_date, project_plan.Row_created_by, project_plan.Row_created_date, project_plan.Row_effective_date, project_plan.Row_expiry_date, project_plan.Row_quality)
	
	if err != nil {
		return err
	}

	return c.JSON(rows)
}

func UpdateProjectPlan(c *fiber.Ctx) error {
	var project_plan dto.Project_plan

	setDefaults(&project_plan)

	if err := json.Unmarshal(c.Body(), &project_plan); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	id := c.Params("id")
	project_plan.Project_plan_id = id

    if project_plan.Row_created_date != nil {
        return c.Status(400).SendString("Cannot update row_created_date")
    }

	stmt, err := db.Prepare("update project_plan set active_ind = :1, description = :2, effective_date = :3, expiry_date = :4, plan_name = :5, ppdm_guid = :6, project_type = :7, remark = :8, source = :9, row_changed_by = :10, row_changed_date = :11, row_created_by = :12, row_effective_date = :13, row_expiry_date = :14, row_quality = :15 where project_plan_id = :17")
	if err != nil {
		return err
	}

	project_plan.Row_changed_date = formatDate(project_plan.Row_changed_date)
	project_plan.Effective_date = formatDateString(project_plan.Effective_date)
	project_plan.Expiry_date = formatDateString(project_plan.Expiry_date)
	project_plan.Row_effective_date = formatDateString(project_plan.Row_effective_date)
	project_plan.Row_expiry_date = formatDateString(project_plan.Row_expiry_date)






	rows, err := stmt.Exec(project_plan.Active_ind, project_plan.Description, project_plan.Effective_date, project_plan.Expiry_date, project_plan.Plan_name, project_plan.Ppdm_guid, project_plan.Project_type, project_plan.Remark, project_plan.Source, project_plan.Row_changed_by, project_plan.Row_changed_date, project_plan.Row_created_by, project_plan.Row_effective_date, project_plan.Row_expiry_date, project_plan.Row_quality, project_plan.Project_plan_id)
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

func PatchProjectPlan(c *fiber.Ctx) error {
	id := c.Params("id")

	updatedFields := make(map[string]interface{})

	if err := c.BodyParser(&updatedFields); err != nil {
		return err
	}
	if _, ok := updatedFields["row_created_date"]; ok {
        return c.Status(400).SendString("Cannot update row_created_date field")
    }

	query := "update project_plan set "
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

func DeleteProjectPlan(c *fiber.Ctx) error {
	id := c.Params("id")
	var project_plan dto.Project_plan
	project_plan.Project_plan_id = id

	stmt, err := db.Prepare("delete from project_plan where project_plan_id = :1")
	if err != nil {
		return err
	}

	rows, err := stmt.Exec(project_plan.Project_plan_id)
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


