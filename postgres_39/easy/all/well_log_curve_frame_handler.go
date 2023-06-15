package apiv1

import (
	dto "github.com/DarrenMannuela/pt_gtn_bibliography/dto"
	"github.com/gofiber/fiber/v2"
	"strconv"
	"encoding/json"
)


func GetWellLogCurveFrame(c *fiber.Ctx) error {
	rows, err := db.Query("select * from well_log_curve_frame")
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}
	defer rows.Close()
	var result []dto.Well_log_curve_frame

	for rows.Next() {
		var well_log_curve_frame dto.Well_log_curve_frame
		if err := rows.Scan(&well_log_curve_frame.Uwi, &well_log_curve_frame.Well_log_id, &well_log_curve_frame.Well_log_source, &well_log_curve_frame.Frame_id, &well_log_curve_frame.Active_ind, &well_log_curve_frame.Base_depth, &well_log_curve_frame.Depth_ouom, &well_log_curve_frame.Effective_date, &well_log_curve_frame.Expiry_date, &well_log_curve_frame.Frame_name, &well_log_curve_frame.Frame_spacing, &well_log_curve_frame.Frame_spacing_ouom, &well_log_curve_frame.Frame_spacing_uom, &well_log_curve_frame.Gaps_ind, &well_log_curve_frame.Index_ouom, &well_log_curve_frame.Index_uom, &well_log_curve_frame.Irregular_desc, &well_log_curve_frame.Irregular_ind, &well_log_curve_frame.Log_direction, &well_log_curve_frame.Max_index, &well_log_curve_frame.Min_index, &well_log_curve_frame.Ppdm_guid, &well_log_curve_frame.Primary_index_type, &well_log_curve_frame.Remark, &well_log_curve_frame.Source, &well_log_curve_frame.Top_depth, &well_log_curve_frame.Row_changed_by, &well_log_curve_frame.Row_changed_date, &well_log_curve_frame.Row_created_by, &well_log_curve_frame.Row_created_date, &well_log_curve_frame.Row_effective_date, &well_log_curve_frame.Row_expiry_date, &well_log_curve_frame.Row_quality); err != nil {
			return err // Exit if we get an error
		}

		// Append well_log_curve_frame to result
		result = append(result, well_log_curve_frame)
	}
	// Return result in JSON format
	return c.JSON(result)
	 
}

func SetWellLogCurveFrame(c *fiber.Ctx) error {
	var well_log_curve_frame dto.Well_log_curve_frame

	setDefaults(&well_log_curve_frame)

	if err := json.Unmarshal(c.Body(), &well_log_curve_frame); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	stmt, err := db.Prepare("insert into well_log_curve_frame values(:1, :2, :3, :4, :5, :6, :7, :8, :9, :10, :11, :12, :13, :14, :15, :16, :17, :18, :19, :20, :21, :22, :23, :24, :25, :26, :27, :28, :29, :30, :31, :32, :33)")
	if err != nil {
		return err
	}
	well_log_curve_frame.Row_created_date = formatDate(well_log_curve_frame.Row_created_date)
	well_log_curve_frame.Row_changed_date = formatDate(well_log_curve_frame.Row_changed_date)
	well_log_curve_frame.Effective_date = formatDateString(well_log_curve_frame.Effective_date)
	well_log_curve_frame.Expiry_date = formatDateString(well_log_curve_frame.Expiry_date)
	well_log_curve_frame.Row_effective_date = formatDateString(well_log_curve_frame.Row_effective_date)
	well_log_curve_frame.Row_expiry_date = formatDateString(well_log_curve_frame.Row_expiry_date)






	rows, err := stmt.Exec(well_log_curve_frame.Uwi, well_log_curve_frame.Well_log_id, well_log_curve_frame.Well_log_source, well_log_curve_frame.Frame_id, well_log_curve_frame.Active_ind, well_log_curve_frame.Base_depth, well_log_curve_frame.Depth_ouom, well_log_curve_frame.Effective_date, well_log_curve_frame.Expiry_date, well_log_curve_frame.Frame_name, well_log_curve_frame.Frame_spacing, well_log_curve_frame.Frame_spacing_ouom, well_log_curve_frame.Frame_spacing_uom, well_log_curve_frame.Gaps_ind, well_log_curve_frame.Index_ouom, well_log_curve_frame.Index_uom, well_log_curve_frame.Irregular_desc, well_log_curve_frame.Irregular_ind, well_log_curve_frame.Log_direction, well_log_curve_frame.Max_index, well_log_curve_frame.Min_index, well_log_curve_frame.Ppdm_guid, well_log_curve_frame.Primary_index_type, well_log_curve_frame.Remark, well_log_curve_frame.Source, well_log_curve_frame.Top_depth, well_log_curve_frame.Row_changed_by, well_log_curve_frame.Row_changed_date, well_log_curve_frame.Row_created_by, well_log_curve_frame.Row_created_date, well_log_curve_frame.Row_effective_date, well_log_curve_frame.Row_expiry_date, well_log_curve_frame.Row_quality)
	
	if err != nil {
		return err
	}

	return c.JSON(rows)
}

func UpdateWellLogCurveFrame(c *fiber.Ctx) error {
	var well_log_curve_frame dto.Well_log_curve_frame

	setDefaults(&well_log_curve_frame)

	if err := json.Unmarshal(c.Body(), &well_log_curve_frame); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	id := c.Params("id")
	well_log_curve_frame.Uwi = id

    if well_log_curve_frame.Row_created_date != nil {
        return c.Status(400).SendString("Cannot update row_created_date")
    }

	stmt, err := db.Prepare("update well_log_curve_frame set well_log_id = :1, well_log_source = :2, frame_id = :3, active_ind = :4, base_depth = :5, depth_ouom = :6, effective_date = :7, expiry_date = :8, frame_name = :9, frame_spacing = :10, frame_spacing_ouom = :11, frame_spacing_uom = :12, gaps_ind = :13, index_ouom = :14, index_uom = :15, irregular_desc = :16, irregular_ind = :17, log_direction = :18, max_index = :19, min_index = :20, ppdm_guid = :21, primary_index_type = :22, remark = :23, source = :24, top_depth = :25, row_changed_by = :26, row_changed_date = :27, row_created_by = :28, row_effective_date = :29, row_expiry_date = :30, row_quality = :31 where uwi = :33")
	if err != nil {
		return err
	}

	well_log_curve_frame.Row_changed_date = formatDate(well_log_curve_frame.Row_changed_date)
	well_log_curve_frame.Effective_date = formatDateString(well_log_curve_frame.Effective_date)
	well_log_curve_frame.Expiry_date = formatDateString(well_log_curve_frame.Expiry_date)
	well_log_curve_frame.Row_effective_date = formatDateString(well_log_curve_frame.Row_effective_date)
	well_log_curve_frame.Row_expiry_date = formatDateString(well_log_curve_frame.Row_expiry_date)






	rows, err := stmt.Exec(well_log_curve_frame.Well_log_id, well_log_curve_frame.Well_log_source, well_log_curve_frame.Frame_id, well_log_curve_frame.Active_ind, well_log_curve_frame.Base_depth, well_log_curve_frame.Depth_ouom, well_log_curve_frame.Effective_date, well_log_curve_frame.Expiry_date, well_log_curve_frame.Frame_name, well_log_curve_frame.Frame_spacing, well_log_curve_frame.Frame_spacing_ouom, well_log_curve_frame.Frame_spacing_uom, well_log_curve_frame.Gaps_ind, well_log_curve_frame.Index_ouom, well_log_curve_frame.Index_uom, well_log_curve_frame.Irregular_desc, well_log_curve_frame.Irregular_ind, well_log_curve_frame.Log_direction, well_log_curve_frame.Max_index, well_log_curve_frame.Min_index, well_log_curve_frame.Ppdm_guid, well_log_curve_frame.Primary_index_type, well_log_curve_frame.Remark, well_log_curve_frame.Source, well_log_curve_frame.Top_depth, well_log_curve_frame.Row_changed_by, well_log_curve_frame.Row_changed_date, well_log_curve_frame.Row_created_by, well_log_curve_frame.Row_effective_date, well_log_curve_frame.Row_expiry_date, well_log_curve_frame.Row_quality, well_log_curve_frame.Uwi)
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

func PatchWellLogCurveFrame(c *fiber.Ctx) error {
	id := c.Params("id")

	updatedFields := make(map[string]interface{})

	if err := c.BodyParser(&updatedFields); err != nil {
		return err
	}
	if _, ok := updatedFields["row_created_date"]; ok {
        return c.Status(400).SendString("Cannot update row_created_date field")
    }

	query := "update well_log_curve_frame set "
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

func DeleteWellLogCurveFrame(c *fiber.Ctx) error {
	id := c.Params("id")
	var well_log_curve_frame dto.Well_log_curve_frame
	well_log_curve_frame.Uwi = id

	stmt, err := db.Prepare("delete from well_log_curve_frame where uwi = :1")
	if err != nil {
		return err
	}

	rows, err := stmt.Exec(well_log_curve_frame.Uwi)
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


