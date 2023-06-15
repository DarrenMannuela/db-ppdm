package apiv1

import (
	dto "github.com/DarrenMannuela/pt_gtn_bibliography/dto"
	"github.com/gofiber/fiber/v2"
	"strconv"
	"encoding/json"
)


func GetContType(c *fiber.Ctx) error {
	rows, err := db.Query("select * from cont_type")
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}
	defer rows.Close()
	var result []dto.Cont_type

	for rows.Next() {
		var cont_type dto.Cont_type
		if err := rows.Scan(&cont_type.Contract_id, &cont_type.Contract_type, &cont_type.Active_ind, &cont_type.Effective_date, &cont_type.Expiry_date, &cont_type.Ppdm_guid, &cont_type.Remark, &cont_type.Source, &cont_type.Row_changed_by, &cont_type.Row_changed_date, &cont_type.Row_created_by, &cont_type.Row_created_date, &cont_type.Row_effective_date, &cont_type.Row_expiry_date, &cont_type.Row_quality); err != nil {
			return err // Exit if we get an error
		}

		// Append cont_type to result
		result = append(result, cont_type)
	}
	// Return result in JSON format
	return c.JSON(result)
	 
}

func SetContType(c *fiber.Ctx) error {
	var cont_type dto.Cont_type

	setDefaults(&cont_type)

	if err := json.Unmarshal(c.Body(), &cont_type); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	stmt, err := db.Prepare("insert into cont_type values(:1, :2, :3, :4, :5, :6, :7, :8, :9, :10, :11, :12, :13, :14, :15)")
	if err != nil {
		return err
	}
	cont_type.Row_created_date = formatDate(cont_type.Row_created_date)
	cont_type.Row_changed_date = formatDate(cont_type.Row_changed_date)
	cont_type.Effective_date = formatDateString(cont_type.Effective_date)
	cont_type.Expiry_date = formatDateString(cont_type.Expiry_date)
	cont_type.Row_effective_date = formatDateString(cont_type.Row_effective_date)
	cont_type.Row_expiry_date = formatDateString(cont_type.Row_expiry_date)






	rows, err := stmt.Exec(cont_type.Contract_id, cont_type.Contract_type, cont_type.Active_ind, cont_type.Effective_date, cont_type.Expiry_date, cont_type.Ppdm_guid, cont_type.Remark, cont_type.Source, cont_type.Row_changed_by, cont_type.Row_changed_date, cont_type.Row_created_by, cont_type.Row_created_date, cont_type.Row_effective_date, cont_type.Row_expiry_date, cont_type.Row_quality)
	
	if err != nil {
		return err
	}

	return c.JSON(rows)
}

func UpdateContType(c *fiber.Ctx) error {
	var cont_type dto.Cont_type

	setDefaults(&cont_type)

	if err := json.Unmarshal(c.Body(), &cont_type); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	id := c.Params("id")
	cont_type.Contract_id = id

    if cont_type.Row_created_date != nil {
        return c.Status(400).SendString("Cannot update row_created_date")
    }

	stmt, err := db.Prepare("update cont_type set contract_type = :1, active_ind = :2, effective_date = :3, expiry_date = :4, ppdm_guid = :5, remark = :6, source = :7, row_changed_by = :8, row_changed_date = :9, row_created_by = :10, row_effective_date = :11, row_expiry_date = :12, row_quality = :13 where contract_id = :15")
	if err != nil {
		return err
	}

	cont_type.Row_changed_date = formatDate(cont_type.Row_changed_date)
	cont_type.Effective_date = formatDateString(cont_type.Effective_date)
	cont_type.Expiry_date = formatDateString(cont_type.Expiry_date)
	cont_type.Row_effective_date = formatDateString(cont_type.Row_effective_date)
	cont_type.Row_expiry_date = formatDateString(cont_type.Row_expiry_date)






	rows, err := stmt.Exec(cont_type.Contract_type, cont_type.Active_ind, cont_type.Effective_date, cont_type.Expiry_date, cont_type.Ppdm_guid, cont_type.Remark, cont_type.Source, cont_type.Row_changed_by, cont_type.Row_changed_date, cont_type.Row_created_by, cont_type.Row_effective_date, cont_type.Row_expiry_date, cont_type.Row_quality, cont_type.Contract_id)
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

func PatchContType(c *fiber.Ctx) error {
	id := c.Params("id")

	updatedFields := make(map[string]interface{})

	if err := c.BodyParser(&updatedFields); err != nil {
		return err
	}
	if _, ok := updatedFields["row_created_date"]; ok {
        return c.Status(400).SendString("Cannot update row_created_date field")
    }

	query := "update cont_type set "
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

func DeleteContType(c *fiber.Ctx) error {
	id := c.Params("id")
	var cont_type dto.Cont_type
	cont_type.Contract_id = id

	stmt, err := db.Prepare("delete from cont_type where contract_id = :1")
	if err != nil {
		return err
	}

	rows, err := stmt.Exec(cont_type.Contract_id)
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


