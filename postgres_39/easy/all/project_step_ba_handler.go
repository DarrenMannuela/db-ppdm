package apiv1

import (
	dto "github.com/DarrenMannuela/pt_gtn_bibliography/dto"
	"github.com/gofiber/fiber/v2"
	"strconv"
	"encoding/json"
)


func GetProjectStepBa(c *fiber.Ctx) error {
	rows, err := db.Query("select * from project_step_ba")
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}
	defer rows.Close()
	var result []dto.Project_step_ba

	for rows.Next() {
		var project_step_ba dto.Project_step_ba
		if err := rows.Scan(&project_step_ba.Project_id, &project_step_ba.Business_associate_id, &project_step_ba.Role, &project_step_ba.Role_seq_no, &project_step_ba.Step_id, &project_step_ba.Step_ba_obs_no, &project_step_ba.Active_ind, &project_step_ba.Actual_ind, &project_step_ba.Effective_date, &project_step_ba.Expiry_date, &project_step_ba.Plan_ind, &project_step_ba.Ppdm_guid, &project_step_ba.Remark, &project_step_ba.Source, &project_step_ba.Row_changed_by, &project_step_ba.Row_changed_date, &project_step_ba.Row_created_by, &project_step_ba.Row_created_date, &project_step_ba.Row_effective_date, &project_step_ba.Row_expiry_date, &project_step_ba.Row_quality); err != nil {
			return err // Exit if we get an error
		}

		// Append project_step_ba to result
		result = append(result, project_step_ba)
	}
	// Return result in JSON format
	return c.JSON(result)
	 
}

func SetProjectStepBa(c *fiber.Ctx) error {
	var project_step_ba dto.Project_step_ba

	setDefaults(&project_step_ba)

	if err := json.Unmarshal(c.Body(), &project_step_ba); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	stmt, err := db.Prepare("insert into project_step_ba values(:1, :2, :3, :4, :5, :6, :7, :8, :9, :10, :11, :12, :13, :14, :15, :16, :17, :18, :19, :20, :21)")
	if err != nil {
		return err
	}
	project_step_ba.Row_created_date = formatDate(project_step_ba.Row_created_date)
	project_step_ba.Row_changed_date = formatDate(project_step_ba.Row_changed_date)
	project_step_ba.Effective_date = formatDateString(project_step_ba.Effective_date)
	project_step_ba.Expiry_date = formatDateString(project_step_ba.Expiry_date)
	project_step_ba.Row_effective_date = formatDateString(project_step_ba.Row_effective_date)
	project_step_ba.Row_expiry_date = formatDateString(project_step_ba.Row_expiry_date)






	rows, err := stmt.Exec(project_step_ba.Project_id, project_step_ba.Business_associate_id, project_step_ba.Role, project_step_ba.Role_seq_no, project_step_ba.Step_id, project_step_ba.Step_ba_obs_no, project_step_ba.Active_ind, project_step_ba.Actual_ind, project_step_ba.Effective_date, project_step_ba.Expiry_date, project_step_ba.Plan_ind, project_step_ba.Ppdm_guid, project_step_ba.Remark, project_step_ba.Source, project_step_ba.Row_changed_by, project_step_ba.Row_changed_date, project_step_ba.Row_created_by, project_step_ba.Row_created_date, project_step_ba.Row_effective_date, project_step_ba.Row_expiry_date, project_step_ba.Row_quality)
	
	if err != nil {
		return err
	}

	return c.JSON(rows)
}

func UpdateProjectStepBa(c *fiber.Ctx) error {
	var project_step_ba dto.Project_step_ba

	setDefaults(&project_step_ba)

	if err := json.Unmarshal(c.Body(), &project_step_ba); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	id := c.Params("id")
	project_step_ba.Project_id = id

    if project_step_ba.Row_created_date != nil {
        return c.Status(400).SendString("Cannot update row_created_date")
    }

	stmt, err := db.Prepare("update project_step_ba set business_associate_id = :1, role = :2, role_seq_no = :3, step_id = :4, step_ba_obs_no = :5, active_ind = :6, actual_ind = :7, effective_date = :8, expiry_date = :9, plan_ind = :10, ppdm_guid = :11, remark = :12, source = :13, row_changed_by = :14, row_changed_date = :15, row_created_by = :16, row_effective_date = :17, row_expiry_date = :18, row_quality = :19 where project_id = :21")
	if err != nil {
		return err
	}

	project_step_ba.Row_changed_date = formatDate(project_step_ba.Row_changed_date)
	project_step_ba.Effective_date = formatDateString(project_step_ba.Effective_date)
	project_step_ba.Expiry_date = formatDateString(project_step_ba.Expiry_date)
	project_step_ba.Row_effective_date = formatDateString(project_step_ba.Row_effective_date)
	project_step_ba.Row_expiry_date = formatDateString(project_step_ba.Row_expiry_date)






	rows, err := stmt.Exec(project_step_ba.Business_associate_id, project_step_ba.Role, project_step_ba.Role_seq_no, project_step_ba.Step_id, project_step_ba.Step_ba_obs_no, project_step_ba.Active_ind, project_step_ba.Actual_ind, project_step_ba.Effective_date, project_step_ba.Expiry_date, project_step_ba.Plan_ind, project_step_ba.Ppdm_guid, project_step_ba.Remark, project_step_ba.Source, project_step_ba.Row_changed_by, project_step_ba.Row_changed_date, project_step_ba.Row_created_by, project_step_ba.Row_effective_date, project_step_ba.Row_expiry_date, project_step_ba.Row_quality, project_step_ba.Project_id)
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

func PatchProjectStepBa(c *fiber.Ctx) error {
	id := c.Params("id")

	updatedFields := make(map[string]interface{})

	if err := c.BodyParser(&updatedFields); err != nil {
		return err
	}
	if _, ok := updatedFields["row_created_date"]; ok {
        return c.Status(400).SendString("Cannot update row_created_date field")
    }

	query := "update project_step_ba set "
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

func DeleteProjectStepBa(c *fiber.Ctx) error {
	id := c.Params("id")
	var project_step_ba dto.Project_step_ba
	project_step_ba.Project_id = id

	stmt, err := db.Prepare("delete from project_step_ba where project_id = :1")
	if err != nil {
		return err
	}

	rows, err := stmt.Exec(project_step_ba.Project_id)
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


