package apiv1

import (
	dto "github.com/DarrenMannuela/pt_gtn_bibliography/dto"
	"github.com/gofiber/fiber/v2"
	"strconv"
	"encoding/json"
)


func GetFacilityLicCond(c *fiber.Ctx) error {
	rows, err := db.Query("select * from facility_lic_cond")
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}
	defer rows.Close()
	var result []dto.Facility_lic_cond

	for rows.Next() {
		var facility_lic_cond dto.Facility_lic_cond
		if err := rows.Scan(&facility_lic_cond.Facility_id, &facility_lic_cond.Facility_type, &facility_lic_cond.License_id, &facility_lic_cond.Condition_id, &facility_lic_cond.Active_ind, &facility_lic_cond.Condition_code, &facility_lic_cond.Condition_type, &facility_lic_cond.Condition_value, &facility_lic_cond.Condition_value_uom, &facility_lic_cond.Contact_ba_id, &facility_lic_cond.Date_format_desc, &facility_lic_cond.Description, &facility_lic_cond.Due_condition, &facility_lic_cond.Due_date, &facility_lic_cond.Due_frequency, &facility_lic_cond.Due_term, &facility_lic_cond.Due_term_uom, &facility_lic_cond.Effective_date, &facility_lic_cond.Exempt_ind, &facility_lic_cond.Expiry_date, &facility_lic_cond.Fulfilled_by_ba_id, &facility_lic_cond.Fulfilled_date, &facility_lic_cond.Fulfilled_ind, &facility_lic_cond.Ppdm_guid, &facility_lic_cond.Product_type, &facility_lic_cond.Remark, &facility_lic_cond.Restriction_id, &facility_lic_cond.Restriction_version, &facility_lic_cond.Source, &facility_lic_cond.Row_changed_by, &facility_lic_cond.Row_changed_date, &facility_lic_cond.Row_created_by, &facility_lic_cond.Row_created_date, &facility_lic_cond.Row_effective_date, &facility_lic_cond.Row_expiry_date, &facility_lic_cond.Row_quality); err != nil {
			return err // Exit if we get an error
		}

		// Append facility_lic_cond to result
		result = append(result, facility_lic_cond)
	}
	// Return result in JSON format
	return c.JSON(result)
	 
}

func SetFacilityLicCond(c *fiber.Ctx) error {
	var facility_lic_cond dto.Facility_lic_cond

	setDefaults(&facility_lic_cond)

	if err := json.Unmarshal(c.Body(), &facility_lic_cond); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	stmt, err := db.Prepare("insert into facility_lic_cond values(:1, :2, :3, :4, :5, :6, :7, :8, :9, :10, :11, :12, :13, :14, :15, :16, :17, :18, :19, :20, :21, :22, :23, :24, :25, :26, :27, :28, :29, :30, :31, :32, :33, :34, :35, :36)")
	if err != nil {
		return err
	}
	facility_lic_cond.Row_created_date = formatDate(facility_lic_cond.Row_created_date)
	facility_lic_cond.Row_changed_date = formatDate(facility_lic_cond.Row_changed_date)
	facility_lic_cond.Date_format_desc = formatDateString(facility_lic_cond.Date_format_desc)
	facility_lic_cond.Due_date = formatDateString(facility_lic_cond.Due_date)
	facility_lic_cond.Effective_date = formatDateString(facility_lic_cond.Effective_date)
	facility_lic_cond.Expiry_date = formatDateString(facility_lic_cond.Expiry_date)
	facility_lic_cond.Fulfilled_date = formatDateString(facility_lic_cond.Fulfilled_date)
	facility_lic_cond.Row_effective_date = formatDateString(facility_lic_cond.Row_effective_date)
	facility_lic_cond.Row_expiry_date = formatDateString(facility_lic_cond.Row_expiry_date)






	rows, err := stmt.Exec(facility_lic_cond.Facility_id, facility_lic_cond.Facility_type, facility_lic_cond.License_id, facility_lic_cond.Condition_id, facility_lic_cond.Active_ind, facility_lic_cond.Condition_code, facility_lic_cond.Condition_type, facility_lic_cond.Condition_value, facility_lic_cond.Condition_value_uom, facility_lic_cond.Contact_ba_id, facility_lic_cond.Date_format_desc, facility_lic_cond.Description, facility_lic_cond.Due_condition, facility_lic_cond.Due_date, facility_lic_cond.Due_frequency, facility_lic_cond.Due_term, facility_lic_cond.Due_term_uom, facility_lic_cond.Effective_date, facility_lic_cond.Exempt_ind, facility_lic_cond.Expiry_date, facility_lic_cond.Fulfilled_by_ba_id, facility_lic_cond.Fulfilled_date, facility_lic_cond.Fulfilled_ind, facility_lic_cond.Ppdm_guid, facility_lic_cond.Product_type, facility_lic_cond.Remark, facility_lic_cond.Restriction_id, facility_lic_cond.Restriction_version, facility_lic_cond.Source, facility_lic_cond.Row_changed_by, facility_lic_cond.Row_changed_date, facility_lic_cond.Row_created_by, facility_lic_cond.Row_created_date, facility_lic_cond.Row_effective_date, facility_lic_cond.Row_expiry_date, facility_lic_cond.Row_quality)
	
	if err != nil {
		return err
	}

	return c.JSON(rows)
}

func UpdateFacilityLicCond(c *fiber.Ctx) error {
	var facility_lic_cond dto.Facility_lic_cond

	setDefaults(&facility_lic_cond)

	if err := json.Unmarshal(c.Body(), &facility_lic_cond); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	id := c.Params("id")
	facility_lic_cond.Facility_id = id

    if facility_lic_cond.Row_created_date != nil {
        return c.Status(400).SendString("Cannot update row_created_date")
    }

	stmt, err := db.Prepare("update facility_lic_cond set facility_type = :1, license_id = :2, condition_id = :3, active_ind = :4, condition_code = :5, condition_type = :6, condition_value = :7, condition_value_uom = :8, contact_ba_id = :9, date_format_desc = :10, description = :11, due_condition = :12, due_date = :13, due_frequency = :14, due_term = :15, due_term_uom = :16, effective_date = :17, exempt_ind = :18, expiry_date = :19, fulfilled_by_ba_id = :20, fulfilled_date = :21, fulfilled_ind = :22, ppdm_guid = :23, product_type = :24, remark = :25, restriction_id = :26, restriction_version = :27, source = :28, row_changed_by = :29, row_changed_date = :30, row_created_by = :31, row_effective_date = :32, row_expiry_date = :33, row_quality = :34 where facility_id = :36")
	if err != nil {
		return err
	}

	facility_lic_cond.Row_changed_date = formatDate(facility_lic_cond.Row_changed_date)
	facility_lic_cond.Date_format_desc = formatDateString(facility_lic_cond.Date_format_desc)
	facility_lic_cond.Due_date = formatDateString(facility_lic_cond.Due_date)
	facility_lic_cond.Effective_date = formatDateString(facility_lic_cond.Effective_date)
	facility_lic_cond.Expiry_date = formatDateString(facility_lic_cond.Expiry_date)
	facility_lic_cond.Fulfilled_date = formatDateString(facility_lic_cond.Fulfilled_date)
	facility_lic_cond.Row_effective_date = formatDateString(facility_lic_cond.Row_effective_date)
	facility_lic_cond.Row_expiry_date = formatDateString(facility_lic_cond.Row_expiry_date)






	rows, err := stmt.Exec(facility_lic_cond.Facility_type, facility_lic_cond.License_id, facility_lic_cond.Condition_id, facility_lic_cond.Active_ind, facility_lic_cond.Condition_code, facility_lic_cond.Condition_type, facility_lic_cond.Condition_value, facility_lic_cond.Condition_value_uom, facility_lic_cond.Contact_ba_id, facility_lic_cond.Date_format_desc, facility_lic_cond.Description, facility_lic_cond.Due_condition, facility_lic_cond.Due_date, facility_lic_cond.Due_frequency, facility_lic_cond.Due_term, facility_lic_cond.Due_term_uom, facility_lic_cond.Effective_date, facility_lic_cond.Exempt_ind, facility_lic_cond.Expiry_date, facility_lic_cond.Fulfilled_by_ba_id, facility_lic_cond.Fulfilled_date, facility_lic_cond.Fulfilled_ind, facility_lic_cond.Ppdm_guid, facility_lic_cond.Product_type, facility_lic_cond.Remark, facility_lic_cond.Restriction_id, facility_lic_cond.Restriction_version, facility_lic_cond.Source, facility_lic_cond.Row_changed_by, facility_lic_cond.Row_changed_date, facility_lic_cond.Row_created_by, facility_lic_cond.Row_effective_date, facility_lic_cond.Row_expiry_date, facility_lic_cond.Row_quality, facility_lic_cond.Facility_id)
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

func PatchFacilityLicCond(c *fiber.Ctx) error {
	id := c.Params("id")

	updatedFields := make(map[string]interface{})

	if err := c.BodyParser(&updatedFields); err != nil {
		return err
	}
	if _, ok := updatedFields["row_created_date"]; ok {
        return c.Status(400).SendString("Cannot update row_created_date field")
    }

	query := "update facility_lic_cond set "
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
		} else if key == "date_format_desc" || key == "due_date" || key == "effective_date" || key == "expiry_date" || key == "fulfilled_date" || key == "row_effective_date" || key == "row_expiry_date"      {
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

func DeleteFacilityLicCond(c *fiber.Ctx) error {
	id := c.Params("id")
	var facility_lic_cond dto.Facility_lic_cond
	facility_lic_cond.Facility_id = id

	stmt, err := db.Prepare("delete from facility_lic_cond where facility_id = :1")
	if err != nil {
		return err
	}

	rows, err := stmt.Exec(facility_lic_cond.Facility_id)
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


