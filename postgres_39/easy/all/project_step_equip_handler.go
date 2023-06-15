package apiv1

import (
	dto "github.com/DarrenMannuela/pt_gtn_bibliography/dto"
	"github.com/gofiber/fiber/v2"
	"strconv"
	"encoding/json"
)


func GetProjectStepEquip(c *fiber.Ctx) error {
	rows, err := db.Query("select * from project_step_equip")
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}
	defer rows.Close()
	var result []dto.Project_step_equip

	for rows.Next() {
		var project_step_equip dto.Project_step_equip
		if err := rows.Scan(&project_step_equip.Project_id, &project_step_equip.Equipment_obs_no, &project_step_equip.Equipment_role, &project_step_equip.Role_seq_no, &project_step_equip.Step_id, &project_step_equip.Active_ind, &project_step_equip.Effective_date, &project_step_equip.Expiry_date, &project_step_equip.Operated_by_ba_id, &project_step_equip.Ppdm_guid, &project_step_equip.Remark, &project_step_equip.Source, &project_step_equip.Row_changed_by, &project_step_equip.Row_changed_date, &project_step_equip.Row_created_by, &project_step_equip.Row_created_date, &project_step_equip.Row_effective_date, &project_step_equip.Row_expiry_date, &project_step_equip.Row_quality); err != nil {
			return err // Exit if we get an error
		}

		// Append project_step_equip to result
		result = append(result, project_step_equip)
	}
	// Return result in JSON format
	return c.JSON(result)
	 
}

func SetProjectStepEquip(c *fiber.Ctx) error {
	var project_step_equip dto.Project_step_equip

	setDefaults(&project_step_equip)

	if err := json.Unmarshal(c.Body(), &project_step_equip); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	stmt, err := db.Prepare("insert into project_step_equip values(:1, :2, :3, :4, :5, :6, :7, :8, :9, :10, :11, :12, :13, :14, :15, :16, :17, :18, :19)")
	if err != nil {
		return err
	}
	project_step_equip.Row_created_date = formatDate(project_step_equip.Row_created_date)
	project_step_equip.Row_changed_date = formatDate(project_step_equip.Row_changed_date)
	project_step_equip.Effective_date = formatDateString(project_step_equip.Effective_date)
	project_step_equip.Expiry_date = formatDateString(project_step_equip.Expiry_date)
	project_step_equip.Row_effective_date = formatDateString(project_step_equip.Row_effective_date)
	project_step_equip.Row_expiry_date = formatDateString(project_step_equip.Row_expiry_date)






	rows, err := stmt.Exec(project_step_equip.Project_id, project_step_equip.Equipment_obs_no, project_step_equip.Equipment_role, project_step_equip.Role_seq_no, project_step_equip.Step_id, project_step_equip.Active_ind, project_step_equip.Effective_date, project_step_equip.Expiry_date, project_step_equip.Operated_by_ba_id, project_step_equip.Ppdm_guid, project_step_equip.Remark, project_step_equip.Source, project_step_equip.Row_changed_by, project_step_equip.Row_changed_date, project_step_equip.Row_created_by, project_step_equip.Row_created_date, project_step_equip.Row_effective_date, project_step_equip.Row_expiry_date, project_step_equip.Row_quality)
	
	if err != nil {
		return err
	}

	return c.JSON(rows)
}

func UpdateProjectStepEquip(c *fiber.Ctx) error {
	var project_step_equip dto.Project_step_equip

	setDefaults(&project_step_equip)

	if err := json.Unmarshal(c.Body(), &project_step_equip); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	id := c.Params("id")
	project_step_equip.Project_id = id

    if project_step_equip.Row_created_date != nil {
        return c.Status(400).SendString("Cannot update row_created_date")
    }

	stmt, err := db.Prepare("update project_step_equip set equipment_obs_no = :1, equipment_role = :2, role_seq_no = :3, step_id = :4, active_ind = :5, effective_date = :6, expiry_date = :7, operated_by_ba_id = :8, ppdm_guid = :9, remark = :10, source = :11, row_changed_by = :12, row_changed_date = :13, row_created_by = :14, row_effective_date = :15, row_expiry_date = :16, row_quality = :17 where project_id = :19")
	if err != nil {
		return err
	}

	project_step_equip.Row_changed_date = formatDate(project_step_equip.Row_changed_date)
	project_step_equip.Effective_date = formatDateString(project_step_equip.Effective_date)
	project_step_equip.Expiry_date = formatDateString(project_step_equip.Expiry_date)
	project_step_equip.Row_effective_date = formatDateString(project_step_equip.Row_effective_date)
	project_step_equip.Row_expiry_date = formatDateString(project_step_equip.Row_expiry_date)






	rows, err := stmt.Exec(project_step_equip.Equipment_obs_no, project_step_equip.Equipment_role, project_step_equip.Role_seq_no, project_step_equip.Step_id, project_step_equip.Active_ind, project_step_equip.Effective_date, project_step_equip.Expiry_date, project_step_equip.Operated_by_ba_id, project_step_equip.Ppdm_guid, project_step_equip.Remark, project_step_equip.Source, project_step_equip.Row_changed_by, project_step_equip.Row_changed_date, project_step_equip.Row_created_by, project_step_equip.Row_effective_date, project_step_equip.Row_expiry_date, project_step_equip.Row_quality, project_step_equip.Project_id)
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

func PatchProjectStepEquip(c *fiber.Ctx) error {
	id := c.Params("id")

	updatedFields := make(map[string]interface{})

	if err := c.BodyParser(&updatedFields); err != nil {
		return err
	}
	if _, ok := updatedFields["row_created_date"]; ok {
        return c.Status(400).SendString("Cannot update row_created_date field")
    }

	query := "update project_step_equip set "
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

func DeleteProjectStepEquip(c *fiber.Ctx) error {
	id := c.Params("id")
	var project_step_equip dto.Project_step_equip
	project_step_equip.Project_id = id

	stmt, err := db.Prepare("delete from project_step_equip where project_id = :1")
	if err != nil {
		return err
	}

	rows, err := stmt.Exec(project_step_equip.Project_id)
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


