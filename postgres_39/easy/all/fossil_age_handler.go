package apiv1

import (
	dto "github.com/DarrenMannuela/pt_gtn_bibliography/dto"
	"github.com/gofiber/fiber/v2"
	"strconv"
	"encoding/json"
)


func GetFossilAge(c *fiber.Ctx) error {
	rows, err := db.Query("select * from fossil_age")
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}
	defer rows.Close()
	var result []dto.Fossil_age

	for rows.Next() {
		var fossil_age dto.Fossil_age
		if err := rows.Scan(&fossil_age.Fossil_id, &fossil_age.Age_interp_id, &fossil_age.Active_ind, &fossil_age.Age_confidence_id, &fossil_age.Average_age, &fossil_age.Average_age_error_minus, &fossil_age.Average_age_error_plus, &fossil_age.Average_rel_strat_name_set, &fossil_age.Average_rel_strat_unit_id, &fossil_age.Effective_date, &fossil_age.Expiry_date, &fossil_age.Lower_max_age, &fossil_age.Lower_max_age_error_minus, &fossil_age.Lower_max_age_error_plus, &fossil_age.Lower_min_age, &fossil_age.Lower_min_age_error_minus, &fossil_age.Lower_min_age_error_plus, &fossil_age.Lower_rel_strat_name_set, &fossil_age.Lower_rel_strat_unit_id, &fossil_age.Max_age, &fossil_age.Max_age_error_minus, &fossil_age.Max_age_error_plus, &fossil_age.Min_age, &fossil_age.Min_age_error_minus, &fossil_age.Min_age_error_plus, &fossil_age.Post_qualifier, &fossil_age.Ppdm_guid, &fossil_age.Preferred_ind, &fossil_age.Pre_qualifier, &fossil_age.Remark, &fossil_age.Source, &fossil_age.Source_document_id, &fossil_age.Upper_max_age, &fossil_age.Upper_max_age_error_minus, &fossil_age.Upper_max_age_error_plus, &fossil_age.Upper_min_age, &fossil_age.Upper_min_age_error_minus, &fossil_age.Upper_min_age_error_plus, &fossil_age.Upper_rel_strat_name_set, &fossil_age.Upper_rel_strat_unit_id, &fossil_age.Row_changed_by, &fossil_age.Row_changed_date, &fossil_age.Row_created_by, &fossil_age.Row_created_date, &fossil_age.Row_effective_date, &fossil_age.Row_expiry_date, &fossil_age.Row_quality); err != nil {
			return err // Exit if we get an error
		}

		// Append fossil_age to result
		result = append(result, fossil_age)
	}
	// Return result in JSON format
	return c.JSON(result)
	 
}

func SetFossilAge(c *fiber.Ctx) error {
	var fossil_age dto.Fossil_age

	setDefaults(&fossil_age)

	if err := json.Unmarshal(c.Body(), &fossil_age); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	stmt, err := db.Prepare("insert into fossil_age values(:1, :2, :3, :4, :5, :6, :7, :8, :9, :10, :11, :12, :13, :14, :15, :16, :17, :18, :19, :20, :21, :22, :23, :24, :25, :26, :27, :28, :29, :30, :31, :32, :33, :34, :35, :36, :37, :38, :39, :40, :41, :42, :43, :44, :45, :46, :47)")
	if err != nil {
		return err
	}
	fossil_age.Row_created_date = formatDate(fossil_age.Row_created_date)
	fossil_age.Row_changed_date = formatDate(fossil_age.Row_changed_date)
	fossil_age.Effective_date = formatDateString(fossil_age.Effective_date)
	fossil_age.Expiry_date = formatDateString(fossil_age.Expiry_date)
	fossil_age.Row_effective_date = formatDateString(fossil_age.Row_effective_date)
	fossil_age.Row_expiry_date = formatDateString(fossil_age.Row_expiry_date)






	rows, err := stmt.Exec(fossil_age.Fossil_id, fossil_age.Age_interp_id, fossil_age.Active_ind, fossil_age.Age_confidence_id, fossil_age.Average_age, fossil_age.Average_age_error_minus, fossil_age.Average_age_error_plus, fossil_age.Average_rel_strat_name_set, fossil_age.Average_rel_strat_unit_id, fossil_age.Effective_date, fossil_age.Expiry_date, fossil_age.Lower_max_age, fossil_age.Lower_max_age_error_minus, fossil_age.Lower_max_age_error_plus, fossil_age.Lower_min_age, fossil_age.Lower_min_age_error_minus, fossil_age.Lower_min_age_error_plus, fossil_age.Lower_rel_strat_name_set, fossil_age.Lower_rel_strat_unit_id, fossil_age.Max_age, fossil_age.Max_age_error_minus, fossil_age.Max_age_error_plus, fossil_age.Min_age, fossil_age.Min_age_error_minus, fossil_age.Min_age_error_plus, fossil_age.Post_qualifier, fossil_age.Ppdm_guid, fossil_age.Preferred_ind, fossil_age.Pre_qualifier, fossil_age.Remark, fossil_age.Source, fossil_age.Source_document_id, fossil_age.Upper_max_age, fossil_age.Upper_max_age_error_minus, fossil_age.Upper_max_age_error_plus, fossil_age.Upper_min_age, fossil_age.Upper_min_age_error_minus, fossil_age.Upper_min_age_error_plus, fossil_age.Upper_rel_strat_name_set, fossil_age.Upper_rel_strat_unit_id, fossil_age.Row_changed_by, fossil_age.Row_changed_date, fossil_age.Row_created_by, fossil_age.Row_created_date, fossil_age.Row_effective_date, fossil_age.Row_expiry_date, fossil_age.Row_quality)
	
	if err != nil {
		return err
	}

	return c.JSON(rows)
}

func UpdateFossilAge(c *fiber.Ctx) error {
	var fossil_age dto.Fossil_age

	setDefaults(&fossil_age)

	if err := json.Unmarshal(c.Body(), &fossil_age); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	id := c.Params("id")
	fossil_age.Fossil_id = id

    if fossil_age.Row_created_date != nil {
        return c.Status(400).SendString("Cannot update row_created_date")
    }

	stmt, err := db.Prepare("update fossil_age set age_interp_id = :1, active_ind = :2, age_confidence_id = :3, average_age = :4, average_age_error_minus = :5, average_age_error_plus = :6, average_rel_strat_name_set = :7, average_rel_strat_unit_id = :8, effective_date = :9, expiry_date = :10, lower_max_age = :11, lower_max_age_error_minus = :12, lower_max_age_error_plus = :13, lower_min_age = :14, lower_min_age_error_minus = :15, lower_min_age_error_plus = :16, lower_rel_strat_name_set = :17, lower_rel_strat_unit_id = :18, max_age = :19, max_age_error_minus = :20, max_age_error_plus = :21, min_age = :22, min_age_error_minus = :23, min_age_error_plus = :24, post_qualifier = :25, ppdm_guid = :26, preferred_ind = :27, pre_qualifier = :28, remark = :29, source = :30, source_document_id = :31, upper_max_age = :32, upper_max_age_error_minus = :33, upper_max_age_error_plus = :34, upper_min_age = :35, upper_min_age_error_minus = :36, upper_min_age_error_plus = :37, upper_rel_strat_name_set = :38, upper_rel_strat_unit_id = :39, row_changed_by = :40, row_changed_date = :41, row_created_by = :42, row_effective_date = :43, row_expiry_date = :44, row_quality = :45 where fossil_id = :47")
	if err != nil {
		return err
	}

	fossil_age.Row_changed_date = formatDate(fossil_age.Row_changed_date)
	fossil_age.Effective_date = formatDateString(fossil_age.Effective_date)
	fossil_age.Expiry_date = formatDateString(fossil_age.Expiry_date)
	fossil_age.Row_effective_date = formatDateString(fossil_age.Row_effective_date)
	fossil_age.Row_expiry_date = formatDateString(fossil_age.Row_expiry_date)






	rows, err := stmt.Exec(fossil_age.Age_interp_id, fossil_age.Active_ind, fossil_age.Age_confidence_id, fossil_age.Average_age, fossil_age.Average_age_error_minus, fossil_age.Average_age_error_plus, fossil_age.Average_rel_strat_name_set, fossil_age.Average_rel_strat_unit_id, fossil_age.Effective_date, fossil_age.Expiry_date, fossil_age.Lower_max_age, fossil_age.Lower_max_age_error_minus, fossil_age.Lower_max_age_error_plus, fossil_age.Lower_min_age, fossil_age.Lower_min_age_error_minus, fossil_age.Lower_min_age_error_plus, fossil_age.Lower_rel_strat_name_set, fossil_age.Lower_rel_strat_unit_id, fossil_age.Max_age, fossil_age.Max_age_error_minus, fossil_age.Max_age_error_plus, fossil_age.Min_age, fossil_age.Min_age_error_minus, fossil_age.Min_age_error_plus, fossil_age.Post_qualifier, fossil_age.Ppdm_guid, fossil_age.Preferred_ind, fossil_age.Pre_qualifier, fossil_age.Remark, fossil_age.Source, fossil_age.Source_document_id, fossil_age.Upper_max_age, fossil_age.Upper_max_age_error_minus, fossil_age.Upper_max_age_error_plus, fossil_age.Upper_min_age, fossil_age.Upper_min_age_error_minus, fossil_age.Upper_min_age_error_plus, fossil_age.Upper_rel_strat_name_set, fossil_age.Upper_rel_strat_unit_id, fossil_age.Row_changed_by, fossil_age.Row_changed_date, fossil_age.Row_created_by, fossil_age.Row_effective_date, fossil_age.Row_expiry_date, fossil_age.Row_quality, fossil_age.Fossil_id)
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

func PatchFossilAge(c *fiber.Ctx) error {
	id := c.Params("id")

	updatedFields := make(map[string]interface{})

	if err := c.BodyParser(&updatedFields); err != nil {
		return err
	}
	if _, ok := updatedFields["row_created_date"]; ok {
        return c.Status(400).SendString("Cannot update row_created_date field")
    }

	query := "update fossil_age set "
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

func DeleteFossilAge(c *fiber.Ctx) error {
	id := c.Params("id")
	var fossil_age dto.Fossil_age
	fossil_age.Fossil_id = id

	stmt, err := db.Prepare("delete from fossil_age where fossil_id = :1")
	if err != nil {
		return err
	}

	rows, err := stmt.Exec(fossil_age.Fossil_id)
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


