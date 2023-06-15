package apiv1

import (
	dto "github.com/DarrenMannuela/pt_gtn_bibliography/dto"
	"github.com/gofiber/fiber/v2"
	"strconv"
	"encoding/json"
)


func GetRmWellLog(c *fiber.Ctx) error {
	rows, err := db.Query("select * from rm_well_log")
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}
	defer rows.Close()
	var result []dto.Rm_well_log

	for rows.Next() {
		var rm_well_log dto.Rm_well_log
		if err := rows.Scan(&rm_well_log.Info_item_subtype, &rm_well_log.Information_item_id, &rm_well_log.Active_ind, &rm_well_log.Display_interval, &rm_well_log.Display_interval_uom, &rm_well_log.Display_scale, &rm_well_log.Display_scale_uom, &rm_well_log.Effective_date, &rm_well_log.Expiry_date, &rm_well_log.Interpreted_ind, &rm_well_log.Ppdm_guid, &rm_well_log.Remark, &rm_well_log.Reported_base_depth, &rm_well_log.Reported_base_depth_ouom, &rm_well_log.Reported_top_depth, &rm_well_log.Reported_top_depth_ouom, &rm_well_log.Source, &rm_well_log.Row_changed_by, &rm_well_log.Row_changed_date, &rm_well_log.Row_created_by, &rm_well_log.Row_created_date, &rm_well_log.Row_effective_date, &rm_well_log.Row_expiry_date, &rm_well_log.Row_quality); err != nil {
			return err // Exit if we get an error
		}

		// Append rm_well_log to result
		result = append(result, rm_well_log)
	}
	// Return result in JSON format
	return c.JSON(result)
	 
}

func SetRmWellLog(c *fiber.Ctx) error {
	var rm_well_log dto.Rm_well_log

	setDefaults(&rm_well_log)

	if err := json.Unmarshal(c.Body(), &rm_well_log); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	stmt, err := db.Prepare("insert into rm_well_log values(:1, :2, :3, :4, :5, :6, :7, :8, :9, :10, :11, :12, :13, :14, :15, :16, :17, :18, :19, :20, :21, :22, :23, :24)")
	if err != nil {
		return err
	}
	rm_well_log.Row_created_date = formatDate(rm_well_log.Row_created_date)
	rm_well_log.Row_changed_date = formatDate(rm_well_log.Row_changed_date)
	rm_well_log.Effective_date = formatDateString(rm_well_log.Effective_date)
	rm_well_log.Expiry_date = formatDateString(rm_well_log.Expiry_date)
	rm_well_log.Row_effective_date = formatDateString(rm_well_log.Row_effective_date)
	rm_well_log.Row_expiry_date = formatDateString(rm_well_log.Row_expiry_date)






	rows, err := stmt.Exec(rm_well_log.Info_item_subtype, rm_well_log.Information_item_id, rm_well_log.Active_ind, rm_well_log.Display_interval, rm_well_log.Display_interval_uom, rm_well_log.Display_scale, rm_well_log.Display_scale_uom, rm_well_log.Effective_date, rm_well_log.Expiry_date, rm_well_log.Interpreted_ind, rm_well_log.Ppdm_guid, rm_well_log.Remark, rm_well_log.Reported_base_depth, rm_well_log.Reported_base_depth_ouom, rm_well_log.Reported_top_depth, rm_well_log.Reported_top_depth_ouom, rm_well_log.Source, rm_well_log.Row_changed_by, rm_well_log.Row_changed_date, rm_well_log.Row_created_by, rm_well_log.Row_created_date, rm_well_log.Row_effective_date, rm_well_log.Row_expiry_date, rm_well_log.Row_quality)
	
	if err != nil {
		return err
	}

	return c.JSON(rows)
}

func UpdateRmWellLog(c *fiber.Ctx) error {
	var rm_well_log dto.Rm_well_log

	setDefaults(&rm_well_log)

	if err := json.Unmarshal(c.Body(), &rm_well_log); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	id := c.Params("id")
	rm_well_log.Info_item_subtype = id

    if rm_well_log.Row_created_date != nil {
        return c.Status(400).SendString("Cannot update row_created_date")
    }

	stmt, err := db.Prepare("update rm_well_log set information_item_id = :1, active_ind = :2, display_interval = :3, display_interval_uom = :4, display_scale = :5, display_scale_uom = :6, effective_date = :7, expiry_date = :8, interpreted_ind = :9, ppdm_guid = :10, remark = :11, reported_base_depth = :12, reported_base_depth_ouom = :13, reported_top_depth = :14, reported_top_depth_ouom = :15, source = :16, row_changed_by = :17, row_changed_date = :18, row_created_by = :19, row_effective_date = :20, row_expiry_date = :21, row_quality = :22 where info_item_subtype = :24")
	if err != nil {
		return err
	}

	rm_well_log.Row_changed_date = formatDate(rm_well_log.Row_changed_date)
	rm_well_log.Effective_date = formatDateString(rm_well_log.Effective_date)
	rm_well_log.Expiry_date = formatDateString(rm_well_log.Expiry_date)
	rm_well_log.Row_effective_date = formatDateString(rm_well_log.Row_effective_date)
	rm_well_log.Row_expiry_date = formatDateString(rm_well_log.Row_expiry_date)






	rows, err := stmt.Exec(rm_well_log.Information_item_id, rm_well_log.Active_ind, rm_well_log.Display_interval, rm_well_log.Display_interval_uom, rm_well_log.Display_scale, rm_well_log.Display_scale_uom, rm_well_log.Effective_date, rm_well_log.Expiry_date, rm_well_log.Interpreted_ind, rm_well_log.Ppdm_guid, rm_well_log.Remark, rm_well_log.Reported_base_depth, rm_well_log.Reported_base_depth_ouom, rm_well_log.Reported_top_depth, rm_well_log.Reported_top_depth_ouom, rm_well_log.Source, rm_well_log.Row_changed_by, rm_well_log.Row_changed_date, rm_well_log.Row_created_by, rm_well_log.Row_effective_date, rm_well_log.Row_expiry_date, rm_well_log.Row_quality, rm_well_log.Info_item_subtype)
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

func PatchRmWellLog(c *fiber.Ctx) error {
	id := c.Params("id")

	updatedFields := make(map[string]interface{})

	if err := c.BodyParser(&updatedFields); err != nil {
		return err
	}
	if _, ok := updatedFields["row_created_date"]; ok {
        return c.Status(400).SendString("Cannot update row_created_date field")
    }

	query := "update rm_well_log set "
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
	query += " where info_item_subtype = :id"

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

func DeleteRmWellLog(c *fiber.Ctx) error {
	id := c.Params("id")
	var rm_well_log dto.Rm_well_log
	rm_well_log.Info_item_subtype = id

	stmt, err := db.Prepare("delete from rm_well_log where info_item_subtype = :1")
	if err != nil {
		return err
	}

	rows, err := stmt.Exec(rm_well_log.Info_item_subtype)
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


