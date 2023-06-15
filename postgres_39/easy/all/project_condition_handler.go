package apiv1

import (
	dto "github.com/DarrenMannuela/pt_gtn_bibliography/dto"
	"github.com/gofiber/fiber/v2"
	"strconv"
	"encoding/json"
)


func GetProjectCondition(c *fiber.Ctx) error {
	rows, err := db.Query("select * from project_condition")
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}
	defer rows.Close()
	var result []dto.Project_condition

	for rows.Next() {
		var project_condition dto.Project_condition
		if err := rows.Scan(&project_condition.Project_id, &project_condition.Condition_obs_no, &project_condition.Active_ind, &project_condition.Business_associate_id, &project_condition.Condition_type, &project_condition.Effective_date, &project_condition.End_date, &project_condition.Expiry_date, &project_condition.Ppdm_guid, &project_condition.Project_type, &project_condition.Remark, &project_condition.Source, &project_condition.Start_date, &project_condition.Row_changed_by, &project_condition.Row_changed_date, &project_condition.Row_created_by, &project_condition.Row_created_date, &project_condition.Row_effective_date, &project_condition.Row_expiry_date, &project_condition.Row_quality); err != nil {
			return err // Exit if we get an error
		}

		// Append project_condition to result
		result = append(result, project_condition)
	}
	// Return result in JSON format
	return c.JSON(result)
	 
}

func SetProjectCondition(c *fiber.Ctx) error {
	var project_condition dto.Project_condition

	setDefaults(&project_condition)

	if err := json.Unmarshal(c.Body(), &project_condition); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	stmt, err := db.Prepare("insert into project_condition values(:1, :2, :3, :4, :5, :6, :7, :8, :9, :10, :11, :12, :13, :14, :15, :16, :17, :18, :19, :20)")
	if err != nil {
		return err
	}
	project_condition.Row_created_date = formatDate(project_condition.Row_created_date)
	project_condition.Row_changed_date = formatDate(project_condition.Row_changed_date)
	project_condition.Effective_date = formatDateString(project_condition.Effective_date)
	project_condition.End_date = formatDateString(project_condition.End_date)
	project_condition.Expiry_date = formatDateString(project_condition.Expiry_date)
	project_condition.Start_date = formatDateString(project_condition.Start_date)
	project_condition.Row_effective_date = formatDateString(project_condition.Row_effective_date)
	project_condition.Row_expiry_date = formatDateString(project_condition.Row_expiry_date)






	rows, err := stmt.Exec(project_condition.Project_id, project_condition.Condition_obs_no, project_condition.Active_ind, project_condition.Business_associate_id, project_condition.Condition_type, project_condition.Effective_date, project_condition.End_date, project_condition.Expiry_date, project_condition.Ppdm_guid, project_condition.Project_type, project_condition.Remark, project_condition.Source, project_condition.Start_date, project_condition.Row_changed_by, project_condition.Row_changed_date, project_condition.Row_created_by, project_condition.Row_created_date, project_condition.Row_effective_date, project_condition.Row_expiry_date, project_condition.Row_quality)
	
	if err != nil {
		return err
	}

	return c.JSON(rows)
}

func UpdateProjectCondition(c *fiber.Ctx) error {
	var project_condition dto.Project_condition

	setDefaults(&project_condition)

	if err := json.Unmarshal(c.Body(), &project_condition); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	id := c.Params("id")
	project_condition.Project_id = id

    if project_condition.Row_created_date != nil {
        return c.Status(400).SendString("Cannot update row_created_date")
    }

	stmt, err := db.Prepare("update project_condition set condition_obs_no = :1, active_ind = :2, business_associate_id = :3, condition_type = :4, effective_date = :5, end_date = :6, expiry_date = :7, ppdm_guid = :8, project_type = :9, remark = :10, source = :11, start_date = :12, row_changed_by = :13, row_changed_date = :14, row_created_by = :15, row_effective_date = :16, row_expiry_date = :17, row_quality = :18 where project_id = :20")
	if err != nil {
		return err
	}

	project_condition.Row_changed_date = formatDate(project_condition.Row_changed_date)
	project_condition.Effective_date = formatDateString(project_condition.Effective_date)
	project_condition.End_date = formatDateString(project_condition.End_date)
	project_condition.Expiry_date = formatDateString(project_condition.Expiry_date)
	project_condition.Start_date = formatDateString(project_condition.Start_date)
	project_condition.Row_effective_date = formatDateString(project_condition.Row_effective_date)
	project_condition.Row_expiry_date = formatDateString(project_condition.Row_expiry_date)






	rows, err := stmt.Exec(project_condition.Condition_obs_no, project_condition.Active_ind, project_condition.Business_associate_id, project_condition.Condition_type, project_condition.Effective_date, project_condition.End_date, project_condition.Expiry_date, project_condition.Ppdm_guid, project_condition.Project_type, project_condition.Remark, project_condition.Source, project_condition.Start_date, project_condition.Row_changed_by, project_condition.Row_changed_date, project_condition.Row_created_by, project_condition.Row_effective_date, project_condition.Row_expiry_date, project_condition.Row_quality, project_condition.Project_id)
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

func PatchProjectCondition(c *fiber.Ctx) error {
	id := c.Params("id")

	updatedFields := make(map[string]interface{})

	if err := c.BodyParser(&updatedFields); err != nil {
		return err
	}
	if _, ok := updatedFields["row_created_date"]; ok {
        return c.Status(400).SendString("Cannot update row_created_date field")
    }

	query := "update project_condition set "
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

func DeleteProjectCondition(c *fiber.Ctx) error {
	id := c.Params("id")
	var project_condition dto.Project_condition
	project_condition.Project_id = id

	stmt, err := db.Prepare("delete from project_condition where project_id = :1")
	if err != nil {
		return err
	}

	rows, err := stmt.Exec(project_condition.Project_id)
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


