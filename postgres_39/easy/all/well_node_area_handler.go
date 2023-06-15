package apiv1

import (
	dto "github.com/DarrenMannuela/pt_gtn_bibliography/dto"
	"github.com/gofiber/fiber/v2"
	"strconv"
	"encoding/json"
)


func GetWellNodeArea(c *fiber.Ctx) error {
	rows, err := db.Query("select * from well_node_area")
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}
	defer rows.Close()
	var result []dto.Well_node_area

	for rows.Next() {
		var well_node_area dto.Well_node_area
		if err := rows.Scan(&well_node_area.Node_id, &well_node_area.Source, &well_node_area.Area_id, &well_node_area.Area_type, &well_node_area.Active_ind, &well_node_area.Effective_date, &well_node_area.Expiry_date, &well_node_area.Ppdm_guid, &well_node_area.Remark, &well_node_area.Row_changed_by, &well_node_area.Row_changed_date, &well_node_area.Row_created_by, &well_node_area.Row_created_date, &well_node_area.Row_effective_date, &well_node_area.Row_expiry_date, &well_node_area.Row_quality); err != nil {
			return err // Exit if we get an error
		}

		// Append well_node_area to result
		result = append(result, well_node_area)
	}
	// Return result in JSON format
	return c.JSON(result)
	 
}

func SetWellNodeArea(c *fiber.Ctx) error {
	var well_node_area dto.Well_node_area

	setDefaults(&well_node_area)

	if err := json.Unmarshal(c.Body(), &well_node_area); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	stmt, err := db.Prepare("insert into well_node_area values(:1, :2, :3, :4, :5, :6, :7, :8, :9, :10, :11, :12, :13, :14, :15, :16)")
	if err != nil {
		return err
	}
	well_node_area.Row_created_date = formatDate(well_node_area.Row_created_date)
	well_node_area.Row_changed_date = formatDate(well_node_area.Row_changed_date)
	well_node_area.Effective_date = formatDateString(well_node_area.Effective_date)
	well_node_area.Expiry_date = formatDateString(well_node_area.Expiry_date)
	well_node_area.Row_effective_date = formatDateString(well_node_area.Row_effective_date)
	well_node_area.Row_expiry_date = formatDateString(well_node_area.Row_expiry_date)






	rows, err := stmt.Exec(well_node_area.Node_id, well_node_area.Source, well_node_area.Area_id, well_node_area.Area_type, well_node_area.Active_ind, well_node_area.Effective_date, well_node_area.Expiry_date, well_node_area.Ppdm_guid, well_node_area.Remark, well_node_area.Row_changed_by, well_node_area.Row_changed_date, well_node_area.Row_created_by, well_node_area.Row_created_date, well_node_area.Row_effective_date, well_node_area.Row_expiry_date, well_node_area.Row_quality)
	
	if err != nil {
		return err
	}

	return c.JSON(rows)
}

func UpdateWellNodeArea(c *fiber.Ctx) error {
	var well_node_area dto.Well_node_area

	setDefaults(&well_node_area)

	if err := json.Unmarshal(c.Body(), &well_node_area); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	id := c.Params("id")
	well_node_area.Node_id = id

    if well_node_area.Row_created_date != nil {
        return c.Status(400).SendString("Cannot update row_created_date")
    }

	stmt, err := db.Prepare("update well_node_area set source = :1, area_id = :2, area_type = :3, active_ind = :4, effective_date = :5, expiry_date = :6, ppdm_guid = :7, remark = :8, row_changed_by = :9, row_changed_date = :10, row_created_by = :11, row_effective_date = :12, row_expiry_date = :13, row_quality = :14 where node_id = :16")
	if err != nil {
		return err
	}

	well_node_area.Row_changed_date = formatDate(well_node_area.Row_changed_date)
	well_node_area.Effective_date = formatDateString(well_node_area.Effective_date)
	well_node_area.Expiry_date = formatDateString(well_node_area.Expiry_date)
	well_node_area.Row_effective_date = formatDateString(well_node_area.Row_effective_date)
	well_node_area.Row_expiry_date = formatDateString(well_node_area.Row_expiry_date)






	rows, err := stmt.Exec(well_node_area.Source, well_node_area.Area_id, well_node_area.Area_type, well_node_area.Active_ind, well_node_area.Effective_date, well_node_area.Expiry_date, well_node_area.Ppdm_guid, well_node_area.Remark, well_node_area.Row_changed_by, well_node_area.Row_changed_date, well_node_area.Row_created_by, well_node_area.Row_effective_date, well_node_area.Row_expiry_date, well_node_area.Row_quality, well_node_area.Node_id)
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

func PatchWellNodeArea(c *fiber.Ctx) error {
	id := c.Params("id")

	updatedFields := make(map[string]interface{})

	if err := c.BodyParser(&updatedFields); err != nil {
		return err
	}
	if _, ok := updatedFields["row_created_date"]; ok {
        return c.Status(400).SendString("Cannot update row_created_date field")
    }

	query := "update well_node_area set "
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
	query += " where node_id = :id"

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

func DeleteWellNodeArea(c *fiber.Ctx) error {
	id := c.Params("id")
	var well_node_area dto.Well_node_area
	well_node_area.Node_id = id

	stmt, err := db.Prepare("delete from well_node_area where node_id = :1")
	if err != nil {
		return err
	}

	rows, err := stmt.Exec(well_node_area.Node_id)
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


