package apiv1

import (
	dto "github.com/DarrenMannuela/pt_gtn_bibliography/dto"
	"github.com/gofiber/fiber/v2"
	"strconv"
	"encoding/json"
)


func GetSampleDescOther(c *fiber.Ctx) error {
	rows, err := db.Query("select * from sample_desc_other")
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}
	defer rows.Close()
	var result []dto.Sample_desc_other

	for rows.Next() {
		var sample_desc_other dto.Sample_desc_other
		if err := rows.Scan(&sample_desc_other.Sample_id, &sample_desc_other.Description_id, &sample_desc_other.Active_ind, &sample_desc_other.Average_ratio_value, &sample_desc_other.Average_value, &sample_desc_other.Average_value_ouom, &sample_desc_other.Average_value_uom, &sample_desc_other.Calculate_method_id, &sample_desc_other.Cost, &sample_desc_other.Currency_conversion, &sample_desc_other.Currency_ouom, &sample_desc_other.Currency_uom, &sample_desc_other.Date_format_desc, &sample_desc_other.Description, &sample_desc_other.Description_code, &sample_desc_other.Description_type, &sample_desc_other.Effective_date, &sample_desc_other.Expiry_date, &sample_desc_other.Max_date, &sample_desc_other.Max_ratio_value, &sample_desc_other.Max_value, &sample_desc_other.Max_value_ouom, &sample_desc_other.Max_value_uom, &sample_desc_other.Min_date, &sample_desc_other.Min_ratio_value, &sample_desc_other.Min_value, &sample_desc_other.Min_value_ouom, &sample_desc_other.Min_value_uom, &sample_desc_other.Ppdm_guid, &sample_desc_other.Reference_value, &sample_desc_other.Reference_value_ouom, &sample_desc_other.Reference_value_type, &sample_desc_other.Reference_value_uom, &sample_desc_other.Remark, &sample_desc_other.Reported_ratio, &sample_desc_other.Reported_ratio_ouom, &sample_desc_other.Source, &sample_desc_other.Row_changed_by, &sample_desc_other.Row_changed_date, &sample_desc_other.Row_created_by, &sample_desc_other.Row_created_date, &sample_desc_other.Row_effective_date, &sample_desc_other.Row_expiry_date, &sample_desc_other.Row_quality); err != nil {
			return err // Exit if we get an error
		}

		// Append sample_desc_other to result
		result = append(result, sample_desc_other)
	}
	// Return result in JSON format
	return c.JSON(result)
	 
}

func SetSampleDescOther(c *fiber.Ctx) error {
	var sample_desc_other dto.Sample_desc_other

	setDefaults(&sample_desc_other)

	if err := json.Unmarshal(c.Body(), &sample_desc_other); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	stmt, err := db.Prepare("insert into sample_desc_other values(:1, :2, :3, :4, :5, :6, :7, :8, :9, :10, :11, :12, :13, :14, :15, :16, :17, :18, :19, :20, :21, :22, :23, :24, :25, :26, :27, :28, :29, :30, :31, :32, :33, :34, :35, :36, :37, :38, :39, :40, :41, :42, :43, :44)")
	if err != nil {
		return err
	}
	sample_desc_other.Row_created_date = formatDate(sample_desc_other.Row_created_date)
	sample_desc_other.Row_changed_date = formatDate(sample_desc_other.Row_changed_date)
	sample_desc_other.Date_format_desc = formatDateString(sample_desc_other.Date_format_desc)
	sample_desc_other.Effective_date = formatDateString(sample_desc_other.Effective_date)
	sample_desc_other.Expiry_date = formatDateString(sample_desc_other.Expiry_date)
	sample_desc_other.Max_date = formatDateString(sample_desc_other.Max_date)
	sample_desc_other.Min_date = formatDateString(sample_desc_other.Min_date)
	sample_desc_other.Row_effective_date = formatDateString(sample_desc_other.Row_effective_date)
	sample_desc_other.Row_expiry_date = formatDateString(sample_desc_other.Row_expiry_date)






	rows, err := stmt.Exec(sample_desc_other.Sample_id, sample_desc_other.Description_id, sample_desc_other.Active_ind, sample_desc_other.Average_ratio_value, sample_desc_other.Average_value, sample_desc_other.Average_value_ouom, sample_desc_other.Average_value_uom, sample_desc_other.Calculate_method_id, sample_desc_other.Cost, sample_desc_other.Currency_conversion, sample_desc_other.Currency_ouom, sample_desc_other.Currency_uom, sample_desc_other.Date_format_desc, sample_desc_other.Description, sample_desc_other.Description_code, sample_desc_other.Description_type, sample_desc_other.Effective_date, sample_desc_other.Expiry_date, sample_desc_other.Max_date, sample_desc_other.Max_ratio_value, sample_desc_other.Max_value, sample_desc_other.Max_value_ouom, sample_desc_other.Max_value_uom, sample_desc_other.Min_date, sample_desc_other.Min_ratio_value, sample_desc_other.Min_value, sample_desc_other.Min_value_ouom, sample_desc_other.Min_value_uom, sample_desc_other.Ppdm_guid, sample_desc_other.Reference_value, sample_desc_other.Reference_value_ouom, sample_desc_other.Reference_value_type, sample_desc_other.Reference_value_uom, sample_desc_other.Remark, sample_desc_other.Reported_ratio, sample_desc_other.Reported_ratio_ouom, sample_desc_other.Source, sample_desc_other.Row_changed_by, sample_desc_other.Row_changed_date, sample_desc_other.Row_created_by, sample_desc_other.Row_created_date, sample_desc_other.Row_effective_date, sample_desc_other.Row_expiry_date, sample_desc_other.Row_quality)
	
	if err != nil {
		return err
	}

	return c.JSON(rows)
}

func UpdateSampleDescOther(c *fiber.Ctx) error {
	var sample_desc_other dto.Sample_desc_other

	setDefaults(&sample_desc_other)

	if err := json.Unmarshal(c.Body(), &sample_desc_other); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	id := c.Params("id")
	sample_desc_other.Sample_id = id

    if sample_desc_other.Row_created_date != nil {
        return c.Status(400).SendString("Cannot update row_created_date")
    }

	stmt, err := db.Prepare("update sample_desc_other set description_id = :1, active_ind = :2, average_ratio_value = :3, average_value = :4, average_value_ouom = :5, average_value_uom = :6, calculate_method_id = :7, cost = :8, currency_conversion = :9, currency_ouom = :10, currency_uom = :11, date_format_desc = :12, description = :13, description_code = :14, description_type = :15, effective_date = :16, expiry_date = :17, max_date = :18, max_ratio_value = :19, max_value = :20, max_value_ouom = :21, max_value_uom = :22, min_date = :23, min_ratio_value = :24, min_value = :25, min_value_ouom = :26, min_value_uom = :27, ppdm_guid = :28, reference_value = :29, reference_value_ouom = :30, reference_value_type = :31, reference_value_uom = :32, remark = :33, reported_ratio = :34, reported_ratio_ouom = :35, source = :36, row_changed_by = :37, row_changed_date = :38, row_created_by = :39, row_effective_date = :40, row_expiry_date = :41, row_quality = :42 where sample_id = :44")
	if err != nil {
		return err
	}

	sample_desc_other.Row_changed_date = formatDate(sample_desc_other.Row_changed_date)
	sample_desc_other.Date_format_desc = formatDateString(sample_desc_other.Date_format_desc)
	sample_desc_other.Effective_date = formatDateString(sample_desc_other.Effective_date)
	sample_desc_other.Expiry_date = formatDateString(sample_desc_other.Expiry_date)
	sample_desc_other.Max_date = formatDateString(sample_desc_other.Max_date)
	sample_desc_other.Min_date = formatDateString(sample_desc_other.Min_date)
	sample_desc_other.Row_effective_date = formatDateString(sample_desc_other.Row_effective_date)
	sample_desc_other.Row_expiry_date = formatDateString(sample_desc_other.Row_expiry_date)






	rows, err := stmt.Exec(sample_desc_other.Description_id, sample_desc_other.Active_ind, sample_desc_other.Average_ratio_value, sample_desc_other.Average_value, sample_desc_other.Average_value_ouom, sample_desc_other.Average_value_uom, sample_desc_other.Calculate_method_id, sample_desc_other.Cost, sample_desc_other.Currency_conversion, sample_desc_other.Currency_ouom, sample_desc_other.Currency_uom, sample_desc_other.Date_format_desc, sample_desc_other.Description, sample_desc_other.Description_code, sample_desc_other.Description_type, sample_desc_other.Effective_date, sample_desc_other.Expiry_date, sample_desc_other.Max_date, sample_desc_other.Max_ratio_value, sample_desc_other.Max_value, sample_desc_other.Max_value_ouom, sample_desc_other.Max_value_uom, sample_desc_other.Min_date, sample_desc_other.Min_ratio_value, sample_desc_other.Min_value, sample_desc_other.Min_value_ouom, sample_desc_other.Min_value_uom, sample_desc_other.Ppdm_guid, sample_desc_other.Reference_value, sample_desc_other.Reference_value_ouom, sample_desc_other.Reference_value_type, sample_desc_other.Reference_value_uom, sample_desc_other.Remark, sample_desc_other.Reported_ratio, sample_desc_other.Reported_ratio_ouom, sample_desc_other.Source, sample_desc_other.Row_changed_by, sample_desc_other.Row_changed_date, sample_desc_other.Row_created_by, sample_desc_other.Row_effective_date, sample_desc_other.Row_expiry_date, sample_desc_other.Row_quality, sample_desc_other.Sample_id)
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

func PatchSampleDescOther(c *fiber.Ctx) error {
	id := c.Params("id")

	updatedFields := make(map[string]interface{})

	if err := c.BodyParser(&updatedFields); err != nil {
		return err
	}
	if _, ok := updatedFields["row_created_date"]; ok {
        return c.Status(400).SendString("Cannot update row_created_date field")
    }

	query := "update sample_desc_other set "
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
	query += " where sample_id = :id"

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

func DeleteSampleDescOther(c *fiber.Ctx) error {
	id := c.Params("id")
	var sample_desc_other dto.Sample_desc_other
	sample_desc_other.Sample_id = id

	stmt, err := db.Prepare("delete from sample_desc_other where sample_id = :1")
	if err != nil {
		return err
	}

	rows, err := stmt.Exec(sample_desc_other.Sample_id)
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


