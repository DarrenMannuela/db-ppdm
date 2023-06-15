package apiv1

import (
	dto "github.com/DarrenMannuela/pt_gtn_bibliography/dto"
	"github.com/gofiber/fiber/v2"
	"strconv"
	"encoding/json"
)


func GetProjectEquipment(c *fiber.Ctx) error {
	rows, err := db.Query("select * from project_equipment")
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}
	defer rows.Close()
	var result []dto.Project_equipment

	for rows.Next() {
		var project_equipment dto.Project_equipment
		if err := rows.Scan(&project_equipment.Project_id, &project_equipment.Equipment_obs_no, &project_equipment.Active_ind, &project_equipment.Catalogue_equip_id, &project_equipment.Effective_date, &project_equipment.Equipment_id, &project_equipment.Expiry_date, &project_equipment.Operated_by_ba_id, &project_equipment.Ppdm_guid, &project_equipment.Remark, &project_equipment.Source, &project_equipment.Row_changed_by, &project_equipment.Row_changed_date, &project_equipment.Row_created_by, &project_equipment.Row_created_date, &project_equipment.Row_effective_date, &project_equipment.Row_expiry_date, &project_equipment.Row_quality); err != nil {
			return err // Exit if we get an error
		}

		// Append project_equipment to result
		result = append(result, project_equipment)
	}
	// Return result in JSON format
	return c.JSON(result)
	 
}

func SetProjectEquipment(c *fiber.Ctx) error {
	var project_equipment dto.Project_equipment

	setDefaults(&project_equipment)

	if err := json.Unmarshal(c.Body(), &project_equipment); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	stmt, err := db.Prepare("insert into project_equipment values(:1, :2, :3, :4, :5, :6, :7, :8, :9, :10, :11, :12, :13, :14, :15, :16, :17, :18)")
	if err != nil {
		return err
	}
	project_equipment.Row_created_date = formatDate(project_equipment.Row_created_date)
	project_equipment.Row_changed_date = formatDate(project_equipment.Row_changed_date)
	project_equipment.Effective_date = formatDateString(project_equipment.Effective_date)
	project_equipment.Expiry_date = formatDateString(project_equipment.Expiry_date)
	project_equipment.Row_effective_date = formatDateString(project_equipment.Row_effective_date)
	project_equipment.Row_expiry_date = formatDateString(project_equipment.Row_expiry_date)






	rows, err := stmt.Exec(project_equipment.Project_id, project_equipment.Equipment_obs_no, project_equipment.Active_ind, project_equipment.Catalogue_equip_id, project_equipment.Effective_date, project_equipment.Equipment_id, project_equipment.Expiry_date, project_equipment.Operated_by_ba_id, project_equipment.Ppdm_guid, project_equipment.Remark, project_equipment.Source, project_equipment.Row_changed_by, project_equipment.Row_changed_date, project_equipment.Row_created_by, project_equipment.Row_created_date, project_equipment.Row_effective_date, project_equipment.Row_expiry_date, project_equipment.Row_quality)
	
	if err != nil {
		return err
	}

	return c.JSON(rows)
}

func UpdateProjectEquipment(c *fiber.Ctx) error {
	var project_equipment dto.Project_equipment

	setDefaults(&project_equipment)

	if err := json.Unmarshal(c.Body(), &project_equipment); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	id := c.Params("id")
	project_equipment.Project_id = id

    if project_equipment.Row_created_date != nil {
        return c.Status(400).SendString("Cannot update row_created_date")
    }

	stmt, err := db.Prepare("update project_equipment set equipment_obs_no = :1, active_ind = :2, catalogue_equip_id = :3, effective_date = :4, equipment_id = :5, expiry_date = :6, operated_by_ba_id = :7, ppdm_guid = :8, remark = :9, source = :10, row_changed_by = :11, row_changed_date = :12, row_created_by = :13, row_effective_date = :14, row_expiry_date = :15, row_quality = :16 where project_id = :18")
	if err != nil {
		return err
	}

	project_equipment.Row_changed_date = formatDate(project_equipment.Row_changed_date)
	project_equipment.Effective_date = formatDateString(project_equipment.Effective_date)
	project_equipment.Expiry_date = formatDateString(project_equipment.Expiry_date)
	project_equipment.Row_effective_date = formatDateString(project_equipment.Row_effective_date)
	project_equipment.Row_expiry_date = formatDateString(project_equipment.Row_expiry_date)






	rows, err := stmt.Exec(project_equipment.Equipment_obs_no, project_equipment.Active_ind, project_equipment.Catalogue_equip_id, project_equipment.Effective_date, project_equipment.Equipment_id, project_equipment.Expiry_date, project_equipment.Operated_by_ba_id, project_equipment.Ppdm_guid, project_equipment.Remark, project_equipment.Source, project_equipment.Row_changed_by, project_equipment.Row_changed_date, project_equipment.Row_created_by, project_equipment.Row_effective_date, project_equipment.Row_expiry_date, project_equipment.Row_quality, project_equipment.Project_id)
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

func PatchProjectEquipment(c *fiber.Ctx) error {
	id := c.Params("id")

	updatedFields := make(map[string]interface{})

	if err := c.BodyParser(&updatedFields); err != nil {
		return err
	}
	if _, ok := updatedFields["row_created_date"]; ok {
        return c.Status(400).SendString("Cannot update row_created_date field")
    }

	query := "update project_equipment set "
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

func DeleteProjectEquipment(c *fiber.Ctx) error {
	id := c.Params("id")
	var project_equipment dto.Project_equipment
	project_equipment.Project_id = id

	stmt, err := db.Prepare("delete from project_equipment where project_id = :1")
	if err != nil {
		return err
	}

	rows, err := stmt.Exec(project_equipment.Project_id)
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


