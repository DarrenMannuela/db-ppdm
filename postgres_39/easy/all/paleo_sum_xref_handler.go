package apiv1

import (
	dto "github.com/DarrenMannuela/pt_gtn_bibliography/dto"
	"github.com/gofiber/fiber/v2"
	"strconv"
	"encoding/json"
)


func GetPaleoSumXref(c *fiber.Ctx) error {
	rows, err := db.Query("select * from paleo_sum_xref")
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}
	defer rows.Close()
	var result []dto.Paleo_sum_xref

	for rows.Next() {
		var paleo_sum_xref dto.Paleo_sum_xref
		if err := rows.Scan(&paleo_sum_xref.Paleo_summary_id1, &paleo_sum_xref.Paleo_summary_id2, &paleo_sum_xref.Active_ind, &paleo_sum_xref.Effective_date, &paleo_sum_xref.Expiry_date, &paleo_sum_xref.Paleo_xref_type, &paleo_sum_xref.Ppdm_guid, &paleo_sum_xref.Remark, &paleo_sum_xref.Source, &paleo_sum_xref.Row_changed_by, &paleo_sum_xref.Row_changed_date, &paleo_sum_xref.Row_created_by, &paleo_sum_xref.Row_created_date, &paleo_sum_xref.Row_effective_date, &paleo_sum_xref.Row_expiry_date, &paleo_sum_xref.Row_quality); err != nil {
			return err // Exit if we get an error
		}

		// Append paleo_sum_xref to result
		result = append(result, paleo_sum_xref)
	}
	// Return result in JSON format
	return c.JSON(result)
	 
}

func SetPaleoSumXref(c *fiber.Ctx) error {
	var paleo_sum_xref dto.Paleo_sum_xref

	setDefaults(&paleo_sum_xref)

	if err := json.Unmarshal(c.Body(), &paleo_sum_xref); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	stmt, err := db.Prepare("insert into paleo_sum_xref values(:1, :2, :3, :4, :5, :6, :7, :8, :9, :10, :11, :12, :13, :14, :15, :16)")
	if err != nil {
		return err
	}
	paleo_sum_xref.Row_created_date = formatDate(paleo_sum_xref.Row_created_date)
	paleo_sum_xref.Row_changed_date = formatDate(paleo_sum_xref.Row_changed_date)
	paleo_sum_xref.Effective_date = formatDateString(paleo_sum_xref.Effective_date)
	paleo_sum_xref.Expiry_date = formatDateString(paleo_sum_xref.Expiry_date)
	paleo_sum_xref.Row_effective_date = formatDateString(paleo_sum_xref.Row_effective_date)
	paleo_sum_xref.Row_expiry_date = formatDateString(paleo_sum_xref.Row_expiry_date)






	rows, err := stmt.Exec(paleo_sum_xref.Paleo_summary_id1, paleo_sum_xref.Paleo_summary_id2, paleo_sum_xref.Active_ind, paleo_sum_xref.Effective_date, paleo_sum_xref.Expiry_date, paleo_sum_xref.Paleo_xref_type, paleo_sum_xref.Ppdm_guid, paleo_sum_xref.Remark, paleo_sum_xref.Source, paleo_sum_xref.Row_changed_by, paleo_sum_xref.Row_changed_date, paleo_sum_xref.Row_created_by, paleo_sum_xref.Row_created_date, paleo_sum_xref.Row_effective_date, paleo_sum_xref.Row_expiry_date, paleo_sum_xref.Row_quality)
	
	if err != nil {
		return err
	}

	return c.JSON(rows)
}

func UpdatePaleoSumXref(c *fiber.Ctx) error {
	var paleo_sum_xref dto.Paleo_sum_xref

	setDefaults(&paleo_sum_xref)

	if err := json.Unmarshal(c.Body(), &paleo_sum_xref); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	id := c.Params("id")
	paleo_sum_xref.Paleo_summary_id1 = id

    if paleo_sum_xref.Row_created_date != nil {
        return c.Status(400).SendString("Cannot update row_created_date")
    }

	stmt, err := db.Prepare("update paleo_sum_xref set paleo_summary_id2 = :1, active_ind = :2, effective_date = :3, expiry_date = :4, paleo_xref_type = :5, ppdm_guid = :6, remark = :7, source = :8, row_changed_by = :9, row_changed_date = :10, row_created_by = :11, row_effective_date = :12, row_expiry_date = :13, row_quality = :14 where paleo_summary_id1 = :16")
	if err != nil {
		return err
	}

	paleo_sum_xref.Row_changed_date = formatDate(paleo_sum_xref.Row_changed_date)
	paleo_sum_xref.Effective_date = formatDateString(paleo_sum_xref.Effective_date)
	paleo_sum_xref.Expiry_date = formatDateString(paleo_sum_xref.Expiry_date)
	paleo_sum_xref.Row_effective_date = formatDateString(paleo_sum_xref.Row_effective_date)
	paleo_sum_xref.Row_expiry_date = formatDateString(paleo_sum_xref.Row_expiry_date)






	rows, err := stmt.Exec(paleo_sum_xref.Paleo_summary_id2, paleo_sum_xref.Active_ind, paleo_sum_xref.Effective_date, paleo_sum_xref.Expiry_date, paleo_sum_xref.Paleo_xref_type, paleo_sum_xref.Ppdm_guid, paleo_sum_xref.Remark, paleo_sum_xref.Source, paleo_sum_xref.Row_changed_by, paleo_sum_xref.Row_changed_date, paleo_sum_xref.Row_created_by, paleo_sum_xref.Row_effective_date, paleo_sum_xref.Row_expiry_date, paleo_sum_xref.Row_quality, paleo_sum_xref.Paleo_summary_id1)
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

func PatchPaleoSumXref(c *fiber.Ctx) error {
	id := c.Params("id")

	updatedFields := make(map[string]interface{})

	if err := c.BodyParser(&updatedFields); err != nil {
		return err
	}
	if _, ok := updatedFields["row_created_date"]; ok {
        return c.Status(400).SendString("Cannot update row_created_date field")
    }

	query := "update paleo_sum_xref set "
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
	query += " where paleo_summary_id1 = :id"

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

func DeletePaleoSumXref(c *fiber.Ctx) error {
	id := c.Params("id")
	var paleo_sum_xref dto.Paleo_sum_xref
	paleo_sum_xref.Paleo_summary_id1 = id

	stmt, err := db.Prepare("delete from paleo_sum_xref where paleo_summary_id1 = :1")
	if err != nil {
		return err
	}

	rows, err := stmt.Exec(paleo_sum_xref.Paleo_summary_id1)
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


