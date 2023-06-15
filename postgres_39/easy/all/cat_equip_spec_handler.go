package apiv1

import (
	dto "github.com/DarrenMannuela/pt_gtn_bibliography/dto"
	"github.com/gofiber/fiber/v2"
	"strconv"
	"encoding/json"
)


func GetCatEquipSpec(c *fiber.Ctx) error {
	rows, err := db.Query("select * from cat_equip_spec")
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}
	defer rows.Close()
	var result []dto.Cat_equip_spec

	for rows.Next() {
		var cat_equip_spec dto.Cat_equip_spec
		if err := rows.Scan(&cat_equip_spec.Catalogue_equip_id, &cat_equip_spec.Spec_id, &cat_equip_spec.Active_ind, &cat_equip_spec.Average_value, &cat_equip_spec.Average_value_ouom, &cat_equip_spec.Average_value_uom, &cat_equip_spec.Cost, &cat_equip_spec.Currency_conversion, &cat_equip_spec.Currency_ouom, &cat_equip_spec.Currency_uom, &cat_equip_spec.Date_format_desc, &cat_equip_spec.Effective_date, &cat_equip_spec.Expiry_date, &cat_equip_spec.Max_date, &cat_equip_spec.Max_value, &cat_equip_spec.Max_value_ouom, &cat_equip_spec.Max_value_uom, &cat_equip_spec.Min_date, &cat_equip_spec.Min_value, &cat_equip_spec.Min_value_ouom, &cat_equip_spec.Min_value_uom, &cat_equip_spec.Ppdm_guid, &cat_equip_spec.Reference_value, &cat_equip_spec.Reference_value_ouom, &cat_equip_spec.Reference_value_type, &cat_equip_spec.Reference_value_uom, &cat_equip_spec.Remark, &cat_equip_spec.Source, &cat_equip_spec.Spec_code, &cat_equip_spec.Spec_desc, &cat_equip_spec.Spec_type, &cat_equip_spec.Row_changed_by, &cat_equip_spec.Row_changed_date, &cat_equip_spec.Row_created_by, &cat_equip_spec.Row_created_date, &cat_equip_spec.Row_effective_date, &cat_equip_spec.Row_expiry_date, &cat_equip_spec.Row_quality); err != nil {
			return err // Exit if we get an error
		}

		// Append cat_equip_spec to result
		result = append(result, cat_equip_spec)
	}
	// Return result in JSON format
	return c.JSON(result)
	 
}

func SetCatEquipSpec(c *fiber.Ctx) error {
	var cat_equip_spec dto.Cat_equip_spec

	setDefaults(&cat_equip_spec)

	if err := json.Unmarshal(c.Body(), &cat_equip_spec); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	stmt, err := db.Prepare("insert into cat_equip_spec values(:1, :2, :3, :4, :5, :6, :7, :8, :9, :10, :11, :12, :13, :14, :15, :16, :17, :18, :19, :20, :21, :22, :23, :24, :25, :26, :27, :28, :29, :30, :31, :32, :33, :34, :35, :36, :37, :38)")
	if err != nil {
		return err
	}
	cat_equip_spec.Row_created_date = formatDate(cat_equip_spec.Row_created_date)
	cat_equip_spec.Row_changed_date = formatDate(cat_equip_spec.Row_changed_date)
	cat_equip_spec.Date_format_desc = formatDateString(cat_equip_spec.Date_format_desc)
	cat_equip_spec.Effective_date = formatDateString(cat_equip_spec.Effective_date)
	cat_equip_spec.Expiry_date = formatDateString(cat_equip_spec.Expiry_date)
	cat_equip_spec.Max_date = formatDateString(cat_equip_spec.Max_date)
	cat_equip_spec.Min_date = formatDateString(cat_equip_spec.Min_date)
	cat_equip_spec.Row_effective_date = formatDateString(cat_equip_spec.Row_effective_date)
	cat_equip_spec.Row_expiry_date = formatDateString(cat_equip_spec.Row_expiry_date)






	rows, err := stmt.Exec(cat_equip_spec.Catalogue_equip_id, cat_equip_spec.Spec_id, cat_equip_spec.Active_ind, cat_equip_spec.Average_value, cat_equip_spec.Average_value_ouom, cat_equip_spec.Average_value_uom, cat_equip_spec.Cost, cat_equip_spec.Currency_conversion, cat_equip_spec.Currency_ouom, cat_equip_spec.Currency_uom, cat_equip_spec.Date_format_desc, cat_equip_spec.Effective_date, cat_equip_spec.Expiry_date, cat_equip_spec.Max_date, cat_equip_spec.Max_value, cat_equip_spec.Max_value_ouom, cat_equip_spec.Max_value_uom, cat_equip_spec.Min_date, cat_equip_spec.Min_value, cat_equip_spec.Min_value_ouom, cat_equip_spec.Min_value_uom, cat_equip_spec.Ppdm_guid, cat_equip_spec.Reference_value, cat_equip_spec.Reference_value_ouom, cat_equip_spec.Reference_value_type, cat_equip_spec.Reference_value_uom, cat_equip_spec.Remark, cat_equip_spec.Source, cat_equip_spec.Spec_code, cat_equip_spec.Spec_desc, cat_equip_spec.Spec_type, cat_equip_spec.Row_changed_by, cat_equip_spec.Row_changed_date, cat_equip_spec.Row_created_by, cat_equip_spec.Row_created_date, cat_equip_spec.Row_effective_date, cat_equip_spec.Row_expiry_date, cat_equip_spec.Row_quality)
	
	if err != nil {
		return err
	}

	return c.JSON(rows)
}

func UpdateCatEquipSpec(c *fiber.Ctx) error {
	var cat_equip_spec dto.Cat_equip_spec

	setDefaults(&cat_equip_spec)

	if err := json.Unmarshal(c.Body(), &cat_equip_spec); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	id := c.Params("id")
	cat_equip_spec.Catalogue_equip_id = id

    if cat_equip_spec.Row_created_date != nil {
        return c.Status(400).SendString("Cannot update row_created_date")
    }

	stmt, err := db.Prepare("update cat_equip_spec set spec_id = :1, active_ind = :2, average_value = :3, average_value_ouom = :4, average_value_uom = :5, cost = :6, currency_conversion = :7, currency_ouom = :8, currency_uom = :9, date_format_desc = :10, effective_date = :11, expiry_date = :12, max_date = :13, max_value = :14, max_value_ouom = :15, max_value_uom = :16, min_date = :17, min_value = :18, min_value_ouom = :19, min_value_uom = :20, ppdm_guid = :21, reference_value = :22, reference_value_ouom = :23, reference_value_type = :24, reference_value_uom = :25, remark = :26, source = :27, spec_code = :28, spec_desc = :29, spec_type = :30, row_changed_by = :31, row_changed_date = :32, row_created_by = :33, row_effective_date = :34, row_expiry_date = :35, row_quality = :36 where catalogue_equip_id = :38")
	if err != nil {
		return err
	}

	cat_equip_spec.Row_changed_date = formatDate(cat_equip_spec.Row_changed_date)
	cat_equip_spec.Date_format_desc = formatDateString(cat_equip_spec.Date_format_desc)
	cat_equip_spec.Effective_date = formatDateString(cat_equip_spec.Effective_date)
	cat_equip_spec.Expiry_date = formatDateString(cat_equip_spec.Expiry_date)
	cat_equip_spec.Max_date = formatDateString(cat_equip_spec.Max_date)
	cat_equip_spec.Min_date = formatDateString(cat_equip_spec.Min_date)
	cat_equip_spec.Row_effective_date = formatDateString(cat_equip_spec.Row_effective_date)
	cat_equip_spec.Row_expiry_date = formatDateString(cat_equip_spec.Row_expiry_date)






	rows, err := stmt.Exec(cat_equip_spec.Spec_id, cat_equip_spec.Active_ind, cat_equip_spec.Average_value, cat_equip_spec.Average_value_ouom, cat_equip_spec.Average_value_uom, cat_equip_spec.Cost, cat_equip_spec.Currency_conversion, cat_equip_spec.Currency_ouom, cat_equip_spec.Currency_uom, cat_equip_spec.Date_format_desc, cat_equip_spec.Effective_date, cat_equip_spec.Expiry_date, cat_equip_spec.Max_date, cat_equip_spec.Max_value, cat_equip_spec.Max_value_ouom, cat_equip_spec.Max_value_uom, cat_equip_spec.Min_date, cat_equip_spec.Min_value, cat_equip_spec.Min_value_ouom, cat_equip_spec.Min_value_uom, cat_equip_spec.Ppdm_guid, cat_equip_spec.Reference_value, cat_equip_spec.Reference_value_ouom, cat_equip_spec.Reference_value_type, cat_equip_spec.Reference_value_uom, cat_equip_spec.Remark, cat_equip_spec.Source, cat_equip_spec.Spec_code, cat_equip_spec.Spec_desc, cat_equip_spec.Spec_type, cat_equip_spec.Row_changed_by, cat_equip_spec.Row_changed_date, cat_equip_spec.Row_created_by, cat_equip_spec.Row_effective_date, cat_equip_spec.Row_expiry_date, cat_equip_spec.Row_quality, cat_equip_spec.Catalogue_equip_id)
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

func PatchCatEquipSpec(c *fiber.Ctx) error {
	id := c.Params("id")

	updatedFields := make(map[string]interface{})

	if err := c.BodyParser(&updatedFields); err != nil {
		return err
	}
	if _, ok := updatedFields["row_created_date"]; ok {
        return c.Status(400).SendString("Cannot update row_created_date field")
    }

	query := "update cat_equip_spec set "
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
		} else if key == "date_format_desc" || key == "effective_date" || key == "expiry_date" || key == "max_date" || key == "min_date" || key == "row_effective_date" || key == "row_expiry_date"      {
			if date, ok := value.(string); ok && len(date) > 0 {
				formattedDate := formatDateString(&date)
				value = formattedDate
			}
		}
		values = append(values, value)
	}
	query += " where catalogue_equip_id = :id"

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

func DeleteCatEquipSpec(c *fiber.Ctx) error {
	id := c.Params("id")
	var cat_equip_spec dto.Cat_equip_spec
	cat_equip_spec.Catalogue_equip_id = id

	stmt, err := db.Prepare("delete from cat_equip_spec where catalogue_equip_id = :1")
	if err != nil {
		return err
	}

	rows, err := stmt.Exec(cat_equip_spec.Catalogue_equip_id)
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


