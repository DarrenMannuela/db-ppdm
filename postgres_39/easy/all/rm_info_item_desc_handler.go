package apiv1

import (
	dto "github.com/DarrenMannuela/pt_gtn_bibliography/dto"
	"github.com/gofiber/fiber/v2"
	"strconv"
	"encoding/json"
)


func GetRmInfoItemDesc(c *fiber.Ctx) error {
	rows, err := db.Query("select * from rm_info_item_desc")
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}
	defer rows.Close()
	var result []dto.Rm_info_item_desc

	for rows.Next() {
		var rm_info_item_desc dto.Rm_info_item_desc
		if err := rows.Scan(&rm_info_item_desc.Info_item_subtype, &rm_info_item_desc.Information_item_id, &rm_info_item_desc.Description_id, &rm_info_item_desc.Active_ind, &rm_info_item_desc.Average_value, &rm_info_item_desc.Average_value_ouom, &rm_info_item_desc.Average_value_uom, &rm_info_item_desc.Cost, &rm_info_item_desc.Currency_conversion, &rm_info_item_desc.Currency_ouom, &rm_info_item_desc.Currency_uom, &rm_info_item_desc.Date_format_desc, &rm_info_item_desc.Description, &rm_info_item_desc.Description_date, &rm_info_item_desc.Description_type, &rm_info_item_desc.Effective_date, &rm_info_item_desc.Expiry_date, &rm_info_item_desc.Max_date, &rm_info_item_desc.Max_value, &rm_info_item_desc.Max_value_ouom, &rm_info_item_desc.Max_value_uom, &rm_info_item_desc.Metadata_code, &rm_info_item_desc.Metadata_type, &rm_info_item_desc.Min_date, &rm_info_item_desc.Min_value, &rm_info_item_desc.Min_value_ouom, &rm_info_item_desc.Min_value_uom, &rm_info_item_desc.Ppdm_guid, &rm_info_item_desc.Remark, &rm_info_item_desc.Source, &rm_info_item_desc.Spec_desc, &rm_info_item_desc.Row_changed_by, &rm_info_item_desc.Row_changed_date, &rm_info_item_desc.Row_created_by, &rm_info_item_desc.Row_created_date, &rm_info_item_desc.Row_effective_date, &rm_info_item_desc.Row_expiry_date, &rm_info_item_desc.Row_quality); err != nil {
			return err // Exit if we get an error
		}

		// Append rm_info_item_desc to result
		result = append(result, rm_info_item_desc)
	}
	// Return result in JSON format
	return c.JSON(result)
	 
}

func SetRmInfoItemDesc(c *fiber.Ctx) error {
	var rm_info_item_desc dto.Rm_info_item_desc

	setDefaults(&rm_info_item_desc)

	if err := json.Unmarshal(c.Body(), &rm_info_item_desc); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	stmt, err := db.Prepare("insert into rm_info_item_desc values(:1, :2, :3, :4, :5, :6, :7, :8, :9, :10, :11, :12, :13, :14, :15, :16, :17, :18, :19, :20, :21, :22, :23, :24, :25, :26, :27, :28, :29, :30, :31, :32, :33, :34, :35, :36, :37, :38)")
	if err != nil {
		return err
	}
	rm_info_item_desc.Row_created_date = formatDate(rm_info_item_desc.Row_created_date)
	rm_info_item_desc.Row_changed_date = formatDate(rm_info_item_desc.Row_changed_date)
	rm_info_item_desc.Date_format_desc = formatDateString(rm_info_item_desc.Date_format_desc)
	rm_info_item_desc.Description_date = formatDateString(rm_info_item_desc.Description_date)
	rm_info_item_desc.Effective_date = formatDateString(rm_info_item_desc.Effective_date)
	rm_info_item_desc.Expiry_date = formatDateString(rm_info_item_desc.Expiry_date)
	rm_info_item_desc.Max_date = formatDateString(rm_info_item_desc.Max_date)
	rm_info_item_desc.Min_date = formatDateString(rm_info_item_desc.Min_date)
	rm_info_item_desc.Row_effective_date = formatDateString(rm_info_item_desc.Row_effective_date)
	rm_info_item_desc.Row_expiry_date = formatDateString(rm_info_item_desc.Row_expiry_date)






	rows, err := stmt.Exec(rm_info_item_desc.Info_item_subtype, rm_info_item_desc.Information_item_id, rm_info_item_desc.Description_id, rm_info_item_desc.Active_ind, rm_info_item_desc.Average_value, rm_info_item_desc.Average_value_ouom, rm_info_item_desc.Average_value_uom, rm_info_item_desc.Cost, rm_info_item_desc.Currency_conversion, rm_info_item_desc.Currency_ouom, rm_info_item_desc.Currency_uom, rm_info_item_desc.Date_format_desc, rm_info_item_desc.Description, rm_info_item_desc.Description_date, rm_info_item_desc.Description_type, rm_info_item_desc.Effective_date, rm_info_item_desc.Expiry_date, rm_info_item_desc.Max_date, rm_info_item_desc.Max_value, rm_info_item_desc.Max_value_ouom, rm_info_item_desc.Max_value_uom, rm_info_item_desc.Metadata_code, rm_info_item_desc.Metadata_type, rm_info_item_desc.Min_date, rm_info_item_desc.Min_value, rm_info_item_desc.Min_value_ouom, rm_info_item_desc.Min_value_uom, rm_info_item_desc.Ppdm_guid, rm_info_item_desc.Remark, rm_info_item_desc.Source, rm_info_item_desc.Spec_desc, rm_info_item_desc.Row_changed_by, rm_info_item_desc.Row_changed_date, rm_info_item_desc.Row_created_by, rm_info_item_desc.Row_created_date, rm_info_item_desc.Row_effective_date, rm_info_item_desc.Row_expiry_date, rm_info_item_desc.Row_quality)
	
	if err != nil {
		return err
	}

	return c.JSON(rows)
}

func UpdateRmInfoItemDesc(c *fiber.Ctx) error {
	var rm_info_item_desc dto.Rm_info_item_desc

	setDefaults(&rm_info_item_desc)

	if err := json.Unmarshal(c.Body(), &rm_info_item_desc); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	id := c.Params("id")
	rm_info_item_desc.Info_item_subtype = id

    if rm_info_item_desc.Row_created_date != nil {
        return c.Status(400).SendString("Cannot update row_created_date")
    }

	stmt, err := db.Prepare("update rm_info_item_desc set information_item_id = :1, description_id = :2, active_ind = :3, average_value = :4, average_value_ouom = :5, average_value_uom = :6, cost = :7, currency_conversion = :8, currency_ouom = :9, currency_uom = :10, date_format_desc = :11, description = :12, description_date = :13, description_type = :14, effective_date = :15, expiry_date = :16, max_date = :17, max_value = :18, max_value_ouom = :19, max_value_uom = :20, metadata_code = :21, metadata_type = :22, min_date = :23, min_value = :24, min_value_ouom = :25, min_value_uom = :26, ppdm_guid = :27, remark = :28, source = :29, spec_desc = :30, row_changed_by = :31, row_changed_date = :32, row_created_by = :33, row_effective_date = :34, row_expiry_date = :35, row_quality = :36 where info_item_subtype = :38")
	if err != nil {
		return err
	}

	rm_info_item_desc.Row_changed_date = formatDate(rm_info_item_desc.Row_changed_date)
	rm_info_item_desc.Date_format_desc = formatDateString(rm_info_item_desc.Date_format_desc)
	rm_info_item_desc.Description_date = formatDateString(rm_info_item_desc.Description_date)
	rm_info_item_desc.Effective_date = formatDateString(rm_info_item_desc.Effective_date)
	rm_info_item_desc.Expiry_date = formatDateString(rm_info_item_desc.Expiry_date)
	rm_info_item_desc.Max_date = formatDateString(rm_info_item_desc.Max_date)
	rm_info_item_desc.Min_date = formatDateString(rm_info_item_desc.Min_date)
	rm_info_item_desc.Row_effective_date = formatDateString(rm_info_item_desc.Row_effective_date)
	rm_info_item_desc.Row_expiry_date = formatDateString(rm_info_item_desc.Row_expiry_date)






	rows, err := stmt.Exec(rm_info_item_desc.Information_item_id, rm_info_item_desc.Description_id, rm_info_item_desc.Active_ind, rm_info_item_desc.Average_value, rm_info_item_desc.Average_value_ouom, rm_info_item_desc.Average_value_uom, rm_info_item_desc.Cost, rm_info_item_desc.Currency_conversion, rm_info_item_desc.Currency_ouom, rm_info_item_desc.Currency_uom, rm_info_item_desc.Date_format_desc, rm_info_item_desc.Description, rm_info_item_desc.Description_date, rm_info_item_desc.Description_type, rm_info_item_desc.Effective_date, rm_info_item_desc.Expiry_date, rm_info_item_desc.Max_date, rm_info_item_desc.Max_value, rm_info_item_desc.Max_value_ouom, rm_info_item_desc.Max_value_uom, rm_info_item_desc.Metadata_code, rm_info_item_desc.Metadata_type, rm_info_item_desc.Min_date, rm_info_item_desc.Min_value, rm_info_item_desc.Min_value_ouom, rm_info_item_desc.Min_value_uom, rm_info_item_desc.Ppdm_guid, rm_info_item_desc.Remark, rm_info_item_desc.Source, rm_info_item_desc.Spec_desc, rm_info_item_desc.Row_changed_by, rm_info_item_desc.Row_changed_date, rm_info_item_desc.Row_created_by, rm_info_item_desc.Row_effective_date, rm_info_item_desc.Row_expiry_date, rm_info_item_desc.Row_quality, rm_info_item_desc.Info_item_subtype)
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

func PatchRmInfoItemDesc(c *fiber.Ctx) error {
	id := c.Params("id")

	updatedFields := make(map[string]interface{})

	if err := c.BodyParser(&updatedFields); err != nil {
		return err
	}
	if _, ok := updatedFields["row_created_date"]; ok {
        return c.Status(400).SendString("Cannot update row_created_date field")
    }

	query := "update rm_info_item_desc set "
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
		} else if key == "date_format_desc" || key == "description_date" || key == "effective_date" || key == "expiry_date" || key == "max_date" || key == "min_date" || key == "row_effective_date" || key == "row_expiry_date"      {
			if date, ok := value.(string); ok && len(date) > 0 {
				formattedDate := formatDateString(&date)
				value = formattedDate
			}
		}
		values = append(values, value)
	}
	query += " where info_item_subtype = :id"

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

func DeleteRmInfoItemDesc(c *fiber.Ctx) error {
	id := c.Params("id")
	var rm_info_item_desc dto.Rm_info_item_desc
	rm_info_item_desc.Info_item_subtype = id

	stmt, err := db.Prepare("delete from rm_info_item_desc where info_item_subtype = :1")
	if err != nil {
		return err
	}

	rows, err := stmt.Exec(rm_info_item_desc.Info_item_subtype)
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


