package apiv1

import (
	dto "github.com/DarrenMannuela/pt_gtn_bibliography/dto"
	"github.com/gofiber/fiber/v2"
	"strconv"
	"encoding/json"
)


func GetFossilEquivalence(c *fiber.Ctx) error {
	rows, err := db.Query("select * from fossil_equivalence")
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}
	defer rows.Close()
	var result []dto.Fossil_equivalence

	for rows.Next() {
		var fossil_equivalence dto.Fossil_equivalence
		if err := rows.Scan(&fossil_equivalence.Fossil_id, &fossil_equivalence.Equiv_fossil_id, &fossil_equivalence.Active_ind, &fossil_equivalence.Effective_date, &fossil_equivalence.Expiry_date, &fossil_equivalence.Ppdm_guid, &fossil_equivalence.Remark, &fossil_equivalence.Source, &fossil_equivalence.Source_document_id, &fossil_equivalence.Row_changed_by, &fossil_equivalence.Row_changed_date, &fossil_equivalence.Row_created_by, &fossil_equivalence.Row_created_date, &fossil_equivalence.Row_effective_date, &fossil_equivalence.Row_expiry_date, &fossil_equivalence.Row_quality); err != nil {
			return err // Exit if we get an error
		}

		// Append fossil_equivalence to result
		result = append(result, fossil_equivalence)
	}
	// Return result in JSON format
	return c.JSON(result)
	 
}

func SetFossilEquivalence(c *fiber.Ctx) error {
	var fossil_equivalence dto.Fossil_equivalence

	setDefaults(&fossil_equivalence)

	if err := json.Unmarshal(c.Body(), &fossil_equivalence); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	stmt, err := db.Prepare("insert into fossil_equivalence values(:1, :2, :3, :4, :5, :6, :7, :8, :9, :10, :11, :12, :13, :14, :15, :16)")
	if err != nil {
		return err
	}
	fossil_equivalence.Row_created_date = formatDate(fossil_equivalence.Row_created_date)
	fossil_equivalence.Row_changed_date = formatDate(fossil_equivalence.Row_changed_date)
	fossil_equivalence.Effective_date = formatDateString(fossil_equivalence.Effective_date)
	fossil_equivalence.Expiry_date = formatDateString(fossil_equivalence.Expiry_date)
	fossil_equivalence.Row_effective_date = formatDateString(fossil_equivalence.Row_effective_date)
	fossil_equivalence.Row_expiry_date = formatDateString(fossil_equivalence.Row_expiry_date)






	rows, err := stmt.Exec(fossil_equivalence.Fossil_id, fossil_equivalence.Equiv_fossil_id, fossil_equivalence.Active_ind, fossil_equivalence.Effective_date, fossil_equivalence.Expiry_date, fossil_equivalence.Ppdm_guid, fossil_equivalence.Remark, fossil_equivalence.Source, fossil_equivalence.Source_document_id, fossil_equivalence.Row_changed_by, fossil_equivalence.Row_changed_date, fossil_equivalence.Row_created_by, fossil_equivalence.Row_created_date, fossil_equivalence.Row_effective_date, fossil_equivalence.Row_expiry_date, fossil_equivalence.Row_quality)
	
	if err != nil {
		return err
	}

	return c.JSON(rows)
}

func UpdateFossilEquivalence(c *fiber.Ctx) error {
	var fossil_equivalence dto.Fossil_equivalence

	setDefaults(&fossil_equivalence)

	if err := json.Unmarshal(c.Body(), &fossil_equivalence); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	id := c.Params("id")
	fossil_equivalence.Fossil_id = id

    if fossil_equivalence.Row_created_date != nil {
        return c.Status(400).SendString("Cannot update row_created_date")
    }

	stmt, err := db.Prepare("update fossil_equivalence set equiv_fossil_id = :1, active_ind = :2, effective_date = :3, expiry_date = :4, ppdm_guid = :5, remark = :6, source = :7, source_document_id = :8, row_changed_by = :9, row_changed_date = :10, row_created_by = :11, row_effective_date = :12, row_expiry_date = :13, row_quality = :14 where fossil_id = :16")
	if err != nil {
		return err
	}

	fossil_equivalence.Row_changed_date = formatDate(fossil_equivalence.Row_changed_date)
	fossil_equivalence.Effective_date = formatDateString(fossil_equivalence.Effective_date)
	fossil_equivalence.Expiry_date = formatDateString(fossil_equivalence.Expiry_date)
	fossil_equivalence.Row_effective_date = formatDateString(fossil_equivalence.Row_effective_date)
	fossil_equivalence.Row_expiry_date = formatDateString(fossil_equivalence.Row_expiry_date)






	rows, err := stmt.Exec(fossil_equivalence.Equiv_fossil_id, fossil_equivalence.Active_ind, fossil_equivalence.Effective_date, fossil_equivalence.Expiry_date, fossil_equivalence.Ppdm_guid, fossil_equivalence.Remark, fossil_equivalence.Source, fossil_equivalence.Source_document_id, fossil_equivalence.Row_changed_by, fossil_equivalence.Row_changed_date, fossil_equivalence.Row_created_by, fossil_equivalence.Row_effective_date, fossil_equivalence.Row_expiry_date, fossil_equivalence.Row_quality, fossil_equivalence.Fossil_id)
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

func PatchFossilEquivalence(c *fiber.Ctx) error {
	id := c.Params("id")

	updatedFields := make(map[string]interface{})

	if err := c.BodyParser(&updatedFields); err != nil {
		return err
	}
	if _, ok := updatedFields["row_created_date"]; ok {
        return c.Status(400).SendString("Cannot update row_created_date field")
    }

	query := "update fossil_equivalence set "
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
	query += " where fossil_id = :id"

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

func DeleteFossilEquivalence(c *fiber.Ctx) error {
	id := c.Params("id")
	var fossil_equivalence dto.Fossil_equivalence
	fossil_equivalence.Fossil_id = id

	stmt, err := db.Prepare("delete from fossil_equivalence where fossil_id = :1")
	if err != nil {
		return err
	}

	rows, err := stmt.Exec(fossil_equivalence.Fossil_id)
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


