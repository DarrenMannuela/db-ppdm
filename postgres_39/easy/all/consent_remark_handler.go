package apiv1

import (
	dto "github.com/DarrenMannuela/pt_gtn_bibliography/dto"
	"github.com/gofiber/fiber/v2"
	"strconv"
	"encoding/json"
)


func GetConsentRemark(c *fiber.Ctx) error {
	rows, err := db.Query("select * from consent_remark")
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}
	defer rows.Close()
	var result []dto.Consent_remark

	for rows.Next() {
		var consent_remark dto.Consent_remark
		if err := rows.Scan(&consent_remark.Consent_id, &consent_remark.Remark_id, &consent_remark.Active_ind, &consent_remark.Effective_date, &consent_remark.Expiry_date, &consent_remark.Made_about_ba_id, &consent_remark.Made_by_ba_id, &consent_remark.Made_for_ba_id, &consent_remark.Ppdm_guid, &consent_remark.Remark, &consent_remark.Remark_type, &consent_remark.Source, &consent_remark.Row_changed_by, &consent_remark.Row_changed_date, &consent_remark.Row_created_by, &consent_remark.Row_created_date, &consent_remark.Row_effective_date, &consent_remark.Row_expiry_date, &consent_remark.Row_quality); err != nil {
			return err // Exit if we get an error
		}

		// Append consent_remark to result
		result = append(result, consent_remark)
	}
	// Return result in JSON format
	return c.JSON(result)
	 
}

func SetConsentRemark(c *fiber.Ctx) error {
	var consent_remark dto.Consent_remark

	setDefaults(&consent_remark)

	if err := json.Unmarshal(c.Body(), &consent_remark); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	stmt, err := db.Prepare("insert into consent_remark values(:1, :2, :3, :4, :5, :6, :7, :8, :9, :10, :11, :12, :13, :14, :15, :16, :17, :18, :19)")
	if err != nil {
		return err
	}
	consent_remark.Row_created_date = formatDate(consent_remark.Row_created_date)
	consent_remark.Row_changed_date = formatDate(consent_remark.Row_changed_date)
	consent_remark.Effective_date = formatDateString(consent_remark.Effective_date)
	consent_remark.Expiry_date = formatDateString(consent_remark.Expiry_date)
	consent_remark.Row_effective_date = formatDateString(consent_remark.Row_effective_date)
	consent_remark.Row_expiry_date = formatDateString(consent_remark.Row_expiry_date)






	rows, err := stmt.Exec(consent_remark.Consent_id, consent_remark.Remark_id, consent_remark.Active_ind, consent_remark.Effective_date, consent_remark.Expiry_date, consent_remark.Made_about_ba_id, consent_remark.Made_by_ba_id, consent_remark.Made_for_ba_id, consent_remark.Ppdm_guid, consent_remark.Remark, consent_remark.Remark_type, consent_remark.Source, consent_remark.Row_changed_by, consent_remark.Row_changed_date, consent_remark.Row_created_by, consent_remark.Row_created_date, consent_remark.Row_effective_date, consent_remark.Row_expiry_date, consent_remark.Row_quality)
	
	if err != nil {
		return err
	}

	return c.JSON(rows)
}

func UpdateConsentRemark(c *fiber.Ctx) error {
	var consent_remark dto.Consent_remark

	setDefaults(&consent_remark)

	if err := json.Unmarshal(c.Body(), &consent_remark); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	id := c.Params("id")
	consent_remark.Consent_id = id

    if consent_remark.Row_created_date != nil {
        return c.Status(400).SendString("Cannot update row_created_date")
    }

	stmt, err := db.Prepare("update consent_remark set remark_id = :1, active_ind = :2, effective_date = :3, expiry_date = :4, made_about_ba_id = :5, made_by_ba_id = :6, made_for_ba_id = :7, ppdm_guid = :8, remark = :9, remark_type = :10, source = :11, row_changed_by = :12, row_changed_date = :13, row_created_by = :14, row_effective_date = :15, row_expiry_date = :16, row_quality = :17 where consent_id = :19")
	if err != nil {
		return err
	}

	consent_remark.Row_changed_date = formatDate(consent_remark.Row_changed_date)
	consent_remark.Effective_date = formatDateString(consent_remark.Effective_date)
	consent_remark.Expiry_date = formatDateString(consent_remark.Expiry_date)
	consent_remark.Row_effective_date = formatDateString(consent_remark.Row_effective_date)
	consent_remark.Row_expiry_date = formatDateString(consent_remark.Row_expiry_date)






	rows, err := stmt.Exec(consent_remark.Remark_id, consent_remark.Active_ind, consent_remark.Effective_date, consent_remark.Expiry_date, consent_remark.Made_about_ba_id, consent_remark.Made_by_ba_id, consent_remark.Made_for_ba_id, consent_remark.Ppdm_guid, consent_remark.Remark, consent_remark.Remark_type, consent_remark.Source, consent_remark.Row_changed_by, consent_remark.Row_changed_date, consent_remark.Row_created_by, consent_remark.Row_effective_date, consent_remark.Row_expiry_date, consent_remark.Row_quality, consent_remark.Consent_id)
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

func PatchConsentRemark(c *fiber.Ctx) error {
	id := c.Params("id")

	updatedFields := make(map[string]interface{})

	if err := c.BodyParser(&updatedFields); err != nil {
		return err
	}
	if _, ok := updatedFields["row_created_date"]; ok {
        return c.Status(400).SendString("Cannot update row_created_date field")
    }

	query := "update consent_remark set "
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

func DeleteConsentRemark(c *fiber.Ctx) error {
	id := c.Params("id")
	var consent_remark dto.Consent_remark
	consent_remark.Consent_id = id

	stmt, err := db.Prepare("delete from consent_remark where consent_id = :1")
	if err != nil {
		return err
	}

	rows, err := stmt.Exec(consent_remark.Consent_id)
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


