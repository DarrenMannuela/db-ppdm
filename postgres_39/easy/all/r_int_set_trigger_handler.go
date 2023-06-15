package apiv1

import (
	dto "github.com/DarrenMannuela/pt_gtn_bibliography/dto"
	"github.com/gofiber/fiber/v2"
	"strconv"
	"encoding/json"
)


func GetRIntSetTrigger(c *fiber.Ctx) error {
	rows, err := db.Query("select * from r_int_set_trigger")
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}
	defer rows.Close()
	var result []dto.R_int_set_trigger

	for rows.Next() {
		var r_int_set_trigger dto.R_int_set_trigger
		if err := rows.Scan(&r_int_set_trigger.Trigger_event, &r_int_set_trigger.Abbreviation, &r_int_set_trigger.Active_ind, &r_int_set_trigger.Effective_date, &r_int_set_trigger.Expiry_date, &r_int_set_trigger.Long_name, &r_int_set_trigger.Ppdm_guid, &r_int_set_trigger.Remark, &r_int_set_trigger.Short_name, &r_int_set_trigger.Source, &r_int_set_trigger.Row_changed_by, &r_int_set_trigger.Row_changed_date, &r_int_set_trigger.Row_created_by, &r_int_set_trigger.Row_created_date, &r_int_set_trigger.Row_effective_date, &r_int_set_trigger.Row_expiry_date, &r_int_set_trigger.Row_quality); err != nil {
			return err // Exit if we get an error
		}

		// Append r_int_set_trigger to result
		result = append(result, r_int_set_trigger)
	}
	// Return result in JSON format
	return c.JSON(result)
	 
}

func SetRIntSetTrigger(c *fiber.Ctx) error {
	var r_int_set_trigger dto.R_int_set_trigger

	setDefaults(&r_int_set_trigger)

	if err := json.Unmarshal(c.Body(), &r_int_set_trigger); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	stmt, err := db.Prepare("insert into r_int_set_trigger values(:1, :2, :3, :4, :5, :6, :7, :8, :9, :10, :11, :12, :13, :14, :15, :16, :17)")
	if err != nil {
		return err
	}
	r_int_set_trigger.Row_created_date = formatDate(r_int_set_trigger.Row_created_date)
	r_int_set_trigger.Row_changed_date = formatDate(r_int_set_trigger.Row_changed_date)
	r_int_set_trigger.Effective_date = formatDateString(r_int_set_trigger.Effective_date)
	r_int_set_trigger.Expiry_date = formatDateString(r_int_set_trigger.Expiry_date)
	r_int_set_trigger.Row_effective_date = formatDateString(r_int_set_trigger.Row_effective_date)
	r_int_set_trigger.Row_expiry_date = formatDateString(r_int_set_trigger.Row_expiry_date)






	rows, err := stmt.Exec(r_int_set_trigger.Trigger_event, r_int_set_trigger.Abbreviation, r_int_set_trigger.Active_ind, r_int_set_trigger.Effective_date, r_int_set_trigger.Expiry_date, r_int_set_trigger.Long_name, r_int_set_trigger.Ppdm_guid, r_int_set_trigger.Remark, r_int_set_trigger.Short_name, r_int_set_trigger.Source, r_int_set_trigger.Row_changed_by, r_int_set_trigger.Row_changed_date, r_int_set_trigger.Row_created_by, r_int_set_trigger.Row_created_date, r_int_set_trigger.Row_effective_date, r_int_set_trigger.Row_expiry_date, r_int_set_trigger.Row_quality)
	
	if err != nil {
		return err
	}

	return c.JSON(rows)
}

func UpdateRIntSetTrigger(c *fiber.Ctx) error {
	var r_int_set_trigger dto.R_int_set_trigger

	setDefaults(&r_int_set_trigger)

	if err := json.Unmarshal(c.Body(), &r_int_set_trigger); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	id := c.Params("id")
	r_int_set_trigger.Trigger_event = id

    if r_int_set_trigger.Row_created_date != nil {
        return c.Status(400).SendString("Cannot update row_created_date")
    }

	stmt, err := db.Prepare("update r_int_set_trigger set abbreviation = :1, active_ind = :2, effective_date = :3, expiry_date = :4, long_name = :5, ppdm_guid = :6, remark = :7, short_name = :8, source = :9, row_changed_by = :10, row_changed_date = :11, row_created_by = :12, row_effective_date = :13, row_expiry_date = :14, row_quality = :15 where trigger_event = :17")
	if err != nil {
		return err
	}

	r_int_set_trigger.Row_changed_date = formatDate(r_int_set_trigger.Row_changed_date)
	r_int_set_trigger.Effective_date = formatDateString(r_int_set_trigger.Effective_date)
	r_int_set_trigger.Expiry_date = formatDateString(r_int_set_trigger.Expiry_date)
	r_int_set_trigger.Row_effective_date = formatDateString(r_int_set_trigger.Row_effective_date)
	r_int_set_trigger.Row_expiry_date = formatDateString(r_int_set_trigger.Row_expiry_date)






	rows, err := stmt.Exec(r_int_set_trigger.Abbreviation, r_int_set_trigger.Active_ind, r_int_set_trigger.Effective_date, r_int_set_trigger.Expiry_date, r_int_set_trigger.Long_name, r_int_set_trigger.Ppdm_guid, r_int_set_trigger.Remark, r_int_set_trigger.Short_name, r_int_set_trigger.Source, r_int_set_trigger.Row_changed_by, r_int_set_trigger.Row_changed_date, r_int_set_trigger.Row_created_by, r_int_set_trigger.Row_effective_date, r_int_set_trigger.Row_expiry_date, r_int_set_trigger.Row_quality, r_int_set_trigger.Trigger_event)
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

func PatchRIntSetTrigger(c *fiber.Ctx) error {
	id := c.Params("id")

	updatedFields := make(map[string]interface{})

	if err := c.BodyParser(&updatedFields); err != nil {
		return err
	}
	if _, ok := updatedFields["row_created_date"]; ok {
        return c.Status(400).SendString("Cannot update row_created_date field")
    }

	query := "update r_int_set_trigger set "
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
	query += " where trigger_event = :id"

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

func DeleteRIntSetTrigger(c *fiber.Ctx) error {
	id := c.Params("id")
	var r_int_set_trigger dto.R_int_set_trigger
	r_int_set_trigger.Trigger_event = id

	stmt, err := db.Prepare("delete from r_int_set_trigger where trigger_event = :1")
	if err != nil {
		return err
	}

	rows, err := stmt.Exec(r_int_set_trigger.Trigger_event)
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


