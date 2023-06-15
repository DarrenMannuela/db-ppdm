package apiv1

import (
	dto "github.com/DarrenMannuela/pt_gtn_bibliography/dto"
	"github.com/gofiber/fiber/v2"
	"strconv"
	"encoding/json"
)


func GetSfDescription(c *fiber.Ctx) error {
	rows, err := db.Query("select * from sf_description")
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}
	defer rows.Close()
	var result []dto.Sf_description

	for rows.Next() {
		var sf_description dto.Sf_description
		if err := rows.Scan(&sf_description.Sf_subtype, &sf_description.Support_facility_id, &sf_description.Description_obs_no, &sf_description.Active_ind, &sf_description.Cost, &sf_description.Cost_ouom, &sf_description.Cost_uom, &sf_description.Currency_conversion, &sf_description.Date_format_desc, &sf_description.Description, &sf_description.Desc_type, &sf_description.Desc_value, &sf_description.Desc_value_code, &sf_description.Desc_value_ouom, &sf_description.Desc_value_uom, &sf_description.Effective_date, &sf_description.Expiry_date, &sf_description.Max_value, &sf_description.Max_value_ouom, &sf_description.Max_value_uom, &sf_description.Min_value, &sf_description.Min_value_ouom, &sf_description.Min_value_uom, &sf_description.Ppdm_guid, &sf_description.Remark, &sf_description.Source, &sf_description.Row_changed_by, &sf_description.Row_changed_date, &sf_description.Row_created_by, &sf_description.Row_created_date, &sf_description.Row_effective_date, &sf_description.Row_expiry_date, &sf_description.Row_quality); err != nil {
			return err // Exit if we get an error
		}

		// Append sf_description to result
		result = append(result, sf_description)
	}
	// Return result in JSON format
	return c.JSON(result)
	 
}

func SetSfDescription(c *fiber.Ctx) error {
	var sf_description dto.Sf_description

	setDefaults(&sf_description)

	if err := json.Unmarshal(c.Body(), &sf_description); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	stmt, err := db.Prepare("insert into sf_description values(:1, :2, :3, :4, :5, :6, :7, :8, :9, :10, :11, :12, :13, :14, :15, :16, :17, :18, :19, :20, :21, :22, :23, :24, :25, :26, :27, :28, :29, :30, :31, :32, :33)")
	if err != nil {
		return err
	}
	sf_description.Row_created_date = formatDate(sf_description.Row_created_date)
	sf_description.Row_changed_date = formatDate(sf_description.Row_changed_date)
	sf_description.Date_format_desc = formatDateString(sf_description.Date_format_desc)
	sf_description.Effective_date = formatDateString(sf_description.Effective_date)
	sf_description.Expiry_date = formatDateString(sf_description.Expiry_date)
	sf_description.Row_effective_date = formatDateString(sf_description.Row_effective_date)
	sf_description.Row_expiry_date = formatDateString(sf_description.Row_expiry_date)






	rows, err := stmt.Exec(sf_description.Sf_subtype, sf_description.Support_facility_id, sf_description.Description_obs_no, sf_description.Active_ind, sf_description.Cost, sf_description.Cost_ouom, sf_description.Cost_uom, sf_description.Currency_conversion, sf_description.Date_format_desc, sf_description.Description, sf_description.Desc_type, sf_description.Desc_value, sf_description.Desc_value_code, sf_description.Desc_value_ouom, sf_description.Desc_value_uom, sf_description.Effective_date, sf_description.Expiry_date, sf_description.Max_value, sf_description.Max_value_ouom, sf_description.Max_value_uom, sf_description.Min_value, sf_description.Min_value_ouom, sf_description.Min_value_uom, sf_description.Ppdm_guid, sf_description.Remark, sf_description.Source, sf_description.Row_changed_by, sf_description.Row_changed_date, sf_description.Row_created_by, sf_description.Row_created_date, sf_description.Row_effective_date, sf_description.Row_expiry_date, sf_description.Row_quality)
	
	if err != nil {
		return err
	}

	return c.JSON(rows)
}

func UpdateSfDescription(c *fiber.Ctx) error {
	var sf_description dto.Sf_description

	setDefaults(&sf_description)

	if err := json.Unmarshal(c.Body(), &sf_description); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	id := c.Params("id")
	sf_description.Sf_subtype = id

    if sf_description.Row_created_date != nil {
        return c.Status(400).SendString("Cannot update row_created_date")
    }

	stmt, err := db.Prepare("update sf_description set support_facility_id = :1, description_obs_no = :2, active_ind = :3, cost = :4, cost_ouom = :5, cost_uom = :6, currency_conversion = :7, date_format_desc = :8, description = :9, desc_type = :10, desc_value = :11, desc_value_code = :12, desc_value_ouom = :13, desc_value_uom = :14, effective_date = :15, expiry_date = :16, max_value = :17, max_value_ouom = :18, max_value_uom = :19, min_value = :20, min_value_ouom = :21, min_value_uom = :22, ppdm_guid = :23, remark = :24, source = :25, row_changed_by = :26, row_changed_date = :27, row_created_by = :28, row_effective_date = :29, row_expiry_date = :30, row_quality = :31 where sf_subtype = :33")
	if err != nil {
		return err
	}

	sf_description.Row_changed_date = formatDate(sf_description.Row_changed_date)
	sf_description.Date_format_desc = formatDateString(sf_description.Date_format_desc)
	sf_description.Effective_date = formatDateString(sf_description.Effective_date)
	sf_description.Expiry_date = formatDateString(sf_description.Expiry_date)
	sf_description.Row_effective_date = formatDateString(sf_description.Row_effective_date)
	sf_description.Row_expiry_date = formatDateString(sf_description.Row_expiry_date)






	rows, err := stmt.Exec(sf_description.Support_facility_id, sf_description.Description_obs_no, sf_description.Active_ind, sf_description.Cost, sf_description.Cost_ouom, sf_description.Cost_uom, sf_description.Currency_conversion, sf_description.Date_format_desc, sf_description.Description, sf_description.Desc_type, sf_description.Desc_value, sf_description.Desc_value_code, sf_description.Desc_value_ouom, sf_description.Desc_value_uom, sf_description.Effective_date, sf_description.Expiry_date, sf_description.Max_value, sf_description.Max_value_ouom, sf_description.Max_value_uom, sf_description.Min_value, sf_description.Min_value_ouom, sf_description.Min_value_uom, sf_description.Ppdm_guid, sf_description.Remark, sf_description.Source, sf_description.Row_changed_by, sf_description.Row_changed_date, sf_description.Row_created_by, sf_description.Row_effective_date, sf_description.Row_expiry_date, sf_description.Row_quality, sf_description.Sf_subtype)
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

func PatchSfDescription(c *fiber.Ctx) error {
	id := c.Params("id")

	updatedFields := make(map[string]interface{})

	if err := c.BodyParser(&updatedFields); err != nil {
		return err
	}
	if _, ok := updatedFields["row_created_date"]; ok {
        return c.Status(400).SendString("Cannot update row_created_date field")
    }

	query := "update sf_description set "
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
	query += " where sf_subtype = :id"

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

func DeleteSfDescription(c *fiber.Ctx) error {
	id := c.Params("id")
	var sf_description dto.Sf_description
	sf_description.Sf_subtype = id

	stmt, err := db.Prepare("delete from sf_description where sf_subtype = :1")
	if err != nil {
		return err
	}

	rows, err := stmt.Exec(sf_description.Sf_subtype)
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


