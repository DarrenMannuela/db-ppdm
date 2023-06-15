package apiv1

import (
	dto "github.com/DarrenMannuela/pt_gtn_bibliography/dto"
	"github.com/gofiber/fiber/v2"
	"strconv"
	"encoding/json"
)


func GetWellLogDgtzCurve(c *fiber.Ctx) error {
	rows, err := db.Query("select * from well_log_dgtz_curve")
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}
	defer rows.Close()
	var result []dto.Well_log_dgtz_curve

	for rows.Next() {
		var well_log_dgtz_curve dto.Well_log_dgtz_curve
		if err := rows.Scan(&well_log_dgtz_curve.Uwi, &well_log_dgtz_curve.Curve_id, &well_log_dgtz_curve.Digital_curve_id, &well_log_dgtz_curve.Active_ind, &well_log_dgtz_curve.Base_depth, &well_log_dgtz_curve.Base_depth_ouom, &well_log_dgtz_curve.Curve_quality, &well_log_dgtz_curve.Depth_correction_method, &well_log_dgtz_curve.Depth_increment, &well_log_dgtz_curve.Depth_increment_ouom, &well_log_dgtz_curve.Digitized_date, &well_log_dgtz_curve.Digitizer, &well_log_dgtz_curve.Effective_date, &well_log_dgtz_curve.Expiry_date, &well_log_dgtz_curve.Null_value, &well_log_dgtz_curve.Ppdm_guid, &well_log_dgtz_curve.Remark, &well_log_dgtz_curve.Source, &well_log_dgtz_curve.Top_depth, &well_log_dgtz_curve.Top_depth_ouom, &well_log_dgtz_curve.Row_changed_by, &well_log_dgtz_curve.Row_changed_date, &well_log_dgtz_curve.Row_created_by, &well_log_dgtz_curve.Row_created_date, &well_log_dgtz_curve.Row_effective_date, &well_log_dgtz_curve.Row_expiry_date, &well_log_dgtz_curve.Row_quality); err != nil {
			return err // Exit if we get an error
		}

		// Append well_log_dgtz_curve to result
		result = append(result, well_log_dgtz_curve)
	}
	// Return result in JSON format
	return c.JSON(result)
	 
}

func SetWellLogDgtzCurve(c *fiber.Ctx) error {
	var well_log_dgtz_curve dto.Well_log_dgtz_curve

	setDefaults(&well_log_dgtz_curve)

	if err := json.Unmarshal(c.Body(), &well_log_dgtz_curve); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	stmt, err := db.Prepare("insert into well_log_dgtz_curve values(:1, :2, :3, :4, :5, :6, :7, :8, :9, :10, :11, :12, :13, :14, :15, :16, :17, :18, :19, :20, :21, :22, :23, :24, :25, :26, :27)")
	if err != nil {
		return err
	}
	well_log_dgtz_curve.Row_created_date = formatDate(well_log_dgtz_curve.Row_created_date)
	well_log_dgtz_curve.Row_changed_date = formatDate(well_log_dgtz_curve.Row_changed_date)
	well_log_dgtz_curve.Digitized_date = formatDateString(well_log_dgtz_curve.Digitized_date)
	well_log_dgtz_curve.Effective_date = formatDateString(well_log_dgtz_curve.Effective_date)
	well_log_dgtz_curve.Expiry_date = formatDateString(well_log_dgtz_curve.Expiry_date)
	well_log_dgtz_curve.Row_effective_date = formatDateString(well_log_dgtz_curve.Row_effective_date)
	well_log_dgtz_curve.Row_expiry_date = formatDateString(well_log_dgtz_curve.Row_expiry_date)






	rows, err := stmt.Exec(well_log_dgtz_curve.Uwi, well_log_dgtz_curve.Curve_id, well_log_dgtz_curve.Digital_curve_id, well_log_dgtz_curve.Active_ind, well_log_dgtz_curve.Base_depth, well_log_dgtz_curve.Base_depth_ouom, well_log_dgtz_curve.Curve_quality, well_log_dgtz_curve.Depth_correction_method, well_log_dgtz_curve.Depth_increment, well_log_dgtz_curve.Depth_increment_ouom, well_log_dgtz_curve.Digitized_date, well_log_dgtz_curve.Digitizer, well_log_dgtz_curve.Effective_date, well_log_dgtz_curve.Expiry_date, well_log_dgtz_curve.Null_value, well_log_dgtz_curve.Ppdm_guid, well_log_dgtz_curve.Remark, well_log_dgtz_curve.Source, well_log_dgtz_curve.Top_depth, well_log_dgtz_curve.Top_depth_ouom, well_log_dgtz_curve.Row_changed_by, well_log_dgtz_curve.Row_changed_date, well_log_dgtz_curve.Row_created_by, well_log_dgtz_curve.Row_created_date, well_log_dgtz_curve.Row_effective_date, well_log_dgtz_curve.Row_expiry_date, well_log_dgtz_curve.Row_quality)
	
	if err != nil {
		return err
	}

	return c.JSON(rows)
}

func UpdateWellLogDgtzCurve(c *fiber.Ctx) error {
	var well_log_dgtz_curve dto.Well_log_dgtz_curve

	setDefaults(&well_log_dgtz_curve)

	if err := json.Unmarshal(c.Body(), &well_log_dgtz_curve); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	id := c.Params("id")
	well_log_dgtz_curve.Uwi = id

    if well_log_dgtz_curve.Row_created_date != nil {
        return c.Status(400).SendString("Cannot update row_created_date")
    }

	stmt, err := db.Prepare("update well_log_dgtz_curve set curve_id = :1, digital_curve_id = :2, active_ind = :3, base_depth = :4, base_depth_ouom = :5, curve_quality = :6, depth_correction_method = :7, depth_increment = :8, depth_increment_ouom = :9, digitized_date = :10, digitizer = :11, effective_date = :12, expiry_date = :13, null_value = :14, ppdm_guid = :15, remark = :16, source = :17, top_depth = :18, top_depth_ouom = :19, row_changed_by = :20, row_changed_date = :21, row_created_by = :22, row_effective_date = :23, row_expiry_date = :24, row_quality = :25 where uwi = :27")
	if err != nil {
		return err
	}

	well_log_dgtz_curve.Row_changed_date = formatDate(well_log_dgtz_curve.Row_changed_date)
	well_log_dgtz_curve.Digitized_date = formatDateString(well_log_dgtz_curve.Digitized_date)
	well_log_dgtz_curve.Effective_date = formatDateString(well_log_dgtz_curve.Effective_date)
	well_log_dgtz_curve.Expiry_date = formatDateString(well_log_dgtz_curve.Expiry_date)
	well_log_dgtz_curve.Row_effective_date = formatDateString(well_log_dgtz_curve.Row_effective_date)
	well_log_dgtz_curve.Row_expiry_date = formatDateString(well_log_dgtz_curve.Row_expiry_date)






	rows, err := stmt.Exec(well_log_dgtz_curve.Curve_id, well_log_dgtz_curve.Digital_curve_id, well_log_dgtz_curve.Active_ind, well_log_dgtz_curve.Base_depth, well_log_dgtz_curve.Base_depth_ouom, well_log_dgtz_curve.Curve_quality, well_log_dgtz_curve.Depth_correction_method, well_log_dgtz_curve.Depth_increment, well_log_dgtz_curve.Depth_increment_ouom, well_log_dgtz_curve.Digitized_date, well_log_dgtz_curve.Digitizer, well_log_dgtz_curve.Effective_date, well_log_dgtz_curve.Expiry_date, well_log_dgtz_curve.Null_value, well_log_dgtz_curve.Ppdm_guid, well_log_dgtz_curve.Remark, well_log_dgtz_curve.Source, well_log_dgtz_curve.Top_depth, well_log_dgtz_curve.Top_depth_ouom, well_log_dgtz_curve.Row_changed_by, well_log_dgtz_curve.Row_changed_date, well_log_dgtz_curve.Row_created_by, well_log_dgtz_curve.Row_effective_date, well_log_dgtz_curve.Row_expiry_date, well_log_dgtz_curve.Row_quality, well_log_dgtz_curve.Uwi)
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

func PatchWellLogDgtzCurve(c *fiber.Ctx) error {
	id := c.Params("id")

	updatedFields := make(map[string]interface{})

	if err := c.BodyParser(&updatedFields); err != nil {
		return err
	}
	if _, ok := updatedFields["row_created_date"]; ok {
        return c.Status(400).SendString("Cannot update row_created_date field")
    }

	query := "update well_log_dgtz_curve set "
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
		} else if key == "digitized_date" || key == "effective_date" || key == "expiry_date" || key == "row_effective_date" || key == "row_expiry_date"      {
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

func DeleteWellLogDgtzCurve(c *fiber.Ctx) error {
	id := c.Params("id")
	var well_log_dgtz_curve dto.Well_log_dgtz_curve
	well_log_dgtz_curve.Uwi = id

	stmt, err := db.Prepare("delete from well_log_dgtz_curve where uwi = :1")
	if err != nil {
		return err
	}

	rows, err := stmt.Exec(well_log_dgtz_curve.Uwi)
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


