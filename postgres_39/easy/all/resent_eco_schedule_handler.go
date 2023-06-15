package apiv1

import (
	dto "github.com/DarrenMannuela/pt_gtn_bibliography/dto"
	"github.com/gofiber/fiber/v2"
	"strconv"
	"encoding/json"
)


func GetResentEcoSchedule(c *fiber.Ctx) error {
	rows, err := db.Query("select * from resent_eco_schedule")
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}
	defer rows.Close()
	var result []dto.Resent_eco_schedule

	for rows.Next() {
		var resent_eco_schedule dto.Resent_eco_schedule
		if err := rows.Scan(&resent_eco_schedule.Resent_id, &resent_eco_schedule.Reserve_class_id, &resent_eco_schedule.Economics_run_id, &resent_eco_schedule.Schedule_id, &resent_eco_schedule.Active_ind, &resent_eco_schedule.Economic_schedule, &resent_eco_schedule.Effective_date, &resent_eco_schedule.Expiry_date, &resent_eco_schedule.Max_value, &resent_eco_schedule.Max_value_ouom, &resent_eco_schedule.Max_value_uom, &resent_eco_schedule.Min_value, &resent_eco_schedule.Min_value_ouom, &resent_eco_schedule.Min_value_uom, &resent_eco_schedule.Period_type, &resent_eco_schedule.Ppdm_guid, &resent_eco_schedule.Product_ind, &resent_eco_schedule.Product_type, &resent_eco_schedule.Remark, &resent_eco_schedule.Schedule_date, &resent_eco_schedule.Schedule_date_desc, &resent_eco_schedule.Schedule_desc, &resent_eco_schedule.Schedule_value, &resent_eco_schedule.Schedule_value_ouom, &resent_eco_schedule.Schedule_value_uom, &resent_eco_schedule.Source, &resent_eco_schedule.Row_changed_by, &resent_eco_schedule.Row_changed_date, &resent_eco_schedule.Row_created_by, &resent_eco_schedule.Row_created_date, &resent_eco_schedule.Row_effective_date, &resent_eco_schedule.Row_expiry_date, &resent_eco_schedule.Row_quality); err != nil {
			return err // Exit if we get an error
		}

		// Append resent_eco_schedule to result
		result = append(result, resent_eco_schedule)
	}
	// Return result in JSON format
	return c.JSON(result)
	 
}

func SetResentEcoSchedule(c *fiber.Ctx) error {
	var resent_eco_schedule dto.Resent_eco_schedule

	setDefaults(&resent_eco_schedule)

	if err := json.Unmarshal(c.Body(), &resent_eco_schedule); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	stmt, err := db.Prepare("insert into resent_eco_schedule values(:1, :2, :3, :4, :5, :6, :7, :8, :9, :10, :11, :12, :13, :14, :15, :16, :17, :18, :19, :20, :21, :22, :23, :24, :25, :26, :27, :28, :29, :30, :31, :32, :33)")
	if err != nil {
		return err
	}
	resent_eco_schedule.Row_created_date = formatDate(resent_eco_schedule.Row_created_date)
	resent_eco_schedule.Row_changed_date = formatDate(resent_eco_schedule.Row_changed_date)
	resent_eco_schedule.Effective_date = formatDateString(resent_eco_schedule.Effective_date)
	resent_eco_schedule.Expiry_date = formatDateString(resent_eco_schedule.Expiry_date)
	resent_eco_schedule.Schedule_date = formatDateString(resent_eco_schedule.Schedule_date)
	resent_eco_schedule.Row_effective_date = formatDateString(resent_eco_schedule.Row_effective_date)
	resent_eco_schedule.Row_expiry_date = formatDateString(resent_eco_schedule.Row_expiry_date)






	rows, err := stmt.Exec(resent_eco_schedule.Resent_id, resent_eco_schedule.Reserve_class_id, resent_eco_schedule.Economics_run_id, resent_eco_schedule.Schedule_id, resent_eco_schedule.Active_ind, resent_eco_schedule.Economic_schedule, resent_eco_schedule.Effective_date, resent_eco_schedule.Expiry_date, resent_eco_schedule.Max_value, resent_eco_schedule.Max_value_ouom, resent_eco_schedule.Max_value_uom, resent_eco_schedule.Min_value, resent_eco_schedule.Min_value_ouom, resent_eco_schedule.Min_value_uom, resent_eco_schedule.Period_type, resent_eco_schedule.Ppdm_guid, resent_eco_schedule.Product_ind, resent_eco_schedule.Product_type, resent_eco_schedule.Remark, resent_eco_schedule.Schedule_date, resent_eco_schedule.Schedule_date_desc, resent_eco_schedule.Schedule_desc, resent_eco_schedule.Schedule_value, resent_eco_schedule.Schedule_value_ouom, resent_eco_schedule.Schedule_value_uom, resent_eco_schedule.Source, resent_eco_schedule.Row_changed_by, resent_eco_schedule.Row_changed_date, resent_eco_schedule.Row_created_by, resent_eco_schedule.Row_created_date, resent_eco_schedule.Row_effective_date, resent_eco_schedule.Row_expiry_date, resent_eco_schedule.Row_quality)
	
	if err != nil {
		return err
	}

	return c.JSON(rows)
}

func UpdateResentEcoSchedule(c *fiber.Ctx) error {
	var resent_eco_schedule dto.Resent_eco_schedule

	setDefaults(&resent_eco_schedule)

	if err := json.Unmarshal(c.Body(), &resent_eco_schedule); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	id := c.Params("id")
	resent_eco_schedule.Resent_id = id

    if resent_eco_schedule.Row_created_date != nil {
        return c.Status(400).SendString("Cannot update row_created_date")
    }

	stmt, err := db.Prepare("update resent_eco_schedule set reserve_class_id = :1, economics_run_id = :2, schedule_id = :3, active_ind = :4, economic_schedule = :5, effective_date = :6, expiry_date = :7, max_value = :8, max_value_ouom = :9, max_value_uom = :10, min_value = :11, min_value_ouom = :12, min_value_uom = :13, period_type = :14, ppdm_guid = :15, product_ind = :16, product_type = :17, remark = :18, schedule_date = :19, schedule_date_desc = :20, schedule_desc = :21, schedule_value = :22, schedule_value_ouom = :23, schedule_value_uom = :24, source = :25, row_changed_by = :26, row_changed_date = :27, row_created_by = :28, row_effective_date = :29, row_expiry_date = :30, row_quality = :31 where resent_id = :33")
	if err != nil {
		return err
	}

	resent_eco_schedule.Row_changed_date = formatDate(resent_eco_schedule.Row_changed_date)
	resent_eco_schedule.Effective_date = formatDateString(resent_eco_schedule.Effective_date)
	resent_eco_schedule.Expiry_date = formatDateString(resent_eco_schedule.Expiry_date)
	resent_eco_schedule.Schedule_date = formatDateString(resent_eco_schedule.Schedule_date)
	resent_eco_schedule.Row_effective_date = formatDateString(resent_eco_schedule.Row_effective_date)
	resent_eco_schedule.Row_expiry_date = formatDateString(resent_eco_schedule.Row_expiry_date)






	rows, err := stmt.Exec(resent_eco_schedule.Reserve_class_id, resent_eco_schedule.Economics_run_id, resent_eco_schedule.Schedule_id, resent_eco_schedule.Active_ind, resent_eco_schedule.Economic_schedule, resent_eco_schedule.Effective_date, resent_eco_schedule.Expiry_date, resent_eco_schedule.Max_value, resent_eco_schedule.Max_value_ouom, resent_eco_schedule.Max_value_uom, resent_eco_schedule.Min_value, resent_eco_schedule.Min_value_ouom, resent_eco_schedule.Min_value_uom, resent_eco_schedule.Period_type, resent_eco_schedule.Ppdm_guid, resent_eco_schedule.Product_ind, resent_eco_schedule.Product_type, resent_eco_schedule.Remark, resent_eco_schedule.Schedule_date, resent_eco_schedule.Schedule_date_desc, resent_eco_schedule.Schedule_desc, resent_eco_schedule.Schedule_value, resent_eco_schedule.Schedule_value_ouom, resent_eco_schedule.Schedule_value_uom, resent_eco_schedule.Source, resent_eco_schedule.Row_changed_by, resent_eco_schedule.Row_changed_date, resent_eco_schedule.Row_created_by, resent_eco_schedule.Row_effective_date, resent_eco_schedule.Row_expiry_date, resent_eco_schedule.Row_quality, resent_eco_schedule.Resent_id)
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

func PatchResentEcoSchedule(c *fiber.Ctx) error {
	id := c.Params("id")

	updatedFields := make(map[string]interface{})

	if err := c.BodyParser(&updatedFields); err != nil {
		return err
	}
	if _, ok := updatedFields["row_created_date"]; ok {
        return c.Status(400).SendString("Cannot update row_created_date field")
    }

	query := "update resent_eco_schedule set "
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
		} else if key == "effective_date" || key == "expiry_date" || key == "schedule_date" || key == "row_effective_date" || key == "row_expiry_date"      {
			if date, ok := value.(string); ok && len(date) > 0 {
				formattedDate := formatDateString(&date)
				value = formattedDate
			}
		}
		values = append(values, value)
	}
	query += " where resent_id = :id"

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

func DeleteResentEcoSchedule(c *fiber.Ctx) error {
	id := c.Params("id")
	var resent_eco_schedule dto.Resent_eco_schedule
	resent_eco_schedule.Resent_id = id

	stmt, err := db.Prepare("delete from resent_eco_schedule where resent_id = :1")
	if err != nil {
		return err
	}

	rows, err := stmt.Exec(resent_eco_schedule.Resent_id)
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


