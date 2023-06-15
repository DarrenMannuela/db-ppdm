package apiv1

import (
	dto "github.com/DarrenMannuela/pt_gtn_bibliography/dto"
	"github.com/gofiber/fiber/v2"
	"strconv"
	"encoding/json"
)


func GetProjectBa(c *fiber.Ctx) error {
	rows, err := db.Query("select * from project_ba")
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}
	defer rows.Close()
	var result []dto.Project_ba

	for rows.Next() {
		var project_ba dto.Project_ba
		if err := rows.Scan(&project_ba.Project_id, &project_ba.Business_associate_id, &project_ba.Active_ind, &project_ba.Effective_date, &project_ba.Expiry_date, &project_ba.Ppdm_guid, &project_ba.Remark, &project_ba.Source, &project_ba.Row_changed_by, &project_ba.Row_changed_date, &project_ba.Row_created_by, &project_ba.Row_created_date, &project_ba.Row_effective_date, &project_ba.Row_expiry_date, &project_ba.Row_quality); err != nil {
			return err // Exit if we get an error
		}

		// Append project_ba to result
		result = append(result, project_ba)
	}
	// Return result in JSON format
	return c.JSON(result)
	 
}

func SetProjectBa(c *fiber.Ctx) error {
	var project_ba dto.Project_ba

	setDefaults(&project_ba)

	if err := json.Unmarshal(c.Body(), &project_ba); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	stmt, err := db.Prepare("insert into project_ba values(:1, :2, :3, :4, :5, :6, :7, :8, :9, :10, :11, :12, :13, :14, :15)")
	if err != nil {
		return err
	}
	project_ba.Row_created_date = formatDate(project_ba.Row_created_date)
	project_ba.Row_changed_date = formatDate(project_ba.Row_changed_date)
	project_ba.Effective_date = formatDateString(project_ba.Effective_date)
	project_ba.Expiry_date = formatDateString(project_ba.Expiry_date)
	project_ba.Row_effective_date = formatDateString(project_ba.Row_effective_date)
	project_ba.Row_expiry_date = formatDateString(project_ba.Row_expiry_date)






	rows, err := stmt.Exec(project_ba.Project_id, project_ba.Business_associate_id, project_ba.Active_ind, project_ba.Effective_date, project_ba.Expiry_date, project_ba.Ppdm_guid, project_ba.Remark, project_ba.Source, project_ba.Row_changed_by, project_ba.Row_changed_date, project_ba.Row_created_by, project_ba.Row_created_date, project_ba.Row_effective_date, project_ba.Row_expiry_date, project_ba.Row_quality)
	
	if err != nil {
		return err
	}

	return c.JSON(rows)
}

func UpdateProjectBa(c *fiber.Ctx) error {
	var project_ba dto.Project_ba

	setDefaults(&project_ba)

	if err := json.Unmarshal(c.Body(), &project_ba); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	id := c.Params("id")
	project_ba.Project_id = id

    if project_ba.Row_created_date != nil {
        return c.Status(400).SendString("Cannot update row_created_date")
    }

	stmt, err := db.Prepare("update project_ba set business_associate_id = :1, active_ind = :2, effective_date = :3, expiry_date = :4, ppdm_guid = :5, remark = :6, source = :7, row_changed_by = :8, row_changed_date = :9, row_created_by = :10, row_effective_date = :11, row_expiry_date = :12, row_quality = :13 where project_id = :15")
	if err != nil {
		return err
	}

	project_ba.Row_changed_date = formatDate(project_ba.Row_changed_date)
	project_ba.Effective_date = formatDateString(project_ba.Effective_date)
	project_ba.Expiry_date = formatDateString(project_ba.Expiry_date)
	project_ba.Row_effective_date = formatDateString(project_ba.Row_effective_date)
	project_ba.Row_expiry_date = formatDateString(project_ba.Row_expiry_date)






	rows, err := stmt.Exec(project_ba.Business_associate_id, project_ba.Active_ind, project_ba.Effective_date, project_ba.Expiry_date, project_ba.Ppdm_guid, project_ba.Remark, project_ba.Source, project_ba.Row_changed_by, project_ba.Row_changed_date, project_ba.Row_created_by, project_ba.Row_effective_date, project_ba.Row_expiry_date, project_ba.Row_quality, project_ba.Project_id)
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

func PatchProjectBa(c *fiber.Ctx) error {
	id := c.Params("id")

	updatedFields := make(map[string]interface{})

	if err := c.BodyParser(&updatedFields); err != nil {
		return err
	}
	if _, ok := updatedFields["row_created_date"]; ok {
        return c.Status(400).SendString("Cannot update row_created_date field")
    }

	query := "update project_ba set "
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

func DeleteProjectBa(c *fiber.Ctx) error {
	id := c.Params("id")
	var project_ba dto.Project_ba
	project_ba.Project_id = id

	stmt, err := db.Prepare("delete from project_ba where project_id = :1")
	if err != nil {
		return err
	}

	rows, err := stmt.Exec(project_ba.Project_id)
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


