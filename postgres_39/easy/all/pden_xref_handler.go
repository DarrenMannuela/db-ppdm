package apiv1

import (
	dto "github.com/DarrenMannuela/pt_gtn_bibliography/dto"
	"github.com/gofiber/fiber/v2"
	"strconv"
	"encoding/json"
)


func GetPdenXref(c *fiber.Ctx) error {
	rows, err := db.Query("select * from pden_xref")
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}
	defer rows.Close()
	var result []dto.Pden_xref

	for rows.Next() {
		var pden_xref dto.Pden_xref
		if err := rows.Scan(&pden_xref.From_pden_subtype, &pden_xref.From_pden_id, &pden_xref.From_pden_source, &pden_xref.To_pden_subtype, &pden_xref.To_pden_id, &pden_xref.To_pden_source, &pden_xref.Xref_obs_no, &pden_xref.Active_ind, &pden_xref.Effective_date, &pden_xref.Expiry_date, &pden_xref.Pden_xref_type, &pden_xref.Ppdm_guid, &pden_xref.Remark, &pden_xref.Source, &pden_xref.Row_changed_by, &pden_xref.Row_changed_date, &pden_xref.Row_created_by, &pden_xref.Row_created_date, &pden_xref.Row_effective_date, &pden_xref.Row_expiry_date, &pden_xref.Row_quality); err != nil {
			return err // Exit if we get an error
		}

		// Append pden_xref to result
		result = append(result, pden_xref)
	}
	// Return result in JSON format
	return c.JSON(result)
	 
}

func SetPdenXref(c *fiber.Ctx) error {
	var pden_xref dto.Pden_xref

	setDefaults(&pden_xref)

	if err := json.Unmarshal(c.Body(), &pden_xref); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	stmt, err := db.Prepare("insert into pden_xref values(:1, :2, :3, :4, :5, :6, :7, :8, :9, :10, :11, :12, :13, :14, :15, :16, :17, :18, :19, :20, :21)")
	if err != nil {
		return err
	}
	pden_xref.Row_created_date = formatDate(pden_xref.Row_created_date)
	pden_xref.Row_changed_date = formatDate(pden_xref.Row_changed_date)
	pden_xref.Effective_date = formatDateString(pden_xref.Effective_date)
	pden_xref.Expiry_date = formatDateString(pden_xref.Expiry_date)
	pden_xref.Row_effective_date = formatDateString(pden_xref.Row_effective_date)
	pden_xref.Row_expiry_date = formatDateString(pden_xref.Row_expiry_date)






	rows, err := stmt.Exec(pden_xref.From_pden_subtype, pden_xref.From_pden_id, pden_xref.From_pden_source, pden_xref.To_pden_subtype, pden_xref.To_pden_id, pden_xref.To_pden_source, pden_xref.Xref_obs_no, pden_xref.Active_ind, pden_xref.Effective_date, pden_xref.Expiry_date, pden_xref.Pden_xref_type, pden_xref.Ppdm_guid, pden_xref.Remark, pden_xref.Source, pden_xref.Row_changed_by, pden_xref.Row_changed_date, pden_xref.Row_created_by, pden_xref.Row_created_date, pden_xref.Row_effective_date, pden_xref.Row_expiry_date, pden_xref.Row_quality)
	
	if err != nil {
		return err
	}

	return c.JSON(rows)
}

func UpdatePdenXref(c *fiber.Ctx) error {
	var pden_xref dto.Pden_xref

	setDefaults(&pden_xref)

	if err := json.Unmarshal(c.Body(), &pden_xref); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	id := c.Params("id")
	pden_xref.From_pden_subtype = id

    if pden_xref.Row_created_date != nil {
        return c.Status(400).SendString("Cannot update row_created_date")
    }

	stmt, err := db.Prepare("update pden_xref set from_pden_id = :1, from_pden_source = :2, to_pden_subtype = :3, to_pden_id = :4, to_pden_source = :5, xref_obs_no = :6, active_ind = :7, effective_date = :8, expiry_date = :9, pden_xref_type = :10, ppdm_guid = :11, remark = :12, source = :13, row_changed_by = :14, row_changed_date = :15, row_created_by = :16, row_effective_date = :17, row_expiry_date = :18, row_quality = :19 where from_pden_subtype = :21")
	if err != nil {
		return err
	}

	pden_xref.Row_changed_date = formatDate(pden_xref.Row_changed_date)
	pden_xref.Effective_date = formatDateString(pden_xref.Effective_date)
	pden_xref.Expiry_date = formatDateString(pden_xref.Expiry_date)
	pden_xref.Row_effective_date = formatDateString(pden_xref.Row_effective_date)
	pden_xref.Row_expiry_date = formatDateString(pden_xref.Row_expiry_date)






	rows, err := stmt.Exec(pden_xref.From_pden_id, pden_xref.From_pden_source, pden_xref.To_pden_subtype, pden_xref.To_pden_id, pden_xref.To_pden_source, pden_xref.Xref_obs_no, pden_xref.Active_ind, pden_xref.Effective_date, pden_xref.Expiry_date, pden_xref.Pden_xref_type, pden_xref.Ppdm_guid, pden_xref.Remark, pden_xref.Source, pden_xref.Row_changed_by, pden_xref.Row_changed_date, pden_xref.Row_created_by, pden_xref.Row_effective_date, pden_xref.Row_expiry_date, pden_xref.Row_quality, pden_xref.From_pden_subtype)
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

func PatchPdenXref(c *fiber.Ctx) error {
	id := c.Params("id")

	updatedFields := make(map[string]interface{})

	if err := c.BodyParser(&updatedFields); err != nil {
		return err
	}
	if _, ok := updatedFields["row_created_date"]; ok {
        return c.Status(400).SendString("Cannot update row_created_date field")
    }

	query := "update pden_xref set "
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
	query += " where from_pden_subtype = :id"

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

func DeletePdenXref(c *fiber.Ctx) error {
	id := c.Params("id")
	var pden_xref dto.Pden_xref
	pden_xref.From_pden_subtype = id

	stmt, err := db.Prepare("delete from pden_xref where from_pden_subtype = :1")
	if err != nil {
		return err
	}

	rows, err := stmt.Exec(pden_xref.From_pden_subtype)
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


