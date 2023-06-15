package apiv1

import (
	dto "github.com/DarrenMannuela/pt_gtn_bibliography/dto"
	"github.com/gofiber/fiber/v2"
	"strconv"
	"encoding/json"
)


func GetStratColumnXref(c *fiber.Ctx) error {
	rows, err := db.Query("select * from strat_column_xref")
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}
	defer rows.Close()
	var result []dto.Strat_column_xref

	for rows.Next() {
		var strat_column_xref dto.Strat_column_xref
		if err := rows.Scan(&strat_column_xref.Strat_column_id_1, &strat_column_xref.Strat_column_source_1, &strat_column_xref.Strat_column_id_2, &strat_column_xref.Strat_column_source_2, &strat_column_xref.Active_ind, &strat_column_xref.Effective_date, &strat_column_xref.Expiry_date, &strat_column_xref.Ppdm_guid, &strat_column_xref.Remark, &strat_column_xref.Source, &strat_column_xref.Strat_column_xref_type, &strat_column_xref.Row_changed_by, &strat_column_xref.Row_changed_date, &strat_column_xref.Row_created_by, &strat_column_xref.Row_created_date, &strat_column_xref.Row_effective_date, &strat_column_xref.Row_expiry_date, &strat_column_xref.Row_quality); err != nil {
			return err // Exit if we get an error
		}

		// Append strat_column_xref to result
		result = append(result, strat_column_xref)
	}
	// Return result in JSON format
	return c.JSON(result)
	 
}

func SetStratColumnXref(c *fiber.Ctx) error {
	var strat_column_xref dto.Strat_column_xref

	setDefaults(&strat_column_xref)

	if err := json.Unmarshal(c.Body(), &strat_column_xref); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	stmt, err := db.Prepare("insert into strat_column_xref values(:1, :2, :3, :4, :5, :6, :7, :8, :9, :10, :11, :12, :13, :14, :15, :16, :17, :18)")
	if err != nil {
		return err
	}
	strat_column_xref.Row_created_date = formatDate(strat_column_xref.Row_created_date)
	strat_column_xref.Row_changed_date = formatDate(strat_column_xref.Row_changed_date)
	strat_column_xref.Effective_date = formatDateString(strat_column_xref.Effective_date)
	strat_column_xref.Expiry_date = formatDateString(strat_column_xref.Expiry_date)
	strat_column_xref.Row_effective_date = formatDateString(strat_column_xref.Row_effective_date)
	strat_column_xref.Row_expiry_date = formatDateString(strat_column_xref.Row_expiry_date)






	rows, err := stmt.Exec(strat_column_xref.Strat_column_id_1, strat_column_xref.Strat_column_source_1, strat_column_xref.Strat_column_id_2, strat_column_xref.Strat_column_source_2, strat_column_xref.Active_ind, strat_column_xref.Effective_date, strat_column_xref.Expiry_date, strat_column_xref.Ppdm_guid, strat_column_xref.Remark, strat_column_xref.Source, strat_column_xref.Strat_column_xref_type, strat_column_xref.Row_changed_by, strat_column_xref.Row_changed_date, strat_column_xref.Row_created_by, strat_column_xref.Row_created_date, strat_column_xref.Row_effective_date, strat_column_xref.Row_expiry_date, strat_column_xref.Row_quality)
	
	if err != nil {
		return err
	}

	return c.JSON(rows)
}

func UpdateStratColumnXref(c *fiber.Ctx) error {
	var strat_column_xref dto.Strat_column_xref

	setDefaults(&strat_column_xref)

	if err := json.Unmarshal(c.Body(), &strat_column_xref); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	id := c.Params("id")
	strat_column_xref.Strat_column_id_1 = id

    if strat_column_xref.Row_created_date != nil {
        return c.Status(400).SendString("Cannot update row_created_date")
    }

	stmt, err := db.Prepare("update strat_column_xref set strat_column_source_1 = :1, strat_column_id_2 = :2, strat_column_source_2 = :3, active_ind = :4, effective_date = :5, expiry_date = :6, ppdm_guid = :7, remark = :8, source = :9, strat_column_xref_type = :10, row_changed_by = :11, row_changed_date = :12, row_created_by = :13, row_effective_date = :14, row_expiry_date = :15, row_quality = :16 where strat_column_id_1 = :18")
	if err != nil {
		return err
	}

	strat_column_xref.Row_changed_date = formatDate(strat_column_xref.Row_changed_date)
	strat_column_xref.Effective_date = formatDateString(strat_column_xref.Effective_date)
	strat_column_xref.Expiry_date = formatDateString(strat_column_xref.Expiry_date)
	strat_column_xref.Row_effective_date = formatDateString(strat_column_xref.Row_effective_date)
	strat_column_xref.Row_expiry_date = formatDateString(strat_column_xref.Row_expiry_date)






	rows, err := stmt.Exec(strat_column_xref.Strat_column_source_1, strat_column_xref.Strat_column_id_2, strat_column_xref.Strat_column_source_2, strat_column_xref.Active_ind, strat_column_xref.Effective_date, strat_column_xref.Expiry_date, strat_column_xref.Ppdm_guid, strat_column_xref.Remark, strat_column_xref.Source, strat_column_xref.Strat_column_xref_type, strat_column_xref.Row_changed_by, strat_column_xref.Row_changed_date, strat_column_xref.Row_created_by, strat_column_xref.Row_effective_date, strat_column_xref.Row_expiry_date, strat_column_xref.Row_quality, strat_column_xref.Strat_column_id_1)
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

func PatchStratColumnXref(c *fiber.Ctx) error {
	id := c.Params("id")

	updatedFields := make(map[string]interface{})

	if err := c.BodyParser(&updatedFields); err != nil {
		return err
	}
	if _, ok := updatedFields["row_created_date"]; ok {
        return c.Status(400).SendString("Cannot update row_created_date field")
    }

	query := "update strat_column_xref set "
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
	query += " where strat_column_id_1 = :id"

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

func DeleteStratColumnXref(c *fiber.Ctx) error {
	id := c.Params("id")
	var strat_column_xref dto.Strat_column_xref
	strat_column_xref.Strat_column_id_1 = id

	stmt, err := db.Prepare("delete from strat_column_xref where strat_column_id_1 = :1")
	if err != nil {
		return err
	}

	rows, err := stmt.Exec(strat_column_xref.Strat_column_id_1)
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


