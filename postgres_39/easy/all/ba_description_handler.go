package apiv1

import (
	dto "github.com/DarrenMannuela/pt_gtn_bibliography/dto"
	"github.com/gofiber/fiber/v2"
	"strconv"
	"encoding/json"
)


func GetBaDescription(c *fiber.Ctx) error {
	rows, err := db.Query("select * from ba_description")
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}
	defer rows.Close()
	var result []dto.Ba_description

	for rows.Next() {
		var ba_description dto.Ba_description
		if err := rows.Scan(&ba_description.Business_associate_id, &ba_description.Description_id, &ba_description.Active_ind, &ba_description.Average_value, &ba_description.Average_value_ouom, &ba_description.Average_value_uom, &ba_description.Ba_desc_code, &ba_description.Ba_desc_type, &ba_description.Cost, &ba_description.Currency_conversion, &ba_description.Currency_ouom, &ba_description.Currency_uom, &ba_description.Description, &ba_description.Effective_date, &ba_description.Expiry_date, &ba_description.Max_date, &ba_description.Max_value, &ba_description.Max_value_ouom, &ba_description.Max_value_uom, &ba_description.Min_date, &ba_description.Min_value, &ba_description.Min_value_ouom, &ba_description.Min_value_uom, &ba_description.Ppdm_guid, &ba_description.Reference_value, &ba_description.Reference_value_ouom, &ba_description.Reference_value_type, &ba_description.Reference_value_uom, &ba_description.Remark, &ba_description.Source, &ba_description.Row_changed_by, &ba_description.Row_changed_date, &ba_description.Row_created_by, &ba_description.Row_created_date, &ba_description.Row_effective_date, &ba_description.Row_expiry_date, &ba_description.Row_quality); err != nil {
			return err // Exit if we get an error
		}

		// Append ba_description to result
		result = append(result, ba_description)
	}
	// Return result in JSON format
	return c.JSON(result)
	 
}

func SetBaDescription(c *fiber.Ctx) error {
	var ba_description dto.Ba_description

	setDefaults(&ba_description)

	if err := json.Unmarshal(c.Body(), &ba_description); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	stmt, err := db.Prepare("insert into ba_description values(:1, :2, :3, :4, :5, :6, :7, :8, :9, :10, :11, :12, :13, :14, :15, :16, :17, :18, :19, :20, :21, :22, :23, :24, :25, :26, :27, :28, :29, :30, :31, :32, :33, :34, :35, :36, :37)")
	if err != nil {
		return err
	}
	ba_description.Row_created_date = formatDate(ba_description.Row_created_date)
	ba_description.Row_changed_date = formatDate(ba_description.Row_changed_date)
	ba_description.Effective_date = formatDateString(ba_description.Effective_date)
	ba_description.Expiry_date = formatDateString(ba_description.Expiry_date)
	ba_description.Max_date = formatDateString(ba_description.Max_date)
	ba_description.Min_date = formatDateString(ba_description.Min_date)
	ba_description.Row_effective_date = formatDateString(ba_description.Row_effective_date)
	ba_description.Row_expiry_date = formatDateString(ba_description.Row_expiry_date)






	rows, err := stmt.Exec(ba_description.Business_associate_id, ba_description.Description_id, ba_description.Active_ind, ba_description.Average_value, ba_description.Average_value_ouom, ba_description.Average_value_uom, ba_description.Ba_desc_code, ba_description.Ba_desc_type, ba_description.Cost, ba_description.Currency_conversion, ba_description.Currency_ouom, ba_description.Currency_uom, ba_description.Description, ba_description.Effective_date, ba_description.Expiry_date, ba_description.Max_date, ba_description.Max_value, ba_description.Max_value_ouom, ba_description.Max_value_uom, ba_description.Min_date, ba_description.Min_value, ba_description.Min_value_ouom, ba_description.Min_value_uom, ba_description.Ppdm_guid, ba_description.Reference_value, ba_description.Reference_value_ouom, ba_description.Reference_value_type, ba_description.Reference_value_uom, ba_description.Remark, ba_description.Source, ba_description.Row_changed_by, ba_description.Row_changed_date, ba_description.Row_created_by, ba_description.Row_created_date, ba_description.Row_effective_date, ba_description.Row_expiry_date, ba_description.Row_quality)
	
	if err != nil {
		return err
	}

	return c.JSON(rows)
}

func UpdateBaDescription(c *fiber.Ctx) error {
	var ba_description dto.Ba_description

	setDefaults(&ba_description)

	if err := json.Unmarshal(c.Body(), &ba_description); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	id := c.Params("id")
	ba_description.Business_associate_id = id

    if ba_description.Row_created_date != nil {
        return c.Status(400).SendString("Cannot update row_created_date")
    }

	stmt, err := db.Prepare("update ba_description set description_id = :1, active_ind = :2, average_value = :3, average_value_ouom = :4, average_value_uom = :5, ba_desc_code = :6, ba_desc_type = :7, cost = :8, currency_conversion = :9, currency_ouom = :10, currency_uom = :11, description = :12, effective_date = :13, expiry_date = :14, max_date = :15, max_value = :16, max_value_ouom = :17, max_value_uom = :18, min_date = :19, min_value = :20, min_value_ouom = :21, min_value_uom = :22, ppdm_guid = :23, reference_value = :24, reference_value_ouom = :25, reference_value_type = :26, reference_value_uom = :27, remark = :28, source = :29, row_changed_by = :30, row_changed_date = :31, row_created_by = :32, row_effective_date = :33, row_expiry_date = :34, row_quality = :35 where business_associate_id = :37")
	if err != nil {
		return err
	}

	ba_description.Row_changed_date = formatDate(ba_description.Row_changed_date)
	ba_description.Effective_date = formatDateString(ba_description.Effective_date)
	ba_description.Expiry_date = formatDateString(ba_description.Expiry_date)
	ba_description.Max_date = formatDateString(ba_description.Max_date)
	ba_description.Min_date = formatDateString(ba_description.Min_date)
	ba_description.Row_effective_date = formatDateString(ba_description.Row_effective_date)
	ba_description.Row_expiry_date = formatDateString(ba_description.Row_expiry_date)






	rows, err := stmt.Exec(ba_description.Description_id, ba_description.Active_ind, ba_description.Average_value, ba_description.Average_value_ouom, ba_description.Average_value_uom, ba_description.Ba_desc_code, ba_description.Ba_desc_type, ba_description.Cost, ba_description.Currency_conversion, ba_description.Currency_ouom, ba_description.Currency_uom, ba_description.Description, ba_description.Effective_date, ba_description.Expiry_date, ba_description.Max_date, ba_description.Max_value, ba_description.Max_value_ouom, ba_description.Max_value_uom, ba_description.Min_date, ba_description.Min_value, ba_description.Min_value_ouom, ba_description.Min_value_uom, ba_description.Ppdm_guid, ba_description.Reference_value, ba_description.Reference_value_ouom, ba_description.Reference_value_type, ba_description.Reference_value_uom, ba_description.Remark, ba_description.Source, ba_description.Row_changed_by, ba_description.Row_changed_date, ba_description.Row_created_by, ba_description.Row_effective_date, ba_description.Row_expiry_date, ba_description.Row_quality, ba_description.Business_associate_id)
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

func PatchBaDescription(c *fiber.Ctx) error {
	id := c.Params("id")

	updatedFields := make(map[string]interface{})

	if err := c.BodyParser(&updatedFields); err != nil {
		return err
	}
	if _, ok := updatedFields["row_created_date"]; ok {
        return c.Status(400).SendString("Cannot update row_created_date field")
    }

	query := "update ba_description set "
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
		} else if key == "effective_date" || key == "expiry_date" || key == "max_date" || key == "min_date" || key == "row_effective_date" || key == "row_expiry_date"      {
			if date, ok := value.(string); ok && len(date) > 0 {
				formattedDate := formatDateString(&date)
				value = formattedDate
			}
		}
		values = append(values, value)
	}
	query += " where business_associate_id = :id"

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

func DeleteBaDescription(c *fiber.Ctx) error {
	id := c.Params("id")
	var ba_description dto.Ba_description
	ba_description.Business_associate_id = id

	stmt, err := db.Prepare("delete from ba_description where business_associate_id = :1")
	if err != nil {
		return err
	}

	rows, err := stmt.Exec(ba_description.Business_associate_id)
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


