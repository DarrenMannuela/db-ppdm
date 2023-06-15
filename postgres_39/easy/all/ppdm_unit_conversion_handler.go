package apiv1

import (
	dto "github.com/DarrenMannuela/pt_gtn_bibliography/dto"
	"github.com/gofiber/fiber/v2"
	"strconv"
	"encoding/json"
)


func GetPpdmUnitConversion(c *fiber.Ctx) error {
	rows, err := db.Query("select * from ppdm_unit_conversion")
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}
	defer rows.Close()
	var result []dto.Ppdm_unit_conversion

	for rows.Next() {
		var ppdm_unit_conversion dto.Ppdm_unit_conversion
		if err := rows.Scan(&ppdm_unit_conversion.From_uom_id, &ppdm_unit_conversion.To_uom_id, &ppdm_unit_conversion.Active_ind, &ppdm_unit_conversion.Effective_date, &ppdm_unit_conversion.Exact_conversion_ind, &ppdm_unit_conversion.Expiry_date, &ppdm_unit_conversion.Factor_denominator, &ppdm_unit_conversion.Factor_numerator, &ppdm_unit_conversion.Post_offset, &ppdm_unit_conversion.Ppdm_guid, &ppdm_unit_conversion.Pre_offset, &ppdm_unit_conversion.Remark, &ppdm_unit_conversion.Source, &ppdm_unit_conversion.Source_document_id, &ppdm_unit_conversion.Unit_expression, &ppdm_unit_conversion.Unit_quantity_type, &ppdm_unit_conversion.Row_changed_by, &ppdm_unit_conversion.Row_changed_date, &ppdm_unit_conversion.Row_created_by, &ppdm_unit_conversion.Row_created_date, &ppdm_unit_conversion.Row_effective_date, &ppdm_unit_conversion.Row_expiry_date, &ppdm_unit_conversion.Row_quality); err != nil {
			return err // Exit if we get an error
		}

		// Append ppdm_unit_conversion to result
		result = append(result, ppdm_unit_conversion)
	}
	// Return result in JSON format
	return c.JSON(result)
	 
}

func SetPpdmUnitConversion(c *fiber.Ctx) error {
	var ppdm_unit_conversion dto.Ppdm_unit_conversion

	setDefaults(&ppdm_unit_conversion)

	if err := json.Unmarshal(c.Body(), &ppdm_unit_conversion); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	stmt, err := db.Prepare("insert into ppdm_unit_conversion values(:1, :2, :3, :4, :5, :6, :7, :8, :9, :10, :11, :12, :13, :14, :15, :16, :17, :18, :19, :20, :21, :22, :23)")
	if err != nil {
		return err
	}
	ppdm_unit_conversion.Row_created_date = formatDate(ppdm_unit_conversion.Row_created_date)
	ppdm_unit_conversion.Row_changed_date = formatDate(ppdm_unit_conversion.Row_changed_date)
	ppdm_unit_conversion.Effective_date = formatDateString(ppdm_unit_conversion.Effective_date)
	ppdm_unit_conversion.Expiry_date = formatDateString(ppdm_unit_conversion.Expiry_date)
	ppdm_unit_conversion.Row_effective_date = formatDateString(ppdm_unit_conversion.Row_effective_date)
	ppdm_unit_conversion.Row_expiry_date = formatDateString(ppdm_unit_conversion.Row_expiry_date)






	rows, err := stmt.Exec(ppdm_unit_conversion.From_uom_id, ppdm_unit_conversion.To_uom_id, ppdm_unit_conversion.Active_ind, ppdm_unit_conversion.Effective_date, ppdm_unit_conversion.Exact_conversion_ind, ppdm_unit_conversion.Expiry_date, ppdm_unit_conversion.Factor_denominator, ppdm_unit_conversion.Factor_numerator, ppdm_unit_conversion.Post_offset, ppdm_unit_conversion.Ppdm_guid, ppdm_unit_conversion.Pre_offset, ppdm_unit_conversion.Remark, ppdm_unit_conversion.Source, ppdm_unit_conversion.Source_document_id, ppdm_unit_conversion.Unit_expression, ppdm_unit_conversion.Unit_quantity_type, ppdm_unit_conversion.Row_changed_by, ppdm_unit_conversion.Row_changed_date, ppdm_unit_conversion.Row_created_by, ppdm_unit_conversion.Row_created_date, ppdm_unit_conversion.Row_effective_date, ppdm_unit_conversion.Row_expiry_date, ppdm_unit_conversion.Row_quality)
	
	if err != nil {
		return err
	}

	return c.JSON(rows)
}

func UpdatePpdmUnitConversion(c *fiber.Ctx) error {
	var ppdm_unit_conversion dto.Ppdm_unit_conversion

	setDefaults(&ppdm_unit_conversion)

	if err := json.Unmarshal(c.Body(), &ppdm_unit_conversion); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	id := c.Params("id")
	ppdm_unit_conversion.From_uom_id = id

    if ppdm_unit_conversion.Row_created_date != nil {
        return c.Status(400).SendString("Cannot update row_created_date")
    }

	stmt, err := db.Prepare("update ppdm_unit_conversion set to_uom_id = :1, active_ind = :2, effective_date = :3, exact_conversion_ind = :4, expiry_date = :5, factor_denominator = :6, factor_numerator = :7, post_offset = :8, ppdm_guid = :9, pre_offset = :10, remark = :11, source = :12, source_document_id = :13, unit_expression = :14, unit_quantity_type = :15, row_changed_by = :16, row_changed_date = :17, row_created_by = :18, row_effective_date = :19, row_expiry_date = :20, row_quality = :21 where from_uom_id = :23")
	if err != nil {
		return err
	}

	ppdm_unit_conversion.Row_changed_date = formatDate(ppdm_unit_conversion.Row_changed_date)
	ppdm_unit_conversion.Effective_date = formatDateString(ppdm_unit_conversion.Effective_date)
	ppdm_unit_conversion.Expiry_date = formatDateString(ppdm_unit_conversion.Expiry_date)
	ppdm_unit_conversion.Row_effective_date = formatDateString(ppdm_unit_conversion.Row_effective_date)
	ppdm_unit_conversion.Row_expiry_date = formatDateString(ppdm_unit_conversion.Row_expiry_date)






	rows, err := stmt.Exec(ppdm_unit_conversion.To_uom_id, ppdm_unit_conversion.Active_ind, ppdm_unit_conversion.Effective_date, ppdm_unit_conversion.Exact_conversion_ind, ppdm_unit_conversion.Expiry_date, ppdm_unit_conversion.Factor_denominator, ppdm_unit_conversion.Factor_numerator, ppdm_unit_conversion.Post_offset, ppdm_unit_conversion.Ppdm_guid, ppdm_unit_conversion.Pre_offset, ppdm_unit_conversion.Remark, ppdm_unit_conversion.Source, ppdm_unit_conversion.Source_document_id, ppdm_unit_conversion.Unit_expression, ppdm_unit_conversion.Unit_quantity_type, ppdm_unit_conversion.Row_changed_by, ppdm_unit_conversion.Row_changed_date, ppdm_unit_conversion.Row_created_by, ppdm_unit_conversion.Row_effective_date, ppdm_unit_conversion.Row_expiry_date, ppdm_unit_conversion.Row_quality, ppdm_unit_conversion.From_uom_id)
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

func PatchPpdmUnitConversion(c *fiber.Ctx) error {
	id := c.Params("id")

	updatedFields := make(map[string]interface{})

	if err := c.BodyParser(&updatedFields); err != nil {
		return err
	}
	if _, ok := updatedFields["row_created_date"]; ok {
        return c.Status(400).SendString("Cannot update row_created_date field")
    }

	query := "update ppdm_unit_conversion set "
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
	query += " where from_uom_id = :id"

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

func DeletePpdmUnitConversion(c *fiber.Ctx) error {
	id := c.Params("id")
	var ppdm_unit_conversion dto.Ppdm_unit_conversion
	ppdm_unit_conversion.From_uom_id = id

	stmt, err := db.Prepare("delete from ppdm_unit_conversion where from_uom_id = :1")
	if err != nil {
		return err
	}

	rows, err := stmt.Exec(ppdm_unit_conversion.From_uom_id)
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


