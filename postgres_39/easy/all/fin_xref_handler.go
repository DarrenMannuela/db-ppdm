package apiv1

import (
	dto "github.com/DarrenMannuela/pt_gtn_bibliography/dto"
	"github.com/gofiber/fiber/v2"
	"strconv"
	"encoding/json"
)


func GetFinXref(c *fiber.Ctx) error {
	rows, err := db.Query("select * from fin_xref")
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}
	defer rows.Close()
	var result []dto.Fin_xref

	for rows.Next() {
		var fin_xref dto.Fin_xref
		if err := rows.Scan(&fin_xref.Finance_id1, &fin_xref.Finance_id2, &fin_xref.Active_ind, &fin_xref.Distribution_percent, &fin_xref.Effective_date, &fin_xref.Expiry_date, &fin_xref.Fin_xref_type, &fin_xref.Ppdm_guid, &fin_xref.Remark, &fin_xref.Source, &fin_xref.Row_changed_by, &fin_xref.Row_changed_date, &fin_xref.Row_created_by, &fin_xref.Row_created_date, &fin_xref.Row_effective_date, &fin_xref.Row_expiry_date, &fin_xref.Row_quality); err != nil {
			return err // Exit if we get an error
		}

		// Append fin_xref to result
		result = append(result, fin_xref)
	}
	// Return result in JSON format
	return c.JSON(result)
	 
}

func SetFinXref(c *fiber.Ctx) error {
	var fin_xref dto.Fin_xref

	setDefaults(&fin_xref)

	if err := json.Unmarshal(c.Body(), &fin_xref); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	stmt, err := db.Prepare("insert into fin_xref values(:1, :2, :3, :4, :5, :6, :7, :8, :9, :10, :11, :12, :13, :14, :15, :16, :17)")
	if err != nil {
		return err
	}
	fin_xref.Row_created_date = formatDate(fin_xref.Row_created_date)
	fin_xref.Row_changed_date = formatDate(fin_xref.Row_changed_date)
	fin_xref.Effective_date = formatDateString(fin_xref.Effective_date)
	fin_xref.Expiry_date = formatDateString(fin_xref.Expiry_date)
	fin_xref.Row_effective_date = formatDateString(fin_xref.Row_effective_date)
	fin_xref.Row_expiry_date = formatDateString(fin_xref.Row_expiry_date)






	rows, err := stmt.Exec(fin_xref.Finance_id1, fin_xref.Finance_id2, fin_xref.Active_ind, fin_xref.Distribution_percent, fin_xref.Effective_date, fin_xref.Expiry_date, fin_xref.Fin_xref_type, fin_xref.Ppdm_guid, fin_xref.Remark, fin_xref.Source, fin_xref.Row_changed_by, fin_xref.Row_changed_date, fin_xref.Row_created_by, fin_xref.Row_created_date, fin_xref.Row_effective_date, fin_xref.Row_expiry_date, fin_xref.Row_quality)
	
	if err != nil {
		return err
	}

	return c.JSON(rows)
}

func UpdateFinXref(c *fiber.Ctx) error {
	var fin_xref dto.Fin_xref

	setDefaults(&fin_xref)

	if err := json.Unmarshal(c.Body(), &fin_xref); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	id := c.Params("id")
	fin_xref.Finance_id1 = id

    if fin_xref.Row_created_date != nil {
        return c.Status(400).SendString("Cannot update row_created_date")
    }

	stmt, err := db.Prepare("update fin_xref set finance_id2 = :1, active_ind = :2, distribution_percent = :3, effective_date = :4, expiry_date = :5, fin_xref_type = :6, ppdm_guid = :7, remark = :8, source = :9, row_changed_by = :10, row_changed_date = :11, row_created_by = :12, row_effective_date = :13, row_expiry_date = :14, row_quality = :15 where finance_id1 = :17")
	if err != nil {
		return err
	}

	fin_xref.Row_changed_date = formatDate(fin_xref.Row_changed_date)
	fin_xref.Effective_date = formatDateString(fin_xref.Effective_date)
	fin_xref.Expiry_date = formatDateString(fin_xref.Expiry_date)
	fin_xref.Row_effective_date = formatDateString(fin_xref.Row_effective_date)
	fin_xref.Row_expiry_date = formatDateString(fin_xref.Row_expiry_date)






	rows, err := stmt.Exec(fin_xref.Finance_id2, fin_xref.Active_ind, fin_xref.Distribution_percent, fin_xref.Effective_date, fin_xref.Expiry_date, fin_xref.Fin_xref_type, fin_xref.Ppdm_guid, fin_xref.Remark, fin_xref.Source, fin_xref.Row_changed_by, fin_xref.Row_changed_date, fin_xref.Row_created_by, fin_xref.Row_effective_date, fin_xref.Row_expiry_date, fin_xref.Row_quality, fin_xref.Finance_id1)
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

func PatchFinXref(c *fiber.Ctx) error {
	id := c.Params("id")

	updatedFields := make(map[string]interface{})

	if err := c.BodyParser(&updatedFields); err != nil {
		return err
	}
	if _, ok := updatedFields["row_created_date"]; ok {
        return c.Status(400).SendString("Cannot update row_created_date field")
    }

	query := "update fin_xref set "
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
	query += " where finance_id1 = :id"

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

func DeleteFinXref(c *fiber.Ctx) error {
	id := c.Params("id")
	var fin_xref dto.Fin_xref
	fin_xref.Finance_id1 = id

	stmt, err := db.Prepare("delete from fin_xref where finance_id1 = :1")
	if err != nil {
		return err
	}

	rows, err := stmt.Exec(fin_xref.Finance_id1)
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


