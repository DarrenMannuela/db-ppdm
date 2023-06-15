package apiv1

import (
	dto "github.com/DarrenMannuela/pt_gtn_bibliography/dto"
	"github.com/gofiber/fiber/v2"
	"strconv"
	"encoding/json"
)


func GetSeisLicense(c *fiber.Ctx) error {
	rows, err := db.Query("select * from seis_license")
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}
	defer rows.Close()
	var result []dto.Seis_license

	for rows.Next() {
		var seis_license dto.Seis_license
		if err := rows.Scan(&seis_license.Seis_set_subtype, &seis_license.Seis_set_id, &seis_license.License_id, &seis_license.Active_ind, &seis_license.Application_id, &seis_license.Description, &seis_license.Effective_date, &seis_license.Expiry_date, &seis_license.Fees_paid_ind, &seis_license.Granted_by_ba_id, &seis_license.Granted_by_contact_id, &seis_license.Granted_date, &seis_license.Granted_to_ba_id, &seis_license.Granted_to_contact_id, &seis_license.License_num, &seis_license.License_term, &seis_license.License_term_uom, &seis_license.Ppdm_guid, &seis_license.Rate_schedule_id, &seis_license.Rel_license_id, &seis_license.Rel_seis_set_id, &seis_license.Rel_seis_set_subtype, &seis_license.Remark, &seis_license.Secondary_term, &seis_license.Secondary_term_uom, &seis_license.Seis_license_type_id, &seis_license.Source, &seis_license.Violation_ind, &seis_license.Row_changed_by, &seis_license.Row_changed_date, &seis_license.Row_created_by, &seis_license.Row_created_date, &seis_license.Row_effective_date, &seis_license.Row_expiry_date, &seis_license.Row_quality); err != nil {
			return err // Exit if we get an error
		}

		// Append seis_license to result
		result = append(result, seis_license)
	}
	// Return result in JSON format
	return c.JSON(result)
	 
}

func SetSeisLicense(c *fiber.Ctx) error {
	var seis_license dto.Seis_license

	setDefaults(&seis_license)

	if err := json.Unmarshal(c.Body(), &seis_license); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	stmt, err := db.Prepare("insert into seis_license values(:1, :2, :3, :4, :5, :6, :7, :8, :9, :10, :11, :12, :13, :14, :15, :16, :17, :18, :19, :20, :21, :22, :23, :24, :25, :26, :27, :28, :29, :30, :31, :32, :33, :34, :35)")
	if err != nil {
		return err
	}
	seis_license.Row_created_date = formatDate(seis_license.Row_created_date)
	seis_license.Row_changed_date = formatDate(seis_license.Row_changed_date)
	seis_license.Effective_date = formatDateString(seis_license.Effective_date)
	seis_license.Expiry_date = formatDateString(seis_license.Expiry_date)
	seis_license.Granted_date = formatDateString(seis_license.Granted_date)
	seis_license.Row_effective_date = formatDateString(seis_license.Row_effective_date)
	seis_license.Row_expiry_date = formatDateString(seis_license.Row_expiry_date)






	rows, err := stmt.Exec(seis_license.Seis_set_subtype, seis_license.Seis_set_id, seis_license.License_id, seis_license.Active_ind, seis_license.Application_id, seis_license.Description, seis_license.Effective_date, seis_license.Expiry_date, seis_license.Fees_paid_ind, seis_license.Granted_by_ba_id, seis_license.Granted_by_contact_id, seis_license.Granted_date, seis_license.Granted_to_ba_id, seis_license.Granted_to_contact_id, seis_license.License_num, seis_license.License_term, seis_license.License_term_uom, seis_license.Ppdm_guid, seis_license.Rate_schedule_id, seis_license.Rel_license_id, seis_license.Rel_seis_set_id, seis_license.Rel_seis_set_subtype, seis_license.Remark, seis_license.Secondary_term, seis_license.Secondary_term_uom, seis_license.Seis_license_type_id, seis_license.Source, seis_license.Violation_ind, seis_license.Row_changed_by, seis_license.Row_changed_date, seis_license.Row_created_by, seis_license.Row_created_date, seis_license.Row_effective_date, seis_license.Row_expiry_date, seis_license.Row_quality)
	
	if err != nil {
		return err
	}

	return c.JSON(rows)
}

func UpdateSeisLicense(c *fiber.Ctx) error {
	var seis_license dto.Seis_license

	setDefaults(&seis_license)

	if err := json.Unmarshal(c.Body(), &seis_license); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	id := c.Params("id")
	seis_license.Seis_set_subtype = id

    if seis_license.Row_created_date != nil {
        return c.Status(400).SendString("Cannot update row_created_date")
    }

	stmt, err := db.Prepare("update seis_license set seis_set_id = :1, license_id = :2, active_ind = :3, application_id = :4, description = :5, effective_date = :6, expiry_date = :7, fees_paid_ind = :8, granted_by_ba_id = :9, granted_by_contact_id = :10, granted_date = :11, granted_to_ba_id = :12, granted_to_contact_id = :13, license_num = :14, license_term = :15, license_term_uom = :16, ppdm_guid = :17, rate_schedule_id = :18, rel_license_id = :19, rel_seis_set_id = :20, rel_seis_set_subtype = :21, remark = :22, secondary_term = :23, secondary_term_uom = :24, seis_license_type_id = :25, source = :26, violation_ind = :27, row_changed_by = :28, row_changed_date = :29, row_created_by = :30, row_effective_date = :31, row_expiry_date = :32, row_quality = :33 where seis_set_subtype = :35")
	if err != nil {
		return err
	}

	seis_license.Row_changed_date = formatDate(seis_license.Row_changed_date)
	seis_license.Effective_date = formatDateString(seis_license.Effective_date)
	seis_license.Expiry_date = formatDateString(seis_license.Expiry_date)
	seis_license.Granted_date = formatDateString(seis_license.Granted_date)
	seis_license.Row_effective_date = formatDateString(seis_license.Row_effective_date)
	seis_license.Row_expiry_date = formatDateString(seis_license.Row_expiry_date)






	rows, err := stmt.Exec(seis_license.Seis_set_id, seis_license.License_id, seis_license.Active_ind, seis_license.Application_id, seis_license.Description, seis_license.Effective_date, seis_license.Expiry_date, seis_license.Fees_paid_ind, seis_license.Granted_by_ba_id, seis_license.Granted_by_contact_id, seis_license.Granted_date, seis_license.Granted_to_ba_id, seis_license.Granted_to_contact_id, seis_license.License_num, seis_license.License_term, seis_license.License_term_uom, seis_license.Ppdm_guid, seis_license.Rate_schedule_id, seis_license.Rel_license_id, seis_license.Rel_seis_set_id, seis_license.Rel_seis_set_subtype, seis_license.Remark, seis_license.Secondary_term, seis_license.Secondary_term_uom, seis_license.Seis_license_type_id, seis_license.Source, seis_license.Violation_ind, seis_license.Row_changed_by, seis_license.Row_changed_date, seis_license.Row_created_by, seis_license.Row_effective_date, seis_license.Row_expiry_date, seis_license.Row_quality, seis_license.Seis_set_subtype)
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

func PatchSeisLicense(c *fiber.Ctx) error {
	id := c.Params("id")

	updatedFields := make(map[string]interface{})

	if err := c.BodyParser(&updatedFields); err != nil {
		return err
	}
	if _, ok := updatedFields["row_created_date"]; ok {
        return c.Status(400).SendString("Cannot update row_created_date field")
    }

	query := "update seis_license set "
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
	query += " where seis_set_subtype = :id"

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

func DeleteSeisLicense(c *fiber.Ctx) error {
	id := c.Params("id")
	var seis_license dto.Seis_license
	seis_license.Seis_set_subtype = id

	stmt, err := db.Prepare("delete from seis_license where seis_set_subtype = :1")
	if err != nil {
		return err
	}

	rows, err := stmt.Exec(seis_license.Seis_set_subtype)
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


