package apiv1

import (
	dto "github.com/DarrenMannuela/pt_gtn_bibliography/dto"
	"github.com/gofiber/fiber/v2"
	"strconv"
	"encoding/json"
)


func GetSfDock(c *fiber.Ctx) error {
	rows, err := db.Query("select * from sf_dock")
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}
	defer rows.Close()
	var result []dto.Sf_dock

	for rows.Next() {
		var sf_dock dto.Sf_dock
		if err := rows.Scan(&sf_dock.Sf_subtype, &sf_dock.Dock_id, &sf_dock.Active_ind, &sf_dock.Dock_length, &sf_dock.Dock_length_ouom, &sf_dock.Dock_type, &sf_dock.Effective_date, &sf_dock.Expiry_date, &sf_dock.Ppdm_guid, &sf_dock.Remark, &sf_dock.Source, &sf_dock.Surface_type, &sf_dock.Water_depth, &sf_dock.Water_depth_ouom, &sf_dock.Row_changed_by, &sf_dock.Row_changed_date, &sf_dock.Row_created_by, &sf_dock.Row_created_date, &sf_dock.Row_effective_date, &sf_dock.Row_expiry_date, &sf_dock.Row_quality); err != nil {
			return err // Exit if we get an error
		}

		// Append sf_dock to result
		result = append(result, sf_dock)
	}
	// Return result in JSON format
	return c.JSON(result)
	 
}

func SetSfDock(c *fiber.Ctx) error {
	var sf_dock dto.Sf_dock

	setDefaults(&sf_dock)

	if err := json.Unmarshal(c.Body(), &sf_dock); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	stmt, err := db.Prepare("insert into sf_dock values(:1, :2, :3, :4, :5, :6, :7, :8, :9, :10, :11, :12, :13, :14, :15, :16, :17, :18, :19, :20, :21)")
	if err != nil {
		return err
	}
	sf_dock.Row_created_date = formatDate(sf_dock.Row_created_date)
	sf_dock.Row_changed_date = formatDate(sf_dock.Row_changed_date)
	sf_dock.Effective_date = formatDateString(sf_dock.Effective_date)
	sf_dock.Expiry_date = formatDateString(sf_dock.Expiry_date)
	sf_dock.Row_effective_date = formatDateString(sf_dock.Row_effective_date)
	sf_dock.Row_expiry_date = formatDateString(sf_dock.Row_expiry_date)






	rows, err := stmt.Exec(sf_dock.Sf_subtype, sf_dock.Dock_id, sf_dock.Active_ind, sf_dock.Dock_length, sf_dock.Dock_length_ouom, sf_dock.Dock_type, sf_dock.Effective_date, sf_dock.Expiry_date, sf_dock.Ppdm_guid, sf_dock.Remark, sf_dock.Source, sf_dock.Surface_type, sf_dock.Water_depth, sf_dock.Water_depth_ouom, sf_dock.Row_changed_by, sf_dock.Row_changed_date, sf_dock.Row_created_by, sf_dock.Row_created_date, sf_dock.Row_effective_date, sf_dock.Row_expiry_date, sf_dock.Row_quality)
	
	if err != nil {
		return err
	}

	return c.JSON(rows)
}

func UpdateSfDock(c *fiber.Ctx) error {
	var sf_dock dto.Sf_dock

	setDefaults(&sf_dock)

	if err := json.Unmarshal(c.Body(), &sf_dock); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	id := c.Params("id")
	sf_dock.Sf_subtype = id

    if sf_dock.Row_created_date != nil {
        return c.Status(400).SendString("Cannot update row_created_date")
    }

	stmt, err := db.Prepare("update sf_dock set dock_id = :1, active_ind = :2, dock_length = :3, dock_length_ouom = :4, dock_type = :5, effective_date = :6, expiry_date = :7, ppdm_guid = :8, remark = :9, source = :10, surface_type = :11, water_depth = :12, water_depth_ouom = :13, row_changed_by = :14, row_changed_date = :15, row_created_by = :16, row_effective_date = :17, row_expiry_date = :18, row_quality = :19 where sf_subtype = :21")
	if err != nil {
		return err
	}

	sf_dock.Row_changed_date = formatDate(sf_dock.Row_changed_date)
	sf_dock.Effective_date = formatDateString(sf_dock.Effective_date)
	sf_dock.Expiry_date = formatDateString(sf_dock.Expiry_date)
	sf_dock.Row_effective_date = formatDateString(sf_dock.Row_effective_date)
	sf_dock.Row_expiry_date = formatDateString(sf_dock.Row_expiry_date)






	rows, err := stmt.Exec(sf_dock.Dock_id, sf_dock.Active_ind, sf_dock.Dock_length, sf_dock.Dock_length_ouom, sf_dock.Dock_type, sf_dock.Effective_date, sf_dock.Expiry_date, sf_dock.Ppdm_guid, sf_dock.Remark, sf_dock.Source, sf_dock.Surface_type, sf_dock.Water_depth, sf_dock.Water_depth_ouom, sf_dock.Row_changed_by, sf_dock.Row_changed_date, sf_dock.Row_created_by, sf_dock.Row_effective_date, sf_dock.Row_expiry_date, sf_dock.Row_quality, sf_dock.Sf_subtype)
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

func PatchSfDock(c *fiber.Ctx) error {
	id := c.Params("id")

	updatedFields := make(map[string]interface{})

	if err := c.BodyParser(&updatedFields); err != nil {
		return err
	}
	if _, ok := updatedFields["row_created_date"]; ok {
        return c.Status(400).SendString("Cannot update row_created_date field")
    }

	query := "update sf_dock set "
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
	query += " where sf_subtype = :id"

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

func DeleteSfDock(c *fiber.Ctx) error {
	id := c.Params("id")
	var sf_dock dto.Sf_dock
	sf_dock.Sf_subtype = id

	stmt, err := db.Prepare("delete from sf_dock where sf_subtype = :1")
	if err != nil {
		return err
	}

	rows, err := stmt.Exec(sf_dock.Sf_subtype)
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


