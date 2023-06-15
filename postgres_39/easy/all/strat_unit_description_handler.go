package apiv1

import (
	dto "github.com/DarrenMannuela/pt_gtn_bibliography/dto"
	"github.com/gofiber/fiber/v2"
	"strconv"
	"encoding/json"
)


func GetStratUnitDescription(c *fiber.Ctx) error {
	rows, err := db.Query("select * from strat_unit_description")
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}
	defer rows.Close()
	var result []dto.Strat_unit_description

	for rows.Next() {
		var strat_unit_description dto.Strat_unit_description
		if err := rows.Scan(&strat_unit_description.Strat_name_set_id, &strat_unit_description.Strat_unit_id, &strat_unit_description.Description_seq_no, &strat_unit_description.Active_ind, &strat_unit_description.Average_value, &strat_unit_description.Average_value_ouom, &strat_unit_description.Average_value_uom, &strat_unit_description.Date_format_desc, &strat_unit_description.Description, &strat_unit_description.Description_date, &strat_unit_description.Description_type, &strat_unit_description.Effective_date, &strat_unit_description.Expiry_date, &strat_unit_description.Max_value, &strat_unit_description.Max_value_ouom, &strat_unit_description.Max_value_uom, &strat_unit_description.Min_value, &strat_unit_description.Min_value_ouom, &strat_unit_description.Min_value_uom, &strat_unit_description.Ppdm_guid, &strat_unit_description.Reference_pages, &strat_unit_description.Remark, &strat_unit_description.Source, &strat_unit_description.Source_document_id, &strat_unit_description.Strat_unit_desc, &strat_unit_description.Row_changed_by, &strat_unit_description.Row_changed_date, &strat_unit_description.Row_created_by, &strat_unit_description.Row_created_date, &strat_unit_description.Row_effective_date, &strat_unit_description.Row_expiry_date, &strat_unit_description.Row_quality); err != nil {
			return err // Exit if we get an error
		}

		// Append strat_unit_description to result
		result = append(result, strat_unit_description)
	}
	// Return result in JSON format
	return c.JSON(result)
	 
}

func SetStratUnitDescription(c *fiber.Ctx) error {
	var strat_unit_description dto.Strat_unit_description

	setDefaults(&strat_unit_description)

	if err := json.Unmarshal(c.Body(), &strat_unit_description); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	stmt, err := db.Prepare("insert into strat_unit_description values(:1, :2, :3, :4, :5, :6, :7, :8, :9, :10, :11, :12, :13, :14, :15, :16, :17, :18, :19, :20, :21, :22, :23, :24, :25, :26, :27, :28, :29, :30, :31, :32)")
	if err != nil {
		return err
	}
	strat_unit_description.Row_created_date = formatDate(strat_unit_description.Row_created_date)
	strat_unit_description.Row_changed_date = formatDate(strat_unit_description.Row_changed_date)
	strat_unit_description.Date_format_desc = formatDateString(strat_unit_description.Date_format_desc)
	strat_unit_description.Description_date = formatDateString(strat_unit_description.Description_date)
	strat_unit_description.Effective_date = formatDateString(strat_unit_description.Effective_date)
	strat_unit_description.Expiry_date = formatDateString(strat_unit_description.Expiry_date)
	strat_unit_description.Row_effective_date = formatDateString(strat_unit_description.Row_effective_date)
	strat_unit_description.Row_expiry_date = formatDateString(strat_unit_description.Row_expiry_date)






	rows, err := stmt.Exec(strat_unit_description.Strat_name_set_id, strat_unit_description.Strat_unit_id, strat_unit_description.Description_seq_no, strat_unit_description.Active_ind, strat_unit_description.Average_value, strat_unit_description.Average_value_ouom, strat_unit_description.Average_value_uom, strat_unit_description.Date_format_desc, strat_unit_description.Description, strat_unit_description.Description_date, strat_unit_description.Description_type, strat_unit_description.Effective_date, strat_unit_description.Expiry_date, strat_unit_description.Max_value, strat_unit_description.Max_value_ouom, strat_unit_description.Max_value_uom, strat_unit_description.Min_value, strat_unit_description.Min_value_ouom, strat_unit_description.Min_value_uom, strat_unit_description.Ppdm_guid, strat_unit_description.Reference_pages, strat_unit_description.Remark, strat_unit_description.Source, strat_unit_description.Source_document_id, strat_unit_description.Strat_unit_desc, strat_unit_description.Row_changed_by, strat_unit_description.Row_changed_date, strat_unit_description.Row_created_by, strat_unit_description.Row_created_date, strat_unit_description.Row_effective_date, strat_unit_description.Row_expiry_date, strat_unit_description.Row_quality)
	
	if err != nil {
		return err
	}

	return c.JSON(rows)
}

func UpdateStratUnitDescription(c *fiber.Ctx) error {
	var strat_unit_description dto.Strat_unit_description

	setDefaults(&strat_unit_description)

	if err := json.Unmarshal(c.Body(), &strat_unit_description); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	id := c.Params("id")
	strat_unit_description.Strat_name_set_id = id

    if strat_unit_description.Row_created_date != nil {
        return c.Status(400).SendString("Cannot update row_created_date")
    }

	stmt, err := db.Prepare("update strat_unit_description set strat_unit_id = :1, description_seq_no = :2, active_ind = :3, average_value = :4, average_value_ouom = :5, average_value_uom = :6, date_format_desc = :7, description = :8, description_date = :9, description_type = :10, effective_date = :11, expiry_date = :12, max_value = :13, max_value_ouom = :14, max_value_uom = :15, min_value = :16, min_value_ouom = :17, min_value_uom = :18, ppdm_guid = :19, reference_pages = :20, remark = :21, source = :22, source_document_id = :23, strat_unit_desc = :24, row_changed_by = :25, row_changed_date = :26, row_created_by = :27, row_effective_date = :28, row_expiry_date = :29, row_quality = :30 where strat_name_set_id = :32")
	if err != nil {
		return err
	}

	strat_unit_description.Row_changed_date = formatDate(strat_unit_description.Row_changed_date)
	strat_unit_description.Date_format_desc = formatDateString(strat_unit_description.Date_format_desc)
	strat_unit_description.Description_date = formatDateString(strat_unit_description.Description_date)
	strat_unit_description.Effective_date = formatDateString(strat_unit_description.Effective_date)
	strat_unit_description.Expiry_date = formatDateString(strat_unit_description.Expiry_date)
	strat_unit_description.Row_effective_date = formatDateString(strat_unit_description.Row_effective_date)
	strat_unit_description.Row_expiry_date = formatDateString(strat_unit_description.Row_expiry_date)






	rows, err := stmt.Exec(strat_unit_description.Strat_unit_id, strat_unit_description.Description_seq_no, strat_unit_description.Active_ind, strat_unit_description.Average_value, strat_unit_description.Average_value_ouom, strat_unit_description.Average_value_uom, strat_unit_description.Date_format_desc, strat_unit_description.Description, strat_unit_description.Description_date, strat_unit_description.Description_type, strat_unit_description.Effective_date, strat_unit_description.Expiry_date, strat_unit_description.Max_value, strat_unit_description.Max_value_ouom, strat_unit_description.Max_value_uom, strat_unit_description.Min_value, strat_unit_description.Min_value_ouom, strat_unit_description.Min_value_uom, strat_unit_description.Ppdm_guid, strat_unit_description.Reference_pages, strat_unit_description.Remark, strat_unit_description.Source, strat_unit_description.Source_document_id, strat_unit_description.Strat_unit_desc, strat_unit_description.Row_changed_by, strat_unit_description.Row_changed_date, strat_unit_description.Row_created_by, strat_unit_description.Row_effective_date, strat_unit_description.Row_expiry_date, strat_unit_description.Row_quality, strat_unit_description.Strat_name_set_id)
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

func PatchStratUnitDescription(c *fiber.Ctx) error {
	id := c.Params("id")

	updatedFields := make(map[string]interface{})

	if err := c.BodyParser(&updatedFields); err != nil {
		return err
	}
	if _, ok := updatedFields["row_created_date"]; ok {
        return c.Status(400).SendString("Cannot update row_created_date field")
    }

	query := "update strat_unit_description set "
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
		} else if key == "date_format_desc" || key == "description_date" || key == "effective_date" || key == "expiry_date" || key == "row_effective_date" || key == "row_expiry_date"      {
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

func DeleteStratUnitDescription(c *fiber.Ctx) error {
	id := c.Params("id")
	var strat_unit_description dto.Strat_unit_description
	strat_unit_description.Strat_name_set_id = id

	stmt, err := db.Prepare("delete from strat_unit_description where strat_name_set_id = :1")
	if err != nil {
		return err
	}

	rows, err := stmt.Exec(strat_unit_description.Strat_name_set_id)
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


