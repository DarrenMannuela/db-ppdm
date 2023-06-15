package apiv1

import (
	dto "github.com/DarrenMannuela/pt_gtn_bibliography/dto"
	"github.com/gofiber/fiber/v2"
	"strconv"
	"encoding/json"
)


func GetWellLogCurveScale(c *fiber.Ctx) error {
	rows, err := db.Query("select * from well_log_curve_scale")
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}
	defer rows.Close()
	var result []dto.Well_log_curve_scale

	for rows.Next() {
		var well_log_curve_scale dto.Well_log_curve_scale
		if err := rows.Scan(&well_log_curve_scale.Uwi, &well_log_curve_scale.Curve_id, &well_log_curve_scale.Digital_curve_id, &well_log_curve_scale.Curve_scale_id, &well_log_curve_scale.Active_ind, &well_log_curve_scale.Backup_curve_scale, &well_log_curve_scale.Base_depth, &well_log_curve_scale.Base_depth_ouom, &well_log_curve_scale.Effective_date, &well_log_curve_scale.Expiry_date, &well_log_curve_scale.Left_scale_value, &well_log_curve_scale.Matrix_lithology_setting, &well_log_curve_scale.Ppdm_guid, &well_log_curve_scale.Remark, &well_log_curve_scale.Right_scale_value, &well_log_curve_scale.Scale_transform_type, &well_log_curve_scale.Source, &well_log_curve_scale.Top_depth, &well_log_curve_scale.Top_depth_ouom, &well_log_curve_scale.Track_num, &well_log_curve_scale.Track_width, &well_log_curve_scale.Track_width_ouom, &well_log_curve_scale.Track_width_uom, &well_log_curve_scale.Vertical_scale_ratio, &well_log_curve_scale.Row_changed_by, &well_log_curve_scale.Row_changed_date, &well_log_curve_scale.Row_created_by, &well_log_curve_scale.Row_created_date, &well_log_curve_scale.Row_effective_date, &well_log_curve_scale.Row_expiry_date, &well_log_curve_scale.Row_quality); err != nil {
			return err // Exit if we get an error
		}

		// Append well_log_curve_scale to result
		result = append(result, well_log_curve_scale)
	}
	// Return result in JSON format
	return c.JSON(result)
	 
}

func SetWellLogCurveScale(c *fiber.Ctx) error {
	var well_log_curve_scale dto.Well_log_curve_scale

	setDefaults(&well_log_curve_scale)

	if err := json.Unmarshal(c.Body(), &well_log_curve_scale); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	stmt, err := db.Prepare("insert into well_log_curve_scale values(:1, :2, :3, :4, :5, :6, :7, :8, :9, :10, :11, :12, :13, :14, :15, :16, :17, :18, :19, :20, :21, :22, :23, :24, :25, :26, :27, :28, :29, :30, :31)")
	if err != nil {
		return err
	}
	well_log_curve_scale.Row_created_date = formatDate(well_log_curve_scale.Row_created_date)
	well_log_curve_scale.Row_changed_date = formatDate(well_log_curve_scale.Row_changed_date)
	well_log_curve_scale.Effective_date = formatDateString(well_log_curve_scale.Effective_date)
	well_log_curve_scale.Expiry_date = formatDateString(well_log_curve_scale.Expiry_date)
	well_log_curve_scale.Row_effective_date = formatDateString(well_log_curve_scale.Row_effective_date)
	well_log_curve_scale.Row_expiry_date = formatDateString(well_log_curve_scale.Row_expiry_date)






	rows, err := stmt.Exec(well_log_curve_scale.Uwi, well_log_curve_scale.Curve_id, well_log_curve_scale.Digital_curve_id, well_log_curve_scale.Curve_scale_id, well_log_curve_scale.Active_ind, well_log_curve_scale.Backup_curve_scale, well_log_curve_scale.Base_depth, well_log_curve_scale.Base_depth_ouom, well_log_curve_scale.Effective_date, well_log_curve_scale.Expiry_date, well_log_curve_scale.Left_scale_value, well_log_curve_scale.Matrix_lithology_setting, well_log_curve_scale.Ppdm_guid, well_log_curve_scale.Remark, well_log_curve_scale.Right_scale_value, well_log_curve_scale.Scale_transform_type, well_log_curve_scale.Source, well_log_curve_scale.Top_depth, well_log_curve_scale.Top_depth_ouom, well_log_curve_scale.Track_num, well_log_curve_scale.Track_width, well_log_curve_scale.Track_width_ouom, well_log_curve_scale.Track_width_uom, well_log_curve_scale.Vertical_scale_ratio, well_log_curve_scale.Row_changed_by, well_log_curve_scale.Row_changed_date, well_log_curve_scale.Row_created_by, well_log_curve_scale.Row_created_date, well_log_curve_scale.Row_effective_date, well_log_curve_scale.Row_expiry_date, well_log_curve_scale.Row_quality)
	
	if err != nil {
		return err
	}

	return c.JSON(rows)
}

func UpdateWellLogCurveScale(c *fiber.Ctx) error {
	var well_log_curve_scale dto.Well_log_curve_scale

	setDefaults(&well_log_curve_scale)

	if err := json.Unmarshal(c.Body(), &well_log_curve_scale); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	id := c.Params("id")
	well_log_curve_scale.Uwi = id

    if well_log_curve_scale.Row_created_date != nil {
        return c.Status(400).SendString("Cannot update row_created_date")
    }

	stmt, err := db.Prepare("update well_log_curve_scale set curve_id = :1, digital_curve_id = :2, curve_scale_id = :3, active_ind = :4, backup_curve_scale = :5, base_depth = :6, base_depth_ouom = :7, effective_date = :8, expiry_date = :9, left_scale_value = :10, matrix_lithology_setting = :11, ppdm_guid = :12, remark = :13, right_scale_value = :14, scale_transform_type = :15, source = :16, top_depth = :17, top_depth_ouom = :18, track_num = :19, track_width = :20, track_width_ouom = :21, track_width_uom = :22, vertical_scale_ratio = :23, row_changed_by = :24, row_changed_date = :25, row_created_by = :26, row_effective_date = :27, row_expiry_date = :28, row_quality = :29 where uwi = :31")
	if err != nil {
		return err
	}

	well_log_curve_scale.Row_changed_date = formatDate(well_log_curve_scale.Row_changed_date)
	well_log_curve_scale.Effective_date = formatDateString(well_log_curve_scale.Effective_date)
	well_log_curve_scale.Expiry_date = formatDateString(well_log_curve_scale.Expiry_date)
	well_log_curve_scale.Row_effective_date = formatDateString(well_log_curve_scale.Row_effective_date)
	well_log_curve_scale.Row_expiry_date = formatDateString(well_log_curve_scale.Row_expiry_date)






	rows, err := stmt.Exec(well_log_curve_scale.Curve_id, well_log_curve_scale.Digital_curve_id, well_log_curve_scale.Curve_scale_id, well_log_curve_scale.Active_ind, well_log_curve_scale.Backup_curve_scale, well_log_curve_scale.Base_depth, well_log_curve_scale.Base_depth_ouom, well_log_curve_scale.Effective_date, well_log_curve_scale.Expiry_date, well_log_curve_scale.Left_scale_value, well_log_curve_scale.Matrix_lithology_setting, well_log_curve_scale.Ppdm_guid, well_log_curve_scale.Remark, well_log_curve_scale.Right_scale_value, well_log_curve_scale.Scale_transform_type, well_log_curve_scale.Source, well_log_curve_scale.Top_depth, well_log_curve_scale.Top_depth_ouom, well_log_curve_scale.Track_num, well_log_curve_scale.Track_width, well_log_curve_scale.Track_width_ouom, well_log_curve_scale.Track_width_uom, well_log_curve_scale.Vertical_scale_ratio, well_log_curve_scale.Row_changed_by, well_log_curve_scale.Row_changed_date, well_log_curve_scale.Row_created_by, well_log_curve_scale.Row_effective_date, well_log_curve_scale.Row_expiry_date, well_log_curve_scale.Row_quality, well_log_curve_scale.Uwi)
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

func PatchWellLogCurveScale(c *fiber.Ctx) error {
	id := c.Params("id")

	updatedFields := make(map[string]interface{})

	if err := c.BodyParser(&updatedFields); err != nil {
		return err
	}
	if _, ok := updatedFields["row_created_date"]; ok {
        return c.Status(400).SendString("Cannot update row_created_date field")
    }

	query := "update well_log_curve_scale set "
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

func DeleteWellLogCurveScale(c *fiber.Ctx) error {
	id := c.Params("id")
	var well_log_curve_scale dto.Well_log_curve_scale
	well_log_curve_scale.Uwi = id

	stmt, err := db.Prepare("delete from well_log_curve_scale where uwi = :1")
	if err != nil {
		return err
	}

	rows, err := stmt.Exec(well_log_curve_scale.Uwi)
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


