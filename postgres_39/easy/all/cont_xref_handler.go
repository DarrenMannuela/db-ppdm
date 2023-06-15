package apiv1

import (
	dto "github.com/DarrenMannuela/pt_gtn_bibliography/dto"
	"github.com/gofiber/fiber/v2"
	"strconv"
	"encoding/json"
)


func GetContXref(c *fiber.Ctx) error {
	rows, err := db.Query("select * from cont_xref")
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}
	defer rows.Close()
	var result []dto.Cont_xref

	for rows.Next() {
		var cont_xref dto.Cont_xref
		if err := rows.Scan(&cont_xref.Contract_id, &cont_xref.Contract_id_2, &cont_xref.Active_ind, &cont_xref.Contract_xref_type, &cont_xref.Effective_date, &cont_xref.Expiry_date, &cont_xref.Order_of_supercedence, &cont_xref.Ppdm_guid, &cont_xref.Remark, &cont_xref.Source, &cont_xref.Row_changed_by, &cont_xref.Row_changed_date, &cont_xref.Row_created_by, &cont_xref.Row_created_date, &cont_xref.Row_effective_date, &cont_xref.Row_expiry_date, &cont_xref.Row_quality); err != nil {
			return err // Exit if we get an error
		}

		// Append cont_xref to result
		result = append(result, cont_xref)
	}
	// Return result in JSON format
	return c.JSON(result)
	 
}

func SetContXref(c *fiber.Ctx) error {
	var cont_xref dto.Cont_xref

	setDefaults(&cont_xref)

	if err := json.Unmarshal(c.Body(), &cont_xref); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	stmt, err := db.Prepare("insert into cont_xref values(:1, :2, :3, :4, :5, :6, :7, :8, :9, :10, :11, :12, :13, :14, :15, :16, :17)")
	if err != nil {
		return err
	}
	cont_xref.Row_created_date = formatDate(cont_xref.Row_created_date)
	cont_xref.Row_changed_date = formatDate(cont_xref.Row_changed_date)
	cont_xref.Effective_date = formatDateString(cont_xref.Effective_date)
	cont_xref.Expiry_date = formatDateString(cont_xref.Expiry_date)
	cont_xref.Row_effective_date = formatDateString(cont_xref.Row_effective_date)
	cont_xref.Row_expiry_date = formatDateString(cont_xref.Row_expiry_date)






	rows, err := stmt.Exec(cont_xref.Contract_id, cont_xref.Contract_id_2, cont_xref.Active_ind, cont_xref.Contract_xref_type, cont_xref.Effective_date, cont_xref.Expiry_date, cont_xref.Order_of_supercedence, cont_xref.Ppdm_guid, cont_xref.Remark, cont_xref.Source, cont_xref.Row_changed_by, cont_xref.Row_changed_date, cont_xref.Row_created_by, cont_xref.Row_created_date, cont_xref.Row_effective_date, cont_xref.Row_expiry_date, cont_xref.Row_quality)
	
	if err != nil {
		return err
	}

	return c.JSON(rows)
}

func UpdateContXref(c *fiber.Ctx) error {
	var cont_xref dto.Cont_xref

	setDefaults(&cont_xref)

	if err := json.Unmarshal(c.Body(), &cont_xref); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	id := c.Params("id")
	cont_xref.Contract_id = id

    if cont_xref.Row_created_date != nil {
        return c.Status(400).SendString("Cannot update row_created_date")
    }

	stmt, err := db.Prepare("update cont_xref set contract_id_2 = :1, active_ind = :2, contract_xref_type = :3, effective_date = :4, expiry_date = :5, order_of_supercedence = :6, ppdm_guid = :7, remark = :8, source = :9, row_changed_by = :10, row_changed_date = :11, row_created_by = :12, row_effective_date = :13, row_expiry_date = :14, row_quality = :15 where contract_id = :17")
	if err != nil {
		return err
	}

	cont_xref.Row_changed_date = formatDate(cont_xref.Row_changed_date)
	cont_xref.Effective_date = formatDateString(cont_xref.Effective_date)
	cont_xref.Expiry_date = formatDateString(cont_xref.Expiry_date)
	cont_xref.Row_effective_date = formatDateString(cont_xref.Row_effective_date)
	cont_xref.Row_expiry_date = formatDateString(cont_xref.Row_expiry_date)






	rows, err := stmt.Exec(cont_xref.Contract_id_2, cont_xref.Active_ind, cont_xref.Contract_xref_type, cont_xref.Effective_date, cont_xref.Expiry_date, cont_xref.Order_of_supercedence, cont_xref.Ppdm_guid, cont_xref.Remark, cont_xref.Source, cont_xref.Row_changed_by, cont_xref.Row_changed_date, cont_xref.Row_created_by, cont_xref.Row_effective_date, cont_xref.Row_expiry_date, cont_xref.Row_quality, cont_xref.Contract_id)
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

func PatchContXref(c *fiber.Ctx) error {
	id := c.Params("id")

	updatedFields := make(map[string]interface{})

	if err := c.BodyParser(&updatedFields); err != nil {
		return err
	}
	if _, ok := updatedFields["row_created_date"]; ok {
        return c.Status(400).SendString("Cannot update row_created_date field")
    }

	query := "update cont_xref set "
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

func DeleteContXref(c *fiber.Ctx) error {
	id := c.Params("id")
	var cont_xref dto.Cont_xref
	cont_xref.Contract_id = id

	stmt, err := db.Prepare("delete from cont_xref where contract_id = :1")
	if err != nil {
		return err
	}

	rows, err := stmt.Exec(cont_xref.Contract_id)
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


