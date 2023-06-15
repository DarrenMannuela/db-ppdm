package apiv1

import (
	dto "github.com/DarrenMannuela/pt_gtn_bibliography/dto"
	"github.com/gofiber/fiber/v2"
	"strconv"
	"encoding/json"
)


func GetProjectBaRole(c *fiber.Ctx) error {
	rows, err := db.Query("select * from project_ba_role")
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}
	defer rows.Close()
	var result []dto.Project_ba_role

	for rows.Next() {
		var project_ba_role dto.Project_ba_role
		if err := rows.Scan(&project_ba_role.Project_id, &project_ba_role.Business_associate_id, &project_ba_role.Role, &project_ba_role.Role_seq_no, &project_ba_role.Active_ind, &project_ba_role.Effective_date, &project_ba_role.Expiry_date, &project_ba_role.Ppdm_guid, &project_ba_role.Remark, &project_ba_role.Source, &project_ba_role.Row_changed_by, &project_ba_role.Row_changed_date, &project_ba_role.Row_created_by, &project_ba_role.Row_created_date, &project_ba_role.Row_effective_date, &project_ba_role.Row_expiry_date, &project_ba_role.Row_quality); err != nil {
			return err // Exit if we get an error
		}

		// Append project_ba_role to result
		result = append(result, project_ba_role)
	}
	// Return result in JSON format
	return c.JSON(result)
	 
}

func SetProjectBaRole(c *fiber.Ctx) error {
	var project_ba_role dto.Project_ba_role

	setDefaults(&project_ba_role)

	if err := json.Unmarshal(c.Body(), &project_ba_role); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	stmt, err := db.Prepare("insert into project_ba_role values(:1, :2, :3, :4, :5, :6, :7, :8, :9, :10, :11, :12, :13, :14, :15, :16, :17)")
	if err != nil {
		return err
	}
	project_ba_role.Row_created_date = formatDate(project_ba_role.Row_created_date)
	project_ba_role.Row_changed_date = formatDate(project_ba_role.Row_changed_date)
	project_ba_role.Effective_date = formatDateString(project_ba_role.Effective_date)
	project_ba_role.Expiry_date = formatDateString(project_ba_role.Expiry_date)
	project_ba_role.Row_effective_date = formatDateString(project_ba_role.Row_effective_date)
	project_ba_role.Row_expiry_date = formatDateString(project_ba_role.Row_expiry_date)





	rows, err := stmt.Exec(project_ba_role.Project_id, project_ba_role.Business_associate_id, project_ba_role.Role, project_ba_role.Role_seq_no, project_ba_role.Active_ind, project_ba_role.Effective_date, project_ba_role.Expiry_date, project_ba_role.Ppdm_guid, project_ba_role.Remark, project_ba_role.Source, project_ba_role.Row_changed_by, project_ba_role.Row_changed_date, project_ba_role.Row_created_by, project_ba_role.Row_created_date, project_ba_role.Row_effective_date, project_ba_role.Row_expiry_date, project_ba_role.Row_quality)
	
	if err != nil {
		return err
	}

	return c.JSON(rows)
}

func UpdateProjectBaRole(c *fiber.Ctx) error {
	var project_ba_role dto.Project_ba_role

	setDefaults(&project_ba_role)

	if err := json.Unmarshal(c.Body(), &project_ba_role); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	id := c.Params("id")
	project_ba_role.Project_id = id

    if project_ba_role.Row_created_date != nil {
        return c.Status(400).SendString("Cannot update row_created_date")
    }

	stmt, err := db.Prepare("update project_ba_role set business_associate_id = :1, role = :2, role_seq_no = :3, active_ind = :4, effective_date = :5, expiry_date = :6, ppdm_guid = :7, remark = :8, source = :9, row_changed_by = :10, row_changed_date = :11, row_created_by = :12, row_effective_date = :13, row_expiry_date = :14, row_quality = :15 where project_id = :17")
	if err != nil {
		return err
	}

	project_ba_role.Row_changed_date = formatDate(project_ba_role.Row_changed_date)
	project_ba_role.Effective_date = formatDateString(project_ba_role.Effective_date)
	project_ba_role.Expiry_date = formatDateString(project_ba_role.Expiry_date)
	project_ba_role.Row_effective_date = formatDateString(project_ba_role.Row_effective_date)
	project_ba_role.Row_expiry_date = formatDateString(project_ba_role.Row_expiry_date)





	rows, err := stmt.Exec(project_ba_role.Business_associate_id, project_ba_role.Role, project_ba_role.Role_seq_no, project_ba_role.Active_ind, project_ba_role.Effective_date, project_ba_role.Expiry_date, project_ba_role.Ppdm_guid, project_ba_role.Remark, project_ba_role.Source, project_ba_role.Row_changed_by, project_ba_role.Row_changed_date, project_ba_role.Row_created_by, project_ba_role.Row_effective_date, project_ba_role.Row_expiry_date, project_ba_role.Row_quality, project_ba_role.Project_id)
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

func PatchProjectBaRole(c *fiber.Ctx) error {
	id := c.Params("id")

	updatedFields := make(map[string]interface{})

	if err := c.BodyParser(&updatedFields); err != nil {
		return err
	}
	if _, ok := updatedFields["row_created_date"]; ok {
        return c.Status(400).SendString("Cannot update row_created_date field")
    }

	query := "update project_ba_role set "
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
		} else if key == "effective_date" || key == "expiry_date" || key == "row_effective_date" || key == "row_expiry_date"     {
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

func DeleteProjectBaRole(c *fiber.Ctx) error {
	id := c.Params("id")
	var project_ba_role dto.Project_ba_role
	project_ba_role.Project_id = id

	stmt, err := db.Prepare("delete from project_ba_role where project_id = :1")
	if err != nil {
		return err
	}

	rows, err := stmt.Exec(project_ba_role.Project_id)
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


