package apiv1

import (
	dto "github.com/DarrenMannuela/pt_gtn_bibliography/dto"
	"github.com/gofiber/fiber/v2"
	"strconv"
	"encoding/json"
)


func GetStratAcqtnMethod(c *fiber.Ctx) error {
	rows, err := db.Query("select * from strat_acqtn_method")
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}
	defer rows.Close()
	var result []dto.Strat_acqtn_method

	for rows.Next() {
		var strat_acqtn_method dto.Strat_acqtn_method
		if err := rows.Scan(&strat_acqtn_method.Strat_acqtn_method_id, &strat_acqtn_method.Acqtn_date, &strat_acqtn_method.Acqtn_method_type, &strat_acqtn_method.Active_ind, &strat_acqtn_method.Effective_date, &strat_acqtn_method.Expiry_date, &strat_acqtn_method.Field_station_id, &strat_acqtn_method.Interp_id, &strat_acqtn_method.Ppdm_guid, &strat_acqtn_method.Remark, &strat_acqtn_method.Source, &strat_acqtn_method.Strat_column_id, &strat_acqtn_method.Strat_column_source, &strat_acqtn_method.Strat_name_set_id, &strat_acqtn_method.Strat_unit_id, &strat_acqtn_method.Uwi, &strat_acqtn_method.Row_changed_by, &strat_acqtn_method.Row_changed_date, &strat_acqtn_method.Row_created_by, &strat_acqtn_method.Row_created_date, &strat_acqtn_method.Row_effective_date, &strat_acqtn_method.Row_expiry_date, &strat_acqtn_method.Row_quality); err != nil {
			return err // Exit if we get an error
		}

		// Append strat_acqtn_method to result
		result = append(result, strat_acqtn_method)
	}
	// Return result in JSON format
	return c.JSON(result)
	 
}

func SetStratAcqtnMethod(c *fiber.Ctx) error {
	var strat_acqtn_method dto.Strat_acqtn_method

	setDefaults(&strat_acqtn_method)

	if err := json.Unmarshal(c.Body(), &strat_acqtn_method); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	stmt, err := db.Prepare("insert into strat_acqtn_method values(:1, :2, :3, :4, :5, :6, :7, :8, :9, :10, :11, :12, :13, :14, :15, :16, :17, :18, :19, :20, :21, :22, :23)")
	if err != nil {
		return err
	}
	strat_acqtn_method.Row_created_date = formatDate(strat_acqtn_method.Row_created_date)
	strat_acqtn_method.Row_changed_date = formatDate(strat_acqtn_method.Row_changed_date)
	strat_acqtn_method.Acqtn_date = formatDateString(strat_acqtn_method.Acqtn_date)
	strat_acqtn_method.Effective_date = formatDateString(strat_acqtn_method.Effective_date)
	strat_acqtn_method.Expiry_date = formatDateString(strat_acqtn_method.Expiry_date)
	strat_acqtn_method.Row_effective_date = formatDateString(strat_acqtn_method.Row_effective_date)
	strat_acqtn_method.Row_expiry_date = formatDateString(strat_acqtn_method.Row_expiry_date)






	rows, err := stmt.Exec(strat_acqtn_method.Strat_acqtn_method_id, strat_acqtn_method.Acqtn_date, strat_acqtn_method.Acqtn_method_type, strat_acqtn_method.Active_ind, strat_acqtn_method.Effective_date, strat_acqtn_method.Expiry_date, strat_acqtn_method.Field_station_id, strat_acqtn_method.Interp_id, strat_acqtn_method.Ppdm_guid, strat_acqtn_method.Remark, strat_acqtn_method.Source, strat_acqtn_method.Strat_column_id, strat_acqtn_method.Strat_column_source, strat_acqtn_method.Strat_name_set_id, strat_acqtn_method.Strat_unit_id, strat_acqtn_method.Uwi, strat_acqtn_method.Row_changed_by, strat_acqtn_method.Row_changed_date, strat_acqtn_method.Row_created_by, strat_acqtn_method.Row_created_date, strat_acqtn_method.Row_effective_date, strat_acqtn_method.Row_expiry_date, strat_acqtn_method.Row_quality)
	
	if err != nil {
		return err
	}

	return c.JSON(rows)
}

func UpdateStratAcqtnMethod(c *fiber.Ctx) error {
	var strat_acqtn_method dto.Strat_acqtn_method

	setDefaults(&strat_acqtn_method)

	if err := json.Unmarshal(c.Body(), &strat_acqtn_method); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	id := c.Params("id")
	strat_acqtn_method.Strat_acqtn_method_id = id

    if strat_acqtn_method.Row_created_date != nil {
        return c.Status(400).SendString("Cannot update row_created_date")
    }

	stmt, err := db.Prepare("update strat_acqtn_method set acqtn_date = :1, acqtn_method_type = :2, active_ind = :3, effective_date = :4, expiry_date = :5, field_station_id = :6, interp_id = :7, ppdm_guid = :8, remark = :9, source = :10, strat_column_id = :11, strat_column_source = :12, strat_name_set_id = :13, strat_unit_id = :14, uwi = :15, row_changed_by = :16, row_changed_date = :17, row_created_by = :18, row_effective_date = :19, row_expiry_date = :20, row_quality = :21 where strat_acqtn_method_id = :23")
	if err != nil {
		return err
	}

	strat_acqtn_method.Row_changed_date = formatDate(strat_acqtn_method.Row_changed_date)
	strat_acqtn_method.Acqtn_date = formatDateString(strat_acqtn_method.Acqtn_date)
	strat_acqtn_method.Effective_date = formatDateString(strat_acqtn_method.Effective_date)
	strat_acqtn_method.Expiry_date = formatDateString(strat_acqtn_method.Expiry_date)
	strat_acqtn_method.Row_effective_date = formatDateString(strat_acqtn_method.Row_effective_date)
	strat_acqtn_method.Row_expiry_date = formatDateString(strat_acqtn_method.Row_expiry_date)






	rows, err := stmt.Exec(strat_acqtn_method.Acqtn_date, strat_acqtn_method.Acqtn_method_type, strat_acqtn_method.Active_ind, strat_acqtn_method.Effective_date, strat_acqtn_method.Expiry_date, strat_acqtn_method.Field_station_id, strat_acqtn_method.Interp_id, strat_acqtn_method.Ppdm_guid, strat_acqtn_method.Remark, strat_acqtn_method.Source, strat_acqtn_method.Strat_column_id, strat_acqtn_method.Strat_column_source, strat_acqtn_method.Strat_name_set_id, strat_acqtn_method.Strat_unit_id, strat_acqtn_method.Uwi, strat_acqtn_method.Row_changed_by, strat_acqtn_method.Row_changed_date, strat_acqtn_method.Row_created_by, strat_acqtn_method.Row_effective_date, strat_acqtn_method.Row_expiry_date, strat_acqtn_method.Row_quality, strat_acqtn_method.Strat_acqtn_method_id)
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

func PatchStratAcqtnMethod(c *fiber.Ctx) error {
	id := c.Params("id")

	updatedFields := make(map[string]interface{})

	if err := c.BodyParser(&updatedFields); err != nil {
		return err
	}
	if _, ok := updatedFields["row_created_date"]; ok {
        return c.Status(400).SendString("Cannot update row_created_date field")
    }

	query := "update strat_acqtn_method set "
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
		} else if key == "acqtn_date" || key == "effective_date" || key == "expiry_date" || key == "row_effective_date" || key == "row_expiry_date"      {
			if date, ok := value.(string); ok && len(date) > 0 {
				formattedDate := formatDateString(&date)
				value = formattedDate
			}
		}
		values = append(values, value)
	}
	query += " where strat_acqtn_method_id = :id"

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

func DeleteStratAcqtnMethod(c *fiber.Ctx) error {
	id := c.Params("id")
	var strat_acqtn_method dto.Strat_acqtn_method
	strat_acqtn_method.Strat_acqtn_method_id = id

	stmt, err := db.Prepare("delete from strat_acqtn_method where strat_acqtn_method_id = :1")
	if err != nil {
		return err
	}

	rows, err := stmt.Exec(strat_acqtn_method.Strat_acqtn_method_id)
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


