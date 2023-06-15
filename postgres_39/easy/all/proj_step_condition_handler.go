package apiv1

import (
	dto "github.com/DarrenMannuela/pt_gtn_bibliography/dto"
	"github.com/gofiber/fiber/v2"
	"strconv"
	"encoding/json"
)


func GetProjStepCondition(c *fiber.Ctx) error {
	rows, err := db.Query("select * from proj_step_condition")
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}
	defer rows.Close()
	var result []dto.Proj_step_condition

	for rows.Next() {
		var proj_step_condition dto.Proj_step_condition
		if err := rows.Scan(&proj_step_condition.Project_id, &proj_step_condition.Project_step_id, &proj_step_condition.Condition_obs_no, &proj_step_condition.Active_ind, &proj_step_condition.Business_associate_id, &proj_step_condition.Effective_date, &proj_step_condition.End_date, &proj_step_condition.Expiry_date, &proj_step_condition.Ppdm_guid, &proj_step_condition.Project_condition_id, &proj_step_condition.Project_type, &proj_step_condition.Remark, &proj_step_condition.Required_step_id, &proj_step_condition.Source, &proj_step_condition.Start_date, &proj_step_condition.Row_changed_by, &proj_step_condition.Row_changed_date, &proj_step_condition.Row_created_by, &proj_step_condition.Row_created_date, &proj_step_condition.Row_effective_date, &proj_step_condition.Row_expiry_date, &proj_step_condition.Row_quality); err != nil {
			return err // Exit if we get an error
		}

		// Append proj_step_condition to result
		result = append(result, proj_step_condition)
	}
	// Return result in JSON format
	return c.JSON(result)
	 
}

func SetProjStepCondition(c *fiber.Ctx) error {
	var proj_step_condition dto.Proj_step_condition

	setDefaults(&proj_step_condition)

	if err := json.Unmarshal(c.Body(), &proj_step_condition); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	stmt, err := db.Prepare("insert into proj_step_condition values(:1, :2, :3, :4, :5, :6, :7, :8, :9, :10, :11, :12, :13, :14, :15, :16, :17, :18, :19, :20, :21, :22)")
	if err != nil {
		return err
	}
	proj_step_condition.Row_created_date = formatDate(proj_step_condition.Row_created_date)
	proj_step_condition.Row_changed_date = formatDate(proj_step_condition.Row_changed_date)
	proj_step_condition.Effective_date = formatDateString(proj_step_condition.Effective_date)
	proj_step_condition.End_date = formatDateString(proj_step_condition.End_date)
	proj_step_condition.Expiry_date = formatDateString(proj_step_condition.Expiry_date)
	proj_step_condition.Start_date = formatDateString(proj_step_condition.Start_date)
	proj_step_condition.Row_effective_date = formatDateString(proj_step_condition.Row_effective_date)
	proj_step_condition.Row_expiry_date = formatDateString(proj_step_condition.Row_expiry_date)






	rows, err := stmt.Exec(proj_step_condition.Project_id, proj_step_condition.Project_step_id, proj_step_condition.Condition_obs_no, proj_step_condition.Active_ind, proj_step_condition.Business_associate_id, proj_step_condition.Effective_date, proj_step_condition.End_date, proj_step_condition.Expiry_date, proj_step_condition.Ppdm_guid, proj_step_condition.Project_condition_id, proj_step_condition.Project_type, proj_step_condition.Remark, proj_step_condition.Required_step_id, proj_step_condition.Source, proj_step_condition.Start_date, proj_step_condition.Row_changed_by, proj_step_condition.Row_changed_date, proj_step_condition.Row_created_by, proj_step_condition.Row_created_date, proj_step_condition.Row_effective_date, proj_step_condition.Row_expiry_date, proj_step_condition.Row_quality)
	
	if err != nil {
		return err
	}

	return c.JSON(rows)
}

func UpdateProjStepCondition(c *fiber.Ctx) error {
	var proj_step_condition dto.Proj_step_condition

	setDefaults(&proj_step_condition)

	if err := json.Unmarshal(c.Body(), &proj_step_condition); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	id := c.Params("id")
	proj_step_condition.Project_id = id

    if proj_step_condition.Row_created_date != nil {
        return c.Status(400).SendString("Cannot update row_created_date")
    }

	stmt, err := db.Prepare("update proj_step_condition set project_step_id = :1, condition_obs_no = :2, active_ind = :3, business_associate_id = :4, effective_date = :5, end_date = :6, expiry_date = :7, ppdm_guid = :8, project_condition_id = :9, project_type = :10, remark = :11, required_step_id = :12, source = :13, start_date = :14, row_changed_by = :15, row_changed_date = :16, row_created_by = :17, row_effective_date = :18, row_expiry_date = :19, row_quality = :20 where project_id = :22")
	if err != nil {
		return err
	}

	proj_step_condition.Row_changed_date = formatDate(proj_step_condition.Row_changed_date)
	proj_step_condition.Effective_date = formatDateString(proj_step_condition.Effective_date)
	proj_step_condition.End_date = formatDateString(proj_step_condition.End_date)
	proj_step_condition.Expiry_date = formatDateString(proj_step_condition.Expiry_date)
	proj_step_condition.Start_date = formatDateString(proj_step_condition.Start_date)
	proj_step_condition.Row_effective_date = formatDateString(proj_step_condition.Row_effective_date)
	proj_step_condition.Row_expiry_date = formatDateString(proj_step_condition.Row_expiry_date)






	rows, err := stmt.Exec(proj_step_condition.Project_step_id, proj_step_condition.Condition_obs_no, proj_step_condition.Active_ind, proj_step_condition.Business_associate_id, proj_step_condition.Effective_date, proj_step_condition.End_date, proj_step_condition.Expiry_date, proj_step_condition.Ppdm_guid, proj_step_condition.Project_condition_id, proj_step_condition.Project_type, proj_step_condition.Remark, proj_step_condition.Required_step_id, proj_step_condition.Source, proj_step_condition.Start_date, proj_step_condition.Row_changed_by, proj_step_condition.Row_changed_date, proj_step_condition.Row_created_by, proj_step_condition.Row_effective_date, proj_step_condition.Row_expiry_date, proj_step_condition.Row_quality, proj_step_condition.Project_id)
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

func PatchProjStepCondition(c *fiber.Ctx) error {
	id := c.Params("id")

	updatedFields := make(map[string]interface{})

	if err := c.BodyParser(&updatedFields); err != nil {
		return err
	}
	if _, ok := updatedFields["row_created_date"]; ok {
        return c.Status(400).SendString("Cannot update row_created_date field")
    }

	query := "update proj_step_condition set "
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

func DeleteProjStepCondition(c *fiber.Ctx) error {
	id := c.Params("id")
	var proj_step_condition dto.Proj_step_condition
	proj_step_condition.Project_id = id

	stmt, err := db.Prepare("delete from proj_step_condition where project_id = :1")
	if err != nil {
		return err
	}

	rows, err := stmt.Exec(proj_step_condition.Project_id)
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


