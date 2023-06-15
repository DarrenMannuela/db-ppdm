package apiv1

import (
	dto "github.com/DarrenMannuela/pt_gtn_bibliography/dto"
	"github.com/gofiber/fiber/v2"
	"strconv"
	"encoding/json"
)


func GetStratUnitAge(c *fiber.Ctx) error {
	rows, err := db.Query("select * from strat_unit_age")
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}
	defer rows.Close()
	var result []dto.Strat_unit_age

	for rows.Next() {
		var strat_unit_age dto.Strat_unit_age
		if err := rows.Scan(&strat_unit_age.Strat_name_set_id, &strat_unit_age.Strat_unit_id, &strat_unit_age.Age_seq_no, &strat_unit_age.Active_ind, &strat_unit_age.Average_age, &strat_unit_age.Average_age_error_minus, &strat_unit_age.Average_age_error_plus, &strat_unit_age.Average_rel_strat_name_set, &strat_unit_age.Average_rel_strat_unit_id, &strat_unit_age.Effective_date, &strat_unit_age.Expiry_date, &strat_unit_age.Lower_max_age, &strat_unit_age.Lower_max_age_error_minus, &strat_unit_age.Lower_max_age_error_plus, &strat_unit_age.Lower_min_age, &strat_unit_age.Lower_min_age_error_minus, &strat_unit_age.Lower_min_age_error_plus, &strat_unit_age.Lower_rel_strat_name_set, &strat_unit_age.Lower_rel_strat_unit_id, &strat_unit_age.Max_age, &strat_unit_age.Max_age_error_minus, &strat_unit_age.Max_age_error_plus, &strat_unit_age.Min_age, &strat_unit_age.Min_age_error_minus, &strat_unit_age.Min_age_error_plus, &strat_unit_age.Post_qualifier, &strat_unit_age.Ppdm_guid, &strat_unit_age.Preferred_ind, &strat_unit_age.Pre_qualifier, &strat_unit_age.Remark, &strat_unit_age.Source, &strat_unit_age.Source_document_id, &strat_unit_age.Strat_age_method, &strat_unit_age.Upper_max_age, &strat_unit_age.Upper_max_age_error_minus, &strat_unit_age.Upper_max_age_error_plus, &strat_unit_age.Upper_min_age, &strat_unit_age.Upper_min_age_error_minus, &strat_unit_age.Upper_min_age_error_plus, &strat_unit_age.Upper_rel_strat_name_set, &strat_unit_age.Upper_rel_strat_unit_id, &strat_unit_age.Row_changed_by, &strat_unit_age.Row_changed_date, &strat_unit_age.Row_created_by, &strat_unit_age.Row_created_date, &strat_unit_age.Row_effective_date, &strat_unit_age.Row_expiry_date, &strat_unit_age.Row_quality); err != nil {
			return err // Exit if we get an error
		}

		// Append strat_unit_age to result
		result = append(result, strat_unit_age)
	}
	// Return result in JSON format
	return c.JSON(result)
	 
}

func SetStratUnitAge(c *fiber.Ctx) error {
	var strat_unit_age dto.Strat_unit_age

	setDefaults(&strat_unit_age)

	if err := json.Unmarshal(c.Body(), &strat_unit_age); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	stmt, err := db.Prepare("insert into strat_unit_age values(:1, :2, :3, :4, :5, :6, :7, :8, :9, :10, :11, :12, :13, :14, :15, :16, :17, :18, :19, :20, :21, :22, :23, :24, :25, :26, :27, :28, :29, :30, :31, :32, :33, :34, :35, :36, :37, :38, :39, :40, :41, :42, :43, :44, :45, :46, :47, :48)")
	if err != nil {
		return err
	}
	strat_unit_age.Row_created_date = formatDate(strat_unit_age.Row_created_date)
	strat_unit_age.Row_changed_date = formatDate(strat_unit_age.Row_changed_date)
	strat_unit_age.Effective_date = formatDateString(strat_unit_age.Effective_date)
	strat_unit_age.Expiry_date = formatDateString(strat_unit_age.Expiry_date)
	strat_unit_age.Row_effective_date = formatDateString(strat_unit_age.Row_effective_date)
	strat_unit_age.Row_expiry_date = formatDateString(strat_unit_age.Row_expiry_date)






	rows, err := stmt.Exec(strat_unit_age.Strat_name_set_id, strat_unit_age.Strat_unit_id, strat_unit_age.Age_seq_no, strat_unit_age.Active_ind, strat_unit_age.Average_age, strat_unit_age.Average_age_error_minus, strat_unit_age.Average_age_error_plus, strat_unit_age.Average_rel_strat_name_set, strat_unit_age.Average_rel_strat_unit_id, strat_unit_age.Effective_date, strat_unit_age.Expiry_date, strat_unit_age.Lower_max_age, strat_unit_age.Lower_max_age_error_minus, strat_unit_age.Lower_max_age_error_plus, strat_unit_age.Lower_min_age, strat_unit_age.Lower_min_age_error_minus, strat_unit_age.Lower_min_age_error_plus, strat_unit_age.Lower_rel_strat_name_set, strat_unit_age.Lower_rel_strat_unit_id, strat_unit_age.Max_age, strat_unit_age.Max_age_error_minus, strat_unit_age.Max_age_error_plus, strat_unit_age.Min_age, strat_unit_age.Min_age_error_minus, strat_unit_age.Min_age_error_plus, strat_unit_age.Post_qualifier, strat_unit_age.Ppdm_guid, strat_unit_age.Preferred_ind, strat_unit_age.Pre_qualifier, strat_unit_age.Remark, strat_unit_age.Source, strat_unit_age.Source_document_id, strat_unit_age.Strat_age_method, strat_unit_age.Upper_max_age, strat_unit_age.Upper_max_age_error_minus, strat_unit_age.Upper_max_age_error_plus, strat_unit_age.Upper_min_age, strat_unit_age.Upper_min_age_error_minus, strat_unit_age.Upper_min_age_error_plus, strat_unit_age.Upper_rel_strat_name_set, strat_unit_age.Upper_rel_strat_unit_id, strat_unit_age.Row_changed_by, strat_unit_age.Row_changed_date, strat_unit_age.Row_created_by, strat_unit_age.Row_created_date, strat_unit_age.Row_effective_date, strat_unit_age.Row_expiry_date, strat_unit_age.Row_quality)
	
	if err != nil {
		return err
	}

	return c.JSON(rows)
}

func UpdateStratUnitAge(c *fiber.Ctx) error {
	var strat_unit_age dto.Strat_unit_age

	setDefaults(&strat_unit_age)

	if err := json.Unmarshal(c.Body(), &strat_unit_age); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	id := c.Params("id")
	strat_unit_age.Strat_name_set_id = id

    if strat_unit_age.Row_created_date != nil {
        return c.Status(400).SendString("Cannot update row_created_date")
    }

	stmt, err := db.Prepare("update strat_unit_age set strat_unit_id = :1, age_seq_no = :2, active_ind = :3, average_age = :4, average_age_error_minus = :5, average_age_error_plus = :6, average_rel_strat_name_set = :7, average_rel_strat_unit_id = :8, effective_date = :9, expiry_date = :10, lower_max_age = :11, lower_max_age_error_minus = :12, lower_max_age_error_plus = :13, lower_min_age = :14, lower_min_age_error_minus = :15, lower_min_age_error_plus = :16, lower_rel_strat_name_set = :17, lower_rel_strat_unit_id = :18, max_age = :19, max_age_error_minus = :20, max_age_error_plus = :21, min_age = :22, min_age_error_minus = :23, min_age_error_plus = :24, post_qualifier = :25, ppdm_guid = :26, preferred_ind = :27, pre_qualifier = :28, remark = :29, source = :30, source_document_id = :31, strat_age_method = :32, upper_max_age = :33, upper_max_age_error_minus = :34, upper_max_age_error_plus = :35, upper_min_age = :36, upper_min_age_error_minus = :37, upper_min_age_error_plus = :38, upper_rel_strat_name_set = :39, upper_rel_strat_unit_id = :40, row_changed_by = :41, row_changed_date = :42, row_created_by = :43, row_effective_date = :44, row_expiry_date = :45, row_quality = :46 where strat_name_set_id = :48")
	if err != nil {
		return err
	}

	strat_unit_age.Row_changed_date = formatDate(strat_unit_age.Row_changed_date)
	strat_unit_age.Effective_date = formatDateString(strat_unit_age.Effective_date)
	strat_unit_age.Expiry_date = formatDateString(strat_unit_age.Expiry_date)
	strat_unit_age.Row_effective_date = formatDateString(strat_unit_age.Row_effective_date)
	strat_unit_age.Row_expiry_date = formatDateString(strat_unit_age.Row_expiry_date)






	rows, err := stmt.Exec(strat_unit_age.Strat_unit_id, strat_unit_age.Age_seq_no, strat_unit_age.Active_ind, strat_unit_age.Average_age, strat_unit_age.Average_age_error_minus, strat_unit_age.Average_age_error_plus, strat_unit_age.Average_rel_strat_name_set, strat_unit_age.Average_rel_strat_unit_id, strat_unit_age.Effective_date, strat_unit_age.Expiry_date, strat_unit_age.Lower_max_age, strat_unit_age.Lower_max_age_error_minus, strat_unit_age.Lower_max_age_error_plus, strat_unit_age.Lower_min_age, strat_unit_age.Lower_min_age_error_minus, strat_unit_age.Lower_min_age_error_plus, strat_unit_age.Lower_rel_strat_name_set, strat_unit_age.Lower_rel_strat_unit_id, strat_unit_age.Max_age, strat_unit_age.Max_age_error_minus, strat_unit_age.Max_age_error_plus, strat_unit_age.Min_age, strat_unit_age.Min_age_error_minus, strat_unit_age.Min_age_error_plus, strat_unit_age.Post_qualifier, strat_unit_age.Ppdm_guid, strat_unit_age.Preferred_ind, strat_unit_age.Pre_qualifier, strat_unit_age.Remark, strat_unit_age.Source, strat_unit_age.Source_document_id, strat_unit_age.Strat_age_method, strat_unit_age.Upper_max_age, strat_unit_age.Upper_max_age_error_minus, strat_unit_age.Upper_max_age_error_plus, strat_unit_age.Upper_min_age, strat_unit_age.Upper_min_age_error_minus, strat_unit_age.Upper_min_age_error_plus, strat_unit_age.Upper_rel_strat_name_set, strat_unit_age.Upper_rel_strat_unit_id, strat_unit_age.Row_changed_by, strat_unit_age.Row_changed_date, strat_unit_age.Row_created_by, strat_unit_age.Row_effective_date, strat_unit_age.Row_expiry_date, strat_unit_age.Row_quality, strat_unit_age.Strat_name_set_id)
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

func PatchStratUnitAge(c *fiber.Ctx) error {
	id := c.Params("id")

	updatedFields := make(map[string]interface{})

	if err := c.BodyParser(&updatedFields); err != nil {
		return err
	}
	if _, ok := updatedFields["row_created_date"]; ok {
        return c.Status(400).SendString("Cannot update row_created_date field")
    }

	query := "update strat_unit_age set "
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
	query += " where strat_name_set_id = :id"

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

func DeleteStratUnitAge(c *fiber.Ctx) error {
	id := c.Params("id")
	var strat_unit_age dto.Strat_unit_age
	strat_unit_age.Strat_name_set_id = id

	stmt, err := db.Prepare("delete from strat_unit_age where strat_name_set_id = :1")
	if err != nil {
		return err
	}

	rows, err := stmt.Exec(strat_unit_age.Strat_name_set_id)
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


