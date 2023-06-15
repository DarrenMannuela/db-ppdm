package apiv1

import (
	dto "github.com/DarrenMannuela/pt_gtn_bibliography/dto"
	"github.com/gofiber/fiber/v2"
	"strconv"
	"encoding/json"
)


func GetProjectEquipRole(c *fiber.Ctx) error {
	rows, err := db.Query("select * from project_equip_role")
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}
	defer rows.Close()
	var result []dto.Project_equip_role

	for rows.Next() {
		var project_equip_role dto.Project_equip_role
		if err := rows.Scan(&project_equip_role.Project_id, &project_equip_role.Equipment_obs_no, &project_equip_role.Equipment_role, &project_equip_role.Role_seq_no, &project_equip_role.Active_ind, &project_equip_role.Effective_date, &project_equip_role.Expiry_date, &project_equip_role.Operated_by_ba_id, &project_equip_role.Ppdm_guid, &project_equip_role.Remark, &project_equip_role.Source, &project_equip_role.Row_changed_by, &project_equip_role.Row_changed_date, &project_equip_role.Row_created_by, &project_equip_role.Row_created_date, &project_equip_role.Row_effective_date, &project_equip_role.Row_expiry_date, &project_equip_role.Row_quality); err != nil {
			return err // Exit if we get an error
		}

		// Append project_equip_role to result
		result = append(result, project_equip_role)
	}
	// Return result in JSON format
	return c.JSON(result)
	 
}

func SetProjectEquipRole(c *fiber.Ctx) error {
	var project_equip_role dto.Project_equip_role

	setDefaults(&project_equip_role)

	if err := json.Unmarshal(c.Body(), &project_equip_role); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	stmt, err := db.Prepare("insert into project_equip_role values(:1, :2, :3, :4, :5, :6, :7, :8, :9, :10, :11, :12, :13, :14, :15, :16, :17, :18)")
	if err != nil {
		return err
	}
	project_equip_role.Row_created_date = formatDate(project_equip_role.Row_created_date)
	project_equip_role.Row_changed_date = formatDate(project_equip_role.Row_changed_date)
	project_equip_role.Effective_date = formatDateString(project_equip_role.Effective_date)
	project_equip_role.Expiry_date = formatDateString(project_equip_role.Expiry_date)
	project_equip_role.Row_effective_date = formatDateString(project_equip_role.Row_effective_date)
	project_equip_role.Row_expiry_date = formatDateString(project_equip_role.Row_expiry_date)





	rows, err := stmt.Exec(project_equip_role.Project_id, project_equip_role.Equipment_obs_no, project_equip_role.Equipment_role, project_equip_role.Role_seq_no, project_equip_role.Active_ind, project_equip_role.Effective_date, project_equip_role.Expiry_date, project_equip_role.Operated_by_ba_id, project_equip_role.Ppdm_guid, project_equip_role.Remark, project_equip_role.Source, project_equip_role.Row_changed_by, project_equip_role.Row_changed_date, project_equip_role.Row_created_by, project_equip_role.Row_created_date, project_equip_role.Row_effective_date, project_equip_role.Row_expiry_date, project_equip_role.Row_quality)
	
	if err != nil {
		return err
	}

	return c.JSON(rows)
}

func UpdateProjectEquipRole(c *fiber.Ctx) error {
	var project_equip_role dto.Project_equip_role

	setDefaults(&project_equip_role)

	if err := json.Unmarshal(c.Body(), &project_equip_role); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	id := c.Params("id")
	project_equip_role.Project_id = id

    if project_equip_role.Row_created_date != nil {
        return c.Status(400).SendString("Cannot update row_created_date")
    }

	stmt, err := db.Prepare("update project_equip_role set equipment_obs_no = :1, equipment_role = :2, role_seq_no = :3, active_ind = :4, effective_date = :5, expiry_date = :6, operated_by_ba_id = :7, ppdm_guid = :8, remark = :9, source = :10, row_changed_by = :11, row_changed_date = :12, row_created_by = :13, row_effective_date = :14, row_expiry_date = :15, row_quality = :16 where project_id = :18")
	if err != nil {
		return err
	}

	project_equip_role.Row_changed_date = formatDate(project_equip_role.Row_changed_date)
	project_equip_role.Effective_date = formatDateString(project_equip_role.Effective_date)
	project_equip_role.Expiry_date = formatDateString(project_equip_role.Expiry_date)
	project_equip_role.Row_effective_date = formatDateString(project_equip_role.Row_effective_date)
	project_equip_role.Row_expiry_date = formatDateString(project_equip_role.Row_expiry_date)





	rows, err := stmt.Exec(project_equip_role.Equipment_obs_no, project_equip_role.Equipment_role, project_equip_role.Role_seq_no, project_equip_role.Active_ind, project_equip_role.Effective_date, project_equip_role.Expiry_date, project_equip_role.Operated_by_ba_id, project_equip_role.Ppdm_guid, project_equip_role.Remark, project_equip_role.Source, project_equip_role.Row_changed_by, project_equip_role.Row_changed_date, project_equip_role.Row_created_by, project_equip_role.Row_effective_date, project_equip_role.Row_expiry_date, project_equip_role.Row_quality, project_equip_role.Project_id)
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

func PatchProjectEquipRole(c *fiber.Ctx) error {
	id := c.Params("id")

	updatedFields := make(map[string]interface{})

	if err := c.BodyParser(&updatedFields); err != nil {
		return err
	}
	if _, ok := updatedFields["row_created_date"]; ok {
        return c.Status(400).SendString("Cannot update row_created_date field")
    }

	query := "update project_equip_role set "
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

func DeleteProjectEquipRole(c *fiber.Ctx) error {
	id := c.Params("id")
	var project_equip_role dto.Project_equip_role
	project_equip_role.Project_id = id

	stmt, err := db.Prepare("delete from project_equip_role where project_id = :1")
	if err != nil {
		return err
	}

	rows, err := stmt.Exec(project_equip_role.Project_id)
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


