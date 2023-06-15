package apiv1

import (
	dto "github.com/DarrenMannuela/pt_gtn_bibliography/dto"
	"github.com/gofiber/fiber/v2"
	"strconv"
	"encoding/json"
)


func GetConsultDiscBa(c *fiber.Ctx) error {
	rows, err := db.Query("select * from consult_disc_ba")
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}
	defer rows.Close()
	var result []dto.Consult_disc_ba

	for rows.Next() {
		var consult_disc_ba dto.Consult_disc_ba
		if err := rows.Scan(&consult_disc_ba.Consult_id, &consult_disc_ba.Discussion_id, &consult_disc_ba.Business_associate_id, &consult_disc_ba.Active_ind, &consult_disc_ba.Attend_type, &consult_disc_ba.Effective_date, &consult_disc_ba.End_date, &consult_disc_ba.Expiry_date, &consult_disc_ba.Ppdm_guid, &consult_disc_ba.Remark, &consult_disc_ba.Source, &consult_disc_ba.Start_date, &consult_disc_ba.Row_changed_by, &consult_disc_ba.Row_changed_date, &consult_disc_ba.Row_created_by, &consult_disc_ba.Row_created_date, &consult_disc_ba.Row_effective_date, &consult_disc_ba.Row_expiry_date, &consult_disc_ba.Row_quality); err != nil {
			return err // Exit if we get an error
		}

		// Append consult_disc_ba to result
		result = append(result, consult_disc_ba)
	}
	// Return result in JSON format
	return c.JSON(result)
	 
}

func SetConsultDiscBa(c *fiber.Ctx) error {
	var consult_disc_ba dto.Consult_disc_ba

	setDefaults(&consult_disc_ba)

	if err := json.Unmarshal(c.Body(), &consult_disc_ba); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	stmt, err := db.Prepare("insert into consult_disc_ba values(:1, :2, :3, :4, :5, :6, :7, :8, :9, :10, :11, :12, :13, :14, :15, :16, :17, :18, :19)")
	if err != nil {
		return err
	}
	consult_disc_ba.Row_created_date = formatDate(consult_disc_ba.Row_created_date)
	consult_disc_ba.Row_changed_date = formatDate(consult_disc_ba.Row_changed_date)
	consult_disc_ba.Effective_date = formatDateString(consult_disc_ba.Effective_date)
	consult_disc_ba.End_date = formatDateString(consult_disc_ba.End_date)
	consult_disc_ba.Expiry_date = formatDateString(consult_disc_ba.Expiry_date)
	consult_disc_ba.Start_date = formatDateString(consult_disc_ba.Start_date)
	consult_disc_ba.Row_effective_date = formatDateString(consult_disc_ba.Row_effective_date)
	consult_disc_ba.Row_expiry_date = formatDateString(consult_disc_ba.Row_expiry_date)






	rows, err := stmt.Exec(consult_disc_ba.Consult_id, consult_disc_ba.Discussion_id, consult_disc_ba.Business_associate_id, consult_disc_ba.Active_ind, consult_disc_ba.Attend_type, consult_disc_ba.Effective_date, consult_disc_ba.End_date, consult_disc_ba.Expiry_date, consult_disc_ba.Ppdm_guid, consult_disc_ba.Remark, consult_disc_ba.Source, consult_disc_ba.Start_date, consult_disc_ba.Row_changed_by, consult_disc_ba.Row_changed_date, consult_disc_ba.Row_created_by, consult_disc_ba.Row_created_date, consult_disc_ba.Row_effective_date, consult_disc_ba.Row_expiry_date, consult_disc_ba.Row_quality)
	
	if err != nil {
		return err
	}

	return c.JSON(rows)
}

func UpdateConsultDiscBa(c *fiber.Ctx) error {
	var consult_disc_ba dto.Consult_disc_ba

	setDefaults(&consult_disc_ba)

	if err := json.Unmarshal(c.Body(), &consult_disc_ba); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	id := c.Params("id")
	consult_disc_ba.Consult_id = id

    if consult_disc_ba.Row_created_date != nil {
        return c.Status(400).SendString("Cannot update row_created_date")
    }

	stmt, err := db.Prepare("update consult_disc_ba set discussion_id = :1, business_associate_id = :2, active_ind = :3, attend_type = :4, effective_date = :5, end_date = :6, expiry_date = :7, ppdm_guid = :8, remark = :9, source = :10, start_date = :11, row_changed_by = :12, row_changed_date = :13, row_created_by = :14, row_effective_date = :15, row_expiry_date = :16, row_quality = :17 where consult_id = :19")
	if err != nil {
		return err
	}

	consult_disc_ba.Row_changed_date = formatDate(consult_disc_ba.Row_changed_date)
	consult_disc_ba.Effective_date = formatDateString(consult_disc_ba.Effective_date)
	consult_disc_ba.End_date = formatDateString(consult_disc_ba.End_date)
	consult_disc_ba.Expiry_date = formatDateString(consult_disc_ba.Expiry_date)
	consult_disc_ba.Start_date = formatDateString(consult_disc_ba.Start_date)
	consult_disc_ba.Row_effective_date = formatDateString(consult_disc_ba.Row_effective_date)
	consult_disc_ba.Row_expiry_date = formatDateString(consult_disc_ba.Row_expiry_date)






	rows, err := stmt.Exec(consult_disc_ba.Discussion_id, consult_disc_ba.Business_associate_id, consult_disc_ba.Active_ind, consult_disc_ba.Attend_type, consult_disc_ba.Effective_date, consult_disc_ba.End_date, consult_disc_ba.Expiry_date, consult_disc_ba.Ppdm_guid, consult_disc_ba.Remark, consult_disc_ba.Source, consult_disc_ba.Start_date, consult_disc_ba.Row_changed_by, consult_disc_ba.Row_changed_date, consult_disc_ba.Row_created_by, consult_disc_ba.Row_effective_date, consult_disc_ba.Row_expiry_date, consult_disc_ba.Row_quality, consult_disc_ba.Consult_id)
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

func PatchConsultDiscBa(c *fiber.Ctx) error {
	id := c.Params("id")

	updatedFields := make(map[string]interface{})

	if err := c.BodyParser(&updatedFields); err != nil {
		return err
	}
	if _, ok := updatedFields["row_created_date"]; ok {
        return c.Status(400).SendString("Cannot update row_created_date field")
    }

	query := "update consult_disc_ba set "
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
		} else if key == "effective_date" || key == "end_date" || key == "expiry_date" || key == "start_date" || key == "row_effective_date" || key == "row_expiry_date"      {
			if date, ok := value.(string); ok && len(date) > 0 {
				formattedDate := formatDateString(&date)
				value = formattedDate
			}
		}
		values = append(values, value)
	}
	query += " where consult_id = :id"

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

func DeleteConsultDiscBa(c *fiber.Ctx) error {
	id := c.Params("id")
	var consult_disc_ba dto.Consult_disc_ba
	consult_disc_ba.Consult_id = id

	stmt, err := db.Prepare("delete from consult_disc_ba where consult_id = :1")
	if err != nil {
		return err
	}

	rows, err := stmt.Exec(consult_disc_ba.Consult_id)
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


