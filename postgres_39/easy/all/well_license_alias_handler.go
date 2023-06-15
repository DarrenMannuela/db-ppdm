package apiv1

import (
	dto "github.com/DarrenMannuela/pt_gtn_bibliography/dto"
	"github.com/gofiber/fiber/v2"
	"strconv"
	"encoding/json"
)


func GetWellLicenseAlias(c *fiber.Ctx) error {
	rows, err := db.Query("select * from well_license_alias")
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}
	defer rows.Close()
	var result []dto.Well_license_alias

	for rows.Next() {
		var well_license_alias dto.Well_license_alias
		if err := rows.Scan(&well_license_alias.Uwi, &well_license_alias.Well_lic_source, &well_license_alias.License_id, &well_license_alias.Alias_source, &well_license_alias.Alias_id, &well_license_alias.Abbreviation, &well_license_alias.Active_ind, &well_license_alias.Alias_long_name, &well_license_alias.Alias_reason_type, &well_license_alias.Alias_short_name, &well_license_alias.Alias_type, &well_license_alias.Amended_date, &well_license_alias.Created_date, &well_license_alias.Effective_date, &well_license_alias.Expiry_date, &well_license_alias.Uwi, &well_license_alias.Original_ind, &well_license_alias.Owner_ba_id, &well_license_alias.Ppdm_guid, &well_license_alias.Preferred_ind, &well_license_alias.Reason_desc, &well_license_alias.Remark, &well_license_alias.Source_document_id, &well_license_alias.Struckoff_date, &well_license_alias.Sw_application_id, &well_license_alias.Use_rule, &well_license_alias.Row_changed_by, &well_license_alias.Row_changed_date, &well_license_alias.Row_created_by, &well_license_alias.Row_created_date, &well_license_alias.Row_effective_date, &well_license_alias.Row_expiry_date, &well_license_alias.Row_quality); err != nil {
			return err // Exit if we get an error
		}

		// Append well_license_alias to result
		result = append(result, well_license_alias)
	}
	// Return result in JSON format
	return c.JSON(result)
	 
}

func SetWellLicenseAlias(c *fiber.Ctx) error {
	var well_license_alias dto.Well_license_alias

	setDefaults(&well_license_alias)

	if err := json.Unmarshal(c.Body(), &well_license_alias); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	stmt, err := db.Prepare("insert into well_license_alias values(:1, :2, :3, :4, :5, :6, :7, :8, :9, :10, :11, :12, :13, :14, :15, :16, :17, :18, :19, :20, :21, :22, :23, :24, :25, :26, :27, :28, :29, :30, :31, :32, :33)")
	if err != nil {
		return err
	}
	well_license_alias.Row_created_date = formatDate(well_license_alias.Row_created_date)
	well_license_alias.Row_changed_date = formatDate(well_license_alias.Row_changed_date)
	well_license_alias.Amended_date = formatDateString(well_license_alias.Amended_date)
	well_license_alias.Created_date = formatDateString(well_license_alias.Created_date)
	well_license_alias.Effective_date = formatDateString(well_license_alias.Effective_date)
	well_license_alias.Expiry_date = formatDateString(well_license_alias.Expiry_date)
	well_license_alias.Struckoff_date = formatDateString(well_license_alias.Struckoff_date)
	well_license_alias.Row_effective_date = formatDateString(well_license_alias.Row_effective_date)
	well_license_alias.Row_expiry_date = formatDateString(well_license_alias.Row_expiry_date)






	rows, err := stmt.Exec(well_license_alias.Uwi, well_license_alias.Well_lic_source, well_license_alias.License_id, well_license_alias.Alias_source, well_license_alias.Alias_id, well_license_alias.Abbreviation, well_license_alias.Active_ind, well_license_alias.Alias_long_name, well_license_alias.Alias_reason_type, well_license_alias.Alias_short_name, well_license_alias.Alias_type, well_license_alias.Amended_date, well_license_alias.Created_date, well_license_alias.Effective_date, well_license_alias.Expiry_date, well_license_alias.Uwi, well_license_alias.Original_ind, well_license_alias.Owner_ba_id, well_license_alias.Ppdm_guid, well_license_alias.Preferred_ind, well_license_alias.Reason_desc, well_license_alias.Remark, well_license_alias.Source_document_id, well_license_alias.Struckoff_date, well_license_alias.Sw_application_id, well_license_alias.Use_rule, well_license_alias.Row_changed_by, well_license_alias.Row_changed_date, well_license_alias.Row_created_by, well_license_alias.Row_created_date, well_license_alias.Row_effective_date, well_license_alias.Row_expiry_date, well_license_alias.Row_quality)
	
	if err != nil {
		return err
	}

	return c.JSON(rows)
}

func UpdateWellLicenseAlias(c *fiber.Ctx) error {
	var well_license_alias dto.Well_license_alias

	setDefaults(&well_license_alias)

	if err := json.Unmarshal(c.Body(), &well_license_alias); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	id := c.Params("id")
	well_license_alias.Uwi = id

    if well_license_alias.Row_created_date != nil {
        return c.Status(400).SendString("Cannot update row_created_date")
    }

	stmt, err := db.Prepare("update well_license_alias set well_lic_source = :1, license_id = :2, alias_source = :3, alias_id = :4, abbreviation = :5, active_ind = :6, alias_long_name = :7, alias_reason_type = :8, alias_short_name = :9, alias_type = :10, amended_date = :11, created_date = :12, effective_date = :13, expiry_date = :14, uwi = :15, original_ind = :16, owner_ba_id = :17, ppdm_guid = :18, preferred_ind = :19, reason_desc = :20, remark = :21, source_document_id = :22, struckoff_date = :23, sw_application_id = :24, use_rule = :25, row_changed_by = :26, row_changed_date = :27, row_created_by = :28, row_effective_date = :29, row_expiry_date = :30, row_quality = :31 where uwi = :33")
	if err != nil {
		return err
	}

	well_license_alias.Row_changed_date = formatDate(well_license_alias.Row_changed_date)
	well_license_alias.Amended_date = formatDateString(well_license_alias.Amended_date)
	well_license_alias.Created_date = formatDateString(well_license_alias.Created_date)
	well_license_alias.Effective_date = formatDateString(well_license_alias.Effective_date)
	well_license_alias.Expiry_date = formatDateString(well_license_alias.Expiry_date)
	well_license_alias.Struckoff_date = formatDateString(well_license_alias.Struckoff_date)
	well_license_alias.Row_effective_date = formatDateString(well_license_alias.Row_effective_date)
	well_license_alias.Row_expiry_date = formatDateString(well_license_alias.Row_expiry_date)






	rows, err := stmt.Exec(well_license_alias.Well_lic_source, well_license_alias.License_id, well_license_alias.Alias_source, well_license_alias.Alias_id, well_license_alias.Abbreviation, well_license_alias.Active_ind, well_license_alias.Alias_long_name, well_license_alias.Alias_reason_type, well_license_alias.Alias_short_name, well_license_alias.Alias_type, well_license_alias.Amended_date, well_license_alias.Created_date, well_license_alias.Effective_date, well_license_alias.Expiry_date, well_license_alias.Uwi, well_license_alias.Original_ind, well_license_alias.Owner_ba_id, well_license_alias.Ppdm_guid, well_license_alias.Preferred_ind, well_license_alias.Reason_desc, well_license_alias.Remark, well_license_alias.Source_document_id, well_license_alias.Struckoff_date, well_license_alias.Sw_application_id, well_license_alias.Use_rule, well_license_alias.Row_changed_by, well_license_alias.Row_changed_date, well_license_alias.Row_created_by, well_license_alias.Row_effective_date, well_license_alias.Row_expiry_date, well_license_alias.Row_quality, well_license_alias.Uwi)
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

func PatchWellLicenseAlias(c *fiber.Ctx) error {
	id := c.Params("id")

	updatedFields := make(map[string]interface{})

	if err := c.BodyParser(&updatedFields); err != nil {
		return err
	}
	if _, ok := updatedFields["row_created_date"]; ok {
        return c.Status(400).SendString("Cannot update row_created_date field")
    }

	query := "update well_license_alias set "
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
		} else if key == "amended_date" || key == "created_date" || key == "effective_date" || key == "expiry_date" || key == "struckoff_date" || key == "row_effective_date" || key == "row_expiry_date"      {
			if date, ok := value.(string); ok && len(date) > 0 {
				formattedDate := formatDateString(&date)
				value = formattedDate
			}
		}
		values = append(values, value)
	}
	query += " where uwi = :id"

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

func DeleteWellLicenseAlias(c *fiber.Ctx) error {
	id := c.Params("id")
	var well_license_alias dto.Well_license_alias
	well_license_alias.Uwi = id

	stmt, err := db.Prepare("delete from well_license_alias where uwi = :1")
	if err != nil {
		return err
	}

	rows, err := stmt.Exec(well_license_alias.Uwi)
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


