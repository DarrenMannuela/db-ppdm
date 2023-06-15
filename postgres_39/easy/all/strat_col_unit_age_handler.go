package apiv1

import (
	dto "github.com/DarrenMannuela/pt_gtn_bibliography/dto"
	"github.com/gofiber/fiber/v2"
	"strconv"
	"encoding/json"
)


func GetStratColUnitAge(c *fiber.Ctx) error {
	rows, err := db.Query("select * from strat_col_unit_age")
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}
	defer rows.Close()
	var result []dto.Strat_col_unit_age

	for rows.Next() {
		var strat_col_unit_age dto.Strat_col_unit_age
		if err := rows.Scan(&strat_col_unit_age.Strat_column_id, &strat_col_unit_age.Strat_column_source, &strat_col_unit_age.Strat_name_set_id, &strat_col_unit_age.Strat_unit_id, &strat_col_unit_age.Interp_id, &strat_col_unit_age.Age_source, &strat_col_unit_age.Age_id, &strat_col_unit_age.Active_ind, &strat_col_unit_age.Average_age, &strat_col_unit_age.Average_age_error_minus, &strat_col_unit_age.Average_age_error_plus, &strat_col_unit_age.Average_rel_strat_name_set, &strat_col_unit_age.Average_rel_strat_unit_id, &strat_col_unit_age.Effective_date, &strat_col_unit_age.Expiry_date, &strat_col_unit_age.Lower_max_age, &strat_col_unit_age.Lower_max_age_error_minus, &strat_col_unit_age.Lower_max_age_error_plus, &strat_col_unit_age.Lower_min_age, &strat_col_unit_age.Lower_min_age_error_minus, &strat_col_unit_age.Lower_min_age_error_plus, &strat_col_unit_age.Lower_rel_strat_name_set, &strat_col_unit_age.Lower_rel_strat_unit_id, &strat_col_unit_age.Max_age, &strat_col_unit_age.Max_age_error_minus, &strat_col_unit_age.Max_age_error_plus, &strat_col_unit_age.Min_age, &strat_col_unit_age.Min_age_error_minus, &strat_col_unit_age.Min_age_error_plus, &strat_col_unit_age.Post_qualifier, &strat_col_unit_age.Ppdm_guid, &strat_col_unit_age.Preferred_ind, &strat_col_unit_age.Pre_qualifier, &strat_col_unit_age.Remark, &strat_col_unit_age.Source_document_id, &strat_col_unit_age.Strat_age_method, &strat_col_unit_age.Upper_max_age, &strat_col_unit_age.Upper_max_age_error_minus, &strat_col_unit_age.Upper_max_age_error_plus, &strat_col_unit_age.Upper_min_age, &strat_col_unit_age.Upper_min_age_error_minus, &strat_col_unit_age.Upper_min_age_error_plus, &strat_col_unit_age.Upper_rel_strat_name_set, &strat_col_unit_age.Upper_rel_strat_unit_id, &strat_col_unit_age.Row_changed_by, &strat_col_unit_age.Row_changed_date, &strat_col_unit_age.Row_created_by, &strat_col_unit_age.Row_created_date, &strat_col_unit_age.Row_effective_date, &strat_col_unit_age.Row_expiry_date, &strat_col_unit_age.Row_quality); err != nil {
			return err // Exit if we get an error
		}

		// Append strat_col_unit_age to result
		result = append(result, strat_col_unit_age)
	}
	// Return result in JSON format
	return c.JSON(result)
	 
}

func SetStratColUnitAge(c *fiber.Ctx) error {
	var strat_col_unit_age dto.Strat_col_unit_age

	setDefaults(&strat_col_unit_age)

	if err := json.Unmarshal(c.Body(), &strat_col_unit_age); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	stmt, err := db.Prepare("insert into strat_col_unit_age values(:1, :2, :3, :4, :5, :6, :7, :8, :9, :10, :11, :12, :13, :14, :15, :16, :17, :18, :19, :20, :21, :22, :23, :24, :25, :26, :27, :28, :29, :30, :31, :32, :33, :34, :35, :36, :37, :38, :39, :40, :41, :42, :43, :44, :45, :46, :47, :48, :49, :50, :51)")
	if err != nil {
		return err
	}
	strat_col_unit_age.Row_created_date = formatDate(strat_col_unit_age.Row_created_date)
	strat_col_unit_age.Row_changed_date = formatDate(strat_col_unit_age.Row_changed_date)
	strat_col_unit_age.Effective_date = formatDateString(strat_col_unit_age.Effective_date)
	strat_col_unit_age.Expiry_date = formatDateString(strat_col_unit_age.Expiry_date)
	strat_col_unit_age.Row_effective_date = formatDateString(strat_col_unit_age.Row_effective_date)
	strat_col_unit_age.Row_expiry_date = formatDateString(strat_col_unit_age.Row_expiry_date)






	rows, err := stmt.Exec(strat_col_unit_age.Strat_column_id, strat_col_unit_age.Strat_column_source, strat_col_unit_age.Strat_name_set_id, strat_col_unit_age.Strat_unit_id, strat_col_unit_age.Interp_id, strat_col_unit_age.Age_source, strat_col_unit_age.Age_id, strat_col_unit_age.Active_ind, strat_col_unit_age.Average_age, strat_col_unit_age.Average_age_error_minus, strat_col_unit_age.Average_age_error_plus, strat_col_unit_age.Average_rel_strat_name_set, strat_col_unit_age.Average_rel_strat_unit_id, strat_col_unit_age.Effective_date, strat_col_unit_age.Expiry_date, strat_col_unit_age.Lower_max_age, strat_col_unit_age.Lower_max_age_error_minus, strat_col_unit_age.Lower_max_age_error_plus, strat_col_unit_age.Lower_min_age, strat_col_unit_age.Lower_min_age_error_minus, strat_col_unit_age.Lower_min_age_error_plus, strat_col_unit_age.Lower_rel_strat_name_set, strat_col_unit_age.Lower_rel_strat_unit_id, strat_col_unit_age.Max_age, strat_col_unit_age.Max_age_error_minus, strat_col_unit_age.Max_age_error_plus, strat_col_unit_age.Min_age, strat_col_unit_age.Min_age_error_minus, strat_col_unit_age.Min_age_error_plus, strat_col_unit_age.Post_qualifier, strat_col_unit_age.Ppdm_guid, strat_col_unit_age.Preferred_ind, strat_col_unit_age.Pre_qualifier, strat_col_unit_age.Remark, strat_col_unit_age.Source_document_id, strat_col_unit_age.Strat_age_method, strat_col_unit_age.Upper_max_age, strat_col_unit_age.Upper_max_age_error_minus, strat_col_unit_age.Upper_max_age_error_plus, strat_col_unit_age.Upper_min_age, strat_col_unit_age.Upper_min_age_error_minus, strat_col_unit_age.Upper_min_age_error_plus, strat_col_unit_age.Upper_rel_strat_name_set, strat_col_unit_age.Upper_rel_strat_unit_id, strat_col_unit_age.Row_changed_by, strat_col_unit_age.Row_changed_date, strat_col_unit_age.Row_created_by, strat_col_unit_age.Row_created_date, strat_col_unit_age.Row_effective_date, strat_col_unit_age.Row_expiry_date, strat_col_unit_age.Row_quality)
	
	if err != nil {
		return err
	}

	return c.JSON(rows)
}

func UpdateStratColUnitAge(c *fiber.Ctx) error {
	var strat_col_unit_age dto.Strat_col_unit_age

	setDefaults(&strat_col_unit_age)

	if err := json.Unmarshal(c.Body(), &strat_col_unit_age); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	id := c.Params("id")
	strat_col_unit_age.Strat_column_id = id

    if strat_col_unit_age.Row_created_date != nil {
        return c.Status(400).SendString("Cannot update row_created_date")
    }

	stmt, err := db.Prepare("update strat_col_unit_age set strat_column_source = :1, strat_name_set_id = :2, strat_unit_id = :3, interp_id = :4, age_source = :5, age_id = :6, active_ind = :7, average_age = :8, average_age_error_minus = :9, average_age_error_plus = :10, average_rel_strat_name_set = :11, average_rel_strat_unit_id = :12, effective_date = :13, expiry_date = :14, lower_max_age = :15, lower_max_age_error_minus = :16, lower_max_age_error_plus = :17, lower_min_age = :18, lower_min_age_error_minus = :19, lower_min_age_error_plus = :20, lower_rel_strat_name_set = :21, lower_rel_strat_unit_id = :22, max_age = :23, max_age_error_minus = :24, max_age_error_plus = :25, min_age = :26, min_age_error_minus = :27, min_age_error_plus = :28, post_qualifier = :29, ppdm_guid = :30, preferred_ind = :31, pre_qualifier = :32, remark = :33, source_document_id = :34, strat_age_method = :35, upper_max_age = :36, upper_max_age_error_minus = :37, upper_max_age_error_plus = :38, upper_min_age = :39, upper_min_age_error_minus = :40, upper_min_age_error_plus = :41, upper_rel_strat_name_set = :42, upper_rel_strat_unit_id = :43, row_changed_by = :44, row_changed_date = :45, row_created_by = :46, row_effective_date = :47, row_expiry_date = :48, row_quality = :49 where strat_column_id = :51")
	if err != nil {
		return err
	}

	strat_col_unit_age.Row_changed_date = formatDate(strat_col_unit_age.Row_changed_date)
	strat_col_unit_age.Effective_date = formatDateString(strat_col_unit_age.Effective_date)
	strat_col_unit_age.Expiry_date = formatDateString(strat_col_unit_age.Expiry_date)
	strat_col_unit_age.Row_effective_date = formatDateString(strat_col_unit_age.Row_effective_date)
	strat_col_unit_age.Row_expiry_date = formatDateString(strat_col_unit_age.Row_expiry_date)






	rows, err := stmt.Exec(strat_col_unit_age.Strat_column_source, strat_col_unit_age.Strat_name_set_id, strat_col_unit_age.Strat_unit_id, strat_col_unit_age.Interp_id, strat_col_unit_age.Age_source, strat_col_unit_age.Age_id, strat_col_unit_age.Active_ind, strat_col_unit_age.Average_age, strat_col_unit_age.Average_age_error_minus, strat_col_unit_age.Average_age_error_plus, strat_col_unit_age.Average_rel_strat_name_set, strat_col_unit_age.Average_rel_strat_unit_id, strat_col_unit_age.Effective_date, strat_col_unit_age.Expiry_date, strat_col_unit_age.Lower_max_age, strat_col_unit_age.Lower_max_age_error_minus, strat_col_unit_age.Lower_max_age_error_plus, strat_col_unit_age.Lower_min_age, strat_col_unit_age.Lower_min_age_error_minus, strat_col_unit_age.Lower_min_age_error_plus, strat_col_unit_age.Lower_rel_strat_name_set, strat_col_unit_age.Lower_rel_strat_unit_id, strat_col_unit_age.Max_age, strat_col_unit_age.Max_age_error_minus, strat_col_unit_age.Max_age_error_plus, strat_col_unit_age.Min_age, strat_col_unit_age.Min_age_error_minus, strat_col_unit_age.Min_age_error_plus, strat_col_unit_age.Post_qualifier, strat_col_unit_age.Ppdm_guid, strat_col_unit_age.Preferred_ind, strat_col_unit_age.Pre_qualifier, strat_col_unit_age.Remark, strat_col_unit_age.Source_document_id, strat_col_unit_age.Strat_age_method, strat_col_unit_age.Upper_max_age, strat_col_unit_age.Upper_max_age_error_minus, strat_col_unit_age.Upper_max_age_error_plus, strat_col_unit_age.Upper_min_age, strat_col_unit_age.Upper_min_age_error_minus, strat_col_unit_age.Upper_min_age_error_plus, strat_col_unit_age.Upper_rel_strat_name_set, strat_col_unit_age.Upper_rel_strat_unit_id, strat_col_unit_age.Row_changed_by, strat_col_unit_age.Row_changed_date, strat_col_unit_age.Row_created_by, strat_col_unit_age.Row_effective_date, strat_col_unit_age.Row_expiry_date, strat_col_unit_age.Row_quality, strat_col_unit_age.Strat_column_id)
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

func PatchStratColUnitAge(c *fiber.Ctx) error {
	id := c.Params("id")

	updatedFields := make(map[string]interface{})

	if err := c.BodyParser(&updatedFields); err != nil {
		return err
	}
	if _, ok := updatedFields["row_created_date"]; ok {
        return c.Status(400).SendString("Cannot update row_created_date field")
    }

	query := "update strat_col_unit_age set "
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

func DeleteStratColUnitAge(c *fiber.Ctx) error {
	id := c.Params("id")
	var strat_col_unit_age dto.Strat_col_unit_age
	strat_col_unit_age.Strat_column_id = id

	stmt, err := db.Prepare("delete from strat_col_unit_age where strat_column_id = :1")
	if err != nil {
		return err
	}

	rows, err := stmt.Exec(strat_col_unit_age.Strat_column_id)
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


