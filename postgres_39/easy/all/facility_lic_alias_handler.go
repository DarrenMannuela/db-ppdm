package apiv1

import (
	dto "github.com/DarrenMannuela/pt_gtn_bibliography/dto"
	"github.com/gofiber/fiber/v2"
	"strconv"
	"encoding/json"
)


func GetFacilityLicAlias(c *fiber.Ctx) error {
	rows, err := db.Query("select * from facility_lic_alias")
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}
	defer rows.Close()
	var result []dto.Facility_lic_alias

	for rows.Next() {
		var facility_lic_alias dto.Facility_lic_alias
		if err := rows.Scan(&facility_lic_alias.Facility_id, &facility_lic_alias.Facility_type, &facility_lic_alias.License_id, &facility_lic_alias.Alias_id, &facility_lic_alias.Abbreviation, &facility_lic_alias.Active_ind, &facility_lic_alias.Alias_long_name, &facility_lic_alias.Alias_reason_type, &facility_lic_alias.Alias_short_name, &facility_lic_alias.Alias_type, &facility_lic_alias.Amended_date, &facility_lic_alias.Created_date, &facility_lic_alias.Effective_date, &facility_lic_alias.Expiry_date, &facility_lic_alias.Facility_id, &facility_lic_alias.Original_ind, &facility_lic_alias.Owner_ba_id, &facility_lic_alias.Ppdm_guid, &facility_lic_alias.Preferred_ind, &facility_lic_alias.Reason_desc, &facility_lic_alias.Remark, &facility_lic_alias.Source, &facility_lic_alias.Source_document_id, &facility_lic_alias.Struckoff_date, &facility_lic_alias.Sw_application_id, &facility_lic_alias.Use_rule, &facility_lic_alias.Row_changed_by, &facility_lic_alias.Row_changed_date, &facility_lic_alias.Row_created_by, &facility_lic_alias.Row_created_date, &facility_lic_alias.Row_effective_date, &facility_lic_alias.Row_expiry_date, &facility_lic_alias.Row_quality); err != nil {
			return err // Exit if we get an error
		}

		// Append facility_lic_alias to result
		result = append(result, facility_lic_alias)
	}
	// Return result in JSON format
	return c.JSON(result)
	 
}

func SetFacilityLicAlias(c *fiber.Ctx) error {
	var facility_lic_alias dto.Facility_lic_alias

	setDefaults(&facility_lic_alias)

	if err := json.Unmarshal(c.Body(), &facility_lic_alias); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	stmt, err := db.Prepare("insert into facility_lic_alias values(:1, :2, :3, :4, :5, :6, :7, :8, :9, :10, :11, :12, :13, :14, :15, :16, :17, :18, :19, :20, :21, :22, :23, :24, :25, :26, :27, :28, :29, :30, :31, :32, :33)")
	if err != nil {
		return err
	}
	facility_lic_alias.Row_created_date = formatDate(facility_lic_alias.Row_created_date)
	facility_lic_alias.Row_changed_date = formatDate(facility_lic_alias.Row_changed_date)
	facility_lic_alias.Amended_date = formatDateString(facility_lic_alias.Amended_date)
	facility_lic_alias.Created_date = formatDateString(facility_lic_alias.Created_date)
	facility_lic_alias.Effective_date = formatDateString(facility_lic_alias.Effective_date)
	facility_lic_alias.Expiry_date = formatDateString(facility_lic_alias.Expiry_date)
	facility_lic_alias.Struckoff_date = formatDateString(facility_lic_alias.Struckoff_date)
	facility_lic_alias.Row_effective_date = formatDateString(facility_lic_alias.Row_effective_date)
	facility_lic_alias.Row_expiry_date = formatDateString(facility_lic_alias.Row_expiry_date)






	rows, err := stmt.Exec(facility_lic_alias.Facility_id, facility_lic_alias.Facility_type, facility_lic_alias.License_id, facility_lic_alias.Alias_id, facility_lic_alias.Abbreviation, facility_lic_alias.Active_ind, facility_lic_alias.Alias_long_name, facility_lic_alias.Alias_reason_type, facility_lic_alias.Alias_short_name, facility_lic_alias.Alias_type, facility_lic_alias.Amended_date, facility_lic_alias.Created_date, facility_lic_alias.Effective_date, facility_lic_alias.Expiry_date, facility_lic_alias.Facility_id, facility_lic_alias.Original_ind, facility_lic_alias.Owner_ba_id, facility_lic_alias.Ppdm_guid, facility_lic_alias.Preferred_ind, facility_lic_alias.Reason_desc, facility_lic_alias.Remark, facility_lic_alias.Source, facility_lic_alias.Source_document_id, facility_lic_alias.Struckoff_date, facility_lic_alias.Sw_application_id, facility_lic_alias.Use_rule, facility_lic_alias.Row_changed_by, facility_lic_alias.Row_changed_date, facility_lic_alias.Row_created_by, facility_lic_alias.Row_created_date, facility_lic_alias.Row_effective_date, facility_lic_alias.Row_expiry_date, facility_lic_alias.Row_quality)
	
	if err != nil {
		return err
	}

	return c.JSON(rows)
}

func UpdateFacilityLicAlias(c *fiber.Ctx) error {
	var facility_lic_alias dto.Facility_lic_alias

	setDefaults(&facility_lic_alias)

	if err := json.Unmarshal(c.Body(), &facility_lic_alias); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	id := c.Params("id")
	facility_lic_alias.Facility_id = id

    if facility_lic_alias.Row_created_date != nil {
        return c.Status(400).SendString("Cannot update row_created_date")
    }

	stmt, err := db.Prepare("update facility_lic_alias set facility_type = :1, license_id = :2, alias_id = :3, abbreviation = :4, active_ind = :5, alias_long_name = :6, alias_reason_type = :7, alias_short_name = :8, alias_type = :9, amended_date = :10, created_date = :11, effective_date = :12, expiry_date = :13, facility_id = :14, original_ind = :15, owner_ba_id = :16, ppdm_guid = :17, preferred_ind = :18, reason_desc = :19, remark = :20, source = :21, source_document_id = :22, struckoff_date = :23, sw_application_id = :24, use_rule = :25, row_changed_by = :26, row_changed_date = :27, row_created_by = :28, row_effective_date = :29, row_expiry_date = :30, row_quality = :31 where facility_id = :33")
	if err != nil {
		return err
	}

	facility_lic_alias.Row_changed_date = formatDate(facility_lic_alias.Row_changed_date)
	facility_lic_alias.Amended_date = formatDateString(facility_lic_alias.Amended_date)
	facility_lic_alias.Created_date = formatDateString(facility_lic_alias.Created_date)
	facility_lic_alias.Effective_date = formatDateString(facility_lic_alias.Effective_date)
	facility_lic_alias.Expiry_date = formatDateString(facility_lic_alias.Expiry_date)
	facility_lic_alias.Struckoff_date = formatDateString(facility_lic_alias.Struckoff_date)
	facility_lic_alias.Row_effective_date = formatDateString(facility_lic_alias.Row_effective_date)
	facility_lic_alias.Row_expiry_date = formatDateString(facility_lic_alias.Row_expiry_date)






	rows, err := stmt.Exec(facility_lic_alias.Facility_type, facility_lic_alias.License_id, facility_lic_alias.Alias_id, facility_lic_alias.Abbreviation, facility_lic_alias.Active_ind, facility_lic_alias.Alias_long_name, facility_lic_alias.Alias_reason_type, facility_lic_alias.Alias_short_name, facility_lic_alias.Alias_type, facility_lic_alias.Amended_date, facility_lic_alias.Created_date, facility_lic_alias.Effective_date, facility_lic_alias.Expiry_date, facility_lic_alias.Facility_id, facility_lic_alias.Original_ind, facility_lic_alias.Owner_ba_id, facility_lic_alias.Ppdm_guid, facility_lic_alias.Preferred_ind, facility_lic_alias.Reason_desc, facility_lic_alias.Remark, facility_lic_alias.Source, facility_lic_alias.Source_document_id, facility_lic_alias.Struckoff_date, facility_lic_alias.Sw_application_id, facility_lic_alias.Use_rule, facility_lic_alias.Row_changed_by, facility_lic_alias.Row_changed_date, facility_lic_alias.Row_created_by, facility_lic_alias.Row_effective_date, facility_lic_alias.Row_expiry_date, facility_lic_alias.Row_quality, facility_lic_alias.Facility_id)
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

func PatchFacilityLicAlias(c *fiber.Ctx) error {
	id := c.Params("id")

	updatedFields := make(map[string]interface{})

	if err := c.BodyParser(&updatedFields); err != nil {
		return err
	}
	if _, ok := updatedFields["row_created_date"]; ok {
        return c.Status(400).SendString("Cannot update row_created_date field")
    }

	query := "update facility_lic_alias set "
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

func DeleteFacilityLicAlias(c *fiber.Ctx) error {
	id := c.Params("id")
	var facility_lic_alias dto.Facility_lic_alias
	facility_lic_alias.Facility_id = id

	stmt, err := db.Prepare("delete from facility_lic_alias where facility_id = :1")
	if err != nil {
		return err
	}

	rows, err := stmt.Exec(facility_lic_alias.Facility_id)
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


