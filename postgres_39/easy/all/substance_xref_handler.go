package apiv1

import (
	dto "github.com/DarrenMannuela/pt_gtn_bibliography/dto"
	"github.com/gofiber/fiber/v2"
	"strconv"
	"encoding/json"
)


func GetSubstanceXref(c *fiber.Ctx) error {
	rows, err := db.Query("select * from substance_xref")
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}
	defer rows.Close()
	var result []dto.Substance_xref

	for rows.Next() {
		var substance_xref dto.Substance_xref
		if err := rows.Scan(&substance_xref.Substance_id1, &substance_xref.Substance_id2, &substance_xref.Substance_xref_obs_no, &substance_xref.Active_ind, &substance_xref.Effective_date, &substance_xref.Expiry_date, &substance_xref.Ppdm_guid, &substance_xref.Remark, &substance_xref.Source, &substance_xref.Substance_xref_type, &substance_xref.Row_changed_by, &substance_xref.Row_changed_date, &substance_xref.Row_created_by, &substance_xref.Row_created_date, &substance_xref.Row_effective_date, &substance_xref.Row_expiry_date, &substance_xref.Row_quality); err != nil {
			return err // Exit if we get an error
		}

		// Append substance_xref to result
		result = append(result, substance_xref)
	}
	// Return result in JSON format
	return c.JSON(result)
	 
}

func SetSubstanceXref(c *fiber.Ctx) error {
	var substance_xref dto.Substance_xref

	setDefaults(&substance_xref)

	if err := json.Unmarshal(c.Body(), &substance_xref); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	stmt, err := db.Prepare("insert into substance_xref values(:1, :2, :3, :4, :5, :6, :7, :8, :9, :10, :11, :12, :13, :14, :15, :16, :17)")
	if err != nil {
		return err
	}
	substance_xref.Row_created_date = formatDate(substance_xref.Row_created_date)
	substance_xref.Row_changed_date = formatDate(substance_xref.Row_changed_date)
	substance_xref.Effective_date = formatDateString(substance_xref.Effective_date)
	substance_xref.Expiry_date = formatDateString(substance_xref.Expiry_date)
	substance_xref.Row_effective_date = formatDateString(substance_xref.Row_effective_date)
	substance_xref.Row_expiry_date = formatDateString(substance_xref.Row_expiry_date)






	rows, err := stmt.Exec(substance_xref.Substance_id1, substance_xref.Substance_id2, substance_xref.Substance_xref_obs_no, substance_xref.Active_ind, substance_xref.Effective_date, substance_xref.Expiry_date, substance_xref.Ppdm_guid, substance_xref.Remark, substance_xref.Source, substance_xref.Substance_xref_type, substance_xref.Row_changed_by, substance_xref.Row_changed_date, substance_xref.Row_created_by, substance_xref.Row_created_date, substance_xref.Row_effective_date, substance_xref.Row_expiry_date, substance_xref.Row_quality)
	
	if err != nil {
		return err
	}

	return c.JSON(rows)
}

func UpdateSubstanceXref(c *fiber.Ctx) error {
	var substance_xref dto.Substance_xref

	setDefaults(&substance_xref)

	if err := json.Unmarshal(c.Body(), &substance_xref); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	id := c.Params("id")
	substance_xref.Substance_id1 = id

    if substance_xref.Row_created_date != nil {
        return c.Status(400).SendString("Cannot update row_created_date")
    }

	stmt, err := db.Prepare("update substance_xref set substance_id2 = :1, substance_xref_obs_no = :2, active_ind = :3, effective_date = :4, expiry_date = :5, ppdm_guid = :6, remark = :7, source = :8, substance_xref_type = :9, row_changed_by = :10, row_changed_date = :11, row_created_by = :12, row_effective_date = :13, row_expiry_date = :14, row_quality = :15 where substance_id1 = :17")
	if err != nil {
		return err
	}

	substance_xref.Row_changed_date = formatDate(substance_xref.Row_changed_date)
	substance_xref.Effective_date = formatDateString(substance_xref.Effective_date)
	substance_xref.Expiry_date = formatDateString(substance_xref.Expiry_date)
	substance_xref.Row_effective_date = formatDateString(substance_xref.Row_effective_date)
	substance_xref.Row_expiry_date = formatDateString(substance_xref.Row_expiry_date)






	rows, err := stmt.Exec(substance_xref.Substance_id2, substance_xref.Substance_xref_obs_no, substance_xref.Active_ind, substance_xref.Effective_date, substance_xref.Expiry_date, substance_xref.Ppdm_guid, substance_xref.Remark, substance_xref.Source, substance_xref.Substance_xref_type, substance_xref.Row_changed_by, substance_xref.Row_changed_date, substance_xref.Row_created_by, substance_xref.Row_effective_date, substance_xref.Row_expiry_date, substance_xref.Row_quality, substance_xref.Substance_id1)
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

func PatchSubstanceXref(c *fiber.Ctx) error {
	id := c.Params("id")

	updatedFields := make(map[string]interface{})

	if err := c.BodyParser(&updatedFields); err != nil {
		return err
	}
	if _, ok := updatedFields["row_created_date"]; ok {
        return c.Status(400).SendString("Cannot update row_created_date field")
    }

	query := "update substance_xref set "
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
	query += " where substance_id1 = :id"

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

func DeleteSubstanceXref(c *fiber.Ctx) error {
	id := c.Params("id")
	var substance_xref dto.Substance_xref
	substance_xref.Substance_id1 = id

	stmt, err := db.Prepare("delete from substance_xref where substance_id1 = :1")
	if err != nil {
		return err
	}

	rows, err := stmt.Exec(substance_xref.Substance_id1)
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


