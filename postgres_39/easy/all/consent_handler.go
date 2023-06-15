package apiv1

import (
	dto "github.com/DarrenMannuela/pt_gtn_bibliography/dto"
	"github.com/gofiber/fiber/v2"
	"strconv"
	"encoding/json"
)


func GetConsent(c *fiber.Ctx) error {
	rows, err := db.Query("select * from consent")
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}
	defer rows.Close()
	var result []dto.Consent

	for rows.Next() {
		var consent dto.Consent
		if err := rows.Scan(&consent.Consent_id, &consent.Active_ind, &consent.Consent_desc, &consent.Consent_method_desc, &consent.Consent_type, &consent.Current_status, &consent.Effective_date, &consent.Expiry_date, &consent.Ppdm_guid, &consent.Receive_date, &consent.Remark, &consent.Request_date, &consent.Source, &consent.Status_remark, &consent.Row_changed_by, &consent.Row_changed_date, &consent.Row_created_by, &consent.Row_created_date, &consent.Row_effective_date, &consent.Row_expiry_date, &consent.Row_quality); err != nil {
			return err // Exit if we get an error
		}

		// Append consent to result
		result = append(result, consent)
	}
	// Return result in JSON format
	return c.JSON(result)
	 
}

func SetConsent(c *fiber.Ctx) error {
	var consent dto.Consent

	setDefaults(&consent)

	if err := json.Unmarshal(c.Body(), &consent); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	stmt, err := db.Prepare("insert into consent values(:1, :2, :3, :4, :5, :6, :7, :8, :9, :10, :11, :12, :13, :14, :15, :16, :17, :18, :19, :20, :21)")
	if err != nil {
		return err
	}
	consent.Row_created_date = formatDate(consent.Row_created_date)
	consent.Row_changed_date = formatDate(consent.Row_changed_date)
	consent.Effective_date = formatDateString(consent.Effective_date)
	consent.Expiry_date = formatDateString(consent.Expiry_date)
	consent.Receive_date = formatDateString(consent.Receive_date)
	consent.Request_date = formatDateString(consent.Request_date)
	consent.Row_effective_date = formatDateString(consent.Row_effective_date)
	consent.Row_expiry_date = formatDateString(consent.Row_expiry_date)






	rows, err := stmt.Exec(consent.Consent_id, consent.Active_ind, consent.Consent_desc, consent.Consent_method_desc, consent.Consent_type, consent.Current_status, consent.Effective_date, consent.Expiry_date, consent.Ppdm_guid, consent.Receive_date, consent.Remark, consent.Request_date, consent.Source, consent.Status_remark, consent.Row_changed_by, consent.Row_changed_date, consent.Row_created_by, consent.Row_created_date, consent.Row_effective_date, consent.Row_expiry_date, consent.Row_quality)
	
	if err != nil {
		return err
	}

	return c.JSON(rows)
}

func UpdateConsent(c *fiber.Ctx) error {
	var consent dto.Consent

	setDefaults(&consent)

	if err := json.Unmarshal(c.Body(), &consent); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	id := c.Params("id")
	consent.Consent_id = id

    if consent.Row_created_date != nil {
        return c.Status(400).SendString("Cannot update row_created_date")
    }

	stmt, err := db.Prepare("update consent set active_ind = :1, consent_desc = :2, consent_method_desc = :3, consent_type = :4, current_status = :5, effective_date = :6, expiry_date = :7, ppdm_guid = :8, receive_date = :9, remark = :10, request_date = :11, source = :12, status_remark = :13, row_changed_by = :14, row_changed_date = :15, row_created_by = :16, row_effective_date = :17, row_expiry_date = :18, row_quality = :19 where consent_id = :21")
	if err != nil {
		return err
	}

	consent.Row_changed_date = formatDate(consent.Row_changed_date)
	consent.Effective_date = formatDateString(consent.Effective_date)
	consent.Expiry_date = formatDateString(consent.Expiry_date)
	consent.Receive_date = formatDateString(consent.Receive_date)
	consent.Request_date = formatDateString(consent.Request_date)
	consent.Row_effective_date = formatDateString(consent.Row_effective_date)
	consent.Row_expiry_date = formatDateString(consent.Row_expiry_date)






	rows, err := stmt.Exec(consent.Active_ind, consent.Consent_desc, consent.Consent_method_desc, consent.Consent_type, consent.Current_status, consent.Effective_date, consent.Expiry_date, consent.Ppdm_guid, consent.Receive_date, consent.Remark, consent.Request_date, consent.Source, consent.Status_remark, consent.Row_changed_by, consent.Row_changed_date, consent.Row_created_by, consent.Row_effective_date, consent.Row_expiry_date, consent.Row_quality, consent.Consent_id)
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

func PatchConsent(c *fiber.Ctx) error {
	id := c.Params("id")

	updatedFields := make(map[string]interface{})

	if err := c.BodyParser(&updatedFields); err != nil {
		return err
	}
	if _, ok := updatedFields["row_created_date"]; ok {
        return c.Status(400).SendString("Cannot update row_created_date field")
    }

	query := "update consent set "
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
		} else if key == "effective_date" || key == "expiry_date" || key == "receive_date" || key == "request_date" || key == "row_effective_date" || key == "row_expiry_date"      {
			if date, ok := value.(string); ok && len(date) > 0 {
				formattedDate := formatDateString(&date)
				value = formattedDate
			}
		}
		values = append(values, value)
	}
	query += " where consent_id = :id"

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

func DeleteConsent(c *fiber.Ctx) error {
	id := c.Params("id")
	var consent dto.Consent
	consent.Consent_id = id

	stmt, err := db.Prepare("delete from consent where consent_id = :1")
	if err != nil {
		return err
	}

	rows, err := stmt.Exec(consent.Consent_id)
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


