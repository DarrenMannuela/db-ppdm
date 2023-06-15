package apiv1

import (
	dto "github.com/DarrenMannuela/pt_gtn_bibliography/dto"
	"github.com/gofiber/fiber/v2"
	"strconv"
	"encoding/json"
)


func GetWellLogAxisCoord(c *fiber.Ctx) error {
	rows, err := db.Query("select * from well_log_axis_coord")
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}
	defer rows.Close()
	var result []dto.Well_log_axis_coord

	for rows.Next() {
		var well_log_axis_coord dto.Well_log_axis_coord
		if err := rows.Scan(&well_log_axis_coord.Uwi, &well_log_axis_coord.Curve_id, &well_log_axis_coord.Axis_id, &well_log_axis_coord.Coordinate_seq_no, &well_log_axis_coord.Active_ind, &well_log_axis_coord.Axis_value, &well_log_axis_coord.Axis_value_ouom, &well_log_axis_coord.Axis_value_uom, &well_log_axis_coord.Effective_date, &well_log_axis_coord.Expiry_date, &well_log_axis_coord.Ppdm_guid, &well_log_axis_coord.Remark, &well_log_axis_coord.Source, &well_log_axis_coord.Text_value, &well_log_axis_coord.Row_changed_by, &well_log_axis_coord.Row_changed_date, &well_log_axis_coord.Row_created_by, &well_log_axis_coord.Row_created_date, &well_log_axis_coord.Row_effective_date, &well_log_axis_coord.Row_expiry_date, &well_log_axis_coord.Row_quality); err != nil {
			return err // Exit if we get an error
		}

		// Append well_log_axis_coord to result
		result = append(result, well_log_axis_coord)
	}
	// Return result in JSON format
	return c.JSON(result)
	 
}

func SetWellLogAxisCoord(c *fiber.Ctx) error {
	var well_log_axis_coord dto.Well_log_axis_coord

	setDefaults(&well_log_axis_coord)

	if err := json.Unmarshal(c.Body(), &well_log_axis_coord); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	stmt, err := db.Prepare("insert into well_log_axis_coord values(:1, :2, :3, :4, :5, :6, :7, :8, :9, :10, :11, :12, :13, :14, :15, :16, :17, :18, :19, :20, :21)")
	if err != nil {
		return err
	}
	well_log_axis_coord.Row_created_date = formatDate(well_log_axis_coord.Row_created_date)
	well_log_axis_coord.Row_changed_date = formatDate(well_log_axis_coord.Row_changed_date)
	well_log_axis_coord.Effective_date = formatDateString(well_log_axis_coord.Effective_date)
	well_log_axis_coord.Expiry_date = formatDateString(well_log_axis_coord.Expiry_date)
	well_log_axis_coord.Row_effective_date = formatDateString(well_log_axis_coord.Row_effective_date)
	well_log_axis_coord.Row_expiry_date = formatDateString(well_log_axis_coord.Row_expiry_date)






	rows, err := stmt.Exec(well_log_axis_coord.Uwi, well_log_axis_coord.Curve_id, well_log_axis_coord.Axis_id, well_log_axis_coord.Coordinate_seq_no, well_log_axis_coord.Active_ind, well_log_axis_coord.Axis_value, well_log_axis_coord.Axis_value_ouom, well_log_axis_coord.Axis_value_uom, well_log_axis_coord.Effective_date, well_log_axis_coord.Expiry_date, well_log_axis_coord.Ppdm_guid, well_log_axis_coord.Remark, well_log_axis_coord.Source, well_log_axis_coord.Text_value, well_log_axis_coord.Row_changed_by, well_log_axis_coord.Row_changed_date, well_log_axis_coord.Row_created_by, well_log_axis_coord.Row_created_date, well_log_axis_coord.Row_effective_date, well_log_axis_coord.Row_expiry_date, well_log_axis_coord.Row_quality)
	
	if err != nil {
		return err
	}

	return c.JSON(rows)
}

func UpdateWellLogAxisCoord(c *fiber.Ctx) error {
	var well_log_axis_coord dto.Well_log_axis_coord

	setDefaults(&well_log_axis_coord)

	if err := json.Unmarshal(c.Body(), &well_log_axis_coord); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	id := c.Params("id")
	well_log_axis_coord.Uwi = id

    if well_log_axis_coord.Row_created_date != nil {
        return c.Status(400).SendString("Cannot update row_created_date")
    }

	stmt, err := db.Prepare("update well_log_axis_coord set curve_id = :1, axis_id = :2, coordinate_seq_no = :3, active_ind = :4, axis_value = :5, axis_value_ouom = :6, axis_value_uom = :7, effective_date = :8, expiry_date = :9, ppdm_guid = :10, remark = :11, source = :12, text_value = :13, row_changed_by = :14, row_changed_date = :15, row_created_by = :16, row_effective_date = :17, row_expiry_date = :18, row_quality = :19 where uwi = :21")
	if err != nil {
		return err
	}

	well_log_axis_coord.Row_changed_date = formatDate(well_log_axis_coord.Row_changed_date)
	well_log_axis_coord.Effective_date = formatDateString(well_log_axis_coord.Effective_date)
	well_log_axis_coord.Expiry_date = formatDateString(well_log_axis_coord.Expiry_date)
	well_log_axis_coord.Row_effective_date = formatDateString(well_log_axis_coord.Row_effective_date)
	well_log_axis_coord.Row_expiry_date = formatDateString(well_log_axis_coord.Row_expiry_date)






	rows, err := stmt.Exec(well_log_axis_coord.Curve_id, well_log_axis_coord.Axis_id, well_log_axis_coord.Coordinate_seq_no, well_log_axis_coord.Active_ind, well_log_axis_coord.Axis_value, well_log_axis_coord.Axis_value_ouom, well_log_axis_coord.Axis_value_uom, well_log_axis_coord.Effective_date, well_log_axis_coord.Expiry_date, well_log_axis_coord.Ppdm_guid, well_log_axis_coord.Remark, well_log_axis_coord.Source, well_log_axis_coord.Text_value, well_log_axis_coord.Row_changed_by, well_log_axis_coord.Row_changed_date, well_log_axis_coord.Row_created_by, well_log_axis_coord.Row_effective_date, well_log_axis_coord.Row_expiry_date, well_log_axis_coord.Row_quality, well_log_axis_coord.Uwi)
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

func PatchWellLogAxisCoord(c *fiber.Ctx) error {
	id := c.Params("id")

	updatedFields := make(map[string]interface{})

	if err := c.BodyParser(&updatedFields); err != nil {
		return err
	}
	if _, ok := updatedFields["row_created_date"]; ok {
        return c.Status(400).SendString("Cannot update row_created_date field")
    }

	query := "update well_log_axis_coord set "
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
	query += " where uwi = :id"

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

func DeleteWellLogAxisCoord(c *fiber.Ctx) error {
	id := c.Params("id")
	var well_log_axis_coord dto.Well_log_axis_coord
	well_log_axis_coord.Uwi = id

	stmt, err := db.Prepare("delete from well_log_axis_coord where uwi = :1")
	if err != nil {
		return err
	}

	rows, err := stmt.Exec(well_log_axis_coord.Uwi)
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


