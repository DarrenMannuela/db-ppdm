package apiv1

import (
	dto "github.com/DarrenMannuela/pt_gtn_bibliography/dto"
	"github.com/gofiber/fiber/v2"
	"strconv"
	"encoding/json"
)


func GetLithLog(c *fiber.Ctx) error {
	rows, err := db.Query("select * from lith_log")
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}
	defer rows.Close()
	var result []dto.Lith_log

	for rows.Next() {
		var lith_log dto.Lith_log
		if err := rows.Scan(&lith_log.Lithology_log_id, &lith_log.Lith_log_source, &lith_log.Active_ind, &lith_log.Base_depth, &lith_log.Base_depth_ouom, &lith_log.Effective_date, &lith_log.Expiry_date, &lith_log.Geologist, &lith_log.Lith_log_class, &lith_log.Lith_log_type, &lith_log.Logged_date, &lith_log.Meas_section_id, &lith_log.Meas_section_source, &lith_log.Ppdm_guid, &lith_log.Remark, &lith_log.Top_depth, &lith_log.Top_depth_ouom, &lith_log.Uwi, &lith_log.Row_changed_by, &lith_log.Row_changed_date, &lith_log.Row_created_by, &lith_log.Row_created_date, &lith_log.Row_effective_date, &lith_log.Row_expiry_date, &lith_log.Row_quality); err != nil {
			return err // Exit if we get an error
		}

		// Append lith_log to result
		result = append(result, lith_log)
	}
	// Return result in JSON format
	return c.JSON(result)
	 
}

func SetLithLog(c *fiber.Ctx) error {
	var lith_log dto.Lith_log

	setDefaults(&lith_log)

	if err := json.Unmarshal(c.Body(), &lith_log); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	stmt, err := db.Prepare("insert into lith_log values(:1, :2, :3, :4, :5, :6, :7, :8, :9, :10, :11, :12, :13, :14, :15, :16, :17, :18, :19, :20, :21, :22, :23, :24, :25)")
	if err != nil {
		return err
	}
	lith_log.Row_created_date = formatDate(lith_log.Row_created_date)
	lith_log.Row_changed_date = formatDate(lith_log.Row_changed_date)
	lith_log.Effective_date = formatDateString(lith_log.Effective_date)
	lith_log.Expiry_date = formatDateString(lith_log.Expiry_date)
	lith_log.Logged_date = formatDateString(lith_log.Logged_date)
	lith_log.Row_effective_date = formatDateString(lith_log.Row_effective_date)
	lith_log.Row_expiry_date = formatDateString(lith_log.Row_expiry_date)






	rows, err := stmt.Exec(lith_log.Lithology_log_id, lith_log.Lith_log_source, lith_log.Active_ind, lith_log.Base_depth, lith_log.Base_depth_ouom, lith_log.Effective_date, lith_log.Expiry_date, lith_log.Geologist, lith_log.Lith_log_class, lith_log.Lith_log_type, lith_log.Logged_date, lith_log.Meas_section_id, lith_log.Meas_section_source, lith_log.Ppdm_guid, lith_log.Remark, lith_log.Top_depth, lith_log.Top_depth_ouom, lith_log.Uwi, lith_log.Row_changed_by, lith_log.Row_changed_date, lith_log.Row_created_by, lith_log.Row_created_date, lith_log.Row_effective_date, lith_log.Row_expiry_date, lith_log.Row_quality)
	
	if err != nil {
		return err
	}

	return c.JSON(rows)
}

func UpdateLithLog(c *fiber.Ctx) error {
	var lith_log dto.Lith_log

	setDefaults(&lith_log)

	if err := json.Unmarshal(c.Body(), &lith_log); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	id := c.Params("id")
	lith_log.Lithology_log_id = id

    if lith_log.Row_created_date != nil {
        return c.Status(400).SendString("Cannot update row_created_date")
    }

	stmt, err := db.Prepare("update lith_log set lith_log_source = :1, active_ind = :2, base_depth = :3, base_depth_ouom = :4, effective_date = :5, expiry_date = :6, geologist = :7, lith_log_class = :8, lith_log_type = :9, logged_date = :10, meas_section_id = :11, meas_section_source = :12, ppdm_guid = :13, remark = :14, top_depth = :15, top_depth_ouom = :16, uwi = :17, row_changed_by = :18, row_changed_date = :19, row_created_by = :20, row_effective_date = :21, row_expiry_date = :22, row_quality = :23 where lithology_log_id = :25")
	if err != nil {
		return err
	}

	lith_log.Row_changed_date = formatDate(lith_log.Row_changed_date)
	lith_log.Effective_date = formatDateString(lith_log.Effective_date)
	lith_log.Expiry_date = formatDateString(lith_log.Expiry_date)
	lith_log.Logged_date = formatDateString(lith_log.Logged_date)
	lith_log.Row_effective_date = formatDateString(lith_log.Row_effective_date)
	lith_log.Row_expiry_date = formatDateString(lith_log.Row_expiry_date)






	rows, err := stmt.Exec(lith_log.Lith_log_source, lith_log.Active_ind, lith_log.Base_depth, lith_log.Base_depth_ouom, lith_log.Effective_date, lith_log.Expiry_date, lith_log.Geologist, lith_log.Lith_log_class, lith_log.Lith_log_type, lith_log.Logged_date, lith_log.Meas_section_id, lith_log.Meas_section_source, lith_log.Ppdm_guid, lith_log.Remark, lith_log.Top_depth, lith_log.Top_depth_ouom, lith_log.Uwi, lith_log.Row_changed_by, lith_log.Row_changed_date, lith_log.Row_created_by, lith_log.Row_effective_date, lith_log.Row_expiry_date, lith_log.Row_quality, lith_log.Lithology_log_id)
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

func PatchLithLog(c *fiber.Ctx) error {
	id := c.Params("id")

	updatedFields := make(map[string]interface{})

	if err := c.BodyParser(&updatedFields); err != nil {
		return err
	}
	if _, ok := updatedFields["row_created_date"]; ok {
        return c.Status(400).SendString("Cannot update row_created_date field")
    }

	query := "update lith_log set "
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
		} else if key == "effective_date" || key == "expiry_date" || key == "logged_date" || key == "row_effective_date" || key == "row_expiry_date"      {
			if date, ok := value.(string); ok && len(date) > 0 {
				formattedDate := formatDateString(&date)
				value = formattedDate
			}
		}
		values = append(values, value)
	}
	query += " where lithology_log_id = :id"

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

func DeleteLithLog(c *fiber.Ctx) error {
	id := c.Params("id")
	var lith_log dto.Lith_log
	lith_log.Lithology_log_id = id

	stmt, err := db.Prepare("delete from lith_log where lithology_log_id = :1")
	if err != nil {
		return err
	}

	rows, err := stmt.Exec(lith_log.Lithology_log_id)
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


