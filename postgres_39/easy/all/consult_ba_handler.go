package apiv1

import (
	dto "github.com/DarrenMannuela/pt_gtn_bibliography/dto"
	"github.com/gofiber/fiber/v2"
	"strconv"
	"encoding/json"
)


func GetConsultBa(c *fiber.Ctx) error {
	rows, err := db.Query("select * from consult_ba")
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}
	defer rows.Close()
	var result []dto.Consult_ba

	for rows.Next() {
		var consult_ba dto.Consult_ba
		if err := rows.Scan(&consult_ba.Consult_id, &consult_ba.Business_associate_id, &consult_ba.Active_ind, &consult_ba.Consult_role, &consult_ba.Effective_date, &consult_ba.Expiry_date, &consult_ba.Ppdm_guid, &consult_ba.Remark, &consult_ba.Represented_ba_id, &consult_ba.Source, &consult_ba.Row_changed_by, &consult_ba.Row_changed_date, &consult_ba.Row_created_by, &consult_ba.Row_created_date, &consult_ba.Row_effective_date, &consult_ba.Row_expiry_date, &consult_ba.Row_quality); err != nil {
			return err // Exit if we get an error
		}

		// Append consult_ba to result
		result = append(result, consult_ba)
	}
	// Return result in JSON format
	return c.JSON(result)
	 
}

func SetConsultBa(c *fiber.Ctx) error {
	var consult_ba dto.Consult_ba

	setDefaults(&consult_ba)

	if err := json.Unmarshal(c.Body(), &consult_ba); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	stmt, err := db.Prepare("insert into consult_ba values(:1, :2, :3, :4, :5, :6, :7, :8, :9, :10, :11, :12, :13, :14, :15, :16, :17)")
	if err != nil {
		return err
	}
	consult_ba.Row_created_date = formatDate(consult_ba.Row_created_date)
	consult_ba.Row_changed_date = formatDate(consult_ba.Row_changed_date)
	consult_ba.Effective_date = formatDateString(consult_ba.Effective_date)
	consult_ba.Expiry_date = formatDateString(consult_ba.Expiry_date)
	consult_ba.Row_effective_date = formatDateString(consult_ba.Row_effective_date)
	consult_ba.Row_expiry_date = formatDateString(consult_ba.Row_expiry_date)






	rows, err := stmt.Exec(consult_ba.Consult_id, consult_ba.Business_associate_id, consult_ba.Active_ind, consult_ba.Consult_role, consult_ba.Effective_date, consult_ba.Expiry_date, consult_ba.Ppdm_guid, consult_ba.Remark, consult_ba.Represented_ba_id, consult_ba.Source, consult_ba.Row_changed_by, consult_ba.Row_changed_date, consult_ba.Row_created_by, consult_ba.Row_created_date, consult_ba.Row_effective_date, consult_ba.Row_expiry_date, consult_ba.Row_quality)
	
	if err != nil {
		return err
	}

	return c.JSON(rows)
}

func UpdateConsultBa(c *fiber.Ctx) error {
	var consult_ba dto.Consult_ba

	setDefaults(&consult_ba)

	if err := json.Unmarshal(c.Body(), &consult_ba); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	id := c.Params("id")
	consult_ba.Consult_id = id

    if consult_ba.Row_created_date != nil {
        return c.Status(400).SendString("Cannot update row_created_date")
    }

	stmt, err := db.Prepare("update consult_ba set business_associate_id = :1, active_ind = :2, consult_role = :3, effective_date = :4, expiry_date = :5, ppdm_guid = :6, remark = :7, represented_ba_id = :8, source = :9, row_changed_by = :10, row_changed_date = :11, row_created_by = :12, row_effective_date = :13, row_expiry_date = :14, row_quality = :15 where consult_id = :17")
	if err != nil {
		return err
	}

	consult_ba.Row_changed_date = formatDate(consult_ba.Row_changed_date)
	consult_ba.Effective_date = formatDateString(consult_ba.Effective_date)
	consult_ba.Expiry_date = formatDateString(consult_ba.Expiry_date)
	consult_ba.Row_effective_date = formatDateString(consult_ba.Row_effective_date)
	consult_ba.Row_expiry_date = formatDateString(consult_ba.Row_expiry_date)






	rows, err := stmt.Exec(consult_ba.Business_associate_id, consult_ba.Active_ind, consult_ba.Consult_role, consult_ba.Effective_date, consult_ba.Expiry_date, consult_ba.Ppdm_guid, consult_ba.Remark, consult_ba.Represented_ba_id, consult_ba.Source, consult_ba.Row_changed_by, consult_ba.Row_changed_date, consult_ba.Row_created_by, consult_ba.Row_effective_date, consult_ba.Row_expiry_date, consult_ba.Row_quality, consult_ba.Consult_id)
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

func PatchConsultBa(c *fiber.Ctx) error {
	id := c.Params("id")

	updatedFields := make(map[string]interface{})

	if err := c.BodyParser(&updatedFields); err != nil {
		return err
	}
	if _, ok := updatedFields["row_created_date"]; ok {
        return c.Status(400).SendString("Cannot update row_created_date field")
    }

	query := "update consult_ba set "
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

func DeleteConsultBa(c *fiber.Ctx) error {
	id := c.Params("id")
	var consult_ba dto.Consult_ba
	consult_ba.Consult_id = id

	stmt, err := db.Prepare("delete from consult_ba where consult_id = :1")
	if err != nil {
		return err
	}

	rows, err := stmt.Exec(consult_ba.Consult_id)
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


