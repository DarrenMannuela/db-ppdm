package apiv1

import (
	dto "github.com/DarrenMannuela/pt_gtn_bibliography/dto"
	"github.com/gofiber/fiber/v2"
	"strconv"
	"encoding/json"
)


func GetPaleoSumInterval(c *fiber.Ctx) error {
	rows, err := db.Query("select * from paleo_sum_interval")
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}
	defer rows.Close()
	var result []dto.Paleo_sum_interval

	for rows.Next() {
		var paleo_sum_interval dto.Paleo_sum_interval
		if err := rows.Scan(&paleo_sum_interval.Paleo_summary_id, &paleo_sum_interval.Interval_id, &paleo_sum_interval.Active_ind, &paleo_sum_interval.Base_depth, &paleo_sum_interval.Base_depth_ouom, &paleo_sum_interval.Effective_date, &paleo_sum_interval.Expiry_date, &paleo_sum_interval.Interval_desc, &paleo_sum_interval.Ppdm_guid, &paleo_sum_interval.Remark, &paleo_sum_interval.Source, &paleo_sum_interval.Top_depth, &paleo_sum_interval.Top_depth_ouom, &paleo_sum_interval.Row_changed_by, &paleo_sum_interval.Row_changed_date, &paleo_sum_interval.Row_created_by, &paleo_sum_interval.Row_created_date, &paleo_sum_interval.Row_effective_date, &paleo_sum_interval.Row_expiry_date, &paleo_sum_interval.Row_quality); err != nil {
			return err // Exit if we get an error
		}

		// Append paleo_sum_interval to result
		result = append(result, paleo_sum_interval)
	}
	// Return result in JSON format
	return c.JSON(result)
	 
}

func SetPaleoSumInterval(c *fiber.Ctx) error {
	var paleo_sum_interval dto.Paleo_sum_interval

	setDefaults(&paleo_sum_interval)

	if err := json.Unmarshal(c.Body(), &paleo_sum_interval); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	stmt, err := db.Prepare("insert into paleo_sum_interval values(:1, :2, :3, :4, :5, :6, :7, :8, :9, :10, :11, :12, :13, :14, :15, :16, :17, :18, :19, :20)")
	if err != nil {
		return err
	}
	paleo_sum_interval.Row_created_date = formatDate(paleo_sum_interval.Row_created_date)
	paleo_sum_interval.Row_changed_date = formatDate(paleo_sum_interval.Row_changed_date)
	paleo_sum_interval.Effective_date = formatDateString(paleo_sum_interval.Effective_date)
	paleo_sum_interval.Expiry_date = formatDateString(paleo_sum_interval.Expiry_date)
	paleo_sum_interval.Row_effective_date = formatDateString(paleo_sum_interval.Row_effective_date)
	paleo_sum_interval.Row_expiry_date = formatDateString(paleo_sum_interval.Row_expiry_date)






	rows, err := stmt.Exec(paleo_sum_interval.Paleo_summary_id, paleo_sum_interval.Interval_id, paleo_sum_interval.Active_ind, paleo_sum_interval.Base_depth, paleo_sum_interval.Base_depth_ouom, paleo_sum_interval.Effective_date, paleo_sum_interval.Expiry_date, paleo_sum_interval.Interval_desc, paleo_sum_interval.Ppdm_guid, paleo_sum_interval.Remark, paleo_sum_interval.Source, paleo_sum_interval.Top_depth, paleo_sum_interval.Top_depth_ouom, paleo_sum_interval.Row_changed_by, paleo_sum_interval.Row_changed_date, paleo_sum_interval.Row_created_by, paleo_sum_interval.Row_created_date, paleo_sum_interval.Row_effective_date, paleo_sum_interval.Row_expiry_date, paleo_sum_interval.Row_quality)
	
	if err != nil {
		return err
	}

	return c.JSON(rows)
}

func UpdatePaleoSumInterval(c *fiber.Ctx) error {
	var paleo_sum_interval dto.Paleo_sum_interval

	setDefaults(&paleo_sum_interval)

	if err := json.Unmarshal(c.Body(), &paleo_sum_interval); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	id := c.Params("id")
	paleo_sum_interval.Paleo_summary_id = id

    if paleo_sum_interval.Row_created_date != nil {
        return c.Status(400).SendString("Cannot update row_created_date")
    }

	stmt, err := db.Prepare("update paleo_sum_interval set interval_id = :1, active_ind = :2, base_depth = :3, base_depth_ouom = :4, effective_date = :5, expiry_date = :6, interval_desc = :7, ppdm_guid = :8, remark = :9, source = :10, top_depth = :11, top_depth_ouom = :12, row_changed_by = :13, row_changed_date = :14, row_created_by = :15, row_effective_date = :16, row_expiry_date = :17, row_quality = :18 where paleo_summary_id = :20")
	if err != nil {
		return err
	}

	paleo_sum_interval.Row_changed_date = formatDate(paleo_sum_interval.Row_changed_date)
	paleo_sum_interval.Effective_date = formatDateString(paleo_sum_interval.Effective_date)
	paleo_sum_interval.Expiry_date = formatDateString(paleo_sum_interval.Expiry_date)
	paleo_sum_interval.Row_effective_date = formatDateString(paleo_sum_interval.Row_effective_date)
	paleo_sum_interval.Row_expiry_date = formatDateString(paleo_sum_interval.Row_expiry_date)






	rows, err := stmt.Exec(paleo_sum_interval.Interval_id, paleo_sum_interval.Active_ind, paleo_sum_interval.Base_depth, paleo_sum_interval.Base_depth_ouom, paleo_sum_interval.Effective_date, paleo_sum_interval.Expiry_date, paleo_sum_interval.Interval_desc, paleo_sum_interval.Ppdm_guid, paleo_sum_interval.Remark, paleo_sum_interval.Source, paleo_sum_interval.Top_depth, paleo_sum_interval.Top_depth_ouom, paleo_sum_interval.Row_changed_by, paleo_sum_interval.Row_changed_date, paleo_sum_interval.Row_created_by, paleo_sum_interval.Row_effective_date, paleo_sum_interval.Row_expiry_date, paleo_sum_interval.Row_quality, paleo_sum_interval.Paleo_summary_id)
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

func PatchPaleoSumInterval(c *fiber.Ctx) error {
	id := c.Params("id")

	updatedFields := make(map[string]interface{})

	if err := c.BodyParser(&updatedFields); err != nil {
		return err
	}
	if _, ok := updatedFields["row_created_date"]; ok {
        return c.Status(400).SendString("Cannot update row_created_date field")
    }

	query := "update paleo_sum_interval set "
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
	query += " where paleo_summary_id = :id"

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

func DeletePaleoSumInterval(c *fiber.Ctx) error {
	id := c.Params("id")
	var paleo_sum_interval dto.Paleo_sum_interval
	paleo_sum_interval.Paleo_summary_id = id

	stmt, err := db.Prepare("delete from paleo_sum_interval where paleo_summary_id = :1")
	if err != nil {
		return err
	}

	rows, err := stmt.Exec(paleo_sum_interval.Paleo_summary_id)
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


