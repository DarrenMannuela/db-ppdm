package apiv1

import (
	dto "github.com/DarrenMannuela/pt_gtn_bibliography/dto"
	"github.com/gofiber/fiber/v2"
	"strconv"
	"encoding/json"
)


func GetWellLogCurveAxis(c *fiber.Ctx) error {
	rows, err := db.Query("select * from well_log_curve_axis")
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}
	defer rows.Close()
	var result []dto.Well_log_curve_axis

	for rows.Next() {
		var well_log_curve_axis dto.Well_log_curve_axis
		if err := rows.Scan(&well_log_curve_axis.Uwi, &well_log_curve_axis.Curve_id, &well_log_curve_axis.Axis_id, &well_log_curve_axis.Active_ind, &well_log_curve_axis.Axis_ouom, &well_log_curve_axis.Axis_seq_no, &well_log_curve_axis.Axis_uom, &well_log_curve_axis.Dimension_count, &well_log_curve_axis.Effective_date, &well_log_curve_axis.Expiry_date, &well_log_curve_axis.Ppdm_guid, &well_log_curve_axis.Remark, &well_log_curve_axis.Reported_axis_name, &well_log_curve_axis.Reported_axis_object_name, &well_log_curve_axis.Source, &well_log_curve_axis.Spacing, &well_log_curve_axis.Spacing_ouom, &well_log_curve_axis.Spacing_uom, &well_log_curve_axis.Value_count, &well_log_curve_axis.Row_changed_by, &well_log_curve_axis.Row_changed_date, &well_log_curve_axis.Row_created_by, &well_log_curve_axis.Row_created_date, &well_log_curve_axis.Row_effective_date, &well_log_curve_axis.Row_expiry_date, &well_log_curve_axis.Row_quality); err != nil {
			return err // Exit if we get an error
		}

		// Append well_log_curve_axis to result
		result = append(result, well_log_curve_axis)
	}
	// Return result in JSON format
	return c.JSON(result)
	 
}

func SetWellLogCurveAxis(c *fiber.Ctx) error {
	var well_log_curve_axis dto.Well_log_curve_axis

	setDefaults(&well_log_curve_axis)

	if err := json.Unmarshal(c.Body(), &well_log_curve_axis); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	stmt, err := db.Prepare("insert into well_log_curve_axis values(:1, :2, :3, :4, :5, :6, :7, :8, :9, :10, :11, :12, :13, :14, :15, :16, :17, :18, :19, :20, :21, :22, :23, :24, :25, :26)")
	if err != nil {
		return err
	}
	well_log_curve_axis.Row_created_date = formatDate(well_log_curve_axis.Row_created_date)
	well_log_curve_axis.Row_changed_date = formatDate(well_log_curve_axis.Row_changed_date)
	well_log_curve_axis.Effective_date = formatDateString(well_log_curve_axis.Effective_date)
	well_log_curve_axis.Expiry_date = formatDateString(well_log_curve_axis.Expiry_date)
	well_log_curve_axis.Row_effective_date = formatDateString(well_log_curve_axis.Row_effective_date)
	well_log_curve_axis.Row_expiry_date = formatDateString(well_log_curve_axis.Row_expiry_date)






	rows, err := stmt.Exec(well_log_curve_axis.Uwi, well_log_curve_axis.Curve_id, well_log_curve_axis.Axis_id, well_log_curve_axis.Active_ind, well_log_curve_axis.Axis_ouom, well_log_curve_axis.Axis_seq_no, well_log_curve_axis.Axis_uom, well_log_curve_axis.Dimension_count, well_log_curve_axis.Effective_date, well_log_curve_axis.Expiry_date, well_log_curve_axis.Ppdm_guid, well_log_curve_axis.Remark, well_log_curve_axis.Reported_axis_name, well_log_curve_axis.Reported_axis_object_name, well_log_curve_axis.Source, well_log_curve_axis.Spacing, well_log_curve_axis.Spacing_ouom, well_log_curve_axis.Spacing_uom, well_log_curve_axis.Value_count, well_log_curve_axis.Row_changed_by, well_log_curve_axis.Row_changed_date, well_log_curve_axis.Row_created_by, well_log_curve_axis.Row_created_date, well_log_curve_axis.Row_effective_date, well_log_curve_axis.Row_expiry_date, well_log_curve_axis.Row_quality)
	
	if err != nil {
		return err
	}

	return c.JSON(rows)
}

func UpdateWellLogCurveAxis(c *fiber.Ctx) error {
	var well_log_curve_axis dto.Well_log_curve_axis

	setDefaults(&well_log_curve_axis)

	if err := json.Unmarshal(c.Body(), &well_log_curve_axis); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	id := c.Params("id")
	well_log_curve_axis.Uwi = id

    if well_log_curve_axis.Row_created_date != nil {
        return c.Status(400).SendString("Cannot update row_created_date")
    }

	stmt, err := db.Prepare("update well_log_curve_axis set curve_id = :1, axis_id = :2, active_ind = :3, axis_ouom = :4, axis_seq_no = :5, axis_uom = :6, dimension_count = :7, effective_date = :8, expiry_date = :9, ppdm_guid = :10, remark = :11, reported_axis_name = :12, reported_axis_object_name = :13, source = :14, spacing = :15, spacing_ouom = :16, spacing_uom = :17, value_count = :18, row_changed_by = :19, row_changed_date = :20, row_created_by = :21, row_effective_date = :22, row_expiry_date = :23, row_quality = :24 where uwi = :26")
	if err != nil {
		return err
	}

	well_log_curve_axis.Row_changed_date = formatDate(well_log_curve_axis.Row_changed_date)
	well_log_curve_axis.Effective_date = formatDateString(well_log_curve_axis.Effective_date)
	well_log_curve_axis.Expiry_date = formatDateString(well_log_curve_axis.Expiry_date)
	well_log_curve_axis.Row_effective_date = formatDateString(well_log_curve_axis.Row_effective_date)
	well_log_curve_axis.Row_expiry_date = formatDateString(well_log_curve_axis.Row_expiry_date)






	rows, err := stmt.Exec(well_log_curve_axis.Curve_id, well_log_curve_axis.Axis_id, well_log_curve_axis.Active_ind, well_log_curve_axis.Axis_ouom, well_log_curve_axis.Axis_seq_no, well_log_curve_axis.Axis_uom, well_log_curve_axis.Dimension_count, well_log_curve_axis.Effective_date, well_log_curve_axis.Expiry_date, well_log_curve_axis.Ppdm_guid, well_log_curve_axis.Remark, well_log_curve_axis.Reported_axis_name, well_log_curve_axis.Reported_axis_object_name, well_log_curve_axis.Source, well_log_curve_axis.Spacing, well_log_curve_axis.Spacing_ouom, well_log_curve_axis.Spacing_uom, well_log_curve_axis.Value_count, well_log_curve_axis.Row_changed_by, well_log_curve_axis.Row_changed_date, well_log_curve_axis.Row_created_by, well_log_curve_axis.Row_effective_date, well_log_curve_axis.Row_expiry_date, well_log_curve_axis.Row_quality, well_log_curve_axis.Uwi)
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

func PatchWellLogCurveAxis(c *fiber.Ctx) error {
	id := c.Params("id")

	updatedFields := make(map[string]interface{})

	if err := c.BodyParser(&updatedFields); err != nil {
		return err
	}
	if _, ok := updatedFields["row_created_date"]; ok {
        return c.Status(400).SendString("Cannot update row_created_date field")
    }

	query := "update well_log_curve_axis set "
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

func DeleteWellLogCurveAxis(c *fiber.Ctx) error {
	id := c.Params("id")
	var well_log_curve_axis dto.Well_log_curve_axis
	well_log_curve_axis.Uwi = id

	stmt, err := db.Prepare("delete from well_log_curve_axis where uwi = :1")
	if err != nil {
		return err
	}

	rows, err := stmt.Exec(well_log_curve_axis.Uwi)
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


