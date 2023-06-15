package apiv1

import (
	dto "github.com/DarrenMannuela/pt_gtn_bibliography/dto"
	"github.com/gofiber/fiber/v2"
	"strconv"
	"encoding/json"
)


func GetPpdmException(c *fiber.Ctx) error {
	rows, err := db.Query("select * from ppdm_exception")
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}
	defer rows.Close()
	var result []dto.Ppdm_exception

	for rows.Next() {
		var ppdm_exception dto.Ppdm_exception
		if err := rows.Scan(&ppdm_exception.Pe_id, &ppdm_exception.Active_ind, &ppdm_exception.Constraint_full_name, &ppdm_exception.Effective_date, &ppdm_exception.Expiry_date, &ppdm_exception.Owner, &ppdm_exception.Ppdm_guid, &ppdm_exception.Remark, &ppdm_exception.Row_id, &ppdm_exception.Source, &ppdm_exception.System_id, &ppdm_exception.Table_name, &ppdm_exception.Row_changed_by, &ppdm_exception.Row_changed_date, &ppdm_exception.Row_created_by, &ppdm_exception.Row_created_date, &ppdm_exception.Row_effective_date, &ppdm_exception.Row_expiry_date, &ppdm_exception.Row_quality); err != nil {
			return err // Exit if we get an error
		}

		// Append ppdm_exception to result
		result = append(result, ppdm_exception)
	}
	// Return result in JSON format
	return c.JSON(result)
	 
}

func SetPpdmException(c *fiber.Ctx) error {
	var ppdm_exception dto.Ppdm_exception

	setDefaults(&ppdm_exception)

	if err := json.Unmarshal(c.Body(), &ppdm_exception); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	stmt, err := db.Prepare("insert into ppdm_exception values(:1, :2, :3, :4, :5, :6, :7, :8, :9, :10, :11, :12, :13, :14, :15, :16, :17, :18, :19)")
	if err != nil {
		return err
	}
	ppdm_exception.Row_created_date = formatDate(ppdm_exception.Row_created_date)
	ppdm_exception.Row_changed_date = formatDate(ppdm_exception.Row_changed_date)
	ppdm_exception.Effective_date = formatDateString(ppdm_exception.Effective_date)
	ppdm_exception.Expiry_date = formatDateString(ppdm_exception.Expiry_date)
	ppdm_exception.Row_effective_date = formatDateString(ppdm_exception.Row_effective_date)
	ppdm_exception.Row_expiry_date = formatDateString(ppdm_exception.Row_expiry_date)






	rows, err := stmt.Exec(ppdm_exception.Pe_id, ppdm_exception.Active_ind, ppdm_exception.Constraint_full_name, ppdm_exception.Effective_date, ppdm_exception.Expiry_date, ppdm_exception.Owner, ppdm_exception.Ppdm_guid, ppdm_exception.Remark, ppdm_exception.Row_id, ppdm_exception.Source, ppdm_exception.System_id, ppdm_exception.Table_name, ppdm_exception.Row_changed_by, ppdm_exception.Row_changed_date, ppdm_exception.Row_created_by, ppdm_exception.Row_created_date, ppdm_exception.Row_effective_date, ppdm_exception.Row_expiry_date, ppdm_exception.Row_quality)
	
	if err != nil {
		return err
	}

	return c.JSON(rows)
}

func UpdatePpdmException(c *fiber.Ctx) error {
	var ppdm_exception dto.Ppdm_exception

	setDefaults(&ppdm_exception)

	if err := json.Unmarshal(c.Body(), &ppdm_exception); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	id := c.Params("id")
	ppdm_exception.Pe_id = id

    if ppdm_exception.Row_created_date != nil {
        return c.Status(400).SendString("Cannot update row_created_date")
    }

	stmt, err := db.Prepare("update ppdm_exception set active_ind = :1, constraint_full_name = :2, effective_date = :3, expiry_date = :4, owner = :5, ppdm_guid = :6, remark = :7, row_id = :8, source = :9, system_id = :10, table_name = :11, row_changed_by = :12, row_changed_date = :13, row_created_by = :14, row_effective_date = :15, row_expiry_date = :16, row_quality = :17 where pe_id = :19")
	if err != nil {
		return err
	}

	ppdm_exception.Row_changed_date = formatDate(ppdm_exception.Row_changed_date)
	ppdm_exception.Effective_date = formatDateString(ppdm_exception.Effective_date)
	ppdm_exception.Expiry_date = formatDateString(ppdm_exception.Expiry_date)
	ppdm_exception.Row_effective_date = formatDateString(ppdm_exception.Row_effective_date)
	ppdm_exception.Row_expiry_date = formatDateString(ppdm_exception.Row_expiry_date)






	rows, err := stmt.Exec(ppdm_exception.Active_ind, ppdm_exception.Constraint_full_name, ppdm_exception.Effective_date, ppdm_exception.Expiry_date, ppdm_exception.Owner, ppdm_exception.Ppdm_guid, ppdm_exception.Remark, ppdm_exception.Row_id, ppdm_exception.Source, ppdm_exception.System_id, ppdm_exception.Table_name, ppdm_exception.Row_changed_by, ppdm_exception.Row_changed_date, ppdm_exception.Row_created_by, ppdm_exception.Row_effective_date, ppdm_exception.Row_expiry_date, ppdm_exception.Row_quality, ppdm_exception.Pe_id)
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

func PatchPpdmException(c *fiber.Ctx) error {
	id := c.Params("id")

	updatedFields := make(map[string]interface{})

	if err := c.BodyParser(&updatedFields); err != nil {
		return err
	}
	if _, ok := updatedFields["row_created_date"]; ok {
        return c.Status(400).SendString("Cannot update row_created_date field")
    }

	query := "update ppdm_exception set "
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
	query += " where pe_id = :id"

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

func DeletePpdmException(c *fiber.Ctx) error {
	id := c.Params("id")
	var ppdm_exception dto.Ppdm_exception
	ppdm_exception.Pe_id = id

	stmt, err := db.Prepare("delete from ppdm_exception where pe_id = :1")
	if err != nil {
		return err
	}

	rows, err := stmt.Exec(ppdm_exception.Pe_id)
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


