package apiv1

import (
	dto "github.com/DarrenMannuela/pt_gtn_bibliography/dto"
	"github.com/gofiber/fiber/v2"
	"strconv"
	"encoding/json"
)


func GetFacilityLicense(c *fiber.Ctx) error {
	rows, err := db.Query("select * from facility_license")
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}
	defer rows.Close()
	var result []dto.Facility_license

	for rows.Next() {
		var facility_license dto.Facility_license
		if err := rows.Scan(&facility_license.Facility_id, &facility_license.Facility_type, &facility_license.License_id, &facility_license.Active_ind, &facility_license.Adjust_13_month_desc, &facility_license.Adjust_13_month_ind, &facility_license.Application_id, &facility_license.Approved_facility_class, &facility_license.Assigned_field_id, &facility_license.Description, &facility_license.Effective_date, &facility_license.Expiry_date, &facility_license.Facility_license_type, &facility_license.Fees_paid_ind, &facility_license.Granted_by_ba_id, &facility_license.Granted_by_contact_id, &facility_license.Granted_date, &facility_license.Granted_to_ba_id, &facility_license.Granted_to_contact_id, &facility_license.License_extension_cond, &facility_license.License_latitude, &facility_license.License_location, &facility_license.License_longitude, &facility_license.License_num, &facility_license.License_term, &facility_license.License_term_uom, &facility_license.Ppdm_guid, &facility_license.Rate_schedule_id, &facility_license.Rel_license_id, &facility_license.Remark, &facility_license.Renewal_condition, &facility_license.Secondary_term, &facility_license.Secondary_term_uom, &facility_license.Source, &facility_license.Violation_ind, &facility_license.Row_changed_by, &facility_license.Row_changed_date, &facility_license.Row_created_by, &facility_license.Row_created_date, &facility_license.Row_effective_date, &facility_license.Row_expiry_date, &facility_license.Row_quality); err != nil {
			return err // Exit if we get an error
		}

		// Append facility_license to result
		result = append(result, facility_license)
	}
	// Return result in JSON format
	return c.JSON(result)
	 
}

func SetFacilityLicense(c *fiber.Ctx) error {
	var facility_license dto.Facility_license

	setDefaults(&facility_license)

	if err := json.Unmarshal(c.Body(), &facility_license); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	stmt, err := db.Prepare("insert into facility_license values(:1, :2, :3, :4, :5, :6, :7, :8, :9, :10, :11, :12, :13, :14, :15, :16, :17, :18, :19, :20, :21, :22, :23, :24, :25, :26, :27, :28, :29, :30, :31, :32, :33, :34, :35, :36, :37, :38, :39, :40, :41, :42)")
	if err != nil {
		return err
	}
	facility_license.Row_created_date = formatDate(facility_license.Row_created_date)
	facility_license.Row_changed_date = formatDate(facility_license.Row_changed_date)
	facility_license.Effective_date = formatDateString(facility_license.Effective_date)
	facility_license.Expiry_date = formatDateString(facility_license.Expiry_date)
	facility_license.Granted_date = formatDateString(facility_license.Granted_date)
	facility_license.Row_effective_date = formatDateString(facility_license.Row_effective_date)
	facility_license.Row_expiry_date = formatDateString(facility_license.Row_expiry_date)






	rows, err := stmt.Exec(facility_license.Facility_id, facility_license.Facility_type, facility_license.License_id, facility_license.Active_ind, facility_license.Adjust_13_month_desc, facility_license.Adjust_13_month_ind, facility_license.Application_id, facility_license.Approved_facility_class, facility_license.Assigned_field_id, facility_license.Description, facility_license.Effective_date, facility_license.Expiry_date, facility_license.Facility_license_type, facility_license.Fees_paid_ind, facility_license.Granted_by_ba_id, facility_license.Granted_by_contact_id, facility_license.Granted_date, facility_license.Granted_to_ba_id, facility_license.Granted_to_contact_id, facility_license.License_extension_cond, facility_license.License_latitude, facility_license.License_location, facility_license.License_longitude, facility_license.License_num, facility_license.License_term, facility_license.License_term_uom, facility_license.Ppdm_guid, facility_license.Rate_schedule_id, facility_license.Rel_license_id, facility_license.Remark, facility_license.Renewal_condition, facility_license.Secondary_term, facility_license.Secondary_term_uom, facility_license.Source, facility_license.Violation_ind, facility_license.Row_changed_by, facility_license.Row_changed_date, facility_license.Row_created_by, facility_license.Row_created_date, facility_license.Row_effective_date, facility_license.Row_expiry_date, facility_license.Row_quality)
	
	if err != nil {
		return err
	}

	return c.JSON(rows)
}

func UpdateFacilityLicense(c *fiber.Ctx) error {
	var facility_license dto.Facility_license

	setDefaults(&facility_license)

	if err := json.Unmarshal(c.Body(), &facility_license); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	id := c.Params("id")
	facility_license.Facility_id = id

    if facility_license.Row_created_date != nil {
        return c.Status(400).SendString("Cannot update row_created_date")
    }

	stmt, err := db.Prepare("update facility_license set facility_type = :1, license_id = :2, active_ind = :3, adjust_13_month_desc = :4, adjust_13_month_ind = :5, application_id = :6, approved_facility_class = :7, assigned_field_id = :8, description = :9, effective_date = :10, expiry_date = :11, facility_license_type = :12, fees_paid_ind = :13, granted_by_ba_id = :14, granted_by_contact_id = :15, granted_date = :16, granted_to_ba_id = :17, granted_to_contact_id = :18, license_extension_cond = :19, license_latitude = :20, license_location = :21, license_longitude = :22, license_num = :23, license_term = :24, license_term_uom = :25, ppdm_guid = :26, rate_schedule_id = :27, rel_license_id = :28, remark = :29, renewal_condition = :30, secondary_term = :31, secondary_term_uom = :32, source = :33, violation_ind = :34, row_changed_by = :35, row_changed_date = :36, row_created_by = :37, row_effective_date = :38, row_expiry_date = :39, row_quality = :40 where facility_id = :42")
	if err != nil {
		return err
	}

	facility_license.Row_changed_date = formatDate(facility_license.Row_changed_date)
	facility_license.Effective_date = formatDateString(facility_license.Effective_date)
	facility_license.Expiry_date = formatDateString(facility_license.Expiry_date)
	facility_license.Granted_date = formatDateString(facility_license.Granted_date)
	facility_license.Row_effective_date = formatDateString(facility_license.Row_effective_date)
	facility_license.Row_expiry_date = formatDateString(facility_license.Row_expiry_date)






	rows, err := stmt.Exec(facility_license.Facility_type, facility_license.License_id, facility_license.Active_ind, facility_license.Adjust_13_month_desc, facility_license.Adjust_13_month_ind, facility_license.Application_id, facility_license.Approved_facility_class, facility_license.Assigned_field_id, facility_license.Description, facility_license.Effective_date, facility_license.Expiry_date, facility_license.Facility_license_type, facility_license.Fees_paid_ind, facility_license.Granted_by_ba_id, facility_license.Granted_by_contact_id, facility_license.Granted_date, facility_license.Granted_to_ba_id, facility_license.Granted_to_contact_id, facility_license.License_extension_cond, facility_license.License_latitude, facility_license.License_location, facility_license.License_longitude, facility_license.License_num, facility_license.License_term, facility_license.License_term_uom, facility_license.Ppdm_guid, facility_license.Rate_schedule_id, facility_license.Rel_license_id, facility_license.Remark, facility_license.Renewal_condition, facility_license.Secondary_term, facility_license.Secondary_term_uom, facility_license.Source, facility_license.Violation_ind, facility_license.Row_changed_by, facility_license.Row_changed_date, facility_license.Row_created_by, facility_license.Row_effective_date, facility_license.Row_expiry_date, facility_license.Row_quality, facility_license.Facility_id)
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

func PatchFacilityLicense(c *fiber.Ctx) error {
	id := c.Params("id")

	updatedFields := make(map[string]interface{})

	if err := c.BodyParser(&updatedFields); err != nil {
		return err
	}
	if _, ok := updatedFields["row_created_date"]; ok {
        return c.Status(400).SendString("Cannot update row_created_date field")
    }

	query := "update facility_license set "
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
		} else if key == "effective_date" || key == "expiry_date" || key == "granted_date" || key == "row_effective_date" || key == "row_expiry_date"      {
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

func DeleteFacilityLicense(c *fiber.Ctx) error {
	id := c.Params("id")
	var facility_license dto.Facility_license
	facility_license.Facility_id = id

	stmt, err := db.Prepare("delete from facility_license where facility_id = :1")
	if err != nil {
		return err
	}

	rows, err := stmt.Exec(facility_license.Facility_id)
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


