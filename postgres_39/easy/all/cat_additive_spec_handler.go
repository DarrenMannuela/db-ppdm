package apiv1

import (
	dto "github.com/DarrenMannuela/pt_gtn_bibliography/dto"
	"github.com/gofiber/fiber/v2"
	"strconv"
	"encoding/json"
)


func GetCatAdditiveSpec(c *fiber.Ctx) error {
	rows, err := db.Query("select * from cat_additive_spec")
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}
	defer rows.Close()
	var result []dto.Cat_additive_spec

	for rows.Next() {
		var cat_additive_spec dto.Cat_additive_spec
		if err := rows.Scan(&cat_additive_spec.Catalogue_additive_id, &cat_additive_spec.Spec_id, &cat_additive_spec.Active_ind, &cat_additive_spec.Additive_component_id, &cat_additive_spec.Additive_spec_type, &cat_additive_spec.Average_value, &cat_additive_spec.Average_value_ouom, &cat_additive_spec.Average_value_uom, &cat_additive_spec.Cost, &cat_additive_spec.Currency_conversion, &cat_additive_spec.Currency_ouom, &cat_additive_spec.Currency_uom, &cat_additive_spec.Date_format_desc, &cat_additive_spec.Effective_date, &cat_additive_spec.Expiry_date, &cat_additive_spec.Max_value, &cat_additive_spec.Max_value_ouom, &cat_additive_spec.Max_value_uom, &cat_additive_spec.Min_value, &cat_additive_spec.Min_value_ouom, &cat_additive_spec.Min_value_uom, &cat_additive_spec.Ppdm_guid, &cat_additive_spec.Remark, &cat_additive_spec.Source, &cat_additive_spec.Spec_code, &cat_additive_spec.Spec_desc, &cat_additive_spec.Row_changed_by, &cat_additive_spec.Row_changed_date, &cat_additive_spec.Row_created_by, &cat_additive_spec.Row_created_date, &cat_additive_spec.Row_effective_date, &cat_additive_spec.Row_expiry_date, &cat_additive_spec.Row_quality); err != nil {
			return err // Exit if we get an error
		}

		// Append cat_additive_spec to result
		result = append(result, cat_additive_spec)
	}
	// Return result in JSON format
	return c.JSON(result)
	 
}

func SetCatAdditiveSpec(c *fiber.Ctx) error {
	var cat_additive_spec dto.Cat_additive_spec

	setDefaults(&cat_additive_spec)

	if err := json.Unmarshal(c.Body(), &cat_additive_spec); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	stmt, err := db.Prepare("insert into cat_additive_spec values(:1, :2, :3, :4, :5, :6, :7, :8, :9, :10, :11, :12, :13, :14, :15, :16, :17, :18, :19, :20, :21, :22, :23, :24, :25, :26, :27, :28, :29, :30, :31, :32, :33)")
	if err != nil {
		return err
	}
	cat_additive_spec.Row_created_date = formatDate(cat_additive_spec.Row_created_date)
	cat_additive_spec.Row_changed_date = formatDate(cat_additive_spec.Row_changed_date)
	cat_additive_spec.Date_format_desc = formatDateString(cat_additive_spec.Date_format_desc)
	cat_additive_spec.Effective_date = formatDateString(cat_additive_spec.Effective_date)
	cat_additive_spec.Expiry_date = formatDateString(cat_additive_spec.Expiry_date)
	cat_additive_spec.Row_effective_date = formatDateString(cat_additive_spec.Row_effective_date)
	cat_additive_spec.Row_expiry_date = formatDateString(cat_additive_spec.Row_expiry_date)






	rows, err := stmt.Exec(cat_additive_spec.Catalogue_additive_id, cat_additive_spec.Spec_id, cat_additive_spec.Active_ind, cat_additive_spec.Additive_component_id, cat_additive_spec.Additive_spec_type, cat_additive_spec.Average_value, cat_additive_spec.Average_value_ouom, cat_additive_spec.Average_value_uom, cat_additive_spec.Cost, cat_additive_spec.Currency_conversion, cat_additive_spec.Currency_ouom, cat_additive_spec.Currency_uom, cat_additive_spec.Date_format_desc, cat_additive_spec.Effective_date, cat_additive_spec.Expiry_date, cat_additive_spec.Max_value, cat_additive_spec.Max_value_ouom, cat_additive_spec.Max_value_uom, cat_additive_spec.Min_value, cat_additive_spec.Min_value_ouom, cat_additive_spec.Min_value_uom, cat_additive_spec.Ppdm_guid, cat_additive_spec.Remark, cat_additive_spec.Source, cat_additive_spec.Spec_code, cat_additive_spec.Spec_desc, cat_additive_spec.Row_changed_by, cat_additive_spec.Row_changed_date, cat_additive_spec.Row_created_by, cat_additive_spec.Row_created_date, cat_additive_spec.Row_effective_date, cat_additive_spec.Row_expiry_date, cat_additive_spec.Row_quality)
	
	if err != nil {
		return err
	}

	return c.JSON(rows)
}

func UpdateCatAdditiveSpec(c *fiber.Ctx) error {
	var cat_additive_spec dto.Cat_additive_spec

	setDefaults(&cat_additive_spec)

	if err := json.Unmarshal(c.Body(), &cat_additive_spec); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	id := c.Params("id")
	cat_additive_spec.Catalogue_additive_id = id

    if cat_additive_spec.Row_created_date != nil {
        return c.Status(400).SendString("Cannot update row_created_date")
    }

	stmt, err := db.Prepare("update cat_additive_spec set spec_id = :1, active_ind = :2, additive_component_id = :3, additive_spec_type = :4, average_value = :5, average_value_ouom = :6, average_value_uom = :7, cost = :8, currency_conversion = :9, currency_ouom = :10, currency_uom = :11, date_format_desc = :12, effective_date = :13, expiry_date = :14, max_value = :15, max_value_ouom = :16, max_value_uom = :17, min_value = :18, min_value_ouom = :19, min_value_uom = :20, ppdm_guid = :21, remark = :22, source = :23, spec_code = :24, spec_desc = :25, row_changed_by = :26, row_changed_date = :27, row_created_by = :28, row_effective_date = :29, row_expiry_date = :30, row_quality = :31 where catalogue_additive_id = :33")
	if err != nil {
		return err
	}

	cat_additive_spec.Row_changed_date = formatDate(cat_additive_spec.Row_changed_date)
	cat_additive_spec.Date_format_desc = formatDateString(cat_additive_spec.Date_format_desc)
	cat_additive_spec.Effective_date = formatDateString(cat_additive_spec.Effective_date)
	cat_additive_spec.Expiry_date = formatDateString(cat_additive_spec.Expiry_date)
	cat_additive_spec.Row_effective_date = formatDateString(cat_additive_spec.Row_effective_date)
	cat_additive_spec.Row_expiry_date = formatDateString(cat_additive_spec.Row_expiry_date)






	rows, err := stmt.Exec(cat_additive_spec.Spec_id, cat_additive_spec.Active_ind, cat_additive_spec.Additive_component_id, cat_additive_spec.Additive_spec_type, cat_additive_spec.Average_value, cat_additive_spec.Average_value_ouom, cat_additive_spec.Average_value_uom, cat_additive_spec.Cost, cat_additive_spec.Currency_conversion, cat_additive_spec.Currency_ouom, cat_additive_spec.Currency_uom, cat_additive_spec.Date_format_desc, cat_additive_spec.Effective_date, cat_additive_spec.Expiry_date, cat_additive_spec.Max_value, cat_additive_spec.Max_value_ouom, cat_additive_spec.Max_value_uom, cat_additive_spec.Min_value, cat_additive_spec.Min_value_ouom, cat_additive_spec.Min_value_uom, cat_additive_spec.Ppdm_guid, cat_additive_spec.Remark, cat_additive_spec.Source, cat_additive_spec.Spec_code, cat_additive_spec.Spec_desc, cat_additive_spec.Row_changed_by, cat_additive_spec.Row_changed_date, cat_additive_spec.Row_created_by, cat_additive_spec.Row_effective_date, cat_additive_spec.Row_expiry_date, cat_additive_spec.Row_quality, cat_additive_spec.Catalogue_additive_id)
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

func PatchCatAdditiveSpec(c *fiber.Ctx) error {
	id := c.Params("id")

	updatedFields := make(map[string]interface{})

	if err := c.BodyParser(&updatedFields); err != nil {
		return err
	}
	if _, ok := updatedFields["row_created_date"]; ok {
        return c.Status(400).SendString("Cannot update row_created_date field")
    }

	query := "update cat_additive_spec set "
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
		} else if key == "date_format_desc" || key == "effective_date" || key == "expiry_date" || key == "row_effective_date" || key == "row_expiry_date"      {
			if date, ok := value.(string); ok && len(date) > 0 {
				formattedDate := formatDateString(&date)
				value = formattedDate
			}
		}
		values = append(values, value)
	}
	query += " where catalogue_additive_id = :id"

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

func DeleteCatAdditiveSpec(c *fiber.Ctx) error {
	id := c.Params("id")
	var cat_additive_spec dto.Cat_additive_spec
	cat_additive_spec.Catalogue_additive_id = id

	stmt, err := db.Prepare("delete from cat_additive_spec where catalogue_additive_id = :1")
	if err != nil {
		return err
	}

	rows, err := stmt.Exec(cat_additive_spec.Catalogue_additive_id)
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


