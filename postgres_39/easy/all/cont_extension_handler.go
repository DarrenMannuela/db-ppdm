package apiv1

import (
	dto "github.com/DarrenMannuela/pt_gtn_bibliography/dto"
	"github.com/gofiber/fiber/v2"
	"strconv"
	"encoding/json"
)


func GetContExtension(c *fiber.Ctx) error {
	rows, err := db.Query("select * from cont_extension")
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}
	defer rows.Close()
	var result []dto.Cont_extension

	for rows.Next() {
		var cont_extension dto.Cont_extension
		if err := rows.Scan(&cont_extension.Contract_id, &cont_extension.Extension_id, &cont_extension.Active_ind, &cont_extension.Description, &cont_extension.Effective_date, &cont_extension.Expiry_date, &cont_extension.Extension_type, &cont_extension.Issued_by, &cont_extension.Issued_date, &cont_extension.Land_right_id, &cont_extension.Land_right_subtype, &cont_extension.Ppdm_guid, &cont_extension.Remark, &cont_extension.Source, &cont_extension.Row_changed_by, &cont_extension.Row_changed_date, &cont_extension.Row_created_by, &cont_extension.Row_created_date, &cont_extension.Row_effective_date, &cont_extension.Row_expiry_date, &cont_extension.Row_quality); err != nil {
			return err // Exit if we get an error
		}

		// Append cont_extension to result
		result = append(result, cont_extension)
	}
	// Return result in JSON format
	return c.JSON(result)
	 
}

func SetContExtension(c *fiber.Ctx) error {
	var cont_extension dto.Cont_extension

	setDefaults(&cont_extension)

	if err := json.Unmarshal(c.Body(), &cont_extension); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	stmt, err := db.Prepare("insert into cont_extension values(:1, :2, :3, :4, :5, :6, :7, :8, :9, :10, :11, :12, :13, :14, :15, :16, :17, :18, :19, :20, :21)")
	if err != nil {
		return err
	}
	cont_extension.Row_created_date = formatDate(cont_extension.Row_created_date)
	cont_extension.Row_changed_date = formatDate(cont_extension.Row_changed_date)
	cont_extension.Effective_date = formatDateString(cont_extension.Effective_date)
	cont_extension.Expiry_date = formatDateString(cont_extension.Expiry_date)
	cont_extension.Issued_date = formatDateString(cont_extension.Issued_date)
	cont_extension.Row_effective_date = formatDateString(cont_extension.Row_effective_date)
	cont_extension.Row_expiry_date = formatDateString(cont_extension.Row_expiry_date)






	rows, err := stmt.Exec(cont_extension.Contract_id, cont_extension.Extension_id, cont_extension.Active_ind, cont_extension.Description, cont_extension.Effective_date, cont_extension.Expiry_date, cont_extension.Extension_type, cont_extension.Issued_by, cont_extension.Issued_date, cont_extension.Land_right_id, cont_extension.Land_right_subtype, cont_extension.Ppdm_guid, cont_extension.Remark, cont_extension.Source, cont_extension.Row_changed_by, cont_extension.Row_changed_date, cont_extension.Row_created_by, cont_extension.Row_created_date, cont_extension.Row_effective_date, cont_extension.Row_expiry_date, cont_extension.Row_quality)
	
	if err != nil {
		return err
	}

	return c.JSON(rows)
}

func UpdateContExtension(c *fiber.Ctx) error {
	var cont_extension dto.Cont_extension

	setDefaults(&cont_extension)

	if err := json.Unmarshal(c.Body(), &cont_extension); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	id := c.Params("id")
	cont_extension.Contract_id = id

    if cont_extension.Row_created_date != nil {
        return c.Status(400).SendString("Cannot update row_created_date")
    }

	stmt, err := db.Prepare("update cont_extension set extension_id = :1, active_ind = :2, description = :3, effective_date = :4, expiry_date = :5, extension_type = :6, issued_by = :7, issued_date = :8, land_right_id = :9, land_right_subtype = :10, ppdm_guid = :11, remark = :12, source = :13, row_changed_by = :14, row_changed_date = :15, row_created_by = :16, row_effective_date = :17, row_expiry_date = :18, row_quality = :19 where contract_id = :21")
	if err != nil {
		return err
	}

	cont_extension.Row_changed_date = formatDate(cont_extension.Row_changed_date)
	cont_extension.Effective_date = formatDateString(cont_extension.Effective_date)
	cont_extension.Expiry_date = formatDateString(cont_extension.Expiry_date)
	cont_extension.Issued_date = formatDateString(cont_extension.Issued_date)
	cont_extension.Row_effective_date = formatDateString(cont_extension.Row_effective_date)
	cont_extension.Row_expiry_date = formatDateString(cont_extension.Row_expiry_date)






	rows, err := stmt.Exec(cont_extension.Extension_id, cont_extension.Active_ind, cont_extension.Description, cont_extension.Effective_date, cont_extension.Expiry_date, cont_extension.Extension_type, cont_extension.Issued_by, cont_extension.Issued_date, cont_extension.Land_right_id, cont_extension.Land_right_subtype, cont_extension.Ppdm_guid, cont_extension.Remark, cont_extension.Source, cont_extension.Row_changed_by, cont_extension.Row_changed_date, cont_extension.Row_created_by, cont_extension.Row_effective_date, cont_extension.Row_expiry_date, cont_extension.Row_quality, cont_extension.Contract_id)
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

func PatchContExtension(c *fiber.Ctx) error {
	id := c.Params("id")

	updatedFields := make(map[string]interface{})

	if err := c.BodyParser(&updatedFields); err != nil {
		return err
	}
	if _, ok := updatedFields["row_created_date"]; ok {
        return c.Status(400).SendString("Cannot update row_created_date field")
    }

	query := "update cont_extension set "
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
		} else if key == "effective_date" || key == "expiry_date" || key == "issued_date" || key == "row_effective_date" || key == "row_expiry_date"      {
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

func DeleteContExtension(c *fiber.Ctx) error {
	id := c.Params("id")
	var cont_extension dto.Cont_extension
	cont_extension.Contract_id = id

	stmt, err := db.Prepare("delete from cont_extension where contract_id = :1")
	if err != nil {
		return err
	}

	rows, err := stmt.Exec(cont_extension.Contract_id)
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


