package apiv1

import (
	dto "github.com/DarrenMannuela/pt_gtn_bibliography/dto"
	"github.com/gofiber/fiber/v2"
	"strconv"
	"encoding/json"
)


func GetPpdmQuantity(c *fiber.Ctx) error {
	rows, err := db.Query("select * from ppdm_quantity")
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}
	defer rows.Close()
	var result []dto.Ppdm_quantity

	for rows.Next() {
		var ppdm_quantity dto.Ppdm_quantity
		if err := rows.Scan(&ppdm_quantity.Quantity_type, &ppdm_quantity.Active_ind, &ppdm_quantity.Effective_date, &ppdm_quantity.Expiry_date, &ppdm_quantity.Long_name, &ppdm_quantity.Ppdm_guid, &ppdm_quantity.Remark, &ppdm_quantity.Short_name, &ppdm_quantity.Source, &ppdm_quantity.Source_document_id, &ppdm_quantity.Row_changed_by, &ppdm_quantity.Row_changed_date, &ppdm_quantity.Row_created_by, &ppdm_quantity.Row_created_date, &ppdm_quantity.Row_effective_date, &ppdm_quantity.Row_expiry_date, &ppdm_quantity.Row_quality); err != nil {
			return err // Exit if we get an error
		}

		// Append ppdm_quantity to result
		result = append(result, ppdm_quantity)
	}
	// Return result in JSON format
	return c.JSON(result)
	 
}

func SetPpdmQuantity(c *fiber.Ctx) error {
	var ppdm_quantity dto.Ppdm_quantity

	setDefaults(&ppdm_quantity)

	if err := json.Unmarshal(c.Body(), &ppdm_quantity); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	stmt, err := db.Prepare("insert into ppdm_quantity values(:1, :2, :3, :4, :5, :6, :7, :8, :9, :10, :11, :12, :13, :14, :15, :16, :17)")
	if err != nil {
		return err
	}
	ppdm_quantity.Row_created_date = formatDate(ppdm_quantity.Row_created_date)
	ppdm_quantity.Row_changed_date = formatDate(ppdm_quantity.Row_changed_date)
	ppdm_quantity.Effective_date = formatDateString(ppdm_quantity.Effective_date)
	ppdm_quantity.Expiry_date = formatDateString(ppdm_quantity.Expiry_date)
	ppdm_quantity.Row_effective_date = formatDateString(ppdm_quantity.Row_effective_date)
	ppdm_quantity.Row_expiry_date = formatDateString(ppdm_quantity.Row_expiry_date)






	rows, err := stmt.Exec(ppdm_quantity.Quantity_type, ppdm_quantity.Active_ind, ppdm_quantity.Effective_date, ppdm_quantity.Expiry_date, ppdm_quantity.Long_name, ppdm_quantity.Ppdm_guid, ppdm_quantity.Remark, ppdm_quantity.Short_name, ppdm_quantity.Source, ppdm_quantity.Source_document_id, ppdm_quantity.Row_changed_by, ppdm_quantity.Row_changed_date, ppdm_quantity.Row_created_by, ppdm_quantity.Row_created_date, ppdm_quantity.Row_effective_date, ppdm_quantity.Row_expiry_date, ppdm_quantity.Row_quality)
	
	if err != nil {
		return err
	}

	return c.JSON(rows)
}

func UpdatePpdmQuantity(c *fiber.Ctx) error {
	var ppdm_quantity dto.Ppdm_quantity

	setDefaults(&ppdm_quantity)

	if err := json.Unmarshal(c.Body(), &ppdm_quantity); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	id := c.Params("id")
	ppdm_quantity.Quantity_type = id

    if ppdm_quantity.Row_created_date != nil {
        return c.Status(400).SendString("Cannot update row_created_date")
    }

	stmt, err := db.Prepare("update ppdm_quantity set active_ind = :1, effective_date = :2, expiry_date = :3, long_name = :4, ppdm_guid = :5, remark = :6, short_name = :7, source = :8, source_document_id = :9, row_changed_by = :10, row_changed_date = :11, row_created_by = :12, row_effective_date = :13, row_expiry_date = :14, row_quality = :15 where quantity_type = :17")
	if err != nil {
		return err
	}

	ppdm_quantity.Row_changed_date = formatDate(ppdm_quantity.Row_changed_date)
	ppdm_quantity.Effective_date = formatDateString(ppdm_quantity.Effective_date)
	ppdm_quantity.Expiry_date = formatDateString(ppdm_quantity.Expiry_date)
	ppdm_quantity.Row_effective_date = formatDateString(ppdm_quantity.Row_effective_date)
	ppdm_quantity.Row_expiry_date = formatDateString(ppdm_quantity.Row_expiry_date)






	rows, err := stmt.Exec(ppdm_quantity.Active_ind, ppdm_quantity.Effective_date, ppdm_quantity.Expiry_date, ppdm_quantity.Long_name, ppdm_quantity.Ppdm_guid, ppdm_quantity.Remark, ppdm_quantity.Short_name, ppdm_quantity.Source, ppdm_quantity.Source_document_id, ppdm_quantity.Row_changed_by, ppdm_quantity.Row_changed_date, ppdm_quantity.Row_created_by, ppdm_quantity.Row_effective_date, ppdm_quantity.Row_expiry_date, ppdm_quantity.Row_quality, ppdm_quantity.Quantity_type)
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

func PatchPpdmQuantity(c *fiber.Ctx) error {
	id := c.Params("id")

	updatedFields := make(map[string]interface{})

	if err := c.BodyParser(&updatedFields); err != nil {
		return err
	}
	if _, ok := updatedFields["row_created_date"]; ok {
        return c.Status(400).SendString("Cannot update row_created_date field")
    }

	query := "update ppdm_quantity set "
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
	query += " where quantity_type = :id"

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

func DeletePpdmQuantity(c *fiber.Ctx) error {
	id := c.Params("id")
	var ppdm_quantity dto.Ppdm_quantity
	ppdm_quantity.Quantity_type = id

	stmt, err := db.Prepare("delete from ppdm_quantity where quantity_type = :1")
	if err != nil {
		return err
	}

	rows, err := stmt.Exec(ppdm_quantity.Quantity_type)
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


