package apiv1

import (
	dto "github.com/DarrenMannuela/pt_gtn_bibliography/dto"
	"github.com/gofiber/fiber/v2"
	"strconv"
	"encoding/json"
)


func GetStratColumn(c *fiber.Ctx) error {
	rows, err := db.Query("select * from strat_column")
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}
	defer rows.Close()
	var result []dto.Strat_column

	for rows.Next() {
		var strat_column dto.Strat_column
		if err := rows.Scan(&strat_column.Strat_column_id, &strat_column.Source, &strat_column.Active_ind, &strat_column.Area_id, &strat_column.Area_type, &strat_column.Business_associate_id, &strat_column.Effective_date, &strat_column.Expiry_date, &strat_column.Ppdm_guid, &strat_column.Remark, &strat_column.Source_document_id, &strat_column.Strat_column_name, &strat_column.Strat_column_type, &strat_column.Row_changed_by, &strat_column.Row_changed_date, &strat_column.Row_created_by, &strat_column.Row_created_date, &strat_column.Row_effective_date, &strat_column.Row_expiry_date, &strat_column.Row_quality); err != nil {
			return err // Exit if we get an error
		}

		// Append strat_column to result
		result = append(result, strat_column)
	}
	// Return result in JSON format
	return c.JSON(result)
	 
}

func SetStratColumn(c *fiber.Ctx) error {
	var strat_column dto.Strat_column

	setDefaults(&strat_column)

	if err := json.Unmarshal(c.Body(), &strat_column); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	stmt, err := db.Prepare("insert into strat_column values(:1, :2, :3, :4, :5, :6, :7, :8, :9, :10, :11, :12, :13, :14, :15, :16, :17, :18, :19, :20)")
	if err != nil {
		return err
	}
	strat_column.Row_created_date = formatDate(strat_column.Row_created_date)
	strat_column.Row_changed_date = formatDate(strat_column.Row_changed_date)
	strat_column.Effective_date = formatDateString(strat_column.Effective_date)
	strat_column.Expiry_date = formatDateString(strat_column.Expiry_date)
	strat_column.Row_effective_date = formatDateString(strat_column.Row_effective_date)
	strat_column.Row_expiry_date = formatDateString(strat_column.Row_expiry_date)






	rows, err := stmt.Exec(strat_column.Strat_column_id, strat_column.Source, strat_column.Active_ind, strat_column.Area_id, strat_column.Area_type, strat_column.Business_associate_id, strat_column.Effective_date, strat_column.Expiry_date, strat_column.Ppdm_guid, strat_column.Remark, strat_column.Source_document_id, strat_column.Strat_column_name, strat_column.Strat_column_type, strat_column.Row_changed_by, strat_column.Row_changed_date, strat_column.Row_created_by, strat_column.Row_created_date, strat_column.Row_effective_date, strat_column.Row_expiry_date, strat_column.Row_quality)
	
	if err != nil {
		return err
	}

	return c.JSON(rows)
}

func UpdateStratColumn(c *fiber.Ctx) error {
	var strat_column dto.Strat_column

	setDefaults(&strat_column)

	if err := json.Unmarshal(c.Body(), &strat_column); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	id := c.Params("id")
	strat_column.Strat_column_id = id

    if strat_column.Row_created_date != nil {
        return c.Status(400).SendString("Cannot update row_created_date")
    }

	stmt, err := db.Prepare("update strat_column set source = :1, active_ind = :2, area_id = :3, area_type = :4, business_associate_id = :5, effective_date = :6, expiry_date = :7, ppdm_guid = :8, remark = :9, source_document_id = :10, strat_column_name = :11, strat_column_type = :12, row_changed_by = :13, row_changed_date = :14, row_created_by = :15, row_effective_date = :16, row_expiry_date = :17, row_quality = :18 where strat_column_id = :20")
	if err != nil {
		return err
	}

	strat_column.Row_changed_date = formatDate(strat_column.Row_changed_date)
	strat_column.Effective_date = formatDateString(strat_column.Effective_date)
	strat_column.Expiry_date = formatDateString(strat_column.Expiry_date)
	strat_column.Row_effective_date = formatDateString(strat_column.Row_effective_date)
	strat_column.Row_expiry_date = formatDateString(strat_column.Row_expiry_date)






	rows, err := stmt.Exec(strat_column.Source, strat_column.Active_ind, strat_column.Area_id, strat_column.Area_type, strat_column.Business_associate_id, strat_column.Effective_date, strat_column.Expiry_date, strat_column.Ppdm_guid, strat_column.Remark, strat_column.Source_document_id, strat_column.Strat_column_name, strat_column.Strat_column_type, strat_column.Row_changed_by, strat_column.Row_changed_date, strat_column.Row_created_by, strat_column.Row_effective_date, strat_column.Row_expiry_date, strat_column.Row_quality, strat_column.Strat_column_id)
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

func PatchStratColumn(c *fiber.Ctx) error {
	id := c.Params("id")

	updatedFields := make(map[string]interface{})

	if err := c.BodyParser(&updatedFields); err != nil {
		return err
	}
	if _, ok := updatedFields["row_created_date"]; ok {
        return c.Status(400).SendString("Cannot update row_created_date field")
    }

	query := "update strat_column set "
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
	query += " where strat_column_id = :id"

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

func DeleteStratColumn(c *fiber.Ctx) error {
	id := c.Params("id")
	var strat_column dto.Strat_column
	strat_column.Strat_column_id = id

	stmt, err := db.Prepare("delete from strat_column where strat_column_id = :1")
	if err != nil {
		return err
	}

	rows, err := stmt.Exec(strat_column.Strat_column_id)
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


