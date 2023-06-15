package apiv1

import (
	dto "github.com/DarrenMannuela/pt_gtn_bibliography/dto"
	"github.com/gofiber/fiber/v2"
	"strconv"
	"encoding/json"
)


func GetFacilityAlias(c *fiber.Ctx) error {
	rows, err := db.Query("select * from facility_alias")
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}
	defer rows.Close()
	var result []dto.Facility_alias

	for rows.Next() {
		var facility_alias dto.Facility_alias
		if err := rows.Scan(&facility_alias.Facility_id, &facility_alias.Facility_type, &facility_alias.Facility_alias_id, &facility_alias.Source, &facility_alias.Abbreviation, &facility_alias.Active_ind, &facility_alias.Alias_long_name, &facility_alias.Alias_reason_type, &facility_alias.Alias_short_name, &facility_alias.Alias_type, &facility_alias.Amended_date, &facility_alias.Created_date, &facility_alias.Effective_date, &facility_alias.Expiry_date, &facility_alias.Facility_id, &facility_alias.Owner_ba_id, &facility_alias.Ppdm_guid, &facility_alias.Preferred_ind, &facility_alias.Reason_desc, &facility_alias.Remark, &facility_alias.Source_document_id, &facility_alias.Struckoff_date, &facility_alias.Sw_application_id, &facility_alias.Use_rule, &facility_alias.Row_changed_by, &facility_alias.Row_changed_date, &facility_alias.Row_created_by, &facility_alias.Row_created_date, &facility_alias.Row_effective_date, &facility_alias.Row_expiry_date, &facility_alias.Row_quality); err != nil {
			return err // Exit if we get an error
		}

		// Append facility_alias to result
		result = append(result, facility_alias)
	}
	// Return result in JSON format
	return c.JSON(result)
	 
}

func SetFacilityAlias(c *fiber.Ctx) error {
	var facility_alias dto.Facility_alias

	setDefaults(&facility_alias)

	if err := json.Unmarshal(c.Body(), &facility_alias); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	stmt, err := db.Prepare("insert into facility_alias values(:1, :2, :3, :4, :5, :6, :7, :8, :9, :10, :11, :12, :13, :14, :15, :16, :17, :18, :19, :20, :21, :22, :23, :24, :25, :26, :27, :28, :29, :30, :31)")
	if err != nil {
		return err
	}
	facility_alias.Row_created_date = formatDate(facility_alias.Row_created_date)
	facility_alias.Row_changed_date = formatDate(facility_alias.Row_changed_date)
	facility_alias.Amended_date = formatDateString(facility_alias.Amended_date)
	facility_alias.Created_date = formatDateString(facility_alias.Created_date)
	facility_alias.Effective_date = formatDateString(facility_alias.Effective_date)
	facility_alias.Expiry_date = formatDateString(facility_alias.Expiry_date)
	facility_alias.Struckoff_date = formatDateString(facility_alias.Struckoff_date)
	facility_alias.Row_effective_date = formatDateString(facility_alias.Row_effective_date)
	facility_alias.Row_expiry_date = formatDateString(facility_alias.Row_expiry_date)






	rows, err := stmt.Exec(facility_alias.Facility_id, facility_alias.Facility_type, facility_alias.Facility_alias_id, facility_alias.Source, facility_alias.Abbreviation, facility_alias.Active_ind, facility_alias.Alias_long_name, facility_alias.Alias_reason_type, facility_alias.Alias_short_name, facility_alias.Alias_type, facility_alias.Amended_date, facility_alias.Created_date, facility_alias.Effective_date, facility_alias.Expiry_date, facility_alias.Facility_id, facility_alias.Owner_ba_id, facility_alias.Ppdm_guid, facility_alias.Preferred_ind, facility_alias.Reason_desc, facility_alias.Remark, facility_alias.Source_document_id, facility_alias.Struckoff_date, facility_alias.Sw_application_id, facility_alias.Use_rule, facility_alias.Row_changed_by, facility_alias.Row_changed_date, facility_alias.Row_created_by, facility_alias.Row_created_date, facility_alias.Row_effective_date, facility_alias.Row_expiry_date, facility_alias.Row_quality)
	
	if err != nil {
		return err
	}

	return c.JSON(rows)
}

func UpdateFacilityAlias(c *fiber.Ctx) error {
	var facility_alias dto.Facility_alias

	setDefaults(&facility_alias)

	if err := json.Unmarshal(c.Body(), &facility_alias); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	id := c.Params("id")
	facility_alias.Facility_id = id

    if facility_alias.Row_created_date != nil {
        return c.Status(400).SendString("Cannot update row_created_date")
    }

	stmt, err := db.Prepare("update facility_alias set facility_type = :1, facility_alias_id = :2, source = :3, abbreviation = :4, active_ind = :5, alias_long_name = :6, alias_reason_type = :7, alias_short_name = :8, alias_type = :9, amended_date = :10, created_date = :11, effective_date = :12, expiry_date = :13, facility_id = :14, owner_ba_id = :15, ppdm_guid = :16, preferred_ind = :17, reason_desc = :18, remark = :19, source_document_id = :20, struckoff_date = :21, sw_application_id = :22, use_rule = :23, row_changed_by = :24, row_changed_date = :25, row_created_by = :26, row_effective_date = :27, row_expiry_date = :28, row_quality = :29 where facility_id = :31")
	if err != nil {
		return err
	}

	facility_alias.Row_changed_date = formatDate(facility_alias.Row_changed_date)
	facility_alias.Amended_date = formatDateString(facility_alias.Amended_date)
	facility_alias.Created_date = formatDateString(facility_alias.Created_date)
	facility_alias.Effective_date = formatDateString(facility_alias.Effective_date)
	facility_alias.Expiry_date = formatDateString(facility_alias.Expiry_date)
	facility_alias.Struckoff_date = formatDateString(facility_alias.Struckoff_date)
	facility_alias.Row_effective_date = formatDateString(facility_alias.Row_effective_date)
	facility_alias.Row_expiry_date = formatDateString(facility_alias.Row_expiry_date)






	rows, err := stmt.Exec(facility_alias.Facility_type, facility_alias.Facility_alias_id, facility_alias.Source, facility_alias.Abbreviation, facility_alias.Active_ind, facility_alias.Alias_long_name, facility_alias.Alias_reason_type, facility_alias.Alias_short_name, facility_alias.Alias_type, facility_alias.Amended_date, facility_alias.Created_date, facility_alias.Effective_date, facility_alias.Expiry_date, facility_alias.Facility_id, facility_alias.Owner_ba_id, facility_alias.Ppdm_guid, facility_alias.Preferred_ind, facility_alias.Reason_desc, facility_alias.Remark, facility_alias.Source_document_id, facility_alias.Struckoff_date, facility_alias.Sw_application_id, facility_alias.Use_rule, facility_alias.Row_changed_by, facility_alias.Row_changed_date, facility_alias.Row_created_by, facility_alias.Row_effective_date, facility_alias.Row_expiry_date, facility_alias.Row_quality, facility_alias.Facility_id)
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

func PatchFacilityAlias(c *fiber.Ctx) error {
	id := c.Params("id")

	updatedFields := make(map[string]interface{})

	if err := c.BodyParser(&updatedFields); err != nil {
		return err
	}
	if _, ok := updatedFields["row_created_date"]; ok {
        return c.Status(400).SendString("Cannot update row_created_date field")
    }

	query := "update facility_alias set "
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

func DeleteFacilityAlias(c *fiber.Ctx) error {
	id := c.Params("id")
	var facility_alias dto.Facility_alias
	facility_alias.Facility_id = id

	stmt, err := db.Prepare("delete from facility_alias where facility_id = :1")
	if err != nil {
		return err
	}

	rows, err := stmt.Exec(facility_alias.Facility_id)
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


