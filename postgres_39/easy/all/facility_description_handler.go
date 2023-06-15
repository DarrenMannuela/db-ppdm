package apiv1

import (
	dto "github.com/DarrenMannuela/pt_gtn_bibliography/dto"
	"github.com/gofiber/fiber/v2"
	"strconv"
	"encoding/json"
)


func GetFacilityDescription(c *fiber.Ctx) error {
	rows, err := db.Query("select * from facility_description")
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}
	defer rows.Close()
	var result []dto.Facility_description

	for rows.Next() {
		var facility_description dto.Facility_description
		if err := rows.Scan(&facility_description.Facility_id, &facility_description.Facility_type, &facility_description.Spec_id, &facility_description.Active_ind, &facility_description.Average_value, &facility_description.Average_value_ouom, &facility_description.Average_value_uom, &facility_description.Cost, &facility_description.Currency_conversion, &facility_description.Currency_ouom, &facility_description.Currency_uom, &facility_description.Date_format_desc, &facility_description.Effective_date, &facility_description.Expiry_date, &facility_description.Max_date, &facility_description.Max_value, &facility_description.Max_value_ouom, &facility_description.Max_value_uom, &facility_description.Min_date, &facility_description.Min_value, &facility_description.Min_value_ouom, &facility_description.Min_value_uom, &facility_description.Ppdm_guid, &facility_description.Reference_value, &facility_description.Reference_value_ouom, &facility_description.Reference_value_type, &facility_description.Reference_value_uom, &facility_description.Remark, &facility_description.Source, &facility_description.Spec_code, &facility_description.Spec_desc, &facility_description.Spec_type, &facility_description.Row_changed_by, &facility_description.Row_changed_date, &facility_description.Row_created_by, &facility_description.Row_created_date, &facility_description.Row_effective_date, &facility_description.Row_expiry_date, &facility_description.Row_quality); err != nil {
			return err // Exit if we get an error
		}

		// Append facility_description to result
		result = append(result, facility_description)
	}
	// Return result in JSON format
	return c.JSON(result)
	 
}

func SetFacilityDescription(c *fiber.Ctx) error {
	var facility_description dto.Facility_description

	setDefaults(&facility_description)

	if err := json.Unmarshal(c.Body(), &facility_description); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	stmt, err := db.Prepare("insert into facility_description values(:1, :2, :3, :4, :5, :6, :7, :8, :9, :10, :11, :12, :13, :14, :15, :16, :17, :18, :19, :20, :21, :22, :23, :24, :25, :26, :27, :28, :29, :30, :31, :32, :33, :34, :35, :36, :37, :38, :39)")
	if err != nil {
		return err
	}
	facility_description.Row_created_date = formatDate(facility_description.Row_created_date)
	facility_description.Row_changed_date = formatDate(facility_description.Row_changed_date)
	facility_description.Date_format_desc = formatDateString(facility_description.Date_format_desc)
	facility_description.Effective_date = formatDateString(facility_description.Effective_date)
	facility_description.Expiry_date = formatDateString(facility_description.Expiry_date)
	facility_description.Max_date = formatDateString(facility_description.Max_date)
	facility_description.Min_date = formatDateString(facility_description.Min_date)
	facility_description.Row_effective_date = formatDateString(facility_description.Row_effective_date)
	facility_description.Row_expiry_date = formatDateString(facility_description.Row_expiry_date)






	rows, err := stmt.Exec(facility_description.Facility_id, facility_description.Facility_type, facility_description.Spec_id, facility_description.Active_ind, facility_description.Average_value, facility_description.Average_value_ouom, facility_description.Average_value_uom, facility_description.Cost, facility_description.Currency_conversion, facility_description.Currency_ouom, facility_description.Currency_uom, facility_description.Date_format_desc, facility_description.Effective_date, facility_description.Expiry_date, facility_description.Max_date, facility_description.Max_value, facility_description.Max_value_ouom, facility_description.Max_value_uom, facility_description.Min_date, facility_description.Min_value, facility_description.Min_value_ouom, facility_description.Min_value_uom, facility_description.Ppdm_guid, facility_description.Reference_value, facility_description.Reference_value_ouom, facility_description.Reference_value_type, facility_description.Reference_value_uom, facility_description.Remark, facility_description.Source, facility_description.Spec_code, facility_description.Spec_desc, facility_description.Spec_type, facility_description.Row_changed_by, facility_description.Row_changed_date, facility_description.Row_created_by, facility_description.Row_created_date, facility_description.Row_effective_date, facility_description.Row_expiry_date, facility_description.Row_quality)
	
	if err != nil {
		return err
	}

	return c.JSON(rows)
}

func UpdateFacilityDescription(c *fiber.Ctx) error {
	var facility_description dto.Facility_description

	setDefaults(&facility_description)

	if err := json.Unmarshal(c.Body(), &facility_description); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	id := c.Params("id")
	facility_description.Facility_id = id

    if facility_description.Row_created_date != nil {
        return c.Status(400).SendString("Cannot update row_created_date")
    }

	stmt, err := db.Prepare("update facility_description set facility_type = :1, spec_id = :2, active_ind = :3, average_value = :4, average_value_ouom = :5, average_value_uom = :6, cost = :7, currency_conversion = :8, currency_ouom = :9, currency_uom = :10, date_format_desc = :11, effective_date = :12, expiry_date = :13, max_date = :14, max_value = :15, max_value_ouom = :16, max_value_uom = :17, min_date = :18, min_value = :19, min_value_ouom = :20, min_value_uom = :21, ppdm_guid = :22, reference_value = :23, reference_value_ouom = :24, reference_value_type = :25, reference_value_uom = :26, remark = :27, source = :28, spec_code = :29, spec_desc = :30, spec_type = :31, row_changed_by = :32, row_changed_date = :33, row_created_by = :34, row_effective_date = :35, row_expiry_date = :36, row_quality = :37 where facility_id = :39")
	if err != nil {
		return err
	}

	facility_description.Row_changed_date = formatDate(facility_description.Row_changed_date)
	facility_description.Date_format_desc = formatDateString(facility_description.Date_format_desc)
	facility_description.Effective_date = formatDateString(facility_description.Effective_date)
	facility_description.Expiry_date = formatDateString(facility_description.Expiry_date)
	facility_description.Max_date = formatDateString(facility_description.Max_date)
	facility_description.Min_date = formatDateString(facility_description.Min_date)
	facility_description.Row_effective_date = formatDateString(facility_description.Row_effective_date)
	facility_description.Row_expiry_date = formatDateString(facility_description.Row_expiry_date)






	rows, err := stmt.Exec(facility_description.Facility_type, facility_description.Spec_id, facility_description.Active_ind, facility_description.Average_value, facility_description.Average_value_ouom, facility_description.Average_value_uom, facility_description.Cost, facility_description.Currency_conversion, facility_description.Currency_ouom, facility_description.Currency_uom, facility_description.Date_format_desc, facility_description.Effective_date, facility_description.Expiry_date, facility_description.Max_date, facility_description.Max_value, facility_description.Max_value_ouom, facility_description.Max_value_uom, facility_description.Min_date, facility_description.Min_value, facility_description.Min_value_ouom, facility_description.Min_value_uom, facility_description.Ppdm_guid, facility_description.Reference_value, facility_description.Reference_value_ouom, facility_description.Reference_value_type, facility_description.Reference_value_uom, facility_description.Remark, facility_description.Source, facility_description.Spec_code, facility_description.Spec_desc, facility_description.Spec_type, facility_description.Row_changed_by, facility_description.Row_changed_date, facility_description.Row_created_by, facility_description.Row_effective_date, facility_description.Row_expiry_date, facility_description.Row_quality, facility_description.Facility_id)
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

func PatchFacilityDescription(c *fiber.Ctx) error {
	id := c.Params("id")

	updatedFields := make(map[string]interface{})

	if err := c.BodyParser(&updatedFields); err != nil {
		return err
	}
	if _, ok := updatedFields["row_created_date"]; ok {
        return c.Status(400).SendString("Cannot update row_created_date field")
    }

	query := "update facility_description set "
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
	query += " where facility_id = :id"

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

func DeleteFacilityDescription(c *fiber.Ctx) error {
	id := c.Params("id")
	var facility_description dto.Facility_description
	facility_description.Facility_id = id

	stmt, err := db.Prepare("delete from facility_description where facility_id = :1")
	if err != nil {
		return err
	}

	rows, err := stmt.Exec(facility_description.Facility_id)
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


