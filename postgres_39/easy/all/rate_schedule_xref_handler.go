package apiv1

import (
	dto "github.com/DarrenMannuela/pt_gtn_bibliography/dto"
	"github.com/gofiber/fiber/v2"
	"strconv"
	"encoding/json"
)


func GetRateScheduleXref(c *fiber.Ctx) error {
	rows, err := db.Query("select * from rate_schedule_xref")
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}
	defer rows.Close()
	var result []dto.Rate_schedule_xref

	for rows.Next() {
		var rate_schedule_xref dto.Rate_schedule_xref
		if err := rows.Scan(&rate_schedule_xref.Rate_schedule_id_1, &rate_schedule_xref.Rate_schedule_id_2, &rate_schedule_xref.Active_ind, &rate_schedule_xref.Effective_date, &rate_schedule_xref.Expiry_date, &rate_schedule_xref.Ppdm_guid, &rate_schedule_xref.Rate_schedule_xref_type, &rate_schedule_xref.Remark, &rate_schedule_xref.Source, &rate_schedule_xref.Row_changed_by, &rate_schedule_xref.Row_changed_date, &rate_schedule_xref.Row_created_by, &rate_schedule_xref.Row_created_date, &rate_schedule_xref.Row_effective_date, &rate_schedule_xref.Row_expiry_date, &rate_schedule_xref.Row_quality); err != nil {
			return err // Exit if we get an error
		}

		// Append rate_schedule_xref to result
		result = append(result, rate_schedule_xref)
	}
	// Return result in JSON format
	return c.JSON(result)
	 
}

func SetRateScheduleXref(c *fiber.Ctx) error {
	var rate_schedule_xref dto.Rate_schedule_xref

	setDefaults(&rate_schedule_xref)

	if err := json.Unmarshal(c.Body(), &rate_schedule_xref); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	stmt, err := db.Prepare("insert into rate_schedule_xref values(:1, :2, :3, :4, :5, :6, :7, :8, :9, :10, :11, :12, :13, :14, :15, :16)")
	if err != nil {
		return err
	}
	rate_schedule_xref.Row_created_date = formatDate(rate_schedule_xref.Row_created_date)
	rate_schedule_xref.Row_changed_date = formatDate(rate_schedule_xref.Row_changed_date)
	rate_schedule_xref.Effective_date = formatDateString(rate_schedule_xref.Effective_date)
	rate_schedule_xref.Expiry_date = formatDateString(rate_schedule_xref.Expiry_date)
	rate_schedule_xref.Row_effective_date = formatDateString(rate_schedule_xref.Row_effective_date)
	rate_schedule_xref.Row_expiry_date = formatDateString(rate_schedule_xref.Row_expiry_date)






	rows, err := stmt.Exec(rate_schedule_xref.Rate_schedule_id_1, rate_schedule_xref.Rate_schedule_id_2, rate_schedule_xref.Active_ind, rate_schedule_xref.Effective_date, rate_schedule_xref.Expiry_date, rate_schedule_xref.Ppdm_guid, rate_schedule_xref.Rate_schedule_xref_type, rate_schedule_xref.Remark, rate_schedule_xref.Source, rate_schedule_xref.Row_changed_by, rate_schedule_xref.Row_changed_date, rate_schedule_xref.Row_created_by, rate_schedule_xref.Row_created_date, rate_schedule_xref.Row_effective_date, rate_schedule_xref.Row_expiry_date, rate_schedule_xref.Row_quality)
	
	if err != nil {
		return err
	}

	return c.JSON(rows)
}

func UpdateRateScheduleXref(c *fiber.Ctx) error {
	var rate_schedule_xref dto.Rate_schedule_xref

	setDefaults(&rate_schedule_xref)

	if err := json.Unmarshal(c.Body(), &rate_schedule_xref); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	id := c.Params("id")
	rate_schedule_xref.Rate_schedule_id_1 = id

    if rate_schedule_xref.Row_created_date != nil {
        return c.Status(400).SendString("Cannot update row_created_date")
    }

	stmt, err := db.Prepare("update rate_schedule_xref set rate_schedule_id_2 = :1, active_ind = :2, effective_date = :3, expiry_date = :4, ppdm_guid = :5, rate_schedule_xref_type = :6, remark = :7, source = :8, row_changed_by = :9, row_changed_date = :10, row_created_by = :11, row_effective_date = :12, row_expiry_date = :13, row_quality = :14 where rate_schedule_id_1 = :16")
	if err != nil {
		return err
	}

	rate_schedule_xref.Row_changed_date = formatDate(rate_schedule_xref.Row_changed_date)
	rate_schedule_xref.Effective_date = formatDateString(rate_schedule_xref.Effective_date)
	rate_schedule_xref.Expiry_date = formatDateString(rate_schedule_xref.Expiry_date)
	rate_schedule_xref.Row_effective_date = formatDateString(rate_schedule_xref.Row_effective_date)
	rate_schedule_xref.Row_expiry_date = formatDateString(rate_schedule_xref.Row_expiry_date)






	rows, err := stmt.Exec(rate_schedule_xref.Rate_schedule_id_2, rate_schedule_xref.Active_ind, rate_schedule_xref.Effective_date, rate_schedule_xref.Expiry_date, rate_schedule_xref.Ppdm_guid, rate_schedule_xref.Rate_schedule_xref_type, rate_schedule_xref.Remark, rate_schedule_xref.Source, rate_schedule_xref.Row_changed_by, rate_schedule_xref.Row_changed_date, rate_schedule_xref.Row_created_by, rate_schedule_xref.Row_effective_date, rate_schedule_xref.Row_expiry_date, rate_schedule_xref.Row_quality, rate_schedule_xref.Rate_schedule_id_1)
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

func PatchRateScheduleXref(c *fiber.Ctx) error {
	id := c.Params("id")

	updatedFields := make(map[string]interface{})

	if err := c.BodyParser(&updatedFields); err != nil {
		return err
	}
	if _, ok := updatedFields["row_created_date"]; ok {
        return c.Status(400).SendString("Cannot update row_created_date field")
    }

	query := "update rate_schedule_xref set "
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
	query += " where rate_schedule_id_1 = :id"

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

func DeleteRateScheduleXref(c *fiber.Ctx) error {
	id := c.Params("id")
	var rate_schedule_xref dto.Rate_schedule_xref
	rate_schedule_xref.Rate_schedule_id_1 = id

	stmt, err := db.Prepare("delete from rate_schedule_xref where rate_schedule_id_1 = :1")
	if err != nil {
		return err
	}

	rows, err := stmt.Exec(rate_schedule_xref.Rate_schedule_id_1)
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


