package apiv1

import (
	dto "github.com/DarrenMannuela/pt_gtn_bibliography/dto"
	"github.com/gofiber/fiber/v2"
	"strconv"
	"encoding/json"
)


func GetRmEncoding(c *fiber.Ctx) error {
	rows, err := db.Query("select * from rm_encoding")
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}
	defer rows.Close()
	var result []dto.Rm_encoding

	for rows.Next() {
		var rm_encoding dto.Rm_encoding
		if err := rows.Scan(&rm_encoding.Physical_item_id, &rm_encoding.Encoding_seq_no, &rm_encoding.Active_ind, &rm_encoding.Decrypt_key_id, &rm_encoding.Effective_date, &rm_encoding.Encoding_type, &rm_encoding.Encoding_version, &rm_encoding.Expiry_date, &rm_encoding.Ppdm_guid, &rm_encoding.Remark, &rm_encoding.Source, &rm_encoding.Sw_application_id, &rm_encoding.Row_changed_by, &rm_encoding.Row_changed_date, &rm_encoding.Row_created_by, &rm_encoding.Row_created_date, &rm_encoding.Row_effective_date, &rm_encoding.Row_expiry_date, &rm_encoding.Row_quality); err != nil {
			return err // Exit if we get an error
		}

		// Append rm_encoding to result
		result = append(result, rm_encoding)
	}
	// Return result in JSON format
	return c.JSON(result)
	 
}

func SetRmEncoding(c *fiber.Ctx) error {
	var rm_encoding dto.Rm_encoding

	setDefaults(&rm_encoding)

	if err := json.Unmarshal(c.Body(), &rm_encoding); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	stmt, err := db.Prepare("insert into rm_encoding values(:1, :2, :3, :4, :5, :6, :7, :8, :9, :10, :11, :12, :13, :14, :15, :16, :17, :18, :19)")
	if err != nil {
		return err
	}
	rm_encoding.Row_created_date = formatDate(rm_encoding.Row_created_date)
	rm_encoding.Row_changed_date = formatDate(rm_encoding.Row_changed_date)
	rm_encoding.Effective_date = formatDateString(rm_encoding.Effective_date)
	rm_encoding.Expiry_date = formatDateString(rm_encoding.Expiry_date)
	rm_encoding.Row_effective_date = formatDateString(rm_encoding.Row_effective_date)
	rm_encoding.Row_expiry_date = formatDateString(rm_encoding.Row_expiry_date)






	rows, err := stmt.Exec(rm_encoding.Physical_item_id, rm_encoding.Encoding_seq_no, rm_encoding.Active_ind, rm_encoding.Decrypt_key_id, rm_encoding.Effective_date, rm_encoding.Encoding_type, rm_encoding.Encoding_version, rm_encoding.Expiry_date, rm_encoding.Ppdm_guid, rm_encoding.Remark, rm_encoding.Source, rm_encoding.Sw_application_id, rm_encoding.Row_changed_by, rm_encoding.Row_changed_date, rm_encoding.Row_created_by, rm_encoding.Row_created_date, rm_encoding.Row_effective_date, rm_encoding.Row_expiry_date, rm_encoding.Row_quality)
	
	if err != nil {
		return err
	}

	return c.JSON(rows)
}

func UpdateRmEncoding(c *fiber.Ctx) error {
	var rm_encoding dto.Rm_encoding

	setDefaults(&rm_encoding)

	if err := json.Unmarshal(c.Body(), &rm_encoding); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	id := c.Params("id")
	rm_encoding.Physical_item_id = id

    if rm_encoding.Row_created_date != nil {
        return c.Status(400).SendString("Cannot update row_created_date")
    }

	stmt, err := db.Prepare("update rm_encoding set encoding_seq_no = :1, active_ind = :2, decrypt_key_id = :3, effective_date = :4, encoding_type = :5, encoding_version = :6, expiry_date = :7, ppdm_guid = :8, remark = :9, source = :10, sw_application_id = :11, row_changed_by = :12, row_changed_date = :13, row_created_by = :14, row_effective_date = :15, row_expiry_date = :16, row_quality = :17 where physical_item_id = :19")
	if err != nil {
		return err
	}

	rm_encoding.Row_changed_date = formatDate(rm_encoding.Row_changed_date)
	rm_encoding.Effective_date = formatDateString(rm_encoding.Effective_date)
	rm_encoding.Expiry_date = formatDateString(rm_encoding.Expiry_date)
	rm_encoding.Row_effective_date = formatDateString(rm_encoding.Row_effective_date)
	rm_encoding.Row_expiry_date = formatDateString(rm_encoding.Row_expiry_date)






	rows, err := stmt.Exec(rm_encoding.Encoding_seq_no, rm_encoding.Active_ind, rm_encoding.Decrypt_key_id, rm_encoding.Effective_date, rm_encoding.Encoding_type, rm_encoding.Encoding_version, rm_encoding.Expiry_date, rm_encoding.Ppdm_guid, rm_encoding.Remark, rm_encoding.Source, rm_encoding.Sw_application_id, rm_encoding.Row_changed_by, rm_encoding.Row_changed_date, rm_encoding.Row_created_by, rm_encoding.Row_effective_date, rm_encoding.Row_expiry_date, rm_encoding.Row_quality, rm_encoding.Physical_item_id)
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

func PatchRmEncoding(c *fiber.Ctx) error {
	id := c.Params("id")

	updatedFields := make(map[string]interface{})

	if err := c.BodyParser(&updatedFields); err != nil {
		return err
	}
	if _, ok := updatedFields["row_created_date"]; ok {
        return c.Status(400).SendString("Cannot update row_created_date field")
    }

	query := "update rm_encoding set "
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
	query += " where physical_item_id = :id"

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

func DeleteRmEncoding(c *fiber.Ctx) error {
	id := c.Params("id")
	var rm_encoding dto.Rm_encoding
	rm_encoding.Physical_item_id = id

	stmt, err := db.Prepare("delete from rm_encoding where physical_item_id = :1")
	if err != nil {
		return err
	}

	rows, err := stmt.Exec(rm_encoding.Physical_item_id)
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


