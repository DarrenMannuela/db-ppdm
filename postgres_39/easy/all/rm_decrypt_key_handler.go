package apiv1

import (
	dto "github.com/DarrenMannuela/pt_gtn_bibliography/dto"
	"github.com/gofiber/fiber/v2"
	"strconv"
	"encoding/json"
)


func GetRmDecryptKey(c *fiber.Ctx) error {
	rows, err := db.Query("select * from rm_decrypt_key")
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}
	defer rows.Close()
	var result []dto.Rm_decrypt_key

	for rows.Next() {
		var rm_decrypt_key dto.Rm_decrypt_key
		if err := rows.Scan(&rm_decrypt_key.Decrypt_key_id, &rm_decrypt_key.Active_ind, &rm_decrypt_key.Decryption_type, &rm_decrypt_key.Decrypt_key, &rm_decrypt_key.Effective_date, &rm_decrypt_key.Expiry_date, &rm_decrypt_key.Ppdm_guid, &rm_decrypt_key.Remark, &rm_decrypt_key.Source, &rm_decrypt_key.Sw_application_id, &rm_decrypt_key.Row_changed_by, &rm_decrypt_key.Row_changed_date, &rm_decrypt_key.Row_created_by, &rm_decrypt_key.Row_created_date, &rm_decrypt_key.Row_effective_date, &rm_decrypt_key.Row_expiry_date, &rm_decrypt_key.Row_quality); err != nil {
			return err // Exit if we get an error
		}

		// Append rm_decrypt_key to result
		result = append(result, rm_decrypt_key)
	}
	// Return result in JSON format
	return c.JSON(result)
	 
}

func SetRmDecryptKey(c *fiber.Ctx) error {
	var rm_decrypt_key dto.Rm_decrypt_key

	setDefaults(&rm_decrypt_key)

	if err := json.Unmarshal(c.Body(), &rm_decrypt_key); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	stmt, err := db.Prepare("insert into rm_decrypt_key values(:1, :2, :3, :4, :5, :6, :7, :8, :9, :10, :11, :12, :13, :14, :15, :16, :17)")
	if err != nil {
		return err
	}
	rm_decrypt_key.Row_created_date = formatDate(rm_decrypt_key.Row_created_date)
	rm_decrypt_key.Row_changed_date = formatDate(rm_decrypt_key.Row_changed_date)
	rm_decrypt_key.Effective_date = formatDateString(rm_decrypt_key.Effective_date)
	rm_decrypt_key.Expiry_date = formatDateString(rm_decrypt_key.Expiry_date)
	rm_decrypt_key.Row_effective_date = formatDateString(rm_decrypt_key.Row_effective_date)
	rm_decrypt_key.Row_expiry_date = formatDateString(rm_decrypt_key.Row_expiry_date)






	rows, err := stmt.Exec(rm_decrypt_key.Decrypt_key_id, rm_decrypt_key.Active_ind, rm_decrypt_key.Decryption_type, rm_decrypt_key.Decrypt_key, rm_decrypt_key.Effective_date, rm_decrypt_key.Expiry_date, rm_decrypt_key.Ppdm_guid, rm_decrypt_key.Remark, rm_decrypt_key.Source, rm_decrypt_key.Sw_application_id, rm_decrypt_key.Row_changed_by, rm_decrypt_key.Row_changed_date, rm_decrypt_key.Row_created_by, rm_decrypt_key.Row_created_date, rm_decrypt_key.Row_effective_date, rm_decrypt_key.Row_expiry_date, rm_decrypt_key.Row_quality)
	
	if err != nil {
		return err
	}

	return c.JSON(rows)
}

func UpdateRmDecryptKey(c *fiber.Ctx) error {
	var rm_decrypt_key dto.Rm_decrypt_key

	setDefaults(&rm_decrypt_key)

	if err := json.Unmarshal(c.Body(), &rm_decrypt_key); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	id := c.Params("id")
	rm_decrypt_key.Decrypt_key_id = id

    if rm_decrypt_key.Row_created_date != nil {
        return c.Status(400).SendString("Cannot update row_created_date")
    }

	stmt, err := db.Prepare("update rm_decrypt_key set active_ind = :1, decryption_type = :2, decrypt_key = :3, effective_date = :4, expiry_date = :5, ppdm_guid = :6, remark = :7, source = :8, sw_application_id = :9, row_changed_by = :10, row_changed_date = :11, row_created_by = :12, row_effective_date = :13, row_expiry_date = :14, row_quality = :15 where decrypt_key_id = :17")
	if err != nil {
		return err
	}

	rm_decrypt_key.Row_changed_date = formatDate(rm_decrypt_key.Row_changed_date)
	rm_decrypt_key.Effective_date = formatDateString(rm_decrypt_key.Effective_date)
	rm_decrypt_key.Expiry_date = formatDateString(rm_decrypt_key.Expiry_date)
	rm_decrypt_key.Row_effective_date = formatDateString(rm_decrypt_key.Row_effective_date)
	rm_decrypt_key.Row_expiry_date = formatDateString(rm_decrypt_key.Row_expiry_date)






	rows, err := stmt.Exec(rm_decrypt_key.Active_ind, rm_decrypt_key.Decryption_type, rm_decrypt_key.Decrypt_key, rm_decrypt_key.Effective_date, rm_decrypt_key.Expiry_date, rm_decrypt_key.Ppdm_guid, rm_decrypt_key.Remark, rm_decrypt_key.Source, rm_decrypt_key.Sw_application_id, rm_decrypt_key.Row_changed_by, rm_decrypt_key.Row_changed_date, rm_decrypt_key.Row_created_by, rm_decrypt_key.Row_effective_date, rm_decrypt_key.Row_expiry_date, rm_decrypt_key.Row_quality, rm_decrypt_key.Decrypt_key_id)
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

func PatchRmDecryptKey(c *fiber.Ctx) error {
	id := c.Params("id")

	updatedFields := make(map[string]interface{})

	if err := c.BodyParser(&updatedFields); err != nil {
		return err
	}
	if _, ok := updatedFields["row_created_date"]; ok {
        return c.Status(400).SendString("Cannot update row_created_date field")
    }

	query := "update rm_decrypt_key set "
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
	query += " where decrypt_key_id = :id"

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

func DeleteRmDecryptKey(c *fiber.Ctx) error {
	id := c.Params("id")
	var rm_decrypt_key dto.Rm_decrypt_key
	rm_decrypt_key.Decrypt_key_id = id

	stmt, err := db.Prepare("delete from rm_decrypt_key where decrypt_key_id = :1")
	if err != nil {
		return err
	}

	rows, err := stmt.Exec(rm_decrypt_key.Decrypt_key_id)
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


