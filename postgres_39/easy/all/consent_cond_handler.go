package apiv1

import (
	dto "github.com/DarrenMannuela/pt_gtn_bibliography/dto"
	"github.com/gofiber/fiber/v2"
	"strconv"
	"encoding/json"
)


func GetConsentCond(c *fiber.Ctx) error {
	rows, err := db.Query("select * from consent_cond")
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}
	defer rows.Close()
	var result []dto.Consent_cond

	for rows.Next() {
		var consent_cond dto.Consent_cond
		if err := rows.Scan(&consent_cond.Consent_id, &consent_cond.Condition_id, &consent_cond.Active_ind, &consent_cond.Component_id, &consent_cond.Consent_condition, &consent_cond.Consent_type, &consent_cond.Effective_date, &consent_cond.Expiry_date, &consent_cond.Ppdm_guid, &consent_cond.Remark, &consent_cond.Source, &consent_cond.Row_changed_by, &consent_cond.Row_changed_date, &consent_cond.Row_created_by, &consent_cond.Row_created_date, &consent_cond.Row_effective_date, &consent_cond.Row_expiry_date, &consent_cond.Row_quality); err != nil {
			return err // Exit if we get an error
		}

		// Append consent_cond to result
		result = append(result, consent_cond)
	}
	// Return result in JSON format
	return c.JSON(result)
	 
}

func SetConsentCond(c *fiber.Ctx) error {
	var consent_cond dto.Consent_cond

	setDefaults(&consent_cond)

	if err := json.Unmarshal(c.Body(), &consent_cond); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	stmt, err := db.Prepare("insert into consent_cond values(:1, :2, :3, :4, :5, :6, :7, :8, :9, :10, :11, :12, :13, :14, :15, :16, :17, :18)")
	if err != nil {
		return err
	}
	consent_cond.Row_created_date = formatDate(consent_cond.Row_created_date)
	consent_cond.Row_changed_date = formatDate(consent_cond.Row_changed_date)
	consent_cond.Effective_date = formatDateString(consent_cond.Effective_date)
	consent_cond.Expiry_date = formatDateString(consent_cond.Expiry_date)
	consent_cond.Row_effective_date = formatDateString(consent_cond.Row_effective_date)
	consent_cond.Row_expiry_date = formatDateString(consent_cond.Row_expiry_date)






	rows, err := stmt.Exec(consent_cond.Consent_id, consent_cond.Condition_id, consent_cond.Active_ind, consent_cond.Component_id, consent_cond.Consent_condition, consent_cond.Consent_type, consent_cond.Effective_date, consent_cond.Expiry_date, consent_cond.Ppdm_guid, consent_cond.Remark, consent_cond.Source, consent_cond.Row_changed_by, consent_cond.Row_changed_date, consent_cond.Row_created_by, consent_cond.Row_created_date, consent_cond.Row_effective_date, consent_cond.Row_expiry_date, consent_cond.Row_quality)
	
	if err != nil {
		return err
	}

	return c.JSON(rows)
}

func UpdateConsentCond(c *fiber.Ctx) error {
	var consent_cond dto.Consent_cond

	setDefaults(&consent_cond)

	if err := json.Unmarshal(c.Body(), &consent_cond); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	id := c.Params("id")
	consent_cond.Consent_id = id

    if consent_cond.Row_created_date != nil {
        return c.Status(400).SendString("Cannot update row_created_date")
    }

	stmt, err := db.Prepare("update consent_cond set condition_id = :1, active_ind = :2, component_id = :3, consent_condition = :4, consent_type = :5, effective_date = :6, expiry_date = :7, ppdm_guid = :8, remark = :9, source = :10, row_changed_by = :11, row_changed_date = :12, row_created_by = :13, row_effective_date = :14, row_expiry_date = :15, row_quality = :16 where consent_id = :18")
	if err != nil {
		return err
	}

	consent_cond.Row_changed_date = formatDate(consent_cond.Row_changed_date)
	consent_cond.Effective_date = formatDateString(consent_cond.Effective_date)
	consent_cond.Expiry_date = formatDateString(consent_cond.Expiry_date)
	consent_cond.Row_effective_date = formatDateString(consent_cond.Row_effective_date)
	consent_cond.Row_expiry_date = formatDateString(consent_cond.Row_expiry_date)






	rows, err := stmt.Exec(consent_cond.Condition_id, consent_cond.Active_ind, consent_cond.Component_id, consent_cond.Consent_condition, consent_cond.Consent_type, consent_cond.Effective_date, consent_cond.Expiry_date, consent_cond.Ppdm_guid, consent_cond.Remark, consent_cond.Source, consent_cond.Row_changed_by, consent_cond.Row_changed_date, consent_cond.Row_created_by, consent_cond.Row_effective_date, consent_cond.Row_expiry_date, consent_cond.Row_quality, consent_cond.Consent_id)
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

func PatchConsentCond(c *fiber.Ctx) error {
	id := c.Params("id")

	updatedFields := make(map[string]interface{})

	if err := c.BodyParser(&updatedFields); err != nil {
		return err
	}
	if _, ok := updatedFields["row_created_date"]; ok {
        return c.Status(400).SendString("Cannot update row_created_date field")
    }

	query := "update consent_cond set "
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

func DeleteConsentCond(c *fiber.Ctx) error {
	id := c.Params("id")
	var consent_cond dto.Consent_cond
	consent_cond.Consent_id = id

	stmt, err := db.Prepare("delete from consent_cond where consent_id = :1")
	if err != nil {
		return err
	}

	rows, err := stmt.Exec(consent_cond.Consent_id)
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


