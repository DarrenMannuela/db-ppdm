package apiv1

import (
	dto "github.com/DarrenMannuela/pt_gtn_bibliography/dto"
	"github.com/gofiber/fiber/v2"
	"strconv"
	"encoding/json"
)


func GetRmImageLoc(c *fiber.Ctx) error {
	rows, err := db.Query("select * from rm_image_loc")
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}
	defer rows.Close()
	var result []dto.Rm_image_loc

	for rows.Next() {
		var rm_image_loc dto.Rm_image_loc
		if err := rows.Scan(&rm_image_loc.Physical_item_id, &rm_image_loc.Log_section_id, &rm_image_loc.Position_id, &rm_image_loc.Active_ind, &rm_image_loc.Effective_date, &rm_image_loc.Expiry_date, &rm_image_loc.Log_depth, &rm_image_loc.Log_depth_ouom, &rm_image_loc.Log_depth_uom, &rm_image_loc.Position_type, &rm_image_loc.Ppdm_guid, &rm_image_loc.Remark, &rm_image_loc.Source, &rm_image_loc.X_position, &rm_image_loc.Y_position, &rm_image_loc.Row_changed_by, &rm_image_loc.Row_changed_date, &rm_image_loc.Row_created_by, &rm_image_loc.Row_created_date, &rm_image_loc.Row_effective_date, &rm_image_loc.Row_expiry_date, &rm_image_loc.Row_quality); err != nil {
			return err // Exit if we get an error
		}

		// Append rm_image_loc to result
		result = append(result, rm_image_loc)
	}
	// Return result in JSON format
	return c.JSON(result)
	 
}

func SetRmImageLoc(c *fiber.Ctx) error {
	var rm_image_loc dto.Rm_image_loc

	setDefaults(&rm_image_loc)

	if err := json.Unmarshal(c.Body(), &rm_image_loc); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	stmt, err := db.Prepare("insert into rm_image_loc values(:1, :2, :3, :4, :5, :6, :7, :8, :9, :10, :11, :12, :13, :14, :15, :16, :17, :18, :19, :20, :21, :22)")
	if err != nil {
		return err
	}
	rm_image_loc.Row_created_date = formatDate(rm_image_loc.Row_created_date)
	rm_image_loc.Row_changed_date = formatDate(rm_image_loc.Row_changed_date)
	rm_image_loc.Effective_date = formatDateString(rm_image_loc.Effective_date)
	rm_image_loc.Expiry_date = formatDateString(rm_image_loc.Expiry_date)
	rm_image_loc.Row_effective_date = formatDateString(rm_image_loc.Row_effective_date)
	rm_image_loc.Row_expiry_date = formatDateString(rm_image_loc.Row_expiry_date)






	rows, err := stmt.Exec(rm_image_loc.Physical_item_id, rm_image_loc.Log_section_id, rm_image_loc.Position_id, rm_image_loc.Active_ind, rm_image_loc.Effective_date, rm_image_loc.Expiry_date, rm_image_loc.Log_depth, rm_image_loc.Log_depth_ouom, rm_image_loc.Log_depth_uom, rm_image_loc.Position_type, rm_image_loc.Ppdm_guid, rm_image_loc.Remark, rm_image_loc.Source, rm_image_loc.X_position, rm_image_loc.Y_position, rm_image_loc.Row_changed_by, rm_image_loc.Row_changed_date, rm_image_loc.Row_created_by, rm_image_loc.Row_created_date, rm_image_loc.Row_effective_date, rm_image_loc.Row_expiry_date, rm_image_loc.Row_quality)
	
	if err != nil {
		return err
	}

	return c.JSON(rows)
}

func UpdateRmImageLoc(c *fiber.Ctx) error {
	var rm_image_loc dto.Rm_image_loc

	setDefaults(&rm_image_loc)

	if err := json.Unmarshal(c.Body(), &rm_image_loc); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	id := c.Params("id")
	rm_image_loc.Physical_item_id = id

    if rm_image_loc.Row_created_date != nil {
        return c.Status(400).SendString("Cannot update row_created_date")
    }

	stmt, err := db.Prepare("update rm_image_loc set log_section_id = :1, position_id = :2, active_ind = :3, effective_date = :4, expiry_date = :5, log_depth = :6, log_depth_ouom = :7, log_depth_uom = :8, position_type = :9, ppdm_guid = :10, remark = :11, source = :12, x_position = :13, y_position = :14, row_changed_by = :15, row_changed_date = :16, row_created_by = :17, row_effective_date = :18, row_expiry_date = :19, row_quality = :20 where physical_item_id = :22")
	if err != nil {
		return err
	}

	rm_image_loc.Row_changed_date = formatDate(rm_image_loc.Row_changed_date)
	rm_image_loc.Effective_date = formatDateString(rm_image_loc.Effective_date)
	rm_image_loc.Expiry_date = formatDateString(rm_image_loc.Expiry_date)
	rm_image_loc.Row_effective_date = formatDateString(rm_image_loc.Row_effective_date)
	rm_image_loc.Row_expiry_date = formatDateString(rm_image_loc.Row_expiry_date)






	rows, err := stmt.Exec(rm_image_loc.Log_section_id, rm_image_loc.Position_id, rm_image_loc.Active_ind, rm_image_loc.Effective_date, rm_image_loc.Expiry_date, rm_image_loc.Log_depth, rm_image_loc.Log_depth_ouom, rm_image_loc.Log_depth_uom, rm_image_loc.Position_type, rm_image_loc.Ppdm_guid, rm_image_loc.Remark, rm_image_loc.Source, rm_image_loc.X_position, rm_image_loc.Y_position, rm_image_loc.Row_changed_by, rm_image_loc.Row_changed_date, rm_image_loc.Row_created_by, rm_image_loc.Row_effective_date, rm_image_loc.Row_expiry_date, rm_image_loc.Row_quality, rm_image_loc.Physical_item_id)
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

func PatchRmImageLoc(c *fiber.Ctx) error {
	id := c.Params("id")

	updatedFields := make(map[string]interface{})

	if err := c.BodyParser(&updatedFields); err != nil {
		return err
	}
	if _, ok := updatedFields["row_created_date"]; ok {
        return c.Status(400).SendString("Cannot update row_created_date field")
    }

	query := "update rm_image_loc set "
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
	query += " where physical_item_id = :id"

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

func DeleteRmImageLoc(c *fiber.Ctx) error {
	id := c.Params("id")
	var rm_image_loc dto.Rm_image_loc
	rm_image_loc.Physical_item_id = id

	stmt, err := db.Prepare("delete from rm_image_loc where physical_item_id = :1")
	if err != nil {
		return err
	}

	rows, err := stmt.Exec(rm_image_loc.Physical_item_id)
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


