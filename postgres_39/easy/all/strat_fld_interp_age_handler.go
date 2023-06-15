package apiv1

import (
	dto "github.com/DarrenMannuela/pt_gtn_bibliography/dto"
	"github.com/gofiber/fiber/v2"
	"strconv"
	"encoding/json"
)


func GetStratFldInterpAge(c *fiber.Ctx) error {
	rows, err := db.Query("select * from strat_fld_interp_age")
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}
	defer rows.Close()
	var result []dto.Strat_fld_interp_age

	for rows.Next() {
		var strat_fld_interp_age dto.Strat_fld_interp_age
		if err := rows.Scan(&strat_fld_interp_age.Field_station_id, &strat_fld_interp_age.Strat_name_set_id, &strat_fld_interp_age.Strat_unit_id, &strat_fld_interp_age.Interp_id, &strat_fld_interp_age.Age_source, &strat_fld_interp_age.Age_id, &strat_fld_interp_age.Active_ind, &strat_fld_interp_age.Average_age, &strat_fld_interp_age.Average_age_error_minus, &strat_fld_interp_age.Average_age_error_plus, &strat_fld_interp_age.Average_rel_strat_name_set, &strat_fld_interp_age.Average_rel_strat_unit_id, &strat_fld_interp_age.Effective_date, &strat_fld_interp_age.Expiry_date, &strat_fld_interp_age.Lower_max_age, &strat_fld_interp_age.Lower_max_age_error_minus, &strat_fld_interp_age.Lower_max_age_error_plus, &strat_fld_interp_age.Lower_min_age, &strat_fld_interp_age.Lower_min_age_error_minus, &strat_fld_interp_age.Lower_min_age_error_plus, &strat_fld_interp_age.Lower_rel_strat_name_set, &strat_fld_interp_age.Lower_rel_strat_unit_id, &strat_fld_interp_age.Post_qualifier, &strat_fld_interp_age.Ppdm_guid, &strat_fld_interp_age.Preferred_ind, &strat_fld_interp_age.Pre_qualifier, &strat_fld_interp_age.Remark, &strat_fld_interp_age.Source_document_id, &strat_fld_interp_age.Strat_age_method, &strat_fld_interp_age.Upper_max_age, &strat_fld_interp_age.Upper_max_age_error_minus, &strat_fld_interp_age.Upper_max_age_error_plus, &strat_fld_interp_age.Upper_min_age, &strat_fld_interp_age.Upper_min_age_error_minus, &strat_fld_interp_age.Upper_min_age_error_plus, &strat_fld_interp_age.Upper_rel_strat_name_set, &strat_fld_interp_age.Upper_rel_strat_unit_id, &strat_fld_interp_age.Row_changed_by, &strat_fld_interp_age.Row_changed_date, &strat_fld_interp_age.Row_created_by, &strat_fld_interp_age.Row_created_date, &strat_fld_interp_age.Row_effective_date, &strat_fld_interp_age.Row_expiry_date, &strat_fld_interp_age.Row_quality); err != nil {
			return err // Exit if we get an error
		}

		// Append strat_fld_interp_age to result
		result = append(result, strat_fld_interp_age)
	}
	// Return result in JSON format
	return c.JSON(result)
	 
}

func SetStratFldInterpAge(c *fiber.Ctx) error {
	var strat_fld_interp_age dto.Strat_fld_interp_age

	setDefaults(&strat_fld_interp_age)

	if err := json.Unmarshal(c.Body(), &strat_fld_interp_age); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	stmt, err := db.Prepare("insert into strat_fld_interp_age values(:1, :2, :3, :4, :5, :6, :7, :8, :9, :10, :11, :12, :13, :14, :15, :16, :17, :18, :19, :20, :21, :22, :23, :24, :25, :26, :27, :28, :29, :30, :31, :32, :33, :34, :35, :36, :37, :38, :39, :40, :41, :42, :43, :44)")
	if err != nil {
		return err
	}
	strat_fld_interp_age.Row_created_date = formatDate(strat_fld_interp_age.Row_created_date)
	strat_fld_interp_age.Row_changed_date = formatDate(strat_fld_interp_age.Row_changed_date)
	strat_fld_interp_age.Effective_date = formatDateString(strat_fld_interp_age.Effective_date)
	strat_fld_interp_age.Expiry_date = formatDateString(strat_fld_interp_age.Expiry_date)
	strat_fld_interp_age.Row_effective_date = formatDateString(strat_fld_interp_age.Row_effective_date)
	strat_fld_interp_age.Row_expiry_date = formatDateString(strat_fld_interp_age.Row_expiry_date)






	rows, err := stmt.Exec(strat_fld_interp_age.Field_station_id, strat_fld_interp_age.Strat_name_set_id, strat_fld_interp_age.Strat_unit_id, strat_fld_interp_age.Interp_id, strat_fld_interp_age.Age_source, strat_fld_interp_age.Age_id, strat_fld_interp_age.Active_ind, strat_fld_interp_age.Average_age, strat_fld_interp_age.Average_age_error_minus, strat_fld_interp_age.Average_age_error_plus, strat_fld_interp_age.Average_rel_strat_name_set, strat_fld_interp_age.Average_rel_strat_unit_id, strat_fld_interp_age.Effective_date, strat_fld_interp_age.Expiry_date, strat_fld_interp_age.Lower_max_age, strat_fld_interp_age.Lower_max_age_error_minus, strat_fld_interp_age.Lower_max_age_error_plus, strat_fld_interp_age.Lower_min_age, strat_fld_interp_age.Lower_min_age_error_minus, strat_fld_interp_age.Lower_min_age_error_plus, strat_fld_interp_age.Lower_rel_strat_name_set, strat_fld_interp_age.Lower_rel_strat_unit_id, strat_fld_interp_age.Post_qualifier, strat_fld_interp_age.Ppdm_guid, strat_fld_interp_age.Preferred_ind, strat_fld_interp_age.Pre_qualifier, strat_fld_interp_age.Remark, strat_fld_interp_age.Source_document_id, strat_fld_interp_age.Strat_age_method, strat_fld_interp_age.Upper_max_age, strat_fld_interp_age.Upper_max_age_error_minus, strat_fld_interp_age.Upper_max_age_error_plus, strat_fld_interp_age.Upper_min_age, strat_fld_interp_age.Upper_min_age_error_minus, strat_fld_interp_age.Upper_min_age_error_plus, strat_fld_interp_age.Upper_rel_strat_name_set, strat_fld_interp_age.Upper_rel_strat_unit_id, strat_fld_interp_age.Row_changed_by, strat_fld_interp_age.Row_changed_date, strat_fld_interp_age.Row_created_by, strat_fld_interp_age.Row_created_date, strat_fld_interp_age.Row_effective_date, strat_fld_interp_age.Row_expiry_date, strat_fld_interp_age.Row_quality)
	
	if err != nil {
		return err
	}

	return c.JSON(rows)
}

func UpdateStratFldInterpAge(c *fiber.Ctx) error {
	var strat_fld_interp_age dto.Strat_fld_interp_age

	setDefaults(&strat_fld_interp_age)

	if err := json.Unmarshal(c.Body(), &strat_fld_interp_age); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	id := c.Params("id")
	strat_fld_interp_age.Field_station_id = id

    if strat_fld_interp_age.Row_created_date != nil {
        return c.Status(400).SendString("Cannot update row_created_date")
    }

	stmt, err := db.Prepare("update strat_fld_interp_age set strat_name_set_id = :1, strat_unit_id = :2, interp_id = :3, age_source = :4, age_id = :5, active_ind = :6, average_age = :7, average_age_error_minus = :8, average_age_error_plus = :9, average_rel_strat_name_set = :10, average_rel_strat_unit_id = :11, effective_date = :12, expiry_date = :13, lower_max_age = :14, lower_max_age_error_minus = :15, lower_max_age_error_plus = :16, lower_min_age = :17, lower_min_age_error_minus = :18, lower_min_age_error_plus = :19, lower_rel_strat_name_set = :20, lower_rel_strat_unit_id = :21, post_qualifier = :22, ppdm_guid = :23, preferred_ind = :24, pre_qualifier = :25, remark = :26, source_document_id = :27, strat_age_method = :28, upper_max_age = :29, upper_max_age_error_minus = :30, upper_max_age_error_plus = :31, upper_min_age = :32, upper_min_age_error_minus = :33, upper_min_age_error_plus = :34, upper_rel_strat_name_set = :35, upper_rel_strat_unit_id = :36, row_changed_by = :37, row_changed_date = :38, row_created_by = :39, row_effective_date = :40, row_expiry_date = :41, row_quality = :42 where field_station_id = :44")
	if err != nil {
		return err
	}

	strat_fld_interp_age.Row_changed_date = formatDate(strat_fld_interp_age.Row_changed_date)
	strat_fld_interp_age.Effective_date = formatDateString(strat_fld_interp_age.Effective_date)
	strat_fld_interp_age.Expiry_date = formatDateString(strat_fld_interp_age.Expiry_date)
	strat_fld_interp_age.Row_effective_date = formatDateString(strat_fld_interp_age.Row_effective_date)
	strat_fld_interp_age.Row_expiry_date = formatDateString(strat_fld_interp_age.Row_expiry_date)






	rows, err := stmt.Exec(strat_fld_interp_age.Strat_name_set_id, strat_fld_interp_age.Strat_unit_id, strat_fld_interp_age.Interp_id, strat_fld_interp_age.Age_source, strat_fld_interp_age.Age_id, strat_fld_interp_age.Active_ind, strat_fld_interp_age.Average_age, strat_fld_interp_age.Average_age_error_minus, strat_fld_interp_age.Average_age_error_plus, strat_fld_interp_age.Average_rel_strat_name_set, strat_fld_interp_age.Average_rel_strat_unit_id, strat_fld_interp_age.Effective_date, strat_fld_interp_age.Expiry_date, strat_fld_interp_age.Lower_max_age, strat_fld_interp_age.Lower_max_age_error_minus, strat_fld_interp_age.Lower_max_age_error_plus, strat_fld_interp_age.Lower_min_age, strat_fld_interp_age.Lower_min_age_error_minus, strat_fld_interp_age.Lower_min_age_error_plus, strat_fld_interp_age.Lower_rel_strat_name_set, strat_fld_interp_age.Lower_rel_strat_unit_id, strat_fld_interp_age.Post_qualifier, strat_fld_interp_age.Ppdm_guid, strat_fld_interp_age.Preferred_ind, strat_fld_interp_age.Pre_qualifier, strat_fld_interp_age.Remark, strat_fld_interp_age.Source_document_id, strat_fld_interp_age.Strat_age_method, strat_fld_interp_age.Upper_max_age, strat_fld_interp_age.Upper_max_age_error_minus, strat_fld_interp_age.Upper_max_age_error_plus, strat_fld_interp_age.Upper_min_age, strat_fld_interp_age.Upper_min_age_error_minus, strat_fld_interp_age.Upper_min_age_error_plus, strat_fld_interp_age.Upper_rel_strat_name_set, strat_fld_interp_age.Upper_rel_strat_unit_id, strat_fld_interp_age.Row_changed_by, strat_fld_interp_age.Row_changed_date, strat_fld_interp_age.Row_created_by, strat_fld_interp_age.Row_effective_date, strat_fld_interp_age.Row_expiry_date, strat_fld_interp_age.Row_quality, strat_fld_interp_age.Field_station_id)
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

func PatchStratFldInterpAge(c *fiber.Ctx) error {
	id := c.Params("id")

	updatedFields := make(map[string]interface{})

	if err := c.BodyParser(&updatedFields); err != nil {
		return err
	}
	if _, ok := updatedFields["row_created_date"]; ok {
        return c.Status(400).SendString("Cannot update row_created_date field")
    }

	query := "update strat_fld_interp_age set "
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
	query += " where field_station_id = :id"

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

func DeleteStratFldInterpAge(c *fiber.Ctx) error {
	id := c.Params("id")
	var strat_fld_interp_age dto.Strat_fld_interp_age
	strat_fld_interp_age.Field_station_id = id

	stmt, err := db.Prepare("delete from strat_fld_interp_age where field_station_id = :1")
	if err != nil {
		return err
	}

	rows, err := stmt.Exec(strat_fld_interp_age.Field_station_id)
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


