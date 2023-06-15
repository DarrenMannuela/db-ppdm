package apiv1

import (
	dto "github.com/DarrenMannuela/pt_gtn_bibliography/dto"
	"github.com/gofiber/fiber/v2"
	"strconv"
	"encoding/json"
)


func GetApplicDesc(c *fiber.Ctx) error {
	rows, err := db.Query("select * from applic_desc")
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}
	defer rows.Close()
	var result []dto.Applic_desc

	for rows.Next() {
		var applic_desc dto.Applic_desc
		if err := rows.Scan(&applic_desc.Application_id, &applic_desc.Description_id, &applic_desc.Active_ind, &applic_desc.Application_type, &applic_desc.Applic_desc_type, &applic_desc.Contact_ba_id, &applic_desc.Currency_conversion, &applic_desc.Currency_ouom, &applic_desc.Currency_uom, &applic_desc.Date_format_desc, &applic_desc.Description, &applic_desc.Desc_date, &applic_desc.Desc_value, &applic_desc.Desc_value_ouom, &applic_desc.Desc_value_uom, &applic_desc.Effective_date, &applic_desc.Expiry_date, &applic_desc.Max_cost, &applic_desc.Max_date, &applic_desc.Max_value, &applic_desc.Max_value_ouom, &applic_desc.Max_value_uom, &applic_desc.Min_cost, &applic_desc.Min_date, &applic_desc.Min_value, &applic_desc.Min_value_ouom, &applic_desc.Min_value_uom, &applic_desc.Ppdm_guid, &applic_desc.Remark, &applic_desc.Source, &applic_desc.Row_changed_by, &applic_desc.Row_changed_date, &applic_desc.Row_created_by, &applic_desc.Row_created_date, &applic_desc.Row_effective_date, &applic_desc.Row_expiry_date, &applic_desc.Row_quality); err != nil {
			return err // Exit if we get an error
		}

		// Append applic_desc to result
		result = append(result, applic_desc)
	}
	// Return result in JSON format
	return c.JSON(result)
	 
}

func SetApplicDesc(c *fiber.Ctx) error {
	var applic_desc dto.Applic_desc

	setDefaults(&applic_desc)

	if err := json.Unmarshal(c.Body(), &applic_desc); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	stmt, err := db.Prepare("insert into applic_desc values(:1, :2, :3, :4, :5, :6, :7, :8, :9, :10, :11, :12, :13, :14, :15, :16, :17, :18, :19, :20, :21, :22, :23, :24, :25, :26, :27, :28, :29, :30, :31, :32, :33, :34, :35, :36, :37)")
	if err != nil {
		return err
	}
	applic_desc.Row_created_date = formatDate(applic_desc.Row_created_date)
	applic_desc.Row_changed_date = formatDate(applic_desc.Row_changed_date)
	applic_desc.Date_format_desc = formatDateString(applic_desc.Date_format_desc)
	applic_desc.Desc_date = formatDateString(applic_desc.Desc_date)
	applic_desc.Effective_date = formatDateString(applic_desc.Effective_date)
	applic_desc.Expiry_date = formatDateString(applic_desc.Expiry_date)
	applic_desc.Max_date = formatDateString(applic_desc.Max_date)
	applic_desc.Min_date = formatDateString(applic_desc.Min_date)
	applic_desc.Row_effective_date = formatDateString(applic_desc.Row_effective_date)
	applic_desc.Row_expiry_date = formatDateString(applic_desc.Row_expiry_date)






	rows, err := stmt.Exec(applic_desc.Application_id, applic_desc.Description_id, applic_desc.Active_ind, applic_desc.Application_type, applic_desc.Applic_desc_type, applic_desc.Contact_ba_id, applic_desc.Currency_conversion, applic_desc.Currency_ouom, applic_desc.Currency_uom, applic_desc.Date_format_desc, applic_desc.Description, applic_desc.Desc_date, applic_desc.Desc_value, applic_desc.Desc_value_ouom, applic_desc.Desc_value_uom, applic_desc.Effective_date, applic_desc.Expiry_date, applic_desc.Max_cost, applic_desc.Max_date, applic_desc.Max_value, applic_desc.Max_value_ouom, applic_desc.Max_value_uom, applic_desc.Min_cost, applic_desc.Min_date, applic_desc.Min_value, applic_desc.Min_value_ouom, applic_desc.Min_value_uom, applic_desc.Ppdm_guid, applic_desc.Remark, applic_desc.Source, applic_desc.Row_changed_by, applic_desc.Row_changed_date, applic_desc.Row_created_by, applic_desc.Row_created_date, applic_desc.Row_effective_date, applic_desc.Row_expiry_date, applic_desc.Row_quality)
	
	if err != nil {
		return err
	}

	return c.JSON(rows)
}

func UpdateApplicDesc(c *fiber.Ctx) error {
	var applic_desc dto.Applic_desc

	setDefaults(&applic_desc)

	if err := json.Unmarshal(c.Body(), &applic_desc); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	id := c.Params("id")
	applic_desc.Application_id = id

    if applic_desc.Row_created_date != nil {
        return c.Status(400).SendString("Cannot update row_created_date")
    }

	stmt, err := db.Prepare("update applic_desc set description_id = :1, active_ind = :2, application_type = :3, applic_desc_type = :4, contact_ba_id = :5, currency_conversion = :6, currency_ouom = :7, currency_uom = :8, date_format_desc = :9, description = :10, desc_date = :11, desc_value = :12, desc_value_ouom = :13, desc_value_uom = :14, effective_date = :15, expiry_date = :16, max_cost = :17, max_date = :18, max_value = :19, max_value_ouom = :20, max_value_uom = :21, min_cost = :22, min_date = :23, min_value = :24, min_value_ouom = :25, min_value_uom = :26, ppdm_guid = :27, remark = :28, source = :29, row_changed_by = :30, row_changed_date = :31, row_created_by = :32, row_effective_date = :33, row_expiry_date = :34, row_quality = :35 where application_id = :37")
	if err != nil {
		return err
	}

	applic_desc.Row_changed_date = formatDate(applic_desc.Row_changed_date)
	applic_desc.Date_format_desc = formatDateString(applic_desc.Date_format_desc)
	applic_desc.Desc_date = formatDateString(applic_desc.Desc_date)
	applic_desc.Effective_date = formatDateString(applic_desc.Effective_date)
	applic_desc.Expiry_date = formatDateString(applic_desc.Expiry_date)
	applic_desc.Max_date = formatDateString(applic_desc.Max_date)
	applic_desc.Min_date = formatDateString(applic_desc.Min_date)
	applic_desc.Row_effective_date = formatDateString(applic_desc.Row_effective_date)
	applic_desc.Row_expiry_date = formatDateString(applic_desc.Row_expiry_date)






	rows, err := stmt.Exec(applic_desc.Description_id, applic_desc.Active_ind, applic_desc.Application_type, applic_desc.Applic_desc_type, applic_desc.Contact_ba_id, applic_desc.Currency_conversion, applic_desc.Currency_ouom, applic_desc.Currency_uom, applic_desc.Date_format_desc, applic_desc.Description, applic_desc.Desc_date, applic_desc.Desc_value, applic_desc.Desc_value_ouom, applic_desc.Desc_value_uom, applic_desc.Effective_date, applic_desc.Expiry_date, applic_desc.Max_cost, applic_desc.Max_date, applic_desc.Max_value, applic_desc.Max_value_ouom, applic_desc.Max_value_uom, applic_desc.Min_cost, applic_desc.Min_date, applic_desc.Min_value, applic_desc.Min_value_ouom, applic_desc.Min_value_uom, applic_desc.Ppdm_guid, applic_desc.Remark, applic_desc.Source, applic_desc.Row_changed_by, applic_desc.Row_changed_date, applic_desc.Row_created_by, applic_desc.Row_effective_date, applic_desc.Row_expiry_date, applic_desc.Row_quality, applic_desc.Application_id)
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

func PatchApplicDesc(c *fiber.Ctx) error {
	id := c.Params("id")

	updatedFields := make(map[string]interface{})

	if err := c.BodyParser(&updatedFields); err != nil {
		return err
	}
	if _, ok := updatedFields["row_created_date"]; ok {
        return c.Status(400).SendString("Cannot update row_created_date field")
    }

	query := "update applic_desc set "
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
		} else if key == "date_format_desc" || key == "desc_date" || key == "effective_date" || key == "expiry_date" || key == "max_date" || key == "min_date" || key == "row_effective_date" || key == "row_expiry_date"      {
			if date, ok := value.(string); ok && len(date) > 0 {
				formattedDate := formatDateString(&date)
				value = formattedDate
			}
		}
		values = append(values, value)
	}
	query += " where application_id = :id"

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

func DeleteApplicDesc(c *fiber.Ctx) error {
	id := c.Params("id")
	var applic_desc dto.Applic_desc
	applic_desc.Application_id = id

	stmt, err := db.Prepare("delete from applic_desc where application_id = :1")
	if err != nil {
		return err
	}

	rows, err := stmt.Exec(applic_desc.Application_id)
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


