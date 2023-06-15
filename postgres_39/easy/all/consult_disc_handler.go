package apiv1

import (
	dto "github.com/DarrenMannuela/pt_gtn_bibliography/dto"
	"github.com/gofiber/fiber/v2"
	"strconv"
	"encoding/json"
)


func GetConsultDisc(c *fiber.Ctx) error {
	rows, err := db.Query("select * from consult_disc")
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}
	defer rows.Close()
	var result []dto.Consult_disc

	for rows.Next() {
		var consult_disc dto.Consult_disc
		if err := rows.Scan(&consult_disc.Consult_id, &consult_disc.Discussion_id, &consult_disc.Active_ind, &consult_disc.Complete_date, &consult_disc.Discussion_type, &consult_disc.Effective_date, &consult_disc.Expiry_date, &consult_disc.Ppdm_guid, &consult_disc.Remark, &consult_disc.Source, &consult_disc.Start_date, &consult_disc.Row_changed_by, &consult_disc.Row_changed_date, &consult_disc.Row_created_by, &consult_disc.Row_created_date, &consult_disc.Row_effective_date, &consult_disc.Row_expiry_date, &consult_disc.Row_quality); err != nil {
			return err // Exit if we get an error
		}

		// Append consult_disc to result
		result = append(result, consult_disc)
	}
	// Return result in JSON format
	return c.JSON(result)
	 
}

func SetConsultDisc(c *fiber.Ctx) error {
	var consult_disc dto.Consult_disc

	setDefaults(&consult_disc)

	if err := json.Unmarshal(c.Body(), &consult_disc); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	stmt, err := db.Prepare("insert into consult_disc values(:1, :2, :3, :4, :5, :6, :7, :8, :9, :10, :11, :12, :13, :14, :15, :16, :17, :18)")
	if err != nil {
		return err
	}
	consult_disc.Row_created_date = formatDate(consult_disc.Row_created_date)
	consult_disc.Row_changed_date = formatDate(consult_disc.Row_changed_date)
	consult_disc.Complete_date = formatDateString(consult_disc.Complete_date)
	consult_disc.Effective_date = formatDateString(consult_disc.Effective_date)
	consult_disc.Expiry_date = formatDateString(consult_disc.Expiry_date)
	consult_disc.Start_date = formatDateString(consult_disc.Start_date)
	consult_disc.Row_effective_date = formatDateString(consult_disc.Row_effective_date)
	consult_disc.Row_expiry_date = formatDateString(consult_disc.Row_expiry_date)






	rows, err := stmt.Exec(consult_disc.Consult_id, consult_disc.Discussion_id, consult_disc.Active_ind, consult_disc.Complete_date, consult_disc.Discussion_type, consult_disc.Effective_date, consult_disc.Expiry_date, consult_disc.Ppdm_guid, consult_disc.Remark, consult_disc.Source, consult_disc.Start_date, consult_disc.Row_changed_by, consult_disc.Row_changed_date, consult_disc.Row_created_by, consult_disc.Row_created_date, consult_disc.Row_effective_date, consult_disc.Row_expiry_date, consult_disc.Row_quality)
	
	if err != nil {
		return err
	}

	return c.JSON(rows)
}

func UpdateConsultDisc(c *fiber.Ctx) error {
	var consult_disc dto.Consult_disc

	setDefaults(&consult_disc)

	if err := json.Unmarshal(c.Body(), &consult_disc); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	id := c.Params("id")
	consult_disc.Consult_id = id

    if consult_disc.Row_created_date != nil {
        return c.Status(400).SendString("Cannot update row_created_date")
    }

	stmt, err := db.Prepare("update consult_disc set discussion_id = :1, active_ind = :2, complete_date = :3, discussion_type = :4, effective_date = :5, expiry_date = :6, ppdm_guid = :7, remark = :8, source = :9, start_date = :10, row_changed_by = :11, row_changed_date = :12, row_created_by = :13, row_effective_date = :14, row_expiry_date = :15, row_quality = :16 where consult_id = :18")
	if err != nil {
		return err
	}

	consult_disc.Row_changed_date = formatDate(consult_disc.Row_changed_date)
	consult_disc.Complete_date = formatDateString(consult_disc.Complete_date)
	consult_disc.Effective_date = formatDateString(consult_disc.Effective_date)
	consult_disc.Expiry_date = formatDateString(consult_disc.Expiry_date)
	consult_disc.Start_date = formatDateString(consult_disc.Start_date)
	consult_disc.Row_effective_date = formatDateString(consult_disc.Row_effective_date)
	consult_disc.Row_expiry_date = formatDateString(consult_disc.Row_expiry_date)






	rows, err := stmt.Exec(consult_disc.Discussion_id, consult_disc.Active_ind, consult_disc.Complete_date, consult_disc.Discussion_type, consult_disc.Effective_date, consult_disc.Expiry_date, consult_disc.Ppdm_guid, consult_disc.Remark, consult_disc.Source, consult_disc.Start_date, consult_disc.Row_changed_by, consult_disc.Row_changed_date, consult_disc.Row_created_by, consult_disc.Row_effective_date, consult_disc.Row_expiry_date, consult_disc.Row_quality, consult_disc.Consult_id)
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

func PatchConsultDisc(c *fiber.Ctx) error {
	id := c.Params("id")

	updatedFields := make(map[string]interface{})

	if err := c.BodyParser(&updatedFields); err != nil {
		return err
	}
	if _, ok := updatedFields["row_created_date"]; ok {
        return c.Status(400).SendString("Cannot update row_created_date field")
    }

	query := "update consult_disc set "
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
		} else if key == "complete_date" || key == "effective_date" || key == "expiry_date" || key == "start_date" || key == "row_effective_date" || key == "row_expiry_date"      {
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

func DeleteConsultDisc(c *fiber.Ctx) error {
	id := c.Params("id")
	var consult_disc dto.Consult_disc
	consult_disc.Consult_id = id

	stmt, err := db.Prepare("delete from consult_disc where consult_id = :1")
	if err != nil {
		return err
	}

	rows, err := stmt.Exec(consult_disc.Consult_id)
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


