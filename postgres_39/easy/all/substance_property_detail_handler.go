package apiv1

import (
	dto "github.com/DarrenMannuela/pt_gtn_bibliography/dto"
	"github.com/gofiber/fiber/v2"
	"strconv"
	"encoding/json"
)


func GetSubstancePropertyDetail(c *fiber.Ctx) error {
	rows, err := db.Query("select * from substance_property_detail")
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}
	defer rows.Close()
	var result []dto.Substance_property_detail

	for rows.Next() {
		var substance_property_detail dto.Substance_property_detail
		if err := rows.Scan(&substance_property_detail.Substance_id, &substance_property_detail.Property_obs_no, &substance_property_detail.Active_ind, &substance_property_detail.Average_value, &substance_property_detail.Average_value_ouom, &substance_property_detail.Average_value_uom, &substance_property_detail.Code_value, &substance_property_detail.Cost, &substance_property_detail.Currency_conversion, &substance_property_detail.Currency_ouom, &substance_property_detail.Currency_uom, &substance_property_detail.Date_format_desc, &substance_property_detail.Description, &substance_property_detail.Effective_date, &substance_property_detail.Expiry_date, &substance_property_detail.Max_date, &substance_property_detail.Max_value, &substance_property_detail.Max_value_ouom, &substance_property_detail.Max_value_uom, &substance_property_detail.Min_date, &substance_property_detail.Min_value, &substance_property_detail.Min_value_ouom, &substance_property_detail.Min_value_uom, &substance_property_detail.Ppdm_guid, &substance_property_detail.Reference_value, &substance_property_detail.Reference_value_type, &substance_property_detail.Reference_value_uom, &substance_property_detail.Remark, &substance_property_detail.Source, &substance_property_detail.Substance_property, &substance_property_detail.Row_changed_by, &substance_property_detail.Row_changed_date, &substance_property_detail.Row_created_by, &substance_property_detail.Row_created_date, &substance_property_detail.Row_effective_date, &substance_property_detail.Row_expiry_date, &substance_property_detail.Row_quality); err != nil {
			return err // Exit if we get an error
		}

		// Append substance_property_detail to result
		result = append(result, substance_property_detail)
	}
	// Return result in JSON format
	return c.JSON(result)
	 
}

func SetSubstancePropertyDetail(c *fiber.Ctx) error {
	var substance_property_detail dto.Substance_property_detail

	setDefaults(&substance_property_detail)

	if err := json.Unmarshal(c.Body(), &substance_property_detail); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	stmt, err := db.Prepare("insert into substance_property_detail values(:1, :2, :3, :4, :5, :6, :7, :8, :9, :10, :11, :12, :13, :14, :15, :16, :17, :18, :19, :20, :21, :22, :23, :24, :25, :26, :27, :28, :29, :30, :31, :32, :33, :34, :35, :36, :37)")
	if err != nil {
		return err
	}
	substance_property_detail.Row_created_date = formatDate(substance_property_detail.Row_created_date)
	substance_property_detail.Row_changed_date = formatDate(substance_property_detail.Row_changed_date)
	substance_property_detail.Date_format_desc = formatDateString(substance_property_detail.Date_format_desc)
	substance_property_detail.Effective_date = formatDateString(substance_property_detail.Effective_date)
	substance_property_detail.Expiry_date = formatDateString(substance_property_detail.Expiry_date)
	substance_property_detail.Max_date = formatDateString(substance_property_detail.Max_date)
	substance_property_detail.Min_date = formatDateString(substance_property_detail.Min_date)
	substance_property_detail.Row_effective_date = formatDateString(substance_property_detail.Row_effective_date)
	substance_property_detail.Row_expiry_date = formatDateString(substance_property_detail.Row_expiry_date)






	rows, err := stmt.Exec(substance_property_detail.Substance_id, substance_property_detail.Property_obs_no, substance_property_detail.Active_ind, substance_property_detail.Average_value, substance_property_detail.Average_value_ouom, substance_property_detail.Average_value_uom, substance_property_detail.Code_value, substance_property_detail.Cost, substance_property_detail.Currency_conversion, substance_property_detail.Currency_ouom, substance_property_detail.Currency_uom, substance_property_detail.Date_format_desc, substance_property_detail.Description, substance_property_detail.Effective_date, substance_property_detail.Expiry_date, substance_property_detail.Max_date, substance_property_detail.Max_value, substance_property_detail.Max_value_ouom, substance_property_detail.Max_value_uom, substance_property_detail.Min_date, substance_property_detail.Min_value, substance_property_detail.Min_value_ouom, substance_property_detail.Min_value_uom, substance_property_detail.Ppdm_guid, substance_property_detail.Reference_value, substance_property_detail.Reference_value_type, substance_property_detail.Reference_value_uom, substance_property_detail.Remark, substance_property_detail.Source, substance_property_detail.Substance_property, substance_property_detail.Row_changed_by, substance_property_detail.Row_changed_date, substance_property_detail.Row_created_by, substance_property_detail.Row_created_date, substance_property_detail.Row_effective_date, substance_property_detail.Row_expiry_date, substance_property_detail.Row_quality)
	
	if err != nil {
		return err
	}

	return c.JSON(rows)
}

func UpdateSubstancePropertyDetail(c *fiber.Ctx) error {
	var substance_property_detail dto.Substance_property_detail

	setDefaults(&substance_property_detail)

	if err := json.Unmarshal(c.Body(), &substance_property_detail); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	id := c.Params("id")
	substance_property_detail.Substance_id = id

    if substance_property_detail.Row_created_date != nil {
        return c.Status(400).SendString("Cannot update row_created_date")
    }

	stmt, err := db.Prepare("update substance_property_detail set property_obs_no = :1, active_ind = :2, average_value = :3, average_value_ouom = :4, average_value_uom = :5, code_value = :6, cost = :7, currency_conversion = :8, currency_ouom = :9, currency_uom = :10, date_format_desc = :11, description = :12, effective_date = :13, expiry_date = :14, max_date = :15, max_value = :16, max_value_ouom = :17, max_value_uom = :18, min_date = :19, min_value = :20, min_value_ouom = :21, min_value_uom = :22, ppdm_guid = :23, reference_value = :24, reference_value_type = :25, reference_value_uom = :26, remark = :27, source = :28, substance_property = :29, row_changed_by = :30, row_changed_date = :31, row_created_by = :32, row_effective_date = :33, row_expiry_date = :34, row_quality = :35 where substance_id = :37")
	if err != nil {
		return err
	}

	substance_property_detail.Row_changed_date = formatDate(substance_property_detail.Row_changed_date)
	substance_property_detail.Date_format_desc = formatDateString(substance_property_detail.Date_format_desc)
	substance_property_detail.Effective_date = formatDateString(substance_property_detail.Effective_date)
	substance_property_detail.Expiry_date = formatDateString(substance_property_detail.Expiry_date)
	substance_property_detail.Max_date = formatDateString(substance_property_detail.Max_date)
	substance_property_detail.Min_date = formatDateString(substance_property_detail.Min_date)
	substance_property_detail.Row_effective_date = formatDateString(substance_property_detail.Row_effective_date)
	substance_property_detail.Row_expiry_date = formatDateString(substance_property_detail.Row_expiry_date)






	rows, err := stmt.Exec(substance_property_detail.Property_obs_no, substance_property_detail.Active_ind, substance_property_detail.Average_value, substance_property_detail.Average_value_ouom, substance_property_detail.Average_value_uom, substance_property_detail.Code_value, substance_property_detail.Cost, substance_property_detail.Currency_conversion, substance_property_detail.Currency_ouom, substance_property_detail.Currency_uom, substance_property_detail.Date_format_desc, substance_property_detail.Description, substance_property_detail.Effective_date, substance_property_detail.Expiry_date, substance_property_detail.Max_date, substance_property_detail.Max_value, substance_property_detail.Max_value_ouom, substance_property_detail.Max_value_uom, substance_property_detail.Min_date, substance_property_detail.Min_value, substance_property_detail.Min_value_ouom, substance_property_detail.Min_value_uom, substance_property_detail.Ppdm_guid, substance_property_detail.Reference_value, substance_property_detail.Reference_value_type, substance_property_detail.Reference_value_uom, substance_property_detail.Remark, substance_property_detail.Source, substance_property_detail.Substance_property, substance_property_detail.Row_changed_by, substance_property_detail.Row_changed_date, substance_property_detail.Row_created_by, substance_property_detail.Row_effective_date, substance_property_detail.Row_expiry_date, substance_property_detail.Row_quality, substance_property_detail.Substance_id)
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

func PatchSubstancePropertyDetail(c *fiber.Ctx) error {
	id := c.Params("id")

	updatedFields := make(map[string]interface{})

	if err := c.BodyParser(&updatedFields); err != nil {
		return err
	}
	if _, ok := updatedFields["row_created_date"]; ok {
        return c.Status(400).SendString("Cannot update row_created_date field")
    }

	query := "update substance_property_detail set "
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
	query += " where substance_id = :id"

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

func DeleteSubstancePropertyDetail(c *fiber.Ctx) error {
	id := c.Params("id")
	var substance_property_detail dto.Substance_property_detail
	substance_property_detail.Substance_id = id

	stmt, err := db.Prepare("delete from substance_property_detail where substance_id = :1")
	if err != nil {
		return err
	}

	rows, err := stmt.Exec(substance_property_detail.Substance_id)
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


