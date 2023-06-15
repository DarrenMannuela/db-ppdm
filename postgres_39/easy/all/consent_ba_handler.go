package apiv1

import (
	dto "github.com/DarrenMannuela/pt_gtn_bibliography/dto"
	"github.com/gofiber/fiber/v2"
	"strconv"
	"encoding/json"
)


func GetConsentBa(c *fiber.Ctx) error {
	rows, err := db.Query("select * from consent_ba")
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}
	defer rows.Close()
	var result []dto.Consent_ba

	for rows.Next() {
		var consent_ba dto.Consent_ba
		if err := rows.Scan(&consent_ba.Consent_id, &consent_ba.Consent_ba_obs_no, &consent_ba.Active_ind, &consent_ba.Business_associate_id, &consent_ba.Description, &consent_ba.Effective_date, &consent_ba.Expiry_date, &consent_ba.Full_name, &consent_ba.Ppdm_guid, &consent_ba.Remark, &consent_ba.Role, &consent_ba.Source, &consent_ba.Row_changed_by, &consent_ba.Row_changed_date, &consent_ba.Row_created_by, &consent_ba.Row_created_date, &consent_ba.Row_effective_date, &consent_ba.Row_expiry_date, &consent_ba.Row_quality); err != nil {
			return err // Exit if we get an error
		}

		// Append consent_ba to result
		result = append(result, consent_ba)
	}
	// Return result in JSON format
	return c.JSON(result)
	 
}

func SetConsentBa(c *fiber.Ctx) error {
	var consent_ba dto.Consent_ba

	setDefaults(&consent_ba)

	if err := json.Unmarshal(c.Body(), &consent_ba); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	stmt, err := db.Prepare("insert into consent_ba values(:1, :2, :3, :4, :5, :6, :7, :8, :9, :10, :11, :12, :13, :14, :15, :16, :17, :18, :19)")
	if err != nil {
		return err
	}
	consent_ba.Row_created_date = formatDate(consent_ba.Row_created_date)
	consent_ba.Row_changed_date = formatDate(consent_ba.Row_changed_date)
	consent_ba.Effective_date = formatDateString(consent_ba.Effective_date)
	consent_ba.Expiry_date = formatDateString(consent_ba.Expiry_date)
	consent_ba.Row_effective_date = formatDateString(consent_ba.Row_effective_date)
	consent_ba.Row_expiry_date = formatDateString(consent_ba.Row_expiry_date)






	rows, err := stmt.Exec(consent_ba.Consent_id, consent_ba.Consent_ba_obs_no, consent_ba.Active_ind, consent_ba.Business_associate_id, consent_ba.Description, consent_ba.Effective_date, consent_ba.Expiry_date, consent_ba.Full_name, consent_ba.Ppdm_guid, consent_ba.Remark, consent_ba.Role, consent_ba.Source, consent_ba.Row_changed_by, consent_ba.Row_changed_date, consent_ba.Row_created_by, consent_ba.Row_created_date, consent_ba.Row_effective_date, consent_ba.Row_expiry_date, consent_ba.Row_quality)
	
	if err != nil {
		return err
	}

	return c.JSON(rows)
}

func UpdateConsentBa(c *fiber.Ctx) error {
	var consent_ba dto.Consent_ba

	setDefaults(&consent_ba)

	if err := json.Unmarshal(c.Body(), &consent_ba); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	id := c.Params("id")
	consent_ba.Consent_id = id

    if consent_ba.Row_created_date != nil {
        return c.Status(400).SendString("Cannot update row_created_date")
    }

	stmt, err := db.Prepare("update consent_ba set consent_ba_obs_no = :1, active_ind = :2, business_associate_id = :3, description = :4, effective_date = :5, expiry_date = :6, full_name = :7, ppdm_guid = :8, remark = :9, role = :10, source = :11, row_changed_by = :12, row_changed_date = :13, row_created_by = :14, row_effective_date = :15, row_expiry_date = :16, row_quality = :17 where consent_id = :19")
	if err != nil {
		return err
	}

	consent_ba.Row_changed_date = formatDate(consent_ba.Row_changed_date)
	consent_ba.Effective_date = formatDateString(consent_ba.Effective_date)
	consent_ba.Expiry_date = formatDateString(consent_ba.Expiry_date)
	consent_ba.Row_effective_date = formatDateString(consent_ba.Row_effective_date)
	consent_ba.Row_expiry_date = formatDateString(consent_ba.Row_expiry_date)






	rows, err := stmt.Exec(consent_ba.Consent_ba_obs_no, consent_ba.Active_ind, consent_ba.Business_associate_id, consent_ba.Description, consent_ba.Effective_date, consent_ba.Expiry_date, consent_ba.Full_name, consent_ba.Ppdm_guid, consent_ba.Remark, consent_ba.Role, consent_ba.Source, consent_ba.Row_changed_by, consent_ba.Row_changed_date, consent_ba.Row_created_by, consent_ba.Row_effective_date, consent_ba.Row_expiry_date, consent_ba.Row_quality, consent_ba.Consent_id)
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

func PatchConsentBa(c *fiber.Ctx) error {
	id := c.Params("id")

	updatedFields := make(map[string]interface{})

	if err := c.BodyParser(&updatedFields); err != nil {
		return err
	}
	if _, ok := updatedFields["row_created_date"]; ok {
        return c.Status(400).SendString("Cannot update row_created_date field")
    }

	query := "update consent_ba set "
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

func DeleteConsentBa(c *fiber.Ctx) error {
	id := c.Params("id")
	var consent_ba dto.Consent_ba
	consent_ba.Consent_id = id

	stmt, err := db.Prepare("delete from consent_ba where consent_id = :1")
	if err != nil {
		return err
	}

	rows, err := stmt.Exec(consent_ba.Consent_id)
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


