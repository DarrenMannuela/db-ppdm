package apiv1

import (
	dto "github.com/DarrenMannuela/pt_gtn_bibliography/dto"
	"github.com/gofiber/fiber/v2"
	"strconv"
	"encoding/json"
)


func GetWellLogCurveValue(c *fiber.Ctx) error {
	rows, err := db.Query("select * from well_log_curve_value")
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}
	defer rows.Close()
	var result []dto.Well_log_curve_value

	for rows.Next() {
		var well_log_curve_value dto.Well_log_curve_value
		if err := rows.Scan(&well_log_curve_value.Uwi, &well_log_curve_value.Curve_id, &well_log_curve_value.Sample_id, &well_log_curve_value.Active_ind, &well_log_curve_value.Effective_date, &well_log_curve_value.Expiry_date, &well_log_curve_value.Index_type, &well_log_curve_value.Index_value, &well_log_curve_value.Index_value_uom, &well_log_curve_value.Measured_value, &well_log_curve_value.Measured_value_uom, &well_log_curve_value.Ppdm_guid, &well_log_curve_value.Remark, &well_log_curve_value.Source, &well_log_curve_value.Row_changed_by, &well_log_curve_value.Row_changed_date, &well_log_curve_value.Row_created_by, &well_log_curve_value.Row_created_date, &well_log_curve_value.Row_effective_date, &well_log_curve_value.Row_expiry_date, &well_log_curve_value.Row_quality); err != nil {
			return err // Exit if we get an error
		}

		// Append well_log_curve_value to result
		result = append(result, well_log_curve_value)
	}
	// Return result in JSON format
	return c.JSON(result)
	 
}

func SetWellLogCurveValue(c *fiber.Ctx) error {
	var well_log_curve_value dto.Well_log_curve_value

	setDefaults(&well_log_curve_value)

	if err := json.Unmarshal(c.Body(), &well_log_curve_value); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	stmt, err := db.Prepare("insert into well_log_curve_value values(:1, :2, :3, :4, :5, :6, :7, :8, :9, :10, :11, :12, :13, :14, :15, :16, :17, :18, :19, :20, :21)")
	if err != nil {
		return err
	}
	well_log_curve_value.Row_created_date = formatDate(well_log_curve_value.Row_created_date)
	well_log_curve_value.Row_changed_date = formatDate(well_log_curve_value.Row_changed_date)
	well_log_curve_value.Effective_date = formatDateString(well_log_curve_value.Effective_date)
	well_log_curve_value.Expiry_date = formatDateString(well_log_curve_value.Expiry_date)
	well_log_curve_value.Row_effective_date = formatDateString(well_log_curve_value.Row_effective_date)
	well_log_curve_value.Row_expiry_date = formatDateString(well_log_curve_value.Row_expiry_date)






	rows, err := stmt.Exec(well_log_curve_value.Uwi, well_log_curve_value.Curve_id, well_log_curve_value.Sample_id, well_log_curve_value.Active_ind, well_log_curve_value.Effective_date, well_log_curve_value.Expiry_date, well_log_curve_value.Index_type, well_log_curve_value.Index_value, well_log_curve_value.Index_value_uom, well_log_curve_value.Measured_value, well_log_curve_value.Measured_value_uom, well_log_curve_value.Ppdm_guid, well_log_curve_value.Remark, well_log_curve_value.Source, well_log_curve_value.Row_changed_by, well_log_curve_value.Row_changed_date, well_log_curve_value.Row_created_by, well_log_curve_value.Row_created_date, well_log_curve_value.Row_effective_date, well_log_curve_value.Row_expiry_date, well_log_curve_value.Row_quality)
	
	if err != nil {
		return err
	}

	return c.JSON(rows)
}

func UpdateWellLogCurveValue(c *fiber.Ctx) error {
	var well_log_curve_value dto.Well_log_curve_value

	setDefaults(&well_log_curve_value)

	if err := json.Unmarshal(c.Body(), &well_log_curve_value); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	id := c.Params("id")
	well_log_curve_value.Uwi = id

    if well_log_curve_value.Row_created_date != nil {
        return c.Status(400).SendString("Cannot update row_created_date")
    }

	stmt, err := db.Prepare("update well_log_curve_value set curve_id = :1, sample_id = :2, active_ind = :3, effective_date = :4, expiry_date = :5, index_type = :6, index_value = :7, index_value_uom = :8, measured_value = :9, measured_value_uom = :10, ppdm_guid = :11, remark = :12, source = :13, row_changed_by = :14, row_changed_date = :15, row_created_by = :16, row_effective_date = :17, row_expiry_date = :18, row_quality = :19 where uwi = :21")
	if err != nil {
		return err
	}

	well_log_curve_value.Row_changed_date = formatDate(well_log_curve_value.Row_changed_date)
	well_log_curve_value.Effective_date = formatDateString(well_log_curve_value.Effective_date)
	well_log_curve_value.Expiry_date = formatDateString(well_log_curve_value.Expiry_date)
	well_log_curve_value.Row_effective_date = formatDateString(well_log_curve_value.Row_effective_date)
	well_log_curve_value.Row_expiry_date = formatDateString(well_log_curve_value.Row_expiry_date)






	rows, err := stmt.Exec(well_log_curve_value.Curve_id, well_log_curve_value.Sample_id, well_log_curve_value.Active_ind, well_log_curve_value.Effective_date, well_log_curve_value.Expiry_date, well_log_curve_value.Index_type, well_log_curve_value.Index_value, well_log_curve_value.Index_value_uom, well_log_curve_value.Measured_value, well_log_curve_value.Measured_value_uom, well_log_curve_value.Ppdm_guid, well_log_curve_value.Remark, well_log_curve_value.Source, well_log_curve_value.Row_changed_by, well_log_curve_value.Row_changed_date, well_log_curve_value.Row_created_by, well_log_curve_value.Row_effective_date, well_log_curve_value.Row_expiry_date, well_log_curve_value.Row_quality, well_log_curve_value.Uwi)
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

func PatchWellLogCurveValue(c *fiber.Ctx) error {
	id := c.Params("id")

	updatedFields := make(map[string]interface{})

	if err := c.BodyParser(&updatedFields); err != nil {
		return err
	}
	if _, ok := updatedFields["row_created_date"]; ok {
        return c.Status(400).SendString("Cannot update row_created_date field")
    }

	query := "update well_log_curve_value set "
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

func DeleteWellLogCurveValue(c *fiber.Ctx) error {
	id := c.Params("id")
	var well_log_curve_value dto.Well_log_curve_value
	well_log_curve_value.Uwi = id

	stmt, err := db.Prepare("delete from well_log_curve_value where uwi = :1")
	if err != nil {
		return err
	}

	rows, err := stmt.Exec(well_log_curve_value.Uwi)
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


