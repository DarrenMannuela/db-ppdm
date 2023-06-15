package apiv1

import (
	dto "github.com/DarrenMannuela/pt_gtn_bibliography/dto"
	"github.com/gofiber/fiber/v2"
	"strconv"
	"encoding/json"
)


func GetProjectStatus(c *fiber.Ctx) error {
	rows, err := db.Query("select * from project_status")
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}
	defer rows.Close()
	var result []dto.Project_status

	for rows.Next() {
		var project_status dto.Project_status
		if err := rows.Scan(&project_status.Project_id, &project_status.Status_id, &project_status.Active_ind, &project_status.Defined_by_ba_id, &project_status.Effective_date, &project_status.Expiry_date, &project_status.Ppdm_guid, &project_status.Remark, &project_status.Source, &project_status.Status, &project_status.Status_type, &project_status.Step_id, &project_status.Row_changed_by, &project_status.Row_changed_date, &project_status.Row_created_by, &project_status.Row_created_date, &project_status.Row_effective_date, &project_status.Row_expiry_date, &project_status.Row_quality); err != nil {
			return err // Exit if we get an error
		}

		// Append project_status to result
		result = append(result, project_status)
	}
	// Return result in JSON format
	return c.JSON(result)
	 
}

func SetProjectStatus(c *fiber.Ctx) error {
	var project_status dto.Project_status

	setDefaults(&project_status)

	if err := json.Unmarshal(c.Body(), &project_status); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	stmt, err := db.Prepare("insert into project_status values(:1, :2, :3, :4, :5, :6, :7, :8, :9, :10, :11, :12, :13, :14, :15, :16, :17, :18, :19)")
	if err != nil {
		return err
	}
	project_status.Row_created_date = formatDate(project_status.Row_created_date)
	project_status.Row_changed_date = formatDate(project_status.Row_changed_date)
	project_status.Effective_date = formatDateString(project_status.Effective_date)
	project_status.Expiry_date = formatDateString(project_status.Expiry_date)
	project_status.Row_effective_date = formatDateString(project_status.Row_effective_date)
	project_status.Row_expiry_date = formatDateString(project_status.Row_expiry_date)






	rows, err := stmt.Exec(project_status.Project_id, project_status.Status_id, project_status.Active_ind, project_status.Defined_by_ba_id, project_status.Effective_date, project_status.Expiry_date, project_status.Ppdm_guid, project_status.Remark, project_status.Source, project_status.Status, project_status.Status_type, project_status.Step_id, project_status.Row_changed_by, project_status.Row_changed_date, project_status.Row_created_by, project_status.Row_created_date, project_status.Row_effective_date, project_status.Row_expiry_date, project_status.Row_quality)
	
	if err != nil {
		return err
	}

	return c.JSON(rows)
}

func UpdateProjectStatus(c *fiber.Ctx) error {
	var project_status dto.Project_status

	setDefaults(&project_status)

	if err := json.Unmarshal(c.Body(), &project_status); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	id := c.Params("id")
	project_status.Project_id = id

    if project_status.Row_created_date != nil {
        return c.Status(400).SendString("Cannot update row_created_date")
    }

	stmt, err := db.Prepare("update project_status set status_id = :1, active_ind = :2, defined_by_ba_id = :3, effective_date = :4, expiry_date = :5, ppdm_guid = :6, remark = :7, source = :8, status = :9, status_type = :10, step_id = :11, row_changed_by = :12, row_changed_date = :13, row_created_by = :14, row_effective_date = :15, row_expiry_date = :16, row_quality = :17 where project_id = :19")
	if err != nil {
		return err
	}

	project_status.Row_changed_date = formatDate(project_status.Row_changed_date)
	project_status.Effective_date = formatDateString(project_status.Effective_date)
	project_status.Expiry_date = formatDateString(project_status.Expiry_date)
	project_status.Row_effective_date = formatDateString(project_status.Row_effective_date)
	project_status.Row_expiry_date = formatDateString(project_status.Row_expiry_date)






	rows, err := stmt.Exec(project_status.Status_id, project_status.Active_ind, project_status.Defined_by_ba_id, project_status.Effective_date, project_status.Expiry_date, project_status.Ppdm_guid, project_status.Remark, project_status.Source, project_status.Status, project_status.Status_type, project_status.Step_id, project_status.Row_changed_by, project_status.Row_changed_date, project_status.Row_created_by, project_status.Row_effective_date, project_status.Row_expiry_date, project_status.Row_quality, project_status.Project_id)
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

func PatchProjectStatus(c *fiber.Ctx) error {
	id := c.Params("id")

	updatedFields := make(map[string]interface{})

	if err := c.BodyParser(&updatedFields); err != nil {
		return err
	}
	if _, ok := updatedFields["row_created_date"]; ok {
        return c.Status(400).SendString("Cannot update row_created_date field")
    }

	query := "update project_status set "
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

func DeleteProjectStatus(c *fiber.Ctx) error {
	id := c.Params("id")
	var project_status dto.Project_status
	project_status.Project_id = id

	stmt, err := db.Prepare("delete from project_status where project_id = :1")
	if err != nil {
		return err
	}

	rows, err := stmt.Exec(project_status.Project_id)
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


