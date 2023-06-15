package apiv1

import (
	dto "github.com/DarrenMannuela/pt_gtn_bibliography/dto"
	"github.com/gofiber/fiber/v2"
	"strconv"
	"encoding/json"
)


func GetRmImageSect(c *fiber.Ctx) error {
	rows, err := db.Query("select * from rm_image_sect")
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}
	defer rows.Close()
	var result []dto.Rm_image_sect

	for rows.Next() {
		var rm_image_sect dto.Rm_image_sect
		if err := rows.Scan(&rm_image_sect.Physical_item_id, &rm_image_sect.Image_section_id, &rm_image_sect.Active_ind, &rm_image_sect.Base_depth, &rm_image_sect.Base_depth_ouom, &rm_image_sect.Calibrated_by_ba_id, &rm_image_sect.Calibrate_application, &rm_image_sect.Calibrate_method, &rm_image_sect.Effective_date, &rm_image_sect.Expiry_date, &rm_image_sect.Image_section_type, &rm_image_sect.Index_type, &rm_image_sect.Log_matrix_type, &rm_image_sect.Ppdm_guid, &rm_image_sect.Remark, &rm_image_sect.Scale_depth_interval, &rm_image_sect.Scale_depth_interval_ouom, &rm_image_sect.Scale_depth_interval_uom, &rm_image_sect.Scale_length, &rm_image_sect.Scale_length_ouom, &rm_image_sect.Scale_length_uom, &rm_image_sect.Scale_ratio, &rm_image_sect.Section_desc, &rm_image_sect.Source, &rm_image_sect.Top_depth, &rm_image_sect.Top_depth_ouom, &rm_image_sect.Well_log_class_id, &rm_image_sect.Row_changed_by, &rm_image_sect.Row_changed_date, &rm_image_sect.Row_created_by, &rm_image_sect.Row_created_date, &rm_image_sect.Row_effective_date, &rm_image_sect.Row_expiry_date, &rm_image_sect.Row_quality); err != nil {
			return err // Exit if we get an error
		}

		// Append rm_image_sect to result
		result = append(result, rm_image_sect)
	}
	// Return result in JSON format
	return c.JSON(result)
	 
}

func SetRmImageSect(c *fiber.Ctx) error {
	var rm_image_sect dto.Rm_image_sect

	setDefaults(&rm_image_sect)

	if err := json.Unmarshal(c.Body(), &rm_image_sect); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	stmt, err := db.Prepare("insert into rm_image_sect values(:1, :2, :3, :4, :5, :6, :7, :8, :9, :10, :11, :12, :13, :14, :15, :16, :17, :18, :19, :20, :21, :22, :23, :24, :25, :26, :27, :28, :29, :30, :31, :32, :33, :34)")
	if err != nil {
		return err
	}
	rm_image_sect.Row_created_date = formatDate(rm_image_sect.Row_created_date)
	rm_image_sect.Row_changed_date = formatDate(rm_image_sect.Row_changed_date)
	rm_image_sect.Effective_date = formatDateString(rm_image_sect.Effective_date)
	rm_image_sect.Expiry_date = formatDateString(rm_image_sect.Expiry_date)
	rm_image_sect.Row_effective_date = formatDateString(rm_image_sect.Row_effective_date)
	rm_image_sect.Row_expiry_date = formatDateString(rm_image_sect.Row_expiry_date)






	rows, err := stmt.Exec(rm_image_sect.Physical_item_id, rm_image_sect.Image_section_id, rm_image_sect.Active_ind, rm_image_sect.Base_depth, rm_image_sect.Base_depth_ouom, rm_image_sect.Calibrated_by_ba_id, rm_image_sect.Calibrate_application, rm_image_sect.Calibrate_method, rm_image_sect.Effective_date, rm_image_sect.Expiry_date, rm_image_sect.Image_section_type, rm_image_sect.Index_type, rm_image_sect.Log_matrix_type, rm_image_sect.Ppdm_guid, rm_image_sect.Remark, rm_image_sect.Scale_depth_interval, rm_image_sect.Scale_depth_interval_ouom, rm_image_sect.Scale_depth_interval_uom, rm_image_sect.Scale_length, rm_image_sect.Scale_length_ouom, rm_image_sect.Scale_length_uom, rm_image_sect.Scale_ratio, rm_image_sect.Section_desc, rm_image_sect.Source, rm_image_sect.Top_depth, rm_image_sect.Top_depth_ouom, rm_image_sect.Well_log_class_id, rm_image_sect.Row_changed_by, rm_image_sect.Row_changed_date, rm_image_sect.Row_created_by, rm_image_sect.Row_created_date, rm_image_sect.Row_effective_date, rm_image_sect.Row_expiry_date, rm_image_sect.Row_quality)
	
	if err != nil {
		return err
	}

	return c.JSON(rows)
}

func UpdateRmImageSect(c *fiber.Ctx) error {
	var rm_image_sect dto.Rm_image_sect

	setDefaults(&rm_image_sect)

	if err := json.Unmarshal(c.Body(), &rm_image_sect); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	id := c.Params("id")
	rm_image_sect.Physical_item_id = id

    if rm_image_sect.Row_created_date != nil {
        return c.Status(400).SendString("Cannot update row_created_date")
    }

	stmt, err := db.Prepare("update rm_image_sect set image_section_id = :1, active_ind = :2, base_depth = :3, base_depth_ouom = :4, calibrated_by_ba_id = :5, calibrate_application = :6, calibrate_method = :7, effective_date = :8, expiry_date = :9, image_section_type = :10, index_type = :11, log_matrix_type = :12, ppdm_guid = :13, remark = :14, scale_depth_interval = :15, scale_depth_interval_ouom = :16, scale_depth_interval_uom = :17, scale_length = :18, scale_length_ouom = :19, scale_length_uom = :20, scale_ratio = :21, section_desc = :22, source = :23, top_depth = :24, top_depth_ouom = :25, well_log_class_id = :26, row_changed_by = :27, row_changed_date = :28, row_created_by = :29, row_effective_date = :30, row_expiry_date = :31, row_quality = :32 where physical_item_id = :34")
	if err != nil {
		return err
	}

	rm_image_sect.Row_changed_date = formatDate(rm_image_sect.Row_changed_date)
	rm_image_sect.Effective_date = formatDateString(rm_image_sect.Effective_date)
	rm_image_sect.Expiry_date = formatDateString(rm_image_sect.Expiry_date)
	rm_image_sect.Row_effective_date = formatDateString(rm_image_sect.Row_effective_date)
	rm_image_sect.Row_expiry_date = formatDateString(rm_image_sect.Row_expiry_date)






	rows, err := stmt.Exec(rm_image_sect.Image_section_id, rm_image_sect.Active_ind, rm_image_sect.Base_depth, rm_image_sect.Base_depth_ouom, rm_image_sect.Calibrated_by_ba_id, rm_image_sect.Calibrate_application, rm_image_sect.Calibrate_method, rm_image_sect.Effective_date, rm_image_sect.Expiry_date, rm_image_sect.Image_section_type, rm_image_sect.Index_type, rm_image_sect.Log_matrix_type, rm_image_sect.Ppdm_guid, rm_image_sect.Remark, rm_image_sect.Scale_depth_interval, rm_image_sect.Scale_depth_interval_ouom, rm_image_sect.Scale_depth_interval_uom, rm_image_sect.Scale_length, rm_image_sect.Scale_length_ouom, rm_image_sect.Scale_length_uom, rm_image_sect.Scale_ratio, rm_image_sect.Section_desc, rm_image_sect.Source, rm_image_sect.Top_depth, rm_image_sect.Top_depth_ouom, rm_image_sect.Well_log_class_id, rm_image_sect.Row_changed_by, rm_image_sect.Row_changed_date, rm_image_sect.Row_created_by, rm_image_sect.Row_effective_date, rm_image_sect.Row_expiry_date, rm_image_sect.Row_quality, rm_image_sect.Physical_item_id)
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

func PatchRmImageSect(c *fiber.Ctx) error {
	id := c.Params("id")

	updatedFields := make(map[string]interface{})

	if err := c.BodyParser(&updatedFields); err != nil {
		return err
	}
	if _, ok := updatedFields["row_created_date"]; ok {
        return c.Status(400).SendString("Cannot update row_created_date field")
    }

	query := "update rm_image_sect set "
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

func DeleteRmImageSect(c *fiber.Ctx) error {
	id := c.Params("id")
	var rm_image_sect dto.Rm_image_sect
	rm_image_sect.Physical_item_id = id

	stmt, err := db.Prepare("delete from rm_image_sect where physical_item_id = :1")
	if err != nil {
		return err
	}

	rows, err := stmt.Exec(rm_image_sect.Physical_item_id)
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


