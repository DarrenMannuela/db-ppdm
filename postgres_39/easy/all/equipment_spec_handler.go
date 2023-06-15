package apiv1

import (
	dto "github.com/DarrenMannuela/pt_gtn_bibliography/dto"
	"github.com/gofiber/fiber/v2"
	"strconv"
	"encoding/json"
)


func GetEquipmentSpec(c *fiber.Ctx) error {
	rows, err := db.Query("select * from equipment_spec")
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}
	defer rows.Close()
	var result []dto.Equipment_spec

	for rows.Next() {
		var equipment_spec dto.Equipment_spec
		if err := rows.Scan(&equipment_spec.Equipment_id, &equipment_spec.Spec_set_id, &equipment_spec.Spec_type, &equipment_spec.Spec_id, &equipment_spec.Active_ind, &equipment_spec.Average_value, &equipment_spec.Average_value_ouom, &equipment_spec.Average_value_uom, &equipment_spec.Cost, &equipment_spec.Currency_conversion, &equipment_spec.Currency_ouom, &equipment_spec.Currency_uom, &equipment_spec.Date_format_desc, &equipment_spec.Effective_date, &equipment_spec.Equip_maint_id, &equipment_spec.Expiry_date, &equipment_spec.Max_date, &equipment_spec.Max_value, &equipment_spec.Max_value_ouom, &equipment_spec.Max_value_uom, &equipment_spec.Min_date, &equipment_spec.Min_value, &equipment_spec.Min_value_ouom, &equipment_spec.Min_value_uom, &equipment_spec.Ppdm_guid, &equipment_spec.Ratio_name, &equipment_spec.Ratio_value_average, &equipment_spec.Ratio_value_maximum, &equipment_spec.Ratio_value_minimum, &equipment_spec.Reference_value, &equipment_spec.Reference_value_ouom, &equipment_spec.Reference_value_type, &equipment_spec.Reference_value_uom, &equipment_spec.Remark, &equipment_spec.Source, &equipment_spec.Spec_code, &equipment_spec.Spec_desc, &equipment_spec.Row_changed_by, &equipment_spec.Row_changed_date, &equipment_spec.Row_created_by, &equipment_spec.Row_created_date, &equipment_spec.Row_effective_date, &equipment_spec.Row_expiry_date, &equipment_spec.Row_quality); err != nil {
			return err // Exit if we get an error
		}

		// Append equipment_spec to result
		result = append(result, equipment_spec)
	}
	// Return result in JSON format
	return c.JSON(result)
	 
}

func SetEquipmentSpec(c *fiber.Ctx) error {
	var equipment_spec dto.Equipment_spec

	setDefaults(&equipment_spec)

	if err := json.Unmarshal(c.Body(), &equipment_spec); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	stmt, err := db.Prepare("insert into equipment_spec values(:1, :2, :3, :4, :5, :6, :7, :8, :9, :10, :11, :12, :13, :14, :15, :16, :17, :18, :19, :20, :21, :22, :23, :24, :25, :26, :27, :28, :29, :30, :31, :32, :33, :34, :35, :36, :37, :38, :39, :40, :41, :42, :43, :44)")
	if err != nil {
		return err
	}
	equipment_spec.Row_created_date = formatDate(equipment_spec.Row_created_date)
	equipment_spec.Row_changed_date = formatDate(equipment_spec.Row_changed_date)
	equipment_spec.Date_format_desc = formatDateString(equipment_spec.Date_format_desc)
	equipment_spec.Effective_date = formatDateString(equipment_spec.Effective_date)
	equipment_spec.Expiry_date = formatDateString(equipment_spec.Expiry_date)
	equipment_spec.Max_date = formatDateString(equipment_spec.Max_date)
	equipment_spec.Min_date = formatDateString(equipment_spec.Min_date)
	equipment_spec.Row_effective_date = formatDateString(equipment_spec.Row_effective_date)
	equipment_spec.Row_expiry_date = formatDateString(equipment_spec.Row_expiry_date)






	rows, err := stmt.Exec(equipment_spec.Equipment_id, equipment_spec.Spec_set_id, equipment_spec.Spec_type, equipment_spec.Spec_id, equipment_spec.Active_ind, equipment_spec.Average_value, equipment_spec.Average_value_ouom, equipment_spec.Average_value_uom, equipment_spec.Cost, equipment_spec.Currency_conversion, equipment_spec.Currency_ouom, equipment_spec.Currency_uom, equipment_spec.Date_format_desc, equipment_spec.Effective_date, equipment_spec.Equip_maint_id, equipment_spec.Expiry_date, equipment_spec.Max_date, equipment_spec.Max_value, equipment_spec.Max_value_ouom, equipment_spec.Max_value_uom, equipment_spec.Min_date, equipment_spec.Min_value, equipment_spec.Min_value_ouom, equipment_spec.Min_value_uom, equipment_spec.Ppdm_guid, equipment_spec.Ratio_name, equipment_spec.Ratio_value_average, equipment_spec.Ratio_value_maximum, equipment_spec.Ratio_value_minimum, equipment_spec.Reference_value, equipment_spec.Reference_value_ouom, equipment_spec.Reference_value_type, equipment_spec.Reference_value_uom, equipment_spec.Remark, equipment_spec.Source, equipment_spec.Spec_code, equipment_spec.Spec_desc, equipment_spec.Row_changed_by, equipment_spec.Row_changed_date, equipment_spec.Row_created_by, equipment_spec.Row_created_date, equipment_spec.Row_effective_date, equipment_spec.Row_expiry_date, equipment_spec.Row_quality)
	
	if err != nil {
		return err
	}

	return c.JSON(rows)
}

func UpdateEquipmentSpec(c *fiber.Ctx) error {
	var equipment_spec dto.Equipment_spec

	setDefaults(&equipment_spec)

	if err := json.Unmarshal(c.Body(), &equipment_spec); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	id := c.Params("id")
	equipment_spec.Equipment_id = id

    if equipment_spec.Row_created_date != nil {
        return c.Status(400).SendString("Cannot update row_created_date")
    }

	stmt, err := db.Prepare("update equipment_spec set spec_set_id = :1, spec_type = :2, spec_id = :3, active_ind = :4, average_value = :5, average_value_ouom = :6, average_value_uom = :7, cost = :8, currency_conversion = :9, currency_ouom = :10, currency_uom = :11, date_format_desc = :12, effective_date = :13, equip_maint_id = :14, expiry_date = :15, max_date = :16, max_value = :17, max_value_ouom = :18, max_value_uom = :19, min_date = :20, min_value = :21, min_value_ouom = :22, min_value_uom = :23, ppdm_guid = :24, ratio_name = :25, ratio_value_average = :26, ratio_value_maximum = :27, ratio_value_minimum = :28, reference_value = :29, reference_value_ouom = :30, reference_value_type = :31, reference_value_uom = :32, remark = :33, source = :34, spec_code = :35, spec_desc = :36, row_changed_by = :37, row_changed_date = :38, row_created_by = :39, row_effective_date = :40, row_expiry_date = :41, row_quality = :42 where equipment_id = :44")
	if err != nil {
		return err
	}

	equipment_spec.Row_changed_date = formatDate(equipment_spec.Row_changed_date)
	equipment_spec.Date_format_desc = formatDateString(equipment_spec.Date_format_desc)
	equipment_spec.Effective_date = formatDateString(equipment_spec.Effective_date)
	equipment_spec.Expiry_date = formatDateString(equipment_spec.Expiry_date)
	equipment_spec.Max_date = formatDateString(equipment_spec.Max_date)
	equipment_spec.Min_date = formatDateString(equipment_spec.Min_date)
	equipment_spec.Row_effective_date = formatDateString(equipment_spec.Row_effective_date)
	equipment_spec.Row_expiry_date = formatDateString(equipment_spec.Row_expiry_date)






	rows, err := stmt.Exec(equipment_spec.Spec_set_id, equipment_spec.Spec_type, equipment_spec.Spec_id, equipment_spec.Active_ind, equipment_spec.Average_value, equipment_spec.Average_value_ouom, equipment_spec.Average_value_uom, equipment_spec.Cost, equipment_spec.Currency_conversion, equipment_spec.Currency_ouom, equipment_spec.Currency_uom, equipment_spec.Date_format_desc, equipment_spec.Effective_date, equipment_spec.Equip_maint_id, equipment_spec.Expiry_date, equipment_spec.Max_date, equipment_spec.Max_value, equipment_spec.Max_value_ouom, equipment_spec.Max_value_uom, equipment_spec.Min_date, equipment_spec.Min_value, equipment_spec.Min_value_ouom, equipment_spec.Min_value_uom, equipment_spec.Ppdm_guid, equipment_spec.Ratio_name, equipment_spec.Ratio_value_average, equipment_spec.Ratio_value_maximum, equipment_spec.Ratio_value_minimum, equipment_spec.Reference_value, equipment_spec.Reference_value_ouom, equipment_spec.Reference_value_type, equipment_spec.Reference_value_uom, equipment_spec.Remark, equipment_spec.Source, equipment_spec.Spec_code, equipment_spec.Spec_desc, equipment_spec.Row_changed_by, equipment_spec.Row_changed_date, equipment_spec.Row_created_by, equipment_spec.Row_effective_date, equipment_spec.Row_expiry_date, equipment_spec.Row_quality, equipment_spec.Equipment_id)
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

func PatchEquipmentSpec(c *fiber.Ctx) error {
	id := c.Params("id")

	updatedFields := make(map[string]interface{})

	if err := c.BodyParser(&updatedFields); err != nil {
		return err
	}
	if _, ok := updatedFields["row_created_date"]; ok {
        return c.Status(400).SendString("Cannot update row_created_date field")
    }

	query := "update equipment_spec set "
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
	query += " where equipment_id = :id"

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

func DeleteEquipmentSpec(c *fiber.Ctx) error {
	id := c.Params("id")
	var equipment_spec dto.Equipment_spec
	equipment_spec.Equipment_id = id

	stmt, err := db.Prepare("delete from equipment_spec where equipment_id = :1")
	if err != nil {
		return err
	}

	rows, err := stmt.Exec(equipment_spec.Equipment_id)
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


