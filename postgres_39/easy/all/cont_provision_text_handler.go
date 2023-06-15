package apiv1

import (
	dto "github.com/DarrenMannuela/pt_gtn_bibliography/dto"
	"github.com/gofiber/fiber/v2"
	"strconv"
	"encoding/json"
)


func GetContProvisionText(c *fiber.Ctx) error {
	rows, err := db.Query("select * from cont_provision_text")
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}
	defer rows.Close()
	var result []dto.Cont_provision_text

	for rows.Next() {
		var cont_provision_text dto.Cont_provision_text
		if err := rows.Scan(&cont_provision_text.Contract_id, &cont_provision_text.Provision_id, &cont_provision_text.Provision_text_seq_no, &cont_provision_text.Active_ind, &cont_provision_text.Effective_date, &cont_provision_text.Expiry_date, &cont_provision_text.Ppdm_guid, &cont_provision_text.Provision_text, &cont_provision_text.Remark, &cont_provision_text.Source, &cont_provision_text.Row_changed_by, &cont_provision_text.Row_changed_date, &cont_provision_text.Row_created_by, &cont_provision_text.Row_created_date, &cont_provision_text.Row_effective_date, &cont_provision_text.Row_expiry_date, &cont_provision_text.Row_quality); err != nil {
			return err // Exit if we get an error
		}

		// Append cont_provision_text to result
		result = append(result, cont_provision_text)
	}
	// Return result in JSON format
	return c.JSON(result)
	 
}

func SetContProvisionText(c *fiber.Ctx) error {
	var cont_provision_text dto.Cont_provision_text

	setDefaults(&cont_provision_text)

	if err := json.Unmarshal(c.Body(), &cont_provision_text); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	stmt, err := db.Prepare("insert into cont_provision_text values(:1, :2, :3, :4, :5, :6, :7, :8, :9, :10, :11, :12, :13, :14, :15, :16, :17)")
	if err != nil {
		return err
	}
	cont_provision_text.Row_created_date = formatDate(cont_provision_text.Row_created_date)
	cont_provision_text.Row_changed_date = formatDate(cont_provision_text.Row_changed_date)
	cont_provision_text.Effective_date = formatDateString(cont_provision_text.Effective_date)
	cont_provision_text.Expiry_date = formatDateString(cont_provision_text.Expiry_date)
	cont_provision_text.Row_effective_date = formatDateString(cont_provision_text.Row_effective_date)
	cont_provision_text.Row_expiry_date = formatDateString(cont_provision_text.Row_expiry_date)






	rows, err := stmt.Exec(cont_provision_text.Contract_id, cont_provision_text.Provision_id, cont_provision_text.Provision_text_seq_no, cont_provision_text.Active_ind, cont_provision_text.Effective_date, cont_provision_text.Expiry_date, cont_provision_text.Ppdm_guid, cont_provision_text.Provision_text, cont_provision_text.Remark, cont_provision_text.Source, cont_provision_text.Row_changed_by, cont_provision_text.Row_changed_date, cont_provision_text.Row_created_by, cont_provision_text.Row_created_date, cont_provision_text.Row_effective_date, cont_provision_text.Row_expiry_date, cont_provision_text.Row_quality)
	
	if err != nil {
		return err
	}

	return c.JSON(rows)
}

func UpdateContProvisionText(c *fiber.Ctx) error {
	var cont_provision_text dto.Cont_provision_text

	setDefaults(&cont_provision_text)

	if err := json.Unmarshal(c.Body(), &cont_provision_text); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	id := c.Params("id")
	cont_provision_text.Contract_id = id

    if cont_provision_text.Row_created_date != nil {
        return c.Status(400).SendString("Cannot update row_created_date")
    }

	stmt, err := db.Prepare("update cont_provision_text set provision_id = :1, provision_text_seq_no = :2, active_ind = :3, effective_date = :4, expiry_date = :5, ppdm_guid = :6, provision_text = :7, remark = :8, source = :9, row_changed_by = :10, row_changed_date = :11, row_created_by = :12, row_effective_date = :13, row_expiry_date = :14, row_quality = :15 where contract_id = :17")
	if err != nil {
		return err
	}

	cont_provision_text.Row_changed_date = formatDate(cont_provision_text.Row_changed_date)
	cont_provision_text.Effective_date = formatDateString(cont_provision_text.Effective_date)
	cont_provision_text.Expiry_date = formatDateString(cont_provision_text.Expiry_date)
	cont_provision_text.Row_effective_date = formatDateString(cont_provision_text.Row_effective_date)
	cont_provision_text.Row_expiry_date = formatDateString(cont_provision_text.Row_expiry_date)






	rows, err := stmt.Exec(cont_provision_text.Provision_id, cont_provision_text.Provision_text_seq_no, cont_provision_text.Active_ind, cont_provision_text.Effective_date, cont_provision_text.Expiry_date, cont_provision_text.Ppdm_guid, cont_provision_text.Provision_text, cont_provision_text.Remark, cont_provision_text.Source, cont_provision_text.Row_changed_by, cont_provision_text.Row_changed_date, cont_provision_text.Row_created_by, cont_provision_text.Row_effective_date, cont_provision_text.Row_expiry_date, cont_provision_text.Row_quality, cont_provision_text.Contract_id)
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

func PatchContProvisionText(c *fiber.Ctx) error {
	id := c.Params("id")

	updatedFields := make(map[string]interface{})

	if err := c.BodyParser(&updatedFields); err != nil {
		return err
	}
	if _, ok := updatedFields["row_created_date"]; ok {
        return c.Status(400).SendString("Cannot update row_created_date field")
    }

	query := "update cont_provision_text set "
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
	query += " where contract_id = :id"

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

func DeleteContProvisionText(c *fiber.Ctx) error {
	id := c.Params("id")
	var cont_provision_text dto.Cont_provision_text
	cont_provision_text.Contract_id = id

	stmt, err := db.Prepare("delete from cont_provision_text where contract_id = :1")
	if err != nil {
		return err
	}

	rows, err := stmt.Exec(cont_provision_text.Contract_id)
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


